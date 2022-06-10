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

type ExportsActivityBody struct {
	// a valid api key
	Key string `json:"key"`
	// an optional email address to notify when the export job has finished
	NotifyEmail string `json:"notify_email,omitempty"`
	// start date as a UTC string in YYYY-MM-DD HH:MM:SS format
	DateFrom time.Time `json:"date_from,omitempty"`
	// end date as a UTC string in YYYY-MM-DD HH:MM:SS format
	DateTo time.Time `json:"date_to,omitempty"`
	// an array of tag names to narrow the export to; will match messages that contain ANY of the tags
	Tags []string `json:"tags,omitempty"`
	// an array of senders to narrow the export to
	Senders []string `json:"senders,omitempty"`
	// an array of message states to narrow the export to; messages with ANY of the states will be included
	States []string `json:"states,omitempty"`
	// an array of api keys to narrow the export to; messsagse sent with ANY of the keys will be included
	ApiKeys []string `json:"api_keys,omitempty"`
}