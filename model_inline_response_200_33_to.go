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

// the message recipient's information
type InlineResponse20033To struct {
	// the email address of the recipient
	Email string `json:"email,omitempty"`
	// the alias of the recipient (if any)
	Name string `json:"name,omitempty"`
}
