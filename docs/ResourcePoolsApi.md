# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryControllerAddResourcePoolUser**](ResourcePoolsApi.md#InventoryControllerAddResourcePoolUser) | **Post** /api/v2/resource-pools/user/{userId}/pool/{resourcePoolId} | Add a user to a Resource Pool
[**InventoryControllerAddServerToResourcePool**](ResourcePoolsApi.md#InventoryControllerAddServerToResourcePool) | **Put** /api/v2/resource-pools/{resourcePoolId}/server/{serverId} | Add a server to a Resource Pool
[**InventoryControllerAddSubnetPoolToResourcePool**](ResourcePoolsApi.md#InventoryControllerAddSubnetPoolToResourcePool) | **Put** /api/v2/resource-pools/{resourcePoolId}/subnet-pool/{subnetPoolId} | Add a subnet pool to a resource pool
[**InventoryControllerCreateResourcePool**](ResourcePoolsApi.md#InventoryControllerCreateResourcePool) | **Post** /api/v2/resource-pools | Creates a Resource Pool
[**InventoryControllerDeleteResourcePool**](ResourcePoolsApi.md#InventoryControllerDeleteResourcePool) | **Delete** /api/v2/resource-pools/{resourcePoolId} | Deletes a Resource Pool
[**InventoryControllerGetResourcePool**](ResourcePoolsApi.md#InventoryControllerGetResourcePool) | **Get** /api/v2/resource-pools/{resourcePoolId} | Get Resource Pool information
[**InventoryControllerGetResourcePoolServers**](ResourcePoolsApi.md#InventoryControllerGetResourcePoolServers) | **Get** /api/v2/resource-pools/{resourcePoolId}/servers | Get all servers that are part of a Resource Pool
[**InventoryControllerGetResourcePoolSubnetPools**](ResourcePoolsApi.md#InventoryControllerGetResourcePoolSubnetPools) | **Get** /api/v2/resource-pools/{resourcePoolId}/subnet-pools | Get all subnet pools that are part of a resource pool
[**InventoryControllerGetResourcePoolUsers**](ResourcePoolsApi.md#InventoryControllerGetResourcePoolUsers) | **Get** /api/v2/resource-pools/{resourcePoolId}/users | Get all users that have access to a Resource Pool
[**InventoryControllerGetResourcePools**](ResourcePoolsApi.md#InventoryControllerGetResourcePools) | **Get** /api/v2/resource-pools | Get all Resource Pools
[**InventoryControllerGetUserResourcePools**](ResourcePoolsApi.md#InventoryControllerGetUserResourcePools) | **Get** /api/v2/resource-pools/user/{userId} | Get all Resource Pools that a user has access to
[**InventoryControllerRemoveResourcePoolUser**](ResourcePoolsApi.md#InventoryControllerRemoveResourcePoolUser) | **Delete** /api/v2/resource-pools/user/{userId}/pool/{resourcePoolId} | Remove a user from a Resource Pool
[**InventoryControllerRemoveServerFromResourcePool**](ResourcePoolsApi.md#InventoryControllerRemoveServerFromResourcePool) | **Delete** /api/v2/resource-pools/{resourcePoolId}/server/{serverId} | Remove a server from a Resource Pool
[**InventoryControllerRemoveSubnetPoolFromResourcePool**](ResourcePoolsApi.md#InventoryControllerRemoveSubnetPoolFromResourcePool) | **Delete** /api/v2/resource-pools/{resourcePoolId}/subnet-pool/{subnetPoolId} | Remove a subnet from a resource pool
[**InventoryControllerUpdateResourcePool**](ResourcePoolsApi.md#InventoryControllerUpdateResourcePool) | **Put** /api/v2/resource-pools/{resourcePoolId} | Updates Resource Pool information

# **InventoryControllerAddResourcePoolUser**
> InventoryControllerAddResourcePoolUser(ctx, resourcePoolId, userId)
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

# **InventoryControllerAddServerToResourcePool**
> InventoryControllerAddServerToResourcePool(ctx, resourcePoolId, serverId)
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

# **InventoryControllerAddSubnetPoolToResourcePool**
> InventoryControllerAddSubnetPoolToResourcePool(ctx, resourcePoolId, subnetPoolId)
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

# **InventoryControllerCreateResourcePool**
> ResourcePoolDto InventoryControllerCreateResourcePool(ctx, body)
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

# **InventoryControllerDeleteResourcePool**
> InventoryControllerDeleteResourcePool(ctx, resourcePoolId)
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

# **InventoryControllerGetResourcePool**
> ResourcePoolDto InventoryControllerGetResourcePool(ctx, resourcePoolId)
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

# **InventoryControllerGetResourcePoolServers**
> InventoryControllerGetResourcePoolServers(ctx, resourcePoolId)
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

# **InventoryControllerGetResourcePoolSubnetPools**
> InventoryControllerGetResourcePoolSubnetPools(ctx, resourcePoolId)
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

# **InventoryControllerGetResourcePoolUsers**
> InventoryControllerGetResourcePoolUsers(ctx, resourcePoolId)
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

# **InventoryControllerGetResourcePools**
> []ResourcePoolWithStatsDto InventoryControllerGetResourcePools(ctx, )
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

# **InventoryControllerGetUserResourcePools**
> InventoryControllerGetUserResourcePools(ctx, userId)
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

# **InventoryControllerRemoveResourcePoolUser**
> InventoryControllerRemoveResourcePoolUser(ctx, resourcePoolId, userId)
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

# **InventoryControllerRemoveServerFromResourcePool**
> InventoryControllerRemoveServerFromResourcePool(ctx, resourcePoolId, serverId)
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

# **InventoryControllerRemoveSubnetPoolFromResourcePool**
> InventoryControllerRemoveSubnetPoolFromResourcePool(ctx, resourcePoolId, subnetPoolId)
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

# **InventoryControllerUpdateResourcePool**
> ResourcePoolDto InventoryControllerUpdateResourcePool(ctx, body, resourcePoolId)
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

