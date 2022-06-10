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

type SubaccountsListBody struct {
	// a valid api key
	Key string `json:"key"`
	// an optional prefix to filter the subaccounts' ids and names
	Q string `json:"q,omitempty"`
}