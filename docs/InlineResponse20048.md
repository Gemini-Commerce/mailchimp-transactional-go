# InlineResponse20048

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | the sender&#x27;s email address | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | the date and time that the sender was first seen by Mandrill as a UTC date string in YYYY-MM-DD HH:MM:SS format | [optional] [default to null]
**Sent** | **int32** | the total number of messages sent by this sender | [optional] [default to null]
**HardBounces** | **int32** | the total number of hard bounces by messages by this sender | [optional] [default to null]
**SoftBounces** | **int32** | the total number of soft bounces by messages by this sender | [optional] [default to null]
**Rejects** | **int32** | the total number of rejected messages by this sender | [optional] [default to null]
**Complaints** | **int32** | the total number of spam complaints received for messages by this sender | [optional] [default to null]
**Unsubs** | **int32** | the total number of unsubscribe requests received for messages by this sender | [optional] [default to null]
**Opens** | **int32** | the total number of times messages by this sender have been opened | [optional] [default to null]
**Clicks** | **int32** | the total number of times tracked URLs in messages by this sender have been clicked | [optional] [default to null]
**Stats** | [***InlineResponse20048Stats**](inline_response_200_48_stats.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

