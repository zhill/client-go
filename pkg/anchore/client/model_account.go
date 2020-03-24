/*
 * Anchore Engine API Server
 *
 * This is the Anchore Engine API. Provides the primary external API for users of the service.
 *
 * API version: 0.1.13
 * Contact: nurmi@anchore.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client
import (
	"time"
)
// Account Account information
type Account struct {
	// The account identifier, not updatable after creation
	Name string `json:"name"`
	// The user type (admin vs user). If not specified in a POST request, 'user' is default
	Type string `json:"type,omitempty"`
	// State of the account. Disabled accounts prevent member users from logging in, deleting accounts are disabled and pending deletion and will be removed once all owned resources are garbage collected by the system
	State string `json:"state,omitempty"`
	// Optional email address associated with the account
	Email string `json:"email,omitempty"`
	// The timestamp when the account was created
	CreatedAt time.Time `json:"created_at,omitempty"`
	// The timestamp of the last update to the account metadata itself (not users or creds)
	LastUpdated time.Time `json:"last_updated,omitempty"`
}
