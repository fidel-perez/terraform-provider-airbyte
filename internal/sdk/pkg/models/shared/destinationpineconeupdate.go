// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type DestinationPineconeUpdateEmbeddingFakeMode string

const (
	DestinationPineconeUpdateEmbeddingFakeModeFake DestinationPineconeUpdateEmbeddingFakeMode = "fake"
)

func (e DestinationPineconeUpdateEmbeddingFakeMode) ToPointer() *DestinationPineconeUpdateEmbeddingFakeMode {
	return &e
}

func (e *DestinationPineconeUpdateEmbeddingFakeMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fake":
		*e = DestinationPineconeUpdateEmbeddingFakeMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPineconeUpdateEmbeddingFakeMode: %v", v)
	}
}

// DestinationPineconeUpdateEmbeddingFake - Use a fake embedding made out of random vectors with 1536 embedding dimensions. This is useful for testing the data pipeline without incurring any costs.
type DestinationPineconeUpdateEmbeddingFake struct {
	Mode *DestinationPineconeUpdateEmbeddingFakeMode `json:"mode,omitempty"`
}

type DestinationPineconeUpdateEmbeddingCohereMode string

const (
	DestinationPineconeUpdateEmbeddingCohereModeCohere DestinationPineconeUpdateEmbeddingCohereMode = "cohere"
)

func (e DestinationPineconeUpdateEmbeddingCohereMode) ToPointer() *DestinationPineconeUpdateEmbeddingCohereMode {
	return &e
}

func (e *DestinationPineconeUpdateEmbeddingCohereMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cohere":
		*e = DestinationPineconeUpdateEmbeddingCohereMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPineconeUpdateEmbeddingCohereMode: %v", v)
	}
}

// DestinationPineconeUpdateEmbeddingCohere - Use the Cohere API to embed text.
type DestinationPineconeUpdateEmbeddingCohere struct {
	CohereKey string                                        `json:"cohere_key"`
	Mode      *DestinationPineconeUpdateEmbeddingCohereMode `json:"mode,omitempty"`
}

type DestinationPineconeUpdateEmbeddingOpenAIMode string

const (
	DestinationPineconeUpdateEmbeddingOpenAIModeOpenai DestinationPineconeUpdateEmbeddingOpenAIMode = "openai"
)

func (e DestinationPineconeUpdateEmbeddingOpenAIMode) ToPointer() *DestinationPineconeUpdateEmbeddingOpenAIMode {
	return &e
}

func (e *DestinationPineconeUpdateEmbeddingOpenAIMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "openai":
		*e = DestinationPineconeUpdateEmbeddingOpenAIMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPineconeUpdateEmbeddingOpenAIMode: %v", v)
	}
}

// DestinationPineconeUpdateEmbeddingOpenAI - Use the OpenAI API to embed text. This option is using the text-embedding-ada-002 model with 1536 embedding dimensions.
type DestinationPineconeUpdateEmbeddingOpenAI struct {
	Mode      *DestinationPineconeUpdateEmbeddingOpenAIMode `json:"mode,omitempty"`
	OpenaiKey string                                        `json:"openai_key"`
}

type DestinationPineconeUpdateEmbeddingType string

const (
	DestinationPineconeUpdateEmbeddingTypeDestinationPineconeUpdateEmbeddingOpenAI DestinationPineconeUpdateEmbeddingType = "destination-pinecone-update_Embedding_OpenAI"
	DestinationPineconeUpdateEmbeddingTypeDestinationPineconeUpdateEmbeddingCohere DestinationPineconeUpdateEmbeddingType = "destination-pinecone-update_Embedding_Cohere"
	DestinationPineconeUpdateEmbeddingTypeDestinationPineconeUpdateEmbeddingFake   DestinationPineconeUpdateEmbeddingType = "destination-pinecone-update_Embedding_Fake"
)

type DestinationPineconeUpdateEmbedding struct {
	DestinationPineconeUpdateEmbeddingOpenAI *DestinationPineconeUpdateEmbeddingOpenAI
	DestinationPineconeUpdateEmbeddingCohere *DestinationPineconeUpdateEmbeddingCohere
	DestinationPineconeUpdateEmbeddingFake   *DestinationPineconeUpdateEmbeddingFake

	Type DestinationPineconeUpdateEmbeddingType
}

func CreateDestinationPineconeUpdateEmbeddingDestinationPineconeUpdateEmbeddingOpenAI(destinationPineconeUpdateEmbeddingOpenAI DestinationPineconeUpdateEmbeddingOpenAI) DestinationPineconeUpdateEmbedding {
	typ := DestinationPineconeUpdateEmbeddingTypeDestinationPineconeUpdateEmbeddingOpenAI

	return DestinationPineconeUpdateEmbedding{
		DestinationPineconeUpdateEmbeddingOpenAI: &destinationPineconeUpdateEmbeddingOpenAI,
		Type:                                     typ,
	}
}

