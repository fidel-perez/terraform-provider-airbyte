// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationMilvusResourceModel) ToCreateSDKType() *shared.DestinationMilvusCreateRequest {
	var embedding shared.DestinationMilvusEmbedding
	var destinationMilvusOpenAI *shared.DestinationMilvusOpenAI
	if r.Configuration.Embedding.OpenAI != nil {
		openaiKey := r.Configuration.Embedding.OpenAI.OpenaiKey.ValueString()
		destinationMilvusOpenAI = &shared.DestinationMilvusOpenAI{
			OpenaiKey: openaiKey,
		}
	}
	if destinationMilvusOpenAI != nil {
		embedding = shared.DestinationMilvusEmbedding{
			DestinationMilvusOpenAI: destinationMilvusOpenAI,
		}
	}
	var destinationMilvusCohere *shared.DestinationMilvusCohere
	if r.Configuration.Embedding.Cohere != nil {
		cohereKey := r.Configuration.Embedding.Cohere.CohereKey.ValueString()
		destinationMilvusCohere = &shared.DestinationMilvusCohere{
			CohereKey: cohereKey,
		}
	}
	if destinationMilvusCohere != nil {
		embedding = shared.DestinationMilvusEmbedding{
			DestinationMilvusCohere: destinationMilvusCohere,
		}
	}
	var destinationMilvusFake *shared.DestinationMilvusFake
	if r.Configuration.Embedding.Fake != nil {
		destinationMilvusFake = &shared.DestinationMilvusFake{}
	}
	if destinationMilvusFake != nil {
		embedding = shared.DestinationMilvusEmbedding{
			DestinationMilvusFake: destinationMilvusFake,
		}
	}
	var destinationMilvusFromField *shared.DestinationMilvusFromField
	if r.Configuration.Embedding.FromField != nil {
		dimensions := r.Configuration.Embedding.FromField.Dimensions.ValueInt64()
		fieldName := r.Configuration.Embedding.FromField.FieldName.ValueString()
		destinationMilvusFromField = &shared.DestinationMilvusFromField{
			Dimensions: dimensions,
			FieldName:  fieldName,
		}
	}
	if destinationMilvusFromField != nil {
		embedding = shared.DestinationMilvusEmbedding{
			DestinationMilvusFromField: destinationMilvusFromField,
		}
	}
	var destinationMilvusAzureOpenAI *shared.DestinationMilvusAzureOpenAI
	if r.Configuration.Embedding.AzureOpenAI != nil {
		apiBase := r.Configuration.Embedding.AzureOpenAI.APIBase.ValueString()
		deployment := r.Configuration.Embedding.AzureOpenAI.Deployment.ValueString()
		openaiKey1 := r.Configuration.Embedding.AzureOpenAI.OpenaiKey.ValueString()
		destinationMilvusAzureOpenAI = &shared.DestinationMilvusAzureOpenAI{
			APIBase:    apiBase,
			Deployment: deployment,
			OpenaiKey:  openaiKey1,
		}
	}
	if destinationMilvusAzureOpenAI != nil {
		embedding = shared.DestinationMilvusEmbedding{
			DestinationMilvusAzureOpenAI: destinationMilvusAzureOpenAI,
		}
	}
	var destinationMilvusOpenAICompatible *shared.DestinationMilvusOpenAICompatible
	if r.Configuration.Embedding.OpenAICompatible != nil {
		apiKey := new(string)
		if !r.Configuration.Embedding.OpenAICompatible.APIKey.IsUnknown() && !r.Configuration.Embedding.OpenAICompatible.APIKey.IsNull() {
			*apiKey = r.Configuration.Embedding.OpenAICompatible.APIKey.ValueString()
		} else {
			apiKey = nil
		}
		baseURL := r.Configuration.Embedding.OpenAICompatible.BaseURL.ValueString()
		dimensions1 := r.Configuration.Embedding.OpenAICompatible.Dimensions.ValueInt64()
		modelName := new(string)
		if !r.Configuration.Embedding.OpenAICompatible.ModelName.IsUnknown() && !r.Configuration.Embedding.OpenAICompatible.ModelName.IsNull() {
			*modelName = r.Configuration.Embedding.OpenAICompatible.ModelName.ValueString()
		} else {
			modelName = nil
		}
		destinationMilvusOpenAICompatible = &shared.DestinationMilvusOpenAICompatible{
			APIKey:     apiKey,
			BaseURL:    baseURL,
			Dimensions: dimensions1,
			ModelName:  modelName,
		}
	}
	if destinationMilvusOpenAICompatible != nil {
		embedding = shared.DestinationMilvusEmbedding{
			DestinationMilvusOpenAICompatible: destinationMilvusOpenAICompatible,
		}
	}
	var auth shared.DestinationMilvusAuthentication
	var destinationMilvusAPIToken *shared.DestinationMilvusAPIToken
	if r.Configuration.Indexing.Auth.APIToken != nil {
		token := r.Configuration.Indexing.Auth.APIToken.Token.ValueString()
		destinationMilvusAPIToken = &shared.DestinationMilvusAPIToken{
			Token: token,
		}
	}
	if destinationMilvusAPIToken != nil {
		auth = shared.DestinationMilvusAuthentication{
			DestinationMilvusAPIToken: destinationMilvusAPIToken,
		}
	}
	var destinationMilvusUsernamePassword *shared.DestinationMilvusUsernamePassword
	if r.Configuration.Indexing.Auth.UsernamePassword != nil {
		password := r.Configuration.Indexing.Auth.UsernamePassword.Password.ValueString()
		username := r.Configuration.Indexing.Auth.UsernamePassword.Username.ValueString()
		destinationMilvusUsernamePassword = &shared.DestinationMilvusUsernamePassword{
			Password: password,
			Username: username,
		}
	}
	if destinationMilvusUsernamePassword != nil {
		auth = shared.DestinationMilvusAuthentication{
			DestinationMilvusUsernamePassword: destinationMilvusUsernamePassword,
		}
	}
	var destinationMilvusNoAuth *shared.DestinationMilvusNoAuth
	if r.Configuration.Indexing.Auth.NoAuth != nil {
		destinationMilvusNoAuth = &shared.DestinationMilvusNoAuth{}
	}
	if destinationMilvusNoAuth != nil {
		auth = shared.DestinationMilvusAuthentication{
			DestinationMilvusNoAuth: destinationMilvusNoAuth,
		}
	}
	collection := r.Configuration.Indexing.Collection.ValueString()
	db := new(string)
	if !r.Configuration.Indexing.Db.IsUnknown() && !r.Configuration.Indexing.Db.IsNull() {
		*db = r.Configuration.Indexing.Db.ValueString()
	} else {
		db = nil
	}
	host := r.Configuration.Indexing.Host.ValueString()
	textField := new(string)
	if !r.Configuration.Indexing.TextField.IsUnknown() && !r.Configuration.Indexing.TextField.IsNull() {
		*textField = r.Configuration.Indexing.TextField.ValueString()
	} else {
		textField = nil
	}
	vectorField := new(string)
	if !r.Configuration.Indexing.VectorField.IsUnknown() && !r.Configuration.Indexing.VectorField.IsNull() {
		*vectorField = r.Configuration.Indexing.VectorField.ValueString()
	} else {
		vectorField = nil
	}
	indexing := shared.DestinationMilvusIndexing{
		Auth:        auth,
		Collection:  collection,
		Db:          db,
		Host:        host,
		TextField:   textField,
		VectorField: vectorField,
	}
	chunkOverlap := new(int64)
	if !r.Configuration.Processing.ChunkOverlap.IsUnknown() && !r.Configuration.Processing.ChunkOverlap.IsNull() {
		*chunkOverlap = r.Configuration.Processing.ChunkOverlap.ValueInt64()
	} else {
		chunkOverlap = nil
	}
	chunkSize := r.Configuration.Processing.ChunkSize.ValueInt64()
	var fieldNameMappings []shared.DestinationMilvusFieldNameMappingConfigModel = nil
	for _, fieldNameMappingsItem := range r.Configuration.Processing.FieldNameMappings {
		fromField := fieldNameMappingsItem.FromField.ValueString()
		toField := fieldNameMappingsItem.ToField.ValueString()
		fieldNameMappings = append(fieldNameMappings, shared.DestinationMilvusFieldNameMappingConfigModel{
			FromField: fromField,
			ToField:   toField,
		})
	}
	var metadataFields []string = nil
	for _, metadataFieldsItem := range r.Configuration.Processing.MetadataFields {
		metadataFields = append(metadataFields, metadataFieldsItem.ValueString())
	}
	var textFields []string = nil
	for _, textFieldsItem := range r.Configuration.Processing.TextFields {
		textFields = append(textFields, textFieldsItem.ValueString())
	}
	var textSplitter *shared.DestinationMilvusTextSplitter
	if r.Configuration.Processing.TextSplitter != nil {
		var destinationMilvusBySeparator *shared.DestinationMilvusBySeparator
		if r.Configuration.Processing.TextSplitter.BySeparator != nil {
			keepSeparator := new(bool)
			if !r.Configuration.Processing.TextSplitter.BySeparator.KeepSeparator.IsUnknown() && !r.Configuration.Processing.TextSplitter.BySeparator.KeepSeparator.IsNull() {
				*keepSeparator = r.Configuration.Processing.TextSplitter.BySeparator.KeepSeparator.ValueBool()
			} else {
				keepSeparator = nil
			}
			var separators []string = nil
			for _, separatorsItem := range r.Configuration.Processing.TextSplitter.BySeparator.Separators {
				separators = append(separators, separatorsItem.ValueString())
			}
			destinationMilvusBySeparator = &shared.DestinationMilvusBySeparator{
				KeepSeparator: keepSeparator,
				Separators:    separators,
			}
		}
		if destinationMilvusBySeparator != nil {
			textSplitter = &shared.DestinationMilvusTextSplitter{
				DestinationMilvusBySeparator: destinationMilvusBySeparator,
			}
		}
		var destinationMilvusByMarkdownHeader *shared.DestinationMilvusByMarkdownHeader
		if r.Configuration.Processing.TextSplitter.ByMarkdownHeader != nil {
			splitLevel := new(int64)
			if !r.Configuration.Processing.TextSplitter.ByMarkdownHeader.SplitLevel.IsUnknown() && !r.Configuration.Processing.TextSplitter.ByMarkdownHeader.SplitLevel.IsNull() {
				*splitLevel = r.Configuration.Processing.TextSplitter.ByMarkdownHeader.SplitLevel.ValueInt64()
			} else {
				splitLevel = nil
			}
			destinationMilvusByMarkdownHeader = &shared.DestinationMilvusByMarkdownHeader{
				SplitLevel: splitLevel,
			}
		}
		if destinationMilvusByMarkdownHeader != nil {
			textSplitter = &shared.DestinationMilvusTextSplitter{
				DestinationMilvusByMarkdownHeader: destinationMilvusByMarkdownHeader,
			}
		}
		var destinationMilvusByProgrammingLanguage *shared.DestinationMilvusByProgrammingLanguage
		if r.Configuration.Processing.TextSplitter.ByProgrammingLanguage != nil {
			language := shared.DestinationMilvusLanguage(r.Configuration.Processing.TextSplitter.ByProgrammingLanguage.Language.ValueString())
			destinationMilvusByProgrammingLanguage = &shared.DestinationMilvusByProgrammingLanguage{
				Language: language,
			}
		}
		if destinationMilvusByProgrammingLanguage != nil {
			textSplitter = &shared.DestinationMilvusTextSplitter{
				DestinationMilvusByProgrammingLanguage: destinationMilvusByProgrammingLanguage,
			}
		}
	}
	processing := shared.DestinationMilvusProcessingConfigModel{
		ChunkOverlap:      chunkOverlap,
		ChunkSize:         chunkSize,
		FieldNameMappings: fieldNameMappings,
		MetadataFields:    metadataFields,
		TextFields:        textFields,
		TextSplitter:      textSplitter,
	}
	configuration := shared.DestinationMilvus{
		Embedding:  embedding,
		Indexing:   indexing,
		Processing: processing,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationMilvusCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationMilvusResourceModel) ToGetSDKType() *shared.DestinationMilvusCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationMilvusResourceModel) ToUpdateSDKType() *shared.DestinationMilvusPutRequest {
	var embedding shared.DestinationMilvusUpdateEmbedding
	var destinationMilvusUpdateOpenAI *shared.DestinationMilvusUpdateOpenAI
	if r.Configuration.Embedding.OpenAI != nil {
		openaiKey := r.Configuration.Embedding.OpenAI.OpenaiKey.ValueString()
		destinationMilvusUpdateOpenAI = &shared.DestinationMilvusUpdateOpenAI{
			OpenaiKey: openaiKey,
		}
	}
	if destinationMilvusUpdateOpenAI != nil {
		embedding = shared.DestinationMilvusUpdateEmbedding{
			DestinationMilvusUpdateOpenAI: destinationMilvusUpdateOpenAI,
		}
	}
	var cohere *shared.Cohere
	if r.Configuration.Embedding.Cohere != nil {
		cohereKey := r.Configuration.Embedding.Cohere.CohereKey.ValueString()
		cohere = &shared.Cohere{
			CohereKey: cohereKey,
		}
	}
	if cohere != nil {
		embedding = shared.DestinationMilvusUpdateEmbedding{
			Cohere: cohere,
		}
	}
	var destinationMilvusUpdateFake *shared.DestinationMilvusUpdateFake
	if r.Configuration.Embedding.Fake != nil {
		destinationMilvusUpdateFake = &shared.DestinationMilvusUpdateFake{}
	}
	if destinationMilvusUpdateFake != nil {
		embedding = shared.DestinationMilvusUpdateEmbedding{
			DestinationMilvusUpdateFake: destinationMilvusUpdateFake,
		}
	}
	var fromField *shared.FromField
	if r.Configuration.Embedding.FromField != nil {
		dimensions := r.Configuration.Embedding.FromField.Dimensions.ValueInt64()
		fieldName := r.Configuration.Embedding.FromField.FieldName.ValueString()
		fromField = &shared.FromField{
			Dimensions: dimensions,
			FieldName:  fieldName,
		}
	}
	if fromField != nil {
		embedding = shared.DestinationMilvusUpdateEmbedding{
			FromField: fromField,
		}
	}
	var azureOpenAI *shared.AzureOpenAI
	if r.Configuration.Embedding.AzureOpenAI != nil {
		apiBase := r.Configuration.Embedding.AzureOpenAI.APIBase.ValueString()
		deployment := r.Configuration.Embedding.AzureOpenAI.Deployment.ValueString()
		openaiKey1 := r.Configuration.Embedding.AzureOpenAI.OpenaiKey.ValueString()
		azureOpenAI = &shared.AzureOpenAI{
			APIBase:    apiBase,
			Deployment: deployment,
			OpenaiKey:  openaiKey1,
		}
	}
	if azureOpenAI != nil {
		embedding = shared.DestinationMilvusUpdateEmbedding{
			AzureOpenAI: azureOpenAI,
		}
	}
	var openAICompatible *shared.OpenAICompatible
	if r.Configuration.Embedding.OpenAICompatible != nil {
		apiKey := new(string)
		if !r.Configuration.Embedding.OpenAICompatible.APIKey.IsUnknown() && !r.Configuration.Embedding.OpenAICompatible.APIKey.IsNull() {
			*apiKey = r.Configuration.Embedding.OpenAICompatible.APIKey.ValueString()
		} else {
			apiKey = nil
		}
		baseURL := r.Configuration.Embedding.OpenAICompatible.BaseURL.ValueString()
		dimensions1 := r.Configuration.Embedding.OpenAICompatible.Dimensions.ValueInt64()
		modelName := new(string)
		if !r.Configuration.Embedding.OpenAICompatible.ModelName.IsUnknown() && !r.Configuration.Embedding.OpenAICompatible.ModelName.IsNull() {
			*modelName = r.Configuration.Embedding.OpenAICompatible.ModelName.ValueString()
		} else {
			modelName = nil
		}
		openAICompatible = &shared.OpenAICompatible{
			APIKey:     apiKey,
			BaseURL:    baseURL,
			Dimensions: dimensions1,
			ModelName:  modelName,
		}
	}
	if openAICompatible != nil {
		embedding = shared.DestinationMilvusUpdateEmbedding{
			OpenAICompatible: openAICompatible,
		}
	}
	var auth shared.DestinationMilvusUpdateAuthentication
	var destinationMilvusUpdateAPIToken *shared.DestinationMilvusUpdateAPIToken
	if r.Configuration.Indexing.Auth.APIToken != nil {
		token := r.Configuration.Indexing.Auth.APIToken.Token.ValueString()
		destinationMilvusUpdateAPIToken = &shared.DestinationMilvusUpdateAPIToken{
			Token: token,
		}
	}
	if destinationMilvusUpdateAPIToken != nil {
		auth = shared.DestinationMilvusUpdateAuthentication{
			DestinationMilvusUpdateAPIToken: destinationMilvusUpdateAPIToken,
		}
	}
	var destinationMilvusUpdateUsernamePassword *shared.DestinationMilvusUpdateUsernamePassword
	if r.Configuration.Indexing.Auth.UsernamePassword != nil {
		password := r.Configuration.Indexing.Auth.UsernamePassword.Password.ValueString()
		username := r.Configuration.Indexing.Auth.UsernamePassword.Username.ValueString()
		destinationMilvusUpdateUsernamePassword = &shared.DestinationMilvusUpdateUsernamePassword{
			Password: password,
			Username: username,
		}
	}
	if destinationMilvusUpdateUsernamePassword != nil {
		auth = shared.DestinationMilvusUpdateAuthentication{
			DestinationMilvusUpdateUsernamePassword: destinationMilvusUpdateUsernamePassword,
		}
	}
	var noAuth *shared.NoAuth
	if r.Configuration.Indexing.Auth.NoAuth != nil {
		noAuth = &shared.NoAuth{}
	}
	if noAuth != nil {
		auth = shared.DestinationMilvusUpdateAuthentication{
			NoAuth: noAuth,
		}
	}
	collection := r.Configuration.Indexing.Collection.ValueString()
	db := new(string)
	if !r.Configuration.Indexing.Db.IsUnknown() && !r.Configuration.Indexing.Db.IsNull() {
		*db = r.Configuration.Indexing.Db.ValueString()
	} else {
		db = nil
	}
	host := r.Configuration.Indexing.Host.ValueString()
	textField := new(string)
	if !r.Configuration.Indexing.TextField.IsUnknown() && !r.Configuration.Indexing.TextField.IsNull() {
		*textField = r.Configuration.Indexing.TextField.ValueString()
	} else {
		textField = nil
	}
	vectorField := new(string)
	if !r.Configuration.Indexing.VectorField.IsUnknown() && !r.Configuration.Indexing.VectorField.IsNull() {
		*vectorField = r.Configuration.Indexing.VectorField.ValueString()
	} else {
		vectorField = nil
	}
	indexing := shared.DestinationMilvusUpdateIndexing{
		Auth:        auth,
		Collection:  collection,
		Db:          db,
		Host:        host,
		TextField:   textField,
		VectorField: vectorField,
	}
	chunkOverlap := new(int64)
	if !r.Configuration.Processing.ChunkOverlap.IsUnknown() && !r.Configuration.Processing.ChunkOverlap.IsNull() {
		*chunkOverlap = r.Configuration.Processing.ChunkOverlap.ValueInt64()
	} else {
		chunkOverlap = nil
	}
	chunkSize := r.Configuration.Processing.ChunkSize.ValueInt64()
	var fieldNameMappings []shared.FieldNameMappingConfigModel = nil
	for _, fieldNameMappingsItem := range r.Configuration.Processing.FieldNameMappings {
		fromField1 := fieldNameMappingsItem.FromField.ValueString()
		toField := fieldNameMappingsItem.ToField.ValueString()
		fieldNameMappings = append(fieldNameMappings, shared.FieldNameMappingConfigModel{
			FromField: fromField1,
			ToField:   toField,
		})
	}
	var metadataFields []string = nil
	for _, metadataFieldsItem := range r.Configuration.Processing.MetadataFields {
		metadataFields = append(metadataFields, metadataFieldsItem.ValueString())
	}
	var textFields []string = nil
	for _, textFieldsItem := range r.Configuration.Processing.TextFields {
		textFields = append(textFields, textFieldsItem.ValueString())
	}
	var textSplitter *shared.TextSplitter
	if r.Configuration.Processing.TextSplitter != nil {
		var bySeparator *shared.BySeparator
		if r.Configuration.Processing.TextSplitter.BySeparator != nil {
			keepSeparator := new(bool)
			if !r.Configuration.Processing.TextSplitter.BySeparator.KeepSeparator.IsUnknown() && !r.Configuration.Processing.TextSplitter.BySeparator.KeepSeparator.IsNull() {
				*keepSeparator = r.Configuration.Processing.TextSplitter.BySeparator.KeepSeparator.ValueBool()
			} else {
				keepSeparator = nil
			}
			var separators []string = nil
			for _, separatorsItem := range r.Configuration.Processing.TextSplitter.BySeparator.Separators {
				separators = append(separators, separatorsItem.ValueString())
			}
			bySeparator = &shared.BySeparator{
				KeepSeparator: keepSeparator,
				Separators:    separators,
			}
		}
		if bySeparator != nil {
			textSplitter = &shared.TextSplitter{
				BySeparator: bySeparator,
			}
		}
		var byMarkdownHeader *shared.ByMarkdownHeader
		if r.Configuration.Processing.TextSplitter.ByMarkdownHeader != nil {
			splitLevel := new(int64)
			if !r.Configuration.Processing.TextSplitter.ByMarkdownHeader.SplitLevel.IsUnknown() && !r.Configuration.Processing.TextSplitter.ByMarkdownHeader.SplitLevel.IsNull() {
				*splitLevel = r.Configuration.Processing.TextSplitter.ByMarkdownHeader.SplitLevel.ValueInt64()
			} else {
				splitLevel = nil
			}
			byMarkdownHeader = &shared.ByMarkdownHeader{
				SplitLevel: splitLevel,
			}
		}
		if byMarkdownHeader != nil {
			textSplitter = &shared.TextSplitter{
				ByMarkdownHeader: byMarkdownHeader,
			}
		}
		var byProgrammingLanguage *shared.ByProgrammingLanguage
		if r.Configuration.Processing.TextSplitter.ByProgrammingLanguage != nil {
			language := shared.DestinationMilvusUpdateLanguage(r.Configuration.Processing.TextSplitter.ByProgrammingLanguage.Language.ValueString())
			byProgrammingLanguage = &shared.ByProgrammingLanguage{
				Language: language,
			}
		}
		if byProgrammingLanguage != nil {
			textSplitter = &shared.TextSplitter{
				ByProgrammingLanguage: byProgrammingLanguage,
			}
		}
	}
	processing := shared.DestinationMilvusUpdateProcessingConfigModel{
		ChunkOverlap:      chunkOverlap,
		ChunkSize:         chunkSize,
		FieldNameMappings: fieldNameMappings,
		MetadataFields:    metadataFields,
		TextFields:        textFields,
		TextSplitter:      textSplitter,
	}
	configuration := shared.DestinationMilvusUpdate{
		Embedding:  embedding,
		Indexing:   indexing,
		Processing: processing,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationMilvusPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationMilvusResourceModel) ToDeleteSDKType() *shared.DestinationMilvusCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationMilvusResourceModel) RefreshFromGetResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationMilvusResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.RefreshFromGetResponse(resp)
}
