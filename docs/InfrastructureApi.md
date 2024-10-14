# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryControllerDeployInfrastructure**](InfrastructureApi.md#InventoryControllerDeployInfrastructure) | **Post** /api/v2/infrastructures/{infrastructureId}/actions/deploy | Deploys the specified infrastructure
[**InventoryControllerGetInfrastructure**](InfrastructureApi.md#InventoryControllerGetInfrastructure) | **Get** /api/v2/infrastructures/{infrastructureId} | Retrieves the specified infrastructure
[**InventoryControllerGetInfrastructures**](InfrastructureApi.md#InventoryControllerGetInfrastructures) | **Get** /api/v2/infrastructures | Get all infrastructures
[**InventoryControllerRevertInfrastructure**](InfrastructureApi.md#InventoryControllerRevertInfrastructure) | **Post** /api/v2/infrastructures/{infrastructureId}/actions/revert | Reverts the specified infrastructure

# **InventoryControllerDeployInfrastructure**
> interface{} InventoryControllerDeployInfrastructure(ctx, body, infrastructureId)
Deploys the specified infrastructure

Deploys the specified infrastructure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InfrastructureDeployOptions**](InfrastructureDeployOptions.md)| The infrastructure deploy options | 
  **infrastructureId** | **float64**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryControllerGetInfrastructure**
> InventoryControllerGetInfrastructure(ctx, infrastructureId)
Retrieves the specified infrastructure

Retrieves the specified infrastructure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryControllerGetInfrastructures**
> []InfrastructureDto InventoryControllerGetInfrastructures(ctx, optional)
Get all infrastructures

Returns list of all infrastructures

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InfrastructureApiInventoryControllerGetInfrastructuresOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InfrastructureApiInventoryControllerGetInfrastructuresOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **optional.Float64**| The account ID to filter user infrastructures by | 

### Return type

[**[]InfrastructureDto**](InfrastructureDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryControllerRevertInfrastructure**
> InventoryControllerRevertInfrastructure(ctx, infrastructureId)
Reverts the specified infrastructure

Reverts the specified infrastructure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

