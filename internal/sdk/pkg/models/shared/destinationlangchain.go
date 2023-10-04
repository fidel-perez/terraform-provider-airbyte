// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
)

type DestinationLangchainLangchain string

const (
	DestinationLangchainLangchainLangchain DestinationLangchainLangchain = "langchain"
)

func (e DestinationLangchainLangchain) ToPointer() *DestinationLangchainLangchain {
	return &e
}

func (e *DestinationLangchainLangchain) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "langchain":
		*e = DestinationLangchainLangchain(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationLangchainLangchain: %v", v)
	}
}

type DestinationLangchainEmbeddingFakeMode string

const (
	DestinationLangchainEmbeddingFakeModeFake DestinationLangchainEmbeddingFakeMode = "fake"
)

func (e DestinationLangchainEmbeddingFakeMode) ToPointer() *DestinationLangchainEmbeddingFakeMode {
	return &e
}

func (e *DestinationLangchainEmbeddingFakeMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fake":
		*e = DestinationLangchainEmbeddingFakeMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationLangchainEmbeddingFakeMode: %v", v)
	}
}

// DestinationLangchainEmbeddingFake - Use a fake embedding made out of random vectors with 1536 embedding dimensions. This is useful for testing the data pipeline without incurring any costs.
type DestinationLangchainEmbeddingFake struct {
	mode *DestinationLangchainEmbeddingFakeMode `const:"fake" json:"mode"`
}

func (d DestinationLangchainEmbeddingFake) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationLangchainEmbeddingFake) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationLangchainEmbeddingFake) GetMode() *DestinationLangchainEmbeddingFakeMode {
	return DestinationLangchainEmbeddingFakeModeFake.ToPointer()
}

type DestinationLangchainEmbeddingOpenAIMode string

const (
	DestinationLangchainEmbeddingOpenAIModeOpenai DestinationLangchainEmbeddingOpenAIMode = "openai"
)

func (e DestinationLangchainEmbeddingOpenAIMode) ToPointer() *DestinationLangchainEmbeddingOpenAIMode {
	return &e
}

func (e *DestinationLangchainEmbeddingOpenAIMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "openai":
		*e = DestinationLangchainEmbeddingOpenAIMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationLangchainEmbeddingOpenAIMode: %v", v)
	}
}

// DestinationLangchainEmbeddingOpenAI - Use the OpenAI API to embed text. This option is using the text-embedding-ada-002 model with 1536 embedding dimensions.
type DestinationLangchainEmbeddingOpenAI struct {
	mode      *DestinationLangchainEmbeddingOpenAIMode `const:"openai" json:"mode"`
	OpenaiKey string                                   `json:"openai_key"`
}

func (d DestinationLangchainEmbeddingOpenAI) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationLangchainEmbeddingOpenAI) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationLangchainEmbeddingOpenAI) GetMode() *DestinationLangchainEmbeddingOpenAIMode {
	return DestinationLangchainEmbeddingOpenAIModeOpenai.ToPointer()
}

func (o *DestinationLangchainEmbeddingOpenAI) GetOpenaiKey() string {
	if o == nil {
		return ""
	}
	return o.OpenaiKey
}

type DestinationLangchainEmbeddingType string

const (
	DestinationLangchainEmbeddingTypeDestinationLangchainEmbeddingOpenAI DestinationLangchainEmbeddingType = "destination-langchain_Embedding_OpenAI"
	DestinationLangchainEmbeddingTypeDestinationLangchainEmbeddingFake   DestinationLangchainEmbeddingType = "destination-langchain_Embedding_Fake"
)

type DestinationLangchainEmbedding struct {
	DestinationLangchainEmbeddingOpenAI *DestinationLangchainEmbeddingOpenAI
	DestinationLangchainEmbeddingFake   *DestinationLangchainEmbeddingFake

	Type DestinationLangchainEmbeddingType
}

func CreateDestinationLangchainEmbeddingDestinationLangchainEmbeddingOpenAI(destinationLangchainEmbeddingOpenAI DestinationLangchainEmbeddingOpenAI) DestinationLangchainEmbedding {
	typ := DestinationLangchainEmbeddingTypeDestinationLangchainEmbeddingOpenAI

	return DestinationLangchainEmbedding{
		DestinationLangchainEmbeddingOpenAI: &destinationLangchainEmbeddingOpenAI,
		Type:                                typ,
	}
}

func CreateDestinationLangchainEmbeddingDestinationLangchainEmbeddingFake(destinationLangchainEmbeddingFake DestinationLangchainEmbeddingFake) DestinationLangchainEmbedding {
	typ := DestinationLangchainEmbeddingTypeDestinationLangchainEmbeddingFake

	return DestinationLangchainEmbedding{
		DestinationLangchainEmbeddingFake: &destinationLangchainEmbeddingFake,
		Type:                              typ,
	}
}

