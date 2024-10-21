# Bucket

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float64** | Id of the Bucket | [default to null]
**ChangeId** | **float64** | Change id of the Bucket | [default to null]
**InfrastructureId** | **float64** | Infrastructure id of the Bucket | [default to null]
**SizeGB** | **float64** | Disk size in GB for Bucket | [default to null]
**CreatedTimestamp** | **string** | Timestamp of the Bucket creation. | [default to null]
**UpdatedTimestamp** | **string** | Timestamp of the Bucket last update. | [default to null]
**StoragePoolId** | **float64** | Id of the storage pool the Bucket is assigned to | [optional] [default to null]
**ServiceStatus** | **string** | Service status of the Bucket | [default to null]
**Label** | **string** | Label of the Bucket. | [default to null]
**Subdomain** | **string** | Subdomain of the Bucket. | [optional] [default to null]
**SubdomainPermanent** | **string** | Subdomain permanent of the Bucket. | [optional] [default to null]
**DnsSubdomainId** | **float64** | Id of the DNS subdomain for the Bucket. | [optional] [default to null]
**NetworkVlanId** | **float64** | Id of the VLAN for the Bucket. | [optional] [default to null]
**GuiSettings** | [***AllOfBucketGuiSettings**](AllOfBucketGuiSettings.md) | GUI settings for the Bucket. This is a JSON object. | [optional] [default to null]
**Endpoint** | **string** | Endpoint of the Bucket. | [optional] [default to null]
**AccessKeyId** | **string** | Endpoint of the Bucket. | [optional] [default to null]
**SecretKey** | **string** | Endpoint of the Bucket. | [optional] [default to null]
**Operation** | **string** | Operation object of the Bucket. | [optional] [default to null]
**Links** | [***interface{}**](interface{}.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

