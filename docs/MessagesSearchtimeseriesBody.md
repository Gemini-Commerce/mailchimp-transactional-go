# MessagesSearchtimeseriesBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | a valid api key | [default to null]
**Query** | **string** | the search terms to find matching messages for | [optional] [default to null]
**DateFrom** | [**time.Time**](time.Time.md) | start date | [optional] [default to null]
**DateTo** | [**time.Time**](time.Time.md) | end date | [optional] [default to null]
**Tags** | **[]string** | an array of tag names to narrow the search to, will return messages that contain ANY of the tags | [optional] [default to null]
**Senders** | **[]string** | an array of sender addresses to narrow the search to, will return messages sent by ANY of the senders | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

