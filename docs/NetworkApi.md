# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInfrastructureNetwork**](NetworkApi.md#CreateInfrastructureNetwork) | **Post** /api/v2/infrastructures/{infrastructureId}/networks | Creates a new LAN network on the infrastructure
[**DeleteInfrastructureNetwork**](NetworkApi.md#DeleteInfrastructureNetwork) | **Delete** /api/v2/infrastructures/{infrastructureId}/networks/{networkId} | Deletes a network from the infrastructure
[**GetInfrastructureNetwork**](NetworkApi.md#GetInfrastructureNetwork) | **Get** /api/v2/infrastructures/{infrastructureId}/networks/{networkId} | Gets the specified network from the infrastructure
[**GetInfrastructureNetworks**](NetworkApi.md#GetInfrastructureNetworks) | **Get** /api/v2/infrastructures/{infrastructureId}/networks | Retrieves all networks on the infrastructure

# **CreateInfrastructureNetwork**
> Network CreateInfrastructureNetwork(ctx, body, infrastructureId)
Creates a new LAN network on the infrastructure

Creates a new LAN network on the infrastructure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateNetwork**](CreateNetwork.md)| The Network create object | 
  **infrastructureId** | **float64**|  | 

### Return type

[**Network**](Network.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInfrastructureNetwork**
> DeleteInfrastructureNetwork(ctx, infrastructureId, networkId)
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

# **GetInfrastructureNetwork**
> GetInfrastructureNetwork(ctx, infrastructureId, networkId)
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

# **GetInfrastructureNetworks**
> GetInfrastructureNetworks(ctx, infrastructureId)
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

