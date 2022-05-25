# \SubaccountsApi

All URIs are relative to *https://mandrillapp.com/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostSubaccountsAdd**](SubaccountsApi.md#PostSubaccountsAdd) | **Post** /subaccounts/add | Add subaccount
[**PostSubaccountsDelete**](SubaccountsApi.md#PostSubaccountsDelete) | **Post** /subaccounts/delete | Delete subaccount
[**PostSubaccountsInfo**](SubaccountsApi.md#PostSubaccountsInfo) | **Post** /subaccounts/info | Get subaccount info
[**PostSubaccountsList**](SubaccountsApi.md#PostSubaccountsList) | **Post** /subaccounts/list | List subaccounts
[**PostSubaccountsPause**](SubaccountsApi.md#PostSubaccountsPause) | **Post** /subaccounts/pause | Pause subaccount
[**PostSubaccountsResume**](SubaccountsApi.md#PostSubaccountsResume) | **Post** /subaccounts/resume | Resume subaccount
[**PostSubaccountsUpdate**](SubaccountsApi.md#PostSubaccountsUpdate) | **Post** /subaccounts/update | Update subaccount


# **PostSubaccountsAdd**
> InlineResponse20051 PostSubaccountsAdd(ctx, body)
Add subaccount

Add a new subaccount.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body57**](Body57.md)|  | 

### Return type

[**InlineResponse20051**](inline_response_200_51.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSubaccountsDelete**
> InlineResponse20054 PostSubaccountsDelete(ctx, body)
Delete subaccount

Delete an existing subaccount. Any email related to the subaccount will be saved, but stats will be removed and any future sending calls to this subaccount will fail.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body60**](Body60.md)|  | 

### Return type

[**InlineResponse20054**](inline_response_200_54.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSubaccountsInfo**
> InlineResponse20052 PostSubaccountsInfo(ctx, body)
Get subaccount info

Given the ID of an existing subaccount, return the data about it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body58**](Body58.md)|  | 

### Return type

[**InlineResponse20052**](inline_response_200_52.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSubaccountsList**
> []InlineResponse20050 PostSubaccountsList(ctx, body)
List subaccounts

Get the list of subaccounts defined for the account, optionally filtered by a prefix.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body56**](Body56.md)|  | 

### Return type

[**[]InlineResponse20050**](inline_response_200_50.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSubaccountsPause**
> InlineResponse20055 PostSubaccountsPause(ctx, body)
Pause subaccount

Pause a subaccount's sending. Any future emails delivered to this subaccount will be queued for a maximum of 3 days until the subaccount is resumed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body61**](Body61.md)|  | 

### Return type

[**InlineResponse20055**](inline_response_200_55.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSubaccountsResume**
> InlineResponse20056 PostSubaccountsResume(ctx, body)
Resume subaccount

Resume a paused subaccount's sending.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body62**](Body62.md)|  | 

### Return type

[**InlineResponse20056**](inline_response_200_56.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSubaccountsUpdate**
> InlineResponse20053 PostSubaccountsUpdate(ctx, body)
Update subaccount

Update an existing subaccount.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body59**](Body59.md)|  | 

### Return type

[**InlineResponse20053**](inline_response_200_53.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

