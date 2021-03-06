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

// VendorDataObject struct for VendorDataObject
type VendorDataObject struct {
	// Vendor Vulnerability ID
	Id     string       `json:"id,omitempty"`
	CvssV2 Cvssv2Scores `json:"cvss_v2,omitempty"`
	CvssV3 Cvssv3Scores `json:"cvss_v3,omitempty"`
}
