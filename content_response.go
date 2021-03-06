/*
 * Anchore Engine API Server
 *
 * This is the Anchore Engine API. Provides the primary external API for users of the service.
 *
 * API version: 0.1.5
 * Contact: nurmi@anchore.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package anchore_engine

// Generic wrapper for content listings from images
type ContentResponse struct {
	ImageDigest string `json:"imageDigest,omitempty"`

	ContentType string `json:"content_type,omitempty"`

	Content []interface{} `json:"content,omitempty"`
}
