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

type MessagessendMessageVars struct {
	// the merge variable's name. Merge variable names are case-insensitive and may not start with _
	Name string `json:"name,omitempty"`
	// the merge variable's content
	Content string `json:"content,omitempty"`
}
