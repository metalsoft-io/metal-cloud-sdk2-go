# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryControllerGetServerInfo**](ServerApi.md#InventoryControllerGetServerInfo) | **Get** /api/v2/servers/{serverId} | Get Server information
[**InventoryControllerRegisterServer**](ServerApi.md#InventoryControllerRegisterServer) | **Post** /api/v2/servers | Initialize server registration
[**ServersControllerEnableSyslog**](ServerApi.md#ServersControllerEnableSyslog) | **Post** /api/v2/servers/{serverId}/actions/syslog-subscribe | Enables remote syslog for a server
[**ServersControllerGetPowerState**](ServerApi.md#ServersControllerGetPowerState) | **Post** /api/v2/servers/{serverId}/actions/get-power | Gets the power state of a server
[**ServersControllerGetRemoteConsoleInfo**](ServerApi.md#ServersControllerGetRemoteConsoleInfo) | **Get** /api/v2/servers/{serverId}/remote-console-info | Get Remote Console information
[**ServersControllerGetVNCInfo**](ServerApi.md#ServersControllerGetVNCInfo) | **Get** /api/v2/servers/{serverId}/vnc-info | Get VNC information
[**ServersControllerResetServerToFactoryDefaults**](ServerApi.md#ServersControllerResetServerToFactoryDefaults) | **Post** /api/v2/servers/{serverId}/actions/factory-reset | Resets a server to factory defaults
[**ServersControllerSetPowerState**](ServerApi.md#ServersControllerSetPowerState) | **Post** /api/v2/servers/{serverId}/actions/set-power | Sets the power state of a server

# **InventoryControllerGetServerInfo**
> Server InventoryControllerGetServerInfo(ctx, serverId)
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

# **InventoryControllerRegisterServer**
> ServerRegistrationResponseDto InventoryControllerRegisterServer(ctx, body)
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

# **ServersControllerEnableSyslog**
> ServersControllerEnableSyslog(ctx, serverId)
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

# **ServersControllerGetPowerState**
> ServersControllerGetPowerState(ctx, serverId)
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

# **ServersControllerGetRemoteConsoleInfo**
> []RemoteConsoleInfoDto ServersControllerGetRemoteConsoleInfo(ctx, serverId)
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

# **ServersControllerGetVNCInfo**
> []ServerVncInfoDto ServersControllerGetVNCInfo(ctx, serverId)
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

# **ServersControllerResetServerToFactoryDefaults**
> ServersControllerResetServerToFactoryDefaults(ctx, serverId)
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

# **ServersControllerSetPowerState**
> ServersControllerSetPowerState(ctx, body, serverId)
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

