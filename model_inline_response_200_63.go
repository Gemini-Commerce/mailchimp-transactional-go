/*
 * Mailchimp Transactional API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.47
 * Contact: apihelp@mailchimp.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package mcapitransactional
import (
	"time"
)

// the template that was published
type InlineResponse20063 struct {
	// the immutable unique code name of the template
	Slug string `json:"slug,omitempty"`
	// the name of the template
	Name string `json:"name,omitempty"`
	// the list of labels applied to the template
	Labels []string `json:"labels,omitempty"`
	// the full HTML code of the template, with mc:edit attributes marking the editable elements - draft version
	Code string `json:"code,omitempty"`
	// the subject line of the template, if provided - draft version
	Subject string `json:"subject,omitempty"`
	// the default sender address for the template, if provided - draft version
	FromEmail string `json:"from_email,omitempty"`
	// the default sender from name for the template, if provided - draft version
	FromName string `json:"from_name,omitempty"`
	// the default text part of messages sent with the template, if provided - draft version
	Text string `json:"text,omitempty"`
	// the same as the template name - kept as a separate field for backwards compatibility
	PublishName string `json:"publish_name,omitempty"`
	// the full HTML code of the template, with mc:edit attributes marking the editable elements that are available as published, if it has been published
	PublishCode string `json:"publish_code,omitempty"`
	// the subject line of the template, if provided
	PublishSubject string `json:"publish_subject,omitempty"`
	// the default sender address for the template, if provided
	PublishFromEmail string `json:"publish_from_email,omitempty"`
	// the default sender from name for the template, if provided
	PublishFromName string `json:"publish_from_name,omitempty"`
	// the default text part of messages sent with the template, if provided
	PublishText string `json:"publish_text,omitempty"`
	// the date and time the template was last published as a UTC string in YYYY-MM-DD HH:MM:SS format, or null if it has not been published
	PublishedAt time.Time `json:"published_at,omitempty"`
	// the date and time the template was first created as a UTC string in YYYY-MM-DD HH:MM:SS format
	CreatedAt time.Time `json:"created_at,omitempty"`
	// the date and time the template was last modified as a UTC string in YYYY-MM-DD HH:MM:SS format
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
