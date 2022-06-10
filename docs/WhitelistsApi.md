# {{classname}}

All URIs are relative to *https://mandrillapp.com/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostWhitelistsAdd**](WhitelistsApi.md#PostWhitelistsAdd) | **Post** /whitelists/add | Add email to allowlist
[**PostWhitelistsDelete**](WhitelistsApi.md#PostWhitelistsDelete) | **Post** /whitelists/delete | Remove email from allowlist
[**PostWhitelistsList**](WhitelistsApi.md#PostWhitelistsList) | **Post** /whitelists/list | List allowlisted emails

# **PostWhitelistsAdd**
> InlineResponse200 PostWhitelistsAdd(ctx, body)
Add email to allowlist

Adds an email to your email rejection allowlist. If the address is currently on your denylist, that denylist entry will be removed automatically.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WhitelistsAddBody**](WhitelistsAddBody.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostWhitelistsDelete**
> InlineResponse2002 PostWhitelistsDelete(ctx, body)
Remove email from allowlist

Removes an email address from the allowlist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WhitelistsDeleteBody**](WhitelistsDeleteBody.md)|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostWhitelistsList**
> []InlineResponse2001 PostWhitelistsList(ctx, body)
List allowlisted emails

Retrieves your email rejection allowlist. You can provide an email address or search prefix to limit the results. Returns up to 1000 results.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WhitelistsListBody**](WhitelistsListBody.md)|  | 

### Return type

[**[]InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

