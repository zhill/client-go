/*
 * Anchore Engine API Server
 *
 * This is the Anchore Engine API. Provides the primary external API for users of the service.
 *
 * API version: 0.1.14
 * Contact: nurmi@anchore.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client
// AccessCredential A login credential mapped to a user identity. For password credentials, the username to present for Basic auth is the user's username from the user record
type AccessCredential struct {
	// The type of credential
	Type string `json:"type"`
	// The credential value (e.g. the password)
	Value string `json:"value"`
	// The timestamp of creation of the credential
	CreatedAt string `json:"created_at,omitempty"`
}
