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

// the individual webhook info
type InlineResponse20074 struct {
	// a unique integer indentifier for the webhook
	Id int32 `json:"id,omitempty"`
	// The URL that the event data will be posted to
	Url string `json:"url,omitempty"`
	// a description of the webhook
	Description string `json:"description,omitempty"`
	// the key used to requests for this webhook
	AuthKey string `json:"auth_key,omitempty"`
	// The message events that will be posted to the hook
	Events []string `json:"events,omitempty"`
	// the date and time that the webhook was created as a UTC string in YYYY-MM-DD HH:MM:SS format
	CreatedAt time.Time `json:"created_at,omitempty"`
	// the date and time that the webhook last successfully received events as a UTC string in YYYY-MM-DD HH:MM:SS format
	LastSentAt time.Time `json:"last_sent_at,omitempty"`
	// the number of event batches that have ever been sent to this webhook
	BatchesSent int32 `json:"batches_sent,omitempty"`
	// the total number of events that have ever been sent to this webhook
	EventsSent int32 `json:"events_sent,omitempty"`
	// if we've ever gotten an error trying to post to this webhook, the last error that we've seen
	LastError string `json:"last_error,omitempty"`
}
