# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryController1GetServerInfo**](ServerApi.md#InventoryController1GetServerInfo) | **Get** /api/v2/servers/{serverId} | Get Server information
[**InventoryController1RegisterServer**](ServerApi.md#InventoryController1RegisterServer) | **Post** /api/v2/servers | Initialize server registration
[**ServersController1EnableSyslog**](ServerApi.md#ServersController1EnableSyslog) | **Post** /api/v2/servers/{serverId}/actions/syslog-subscribe | Enables remote syslog for a server
[**ServersController1GetPowerState**](ServerApi.md#ServersController1GetPowerState) | **Post** /api/v2/servers/{serverId}/actions/get-power | Gets the power state of a server
[**ServersController1GetRemoteConsoleInfo**](ServerApi.md#ServersController1GetRemoteConsoleInfo) | **Get** /api/v2/servers/{serverId}/remote-console-info | Get Remote Console information
[**ServersController1GetVNCInfo**](ServerApi.md#ServersController1GetVNCInfo) | **Get** /api/v2/servers/{serverId}/vnc-info | Get VNC information
[**ServersController1ResetServerToFactoryDefaults**](ServerApi.md#ServersController1ResetServerToFactoryDefaults) | **Post** /api/v2/servers/{serverId}/actions/factory-reset | Resets a server to factory defaults
[**ServersController1SetPowerState**](ServerApi.md#ServersController1SetPowerState) | **Post** /api/v2/servers/{serverId}/actions/set-power | Sets the power state of a server

# **InventoryController1GetServerInfo**
> Server InventoryController1GetServerInfo(ctx, serverId)
Get Server information

Returns Server information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **float64**|  | 

### Return type

[**Server**](Server.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryController1RegisterServer**
> ServerRegistrationResponseDto InventoryController1RegisterServer(ctx, body)
Initialize server registration

Initializes server registration process

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerRegistrationDto**](ServerRegistrationDto.md)| The server registration information | 

### Return type

[**ServerRegistrationResponseDto**](ServerRegistrationResponseDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServersController1EnableSyslog**
> ServersController1EnableSyslog(ctx, serverId)
Enables remote syslog for a server

Enables remote syslog for a server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServersController1GetPowerState**
> ServersController1GetPowerState(ctx, serverId)
Gets the power state of a server

Gets the power state of a server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServersController1GetRemoteConsoleInfo**
> []RemoteConsoleInfoDto ServersController1GetRemoteConsoleInfo(ctx, serverId)
Get Remote Console information

Returns Remote Console information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **float64**|  | 

### Return type

[**[]RemoteConsoleInfoDto**](RemoteConsoleInfoDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServersController1GetVNCInfo**
> []ServerVncInfoDto ServersController1GetVNCInfo(ctx, serverId)
Get VNC information

Returns VNC information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **float64**|  | 

### Return type

[**[]ServerVncInfoDto**](ServerVNCInfoDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServersController1ResetServerToFactoryDefaults**
> ServersController1ResetServerToFactoryDefaults(ctx, serverId)
Resets a server to factory defaults

Resets a server to factory defaults

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServersController1SetPowerState**
> ServersController1SetPowerState(ctx, body, serverId)
Sets the power state of a server

Sets the power state of a server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerPowerSetDto**](ServerPowerSetDto.md)|  | 
  **serverId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

