# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnableServerSyslog**](ServerApi.md#EnableServerSyslog) | **Post** /api/v2/servers/{serverId}/actions/syslog-subscribe | Enables remote syslog for a server
[**GetServerInfo**](ServerApi.md#GetServerInfo) | **Get** /api/v2/servers/{serverId} | Get Server information
[**GetServerPowerState**](ServerApi.md#GetServerPowerState) | **Post** /api/v2/servers/{serverId}/actions/get-power | Gets the power state of a server
[**GetServerRemoteConsoleInfo**](ServerApi.md#GetServerRemoteConsoleInfo) | **Get** /api/v2/servers/{serverId}/remote-console-info | Get Remote Console information
[**GetServerVNCInfo**](ServerApi.md#GetServerVNCInfo) | **Get** /api/v2/servers/{serverId}/vnc-info | Get VNC information
[**GetServers**](ServerApi.md#GetServers) | **Get** /api/v2/servers | Get a list of Servers
[**ReRegisterServer**](ServerApi.md#ReRegisterServer) | **Post** /api/v2/servers/{serverId}/actions/re-register | Re-register a server
[**RegisterServer**](ServerApi.md#RegisterServer) | **Post** /api/v2/servers | Initialize server registration
[**ResetServerToFactoryDefaults**](ServerApi.md#ResetServerToFactoryDefaults) | **Post** /api/v2/servers/{serverId}/actions/factory-reset | Resets a server to factory defaults
[**SetServerPowerState**](ServerApi.md#SetServerPowerState) | **Post** /api/v2/servers/{serverId}/actions/set-power | Sets the power state of a server

# **EnableServerSyslog**
> EnableServerSyslog(ctx, serverId)
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

# **GetServerInfo**
> Server GetServerInfo(ctx, serverId)
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

# **GetServerPowerState**
> GetServerPowerState(ctx, serverId)
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

# **GetServerRemoteConsoleInfo**
> []RemoteConsoleInfoDto GetServerRemoteConsoleInfo(ctx, serverId)
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

# **GetServerVNCInfo**
> []ServerVncInfoDto GetServerVNCInfo(ctx, serverId)
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

# **GetServers**
> ServerList GetServers(ctx, )
Get a list of Servers

Returns a list of Servers

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ServerList**](ServerList.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReRegisterServer**
> ServerReRegistrationResponseDto ReRegisterServer(ctx, serverId)
Re-register a server

Re-registers a server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **float64**|  | 

### Return type

[**ServerReRegistrationResponseDto**](ServerReRegistrationResponseDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterServer**
> ServerRegistrationResponseDto RegisterServer(ctx, body)
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

# **ResetServerToFactoryDefaults**
> ResetServerToFactoryDefaults(ctx, serverId)
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

# **SetServerPowerState**
> SetServerPowerState(ctx, body, serverId)
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

