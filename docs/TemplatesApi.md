# \TemplatesApi

All URIs are relative to *https://mandrillapp.com/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostTemplatesAdd**](TemplatesApi.md#PostTemplatesAdd) | **Post** /templates/add | Add template
[**PostTemplatesDelete**](TemplatesApi.md#PostTemplatesDelete) | **Post** /templates/delete | Delete template
[**PostTemplatesInfo**](TemplatesApi.md#PostTemplatesInfo) | **Post** /templates/info | Get template info
[**PostTemplatesList**](TemplatesApi.md#PostTemplatesList) | **Post** /templates/list | List templates
[**PostTemplatesPublish**](TemplatesApi.md#PostTemplatesPublish) | **Post** /templates/publish | Publish template content
[**PostTemplatesRender**](TemplatesApi.md#PostTemplatesRender) | **Post** /templates/render | Render html template
[**PostTemplatesTimeSeries**](TemplatesApi.md#PostTemplatesTimeSeries) | **Post** /templates/time-series | Get template history
[**PostTemplatesUpdate**](TemplatesApi.md#PostTemplatesUpdate) | **Post** /templates/update | Update template


# **PostTemplatesAdd**
> InlineResponse20060 PostTemplatesAdd(ctx, body)
Add template

Add a new template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body68**](Body68.md)|  | 

### Return type

[**InlineResponse20060**](inline_response_200_60.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTemplatesDelete**
> InlineResponse20064 PostTemplatesDelete(ctx, body)
Delete template

Delete a template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body72**](Body72.md)|  | 

### Return type

[**InlineResponse20064**](inline_response_200_64.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTemplatesInfo**
> InlineResponse20061 PostTemplatesInfo(ctx, body)
Get template info

Get the information for an existing template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body69**](Body69.md)|  | 

### Return type

[**InlineResponse20061**](inline_response_200_61.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTemplatesList**
> []InlineResponse20065 PostTemplatesList(ctx, body)
List templates

Return a list of all the templates available to this user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body73**](Body73.md)|  | 

### Return type

[**[]InlineResponse20065**](inline_response_200_65.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTemplatesPublish**
> InlineResponse20063 PostTemplatesPublish(ctx, body)
Publish template content

Publish the content for the template. Any new messages sent using this template will start using the content that was previously in draft.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body71**](Body71.md)|  | 

### Return type

[**InlineResponse20063**](inline_response_200_63.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTemplatesRender**
> InlineResponse20066 PostTemplatesRender(ctx, body)
Render html template

Inject content and optionally merge fields into a template, returning the HTML that results.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body75**](Body75.md)|  | 

### Return type

[**InlineResponse20066**](inline_response_200_66.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTemplatesTimeSeries**
> []InlineResponse20049 PostTemplatesTimeSeries(ctx, body)
Get template history

Return the recent history (hourly stats for the last 30 days) for a template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body74**](Body74.md)|  | 

### Return type

[**[]InlineResponse20049**](inline_response_200_49.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTemplatesUpdate**
> InlineResponse20062 PostTemplatesUpdate(ctx, body)
Update template

Update the code for an existing template. If null is provided for any fields, the values will remain unchanged.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body70**](Body70.md)|  | 

### Return type

[**InlineResponse20062**](inline_response_200_62.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

