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

type Body15 struct {
	// a valid api key
	Key string `json:"key"`
	// the unique identifier of an existing mailbox route
	Id string `json:"id"`
	// the search pattern that the mailbox name should match
	Pattern string `json:"pattern,omitempty"`
	// the webhook URL where the inbound messages will be published; Validation: webhookexists
	Url string `json:"url,omitempty"`
}
