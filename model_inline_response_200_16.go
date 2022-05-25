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

// the individual recipient information
type InlineResponse20016 struct {
	// the email address of the matching recipient
	Email string `json:"email,omitempty"`
	// the mailbox route pattern that the recipient matched
	Pattern string `json:"pattern,omitempty"`
	// the webhook URL that the message was posted to
	Url string `json:"url,omitempty"`
}