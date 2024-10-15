# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateEliResponse**](AIApi.md#GenerateEliResponse) | **Post** /api/v2/ai/generate | Request from AI a response for the given input

# **GenerateEliResponse**
> AiGenerateResponseDto GenerateEliResponse(ctx, body)
Request from AI a response for the given input

Returns response with potential actions to execute

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AiGenerateRequestDto**](AiGenerateRequestDto.md)| The AI generate request | 

### Return type

[**AiGenerateResponseDto**](AIGenerateResponseDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

