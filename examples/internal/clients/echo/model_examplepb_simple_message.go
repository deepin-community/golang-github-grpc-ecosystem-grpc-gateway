/*
 * Echo Service
 *
 * Echo Service API consists of a single service which returns a message.
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package echo

// SimpleMessage represents a simple message sent to the Echo service.
type ExamplepbSimpleMessage struct {
	// Id represents the message identifier.
	Id string `json:"id,omitempty"`
	Num string `json:"num,omitempty"`
	LineNum string `json:"lineNum,omitempty"`
	Lang string `json:"lang,omitempty"`
	Status *ExamplepbEmbedded `json:"status,omitempty"`
	En string `json:"en,omitempty"`
	No *ExamplepbEmbedded `json:"no,omitempty"`
	ResourceId string `json:"resourceId,omitempty"`
	NId *ExamplepbNestedMessage `json:"nId,omitempty"`
}
