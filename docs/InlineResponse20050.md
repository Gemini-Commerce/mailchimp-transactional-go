# InlineResponse20050

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | a unique indentifier for the subaccount | [optional] [default to null]
**Name** | **string** | an optional display name for the subaccount | [optional] [default to null]
**CustomQuota** | **int32** | an optional manual hourly quota for the subaccount. If not specified, the hourly quota will be managed based on reputation | [optional] [default to null]
**Status** | **string** | the current sending status of the subaccount | [optional] [default to null]
**Reputation** | **int32** | the subaccount&#39;s current reputation on a scale from 0 to 100 | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | the date and time that the subaccount was created as a UTC string in YYYY-MM-DD HH:MM:SS format | [optional] [default to null]
**FirstSentAt** | [**time.Time**](time.Time.md) | the date and time that the subaccount first sent as a UTC string in YYYY-MM-DD HH:MM:SS format | [optional] [default to null]
**SentWeekly** | **int32** | the number of emails the subaccount has sent so far this week (weeks start on midnight Monday, UTC) | [optional] [default to null]
**SentMonthly** | **int32** | the number of emails the subaccount has sent so far this month (months start on midnight of the 1st, UTC) | [optional] [default to null]
**SentTotal** | **int32** | the number of emails the subaccount has sent since it was created | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


