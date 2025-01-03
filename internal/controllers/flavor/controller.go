/*
Copyright 2024 The ORC Authors.

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

package flavor

import (
	"context"

	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	orcv1alpha1 "github.com/k-orc/openstack-resource-controller/api/v1alpha1"

	ctrlexport "github.com/k-orc/openstack-resource-controller/internal/controllers/export"
	"github.com/k-orc/openstack-resource-controller/internal/controllers/generic"
	"github.com/k-orc/openstack-resource-controller/internal/scope"
)

const (
	FieldOwner = "openstack.k-orc.cloud/flavorcontroller"
	// Field owner of transient status.
	SSAStatusTxn = "status"
)

// ssaFieldOwner returns the field owner for a specific named SSA transaction.
func ssaFieldOwner(txn string) client.FieldOwner {
	return client.FieldOwner(FieldOwner + "/" + txn)
}

type flavorReconcilerConstructor struct {
	scopeFactory scope.Factory
}

func New(scopeFactory scope.Factory) ctrlexport.Controller {
	return flavorReconcilerConstructor{scopeFactory: scopeFactory}
}

func (flavorReconcilerConstructor) GetName() string {
	return "flavor"
}

// orcFlavorReconciler reconciles an ORC Flavor.
type orcFlavorReconciler struct {
	client   client.Client
	recorder record.EventRecorder

	flavorReconcilerConstructor
}

var _ generic.ResourceControllerCommon = &orcFlavorReconciler{}

func (r orcFlavorReconciler) GetK8sClient() client.Client {
	return r.client
}

func (r orcFlavorReconciler) GetScopeFactory() scope.Factory {
	return r.scopeFactory
}

// SetupWithManager sets up the controller with the Manager.
func (c flavorReconcilerConstructor) SetupWithManager(ctx context.Context, mgr ctrl.Manager, options controller.Options) error {
	reconciler := orcFlavorReconciler{
		client:   mgr.GetClient(),
		recorder: mgr.GetEventRecorderFor("orc-flavor-controller"),

		flavorReconcilerConstructor: c,
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&orcv1alpha1.Flavor{}).
		WithOptions(options).
		Complete(&reconciler)
}
