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

// the individual mailbox route
type InlineResponse20012 struct {
	// the unique identifier of the route
	Id string `json:"id,omitempty"`
	// the search pattern that the mailbox name should match
	Pattern string `json:"pattern,omitempty"`
	// the webhook URL where inbound messages will be published
	Url string `json:"url,omitempty"`
}
