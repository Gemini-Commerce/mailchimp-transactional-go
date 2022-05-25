# InlineResponse20033

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | **int32** | the Unix timestamp from when this message was sent | [optional] [default to null]
**Id** | **string** | the message&#39;s unique id | [optional] [default to null]
**FromEmail** | **string** | the email address of the sender | [optional] [default to null]
**FromName** | **string** | the alias of the sender (if any) | [optional] [default to null]
**Subject** | **string** | the message&#39;s subject line | [optional] [default to null]
**To** | [***InlineResponse20033To**](inline_response_200_33_to.md) |  | [optional] [default to null]
**Tags** | **[]string** | list of tags on this message | [optional] [default to null]
**Headers** | [***interface{}**](interface{}.md) | the key-value pairs of the custom MIME headers for the message&#39;s main document | [optional] [default to null]
**Text** | **string** | the text part of the message, if any | [optional] [default to null]
**Html** | **string** | the HTML part of the message, if any | [optional] [default to null]
**Attachments** | [**[]InlineResponse20033Attachments**](inline_response_200_33_attachments.md) | an array of any attachments that can be found in the message | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


