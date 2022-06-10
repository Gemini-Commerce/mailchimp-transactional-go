# UrlstrackingdomainsCname

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Valid** | **bool** | whether the domain&#x27;s CNAME record is valid for use with Mandrill | [optional] [default to null]
**ValidAfter** | **string** | when the domain&#x27;s CNAME record will be considered valid for use with Mandrill as a UTC string in YYYY-MM-DD HH:MM:SS format. If set, this indicates that the record is valid now, but was previously invalid, and Mandrill will wait until the record&#x27;s TTL elapses to start using it. | [optional] [default to null]
**Error_** | **string** | an error describing the CNAME record, or null if the record is correct | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

