# {{classname}}

All URIs are relative to *https://mandrillapp.com/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostInboundAddDomain**](InboundApi.md#PostInboundAddDomain) | **Post** /inbound/add-domain | Add inbound domain
[**PostInboundAddRoute**](InboundApi.md#PostInboundAddRoute) | **Post** /inbound/add-route | Add mailbox route
[**PostInboundCheckDomain**](InboundApi.md#PostInboundCheckDomain) | **Post** /inbound/check-domain | Check domain settings
[**PostInboundDeleteDomain**](InboundApi.md#PostInboundDeleteDomain) | **Post** /inbound/delete-domain | Delete inbound domain
[**PostInboundDeleteRoute**](InboundApi.md#PostInboundDeleteRoute) | **Post** /inbound/delete-route | Delete mailbox route
[**PostInboundDomains**](InboundApi.md#PostInboundDomains) | **Post** /inbound/domains | List inbound domains
[**PostInboundRoutes**](InboundApi.md#PostInboundRoutes) | **Post** /inbound/routes | List mailbox routes
[**PostInboundSendRaw**](InboundApi.md#PostInboundSendRaw) | **Post** /inbound/send-raw | Send mime document
[**PostInboundUpdateRoute**](InboundApi.md#PostInboundUpdateRoute) | **Post** /inbound/update-route | Update mailbox route

# **PostInboundAddDomain**
> InlineResponse2009 PostInboundAddDomain(ctx, body)
Add inbound domain

Add an inbound domain to your account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InboundAdddomainBody**](InboundAdddomainBody.md)|  | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostInboundAddRoute**
> InlineResponse20013 PostInboundAddRoute(ctx, body)
Add mailbox route

Add a new mailbox route to an inbound domain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InboundAddrouteBody**](InboundAddrouteBody.md)|  | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostInboundCheckDomain**
> InlineResponse20010 PostInboundCheckDomain(ctx, body)
Check domain settings

Check the MX settings for an inbound domain. The domain must have already been added with the add-domain call.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InboundCheckdomainBody**](InboundCheckdomainBody.md)|  | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostInboundDeleteDomain**
> InlineResponse20011 PostInboundDeleteDomain(ctx, body)
Delete inbound domain

Delete an inbound domain from the account. All mail will stop routing for this domain immediately.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InboundDeletedomainBody**](InboundDeletedomainBody.md)|  | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostInboundDeleteRoute**
> InlineResponse20015 PostInboundDeleteRoute(ctx, body)
Delete mailbox route

Delete an existing inbound mailbox route.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InboundDeleterouteBody**](InboundDeleterouteBody.md)|  | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostInboundDomains**
> []InlineResponse2008 PostInboundDomains(ctx, body)
List inbound domains

List the domains that have been configured for inbound delivery.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InboundDomainsBody**](InboundDomainsBody.md)|  | 

### Return type

[**[]InlineResponse2008**](inline_response_200_8.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostInboundRoutes**
> []InlineResponse20012 PostInboundRoutes(ctx, body)
List mailbox routes

List the mailbox routes defined for an inbound domain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InboundRoutesBody**](InboundRoutesBody.md)|  | 

### Return type

[**[]InlineResponse20012**](inline_response_200_12.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostInboundSendRaw**
> []InlineResponse20016 PostInboundSendRaw(ctx, body)
Send mime document

Take a raw MIME document destined for a domain with inbound domains set up, and send it to the inbound hook exactly as if it had been sent over SMTP.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InboundSendrawBody**](InboundSendrawBody.md)|  | 

### Return type

[**[]InlineResponse20016**](inline_response_200_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostInboundUpdateRoute**
> InlineResponse20014 PostInboundUpdateRoute(ctx, body)
Update mailbox route

Update the pattern or webhook of an existing inbound mailbox route. If null is provided for any fields, the values will remain unchanged.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InboundUpdaterouteBody**](InboundUpdaterouteBody.md)|  | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

