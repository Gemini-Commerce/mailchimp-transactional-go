# MessagesRescheduleBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | a valid api key | [default to null]
**Id** | **string** | a scheduled email id, as returned by any of the messages/send calls or messages/list-scheduled | [default to null]
**SendAt** | [**time.Time**](time.Time.md) | the new UTC timestamp when the message should sent. Mandrill can&#x27;t time travel, so if you specify a time in past the message will be sent immediately | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

