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

type BucketExtendedInfo struct {
	// Id of the Bucket
	Id float64 `json:"id,omitempty"`
	// Change id of the Bucket
	ChangeId float64 `json:"changeId,omitempty"`
	// Infrastructure id of the Bucket
	InfrastructureId float64 `json:"infrastructureId,omitempty"`
	// Disk size in GB for Bucket
	SizeGB float64 `json:"sizeGB,omitempty"`
	// Timestamp of the Bucket creation.
	CreatedTimestamp string `json:"createdTimestamp,omitempty"`
	// Timestamp of the Bucket last update.
	UpdatedTimestamp string `json:"updatedTimestamp,omitempty"`
	// Id of the storage pool the Bucket is assigned to
	StoragePoolId float64 `json:"storagePoolId,omitempty"`
	// Service status of the Bucket
	ServiceStatus string `json:"serviceStatus,omitempty"`
	// Label of the Bucket.
	Label string `json:"label,omitempty"`
	// Subdomain of the Bucket.
	Subdomain string `json:"subdomain,omitempty"`
	// Subdomain permanent of the Bucket.
	SubdomainPermanent string `json:"subdomainPermanent,omitempty"`
	// Id of the DNS subdomain for the Bucket.
	DnsSubdomainId float64 `json:"dnsSubdomainId,omitempty"`
	// Id of the VLAN for the Bucket.
	NetworkVlanId float64 `json:"networkVlanId,omitempty"`
	// GUI settings for the Bucket. This is a JSON object.
	GuiSettings *AllOfBucketExtendedInfoGuiSettings `json:"guiSettings,omitempty"`
	// Endpoint of the Bucket.
	Endpoint string `json:"endpoint,omitempty"`
	// Endpoint of the Bucket.
	AccessKeyId string `json:"accessKeyId,omitempty"`
	// Endpoint of the Bucket.
	SecretKeyEncrypted string `json:"secretKeyEncrypted,omitempty"`
	// Infrastructure information
	Infrastructure *interface{} `json:"infrastructure"`
	Links *interface{} `json:"links,omitempty"`
}
