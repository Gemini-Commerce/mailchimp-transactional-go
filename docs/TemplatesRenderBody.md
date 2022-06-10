# TemplatesRenderBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | a valid api key | [default to null]
**TemplateName** | **string** | the immutable name of a template that exists in the user&#x27;s account | [default to null]
**TemplateContent** | [**[]TemplatesrenderTemplateContent**](templatesrender_template_content.md) | an array of template content to render. Each item in the array should be a struct with two keys - name: the name of the content block to set the content for, and content: the actual content to put into the block | [default to null]
**MergeVars** | [**[]TemplatesrenderMergeVars**](templatesrender_merge_vars.md) | optional merge variables to use for injecting merge field content. If this is not provided, no merge fields will be replaced. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

