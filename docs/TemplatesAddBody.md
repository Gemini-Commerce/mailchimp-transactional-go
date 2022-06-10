# TemplatesAddBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | a valid api key | [default to null]
**Name** | **string** | the name for the new template - must be unique | [default to null]
**FromEmail** | **string** | a default sending address for emails sent using this template | [optional] [default to null]
**FromName** | **string** | a default from name to be used | [optional] [default to null]
**Subject** | **string** | a default subject line to be used | [optional] [default to null]
**Code** | **string** | the HTML code for the template with mc:edit attributes for the editable elements | [optional] [default to null]
**Text** | **string** | a default text part to be used when sending with this template | [optional] [default to null]
**Publish** | **bool** | set to false to add a draft template without publishing | [optional] [default to null]
**Labels** | **[]string** | an optional array of up to 10 labels to use for filtering templates | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

