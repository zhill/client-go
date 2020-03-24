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
// AnalysisArchiveSource An image reference in the analysis archive for the purposes of loading analysis from the archive into th working set
type AnalysisArchiveSource struct {
	// The image digest identify the analysis. Archived analyses are based on digest, tag records are restored as analysis is restored.
	Digest string `json:"digest"`
}
