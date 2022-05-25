# SendersdomainsSpf

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Valid** | **bool** | whether the domain&#39;s SPF record is valid for use with Mandrill | [optional] [default to null]
**ValidAfter** | [**time.Time**](time.Time.md) | when the domain&#39;s SPF record will be considered valid for use with Mandrill as a UTC string in YYYY-MM-DD HH:MM:SS format. If set, this indicates that the record is valid now, but was previously invalid, and Mandrill will wait until the record&#39;s TTL elapses to start using it. | [optional] [default to null]
**Error_** | **string** | an error describing the spf record, or null if the record is correct | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


