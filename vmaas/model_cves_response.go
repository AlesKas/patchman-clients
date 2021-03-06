/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.20.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vmaas
// CvesResponse struct for CvesResponse
type CvesResponse struct {
	Page float32 `json:"page,omitempty"`
	PageSize float32 `json:"page_size,omitempty"`
	Pages float32 `json:"pages,omitempty"`
	CveList map[string]CvesResponseCveList `json:"cve_list,omitempty"`
	ModifiedSince string `json:"modified_since,omitempty"`
}
