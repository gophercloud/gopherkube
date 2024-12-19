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

// SubnetResourceSpecApplyConfiguration represents a declarative configuration of the SubnetResourceSpec type for use
// with apply.
type SubnetResourceSpecApplyConfiguration struct {
	Name              *v1alpha1.OpenStackName            `json:"name,omitempty"`
	Description       *v1alpha1.OpenStackDescription     `json:"description,omitempty"`
	Tags              []v1alpha1.NeutronTag              `json:"tags,omitempty"`
	IPVersion         *v1alpha1.IPVersion                `json:"ipVersion,omitempty"`
	CIDR              *v1alpha1.CIDR                     `json:"cidr,omitempty"`
	ProjectID         *v1alpha1.UUID                     `json:"projectID,omitempty"`
	AllocationPools   []AllocationPoolApplyConfiguration `json:"allocationPools,omitempty"`
	Gateway           *SubnetGatewayApplyConfiguration   `json:"gateway,omitempty"`
	EnableDHCP        *bool                              `json:"enableDHCP,omitempty"`
	DNSNameservers    []v1alpha1.IPvAny                  `json:"dnsNameservers,omitempty"`
	DNSPublishFixedIP *bool                              `json:"dnsPublishFixedIP,omitempty"`
	HostRoutes        []HostRouteApplyConfiguration      `json:"hostRoutes,omitempty"`
	IPv6              *IPv6OptionsApplyConfiguration     `json:"ipv6,omitempty"`
	RouterRef         *v1alpha1.KubernetesNameRef        `json:"routerRef,omitempty"`
}

// SubnetResourceSpecApplyConfiguration constructs a declarative configuration of the SubnetResourceSpec type for use with
// apply.
func SubnetResourceSpec() *SubnetResourceSpecApplyConfiguration {
	return &SubnetResourceSpecApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *SubnetResourceSpecApplyConfiguration) WithName(value v1alpha1.OpenStackName) *SubnetResourceSpecApplyConfiguration {
	b.Name = &value
	return b
}

// WithDescription sets the Description field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Description field is set to the value of the last call.
func (b *SubnetResourceSpecApplyConfiguration) WithDescription(value v1alpha1.OpenStackDescription) *SubnetResourceSpecApplyConfiguration {
	b.Description = &value
	return b
}

// WithTags adds the given value to the Tags field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tags field.
func (b *SubnetResourceSpecApplyConfiguration) WithTags(values ...v1alpha1.NeutronTag) *SubnetResourceSpecApplyConfiguration {
	for i := range values {
		b.Tags = append(b.Tags, values[i])
	}
	return b
}

// WithIPVersion sets the IPVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IPVersion field is set to the value of the last call.
func (b *SubnetResourceSpecApplyConfiguration) WithIPVersion(value v1alpha1.IPVersion) *SubnetResourceSpecApplyConfiguration {
	b.IPVersion = &value
	return b
}

// WithCIDR sets the CIDR field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CIDR field is set to the value of the last call.
func (b *SubnetResourceSpecApplyConfiguration) WithCIDR(value v1alpha1.CIDR) *SubnetResourceSpecApplyConfiguration {
	b.CIDR = &value
	return b
}

// WithProjectID sets the ProjectID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProjectID field is set to the value of the last call.
func (b *SubnetResourceSpecApplyConfiguration) WithProjectID(value v1alpha1.UUID) *SubnetResourceSpecApplyConfiguration {
	b.ProjectID = &value
	return b
}

// WithAllocationPools adds the given value to the AllocationPools field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AllocationPools field.
func (b *SubnetResourceSpecApplyConfiguration) WithAllocationPools(values ...*AllocationPoolApplyConfiguration) *SubnetResourceSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithAllocationPools")
		}
		b.AllocationPools = append(b.AllocationPools, *values[i])
	}
	return b
}

// WithGateway sets the Gateway field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Gateway field is set to the value of the last call.
func (b *SubnetResourceSpecApplyConfiguration) WithGateway(value *SubnetGatewayApplyConfiguration) *SubnetResourceSpecApplyConfiguration {
	b.Gateway = value
	return b
}

// WithEnableDHCP sets the EnableDHCP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnableDHCP field is set to the value of the last call.
func (b *SubnetResourceSpecApplyConfiguration) WithEnableDHCP(value bool) *SubnetResourceSpecApplyConfiguration {
	b.EnableDHCP = &value
	return b
}

// WithDNSNameservers adds the given value to the DNSNameservers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the DNSNameservers field.
func (b *SubnetResourceSpecApplyConfiguration) WithDNSNameservers(values ...v1alpha1.IPvAny) *SubnetResourceSpecApplyConfiguration {
	for i := range values {
		b.DNSNameservers = append(b.DNSNameservers, values[i])
	}
	return b
}

// WithDNSPublishFixedIP sets the DNSPublishFixedIP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DNSPublishFixedIP field is set to the value of the last call.
func (b *SubnetResourceSpecApplyConfiguration) WithDNSPublishFixedIP(value bool) *SubnetResourceSpecApplyConfiguration {
	b.DNSPublishFixedIP = &value
	return b
}

// WithHostRoutes adds the given value to the HostRoutes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the HostRoutes field.
func (b *SubnetResourceSpecApplyConfiguration) WithHostRoutes(values ...*HostRouteApplyConfiguration) *SubnetResourceSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithHostRoutes")
		}
		b.HostRoutes = append(b.HostRoutes, *values[i])
	}
	return b
}

// WithIPv6 sets the IPv6 field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IPv6 field is set to the value of the last call.
func (b *SubnetResourceSpecApplyConfiguration) WithIPv6(value *IPv6OptionsApplyConfiguration) *SubnetResourceSpecApplyConfiguration {
	b.IPv6 = value
	return b
}

// WithRouterRef sets the RouterRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RouterRef field is set to the value of the last call.
func (b *SubnetResourceSpecApplyConfiguration) WithRouterRef(value v1alpha1.KubernetesNameRef) *SubnetResourceSpecApplyConfiguration {
	b.RouterRef = &value
	return b
}
