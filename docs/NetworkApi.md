# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlueprintControllerCreateNetwork**](NetworkApi.md#BlueprintControllerCreateNetwork) | **Post** /api/v2/infrastructures/{infrastructureId}/networks | Creates a new LAN network on the infrastructure
[**BlueprintControllerDeleteNetwork**](NetworkApi.md#BlueprintControllerDeleteNetwork) | **Delete** /api/v2/infrastructures/{infrastructureId}/networks/{networkId} | Deletes a network from the infrastructure
[**BlueprintControllerGetNetwork**](NetworkApi.md#BlueprintControllerGetNetwork) | **Get** /api/v2/infrastructures/{infrastructureId}/networks/{networkId} | Gets the specified network from the infrastructure
[**BlueprintControllerGetNetworks**](NetworkApi.md#BlueprintControllerGetNetworks) | **Get** /api/v2/infrastructures/{infrastructureId}/networks | Retrieves all networks on the infrastructure

# **BlueprintControllerCreateNetwork**
> NetworkDto BlueprintControllerCreateNetwork(ctx, body, infrastructureId)
Creates a new LAN network on the infrastructure

Creates a new LAN network on the infrastructure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateNetworkDto**](CreateNetworkDto.md)| The Network create object | 
  **infrastructureId** | **float64**|  | 

### Return type

[**NetworkDto**](NetworkDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BlueprintControllerDeleteNetwork**
> BlueprintControllerDeleteNetwork(ctx, infrastructureId, networkId)
Deletes a network from the infrastructure

Deletes a network from the infrastructure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 
  **networkId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BlueprintControllerGetNetwork**
> BlueprintControllerGetNetwork(ctx, infrastructureId, networkId)
Gets the specified network from the infrastructure

Gets the specified network from the infrastructure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 
  **networkId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BlueprintControllerGetNetworks**
> BlueprintControllerGetNetworks(ctx, infrastructureId)
Retrieves all networks on the infrastructure

Retrieves all networks on the infrastructure

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

