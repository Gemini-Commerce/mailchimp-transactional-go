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

type Body30 struct {
	// a valid api key
	Key string `json:"key"`
	// a dedicated ip address
	Ip string `json:"ip"`
	// a domain name to set as the dedicated IP's custom dns name.
	Domain string `json:"domain"`
}
