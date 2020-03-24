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
import (
	"time"
)
// RegistryDigestSource An image reference using a digest in a registry, includes some extra tag and timestamp info in addition to the pull string to allow proper tag history reconstruction.
type RegistryDigestSource struct {
	// A digest-based pullstring (e.g. docker.io/nginx@sha256:123abc)
	Pullstring string `json:"pullstring"`
	// A docker pull string (e.g. docker.io/nginx:latest, or docker.io/nginx@sha256:abd) to retrieve the image
	Tag string `json:"tag"`
	// Optional override of the image creation time to support proper tag history construction in cases of out-of-order analysis compared to registry history for the tag
	CreationTimestampOverride time.Time `json:"creation_timestamp_override"`
	// Base64 encoded content of the dockerfile used to build the image, if available.
	Dockerfile string `json:"dockerfile,omitempty"`
}
