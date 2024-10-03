# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**SwitchControllerChangeStatus**](NetworkDevicesApi.md#SwitchControllerChangeStatus) | **Patch** /api/v2/network-devices/{networkDeviceId}/actions/change-status | Change status of a network device
[**SwitchControllerDiscover**](NetworkDevicesApi.md#SwitchControllerDiscover) | **Post** /api/v2/network-devices/{networkDeviceId}/discover | Discover network device interfaces, hardware and software configuration
[**SwitchControllerEnableSyslog**](NetworkDevicesApi.md#SwitchControllerEnableSyslog) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/syslog-subscribe | Enables remote syslog for a network device
[**SwitchControllerGetPorts**](NetworkDevicesApi.md#SwitchControllerGetPorts) | **Get** /api/v2/network-devices/{networkDeviceId}/ports | Get all ports for network device
[**SwitchControllerResetNetworkDevice**](NetworkDevicesApi.md#SwitchControllerResetNetworkDevice) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/reset | Resets a network device to default state
[**SwitchControllerSetPortStatus**](NetworkDevicesApi.md#SwitchControllerSetPortStatus) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/set-port-status | Set port status

# **SwitchControllerChangeStatus**
> SwitchControllerChangeStatus(ctx, networkDeviceId)
Change status of a network device

Change status of a network device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkDeviceId** | **float64**| Network device ID | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwitchControllerDiscover**
> SwitchControllerDiscover(ctx, body, networkDeviceId)
Discover network device interfaces, hardware and software configuration

Discover network device interfaces, hardware and software configuration and return them and/or persist them

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DiscoveryQueryDto**](DiscoveryQueryDto.md)|  | 
  **networkDeviceId** | **float64**| Network device identifier | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwitchControllerEnableSyslog**
> SwitchControllerEnableSyslog(ctx, networkDeviceId)
Enables remote syslog for a network device

Enables remote syslog for a network device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkDeviceId** | **float64**| Network device ID | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwitchControllerGetPorts**
> SwitchControllerGetPorts(ctx, networkDeviceId)
Get all ports for network device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkDeviceId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwitchControllerResetNetworkDevice**
> SwitchControllerResetNetworkDevice(ctx, networkDeviceId)
Resets a network device to default state

Resets a network device to default state and destroy all configurations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkDeviceId** | **float64**| Network device ID | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwitchControllerSetPortStatus**
> SwitchControllerSetPortStatus(ctx, body, networkDeviceId)
Set port status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NetworkDevicePortStatusDto**](NetworkDevicePortStatusDto.md)| Port status | 
  **networkDeviceId** | **float64**| Network device ID | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

