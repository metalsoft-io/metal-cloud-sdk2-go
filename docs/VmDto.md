# VmDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float64** | VM ID | [default to null]
**Name** | **string** | Name of the VM | [default to null]
**SiteId** | **float64** | Id of the site for the VM | [default to null]
**DatacenterName** | **string** | Datacenter of the VM | [default to null]
**InfrastructureId** | **float64** | ID of the infrastructure where this VM is deployed | [default to null]
**UserId** | **float64** | ID of the user that owns this VM | [default to null]
**UserEmail** | **string** | Email of the user that owns this VM | [default to null]
**InstanceId** | **float64** | ID of the instance where this VM is deployed | [default to null]
**VmInstanceId** | **float64** | ID of the VM instance where this VM is deployed | [default to null]
**Host** | **string** | Name of the host | [default to null]
**Hosts** | **[]string** | List of hosts | [default to null]
**CpuCores** | **float64** | Number of CPU cores for the VM | [default to null]
**RamGB** | **float64** | RAM in GB for the VM | [default to null]
**DiskSizeGB** | **float64** | Disk size in GB for the VM | [default to null]
**TypeId** | **float64** | The id of the VM Type. This is a number. | [default to null]
**PoolId** | **float64** | The id of the VM Pool. This is a number. | [default to null]
**AdministrationState** | [***interface{}**](interface{}.md) | The administration state of the VM. | [default to null]
**Comments** | **float64** | VM comments. | [default to null]
**PowerState** | [***interface{}**](interface{}.md) | The power state of the VM. | [default to null]
**PowerStateLastUpdatedTimestamp** | **string** | Timestamp when the VM power state was last updated. | [default to null]
**CreatedTimestamp** | **string** | Timestamp when the VM was created | [default to null]
**AllocationTimestamp** | **string** | Timestamp when the VM was allocated | [default to null]
**Tags** | **[]string** | Tags for the VM. | [default to null]
**VncPort** | **float64** | The port number for the VNC server. | [default to null]
**VncPassword** | **string** | The password for the VNC server. | [default to null]
**Disks** | [**[]VmDiskDto**](VMDiskDto.md) | The disks of the VM. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

