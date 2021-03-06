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

// a description of the provisioning request that was created
type InlineResponse20019 struct {
	// the date and time that the request was created as a UTC timestamp in YYYY-MM-DD HH:MM:SS format
	RequestedAt time.Time `json:"requested_at,omitempty"`
}
