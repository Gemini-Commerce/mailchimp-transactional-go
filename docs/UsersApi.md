# {{classname}}

All URIs are relative to *https://mandrillapp.com/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostUsersInfo**](UsersApi.md#PostUsersInfo) | **Post** /users/info | Get user info
[**PostUsersPing**](UsersApi.md#PostUsersPing) | **Post** /users/ping | Ping
[**PostUsersPing2**](UsersApi.md#PostUsersPing2) | **Post** /users/ping2 | Ping 2
[**PostUsersSenders**](UsersApi.md#PostUsersSenders) | **Post** /users/senders | List account senders

# **PostUsersInfo**
> InlineResponse20072 PostUsersInfo(ctx, body)
Get user info

Return the information about the API-connected user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UsersInfoBody**](UsersInfoBody.md)|  | 

### Return type

[**InlineResponse20072**](inline_response_200_72.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUsersPing**
> string PostUsersPing(ctx, body)
Ping

Validate an API key and respond to a ping.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UsersPingBody**](UsersPingBody.md)|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUsersPing2**
> InlineResponse20073 PostUsersPing2(ctx, body)
Ping 2

Validate an API key and respond to a ping (JSON parser version).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UsersPing2Body**](UsersPing2Body.md)|  | 

### Return type

[**InlineResponse20073**](inline_response_200_73.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUsersSenders**
> []InlineResponse20043 PostUsersSenders(ctx, body)
List account senders

Return the senders that have tried to use this account, both verified and unverified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UsersSendersBody**](UsersSendersBody.md)|  | 

### Return type

[**[]InlineResponse20043**](inline_response_200_43.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

