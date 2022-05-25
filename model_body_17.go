/*
 * Mailchimp Transactional API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.47
 * Contact: apihelp@mailchimp.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package MailchimpTransactional

type Body17 struct {
	// a valid api key
	Key string `json:"key"`
	// the full MIME document of an email message
	RawMessage string `json:"raw_message"`
	// optionally define the recipients to receive the message - otherwise we'll use the To, Cc, and Bcc headers provided in the document
	To []string `json:"to,omitempty"`
	// the address specified in the MAIL FROM stage of the SMTP conversation. Required for the SPF check.
	MailFrom string `json:"mail_from,omitempty"`
	// the identification provided by the client mta in the MTA state of the SMTP conversation. Required for the SPF check.
	Helo string `json:"helo,omitempty"`
	// the remote MTA's ip address. Optional; required for the SPF check.
	ClientAddress string `json:"client_address,omitempty"`
}