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

type PatchStorage struct {
	Options *interface{} `json:"options,omitempty"`
	InMaintenance float64 `json:"inMaintenance,omitempty"`
	IsExperimental float64 `json:"isExperimental,omitempty"`
	DrivePriority float64 `json:"drivePriority,omitempty"`
	SharedDrivePriority float64 `json:"sharedDrivePriority,omitempty"`
	Tags []string `json:"tags,omitempty"`
	PortGroupAllocationOrder []string `json:"portGroupAllocationOrder,omitempty"`
	PortGroupPhysicalPorts *interface{} `json:"portGroupPhysicalPorts,omitempty"`
	DefaultIoLimitPolicy string `json:"defaultIoLimitPolicy,omitempty"`
	ArrayId string `json:"arrayId,omitempty"`
	DirectorId string `json:"directorId,omitempty"`
	S3Hostname string `json:"s3Hostname,omitempty"`
	S3Port float64 `json:"s3Port,omitempty"`
}
