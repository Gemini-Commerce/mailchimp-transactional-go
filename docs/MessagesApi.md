# \MessagesApi

All URIs are relative to *https://mandrillapp.com/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostMessagesCancelScheduled**](MessagesApi.md#PostMessagesCancelScheduled) | **Post** /messages/cancel-scheduled | Cancel scheduled email
[**PostMessagesContent**](MessagesApi.md#PostMessagesContent) | **Post** /messages/content | Get message content
[**PostMessagesInfo**](MessagesApi.md#PostMessagesInfo) | **Post** /messages/info | Get message info
[**PostMessagesListScheduled**](MessagesApi.md#PostMessagesListScheduled) | **Post** /messages/list-scheduled | List scheduled emails
[**PostMessagesParse**](MessagesApi.md#PostMessagesParse) | **Post** /messages/parse | Parse mime document
[**PostMessagesReschedule**](MessagesApi.md#PostMessagesReschedule) | **Post** /messages/reschedule | Reschedule email
[**PostMessagesSearch**](MessagesApi.md#PostMessagesSearch) | **Post** /messages/search | Search messages by date
[**PostMessagesSearchTimeSeries**](MessagesApi.md#PostMessagesSearchTimeSeries) | **Post** /messages/search-time-series | Search messages by hour
[**PostMessagesSend**](MessagesApi.md#PostMessagesSend) | **Post** /messages/send | Send new message
[**PostMessagesSendRaw**](MessagesApi.md#PostMessagesSendRaw) | **Post** /messages/send-raw | Send mime document
[**PostMessagesSendTemplate**](MessagesApi.md#PostMessagesSendTemplate) | **Post** /messages/send-template | Send using message template


# **PostMessagesCancelScheduled**
> []InlineResponse20035 PostMessagesCancelScheduled(ctx, body)
Cancel scheduled email

Cancels a scheduled email.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body40**](Body40.md)|  | 

### Return type

[**[]InlineResponse20035**](inline_response_200_35.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMessagesContent**
> InlineResponse20033 PostMessagesContent(ctx, body)
Get message content

Get the full content of a recently sent message.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body36**](Body36.md)|  | 

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMessagesInfo**
> InlineResponse20032 PostMessagesInfo(ctx, body)
Get message info

Get the information for a single recently sent message.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body35**](Body35.md)|  | 

### Return type

[**InlineResponse20032**](inline_response_200_32.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMessagesListScheduled**
> []InlineResponse20035 PostMessagesListScheduled(ctx, body)
List scheduled emails

Queries your scheduled emails.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body39**](Body39.md)|  | 

### Return type

[**[]InlineResponse20035**](inline_response_200_35.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMessagesParse**
> InlineResponse20034 PostMessagesParse(ctx, body)
Parse mime document

Parse the full MIME document for an email message, returning the content of the message broken into its constituent pieces.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body37**](Body37.md)|  | 

### Return type

[**InlineResponse20034**](inline_response_200_34.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMessagesReschedule**
> []InlineResponse20035 PostMessagesReschedule(ctx, body)
Reschedule email

Reschedules a scheduled email.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body41**](Body41.md)|  | 

### Return type

[**[]InlineResponse20035**](inline_response_200_35.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMessagesSearch**
> []InlineResponse20030 PostMessagesSearch(ctx, body)
Search messages by date

Search recently sent messages and optionally narrow by date range, tags, senders, and API keys. If no date range is specified, results within the last 7 days are returned. This method may be called up to 20 times per minute. If you need the data more often, you can use /messages/info.json to get the information for a single message, or webhooks to push activity to your own application for querying.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body33**](Body33.md)|  | 

### Return type

[**[]InlineResponse20030**](inline_response_200_30.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMessagesSearchTimeSeries**
> []InlineResponse20031 PostMessagesSearchTimeSeries(ctx, body)
Search messages by hour

Search the content of recently sent messages and return the aggregated hourly stats for matching messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body34**](Body34.md)|  | 

### Return type

[**[]InlineResponse20031**](inline_response_200_31.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMessagesSend**
> []InlineResponse20028 PostMessagesSend(ctx, body)
Send new message

Send a new transactional message through the Transactional API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body31**](Body31.md)|  | 

### Return type

[**[]InlineResponse20028**](inline_response_200_28.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMessagesSendRaw**
> interface{} PostMessagesSendRaw(ctx, body)
Send mime document

Take a raw MIME document for a message, and send it exactly as if it were sent through the Transactional API's SMTP servers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body38**](Body38.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMessagesSendTemplate**
> []InlineResponse20029 PostMessagesSendTemplate(ctx, body)
Send using message template

Send a new transactional message through the Transactional API using a template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body32**](Body32.md)|  | 

### Return type

[**[]InlineResponse20029**](inline_response_200_29.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

