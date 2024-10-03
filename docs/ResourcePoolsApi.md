# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryController1AddResourcePoolUser**](ResourcePoolsApi.md#InventoryController1AddResourcePoolUser) | **Post** /api/v2/resource-pools/user/{userId}/pool/{resourcePoolId} | Add a user to a Resource Pool
[**InventoryController1AddServerToResourcePool**](ResourcePoolsApi.md#InventoryController1AddServerToResourcePool) | **Put** /api/v2/resource-pools/{resourcePoolId}/server/{serverId} | Add a server to a Resource Pool
[**InventoryController1AddSubnetPoolToResourcePool**](ResourcePoolsApi.md#InventoryController1AddSubnetPoolToResourcePool) | **Put** /api/v2/resource-pools/{resourcePoolId}/subnet-pool/{subnetPoolId} | Add a subnet pool to a resource pool
[**InventoryController1CreateResourcePool**](ResourcePoolsApi.md#InventoryController1CreateResourcePool) | **Post** /api/v2/resource-pools | Creates a Resource Pool
[**InventoryController1DeleteResourcePool**](ResourcePoolsApi.md#InventoryController1DeleteResourcePool) | **Delete** /api/v2/resource-pools/{resourcePoolId} | Deletes a Resource Pool
[**InventoryController1GetResourcePool**](ResourcePoolsApi.md#InventoryController1GetResourcePool) | **Get** /api/v2/resource-pools/{resourcePoolId} | Get Resource Pool information
[**InventoryController1GetResourcePoolServers**](ResourcePoolsApi.md#InventoryController1GetResourcePoolServers) | **Get** /api/v2/resource-pools/{resourcePoolId}/servers | Get all servers that are part of a Resource Pool
[**InventoryController1GetResourcePoolSubnetPools**](ResourcePoolsApi.md#InventoryController1GetResourcePoolSubnetPools) | **Get** /api/v2/resource-pools/{resourcePoolId}/subnet-pools | Get all subnet pools that are part of a resource pool
[**InventoryController1GetResourcePoolUsers**](ResourcePoolsApi.md#InventoryController1GetResourcePoolUsers) | **Get** /api/v2/resource-pools/{resourcePoolId}/users | Get all users that have access to a Resource Pool
[**InventoryController1GetResourcePools**](ResourcePoolsApi.md#InventoryController1GetResourcePools) | **Get** /api/v2/resource-pools | Get all Resource Pools
[**InventoryController1GetUserResourcePools**](ResourcePoolsApi.md#InventoryController1GetUserResourcePools) | **Get** /api/v2/resource-pools/user/{userId} | Get all Resource Pools that a user has access to
[**InventoryController1RemoveResourcePoolUser**](ResourcePoolsApi.md#InventoryController1RemoveResourcePoolUser) | **Delete** /api/v2/resource-pools/user/{userId}/pool/{resourcePoolId} | Remove a user from a Resource Pool
[**InventoryController1RemoveServerFromResourcePool**](ResourcePoolsApi.md#InventoryController1RemoveServerFromResourcePool) | **Delete** /api/v2/resource-pools/{resourcePoolId}/server/{serverId} | Remove a server from a Resource Pool
[**InventoryController1RemoveSubnetPoolFromResourcePool**](ResourcePoolsApi.md#InventoryController1RemoveSubnetPoolFromResourcePool) | **Delete** /api/v2/resource-pools/{resourcePoolId}/subnet-pool/{subnetPoolId} | Remove a subnet from a resource pool
[**InventoryController1UpdateResourcePool**](ResourcePoolsApi.md#InventoryController1UpdateResourcePool) | **Put** /api/v2/resource-pools/{resourcePoolId} | Updates Resource Pool information

# **InventoryController1AddResourcePoolUser**
> InventoryController1AddResourcePoolUser(ctx, resourcePoolId, userId)
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

# **InventoryController1AddServerToResourcePool**
> InventoryController1AddServerToResourcePool(ctx, resourcePoolId, serverId)
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

# **InventoryController1AddSubnetPoolToResourcePool**
> InventoryController1AddSubnetPoolToResourcePool(ctx, resourcePoolId, subnetPoolId)
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

# **InventoryController1CreateResourcePool**
> ResourcePoolDto InventoryController1CreateResourcePool(ctx, body)
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

# **InventoryController1DeleteResourcePool**
> InventoryController1DeleteResourcePool(ctx, resourcePoolId)
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

# **InventoryController1GetResourcePool**
> ResourcePoolDto InventoryController1GetResourcePool(ctx, resourcePoolId)
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

# **InventoryController1GetResourcePoolServers**
> InventoryController1GetResourcePoolServers(ctx, resourcePoolId)
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

# **InventoryController1GetResourcePoolSubnetPools**
> InventoryController1GetResourcePoolSubnetPools(ctx, resourcePoolId)
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

# **InventoryController1GetResourcePoolUsers**
> InventoryController1GetResourcePoolUsers(ctx, resourcePoolId)
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

# **InventoryController1GetResourcePools**
> []ResourcePoolWithStatsDto InventoryController1GetResourcePools(ctx, )
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

# **InventoryController1GetUserResourcePools**
> InventoryController1GetUserResourcePools(ctx, userId)
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

# **InventoryController1RemoveResourcePoolUser**
> InventoryController1RemoveResourcePoolUser(ctx, resourcePoolId, userId)
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

# **InventoryController1RemoveServerFromResourcePool**
> InventoryController1RemoveServerFromResourcePool(ctx, resourcePoolId, serverId)
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

# **InventoryController1RemoveSubnetPoolFromResourcePool**
> InventoryController1RemoveSubnetPoolFromResourcePool(ctx, resourcePoolId, subnetPoolId)
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

# **InventoryController1UpdateResourcePool**
> ResourcePoolDto InventoryController1UpdateResourcePool(ctx, body, resourcePoolId)
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

