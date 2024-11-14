# ExtensionDefinitionDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The kind of extension. | [default to null]
**SchemaVersion** | **string** | Schema version of the extension. | [default to null]
**Name** | **string** | Name of the extension. | [default to null]
**Label** | **string** | Label of the extension. | [optional] [default to null]
**ExtensionType** | **string** | Type of the extension. | [default to null]
**Vendor** | **string** | Vendor of the extension. | [default to null]
**ExtensionVersion** | **string** | Version of the extension. | [default to null]
**Description** | **string** | Description of the extension. | [optional] [default to null]
**Icon** | **string** | Icon of the extension. | [default to null]
**Dependencies** | [***ExtensionDependency**](ExtensionDependency.md) |  | [default to null]
**Inputs** | [**[]ExtensionInput**](ExtensionInput.md) | List of inputs for the platform service. | [default to null]
**Outputs** | [**[]ExtensionOutput**](ExtensionOutput.md) | List of outputs for the platform service. | [default to null]
**Infrastructure** | [***ExtensionInfrastructure**](ExtensionInfrastructure.md) |  | [default to null]
**Assets** | [**[]ExtensionAsset**](ExtensionAsset.md) | List of assets for the platform service. | [default to null]
**OnCreate** | [***ExtensionActions**](ExtensionActions.md) |  | [optional] [default to null]
**OnEdit** | [***ExtensionActions**](ExtensionActions.md) |  | [optional] [default to null]
**OnDelete** | [***ExtensionActions**](ExtensionActions.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

