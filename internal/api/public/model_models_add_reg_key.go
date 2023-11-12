/*
Nexodus API

This is the Nexodus API Server.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package public

// ModelsAddRegKey struct for ModelsAddRegKey
type ModelsAddRegKey struct {
	Description string `json:"description,omitempty"`
	// ExpiresAt is optional, if set the registration key is only valid until the ExpiresAt time.
	ExpiresAt string `json:"expires_at,omitempty"`
	// SingleUse only allows the registration key to be used once.
	SingleUse bool   `json:"single_use,omitempty"`
	VpcId     string `json:"vpc_id,omitempty"`
}