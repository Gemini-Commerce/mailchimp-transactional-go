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

// the parsed message
type InlineResponse20034 struct {
	// the subject of the message
	Subject string `json:"subject,omitempty"`
	// the email address of the sender
	FromEmail string `json:"from_email,omitempty"`
	// the alias of the sender (if any)
	FromName string `json:"from_name,omitempty"`
	// an array of any recipients in the message
	To []InlineResponse20034To `json:"to,omitempty"`
	// the key-value pairs of the MIME headers for the message's main document
	Headers *interface{} `json:"headers,omitempty"`
	// the text part of the message, if any
	Text string `json:"text,omitempty"`
	// the HTML part of the message, if any
	Html string `json:"html,omitempty"`
	// an array of any attachments that can be found in the message
	Attachments []InlineResponse20034Attachments `json:"attachments,omitempty"`
	// an array of any embedded images that can be found in the message
	Images []InlineResponse20034Images `json:"images,omitempty"`
}
