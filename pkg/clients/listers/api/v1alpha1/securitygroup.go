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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	apiv1alpha1 "github.com/k-orc/openstack-resource-controller/api/v1alpha1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// SecurityGroupLister helps list SecurityGroups.
// All objects returned here must be treated as read-only.
type SecurityGroupLister interface {
	// List lists all SecurityGroups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*apiv1alpha1.SecurityGroup, err error)
	// SecurityGroups returns an object that can list and get SecurityGroups.
	SecurityGroups(namespace string) SecurityGroupNamespaceLister
	SecurityGroupListerExpansion
}

// securityGroupLister implements the SecurityGroupLister interface.
type securityGroupLister struct {
	listers.ResourceIndexer[*apiv1alpha1.SecurityGroup]
}

// NewSecurityGroupLister returns a new SecurityGroupLister.
func NewSecurityGroupLister(indexer cache.Indexer) SecurityGroupLister {
	return &securityGroupLister{listers.New[*apiv1alpha1.SecurityGroup](indexer, apiv1alpha1.Resource("securitygroup"))}
}

// SecurityGroups returns an object that can list and get SecurityGroups.
func (s *securityGroupLister) SecurityGroups(namespace string) SecurityGroupNamespaceLister {
	return securityGroupNamespaceLister{listers.NewNamespaced[*apiv1alpha1.SecurityGroup](s.ResourceIndexer, namespace)}
}

// SecurityGroupNamespaceLister helps list and get SecurityGroups.
// All objects returned here must be treated as read-only.
type SecurityGroupNamespaceLister interface {
	// List lists all SecurityGroups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*apiv1alpha1.SecurityGroup, err error)
	// Get retrieves the SecurityGroup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*apiv1alpha1.SecurityGroup, error)
	SecurityGroupNamespaceListerExpansion
}

// securityGroupNamespaceLister implements the SecurityGroupNamespaceLister
// interface.
type securityGroupNamespaceLister struct {
	listers.ResourceIndexer[*apiv1alpha1.SecurityGroup]
}
