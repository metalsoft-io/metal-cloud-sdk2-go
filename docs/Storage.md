# Storage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageId** | **float64** | Id of the Storage | [default to null]
**UserId** | **float64** | Id of the owner | [optional] [default to null]
**SiteId** | **float64** | Id of the site | [default to null]
**StorageDriver** | **string** | Storage driver | [default to null]
**StorageTechnology** | **string** | Storage technology | [default to null]
**StorageType** | **string** | Storage type | [default to null]
**Name** | **string** | Name of the storage | [default to null]
**IscsiHost** | **string** | ISCSI host | [optional] [default to null]
**IscsiPort** | **float64** | ISCSI port | [optional] [default to null]
**ManagementHost** | **string** | Management host | [default to null]
**ManagementUsername** | **string** | Username | [default to null]
**ManagementPassword** | **string** | Password | [default to null]
**Options** | [***interface{}**](interface{}.md) | Options for the storage | [optional] [default to null]
**InMaintenance** | **float64** | Specifies if the storage is in maintenance | [default to null]
**TargetIQN** | **string** | Target IQN | [optional] [default to null]
**IsExperimental** | **float64** | Specifies if the storage is experimental | [optional] [default to null]
**DrivePriority** | **float64** | Specifies if the drive priority | [default to 50]
**SharedDrivePriority** | **float64** | Specifies if the drive priority | [default to 50]
**AlternateSanIPs** | **[]string** | Alternate SAN IPs | [optional] [default to null]
**Tags** | **[]string** | Tags | [optional] [default to null]
**PortGroupAllocationOrder** | [***interface{}**](interface{}.md) | Port group allocation order | [optional] [default to null]
**PortGroupPhysicalPorts** | [***interface{}**](interface{}.md) | Port group physical ports | [optional] [default to null]
**DefaultIoLimitPolicy** | **string** | Default IO limit policy | [optional] [default to null]
**SubnetType** | **string** | Subnet type | [default to null]
**ArrayId** | **string** | Array id | [optional] [default to null]
**DirectorId** | **string** | Director id | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

