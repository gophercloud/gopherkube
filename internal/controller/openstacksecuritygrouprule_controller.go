/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"errors"
	"fmt"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/cluster-api/util/patch"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/gophercloud/gophercloud"
	openstackv1 "github.com/gophercloud/gopherkube/api/v1alpha1"
	"github.com/gophercloud/gopherkube/pkg/cloud"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/rules"
)

// OpenStackSecurityGroupRuleReconciler reconciles a OpenStackSecurityGroupRule object
type OpenStackSecurityGroupRuleReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=gopherkube.dev,resources=openstacksecuritygrouprules,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=gopherkube.dev,resources=openstacksecuritygrouprules/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=gopherkube.dev,resources=openstacksecuritygrouprules/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the OpenStackSecurityGroupRule object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.0/pkg/reconcile
func (r *OpenStackSecurityGroupRuleReconciler) Reconcile(ctx context.Context, req ctrl.Request) (_ ctrl.Result, reterr error) {
	logger := log.FromContext(ctx)
	logger = logger.WithValues("OpenStackSecurityGroupRule", req.Name)

	resource := &openstackv1.OpenStackSecurityGroupRule{}
	err := r.Client.Get(ctx, req.NamespacedName, resource)
	if err != nil {
		if apierrors.IsNotFound(err) {
			logger.Info("resource not found in the API")
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	openStackCloud := &openstackv1.OpenStackCloud{}
	if err := r.Client.Get(ctx, types.NamespacedName{
		Namespace: req.Namespace,
		Name:      resource.Spec.Cloud,
	}, openStackCloud); err != nil {
		if apierrors.IsNotFound(err) {
			err = fmt.Errorf("OpenStackCloud %q not found: %w", resource.Spec.Cloud, err)
			logger.Info(err.Error())
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, err
	}

	// Initialize the patch helper
	resourcePatchHelper, err := patch.NewHelper(resource, r.Client)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Always patch the resource when exiting this function.
	defer func() {
		reterr = kerrors.NewAggregate([]error{
			reterr,
			resourcePatchHelper.Patch(ctx, resource),
		})
	}()

	networkClient, err := cloud.NewClient(log.IntoContext(ctx, logger), r.Client, openStackCloud, "network")
	if err != nil {
		err = fmt.Errorf("unable to build an OpenStack client: %w", err)
		logger.Info(err.Error())
		return ctrl.Result{}, err
	}

	if !resource.DeletionTimestamp.IsZero() {
		return r.reconcileDelete(log.IntoContext(ctx, logger), networkClient, resource)
	}

	return r.reconcile(log.IntoContext(ctx, logger), networkClient, resource)
}

// reconcile handles creation. No modification is accepted.
// TODO: restrict unhandled modification through a webhook
// TODO: potentially handle (some?) modifications accepted in OpenStack, as in `openstack network set`
func (r *OpenStackSecurityGroupRuleReconciler) reconcile(ctx context.Context, networkClient *gophercloud.ServiceClient, resource *openstackv1.OpenStackSecurityGroupRule) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// If the resource doesn't have our finalizer, add it.
	if controllerutil.AddFinalizer(resource, openstackv1.Finalizer) {
		// Register the finalizer immediately to avoid orphaning OpenStack resources on delete
		return ctrl.Result{}, nil
	}

	// If the resource has an ID set but hasn't been created by us, then
	// it's unmanaged by default.
	if resource.Spec.Unmanaged == nil {
		var unmanaged bool
		if resource.Spec.ID != "" && resource.Status.ID == "" {
			unmanaged = true
		}
		resource.Spec.Unmanaged = &unmanaged
		return ctrl.Result{}, nil
	}

	var openstackResource *rules.SecGroupRule
	if resource.Spec.ID != "" {
		logger = logger.WithValues("OpenStackID", resource.Spec.ID)

		var err error
		openstackResource, err = rules.Get(networkClient, resource.Spec.ID).Extract()
		if err != nil {
			return ctrl.Result{}, err
		}
	} else {
		securityGroup := &openstackv1.OpenStackSecurityGroup{}
		err := r.Client.Get(ctx, types.NamespacedName{
			Namespace: resource.GetNamespace(),
			Name:      resource.Spec.SecurityGroup,
		}, securityGroup)
		if err != nil {
			if apierrors.IsNotFound(err) {
				logger.Info("parent security group %q not found in the API. Requeuing", resource.Spec.SecurityGroup)
				return ctrl.Result{}, fmt.Errorf("parent security group %q not found", resource.Spec.SecurityGroup)
			}
			return ctrl.Result{}, err
		}
		if securityGroup.Status.ID == "" {
			return ctrl.Result{}, fmt.Errorf("parent security group %q not found in OpenStack", resource.Spec.SecurityGroup)
		}

		openstackResource, err = rules.Create(networkClient, rules.CreateOpts{
			Direction:      rules.RuleDirection(resource.Spec.Direction),
			Description:    resource.Spec.Description,
			EtherType:      rules.RuleEtherType(resource.Spec.EtherType),
			SecGroupID:     securityGroup.Status.ID,
			PortRangeMax:   resource.Spec.PortRangeMax,
			PortRangeMin:   resource.Spec.PortRangeMin,
			Protocol:       rules.RuleProtocol(resource.Spec.Protocol),
			RemoteGroupID:  resource.Spec.RemoteGroupID,
			RemoteIPPrefix: resource.Spec.RemoteIPPrefix,
			ProjectID:      resource.Spec.ProjectID,
		}).Extract()
		if err != nil {
			return ctrl.Result{}, err
		}
		resource.Spec.ID = openstackResource.ID
		logger = logger.WithValues("OpenStackID", openstackResource.ID)
	}

	resource.Status = openstackv1.OpenStackSecurityGroupRuleStatus{
		ID:              openstackResource.ID,
		Direction:       openstackResource.Direction,
		Description:     openstackResource.Description,
		EtherType:       openstackResource.EtherType,
		SecurityGroupID: openstackResource.SecGroupID,
		PortRangeMin:    openstackResource.PortRangeMin,
		PortRangeMax:    openstackResource.PortRangeMax,
		Protocol:        openstackResource.Protocol,
		RemoteGroupID:   openstackResource.RemoteGroupID,
		RemoteIPPrefix:  openstackResource.RemoteIPPrefix,
		TenantID:        openstackResource.TenantID,
		ProjectID:       openstackResource.ProjectID,
	}

	logger.Info("reconcile succeeded.")
	return ctrl.Result{}, nil
}

func (r *OpenStackSecurityGroupRuleReconciler) reconcileDelete(ctx context.Context, networkClient *gophercloud.ServiceClient, resource *openstackv1.OpenStackSecurityGroupRule) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	if resource.Spec.ID == "" {
		logger.Info("deletion was requested on a resource that hasn't been created yet.")
	} else {
		logger = logger.WithValues("OpenStackID", resource.Spec.ID)
		if resource.Spec.Unmanaged != nil && !*resource.Spec.Unmanaged {
			if err := rules.Delete(networkClient, resource.Spec.ID).ExtractErr(); err != nil {
				var gerr gophercloud.ErrDefault404
				if errors.As(err, &gerr) {
					logger.Info("deletion was requested on a resource that can't be found in OpenStack.")
				} else {
					logger.Info("failed to delete resouce in OpenStack; requeuing.")
					return ctrl.Result{}, err
				}
			}
		}
	}

	controllerutil.RemoveFinalizer(resource, openstackv1.Finalizer)
	logger.Info("reconcileDelete succeeded.")
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *OpenStackSecurityGroupRuleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&openstackv1.OpenStackSecurityGroupRule{}).
		Complete(r)
}
