# \IpsApi

All URIs are relative to *https://mandrillapp.com/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostIpsCancelWarmup**](IpsApi.md#PostIpsCancelWarmup) | **Post** /ips/cancel-warmup | Cancel ip warmup
[**PostIpsCheckCustomDns**](IpsApi.md#PostIpsCheckCustomDns) | **Post** /ips/check-custom-dns | Test custom dns
[**PostIpsCreatePool**](IpsApi.md#PostIpsCreatePool) | **Post** /ips/create-pool | Add ip pool
[**PostIpsDelete**](IpsApi.md#PostIpsDelete) | **Post** /ips/delete | Delete ip address
[**PostIpsDeletePool**](IpsApi.md#PostIpsDeletePool) | **Post** /ips/delete-pool | Delete ip pool
[**PostIpsInfo**](IpsApi.md#PostIpsInfo) | **Post** /ips/info | Get ip info
[**PostIpsList**](IpsApi.md#PostIpsList) | **Post** /ips/list | List ip addresses
[**PostIpsListPools**](IpsApi.md#PostIpsListPools) | **Post** /ips/list-pools | List ip pools
[**PostIpsPoolInfo**](IpsApi.md#PostIpsPoolInfo) | **Post** /ips/pool-info | Get ip pool info
[**PostIpsProvision**](IpsApi.md#PostIpsProvision) | **Post** /ips/provision | Request additional ip
[**PostIpsSetCustomDns**](IpsApi.md#PostIpsSetCustomDns) | **Post** /ips/set-custom-dns | Set custom dns
[**PostIpsSetPool**](IpsApi.md#PostIpsSetPool) | **Post** /ips/set-pool | Move ip to different pool
[**PostIpsStartWarmup**](IpsApi.md#PostIpsStartWarmup) | **Post** /ips/start-warmup | Start ip warmup


# **PostIpsCancelWarmup**
> InlineResponse20020 PostIpsCancelWarmup(ctx, body)
Cancel ip warmup

Cancels the warmup process for a dedicated IP.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body22**](Body22.md)|  | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIpsCheckCustomDns**
> InlineResponse20026 PostIpsCheckCustomDns(ctx, body)
Test custom dns

Tests whether a domain name is valid for use as the custom reverse DNS for a dedicated IP.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body29**](Body29.md)|  | 

### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIpsCreatePool**
> InlineResponse20024 PostIpsCreatePool(ctx, body)
Add ip pool

Creates a pool and returns it. If a pool already exists with this name, no action will be performed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body27**](Body27.md)|  | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIpsDelete**
> InlineResponse20022 PostIpsDelete(ctx, body)
Delete ip address

Deletes a dedicated IP. This is permanent and cannot be undone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body24**](Body24.md)|  | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIpsDeletePool**
> InlineResponse20025 PostIpsDeletePool(ctx, body)
Delete ip pool

Deletes a pool. A pool must be empty before you can delete it, and you cannot delete your default pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body28**](Body28.md)|  | 

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIpsInfo**
> InlineResponse20018 PostIpsInfo(ctx, body)
Get ip info

Retrieves information about a single dedicated IP.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body19**](Body19.md)|  | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIpsList**
> []InlineResponse20017 PostIpsList(ctx, body)
List ip addresses

Lists your dedicated IPs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body18**](Body18.md)|  | 

### Return type

[**[]InlineResponse20017**](inline_response_200_17.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIpsListPools**
> []InlineResponse20023 PostIpsListPools(ctx, body)
List ip pools

Lists your dedicated IP pools.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body25**](Body25.md)|  | 

### Return type

[**[]InlineResponse20023**](inline_response_200_23.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIpsPoolInfo**
> InlineResponse20024 PostIpsPoolInfo(ctx, body)
Get ip pool info

Describes a single dedicated IP pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body26**](Body26.md)|  | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIpsProvision**
> InlineResponse20019 PostIpsProvision(ctx, body)
Request additional ip

Requests an additional dedicated IP for your account. Accounts may have one outstanding request at any time, and provisioning requests are processed within 24 hours.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body20**](Body20.md)|  | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIpsSetCustomDns**
> InlineResponse20027 PostIpsSetCustomDns(ctx, body)
Set custom dns

Configures the custom DNS name for a dedicated IP.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body30**](Body30.md)|  | 

### Return type

[**InlineResponse20027**](inline_response_200_27.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIpsSetPool**
> InlineResponse20021 PostIpsSetPool(ctx, body)
Move ip to different pool

Moves a dedicated IP to a different pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body23**](Body23.md)|  | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIpsStartWarmup**
> InlineResponse20020 PostIpsStartWarmup(ctx, body)
Start ip warmup

Begins the warmup process for a dedicated IP. During the warmup process, the Transactional API will gradually increase the percentage of your mail that is sent over the warming-up IP, over a period of roughly 30 days. The rest of your mail will be sent over shared IPs or other dedicated IPs in the same pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body21**](Body21.md)|  | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

