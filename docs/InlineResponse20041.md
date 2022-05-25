# InlineResponse20041

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | the email that is blocked | [optional] [default to null]
**Reason** | **string** | the type of event (hard-bounce, soft-bounce, spam, unsub, custom) that caused this rejection | [optional] [default to null]
**Detail** | **string** | extended details about the event, such as the SMTP diagnostic for bounces or the comment for manually-created rejections | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | when the email was added to the denylist | [optional] [default to null]
**LastEventAt** | [**time.Time**](time.Time.md) | the timestamp of the most recent event that either created or renewed this rejection | [optional] [default to null]
**ExpiresAt** | [**time.Time**](time.Time.md) | when the denylist entry will expire (this may be in the past) | [optional] [default to null]
**Expired** | **bool** | whether the denylist entry has expired | [optional] [default to null]
**Sender** | [***RejectslistSender**](rejectslist_sender.md) |  | [optional] [default to null]
**Subaccount** | **string** | the subaccount that this denylist entry applies to, or null if none. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


