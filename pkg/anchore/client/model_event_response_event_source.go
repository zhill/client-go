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

// EventResponseEventSource struct for EventResponseEventSource
type EventResponseEventSource struct {
	Servicename string `json:"servicename,omitempty"`
	Hostid      string `json:"hostid,omitempty"`
	BaseUrl     string `json:"base_url,omitempty"`
	RequestId   string `json:"request_id,omitempty"`
}
