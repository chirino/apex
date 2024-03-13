/*
Nexodus API

This is the Nexodus API Server.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ModelsRegKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsRegKey{}

// ModelsRegKey struct for ModelsRegKey
type ModelsRegKey struct {
	// BearerToken is the bearer token the client should use to authenticate the device registration request.
	BearerToken *string `json:"bearer_token,omitempty"`
	// Description of the registration key.
	Description *string `json:"description,omitempty"`
	// DeviceId is set if the RegKey was created for single use
	DeviceId *string `json:"device_id,omitempty"`
	// ExpiresAt is optional, if set the registration key is only valid until the ExpiresAt time.
	ExpiresAt *string `json:"expires_at,omitempty"`
	Id        *string `json:"id,omitempty"`
	// OwnerID is the ID of the user that created the registration key.
	OwnerId *string `json:"owner_id,omitempty"`
	// SecurityGroupId is the ID of the security group to assign to the device.
	SecurityGroupId *string `json:"security_group_id,omitempty"`
	// ServiceNetworkID is the ID of the Service Network the device can join.
	ServiceNetworkId *string `json:"service_network_id,omitempty"`
	// Settings contains general settings for the device.
	Settings map[string]interface{} `json:"settings,omitempty"`
	// VpcID is the ID of the VPC the device can join.
	VpcId *string `json:"vpc_id,omitempty"`
}

// NewModelsRegKey instantiates a new ModelsRegKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsRegKey() *ModelsRegKey {
	this := ModelsRegKey{}
	return &this
}

// NewModelsRegKeyWithDefaults instantiates a new ModelsRegKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsRegKeyWithDefaults() *ModelsRegKey {
	this := ModelsRegKey{}
	return &this
}

// GetBearerToken returns the BearerToken field value if set, zero value otherwise.
func (o *ModelsRegKey) GetBearerToken() string {
	if o == nil || IsNil(o.BearerToken) {
		var ret string
		return ret
	}
	return *o.BearerToken
}

// GetBearerTokenOk returns a tuple with the BearerToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRegKey) GetBearerTokenOk() (*string, bool) {
	if o == nil || IsNil(o.BearerToken) {
		return nil, false
	}
	return o.BearerToken, true
}

// HasBearerToken returns a boolean if a field has been set.
func (o *ModelsRegKey) HasBearerToken() bool {
	if o != nil && !IsNil(o.BearerToken) {
		return true
	}

	return false
}

// SetBearerToken gets a reference to the given string and assigns it to the BearerToken field.
func (o *ModelsRegKey) SetBearerToken(v string) {
	o.BearerToken = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModelsRegKey) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRegKey) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelsRegKey) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModelsRegKey) SetDescription(v string) {
	o.Description = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *ModelsRegKey) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRegKey) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *ModelsRegKey) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *ModelsRegKey) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *ModelsRegKey) GetExpiresAt() string {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRegKey) GetExpiresAtOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ModelsRegKey) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *ModelsRegKey) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsRegKey) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRegKey) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsRegKey) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelsRegKey) SetId(v string) {
	o.Id = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *ModelsRegKey) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRegKey) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *ModelsRegKey) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *ModelsRegKey) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetSecurityGroupId returns the SecurityGroupId field value if set, zero value otherwise.
func (o *ModelsRegKey) GetSecurityGroupId() string {
	if o == nil || IsNil(o.SecurityGroupId) {
		var ret string
		return ret
	}
	return *o.SecurityGroupId
}

// GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRegKey) GetSecurityGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.SecurityGroupId) {
		return nil, false
	}
	return o.SecurityGroupId, true
}

// HasSecurityGroupId returns a boolean if a field has been set.
func (o *ModelsRegKey) HasSecurityGroupId() bool {
	if o != nil && !IsNil(o.SecurityGroupId) {
		return true
	}

	return false
}

// SetSecurityGroupId gets a reference to the given string and assigns it to the SecurityGroupId field.
func (o *ModelsRegKey) SetSecurityGroupId(v string) {
	o.SecurityGroupId = &v
}

// GetServiceNetworkId returns the ServiceNetworkId field value if set, zero value otherwise.
func (o *ModelsRegKey) GetServiceNetworkId() string {
	if o == nil || IsNil(o.ServiceNetworkId) {
		var ret string
		return ret
	}
	return *o.ServiceNetworkId
}

// GetServiceNetworkIdOk returns a tuple with the ServiceNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRegKey) GetServiceNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceNetworkId) {
		return nil, false
	}
	return o.ServiceNetworkId, true
}

// HasServiceNetworkId returns a boolean if a field has been set.
func (o *ModelsRegKey) HasServiceNetworkId() bool {
	if o != nil && !IsNil(o.ServiceNetworkId) {
		return true
	}

	return false
}

// SetServiceNetworkId gets a reference to the given string and assigns it to the ServiceNetworkId field.
func (o *ModelsRegKey) SetServiceNetworkId(v string) {
	o.ServiceNetworkId = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *ModelsRegKey) GetSettings() map[string]interface{} {
	if o == nil || IsNil(o.Settings) {
		var ret map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRegKey) GetSettingsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Settings) {
		return map[string]interface{}{}, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *ModelsRegKey) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given map[string]interface{} and assigns it to the Settings field.
func (o *ModelsRegKey) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// GetVpcId returns the VpcId field value if set, zero value otherwise.
func (o *ModelsRegKey) GetVpcId() string {
	if o == nil || IsNil(o.VpcId) {
		var ret string
		return ret
	}
	return *o.VpcId
}

// GetVpcIdOk returns a tuple with the VpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRegKey) GetVpcIdOk() (*string, bool) {
	if o == nil || IsNil(o.VpcId) {
		return nil, false
	}
	return o.VpcId, true
}

// HasVpcId returns a boolean if a field has been set.
func (o *ModelsRegKey) HasVpcId() bool {
	if o != nil && !IsNil(o.VpcId) {
		return true
	}

	return false
}

// SetVpcId gets a reference to the given string and assigns it to the VpcId field.
func (o *ModelsRegKey) SetVpcId(v string) {
	o.VpcId = &v
}

func (o ModelsRegKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsRegKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BearerToken) {
		toSerialize["bearer_token"] = o.BearerToken
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DeviceId) {
		toSerialize["device_id"] = o.DeviceId
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.OwnerId) {
		toSerialize["owner_id"] = o.OwnerId
	}
	if !IsNil(o.SecurityGroupId) {
		toSerialize["security_group_id"] = o.SecurityGroupId
	}
	if !IsNil(o.ServiceNetworkId) {
		toSerialize["service_network_id"] = o.ServiceNetworkId
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.VpcId) {
		toSerialize["vpc_id"] = o.VpcId
	}
	return toSerialize, nil
}

type NullableModelsRegKey struct {
	value *ModelsRegKey
	isSet bool
}

func (v NullableModelsRegKey) Get() *ModelsRegKey {
	return v.value
}

func (v *NullableModelsRegKey) Set(val *ModelsRegKey) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsRegKey) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsRegKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsRegKey(val *ModelsRegKey) *NullableModelsRegKey {
	return &NullableModelsRegKey{value: val, isSet: true}
}

func (v NullableModelsRegKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsRegKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}