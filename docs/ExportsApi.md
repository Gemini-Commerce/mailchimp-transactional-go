# {{classname}}

All URIs are relative to *https://mandrillapp.com/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostExportsActivity**](ExportsApi.md#PostExportsActivity) | **Post** /exports/activity | Export activity history
[**PostExportsAllowlist**](ExportsApi.md#PostExportsAllowlist) | **Post** /exports/allowlist | Export Allowlist
[**PostExportsInfo**](ExportsApi.md#PostExportsInfo) | **Post** /exports/info | View export info
[**PostExportsList**](ExportsApi.md#PostExportsList) | **Post** /exports/list | List exports
[**PostExportsRejects**](ExportsApi.md#PostExportsRejects) | **Post** /exports/rejects | Export denylist
[**PostExportsWhitelist**](ExportsApi.md#PostExportsWhitelist) | **Post** /exports/whitelist | Export Allowlist

# **PostExportsActivity**
> InlineResponse2007 PostExportsActivity(ctx, body)
Export activity history

Begins an export of your activity history. The activity will be exported to a zip archive containing a single file named activity.csv in the same format as you would be able to export from your account's activity view. It includes the following fields: Date, Email Address, Sender, Subject, Status, Tags, Opens, Clicks, Bounce Detail. If you have configured any custom metadata fields, they will be included in the exported data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExportsActivityBody**](ExportsActivityBody.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostExportsAllowlist**
> InlineResponse2006 PostExportsAllowlist(ctx, body)
Export Allowlist

Begins an export of your rejection allowlist. The allowlist will be exported to a zip archive containing a single file named allowlist.csv that includes the following fields: email, detail, created_at.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExportsAllowlistBody**](ExportsAllowlistBody.md)|  | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostExportsInfo**
> InlineResponse2003 PostExportsInfo(ctx, body)
View export info

Returns information about an export job. If the export job's state is 'complete', the returned data will include a URL you can use to fetch the results. Every export job produces a zip archive, but the format of the archive is distinct for each job type. The api calls that initiate exports include more details about the output format for that job type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExportsInfoBody**](ExportsInfoBody.md)|  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostExportsList**
> []InlineResponse2004 PostExportsList(ctx, body)
List exports

Returns a list of your exports.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExportsListBody**](ExportsListBody.md)|  | 

### Return type

[**[]InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostExportsRejects**
> InlineResponse2005 PostExportsRejects(ctx, body)
Export denylist

Begins an export of your rejection denylist. The denylist will be exported to a zip archive containing a single file named rejects.csv that includes the following fields: email, reason, detail, created_at, expires_at, last_event_at, expires_at.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExportsRejectsBody**](ExportsRejectsBody.md)|  | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostExportsWhitelist**
> InlineResponse2006 PostExportsWhitelist(ctx, body)
Export Allowlist

Begins an export of your rejection allowlist. The allowlist will be exported to a zip archive containing a single file named allowlist.csv that includes the following fields: email, detail, created_at.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExportsWhitelistBody**](ExportsWhitelistBody.md)|  | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml, application/x-php, application/x-yaml; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

