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
// EventResponse A record of occurance of an asynchronous event triggered either by system or by user activity
type EventResponse struct {
	GeneratedUuid string `json:"generated_uuid,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Event EventResponseEvent `json:"event,omitempty"`
}
