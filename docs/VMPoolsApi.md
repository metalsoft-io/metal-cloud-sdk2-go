# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryController1CreateVMPool**](VMPoolsApi.md#InventoryController1CreateVMPool) | **Post** /api/v2/vm-pools | Creates a VM Pool
[**InventoryController1DeleteVMPool**](VMPoolsApi.md#InventoryController1DeleteVMPool) | **Delete** /api/v2/vm-pools/{vmPoolId} | Deletes a VM Pool
[**InventoryController1GetVMPool**](VMPoolsApi.md#InventoryController1GetVMPool) | **Get** /api/v2/vm-pools/{vmPoolId} | Get VM Pool information
[**InventoryController1GetVMPoolClusterHost**](VMPoolsApi.md#InventoryController1GetVMPoolClusterHost) | **Get** /api/v2/vm-pools/{vmPoolId}/cluster-hosts/{vmPoolClusterHostId} | Retrieves a VM Cluster Host
[**InventoryController1GetVMPoolClusterHostInterface**](VMPoolsApi.md#InventoryController1GetVMPoolClusterHostInterface) | **Get** /api/v2/vm-pools/{vmPoolId}/cluster-hosts/{vmPoolClusterHostId}/interfaces/{vmPoolClusterHostInterfaceId} | Retrieves a VM Cluster Host Interface
[**InventoryController1GetVMPoolClusterHostInterfaces**](VMPoolsApi.md#InventoryController1GetVMPoolClusterHostInterfaces) | **Get** /api/v2/vm-pools/{vmPoolId}/cluster-hosts/{vmPoolClusterHostId}/interfaces | Retrieves a list of VM Cluster Host Interfaces
[**InventoryController1GetVMPoolClusterHostVMs**](VMPoolsApi.md#InventoryController1GetVMPoolClusterHostVMs) | **Get** /api/v2/vm-pools/{vmPoolId}/cluster-hosts/{vmPoolClusterHostId}/vms | Retrieves a list of VM Cluster Host VMs
[**InventoryController1GetVMPoolClusterHosts**](VMPoolsApi.md#InventoryController1GetVMPoolClusterHosts) | **Get** /api/v2/vm-pools/{vmPoolId}/cluster-hosts | Get list of VM Cluster Hosts linked to the VM Pool
[**InventoryController1GetVMPoolVMs**](VMPoolsApi.md#InventoryController1GetVMPoolVMs) | **Get** /api/v2/vm-pools/{vmPoolId}/vms | Returns all VMs linked to the VM Pool
[**InventoryController1GetVMPools**](VMPoolsApi.md#InventoryController1GetVMPools) | **Get** /api/v2/vm-pools | Get all VM Pools
[**InventoryController1PatchVMPoolClusterHostInterface**](VMPoolsApi.md#InventoryController1PatchVMPoolClusterHostInterface) | **Patch** /api/v2/vm-pools/{vmPoolId}/cluster-hosts/{vmPoolClusterHostId}/interfaces/{vmPoolClusterHostInterfaceId} | Updates a VM Cluster Host Interface
[**InventoryController1UpdateVMPool**](VMPoolsApi.md#InventoryController1UpdateVMPool) | **Patch** /api/v2/vm-pools/{vmPoolId} | Updates VM Pool information

# **InventoryController1CreateVMPool**
> VmPoolDto InventoryController1CreateVMPool(ctx, body)
Creates a VM Pool

Creates a VM Pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateVmPoolDto**](CreateVmPoolDto.md)| The VM Pool create object | 

### Return type

[**VmPoolDto**](VMPoolDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryController1DeleteVMPool**
> InventoryController1DeleteVMPool(ctx, vmPoolId)
Deletes a VM Pool

Deletes a VM Pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmPoolId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryController1GetVMPool**
> VmPoolDto InventoryController1GetVMPool(ctx, vmPoolId)
Get VM Pool information

Returns VM Pool information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmPoolId** | **float64**|  | 

### Return type

[**VmPoolDto**](VMPoolDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryController1GetVMPoolClusterHost**
> VmPoolHostsDto InventoryController1GetVMPoolClusterHost(ctx, vmPoolId, vmPoolClusterHostId)
Retrieves a VM Cluster Host

Returns a VM Cluster Host

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmPoolId** | **float64**|  | 
  **vmPoolClusterHostId** | **float64**|  | 

### Return type

[**VmPoolHostsDto**](VMPoolHostsDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryController1GetVMPoolClusterHostInterface**
> VmPoolHostInterfacesDto InventoryController1GetVMPoolClusterHostInterface(ctx, vmPoolId, vmPoolClusterHostId, vmPoolClusterHostInterfaceId)
Retrieves a VM Cluster Host Interface

Returns a VM Cluster Host Interface

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmPoolId** | **float64**|  | 
  **vmPoolClusterHostId** | **float64**|  | 
  **vmPoolClusterHostInterfaceId** | **float64**|  | 

### Return type

[**VmPoolHostInterfacesDto**](VMPoolHostInterfacesDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryController1GetVMPoolClusterHostInterfaces**
> []VmPoolHostInterfacesDto InventoryController1GetVMPoolClusterHostInterfaces(ctx, vmPoolId, vmPoolClusterHostId)
Retrieves a list of VM Cluster Host Interfaces

Returns a list of VM Cluster Host Interfaces

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmPoolId** | **float64**|  | 
  **vmPoolClusterHostId** | **float64**|  | 

### Return type

[**[]VmPoolHostInterfacesDto**](VMPoolHostInterfacesDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryController1GetVMPoolClusterHostVMs**
> []VmDto InventoryController1GetVMPoolClusterHostVMs(ctx, vmPoolId, vmPoolClusterHostId)
Retrieves a list of VM Cluster Host VMs

Returns a list of VM Cluster Host VMs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmPoolId** | **float64**|  | 
  **vmPoolClusterHostId** | **float64**|  | 

### Return type

[**[]VmDto**](VMDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryController1GetVMPoolClusterHosts**
> []VmPoolHostsDto InventoryController1GetVMPoolClusterHosts(ctx, vmPoolId)
Get list of VM Cluster Hosts linked to the VM Pool

Returns list of VM Cluster Hosts linked to the VM Pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmPoolId** | **float64**|  | 

### Return type

[**[]VmPoolHostsDto**](VMPoolHostsDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryController1GetVMPoolVMs**
> InventoryController1GetVMPoolVMs(ctx, vmPoolId)
Returns all VMs linked to the VM Pool

Returns all VMs linked to the VM Pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmPoolId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryController1GetVMPools**
> []VmPoolDto InventoryController1GetVMPools(ctx, )
Get all VM Pools

Returns list of all VM Pools

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]VmPoolDto**](VMPoolDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryController1PatchVMPoolClusterHostInterface**
> VmPoolHostInterfacesDto InventoryController1PatchVMPoolClusterHostInterface(ctx, body, vmPoolId, vmPoolClusterHostId, vmPoolClusterHostInterfaceId)
Updates a VM Cluster Host Interface

Updates a VM Cluster Host Interface

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateVmPoolClusterHostInterfaceDto**](UpdateVmPoolClusterHostInterfaceDto.md)| The VM Pool Cluster Host Interface update object | 
  **vmPoolId** | **float64**|  | 
  **vmPoolClusterHostId** | **float64**|  | 
  **vmPoolClusterHostInterfaceId** | **float64**|  | 

### Return type

[**VmPoolHostInterfacesDto**](VMPoolHostInterfacesDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryController1UpdateVMPool**
> VmPoolDto InventoryController1UpdateVMPool(ctx, body, vmPoolId)
Updates VM Pool information

Updates VM Pool information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateVmPoolDto**](UpdateVmPoolDto.md)| The VM Pool update object | 
  **vmPoolId** | **float64**|  | 

### Return type

[**VmPoolDto**](VMPoolDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

