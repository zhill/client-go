/*
 * Anchore Engine API Server
 *
 * This is the Anchore Engine API. Provides the primary external API for users of the service.
 *
 * API version: 0.1.12
 * Contact: nurmi@anchore.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client
// AccountCreationRequest An account to create/add to the system. If already exists will return 400.
type AccountCreationRequest struct {
	// The account name to use. This will identify the account and must be globally unique in the system.
	Name string `json:"name"`
	// An optional email to associate with the account for contact purposes
	Email string `json:"email,omitempty"`
}
