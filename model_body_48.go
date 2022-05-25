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

type Body48 struct {
	// a valid api key
	Key string `json:"key"`
	// an email address
	Email string `json:"email"`
	// an optional unique identifier for the subaccount to limit the denylist deletion
	Subaccount string `json:"subaccount,omitempty"`
}
