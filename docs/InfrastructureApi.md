# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryController1DeployInfrastructure**](InfrastructureApi.md#InventoryController1DeployInfrastructure) | **Post** /api/v2/infrastructures/{infrastructureId}/actions/deploy | Deploys the specified infrastructure
[**InventoryController1GetInfrastructure**](InfrastructureApi.md#InventoryController1GetInfrastructure) | **Get** /api/v2/infrastructures/{infrastructureId} | Retrieves the specified infrastructure
[**InventoryController1GetInfrastructures**](InfrastructureApi.md#InventoryController1GetInfrastructures) | **Get** /api/v2/infrastructures | Get all infrastructures
[**InventoryController1RevertInfrastructure**](InfrastructureApi.md#InventoryController1RevertInfrastructure) | **Post** /api/v2/infrastructures/{infrastructureId}/actions/revert | Reverts the specified infrastructure

# **InventoryController1DeployInfrastructure**
> interface{} InventoryController1DeployInfrastructure(ctx, body, infrastructureId)
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

# **InventoryController1GetInfrastructure**
> InventoryController1GetInfrastructure(ctx, infrastructureId)
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

# **InventoryController1GetInfrastructures**
> []InfrastructureDto InventoryController1GetInfrastructures(ctx, optional)
Get all infrastructures

Returns list of all infrastructures

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InfrastructureApiInventoryController1GetInfrastructuresOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InfrastructureApiInventoryController1GetInfrastructuresOpts struct
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

# **InventoryController1RevertInfrastructure**
> InventoryController1RevertInfrastructure(ctx, infrastructureId)
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

