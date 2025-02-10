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

package fake

import (
	v1alpha1 "github.com/k-orc/openstack-resource-controller/api/v1alpha1"
	apiv1alpha1 "github.com/k-orc/openstack-resource-controller/pkg/clients/applyconfiguration/api/v1alpha1"
	typedapiv1alpha1 "github.com/k-orc/openstack-resource-controller/pkg/clients/clientset/clientset/typed/api/v1alpha1"
	gentype "k8s.io/client-go/gentype"
)

// fakeSecurityGroups implements SecurityGroupInterface
type fakeSecurityGroups struct {
	*gentype.FakeClientWithListAndApply[*v1alpha1.SecurityGroup, *v1alpha1.SecurityGroupList, *apiv1alpha1.SecurityGroupApplyConfiguration]
	Fake *FakeOpenstackV1alpha1
}

func newFakeSecurityGroups(fake *FakeOpenstackV1alpha1, namespace string) typedapiv1alpha1.SecurityGroupInterface {
	return &fakeSecurityGroups{
		gentype.NewFakeClientWithListAndApply[*v1alpha1.SecurityGroup, *v1alpha1.SecurityGroupList, *apiv1alpha1.SecurityGroupApplyConfiguration](
			fake.Fake,
			namespace,
			v1alpha1.SchemeGroupVersion.WithResource("securitygroups"),
			v1alpha1.SchemeGroupVersion.WithKind("SecurityGroup"),
			func() *v1alpha1.SecurityGroup { return &v1alpha1.SecurityGroup{} },
			func() *v1alpha1.SecurityGroupList { return &v1alpha1.SecurityGroupList{} },
			func(dst, src *v1alpha1.SecurityGroupList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.SecurityGroupList) []*v1alpha1.SecurityGroup {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1alpha1.SecurityGroupList, items []*v1alpha1.SecurityGroup) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
