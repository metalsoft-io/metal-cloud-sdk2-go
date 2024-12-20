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

type ServerReRegistrationResponseDto struct {
	// The id of the server.
	Id float64 `json:"id" yaml:"id"`
	JobInfo *AfcJobInfo `json:"jobInfo,omitempty" yaml:"jobInfo,omitempty"`
}
