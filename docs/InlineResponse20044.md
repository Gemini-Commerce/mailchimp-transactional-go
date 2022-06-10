# InlineResponse20044

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | the sender domain name | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | the date and time that the sending domain was first seen as a UTC string in YYYY-MM-DD HH:MM:SS format | [optional] [default to null]
**LastTestedAt** | [**time.Time**](time.Time.md) | when the domain&#x27;s DNS settings were last tested as a UTC string in YYYY-MM-DD HH:MM:SS format | [optional] [default to null]
**Spf** | [***SendersdomainsSpf**](sendersdomains_spf.md) |  | [optional] [default to null]
**Dkim** | [***SendersdomainsDkim**](sendersdomains_dkim.md) |  | [optional] [default to null]
**VerifiedAt** | [**time.Time**](time.Time.md) | if the domain has been verified, this indicates when that verification occurred as a UTC string in YYYY-MM-DD HH:MM:SS format | [optional] [default to null]
**ValidSigning** | **bool** | whether this domain can be used to authenticate mail, either for itself or as a custom signing domain. If this is false but spf and dkim are both valid, you will need to verify the domain before using it to authenticate mail | [optional] [default to null]
**VerifyTxtKey** | **string** | a unique key used to verify a domain by adding a TXT record. Append this key to &#x27;mandrill_verify.&#x27; and add it to your domain&#x27;s TXT records to verify | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

