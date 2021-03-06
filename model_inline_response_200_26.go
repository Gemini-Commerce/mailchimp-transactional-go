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

// information about the dedicated IP's new configuration
type InlineResponse20026 struct {
	// whether the domain name has a correctly-configured A record pointing to the ip address
	Valid string `json:"valid,omitempty"`
	// if valid is false, this will contain details about why the domain's A record is incorrect
	Error_ string `json:"error,omitempty"`
}
