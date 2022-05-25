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

type Body53 struct {
	// a valid api key
	Key string `json:"key"`
	// domain name at which you can receive email
	Domain string `json:"domain"`
	// a mailbox at the domain where the verification email should be sent
	Mailbox string `json:"mailbox"`
}
