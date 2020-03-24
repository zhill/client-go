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
// WhitelistItem Identifies a specific gate and trigger match from a policy against an image and indicates it should be ignored in final policy decisions
type WhitelistItem struct {
	Id string `json:"id,omitempty"`
	Gate string `json:"gate"`
	TriggerId string `json:"trigger_id"`
}
