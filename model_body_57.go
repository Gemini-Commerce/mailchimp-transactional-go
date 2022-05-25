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

type Body57 struct {
	// a valid api key
	Key string `json:"key"`
	// a unique identifier for the subaccount to be used in sending calls
	Id string `json:"id"`
	// an optional display name to further identify the subaccount
	Name string `json:"name,omitempty"`
	// optional extra text to associate with the subaccount
	Notes string `json:"notes,omitempty"`
	// an optional manual hourly quota for the subaccount. If not specified, Mandrill will manage this based on reputation
	CustomQuota int32 `json:"custom_quota,omitempty"`
}
