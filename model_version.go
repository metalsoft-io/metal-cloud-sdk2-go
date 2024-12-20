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
import (
	"time"
)

type Version struct {
	// The version of the installed MetalSoft software
	Version string `json:"version" yaml:"version"`
	// The release date of the version
	VersionDate time.Time `json:"versionDate,omitempty" yaml:"versionDate,omitempty"`
	// The build number of the version
	VersionBuild string `json:"versionBuild,omitempty" yaml:"versionBuild,omitempty"`
}
