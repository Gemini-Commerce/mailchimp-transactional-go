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

// a single supported attachment
type MessagessendMessageAttachments struct {
	// the MIME type of the attachment
	Type_ string `json:"type,omitempty"`
	// the file name of the attachment
	Name string `json:"name,omitempty"`
	// the content of the attachment as a base64-encoded string
	Content string `json:"content,omitempty"`
}