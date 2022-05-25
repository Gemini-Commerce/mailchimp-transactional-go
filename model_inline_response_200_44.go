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

import (
	"time"
)

// the information on each sending domain for the account
type InlineResponse20044 struct {
	// the sender domain name
	Domain string `json:"domain,omitempty"`
	// the date and time that the sending domain was first seen as a UTC string in YYYY-MM-DD HH:MM:SS format
	CreatedAt time.Time `json:"created_at,omitempty"`
	// when the domain's DNS settings were last tested as a UTC string in YYYY-MM-DD HH:MM:SS format
	LastTestedAt time.Time `json:"last_tested_at,omitempty"`
	Spf *SendersdomainsSpf `json:"spf,omitempty"`
	Dkim *SendersdomainsDkim `json:"dkim,omitempty"`
	// if the domain has been verified, this indicates when that verification occurred as a UTC string in YYYY-MM-DD HH:MM:SS format
	VerifiedAt time.Time `json:"verified_at,omitempty"`
	// whether this domain can be used to authenticate mail, either for itself or as a custom signing domain. If this is false but spf and dkim are both valid, you will need to verify the domain before using it to authenticate mail
	ValidSigning bool `json:"valid_signing,omitempty"`
	// a unique key used to verify a domain by adding a TXT record. Append this key to 'mandrill_verify.' and add it to your domain's TXT records to verify
	VerifyTxtKey string `json:"verify_txt_key,omitempty"`
}