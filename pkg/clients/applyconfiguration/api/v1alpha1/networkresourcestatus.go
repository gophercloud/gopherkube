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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetworkResourceStatusApplyConfiguration represents a declarative configuration of the NetworkResourceStatus type for use
// with apply.
type NetworkResourceStatusApplyConfiguration struct {
	Name                                    *string  `json:"name,omitempty"`
	Description                             *string  `json:"description,omitempty"`
	ProjectID                               *string  `json:"projectID,omitempty"`
	Status                                  *string  `json:"status,omitempty"`
	Tags                                    []string `json:"tags,omitempty"`
	NeutronStatusMetadataApplyConfiguration `json:",inline"`
	AdminStateUp                            *bool                                       `json:"adminStateUp,omitempty"`
	AvailabilityZoneHints                   []string                                    `json:"availabilityZoneHints,omitempty"`
	DNSDomain                               *string                                     `json:"dnsDomain,omitempty"`
	MTU                                     *int32                                      `json:"mtu,omitempty"`
	PortSecurityEnabled                     *bool                                       `json:"portSecurityEnabled,omitempty"`
	Provider                                *ProviderPropertiesStatusApplyConfiguration `json:"provider,omitempty"`
	External                                *bool                                       `json:"external,omitempty"`
	Shared                                  *bool                                       `json:"shared,omitempty"`
	Subnets                                 []string                                    `json:"subnets,omitempty"`
}

// NetworkResourceStatusApplyConfiguration constructs a declarative configuration of the NetworkResourceStatus type for use with
// apply.
func NetworkResourceStatus() *NetworkResourceStatusApplyConfiguration {
	return &NetworkResourceStatusApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *NetworkResourceStatusApplyConfiguration) WithName(value string) *NetworkResourceStatusApplyConfiguration {
	b.Name = &value
	return b
}

// WithDescription sets the Description field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Description field is set to the value of the last call.
func (b *NetworkResourceStatusApplyConfiguration) WithDescription(value string) *NetworkResourceStatusApplyConfiguration {
	b.Description = &value
	return b
}

// WithProjectID sets the ProjectID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProjectID field is set to the value of the last call.
func (b *NetworkResourceStatusApplyConfiguration) WithProjectID(value string) *NetworkResourceStatusApplyConfiguration {
	b.ProjectID = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *NetworkResourceStatusApplyConfiguration) WithStatus(value string) *NetworkResourceStatusApplyConfiguration {
	b.Status = &value
	return b
}

// WithTags adds the given value to the Tags field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tags field.
func (b *NetworkResourceStatusApplyConfiguration) WithTags(values ...string) *NetworkResourceStatusApplyConfiguration {
	for i := range values {
		b.Tags = append(b.Tags, values[i])
	}
	return b
}

// WithCreatedAt sets the CreatedAt field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreatedAt field is set to the value of the last call.
func (b *NetworkResourceStatusApplyConfiguration) WithCreatedAt(value v1.Time) *NetworkResourceStatusApplyConfiguration {
	b.NeutronStatusMetadataApplyConfiguration.CreatedAt = &value
	return b
}

// WithUpdatedAt sets the UpdatedAt field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UpdatedAt field is set to the value of the last call.
func (b *NetworkResourceStatusApplyConfiguration) WithUpdatedAt(value v1.Time) *NetworkResourceStatusApplyConfiguration {
	b.NeutronStatusMetadataApplyConfiguration.UpdatedAt = &value
	return b
}

// WithRevisionNumber sets the RevisionNumber field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RevisionNumber field is set to the value of the last call.
func (b *NetworkResourceStatusApplyConfiguration) WithRevisionNumber(value int64) *NetworkResourceStatusApplyConfiguration {
	b.NeutronStatusMetadataApplyConfiguration.RevisionNumber = &value
	return b
}

// WithAdminStateUp sets the AdminStateUp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AdminStateUp field is set to the value of the last call.
func (b *NetworkResourceStatusApplyConfiguration) WithAdminStateUp(value bool) *NetworkResourceStatusApplyConfiguration {
	b.AdminStateUp = &value
	return b
}

// WithAvailabilityZoneHints adds the given value to the AvailabilityZoneHints field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AvailabilityZoneHints field.
func (b *NetworkResourceStatusApplyConfiguration) WithAvailabilityZoneHints(values ...string) *NetworkResourceStatusApplyConfiguration {
	for i := range values {
		b.AvailabilityZoneHints = append(b.AvailabilityZoneHints, values[i])
	}
	return b
}

// WithDNSDomain sets the DNSDomain field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DNSDomain field is set to the value of the last call.
func (b *NetworkResourceStatusApplyConfiguration) WithDNSDomain(value string) *NetworkResourceStatusApplyConfiguration {
	b.DNSDomain = &value
	return b
}

// WithMTU sets the MTU field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MTU field is set to the value of the last call.
func (b *NetworkResourceStatusApplyConfiguration) WithMTU(value int32) *NetworkResourceStatusApplyConfiguration {
	b.MTU = &value
	return b
}

// WithPortSecurityEnabled sets the PortSecurityEnabled field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PortSecurityEnabled field is set to the value of the last call.
func (b *NetworkResourceStatusApplyConfiguration) WithPortSecurityEnabled(value bool) *NetworkResourceStatusApplyConfiguration {
	b.PortSecurityEnabled = &value
	return b
}

// WithProvider sets the Provider field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Provider field is set to the value of the last call.
func (b *NetworkResourceStatusApplyConfiguration) WithProvider(value *ProviderPropertiesStatusApplyConfiguration) *NetworkResourceStatusApplyConfiguration {
	b.Provider = value
	return b
}

// WithExternal sets the External field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the External field is set to the value of the last call.
func (b *NetworkResourceStatusApplyConfiguration) WithExternal(value bool) *NetworkResourceStatusApplyConfiguration {
	b.External = &value
	return b
}

// WithShared sets the Shared field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Shared field is set to the value of the last call.
func (b *NetworkResourceStatusApplyConfiguration) WithShared(value bool) *NetworkResourceStatusApplyConfiguration {
	b.Shared = &value
	return b
}

// WithSubnets adds the given value to the Subnets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Subnets field.
func (b *NetworkResourceStatusApplyConfiguration) WithSubnets(values ...string) *NetworkResourceStatusApplyConfiguration {
	for i := range values {
		b.Subnets = append(b.Subnets, values[i])
	}
	return b
}
