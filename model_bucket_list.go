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

type BucketList struct {
	// List of Buckets
	Data []BucketExtendedInfo `json:"data"`
	// Meta information
	Meta *interface{} `json:"meta"`
	// Links to other resources
	Links *interface{} `json:"links"`
}
