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

// stats with this tag so far today
type InlineResponse20059StatsToday struct {
	// the number of emails sent with this tag so far today
	Sent int32 `json:"sent,omitempty"`
	// the number of emails hard bounced with this tag so far today
	HardBounces int32 `json:"hard_bounces,omitempty"`
	// the number of emails soft bounced with this tag so far today
	SoftBounces int32 `json:"soft_bounces,omitempty"`
	// the number of emails rejected for sending this sender so far today
	Rejects int32 `json:"rejects,omitempty"`
	// the number of spam complaints with this tag so far today
	Complaints int32 `json:"complaints,omitempty"`
	// the number of unsubscribes with this tag so far today
	Unsubs int32 `json:"unsubs,omitempty"`
	// the number of times emails have been opened with this tag so far today
	Opens int32 `json:"opens,omitempty"`
	// the number of unique opens for emails sent with this tag so far today
	UniqueOpens int32 `json:"unique_opens,omitempty"`
	// the number of URLs that have been clicked with this tag so far today
	Clicks int32 `json:"clicks,omitempty"`
	// the number of unique clicks for emails sent with this tag so far today
	UniqueClicks int32 `json:"unique_clicks,omitempty"`
}