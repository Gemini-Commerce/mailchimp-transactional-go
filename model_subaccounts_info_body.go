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

type SubaccountsInfoBody struct {
	// a valid api key
	Key string `json:"key"`
	// the unique identifier of the subaccount to query
	Id string `json:"id"`
}
