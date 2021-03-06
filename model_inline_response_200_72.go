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

// the user information including username, key, reputation, quota, and historical sending stats
type InlineResponse20072 struct {
	// the username of the user (used for SMTP authentication)
	Username string `json:"username,omitempty"`
	// the date and time that the user's Mandrill account was created as a UTC string in YYYY-MM-DD HH:MM:SS format
	CreatedAt string `json:"created_at,omitempty"`
	// a unique, permanent identifier for this user
	PublicId string `json:"public_id,omitempty"`
	// the reputation of the user on a scale from 0 to 100, with 75 generally being a \"good\" reputation
	Reputation int32 `json:"reputation,omitempty"`
	// the maximum number of emails Mandrill will deliver for this user each hour. Any emails beyond that will be accepted and queued for later delivery. Users with higher reputations will have higher hourly quotas
	HourlyQuota int32 `json:"hourly_quota,omitempty"`
	// the number of emails that are queued for delivery due to exceeding your monthly or hourly quotas
	Backlog int32 `json:"backlog,omitempty"`
	Stats *InlineResponse20072Stats `json:"stats,omitempty"`
}
