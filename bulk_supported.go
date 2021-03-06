/*
 * SCIM API
 *
 * SCIM V2 API implemented by RingCentral
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package scim

type BulkSupported struct {
	MaxOperations int32 `json:"maxOperations,omitempty"`

	MaxPayloadSize int32 `json:"maxPayloadSize,omitempty"`

	Supported bool `json:"supported,omitempty"`
}
