# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddResourcePoolUser**](ResourcePoolsApi.md#AddResourcePoolUser) | **Post** /api/v2/resource-pools/user/{userId}/pool/{resourcePoolId} | Add a user to a Resource Pool
[**AddServerToResourcePool**](ResourcePoolsApi.md#AddServerToResourcePool) | **Put** /api/v2/resource-pools/{resourcePoolId}/server/{serverId} | Add a server to a Resource Pool
[**AddSubnetPoolToResourcePool**](ResourcePoolsApi.md#AddSubnetPoolToResourcePool) | **Put** /api/v2/resource-pools/{resourcePoolId}/subnet-pool/{subnetPoolId} | Add a subnet pool to a resource pool
[**CreateResourcePool**](ResourcePoolsApi.md#CreateResourcePool) | **Post** /api/v2/resource-pools | Creates a Resource Pool
[**DeleteResourcePool**](ResourcePoolsApi.md#DeleteResourcePool) | **Delete** /api/v2/resource-pools/{resourcePoolId} | Deletes a Resource Pool
[**GetResourcePool**](ResourcePoolsApi.md#GetResourcePool) | **Get** /api/v2/resource-pools/{resourcePoolId} | Get Resource Pool information
[**GetResourcePoolServers**](ResourcePoolsApi.md#GetResourcePoolServers) | **Get** /api/v2/resource-pools/{resourcePoolId}/servers | Get all servers that are part of a Resource Pool
[**GetResourcePoolSubnetPools**](ResourcePoolsApi.md#GetResourcePoolSubnetPools) | **Get** /api/v2/resource-pools/{resourcePoolId}/subnet-pools | Get all subnet pools that are part of a resource pool
[**GetResourcePoolUsers**](ResourcePoolsApi.md#GetResourcePoolUsers) | **Get** /api/v2/resource-pools/{resourcePoolId}/users | Get all users that have access to a Resource Pool
[**GetResourcePools**](ResourcePoolsApi.md#GetResourcePools) | **Get** /api/v2/resource-pools | Get all Resource Pools
[**GetUserResourcePools**](ResourcePoolsApi.md#GetUserResourcePools) | **Get** /api/v2/resource-pools/user/{userId} | Get all Resource Pools that a user has access to
[**RemoveResourcePoolUser**](ResourcePoolsApi.md#RemoveResourcePoolUser) | **Delete** /api/v2/resource-pools/user/{userId}/pool/{resourcePoolId} | Remove a user from a Resource Pool
[**RemoveServerFromResourcePool**](ResourcePoolsApi.md#RemoveServerFromResourcePool) | **Delete** /api/v2/resource-pools/{resourcePoolId}/server/{serverId} | Remove a server from a Resource Pool
[**RemoveSubnetPoolFromResourcePool**](ResourcePoolsApi.md#RemoveSubnetPoolFromResourcePool) | **Delete** /api/v2/resource-pools/{resourcePoolId}/subnet-pool/{subnetPoolId} | Remove a subnet from a resource pool
[**UpdateResourcePool**](ResourcePoolsApi.md#UpdateResourcePool) | **Put** /api/v2/resource-pools/{resourcePoolId} | Updates Resource Pool information

# **AddResourcePoolUser**
> AddResourcePoolUser(ctx, resourcePoolId, userId)
Add a user to a Resource Pool

Add a user to a Resource Pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourcePoolId** | **float64**|  | 
  **userId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddServerToResourcePool**
> AddServerToResourcePool(ctx, resourcePoolId, serverId)
Add a server to a Resource Pool

Add a server to a Resource Pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourcePoolId** | **float64**|  | 
  **serverId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddSubnetPoolToResourcePool**
> AddSubnetPoolToResourcePool(ctx, resourcePoolId, subnetPoolId)
Add a subnet pool to a resource pool

Add a subnet pool to a resource pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourcePoolId** | **float64**|  | 
  **subnetPoolId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateResourcePool**
> ResourcePoolDto CreateResourcePool(ctx, body)
Creates a Resource Pool

Creates a Resource Pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateResourcePoolDto**](CreateResourcePoolDto.md)| The Resource Pool create object | 

### Return type

[**ResourcePoolDto**](ResourcePoolDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteResourcePool**
> DeleteResourcePool(ctx, resourcePoolId)
Deletes a Resource Pool

Deletes a Resource Pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourcePoolId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourcePool**
> ResourcePoolDto GetResourcePool(ctx, resourcePoolId)
Get Resource Pool information

Returns Resource Pool information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourcePoolId** | **float64**|  | 

### Return type

[**ResourcePoolDto**](ResourcePoolDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourcePoolServers**
> GetResourcePoolServers(ctx, resourcePoolId)
Get all servers that are part of a Resource Pool

Returns list of all servers that are part of a Resource Pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourcePoolId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourcePoolSubnetPools**
> GetResourcePoolSubnetPools(ctx, resourcePoolId)
Get all subnet pools that are part of a resource pool

Returns list of all subnet pools that are part of a resource pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourcePoolId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourcePoolUsers**
> GetResourcePoolUsers(ctx, resourcePoolId)
Get all users that have access to a Resource Pool

Returns list of all users that have access to a Resource Pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourcePoolId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourcePools**
> []ResourcePoolWithStatsDto GetResourcePools(ctx, )
Get all Resource Pools

Returns list of all Resource Pools

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ResourcePoolWithStatsDto**](ResourcePoolWithStatsDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserResourcePools**
> GetUserResourcePools(ctx, userId)
Get all Resource Pools that a user has access to

Returns list of all Resource Pools that a user has access to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveResourcePoolUser**
> RemoveResourcePoolUser(ctx, resourcePoolId, userId)
Remove a user from a Resource Pool

Remove a user from a Resource Pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourcePoolId** | **float64**|  | 
  **userId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveServerFromResourcePool**
> RemoveServerFromResourcePool(ctx, resourcePoolId, serverId)
Remove a server from a Resource Pool

Remove a server from a Resource Pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourcePoolId** | **float64**|  | 
  **serverId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveSubnetPoolFromResourcePool**
> RemoveSubnetPoolFromResourcePool(ctx, resourcePoolId, subnetPoolId)
Remove a subnet from a resource pool

Remove a subnet from a resource pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourcePoolId** | **float64**|  | 
  **subnetPoolId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateResourcePool**
> ResourcePoolDto UpdateResourcePool(ctx, body, resourcePoolId)
Updates Resource Pool information

Updates Resource Pool information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateResourcePoolDto**](UpdateResourcePoolDto.md)| The Resource Pool update object | 
  **resourcePoolId** | **float64**|  | 

### Return type

[**ResourcePoolDto**](ResourcePoolDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

