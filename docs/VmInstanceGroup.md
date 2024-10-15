# VmInstanceGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float64** | Id of the VM Instance Group. | [default to null]
**InfrastructureId** | **float64** | Id of the Infrastructure. | [default to null]
**ServiceStatus** | **string** | Status of the VM Instance Group. | [default to null]
**ChangeId** | **float64** | Id of the VM Instance Group change object. | [default to null]
**Label** | **string** | Name of the VM Instance Group. | [default to null]
**InstanceCount** | **float64** | Number of VM instances in the VM Instance Group. | [default to null]
**CreatedTimestamp** | **string** | Timestamp of the VM Instance Group creation. | [default to null]
**UpdatedTimestamp** | **string** | Timestamp of the VM Instance Group update. | [default to null]
**Operation** | [***interface{}**](interface{}.md) | Operation object for the VM Instance Group. | [default to null]
**VmInstance** | **[]string** | Array of VM instances in the VM Instance Group. | [default to null]
**Tags** | **[]string** | Tags for the VM Instance Group. | [default to null]
**DiskSizeGB** | **float64** | Disk size in GB for each VM Instance in the VM Instance Group. | [default to null]
**VolumeTemplateId** | **float64** | Id of the template used by the VM Instance Group. | [optional] [default to null]
**NetworkIdToNetworkProfileId** | [***interface{}**](interface{}.md) | Network Id to Network Profile Id for the VM Instance Group. This is a JSON object. | [optional] [default to null]
**GuiSettings** | [***AllOfVmInstanceGroupGuiSettings**](AllOfVmInstanceGroupGuiSettings.md) | GUI settings for the VM Instance Group. This is a JSON object. | [optional] [default to null]
**Links** | [***interface{}**](interface{}.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

