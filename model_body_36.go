/*
 * Mailchimp Transactional API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.47
 * Contact: apihelp@mailchimp.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package MailchimpTransactional

type Body36 struct {
	// a valid api key
	Key string `json:"key"`
	// the unique id of the message to get - passed as the \"_id\" field in webhooks, send calls, or search calls
	Id string `json:"id"`
}
