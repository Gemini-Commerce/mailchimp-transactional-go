# MessagesSendtemplateBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | a valid api key | [default to null]
**TemplateName** | **string** | the immutable name or slug of a template that exists in the user&#x27;s account. For backwards-compatibility, the template name may also be used but the immutable slug is preferred. | [default to null]
**TemplateContent** | [**[]MessagessendtemplateTemplateContent**](messagessendtemplate_template_content.md) | an array of template content to send. Each item in the array should be a struct with two keys - name: the name of the content block to set the content for, and content: the actual content to put into the block | [default to null]
**Message** | [***MessagessendtemplateMessage**](messagessendtemplate_message.md) |  | [default to null]
**Async** | **bool** | enable a background sending mode that is optimized for bulk sending. In async mode, messages/send will immediately return a status of \&quot;queued\&quot; for every recipient. To handle rejections when sending in async mode, set up a webhook for the &#x27;reject&#x27; event. Defaults to false for messages with no more than 10 recipients; messages with more than 10 recipients are always sent asynchronously, regardless of the value of async. | [optional] [default to null]
**IpPool** | **string** | the name of the dedicated ip pool that should be used to send the message. If you do not have any dedicated IPs, this parameter has no effect. If you specify a pool that does not exist, your default pool will be used instead. | [optional] [default to null]
**SendAt** | [**time.Time**](time.Time.md) | when this message should be sent as a UTC timestamp in YYYY-MM-DD HH:MM:SS format. If you specify a time in the past, the message will be sent immediately. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

