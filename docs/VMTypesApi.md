# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryControllerCreateVMType**](VMTypesApi.md#InventoryControllerCreateVMType) | **Post** /api/v2/vm-types | Creates a VM Type
[**InventoryControllerDeleteVMType**](VMTypesApi.md#InventoryControllerDeleteVMType) | **Delete** /api/v2/vm-types/{vmTypeId} | Deletes a VM Type
[**InventoryControllerGetVMType**](VMTypesApi.md#InventoryControllerGetVMType) | **Get** /api/v2/vm-types/{vmTypeId} | Get VM Type information
[**InventoryControllerGetVMTypes**](VMTypesApi.md#InventoryControllerGetVMTypes) | **Get** /api/v2/vm-types | Get all VM Types
[**InventoryControllerGetVMsByVMType**](VMTypesApi.md#InventoryControllerGetVMsByVMType) | **Get** /api/v2/vm-types/{vmTypeId}/vms | Returns all VMs linked to the VM Type
[**InventoryControllerUpdateVMType**](VMTypesApi.md#InventoryControllerUpdateVMType) | **Patch** /api/v2/vm-types/{vmTypeId} | Updates VM Type information

# **InventoryControllerCreateVMType**
> VmType InventoryControllerCreateVMType(ctx, body)
Creates a VM Type

Creates a VM Type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateVmType**](CreateVmType.md)| The VM Type create object | 

### Return type

[**VmType**](VMType.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryControllerDeleteVMType**
> InventoryControllerDeleteVMType(ctx, vmTypeId)
Deletes a VM Type

Deletes a VM Type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmTypeId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryControllerGetVMType**
> VmType InventoryControllerGetVMType(ctx, vmTypeId)
Get VM Type information

Returns VM Type information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmTypeId** | **float64**|  | 

### Return type

[**VmType**](VMType.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryControllerGetVMTypes**
> VmTypeList InventoryControllerGetVMTypes(ctx, )
Get all VM Types

Returns list of all VM Types

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**VmTypeList**](VMTypeList.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryControllerGetVMsByVMType**
> InventoryControllerGetVMsByVMType(ctx, vmTypeId)
Returns all VMs linked to the VM Type

Returns all VMs linked to the VM Type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmTypeId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryControllerUpdateVMType**
> VmType InventoryControllerUpdateVMType(ctx, body, vmTypeId)
Updates VM Type information

Updates VM Type information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateVmType**](UpdateVmType.md)| The VM Type update object | 
  **vmTypeId** | **float64**|  | 

### Return type

[**VmType**](VMType.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

