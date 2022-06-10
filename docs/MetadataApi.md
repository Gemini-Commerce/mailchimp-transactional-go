# {{classname}}

All URIs are relative to *https://mandrillapp.com/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostMetadataAdd**](MetadataApi.md#PostMetadataAdd) | **Post** /metadata/add | Add metadata field
[**PostMetadataDelete**](MetadataApi.md#PostMetadataDelete) | **Post** /metadata/delete | Delete metadata field
[**PostMetadataList**](MetadataApi.md#PostMetadataList) | **Post** /metadata/list | List metadata fields
[**PostMetadataUpdate**](MetadataApi.md#PostMetadataUpdate) | **Post** /metadata/update | Update metadata field

# **PostMetadataAdd**
> InlineResponse20037 PostMetadataAdd(ctx, body)
Add metadata field

Add a new custom metadata field to be indexed for the account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MetadataAddBody**](MetadataAddBody.md)|  | 

### Return type

[**InlineResponse20037**](inline_response_200_37.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMetadataDelete**
> InlineResponse20039 PostMetadataDelete(ctx, body)
Delete metadata field

Delete an existing custom metadata field. Deletion isn't instataneous, and /metadata/list will continue to return the field until the asynchronous deletion process is complete.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MetadataDeleteBody**](MetadataDeleteBody.md)|  | 

### Return type

[**InlineResponse20039**](inline_response_200_39.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMetadataList**
> []InlineResponse20036 PostMetadataList(ctx, body)
List metadata fields

Get the list of custom metadata fields indexed for the account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MetadataListBody**](MetadataListBody.md)|  | 

### Return type

[**[]InlineResponse20036**](inline_response_200_36.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMetadataUpdate**
> InlineResponse20038 PostMetadataUpdate(ctx, body)
Update metadata field

Update an existing custom metadata field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MetadataUpdateBody**](MetadataUpdateBody.md)|  | 

### Return type

[**InlineResponse20038**](inline_response_200_38.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

