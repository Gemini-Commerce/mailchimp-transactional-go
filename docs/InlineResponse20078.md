# InlineResponse20078

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | a unique integer indentifier for the webhook | [optional] [default to null]
**Url** | **string** | The URL that the event data will be posted to | [optional] [default to null]
**Description** | **string** | a description of the webhook | [optional] [default to null]
**AuthKey** | **string** | the key used to requests for this webhook | [optional] [default to null]
**Events** | **[]string** | The message events that will be posted to the hook | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | the date and time that the webhook was created as a UTC string in YYYY-MM-DD HH:MM:SS format | [optional] [default to null]
**LastSentAt** | [**time.Time**](time.Time.md) | the date and time that the webhook last successfully received events as a UTC string in YYYY-MM-DD HH:MM:SS format | [optional] [default to null]
**BatchesSent** | **int32** | the number of event batches that have ever been sent to this webhook | [optional] [default to null]
**EventsSent** | **int32** | the total number of events that have ever been sent to this webhook | [optional] [default to null]
**LastError** | **string** | if we&#39;ve ever gotten an error trying to post to this webhook, the last error that we&#39;ve seen | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


