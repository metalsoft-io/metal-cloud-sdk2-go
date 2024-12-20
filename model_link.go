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

type Link struct {
	// The relation
	Rel string `json:"rel" yaml:"rel"`
	// The link URI
	Href string `json:"href" yaml:"href"`
}
