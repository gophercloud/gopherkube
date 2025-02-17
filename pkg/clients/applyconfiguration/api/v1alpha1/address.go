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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	apiv1alpha1 "github.com/k-orc/openstack-resource-controller/api/v1alpha1"
)

// AddressApplyConfiguration represents a declarative configuration of the Address type for use
// with apply.
type AddressApplyConfiguration struct {
	IP        *apiv1alpha1.IPvAny            `json:"ip,omitempty"`
	SubnetRef *apiv1alpha1.KubernetesNameRef `json:"subnetRef,omitempty"`
}

// AddressApplyConfiguration constructs a declarative configuration of the Address type for use with
// apply.
func Address() *AddressApplyConfiguration {
	return &AddressApplyConfiguration{}
}

// WithIP sets the IP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IP field is set to the value of the last call.
func (b *AddressApplyConfiguration) WithIP(value apiv1alpha1.IPvAny) *AddressApplyConfiguration {
	b.IP = &value
	return b
}

// WithSubnetRef sets the SubnetRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SubnetRef field is set to the value of the last call.
func (b *AddressApplyConfiguration) WithSubnetRef(value apiv1alpha1.KubernetesNameRef) *AddressApplyConfiguration {
	b.SubnetRef = &value
	return b
}