func (u *DestinationLangchainEmbedding) UnmarshalJSON(data []byte) error {

	destinationLangchainEmbeddingFake := new(DestinationLangchainEmbeddingFake)
	if err := utils.UnmarshalJSON(data, &destinationLangchainEmbeddingFake, "", true, true); err == nil {
		u.DestinationLangchainEmbeddingFake = destinationLangchainEmbeddingFake
		u.Type = DestinationLangchainEmbeddingTypeDestinationLangchainEmbeddingFake
		return nil
	}

	destinationLangchainEmbeddingOpenAI := new(DestinationLangchainEmbeddingOpenAI)
	if err := utils.UnmarshalJSON(data, &destinationLangchainEmbeddingOpenAI, "", true, true); err == nil {
		u.DestinationLangchainEmbeddingOpenAI = destinationLangchainEmbeddingOpenAI
		u.Type = DestinationLangchainEmbeddingTypeDestinationLangchainEmbeddingOpenAI
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationLangchainEmbedding) MarshalJSON() ([]byte, error) {
	if u.DestinationLangchainEmbeddingOpenAI != nil {
		return utils.MarshalJSON(u.DestinationLangchainEmbeddingOpenAI, "", true)
	}

	if u.DestinationLangchainEmbeddingFake != nil {
		return utils.MarshalJSON(u.DestinationLangchainEmbeddingFake, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type DestinationLangchainIndexingChromaLocalPersistanceMode string

const (
	DestinationLangchainIndexingChromaLocalPersistanceModeChromaLocal DestinationLangchainIndexingChromaLocalPersistanceMode = "chroma_local"
)

func (e DestinationLangchainIndexingChromaLocalPersistanceMode) ToPointer() *DestinationLangchainIndexingChromaLocalPersistanceMode {
	return &e
}

func (e *DestinationLangchainIndexingChromaLocalPersistanceMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "chroma_local":
		*e = DestinationLangchainIndexingChromaLocalPersistanceMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationLangchainIndexingChromaLocalPersistanceMode: %v", v)
	}
}

// DestinationLangchainIndexingChromaLocalPersistance - Chroma is a popular vector store that can be used to store and retrieve embeddings. It will build its index in memory and persist it to disk by the end of the sync.
type DestinationLangchainIndexingChromaLocalPersistance struct {
	// Name of the collection to use.
	CollectionName *string `default:"langchain" json:"collection_name"`
	// Path to the directory where chroma files will be written. The files will be placed inside that local mount.
	DestinationPath string                                                  `json:"destination_path"`
	mode            *DestinationLangchainIndexingChromaLocalPersistanceMode `const:"chroma_local" json:"mode"`
}

func (d DestinationLangchainIndexingChromaLocalPersistance) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationLangchainIndexingChromaLocalPersistance) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationLangchainIndexingChromaLocalPersistance) GetCollectionName() *string {
	if o == nil {
		return nil
	}
	return o.CollectionName
}

func (o *DestinationLangchainIndexingChromaLocalPersistance) GetDestinationPath() string {
	if o == nil {
		return ""
	}
	return o.DestinationPath
}

func (o *DestinationLangchainIndexingChromaLocalPersistance) GetMode() *DestinationLangchainIndexingChromaLocalPersistanceMode {
	return DestinationLangchainIndexingChromaLocalPersistanceModeChromaLocal.ToPointer()
}

type DestinationLangchainIndexingDocArrayHnswSearchMode string

const (
	DestinationLangchainIndexingDocArrayHnswSearchModeDocArrayHnswSearch DestinationLangchainIndexingDocArrayHnswSearchMode = "DocArrayHnswSearch"
)

func (e DestinationLangchainIndexingDocArrayHnswSearchMode) ToPointer() *DestinationLangchainIndexingDocArrayHnswSearchMode {
	return &e
}

func (e *DestinationLangchainIndexingDocArrayHnswSearchMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DocArrayHnswSearch":
		*e = DestinationLangchainIndexingDocArrayHnswSearchMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationLangchainIndexingDocArrayHnswSearchMode: %v", v)
	}
}

// DestinationLangchainIndexingDocArrayHnswSearch - DocArrayHnswSearch is a lightweight Document Index implementation provided by Docarray that runs fully locally and is best suited for small- to medium-sized datasets. It stores vectors on disk in hnswlib, and stores all other data in SQLite.
type DestinationLangchainIndexingDocArrayHnswSearch struct {
	// Path to the directory where hnswlib and meta data files will be written. The files will be placed inside that local mount. All files in the specified destination directory will be deleted on each run.
	DestinationPath string                                              `json:"destination_path"`
	mode            *DestinationLangchainIndexingDocArrayHnswSearchMode `const:"DocArrayHnswSearch" json:"mode"`
}

