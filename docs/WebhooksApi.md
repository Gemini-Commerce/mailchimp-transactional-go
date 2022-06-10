# {{classname}}

All URIs are relative to *https://mandrillapp.com/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostWebhooksAdd**](WebhooksApi.md#PostWebhooksAdd) | **Post** /webhooks/add | Add webhook
[**PostWebhooksDelete**](WebhooksApi.md#PostWebhooksDelete) | **Post** /webhooks/delete | Delete webhook
[**PostWebhooksInfo**](WebhooksApi.md#PostWebhooksInfo) | **Post** /webhooks/info | Get webhook info
[**PostWebhooksList**](WebhooksApi.md#PostWebhooksList) | **Post** /webhooks/list | List webhooks
[**PostWebhooksUpdate**](WebhooksApi.md#PostWebhooksUpdate) | **Post** /webhooks/update | Update webhook

# **PostWebhooksAdd**
> InlineResponse20075 PostWebhooksAdd(ctx, body)
Add webhook

Add a new webhook.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WebhooksAddBody**](WebhooksAddBody.md)|  | 

### Return type

[**InlineResponse20075**](inline_response_200_75.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostWebhooksDelete**
> InlineResponse20078 PostWebhooksDelete(ctx, body)
Delete webhook

Delete an existing webhook.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WebhooksDeleteBody**](WebhooksDeleteBody.md)|  | 

### Return type

[**InlineResponse20078**](inline_response_200_78.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostWebhooksInfo**
> InlineResponse20076 PostWebhooksInfo(ctx, body)
Get webhook info

Given the ID of an existing webhook, return the data about it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WebhooksInfoBody**](WebhooksInfoBody.md)|  | 

### Return type

[**InlineResponse20076**](inline_response_200_76.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostWebhooksList**
> []InlineResponse20074 PostWebhooksList(ctx, body)
List webhooks

Get the list of all webhooks defined on the account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WebhooksListBody**](WebhooksListBody.md)|  | 

### Return type

[**[]InlineResponse20074**](inline_response_200_74.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostWebhooksUpdate**
> InlineResponse20077 PostWebhooksUpdate(ctx, body)
Update webhook

Update an existing webhook.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WebhooksUpdateBody**](WebhooksUpdateBody.md)|  | 

### Return type

[**InlineResponse20077**](inline_response_200_77.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

