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

import (
	"time"
)

type Body32 struct {
	// a valid api key
	Key string `json:"key"`
	// the immutable name or slug of a template that exists in the user's account. For backwards-compatibility, the template name may also be used but the immutable slug is preferred.
	TemplateName string `json:"template_name"`
	// an array of template content to send. Each item in the array should be a struct with two keys - name: the name of the content block to set the content for, and content: the actual content to put into the block
	TemplateContent []MessagessendtemplateTemplateContent `json:"template_content"`
	Message *MessagessendtemplateMessage `json:"message"`
	// enable a background sending mode that is optimized for bulk sending. In async mode, messages/send will immediately return a status of \"queued\" for every recipient. To handle rejections when sending in async mode, set up a webhook for the 'reject' event. Defaults to false for messages with no more than 10 recipients; messages with more than 10 recipients are always sent asynchronously, regardless of the value of async.
	Async bool `json:"async,omitempty"`
	// the name of the dedicated ip pool that should be used to send the message. If you do not have any dedicated IPs, this parameter has no effect. If you specify a pool that does not exist, your default pool will be used instead.
	IpPool string `json:"ip_pool,omitempty"`
	// when this message should be sent as a UTC timestamp in YYYY-MM-DD HH:MM:SS format. If you specify a time in the past, the message will be sent immediately.
	SendAt time.Time `json:"send_at,omitempty"`
}