# InlineResponse20030

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | **int32** | the Unix timestamp from when this message was sent | [optional] [default to null]
**Id** | **string** | the message&#x27;s unique id | [optional] [default to null]
**Sender** | **string** | the email address of the sender | [optional] [default to null]
**Template** | **string** | the unique name of the template used, if any | [optional] [default to null]
**Subject** | **string** | the message&#x27;s subject line | [optional] [default to null]
**Email** | **string** | the recipient email address | [optional] [default to null]
**Tags** | **[]string** | list of tags on this message | [optional] [default to null]
**Opens** | **int32** | how many times has this message been opened | [optional] [default to null]
**OpensDetail** | [**[]MessagessearchOpensDetail**](messagessearch_opens_detail.md) | list of individual opens for the message | [optional] [default to null]
**Clicks** | **int32** | how many times has a link been clicked in this message | [optional] [default to null]
**ClicksDetail** | [**[]MessagessearchClicksDetail**](messagessearch_clicks_detail.md) | list of individual clicks for the message | [optional] [default to null]
**State** | **string** | sending status of this message | [optional] [default to null]
**Metadata** | [***interface{}**](interface{}.md) | any custom metadata provided when the message was sent | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

