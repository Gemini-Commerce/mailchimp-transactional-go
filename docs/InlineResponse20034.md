# InlineResponse20034

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | **string** | the subject of the message | [optional] [default to null]
**FromEmail** | **string** | the email address of the sender | [optional] [default to null]
**FromName** | **string** | the alias of the sender (if any) | [optional] [default to null]
**To** | [**[]InlineResponse20034To**](inline_response_200_34_to.md) | an array of any recipients in the message | [optional] [default to null]
**Headers** | [***interface{}**](interface{}.md) | the key-value pairs of the MIME headers for the message&#x27;s main document | [optional] [default to null]
**Text** | **string** | the text part of the message, if any | [optional] [default to null]
**Html** | **string** | the HTML part of the message, if any | [optional] [default to null]
**Attachments** | [**[]InlineResponse20034Attachments**](inline_response_200_34_attachments.md) | an array of any attachments that can be found in the message | [optional] [default to null]
**Images** | [**[]InlineResponse20034Images**](inline_response_200_34_images.md) | an array of any embedded images that can be found in the message | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

