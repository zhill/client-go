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
// EventDescription A description of an event type
type EventDescription struct {
	// The event type. The last component of the fully-qualified event_type (category.subcategory.event)
	Name string `json:"name,omitempty"`
	// The fully qualified event type as would be seen in the event payload
	Type string `json:"type,omitempty"`
	// The message associated with the event type
	Message string `json:"message,omitempty"`
	// The type of resource this event is generated from
	ResourceType string `json:"resource_type,omitempty"`
}