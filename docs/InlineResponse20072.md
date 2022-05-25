# InlineResponse20072

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | the username of the user (used for SMTP authentication) | [optional] [default to null]
**CreatedAt** | **string** | the date and time that the user&#39;s Mandrill account was created as a UTC string in YYYY-MM-DD HH:MM:SS format | [optional] [default to null]
**PublicId** | **string** | a unique, permanent identifier for this user | [optional] [default to null]
**Reputation** | **int32** | the reputation of the user on a scale from 0 to 100, with 75 generally being a \&quot;good\&quot; reputation | [optional] [default to null]
**HourlyQuota** | **int32** | the maximum number of emails Mandrill will deliver for this user each hour. Any emails beyond that will be accepted and queued for later delivery. Users with higher reputations will have higher hourly quotas | [optional] [default to null]
**Backlog** | **int32** | the number of emails that are queued for delivery due to exceeding your monthly or hourly quotas | [optional] [default to null]
**Stats** | [***InlineResponse20072Stats**](inline_response_200_72_stats.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


