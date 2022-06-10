# {{classname}}

All URIs are relative to *https://mandrillapp.com/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostUrlsAddTrackingDomain**](UrlsApi.md#PostUrlsAddTrackingDomain) | **Post** /urls/add-tracking-domain | Add tracking domains
[**PostUrlsCheckTrackingDomain**](UrlsApi.md#PostUrlsCheckTrackingDomain) | **Post** /urls/check-tracking-domain | Check cname settings
[**PostUrlsList**](UrlsApi.md#PostUrlsList) | **Post** /urls/list | List most clicked urls
[**PostUrlsSearchDeprecated**](UrlsApi.md#PostUrlsSearchDeprecated) | **Post** /urls/search | Search most clicked urls
[**PostUrlsTimeSeries**](UrlsApi.md#PostUrlsTimeSeries) | **Post** /urls/time-series | Get url history
[**PostUrlsTrackingDomains**](UrlsApi.md#PostUrlsTrackingDomains) | **Post** /urls/tracking-domains | List tracking domains

# **PostUrlsAddTrackingDomain**
> InlineResponse20071 PostUrlsAddTrackingDomain(ctx, body)
Add tracking domains

Add a tracking domain to your account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UrlsAddtrackingdomainBody**](UrlsAddtrackingdomainBody.md)|  | 

### Return type

[**InlineResponse20071**](inline_response_200_71.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUrlsCheckTrackingDomain**
> InlineResponse20071 PostUrlsCheckTrackingDomain(ctx, body)
Check cname settings

Checks the CNAME settings for a tracking domain. The domain must have been added already with the add-tracking-domain call.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UrlsChecktrackingdomainBody**](UrlsChecktrackingdomainBody.md)|  | 

### Return type

[**InlineResponse20071**](inline_response_200_71.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUrlsList**
> []InlineResponse20067 PostUrlsList(ctx, body)
List most clicked urls

Get the 100 most clicked URLs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UrlsListBody**](UrlsListBody.md)|  | 

### Return type

[**[]InlineResponse20067**](inline_response_200_67.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUrlsSearchDeprecated**
> []InlineResponse20068 PostUrlsSearchDeprecated(ctx, body)
Search most clicked urls

Return the 100 most clicked URLs that match the search query given.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UrlsSearchBody**](UrlsSearchBody.md)|  | 

### Return type

[**[]InlineResponse20068**](inline_response_200_68.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUrlsTimeSeries**
> []InlineResponse20069 PostUrlsTimeSeries(ctx, body)
Get url history

Return the recent history (hourly stats for the last 30 days) for a URL

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UrlsTimeseriesBody**](UrlsTimeseriesBody.md)|  | 

### Return type

[**[]InlineResponse20069**](inline_response_200_69.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUrlsTrackingDomains**
> []InlineResponse20070 PostUrlsTrackingDomains(ctx, body)
List tracking domains

Get the list of tracking domains set up for this account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UrlsTrackingdomainsBody**](UrlsTrackingdomainsBody.md)|  | 

### Return type

[**[]InlineResponse20070**](inline_response_200_70.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