func (d DestinationLangchainIndexingDocArrayHnswSearch) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationLangchainIndexingDocArrayHnswSearch) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationLangchainIndexingDocArrayHnswSearch) GetDestinationPath() string {
	if o == nil {
		return ""
	}
	return o.DestinationPath
}

func (o *DestinationLangchainIndexingDocArrayHnswSearch) GetMode() *DestinationLangchainIndexingDocArrayHnswSearchMode {
	return DestinationLangchainIndexingDocArrayHnswSearchModeDocArrayHnswSearch.ToPointer()
}

type DestinationLangchainIndexingPineconeMode string

const (
	DestinationLangchainIndexingPineconeModePinecone DestinationLangchainIndexingPineconeMode = "pinecone"
)

func (e DestinationLangchainIndexingPineconeMode) ToPointer() *DestinationLangchainIndexingPineconeMode {
	return &e
}

func (e *DestinationLangchainIndexingPineconeMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pinecone":
		*e = DestinationLangchainIndexingPineconeMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationLangchainIndexingPineconeMode: %v", v)
	}
}

// DestinationLangchainIndexingPinecone - Pinecone is a popular vector store that can be used to store and retrieve embeddings. It is a managed service and can also be queried from outside of langchain.
type DestinationLangchainIndexingPinecone struct {
	// Pinecone index to use
	Index string                                    `json:"index"`
	mode  *DestinationLangchainIndexingPineconeMode `const:"pinecone" json:"mode"`
	// Pinecone environment to use
	PineconeEnvironment string `json:"pinecone_environment"`
	PineconeKey         string `json:"pinecone_key"`
}

func (d DestinationLangchainIndexingPinecone) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationLangchainIndexingPinecone) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationLangchainIndexingPinecone) GetIndex() string {
	if o == nil {
		return ""
	}
	return o.Index
}

func (o *DestinationLangchainIndexingPinecone) GetMode() *DestinationLangchainIndexingPineconeMode {
	return DestinationLangchainIndexingPineconeModePinecone.ToPointer()
}

func (o *DestinationLangchainIndexingPinecone) GetPineconeEnvironment() string {
	if o == nil {
		return ""
	}
	return o.PineconeEnvironment
}

func (o *DestinationLangchainIndexingPinecone) GetPineconeKey() string {
	if o == nil {
		return ""
	}
	return o.PineconeKey
}

type DestinationLangchainIndexingType string

const (
	DestinationLangchainIndexingTypeDestinationLangchainIndexingPinecone               DestinationLangchainIndexingType = "destination-langchain_Indexing_Pinecone"
	DestinationLangchainIndexingTypeDestinationLangchainIndexingDocArrayHnswSearch     DestinationLangchainIndexingType = "destination-langchain_Indexing_DocArrayHnswSearch"
	DestinationLangchainIndexingTypeDestinationLangchainIndexingChromaLocalPersistance DestinationLangchainIndexingType = "destination-langchain_Indexing_Chroma (local persistance)"
)

type DestinationLangchainIndexing struct {
	DestinationLangchainIndexingPinecone               *DestinationLangchainIndexingPinecone
	DestinationLangchainIndexingDocArrayHnswSearch     *DestinationLangchainIndexingDocArrayHnswSearch
	DestinationLangchainIndexingChromaLocalPersistance *DestinationLangchainIndexingChromaLocalPersistance

	Type DestinationLangchainIndexingType
}

func CreateDestinationLangchainIndexingDestinationLangchainIndexingPinecone(destinationLangchainIndexingPinecone DestinationLangchainIndexingPinecone) DestinationLangchainIndexing {
	typ := DestinationLangchainIndexingTypeDestinationLangchainIndexingPinecone

	return DestinationLangchainIndexing{
		DestinationLangchainIndexingPinecone: &destinationLangchainIndexingPinecone,
		Type:                                 typ,
	}
}

func CreateDestinationLangchainIndexingDestinationLangchainIndexingDocArrayHnswSearch(destinationLangchainIndexingDocArrayHnswSearch DestinationLangchainIndexingDocArrayHnswSearch) DestinationLangchainIndexing {
	typ := DestinationLangchainIndexingTypeDestinationLangchainIndexingDocArrayHnswSearch

	return DestinationLangchainIndexing{
		DestinationLangchainIndexingDocArrayHnswSearch: &destinationLangchainIndexingDocArrayHnswSearch,
		Type: typ,
	}
}

func CreateDestinationLangchainIndexingDestinationLangchainIndexingChromaLocalPersistance(destinationLangchainIndexingChromaLocalPersistance DestinationLangchainIndexingChromaLocalPersistance) DestinationLangchainIndexing {
	typ := DestinationLangchainIndexingTypeDestinationLangchainIndexingChromaLocalPersistance

	return DestinationLangchainIndexing{
		DestinationLangchainIndexingChromaLocalPersistance: &destinationLangchainIndexingChromaLocalPersistance,
		Type: typ,
	}
}

