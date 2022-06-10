# InlineResponse20063

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | **string** | the immutable unique code name of the template | [optional] [default to null]
**Name** | **string** | the name of the template | [optional] [default to null]
**Labels** | **[]string** | the list of labels applied to the template | [optional] [default to null]
**Code** | **string** | the full HTML code of the template, with mc:edit attributes marking the editable elements - draft version | [optional] [default to null]
**Subject** | **string** | the subject line of the template, if provided - draft version | [optional] [default to null]
**FromEmail** | **string** | the default sender address for the template, if provided - draft version | [optional] [default to null]
**FromName** | **string** | the default sender from name for the template, if provided - draft version | [optional] [default to null]
**Text** | **string** | the default text part of messages sent with the template, if provided - draft version | [optional] [default to null]
**PublishName** | **string** | the same as the template name - kept as a separate field for backwards compatibility | [optional] [default to null]
**PublishCode** | **string** | the full HTML code of the template, with mc:edit attributes marking the editable elements that are available as published, if it has been published | [optional] [default to null]
**PublishSubject** | **string** | the subject line of the template, if provided | [optional] [default to null]
**PublishFromEmail** | **string** | the default sender address for the template, if provided | [optional] [default to null]
**PublishFromName** | **string** | the default sender from name for the template, if provided | [optional] [default to null]
**PublishText** | **string** | the default text part of messages sent with the template, if provided | [optional] [default to null]
**PublishedAt** | [**time.Time**](time.Time.md) | the date and time the template was last published as a UTC string in YYYY-MM-DD HH:MM:SS format, or null if it has not been published | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | the date and time the template was first created as a UTC string in YYYY-MM-DD HH:MM:SS format | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | the date and time the template was last modified as a UTC string in YYYY-MM-DD HH:MM:SS format | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

