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

type WebhooksAddBody struct {
	// a valid api key
	Key string `json:"key"`
	// the URL to POST batches of events
	Url string `json:"url"`
	// an optional description of the webhook
	Description string `json:"description,omitempty"`
	// an optional list of events that will be posted to the webhook
	Events []string `json:"events,omitempty"`
}
