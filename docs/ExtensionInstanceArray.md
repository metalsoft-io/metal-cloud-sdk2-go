# ExtensionInstanceArray

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label of the instance array. | [default to null]
**InstanceCount** | **string** | Instance count value or reference to variable. | [default to null]
**ServerType** | **string** | Server type variable reference. | [default to null]
**OsTemplate** | **string** | OS template variable reference. | [default to null]
**ConnectedSharedDrives** | **[]string** | Connected shared drive arrays. | [optional] [default to null]
**CustomVariables** | [**[]CustomVariable**](CustomVariable.md) | Custom variables. The value may be a reference to an input variable. | [optional] [default to null]
**Dependencies** | **[]string** | Labels of instance arrays this one depends on. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

