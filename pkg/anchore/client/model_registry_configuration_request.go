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
// RegistryConfigurationRequest A registry record describing the endpoint and credentials for a registry
type RegistryConfigurationRequest struct {
	// Username portion of credential to use for this registry
	RegistryUser string `json:"registry_user,omitempty"`
	// Password portion of credential to use for this registry
	RegistryPass string `json:"registry_pass,omitempty"`
	// Type of registry
	RegistryType string `json:"registry_type,omitempty"`
	// hostname:port string for accessing the registry, as would be used in a docker pull operation. May include some or all of a repository and wildcards (e.g. docker.io/library/_* or gcr.io/myproject/myrepository)
	Registry string `json:"registry,omitempty"`
	// human readable name associated with registry record
	RegistryName string `json:"registry_name,omitempty"`
	// Use TLS/SSL verification for the registry URL
	RegistryVerify bool `json:"registry_verify,omitempty"`
}
