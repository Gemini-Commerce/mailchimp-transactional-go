# InlineResponse2006

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | the unique identifier for this Export. Use this identifier when checking the export job&#39;s status | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | the date and time that the export job was created as a UTC string in YYYY-MM-DD HH:MM:SS format | [optional] [default to null]
**Type_** | **string** | the type of the export job | [optional] [default to null]
**FinishedAt** | [**time.Time**](time.Time.md) | the date and time that the export job was finished as a UTC string in YYYY-MM-DD HH:MM:SS format, or null for jobs that have not run | [optional] [default to null]
**State** | **string** | the export job&#39;s state | [optional] [default to null]
**ResultUrl** | **string** | the url for the export job&#39;s results, if the job is complete | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