func CreateDestinationPineconeUpdateEmbeddingDestinationPineconeUpdateEmbeddingCohere(destinationPineconeUpdateEmbeddingCohere DestinationPineconeUpdateEmbeddingCohere) DestinationPineconeUpdateEmbedding {
	typ := DestinationPineconeUpdateEmbeddingTypeDestinationPineconeUpdateEmbeddingCohere

	return DestinationPineconeUpdateEmbedding{
		DestinationPineconeUpdateEmbeddingCohere: &destinationPineconeUpdateEmbeddingCohere,
		Type:                                     typ,
	}
}

func CreateDestinationPineconeUpdateEmbeddingDestinationPineconeUpdateEmbeddingFake(destinationPineconeUpdateEmbeddingFake DestinationPineconeUpdateEmbeddingFake) DestinationPineconeUpdateEmbedding {
	typ := DestinationPineconeUpdateEmbeddingTypeDestinationPineconeUpdateEmbeddingFake

	return DestinationPineconeUpdateEmbedding{
		DestinationPineconeUpdateEmbeddingFake: &destinationPineconeUpdateEmbeddingFake,
		Type:                                   typ,
	}
}

func (u *DestinationPineconeUpdateEmbedding) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationPineconeUpdateEmbeddingFake := new(DestinationPineconeUpdateEmbeddingFake)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPineconeUpdateEmbeddingFake); err == nil {
		u.DestinationPineconeUpdateEmbeddingFake = destinationPineconeUpdateEmbeddingFake
		u.Type = DestinationPineconeUpdateEmbeddingTypeDestinationPineconeUpdateEmbeddingFake
		return nil
	}

	destinationPineconeUpdateEmbeddingOpenAI := new(DestinationPineconeUpdateEmbeddingOpenAI)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPineconeUpdateEmbeddingOpenAI); err == nil {
		u.DestinationPineconeUpdateEmbeddingOpenAI = destinationPineconeUpdateEmbeddingOpenAI
		u.Type = DestinationPineconeUpdateEmbeddingTypeDestinationPineconeUpdateEmbeddingOpenAI
		return nil
	}

	destinationPineconeUpdateEmbeddingCohere := new(DestinationPineconeUpdateEmbeddingCohere)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPineconeUpdateEmbeddingCohere); err == nil {
		u.DestinationPineconeUpdateEmbeddingCohere = destinationPineconeUpdateEmbeddingCohere
		u.Type = DestinationPineconeUpdateEmbeddingTypeDestinationPineconeUpdateEmbeddingCohere
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationPineconeUpdateEmbedding) MarshalJSON() ([]byte, error) {
	if u.DestinationPineconeUpdateEmbeddingFake != nil {
		return json.Marshal(u.DestinationPineconeUpdateEmbeddingFake)
	}

	if u.DestinationPineconeUpdateEmbeddingOpenAI != nil {
		return json.Marshal(u.DestinationPineconeUpdateEmbeddingOpenAI)
	}

	if u.DestinationPineconeUpdateEmbeddingCohere != nil {
		return json.Marshal(u.DestinationPineconeUpdateEmbeddingCohere)
	}

	return nil, nil
}

// DestinationPineconeUpdateIndexing - Pinecone is a popular vector store that can be used to store and retrieve embeddings.
type DestinationPineconeUpdateIndexing struct {
	// Pinecone index to use
	Index string `json:"index"`
	// Pinecone environment to use
	PineconeEnvironment string `json:"pinecone_environment"`
	PineconeKey         string `json:"pinecone_key"`
}

type DestinationPineconeUpdateProcessingConfigModel struct {
	// Size of overlap between chunks in tokens to store in vector store to better capture relevant context
	ChunkOverlap *int64 `json:"chunk_overlap,omitempty"`
	// Size of chunks in tokens to store in vector store (make sure it is not too big for the context if your LLM)
	ChunkSize int64 `json:"chunk_size"`
	// List of fields in the record that should be stored as metadata. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered metadata fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. `user.name` will access the `name` field in the `user` object. It's also possible to use wildcards to access all fields in an object, e.g. `users.*.name` will access all `names` fields in all entries of the `users` array. When specifying nested paths, all matching values are flattened into an array set to a field named by the path.
	MetadataFields []string `json:"metadata_fields,omitempty"`
	// List of fields in the record that should be used to calculate the embedding. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered text fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. `user.name` will access the `name` field in the `user` object. It's also possible to use wildcards to access all fields in an object, e.g. `users.*.name` will access all `names` fields in all entries of the `users` array.
	TextFields []string `json:"text_fields,omitempty"`
}

type DestinationPineconeUpdate struct {
	// Embedding configuration
	Embedding DestinationPineconeUpdateEmbedding `json:"embedding"`
	// Pinecone is a popular vector store that can be used to store and retrieve embeddings.
	Indexing   DestinationPineconeUpdateIndexing              `json:"indexing"`
	Processing DestinationPineconeUpdateProcessingConfigModel `json:"processing"`
}
