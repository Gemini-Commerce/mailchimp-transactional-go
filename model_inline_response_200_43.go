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

// the information on each sending address in the account
type InlineResponse20043 struct {
	// the sender's email address
	Address string `json:"address,omitempty"`
	// the date and time that the sender was first seen by Mandrill as a UTC date string in YYYY-MM-DD HH:MM:SS format
	CreatedAt time.Time `json:"created_at,omitempty"`
	// the total number of messages sent by this sender
	Sent int32 `json:"sent,omitempty"`
	// the total number of hard bounces by messages by this sender
	HardBounces int32 `json:"hard_bounces,omitempty"`
	// the total number of soft bounces by messages by this sender
	SoftBounces int32 `json:"soft_bounces,omitempty"`
	// the total number of rejected messages by this sender
	Rejects int32 `json:"rejects,omitempty"`
	// the total number of spam complaints received for messages by this sender
	Complaints int32 `json:"complaints,omitempty"`
	// the total number of unsubscribe requests received for messages by this sender
	Unsubs int32 `json:"unsubs,omitempty"`
	// the total number of times messages by this sender have been opened
	Opens int32 `json:"opens,omitempty"`
	// the total number of times tracked URLs in messages by this sender have been clicked
	Clicks int32 `json:"clicks,omitempty"`
	// the number of unique opens for emails sent for this sender
	UniqueOpens int32 `json:"unique_opens,omitempty"`
	// the number of unique clicks for emails sent for this sender
	UniqueClicks int32 `json:"unique_clicks,omitempty"`
}
