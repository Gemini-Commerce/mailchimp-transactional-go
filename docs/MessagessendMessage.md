# MessagessendMessage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Html** | **string** | the full HTML content to be sent | [optional] [default to null]
**Text** | **string** | optional full text content to be sent | [optional] [default to null]
**Subject** | **string** | the message subject | [optional] [default to null]
**FromEmail** | **string** | the sender email address | [optional] [default to null]
**FromName** | **string** | optional from name to be used | [optional] [default to null]
**To** | [**[]MessagessendMessageTo**](messagessend_message_to.md) | an array of recipient information. | [optional] [default to null]
**Headers** | [***interface{}**](interface{}.md) | optional extra headers to add to the message (most headers are allowed) | [optional] [default to null]
**Important** | **bool** | whether or not this message is important, and should be delivered ahead of non-important messages | [optional] [default to null]
**TrackOpens** | **bool** | whether or not to turn on open tracking for the message | [optional] [default to null]
**TrackClicks** | **bool** | whether or not to turn on click tracking for the message | [optional] [default to null]
**AutoText** | **bool** | whether or not to automatically generate a text part for messages that are not given text | [optional] [default to null]
**AutoHtml** | **bool** | whether or not to automatically generate an HTML part for messages that are not given HTML | [optional] [default to null]
**InlineCss** | **bool** | whether or not to automatically inline all CSS styles provided in the message HTML - only for HTML documents less than 256KB in size | [optional] [default to null]
**UrlStripQs** | **bool** | whether or not to strip the query string from URLs when aggregating tracked URL data | [optional] [default to null]
**PreserveRecipients** | **bool** | whether or not to expose all recipients in to \&quot;To\&quot; header for each email | [optional] [default to null]
**ViewContentLink** | **bool** | set to false to remove content logging for sensitive emails | [optional] [default to null]
**BccAddress** | **string** | an optional address to receive an exact copy of each recipient&#39;s email | [optional] [default to null]
**TrackingDomain** | **string** | a custom domain to use for tracking opens and clicks instead of mandrillapp.com | [optional] [default to null]
**SigningDomain** | **string** | a custom domain to use for SPF/DKIM signing instead of mandrill (for \&quot;via\&quot; or \&quot;on behalf of\&quot; in email clients) | [optional] [default to null]
**ReturnPathDomain** | **string** | a custom domain to use for the messages&#39;s return-path | [optional] [default to null]
**Merge** | **bool** | whether to evaluate merge tags in the message. Will automatically be set to true if either merge_vars or global_merge_vars are provided. | [optional] [default to null]
**MergeLanguage** | **string** | the merge tag language to use when evaluating merge tags, either mailchimp or handlebars | [optional] [default to null]
**GlobalMergeVars** | [**[]MessagessendMessageGlobalMergeVars**](messagessend_message_global_merge_vars.md) | global merge variables to use for all recipients. You can override these per recipient. | [optional] [default to null]
**MergeVars** | [**[]MessagessendMessageMergeVars**](messagessend_message_merge_vars.md) | per-recipient merge variables, which override global merge variables with the same name. | [optional] [default to null]
**Tags** | **[]string** | an array of string to tag the message with. Stats are accumulated using tags, though we only store the first 100 we see, so this should not be unique or change frequently. Tags should be 50 characters or less. Any tags starting with an underscore are reserved for internal use and will cause errors. | [optional] [default to null]
**Subaccount** | **string** | the unique id of a subaccount for this message - must already exist or will fail with an error | [optional] [default to null]
**GoogleAnalyticsDomains** | **[]string** | an array of strings indicating for which any matching URLs will automatically have Google Analytics parameters appended to their query string automatically. | [optional] [default to null]
**GoogleAnalyticsCampaign** | **string** | optional string indicating the value to set for the utm_campaign tracking parameter. If this isn&#39;t provided the email&#39;s from address will be used instead. | [optional] [default to null]
**Metadata** | [***MessagessendMessageMetadata**](messagessend_message_metadata.md) |  | [optional] [default to null]
**RecipientMetadata** | [**[]MessagessendMessageRecipientMetadata**](messagessend_message_recipient_metadata.md) | Per-recipient metadata that will override the global values specified in the metadata parameter. | [optional] [default to null]
**Attachments** | [**[]MessagessendMessageAttachments**](messagessend_message_attachments.md) | an array of supported attachments to add to the message | [optional] [default to null]
**Images** | [**[]MessagessendMessageImages**](messagessend_message_images.md) | an array of embedded images to add to the message | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


