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

// the stats for a single hour
type InlineResponse20031 struct {
	// the hour as a UTC date string in YYYY-MM-DD HH:MM:SS format
	Time time.Time `json:"time,omitempty"`
	// the number of emails that were sent during the hour
	Sent int32 `json:"sent,omitempty"`
	// the number of emails that hard bounced during the hour
	HardBounces int32 `json:"hard_bounces,omitempty"`
	// the number of emails that soft bounced during the hour
	SoftBounces int32 `json:"soft_bounces,omitempty"`
	// the number of emails that were rejected during the hour
	Rejects int32 `json:"rejects,omitempty"`
	// the number of spam complaints received during the hour
	Complaints int32 `json:"complaints,omitempty"`
	// the number of unsubscribes received during the hour
	Unsubs int32 `json:"unsubs,omitempty"`
	// the number of emails opened during the hour
	Opens int32 `json:"opens,omitempty"`
	// the number of unique opens generated by messages sent during the hour
	UniqueOpens int32 `json:"unique_opens,omitempty"`
	// the number of tracked URLs clicked during the hour
	Clicks int32 `json:"clicks,omitempty"`
	// the number of unique clicks generated by messages sent during the hour
	UniqueClicks int32 `json:"unique_clicks,omitempty"`
}