func (u *DestinationLangchainIndexing) UnmarshalJSON(data []byte) error {

	destinationLangchainIndexingDocArrayHnswSearch := new(DestinationLangchainIndexingDocArrayHnswSearch)
	if err := utils.UnmarshalJSON(data, &destinationLangchainIndexingDocArrayHnswSearch, "", true, true); err == nil {
		u.DestinationLangchainIndexingDocArrayHnswSearch = destinationLangchainIndexingDocArrayHnswSearch
		u.Type = DestinationLangchainIndexingTypeDestinationLangchainIndexingDocArrayHnswSearch
		return nil
	}

	destinationLangchainIndexingChromaLocalPersistance := new(DestinationLangchainIndexingChromaLocalPersistance)
	if err := utils.UnmarshalJSON(data, &destinationLangchainIndexingChromaLocalPersistance, "", true, true); err == nil {
		u.DestinationLangchainIndexingChromaLocalPersistance = destinationLangchainIndexingChromaLocalPersistance
		u.Type = DestinationLangchainIndexingTypeDestinationLangchainIndexingChromaLocalPersistance
		return nil
	}

	destinationLangchainIndexingPinecone := new(DestinationLangchainIndexingPinecone)
	if err := utils.UnmarshalJSON(data, &destinationLangchainIndexingPinecone, "", true, true); err == nil {
		u.DestinationLangchainIndexingPinecone = destinationLangchainIndexingPinecone
		u.Type = DestinationLangchainIndexingTypeDestinationLangchainIndexingPinecone
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationLangchainIndexing) MarshalJSON() ([]byte, error) {
	if u.DestinationLangchainIndexingPinecone != nil {
		return utils.MarshalJSON(u.DestinationLangchainIndexingPinecone, "", true)
	}

	if u.DestinationLangchainIndexingDocArrayHnswSearch != nil {
		return utils.MarshalJSON(u.DestinationLangchainIndexingDocArrayHnswSearch, "", true)
	}

	if u.DestinationLangchainIndexingChromaLocalPersistance != nil {
		return utils.MarshalJSON(u.DestinationLangchainIndexingChromaLocalPersistance, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type DestinationLangchainProcessingConfigModel struct {
	// Size of overlap between chunks in tokens to store in vector store to better capture relevant context
	ChunkOverlap *int64 `default:"0" json:"chunk_overlap"`
	// Size of chunks in tokens to store in vector store (make sure it is not too big for the context if your LLM)
	ChunkSize int64 `json:"chunk_size"`
	// List of fields in the record that should be used to calculate the embedding. All other fields are passed along as meta fields. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered text fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. `user.name` will access the `name` field in the `user` object. It's also possible to use wildcards to access all fields in an object, e.g. `users.*.name` will access all `names` fields in all entries of the `users` array.
	TextFields []string `json:"text_fields"`
}

func (d DestinationLangchainProcessingConfigModel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationLangchainProcessingConfigModel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationLangchainProcessingConfigModel) GetChunkOverlap() *int64 {
	if o == nil {
		return nil
	}
	return o.ChunkOverlap
}

func (o *DestinationLangchainProcessingConfigModel) GetChunkSize() int64 {
	if o == nil {
		return 0
	}
	return o.ChunkSize
}

func (o *DestinationLangchainProcessingConfigModel) GetTextFields() []string {
	if o == nil {
		return []string{}
	}
	return o.TextFields
}

type DestinationLangchain struct {
	destinationType DestinationLangchainLangchain `const:"langchain" json:"destinationType"`
	// Embedding configuration
	Embedding DestinationLangchainEmbedding `json:"embedding"`
	// Indexing configuration
	Indexing   DestinationLangchainIndexing              `json:"indexing"`
	Processing DestinationLangchainProcessingConfigModel `json:"processing"`
}

func (d DestinationLangchain) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationLangchain) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationLangchain) GetDestinationType() DestinationLangchainLangchain {
	return DestinationLangchainLangchainLangchain
}

func (o *DestinationLangchain) GetEmbedding() DestinationLangchainEmbedding {
	if o == nil {
		return DestinationLangchainEmbedding{}
	}
	return o.Embedding
}

func (o *DestinationLangchain) GetIndexing() DestinationLangchainIndexing {
	if o == nil {
		return DestinationLangchainIndexing{}
	}
	return o.Indexing
}

func (o *DestinationLangchain) GetProcessing() DestinationLangchainProcessingConfigModel {
	if o == nil {
		return DestinationLangchainProcessingConfigModel{}
	}
	return o.Processing
}
