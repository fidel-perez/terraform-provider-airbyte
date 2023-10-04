// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceFileSecureResourceModel) ToCreateSDKType() *shared.SourceFileSecureCreateRequest {
	datasetName := r.Configuration.DatasetName.ValueString()
	format := new(shared.SourceFileSecureFileFormat)
	if !r.Configuration.Format.IsUnknown() && !r.Configuration.Format.IsNull() {
		*format = shared.SourceFileSecureFileFormat(r.Configuration.Format.ValueString())
	} else {
		format = nil
	}
	var provider shared.SourceFileSecureStorageProvider
	var sourceFileSecureStorageProviderHTTPSPublicWeb *shared.SourceFileSecureStorageProviderHTTPSPublicWeb
	if r.Configuration.Provider.SourceFileSecureStorageProviderHTTPSPublicWeb != nil {
		userAgent := new(bool)
		if !r.Configuration.Provider.SourceFileSecureStorageProviderHTTPSPublicWeb.UserAgent.IsUnknown() && !r.Configuration.Provider.SourceFileSecureStorageProviderHTTPSPublicWeb.UserAgent.IsNull() {
			*userAgent = r.Configuration.Provider.SourceFileSecureStorageProviderHTTPSPublicWeb.UserAgent.ValueBool()
		} else {
			userAgent = nil
		}
		sourceFileSecureStorageProviderHTTPSPublicWeb = &shared.SourceFileSecureStorageProviderHTTPSPublicWeb{
			UserAgent: userAgent,
		}
	}
	if sourceFileSecureStorageProviderHTTPSPublicWeb != nil {
		provider = shared.SourceFileSecureStorageProvider{
			SourceFileSecureStorageProviderHTTPSPublicWeb: sourceFileSecureStorageProviderHTTPSPublicWeb,
		}
	}
	var sourceFileSecureStorageProviderGCSGoogleCloudStorage *shared.SourceFileSecureStorageProviderGCSGoogleCloudStorage
	if r.Configuration.Provider.SourceFileSecureStorageProviderGCSGoogleCloudStorage != nil {
		serviceAccountJSON := new(string)
		if !r.Configuration.Provider.SourceFileSecureStorageProviderGCSGoogleCloudStorage.ServiceAccountJSON.IsUnknown() && !r.Configuration.Provider.SourceFileSecureStorageProviderGCSGoogleCloudStorage.ServiceAccountJSON.IsNull() {
			*serviceAccountJSON = r.Configuration.Provider.SourceFileSecureStorageProviderGCSGoogleCloudStorage.ServiceAccountJSON.ValueString()
		} else {
			serviceAccountJSON = nil
		}
		sourceFileSecureStorageProviderGCSGoogleCloudStorage = &shared.SourceFileSecureStorageProviderGCSGoogleCloudStorage{
			ServiceAccountJSON: serviceAccountJSON,
		}
	}
	if sourceFileSecureStorageProviderGCSGoogleCloudStorage != nil {
		provider = shared.SourceFileSecureStorageProvider{
			SourceFileSecureStorageProviderGCSGoogleCloudStorage: sourceFileSecureStorageProviderGCSGoogleCloudStorage,
		}
	}
	var sourceFileSecureStorageProviderS3AmazonWebServices *shared.SourceFileSecureStorageProviderS3AmazonWebServices
	if r.Configuration.Provider.SourceFileSecureStorageProviderS3AmazonWebServices != nil {
		awsAccessKeyID := new(string)
		if !r.Configuration.Provider.SourceFileSecureStorageProviderS3AmazonWebServices.AwsAccessKeyID.IsUnknown() && !r.Configuration.Provider.SourceFileSecureStorageProviderS3AmazonWebServices.AwsAccessKeyID.IsNull() {
			*awsAccessKeyID = r.Configuration.Provider.SourceFileSecureStorageProviderS3AmazonWebServices.AwsAccessKeyID.ValueString()
		} else {
			awsAccessKeyID = nil
		}
		awsSecretAccessKey := new(string)
		if !r.Configuration.Provider.SourceFileSecureStorageProviderS3AmazonWebServices.AwsSecretAccessKey.IsUnknown() && !r.Configuration.Provider.SourceFileSecureStorageProviderS3AmazonWebServices.AwsSecretAccessKey.IsNull() {
			*awsSecretAccessKey = r.Configuration.Provider.SourceFileSecureStorageProviderS3AmazonWebServices.AwsSecretAccessKey.ValueString()
		} else {
			awsSecretAccessKey = nil
		}
		sourceFileSecureStorageProviderS3AmazonWebServices = &shared.SourceFileSecureStorageProviderS3AmazonWebServices{
			AwsAccessKeyID:     awsAccessKeyID,
			AwsSecretAccessKey: awsSecretAccessKey,
		}
	}
	if sourceFileSecureStorageProviderS3AmazonWebServices != nil {
		provider = shared.SourceFileSecureStorageProvider{
			SourceFileSecureStorageProviderS3AmazonWebServices: sourceFileSecureStorageProviderS3AmazonWebServices,
		}
	}
	var sourceFileSecureStorageProviderAzBlobAzureBlobStorage *shared.SourceFileSecureStorageProviderAzBlobAzureBlobStorage
	if r.Configuration.Provider.SourceFileSecureStorageProviderAzBlobAzureBlobStorage != nil {
		sasToken := new(string)
		if !r.Configuration.Provider.SourceFileSecureStorageProviderAzBlobAzureBlobStorage.SasToken.IsUnknown() && !r.Configuration.Provider.SourceFileSecureStorageProviderAzBlobAzureBlobStorage.SasToken.IsNull() {
			*sasToken = r.Configuration.Provider.SourceFileSecureStorageProviderAzBlobAzureBlobStorage.SasToken.ValueString()
		} else {
			sasToken = nil
		}
		sharedKey := new(string)
		if !r.Configuration.Provider.SourceFileSecureStorageProviderAzBlobAzureBlobStorage.SharedKey.IsUnknown() && !r.Configuration.Provider.SourceFileSecureStorageProviderAzBlobAzureBlobStorage.SharedKey.IsNull() {
			*sharedKey = r.Configuration.Provider.SourceFileSecureStorageProviderAzBlobAzureBlobStorage.SharedKey.ValueString()
		} else {
			sharedKey = nil
		}
		storageAccount := r.Configuration.Provider.SourceFileSecureStorageProviderAzBlobAzureBlobStorage.StorageAccount.ValueString()
		sourceFileSecureStorageProviderAzBlobAzureBlobStorage = &shared.SourceFileSecureStorageProviderAzBlobAzureBlobStorage{
			SasToken:       sasToken,
			SharedKey:      sharedKey,
			StorageAccount: storageAccount,
		}
	}
	if sourceFileSecureStorageProviderAzBlobAzureBlobStorage != nil {
		provider = shared.SourceFileSecureStorageProvider{
			SourceFileSecureStorageProviderAzBlobAzureBlobStorage: sourceFileSecureStorageProviderAzBlobAzureBlobStorage,
		}
	}
	var sourceFileSecureStorageProviderSSHSecureShell *shared.SourceFileSecureStorageProviderSSHSecureShell
	if r.Configuration.Provider.SourceFileSecureStorageProviderSSHSecureShell != nil {
		host := r.Configuration.Provider.SourceFileSecureStorageProviderSSHSecureShell.Host.ValueString()
		password := new(string)
		if !r.Configuration.Provider.SourceFileSecureStorageProviderSSHSecureShell.Password.IsUnknown() && !r.Configuration.Provider.SourceFileSecureStorageProviderSSHSecureShell.Password.IsNull() {
			*password = r.Configuration.Provider.SourceFileSecureStorageProviderSSHSecureShell.Password.ValueString()
		} else {
			password = nil
		}
		port := new(string)
		if !r.Configuration.Provider.SourceFileSecureStorageProviderSSHSecureShell.Port.IsUnknown() && !r.Configuration.Provider.SourceFileSecureStorageProviderSSHSecureShell.Port.IsNull() {
			*port = r.Configuration.Provider.SourceFileSecureStorageProviderSSHSecureShell.Port.ValueString()
		} else {
			port = nil
		}
		user := r.Configuration.Provider.SourceFileSecureStorageProviderSSHSecureShell.User.ValueString()
		sourceFileSecureStorageProviderSSHSecureShell = &shared.SourceFileSecureStorageProviderSSHSecureShell{
			Host:     host,
			Password: password,
			Port:     port,
			User:     user,
		}
	}
	if sourceFileSecureStorageProviderSSHSecureShell != nil {
		provider = shared.SourceFileSecureStorageProvider{
			SourceFileSecureStorageProviderSSHSecureShell: sourceFileSecureStorageProviderSSHSecureShell,
		}
	}
	var sourceFileSecureStorageProviderSCPSecureCopyProtocol *shared.SourceFileSecureStorageProviderSCPSecureCopyProtocol
	if r.Configuration.Provider.SourceFileSecureStorageProviderSCPSecureCopyProtocol != nil {
		host1 := r.Configuration.Provider.SourceFileSecureStorageProviderSCPSecureCopyProtocol.Host.ValueString()
		password1 := new(string)
		if !r.Configuration.Provider.SourceFileSecureStorageProviderSCPSecureCopyProtocol.Password.IsUnknown() && !r.Configuration.Provider.SourceFileSecureStorageProviderSCPSecureCopyProtocol.Password.IsNull() {
			*password1 = r.Configuration.Provider.SourceFileSecureStorageProviderSCPSecureCopyProtocol.Password.ValueString()
		} else {
			password1 = nil
		}
		port1 := new(string)
		if !r.Configuration.Provider.SourceFileSecureStorageProviderSCPSecureCopyProtocol.Port.IsUnknown() && !r.Configuration.Provider.SourceFileSecureStorageProviderSCPSecureCopyProtocol.Port.IsNull() {
			*port1 = r.Configuration.Provider.SourceFileSecureStorageProviderSCPSecureCopyProtocol.Port.ValueString()
		} else {
			port1 = nil
		}
		user1 := r.Configuration.Provider.SourceFileSecureStorageProviderSCPSecureCopyProtocol.User.ValueString()
		sourceFileSecureStorageProviderSCPSecureCopyProtocol = &shared.SourceFileSecureStorageProviderSCPSecureCopyProtocol{
			Host:     host1,
			Password: password1,
			Port:     port1,
			User:     user1,
		}
	}
	if sourceFileSecureStorageProviderSCPSecureCopyProtocol != nil {
		provider = shared.SourceFileSecureStorageProvider{
			SourceFileSecureStorageProviderSCPSecureCopyProtocol: sourceFileSecureStorageProviderSCPSecureCopyProtocol,
		}
	}
	var sourceFileSecureStorageProviderSFTPSecureFileTransferProtocol *shared.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol
	if r.Configuration.Provider.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol != nil {
		host2 := r.Configuration.Provider.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol.Host.ValueString()
		password2 := new(string)
		if !r.Configuration.Provider.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol.Password.IsUnknown() && !r.Configuration.Provider.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol.Password.IsNull() {
			*password2 = r.Configuration.Provider.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol.Password.ValueString()
		} else {
			password2 = nil
		}
		port2 := new(string)
		if !r.Configuration.Provider.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol.Port.IsUnknown() && !r.Configuration.Provider.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol.Port.IsNull() {
			*port2 = r.Configuration.Provider.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol.Port.ValueString()
		} else {
			port2 = nil
		}
		user2 := r.Configuration.Provider.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol.User.ValueString()
		sourceFileSecureStorageProviderSFTPSecureFileTransferProtocol = &shared.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol{
			Host:     host2,
			Password: password2,
			Port:     port2,
			User:     user2,
		}
	}
	if sourceFileSecureStorageProviderSFTPSecureFileTransferProtocol != nil {
		provider = shared.SourceFileSecureStorageProvider{
			SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol: sourceFileSecureStorageProviderSFTPSecureFileTransferProtocol,
		}
	}
	readerOptions := new(string)
	if !r.Configuration.ReaderOptions.IsUnknown() && !r.Configuration.ReaderOptions.IsNull() {
		*readerOptions = r.Configuration.ReaderOptions.ValueString()
	} else {
		readerOptions = nil
	}
	url := r.Configuration.URL.ValueString()
	configuration := shared.SourceFileSecure{
		DatasetName:   datasetName,
		Format:        format,
		Provider:      provider,
		ReaderOptions: readerOptions,
		URL:           url,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceFileSecureCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceFileSecureResourceModel) ToGetSDKType() *shared.SourceFileSecureCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceFileSecureResourceModel) ToUpdateSDKType() *shared.SourceFileSecurePutRequest {
	datasetName := r.Configuration.DatasetName.ValueString()
	format := new(shared.SourceFileSecureUpdateFileFormat)
	if !r.Configuration.Format.IsUnknown() && !r.Configuration.Format.IsNull() {
		*format = shared.SourceFileSecureUpdateFileFormat(r.Configuration.Format.ValueString())
	} else {
		format = nil
	}
	var provider shared.SourceFileSecureUpdateStorageProvider
	var sourceFileSecureUpdateStorageProviderHTTPSPublicWeb *shared.SourceFileSecureUpdateStorageProviderHTTPSPublicWeb
	if r.Configuration.Provider.SourceFileSecureUpdateStorageProviderHTTPSPublicWeb != nil {
		userAgent := new(bool)
		if !r.Configuration.Provider.SourceFileSecureUpdateStorageProviderHTTPSPublicWeb.UserAgent.IsUnknown() && !r.Configuration.Provider.SourceFileSecureUpdateStorageProviderHTTPSPublicWeb.UserAgent.IsNull() {
			*userAgent = r.Configuration.Provider.SourceFileSecureUpdateStorageProviderHTTPSPublicWeb.UserAgent.ValueBool()
		} else {
			userAgent = nil
		}
		sourceFileSecureUpdateStorageProviderHTTPSPublicWeb = &shared.SourceFileSecureUpdateStorageProviderHTTPSPublicWeb{
			UserAgent: userAgent,
		}
	}
	if sourceFileSecureUpdateStorageProviderHTTPSPublicWeb != nil {
		provider = shared.SourceFileSecureUpdateStorageProvider{
			SourceFileSecureUpdateStorageProviderHTTPSPublicWeb: sourceFileSecureUpdateStorageProviderHTTPSPublicWeb,
		}
	}
	var sourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage *shared.SourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage
	if r.Configuration.Provider.SourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage != nil {
		serviceAccountJSON := new(string)
		if !r.Configuration.Provider.SourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage.ServiceAccountJSON.IsUnknown() && !r.Configuration.Provider.SourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage.ServiceAccountJSON.IsNull() {
			*serviceAccountJSON = r.Configuration.Provider.SourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage.ServiceAccountJSON.ValueString()
		} else {
			serviceAccountJSON = nil
		}
		sourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage = &shared.SourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage{
			ServiceAccountJSON: serviceAccountJSON,
		}
	}
	if sourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage != nil {
		provider = shared.SourceFileSecureUpdateStorageProvider{
			SourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage: sourceFileSecureUpdateStorageProviderGCSGoogleCloudStorage,
		}
	}
	var sourceFileSecureUpdateStorageProviderS3AmazonWebServices *shared.SourceFileSecureUpdateStorageProviderS3AmazonWebServices
	if r.Configuration.Provider.SourceFileSecureUpdateStorageProviderS3AmazonWebServices != nil {
		awsAccessKeyID := new(string)
		if !r.Configuration.Provider.SourceFileSecureUpdateStorageProviderS3AmazonWebServices.AwsAccessKeyID.IsUnknown() && !r.Configuration.Provider.SourceFileSecureUpdateStorageProviderS3AmazonWebServices.AwsAccessKeyID.IsNull() {
			*awsAccessKeyID = r.Configuration.Provider.SourceFileSecureUpdateStorageProviderS3AmazonWebServices.AwsAccessKeyID.ValueString()
		} else {
			awsAccessKeyID = nil
		}
		awsSecretAccessKey := new(string)
		if !r.Configuration.Provider.SourceFileSecureUpdateStorageProviderS3AmazonWebServices.AwsSecretAccessKey.IsUnknown() && !r.Configuration.Provider.SourceFileSecureUpdateStorageProviderS3AmazonWebServices.AwsSecretAccessKey.IsNull() {
			*awsSecretAccessKey = r.Configuration.Provider.SourceFileSecureUpdateStorageProviderS3AmazonWebServices.AwsSecretAccessKey.ValueString()
		} else {
			awsSecretAccessKey = nil
		}
		sourceFileSecureUpdateStorageProviderS3AmazonWebServices = &shared.SourceFileSecureUpdateStorageProviderS3AmazonWebServices{
			AwsAccessKeyID:     awsAccessKeyID,
			AwsSecretAccessKey: awsSecretAccessKey,
		}
	}
	if sourceFileSecureUpdateStorageProviderS3AmazonWebServices != nil {
		provider = shared.SourceFileSecureUpdateStorageProvider{
			SourceFileSecureUpdateStorageProviderS3AmazonWebServices: sourceFileSecureUpdateStorageProviderS3AmazonWebServices,
		}
	}
	var sourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage *shared.SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage
	if r.Configuration.Provider.SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage != nil {
		sasToken := new(string)
		if !r.Configuration.Provider.SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage.SasToken.IsUnknown() && !r.Configuration.Provider.SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage.SasToken.IsNull() {
			*sasToken = r.Configuration.Provider.SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage.SasToken.ValueString()
		} else {
			sasToken = nil
		}
		sharedKey := new(string)
		if !r.Configuration.Provider.SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage.SharedKey.IsUnknown() && !r.Configuration.Provider.SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage.SharedKey.IsNull() {
			*sharedKey = r.Configuration.Provider.SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage.SharedKey.ValueString()
		} else {
			sharedKey = nil
		}
		storageAccount := r.Configuration.Provider.SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage.StorageAccount.ValueString()
		sourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage = &shared.SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage{
			SasToken:       sasToken,
			SharedKey:      sharedKey,
			StorageAccount: storageAccount,
		}
	}
	if sourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage != nil {
		provider = shared.SourceFileSecureUpdateStorageProvider{
			SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage: sourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage,
		}
	}
	var sourceFileSecureUpdateStorageProviderSSHSecureShell *shared.SourceFileSecureUpdateStorageProviderSSHSecureShell
	if r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSSHSecureShell != nil {
		host := r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSSHSecureShell.Host.ValueString()
		password := new(string)
		if !r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSSHSecureShell.Password.IsUnknown() && !r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSSHSecureShell.Password.IsNull() {
			*password = r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSSHSecureShell.Password.ValueString()
		} else {
			password = nil
		}
		port := new(string)
		if !r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSSHSecureShell.Port.IsUnknown() && !r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSSHSecureShell.Port.IsNull() {
			*port = r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSSHSecureShell.Port.ValueString()
		} else {
			port = nil
		}
		user := r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSSHSecureShell.User.ValueString()
		sourceFileSecureUpdateStorageProviderSSHSecureShell = &shared.SourceFileSecureUpdateStorageProviderSSHSecureShell{
			Host:     host,
			Password: password,
			Port:     port,
			User:     user,
		}
	}
	if sourceFileSecureUpdateStorageProviderSSHSecureShell != nil {
		provider = shared.SourceFileSecureUpdateStorageProvider{
			SourceFileSecureUpdateStorageProviderSSHSecureShell: sourceFileSecureUpdateStorageProviderSSHSecureShell,
		}
	}
	var sourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol *shared.SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol
	if r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol != nil {
		host1 := r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol.Host.ValueString()
		password1 := new(string)
		if !r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol.Password.IsUnknown() && !r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol.Password.IsNull() {
			*password1 = r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol.Password.ValueString()
		} else {
			password1 = nil
		}
		port1 := new(string)
		if !r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol.Port.IsUnknown() && !r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol.Port.IsNull() {
			*port1 = r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol.Port.ValueString()
		} else {
			port1 = nil
		}
		user1 := r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol.User.ValueString()
		sourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol = &shared.SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol{
			Host:     host1,
			Password: password1,
			Port:     port1,
			User:     user1,
		}
	}
	if sourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol != nil {
		provider = shared.SourceFileSecureUpdateStorageProvider{
			SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol: sourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol,
		}
	}
	var sourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol *shared.SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol
	if r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol != nil {
		host2 := r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol.Host.ValueString()
		password2 := new(string)
		if !r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol.Password.IsUnknown() && !r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol.Password.IsNull() {
			*password2 = r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol.Password.ValueString()
		} else {
			password2 = nil
		}
		port2 := new(string)
		if !r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol.Port.IsUnknown() && !r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol.Port.IsNull() {
			*port2 = r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol.Port.ValueString()
		} else {
			port2 = nil
		}
		user2 := r.Configuration.Provider.SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol.User.ValueString()
		sourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol = &shared.SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol{
			Host:     host2,
			Password: password2,
			Port:     port2,
			User:     user2,
		}
	}
	if sourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol != nil {
		provider = shared.SourceFileSecureUpdateStorageProvider{
			SourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol: sourceFileSecureUpdateStorageProviderSFTPSecureFileTransferProtocol,
		}
	}
	readerOptions := new(string)
	if !r.Configuration.ReaderOptions.IsUnknown() && !r.Configuration.ReaderOptions.IsNull() {
		*readerOptions = r.Configuration.ReaderOptions.ValueString()
	} else {
		readerOptions = nil
	}
	url := r.Configuration.URL.ValueString()
	configuration := shared.SourceFileSecureUpdate{
		DatasetName:   datasetName,
		Format:        format,
		Provider:      provider,
		ReaderOptions: readerOptions,
		URL:           url,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceFileSecurePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceFileSecureResourceModel) ToDeleteSDKType() *shared.SourceFileSecureCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceFileSecureResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceFileSecureResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
