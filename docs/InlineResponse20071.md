# InlineResponse20071

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | the tracking domain name | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | the date and time that the tracking domain was added as a UTC string in YYYY-MM-DD HH:MM:SS format | [optional] [default to null]
**LastTestedAt** | [**time.Time**](time.Time.md) | when the domain&#x27;s DNS settings were last tested as a UTC string in YYYY-MM-DD HH:MM:SS format | [optional] [default to null]
**Cname** | [***UrlstrackingdomainsCname**](urlstrackingdomains_cname.md) |  | [optional] [default to null]
**ValidTracking** | **bool** | whether this domain can be used as a tracking domain for email. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

