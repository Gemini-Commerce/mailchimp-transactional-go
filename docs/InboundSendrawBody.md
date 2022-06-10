# InboundSendrawBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | a valid api key | [default to null]
**RawMessage** | **string** | the full MIME document of an email message | [default to null]
**To** | **[]string** | optionally define the recipients to receive the message - otherwise we&#x27;ll use the To, Cc, and Bcc headers provided in the document | [optional] [default to null]
**MailFrom** | **string** | the address specified in the MAIL FROM stage of the SMTP conversation. Required for the SPF check. | [optional] [default to null]
**Helo** | **string** | the identification provided by the client mta in the MTA state of the SMTP conversation. Required for the SPF check. | [optional] [default to null]
**ClientAddress** | **string** | the remote MTA&#x27;s ip address. Optional; required for the SPF check. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

