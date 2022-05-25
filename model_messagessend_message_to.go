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

// a single recipient's information
type MessagessendMessageTo struct {
	// the email address of the recipient
	Email string `json:"email"`
	// the optional display name to use for the recipient
	Name string `json:"name,omitempty"`
	// the header type to use for the recipient, defaults to \"to\" if not provided
	Type_ string `json:"type,omitempty"`
}
