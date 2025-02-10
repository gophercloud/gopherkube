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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"

	apiv1alpha1 "github.com/k-orc/openstack-resource-controller/api/v1alpha1"
	applyconfigurationapiv1alpha1 "github.com/k-orc/openstack-resource-controller/pkg/clients/applyconfiguration/api/v1alpha1"
	scheme "github.com/k-orc/openstack-resource-controller/pkg/clients/clientset/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// SecurityGroupsGetter has a method to return a SecurityGroupInterface.
// A group's client should implement this interface.
type SecurityGroupsGetter interface {
	SecurityGroups(namespace string) SecurityGroupInterface
}

// SecurityGroupInterface has methods to work with SecurityGroup resources.
type SecurityGroupInterface interface {
	Create(ctx context.Context, securityGroup *apiv1alpha1.SecurityGroup, opts v1.CreateOptions) (*apiv1alpha1.SecurityGroup, error)
	Update(ctx context.Context, securityGroup *apiv1alpha1.SecurityGroup, opts v1.UpdateOptions) (*apiv1alpha1.SecurityGroup, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, securityGroup *apiv1alpha1.SecurityGroup, opts v1.UpdateOptions) (*apiv1alpha1.SecurityGroup, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*apiv1alpha1.SecurityGroup, error)
	List(ctx context.Context, opts v1.ListOptions) (*apiv1alpha1.SecurityGroupList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *apiv1alpha1.SecurityGroup, err error)
	Apply(ctx context.Context, securityGroup *applyconfigurationapiv1alpha1.SecurityGroupApplyConfiguration, opts v1.ApplyOptions) (result *apiv1alpha1.SecurityGroup, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, securityGroup *applyconfigurationapiv1alpha1.SecurityGroupApplyConfiguration, opts v1.ApplyOptions) (result *apiv1alpha1.SecurityGroup, err error)
	SecurityGroupExpansion
}

// securityGroups implements SecurityGroupInterface
type securityGroups struct {
	*gentype.ClientWithListAndApply[*apiv1alpha1.SecurityGroup, *apiv1alpha1.SecurityGroupList, *applyconfigurationapiv1alpha1.SecurityGroupApplyConfiguration]
}

// newSecurityGroups returns a SecurityGroups
func newSecurityGroups(c *OpenstackV1alpha1Client, namespace string) *securityGroups {
	return &securityGroups{
		gentype.NewClientWithListAndApply[*apiv1alpha1.SecurityGroup, *apiv1alpha1.SecurityGroupList, *applyconfigurationapiv1alpha1.SecurityGroupApplyConfiguration](
			"securitygroups",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *apiv1alpha1.SecurityGroup { return &apiv1alpha1.SecurityGroup{} },
			func() *apiv1alpha1.SecurityGroupList { return &apiv1alpha1.SecurityGroupList{} },
		),
	}
}
