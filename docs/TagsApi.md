# {{classname}}

All URIs are relative to *https://mandrillapp.com/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostTagsAllTimeSeries**](TagsApi.md#PostTagsAllTimeSeries) | **Post** /tags/all-time-series | View all tags history
[**PostTagsDelete**](TagsApi.md#PostTagsDelete) | **Post** /tags/delete | Delete tag
[**PostTagsInfo**](TagsApi.md#PostTagsInfo) | **Post** /tags/info | Get tag info
[**PostTagsList**](TagsApi.md#PostTagsList) | **Post** /tags/list | List tags
[**PostTagsTimeSeries**](TagsApi.md#PostTagsTimeSeries) | **Post** /tags/time-series | View tag history

# **PostTagsAllTimeSeries**
> []InlineResponse20031 PostTagsAllTimeSeries(ctx, body)
View all tags history

Return the recent history (hourly stats for the last 30 days) for all tags.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TagsAlltimeseriesBody**](TagsAlltimeseriesBody.md)|  | 

### Return type

[**[]InlineResponse20031**](inline_response_200_31.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTagsDelete**
> InlineResponse20058 PostTagsDelete(ctx, body)
Delete tag

Deletes a tag permanently. Deleting a tag removes the tag from any messages that have been sent, and also deletes the tag's stats. There is no way to undo this operation, so use it carefully.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TagsDeleteBody**](TagsDeleteBody.md)|  | 

### Return type

[**InlineResponse20058**](inline_response_200_58.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTagsInfo**
> InlineResponse20059 PostTagsInfo(ctx, body)
Get tag info

Return more detailed information about a single tag, including aggregates of recent stats.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TagsInfoBody**](TagsInfoBody.md)|  | 

### Return type

[**InlineResponse20059**](inline_response_200_59.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTagsList**
> []InlineResponse20057 PostTagsList(ctx, body)
List tags

Return all of the user-defined tag information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TagsListBody**](TagsListBody.md)|  | 

### Return type

[**[]InlineResponse20057**](inline_response_200_57.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTagsTimeSeries**
> []InlineResponse20031 PostTagsTimeSeries(ctx, body)
View tag history

Return the recent history (hourly stats for the last 30 days) for a tag.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TagsTimeseriesBody**](TagsTimeseriesBody.md)|  | 

### Return type

[**[]InlineResponse20031**](inline_response_200_31.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

