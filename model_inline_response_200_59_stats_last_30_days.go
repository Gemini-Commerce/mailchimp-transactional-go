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

// stats with this tag in the last 30 days
type InlineResponse20059StatsLast30Days struct {
	// the number of emails sent with this tag in the last 30 days
	Sent int32 `json:"sent,omitempty"`
	// the number of emails hard bounced with this tag in the last 30 days
	HardBounces int32 `json:"hard_bounces,omitempty"`
	// the number of emails soft bounced with this tag in the last 30 days
	SoftBounces int32 `json:"soft_bounces,omitempty"`
	// the number of emails rejected for sending this sender in the last 30 days
	Rejects int32 `json:"rejects,omitempty"`
	// the number of spam complaints with this tag in the last 30 days
	Complaints int32 `json:"complaints,omitempty"`
	// the number of unsubscribes with this tag in the last 30 days
	Unsubs int32 `json:"unsubs,omitempty"`
	// the number of times emails have been opened with this tag in the last 30 days
	Opens int32 `json:"opens,omitempty"`
	// the number of unique opens for emails sent with this tag in the last 30 days
	UniqueOpens int32 `json:"unique_opens,omitempty"`
	// the number of URLs that have been clicked with this tag in the last 30 days
	Clicks int32 `json:"clicks,omitempty"`
	// the number of unique clicks for emails sent with this tag in the last 30 days
	UniqueClicks int32 `json:"unique_clicks,omitempty"`
}
