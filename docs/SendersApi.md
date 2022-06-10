# {{classname}}

All URIs are relative to *https://mandrillapp.com/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostSendersAddDomain**](SendersApi.md#PostSendersAddDomain) | **Post** /senders/add-domain | Add sender domain
[**PostSendersCheckDomain**](SendersApi.md#PostSendersCheckDomain) | **Post** /senders/check-domain | Check domain settings
[**PostSendersDomains**](SendersApi.md#PostSendersDomains) | **Post** /senders/domains | List sender domains
[**PostSendersInfo**](SendersApi.md#PostSendersInfo) | **Post** /senders/info | Get sender info
[**PostSendersList**](SendersApi.md#PostSendersList) | **Post** /senders/list | List account senders
[**PostSendersTimeSeries**](SendersApi.md#PostSendersTimeSeries) | **Post** /senders/time-series | View sender history
[**PostSendersVerifyDomain**](SendersApi.md#PostSendersVerifyDomain) | **Post** /senders/verify-domain | Verify domain

# **PostSendersAddDomain**
> InlineResponse20045 PostSendersAddDomain(ctx, body)
Add sender domain

Adds a sender domain to your account. Sender domains are added automatically as you send, but you can use this call to add them ahead of time.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SendersAdddomainBody**](SendersAdddomainBody.md)|  | 

### Return type

[**InlineResponse20045**](inline_response_200_45.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSendersCheckDomain**
> InlineResponse20046 PostSendersCheckDomain(ctx, body)
Check domain settings

Checks the SPF and DKIM settings for a domain, as well the domain verification. If you haven't already added this domain to your account, it will be added automatically.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SendersCheckdomainBody**](SendersCheckdomainBody.md)|  | 

### Return type

[**InlineResponse20046**](inline_response_200_46.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSendersDomains**
> []InlineResponse20044 PostSendersDomains(ctx, body)
List sender domains

Returns the sender domains that have been added to this account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SendersDomainsBody**](SendersDomainsBody.md)|  | 

### Return type

[**[]InlineResponse20044**](inline_response_200_44.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSendersInfo**
> InlineResponse20048 PostSendersInfo(ctx, body)
Get sender info

Return more detailed information about a single sender, including aggregates of recent stats.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SendersInfoBody**](SendersInfoBody.md)|  | 

### Return type

[**InlineResponse20048**](inline_response_200_48.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSendersList**
> []InlineResponse20043 PostSendersList(ctx, body)
List account senders

Return the senders that have tried to use this account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SendersListBody**](SendersListBody.md)|  | 

### Return type

[**[]InlineResponse20043**](inline_response_200_43.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSendersTimeSeries**
> []InlineResponse20049 PostSendersTimeSeries(ctx, body)
View sender history

Return the recent history (hourly stats for the last 30 days) for a sender.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SendersTimeseriesBody**](SendersTimeseriesBody.md)|  | 

### Return type

[**[]InlineResponse20049**](inline_response_200_49.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSendersVerifyDomain**
> InlineResponse20047 PostSendersVerifyDomain(ctx, body)
Verify domain

Sends a verification email in order to verify ownership of a domain. Domain verification is a required step to confirm ownership of a domain. Once a domain has been verified in a Transactional API account, other accounts may not have their messages signed by that domain unless they also verify the domain. This prevents other Transactional API accounts from sending mail signed by your domain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SendersVerifydomainBody**](SendersVerifydomainBody.md)|  | 

### Return type

[**InlineResponse20047**](inline_response_200_47.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

