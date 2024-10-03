# UpdateVmInstanceGroupDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label for the VM Instance Group. | [optional] [default to null]
**Tags** | **[]string** | Tags for the VM Instance Group. | [optional] [default to null]
**GuiSettings** | [***AllOfUpdateVmInstanceGroupDtoGuiSettings**](AllOfUpdateVmInstanceGroupDtoGuiSettings.md) | GUI settings for the VM Instance Group. This is a JSON object. | [optional] [default to null]
**VmInstanceGroupInterfaces** | [**[]UpdateVmInstanceGroupInterfaceDto**](UpdateVMInstanceGroupInterfaceDto.md) | Interfaces for the VM Instance Group | [optional] [default to null]
**CustomVariables** | [***interface{}**](interface{}.md) | Custom variables for the VM Instance. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

