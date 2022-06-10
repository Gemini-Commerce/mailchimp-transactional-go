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

// a status object containing the address and the result of the operation
type InlineResponse20040 struct {
	// the email address you provided
	Email string `json:"email,omitempty"`
	// whether the operation succeeded
	Added bool `json:"added,omitempty"`
}
