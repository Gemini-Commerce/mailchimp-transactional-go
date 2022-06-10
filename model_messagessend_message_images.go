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

// a single embedded image
type MessagessendMessageImages struct {
	// the MIME type of the image - must start with \"image/\"
	Type_ string `json:"type,omitempty"`
	// the Content ID of the image - use <img src=\"cid:THIS_VALUE\"> to reference the image in your HTML content
	Name string `json:"name,omitempty"`
	// the content of the image as a base64-encoded string
	Content string `json:"content,omitempty"`
}
