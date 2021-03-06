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

type MessagesSendBody struct {
	// a valid api key
	Key string `json:"key"`
	Message *MessagessendMessage `json:"message"`
	// enable a background sending mode that is optimized for bulk sending. In async mode, messages/send will immediately return a status of \"queued\" for every recipient. To handle rejections when sending in async mode, set up a webhook for the 'reject' event. Defaults to false for messages with no more than 10 recipients; messages with more than 10 recipients are always sent asynchronously, regardless of the value of async.
	Async bool `json:"async,omitempty"`
	// the name of the dedicated ip pool that should be used to send the message. If you do not have any dedicated IPs, this parameter has no effect. If you specify a pool that does not exist, your default pool will be used instead.
	IpPool string `json:"ip_pool,omitempty"`
	// when this message should be sent as a UTC timestamp in YYYY-MM-DD HH:MM:SS format. If you specify a time in the past, the message will be sent immediately.
	SendAt time.Time `json:"send_at,omitempty"`
}
