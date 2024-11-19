# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyVMTypeOnVMInstanceGroup**](VMInstanceGroupApi.md#ApplyVMTypeOnVMInstanceGroup) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/actions/apply-type/{vmTypeId} | Applies a VM Type to a VM Instance Group
[**CreateVMInstanceGroup**](VMInstanceGroupApi.md#CreateVMInstanceGroup) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups | Creates a VM Instance Group
[**CreateVMInterfaceOnVMInstanceGroup**](VMInstanceGroupApi.md#CreateVMInterfaceOnVMInstanceGroup) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/interfaces | Creates a new Virtual Interface for the VM Instance Group
[**DeleteVMInstanceGroup**](VMInstanceGroupApi.md#DeleteVMInstanceGroup) | **Delete** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId} | Deletes a VM Instance Group
[**GetVMInstanceGroup**](VMInstanceGroupApi.md#GetVMInstanceGroup) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId} | Get VM Instance Group information
[**GetVMInstanceGroupVMInstances**](VMInstanceGroupApi.md#GetVMInstanceGroupVMInstances) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/vm-instances | Get the VM Instances of VM Instance Group
[**GetVMInstanceGroups**](VMInstanceGroupApi.md#GetVMInstanceGroups) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups | Get all VM Instance Groups
[**UpdateNetworkProfileOnVMInstanceGroupNetwork**](VMInstanceGroupApi.md#UpdateNetworkProfileOnVMInstanceGroupNetwork) | **Patch** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/network/{networkId} | Applies the given Network Profile to the specified VM Instance Group Network
[**UpdateVMInstanceGroup**](VMInstanceGroupApi.md#UpdateVMInstanceGroup) | **Patch** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId} | Updates VM Instance Group information

# **ApplyVMTypeOnVMInstanceGroup**
> VmInstanceGroup ApplyVMTypeOnVMInstanceGroup(ctx, infrastructureId, vmInstanceGroupId, vmTypeId)
Applies a VM Type to a VM Instance Group

Applies a VM Type to a VM Instance Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 
  **vmInstanceGroupId** | **float64**|  | 
  **vmTypeId** | **float64**|  | 

### Return type

[**VmInstanceGroup**](VMInstanceGroup.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVMInstanceGroup**
> VmInstanceGroup CreateVMInstanceGroup(ctx, body, infrastructureId)
Creates a VM Instance Group

Creates a VM Instance Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateVmInstanceGroup**](CreateVmInstanceGroup.md)| The VM Instance Group create object | 
  **infrastructureId** | **float64**|  | 

### Return type

[**VmInstanceGroup**](VMInstanceGroup.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVMInterfaceOnVMInstanceGroup**
> VmInstanceGroupInterface CreateVMInterfaceOnVMInstanceGroup(ctx, body, infrastructureId, vmInstanceGroupId)
Creates a new Virtual Interface for the VM Instance Group

Creates a new Virtual Interface for the VM Instance Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateVmInstanceGroupInterface**](CreateVmInstanceGroupInterface.md)|  | 
  **infrastructureId** | **float64**|  | 
  **vmInstanceGroupId** | **float64**|  | 

### Return type

[**VmInstanceGroupInterface**](VMInstanceGroupInterface.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVMInstanceGroup**
> DeleteVMInstanceGroup(ctx, infrastructureId, vmInstanceGroupId)
Deletes a VM Instance Group

Deletes a VM Instance Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 
  **vmInstanceGroupId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVMInstanceGroup**
> VmInstanceGroup GetVMInstanceGroup(ctx, infrastructureId, vmInstanceGroupId)
Get VM Instance Group information

Returns VM Instance Group information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 
  **vmInstanceGroupId** | **float64**|  | 

### Return type

[**VmInstanceGroup**](VMInstanceGroup.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVMInstanceGroupVMInstances**
> map[string]VmInstance GetVMInstanceGroupVMInstances(ctx, infrastructureId, vmInstanceGroupId)
Get the VM Instances of VM Instance Group

Returns the VM Instances of VM Instance Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 
  **vmInstanceGroupId** | **float64**|  | 

### Return type

[**map[string]VmInstance**](VMInstance.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVMInstanceGroups**
> []VmInstanceGroup GetVMInstanceGroups(ctx, infrastructureId)
Get all VM Instance Groups

Returns list of all VM Instance Groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 

### Return type

[**[]VmInstanceGroup**](VMInstanceGroup.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNetworkProfileOnVMInstanceGroupNetwork**
> VmInstanceGroup UpdateNetworkProfileOnVMInstanceGroupNetwork(ctx, body, infrastructureId, vmInstanceGroupId, networkId)
Applies the given Network Profile to the specified VM Instance Group Network

Applies the given Network Profile to the specified VM Instance Group Network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateVmInstanceGroupNetwork**](UpdateVmInstanceGroupNetwork.md)| The VM Instance Group Network update object | 
  **infrastructureId** | **float64**|  | 
  **vmInstanceGroupId** | **float64**|  | 
  **networkId** | **float64**|  | 

### Return type

[**VmInstanceGroup**](VMInstanceGroup.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVMInstanceGroup**
> VmInstanceGroup UpdateVMInstanceGroup(ctx, body, infrastructureId, vmInstanceGroupId)
Updates VM Instance Group information

Updates VM Instance Group information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateVmInstanceGroup**](UpdateVmInstanceGroup.md)| The VM Instance Group update object | 
  **infrastructureId** | **float64**|  | 
  **vmInstanceGroupId** | **float64**|  | 

### Return type

[**VmInstanceGroup**](VMInstanceGroup.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

