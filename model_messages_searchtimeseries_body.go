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

type MessagesSearchtimeseriesBody struct {
	// a valid api key
	Key string `json:"key"`
	// the search terms to find matching messages for
	Query string `json:"query,omitempty"`
	// start date
	DateFrom time.Time `json:"date_from,omitempty"`
	// end date
	DateTo time.Time `json:"date_to,omitempty"`
	// an array of tag names to narrow the search to, will return messages that contain ANY of the tags
	Tags []string `json:"tags,omitempty"`
	// an array of sender addresses to narrow the search to, will return messages sent by ANY of the senders
	Senders []string `json:"senders,omitempty"`
}
