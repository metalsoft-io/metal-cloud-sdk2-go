/*
 * MetalSoft REST API
 *
 * MetalSoft REST API documentation
 *
 * API version: 2.0
 * Contact: support@metalsoft.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package sdk2

type AssetDto struct {
	// Name of the asset.
	Name string `json:"name" yaml:"name"`
	// URL of the asset.
	Url string `json:"url" yaml:"url"`
	// Label of the asset.
	Label string `json:"label" yaml:"label"`
	// Type of the asset.
	Type_ string `json:"type" yaml:"type"`
	// Required assets for the asset.
	RequiredAssets string `json:"required_assets" yaml:"required_assets"`
}
