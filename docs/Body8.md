# Body8

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | a valid api key | [default to null]
**NotifyEmail** | **string** | an optional email address to notify when the export job has finished | [optional] [default to null]
**DateFrom** | [**time.Time**](time.Time.md) | start date as a UTC string in YYYY-MM-DD HH:MM:SS format | [optional] [default to null]
**DateTo** | [**time.Time**](time.Time.md) | end date as a UTC string in YYYY-MM-DD HH:MM:SS format | [optional] [default to null]
**Tags** | **[]string** | an array of tag names to narrow the export to; will match messages that contain ANY of the tags | [optional] [default to null]
**Senders** | **[]string** | an array of senders to narrow the export to | [optional] [default to null]
**States** | **[]string** | an array of message states to narrow the export to; messages with ANY of the states will be included | [optional] [default to null]
**ApiKeys** | **[]string** | an array of api keys to narrow the export to; messsagse sent with ANY of the keys will be included | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


