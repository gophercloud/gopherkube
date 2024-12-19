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
	v1alpha1 "github.com/k-orc/openstack-resource-controller/api/v1alpha1"
)

// ServerResourceSpecApplyConfiguration represents a declarative configuration of the ServerResourceSpec type for use
// with apply.
type ServerResourceSpecApplyConfiguration struct {
	Name      *v1alpha1.OpenStackName            `json:"name,omitempty"`
	ImageRef  *v1alpha1.KubernetesNameRef        `json:"imageRef,omitempty"`
	FlavorRef *v1alpha1.KubernetesNameRef        `json:"flavorRef,omitempty"`
	UserData  *UserDataSpecApplyConfiguration    `json:"userData,omitempty"`
	Ports     []ServerPortSpecApplyConfiguration `json:"ports,omitempty"`
}

// ServerResourceSpecApplyConfiguration constructs a declarative configuration of the ServerResourceSpec type for use with
// apply.
func ServerResourceSpec() *ServerResourceSpecApplyConfiguration {
	return &ServerResourceSpecApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ServerResourceSpecApplyConfiguration) WithName(value v1alpha1.OpenStackName) *ServerResourceSpecApplyConfiguration {
	b.Name = &value
	return b
}

// WithImageRef sets the ImageRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ImageRef field is set to the value of the last call.
func (b *ServerResourceSpecApplyConfiguration) WithImageRef(value v1alpha1.KubernetesNameRef) *ServerResourceSpecApplyConfiguration {
	b.ImageRef = &value
	return b
}

// WithFlavorRef sets the FlavorRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FlavorRef field is set to the value of the last call.
func (b *ServerResourceSpecApplyConfiguration) WithFlavorRef(value v1alpha1.KubernetesNameRef) *ServerResourceSpecApplyConfiguration {
	b.FlavorRef = &value
	return b
}

// WithUserData sets the UserData field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UserData field is set to the value of the last call.
func (b *ServerResourceSpecApplyConfiguration) WithUserData(value *UserDataSpecApplyConfiguration) *ServerResourceSpecApplyConfiguration {
	b.UserData = value
	return b
}

// WithPorts adds the given value to the Ports field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Ports field.
func (b *ServerResourceSpecApplyConfiguration) WithPorts(values ...*ServerPortSpecApplyConfiguration) *ServerResourceSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPorts")
		}
		b.Ports = append(b.Ports, *values[i])
	}
	return b
}
