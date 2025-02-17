// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &AirbyteProvider{}

type AirbyteProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// AirbyteProviderModel describes the provider data model.
type AirbyteProviderModel struct {
	ServerURL  types.String `tfsdk:"server_url"`
	Password   types.String `tfsdk:"password"`
	Username   types.String `tfsdk:"username"`
	BearerAuth types.String `tfsdk:"bearer_auth"`
}

func (p *AirbyteProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "airbyte"
	resp.Version = p.version
}

func (p *AirbyteProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: `airbyte-api: Programatically control Airbyte Cloud, OSS & Enterprise.`,
		Attributes: map[string]schema.Attribute{
			"server_url": schema.StringAttribute{
				MarkdownDescription: "Server URL (defaults to https://api.airbyte.com/v1)",
				Optional:            true,
				Required:            false,
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"username": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"bearer_auth": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *AirbyteProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data AirbyteProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ServerURL := data.ServerURL.ValueString()

	if ServerURL == "" {
		ServerURL = "https://api.airbyte.com/v1"
	}

	var basicAuth *shared.SchemeBasicAuth
	password := data.Password.ValueString()
	username := data.Username.ValueString()
	basicAuth = &shared.SchemeBasicAuth{
		Password: password,
		Username: username,
	}
	bearerAuth := new(string)
	if !data.BearerAuth.IsUnknown() && !data.BearerAuth.IsNull() {
		*bearerAuth = data.BearerAuth.ValueString()
	} else {
		bearerAuth = nil
	}
	security := shared.Security{
		BasicAuth:  basicAuth,
		BearerAuth: bearerAuth,
	}

	opts := []sdk.SDKOption{
		sdk.WithServerURL(ServerURL),
		sdk.WithSecurity(security),
	}
	client := sdk.New(opts...)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *AirbyteProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewConnectionResource,
		NewDestinationAwsDatalakeResource,
		NewDestinationAzureBlobStorageResource,
		NewDestinationBigqueryResource,
		NewDestinationClickhouseResource,
		NewDestinationConvexResource,
		NewDestinationCumulioResource,
		NewDestinationDatabendResource,
		NewDestinationDatabricksResource,
		NewDestinationDevNullResource,
		NewDestinationDuckdbResource,
		NewDestinationDynamodbResource,
		NewDestinationElasticsearchResource,
		NewDestinationFireboltResource,
		NewDestinationFirestoreResource,
		NewDestinationGcsResource,
		NewDestinationGoogleSheetsResource,
		NewDestinationKeenResource,
		NewDestinationKinesisResource,
		NewDestinationLangchainResource,
		NewDestinationMilvusResource,
		NewDestinationMongodbResource,
		NewDestinationMssqlResource,
		NewDestinationMysqlResource,
		NewDestinationOracleResource,
		NewDestinationPineconeResource,
		NewDestinationPostgresResource,
		NewDestinationPubsubResource,
		NewDestinationQdrantResource,
		NewDestinationRedisResource,
		NewDestinationRedshiftResource,
		NewDestinationS3Resource,
		NewDestinationS3GlueResource,
		NewDestinationSftpJSONResource,
		NewDestinationSnowflakeResource,
		NewDestinationTimeplusResource,
		NewDestinationTypesenseResource,
		NewDestinationVerticaResource,
		NewDestinationWeaviateResource,
		NewDestinationXataResource,
		NewSourceAhaResource,
		NewSourceAircallResource,
		NewSourceAirtableResource,
		NewSourceAlloydbResource,
		NewSourceAmazonAdsResource,
		NewSourceAmazonSellerPartnerResource,
		NewSourceAmazonSqsResource,
		NewSourceAmplitudeResource,
		NewSourceApifyDatasetResource,
		NewSourceAppfollowResource,
		NewSourceAsanaResource,
		NewSourceAuth0Resource,
		NewSourceAwsCloudtrailResource,
		NewSourceAzureBlobStorageResource,
		NewSourceAzureTableResource,
		NewSourceBambooHrResource,
		NewSourceBigqueryResource,
		NewSourceBingAdsResource,
		NewSourceBraintreeResource,
		NewSourceBrazeResource,
		NewSourceCartResource,
		NewSourceChargebeeResource,
		NewSourceChartmogulResource,
		NewSourceClickhouseResource,
		NewSourceClickupAPIResource,
		NewSourceClockifyResource,
		NewSourceCloseComResource,
		NewSourceCodaResource,
		NewSourceCoinAPIResource,
		NewSourceCoinmarketcapResource,
		NewSourceConfigcatResource,
		NewSourceConfluenceResource,
		NewSourceConvexResource,
		NewSourceDatascopeResource,
		NewSourceDelightedResource,
		NewSourceDixaResource,
		NewSourceDockerhubResource,
		NewSourceDremioResource,
		NewSourceDynamodbResource,
		NewSourceEmailoctopusResource,
		NewSourceExchangeRatesResource,
		NewSourceFacebookMarketingResource,
		NewSourceFacebookPagesResource,
		NewSourceFakerResource,
		NewSourceFaunaResource,
		NewSourceFileResource,
		NewSourceFireboltResource,
		NewSourceFreshcallerResource,
		NewSourceFreshdeskResource,
		NewSourceFreshsalesResource,
		NewSourceGainsightPxResource,
		NewSourceGcsResource,
		NewSourceGetlagoResource,
		NewSourceGithubResource,
		NewSourceGitlabResource,
		NewSourceGlassfrogResource,
		NewSourceGnewsResource,
		NewSourceGoogleAdsResource,
		NewSourceGoogleAnalyticsDataAPIResource,
		NewSourceGoogleDirectoryResource,
		NewSourceGoogleDriveResource,
		NewSourceGooglePagespeedInsightsResource,
		NewSourceGoogleSearchConsoleResource,
		NewSourceGoogleSheetsResource,
		NewSourceGoogleWebfontsResource,
		NewSourceGoogleWorkspaceAdminReportsResource,
		NewSourceGreenhouseResource,
		NewSourceGridlyResource,
		NewSourceHarvestResource,
		NewSourceHubplannerResource,
		NewSourceHubspotResource,
		NewSourceInsightlyResource,
		NewSourceInstagramResource,
		NewSourceInstatusResource,
		NewSourceIntercomResource,
		NewSourceIp2whoisResource,
		NewSourceIterableResource,
		NewSourceJiraResource,
		NewSourceK6CloudResource,
		NewSourceKlarnaResource,
		NewSourceKlaviyoResource,
		NewSourceKustomerSingerResource,
		NewSourceKyveResource,
		NewSourceLaunchdarklyResource,
		NewSourceLemlistResource,
		NewSourceLeverHiringResource,
		NewSourceLinkedinAdsResource,
		NewSourceLinkedinPagesResource,
		NewSourceLinnworksResource,
		NewSourceLokaliseResource,
		NewSourceMailchimpResource,
		NewSourceMailgunResource,
		NewSourceMailjetSmsResource,
		NewSourceMarketoResource,
		NewSourceMetabaseResource,
		NewSourceMicrosoftTeamsResource,
		NewSourceMixpanelResource,
		NewSourceMondayResource,
		NewSourceMongodbInternalPocResource,
		NewSourceMongodbV2Resource,
		NewSourceMssqlResource,
		NewSourceMyHoursResource,
		NewSourceMysqlResource,
		NewSourceNetsuiteResource,
		NewSourceNotionResource,
		NewSourceNytimesResource,
		NewSourceOktaResource,
		NewSourceOmnisendResource,
		NewSourceOnesignalResource,
		NewSourceOracleResource,
		NewSourceOrbResource,
		NewSourceOrbitResource,
		NewSourceOutbrainAmplifyResource,
		NewSourceOutreachResource,
		NewSourcePaypalTransactionResource,
		NewSourcePaystackResource,
		NewSourcePendoResource,
		NewSourcePersistiqResource,
		NewSourcePexelsAPIResource,
		NewSourcePinterestResource,
		NewSourcePipedriveResource,
		NewSourcePocketResource,
		NewSourcePokeapiResource,
		NewSourcePolygonStockAPIResource,
		NewSourcePostgresResource,
		NewSourcePosthogResource,
		NewSourcePostmarkappResource,
		NewSourcePrestashopResource,
		NewSourcePunkAPIResource,
		NewSourcePypiResource,
		NewSourceQualarooResource,
		NewSourceQuickbooksResource,
		NewSourceRailzResource,
		NewSourceRechargeResource,
		NewSourceRecreationResource,
		NewSourceRecruiteeResource,
		NewSourceRecurlyResource,
		NewSourceRedshiftResource,
		NewSourceRetentlyResource,
		NewSourceRkiCovidResource,
		NewSourceRssResource,
		NewSourceS3Resource,
		NewSourceSalesforceResource,
		NewSourceSalesloftResource,
		NewSourceSapFieldglassResource,
		NewSourceSecodaResource,
		NewSourceSendgridResource,
		NewSourceSendinblueResource,
		NewSourceSenseforceResource,
		NewSourceSentryResource,
		NewSourceSftpResource,
		NewSourceSftpBulkResource,
		NewSourceShopifyResource,
		NewSourceShortioResource,
		NewSourceSlackResource,
		NewSourceSmailyResource,
		NewSourceSmartengageResource,
		NewSourceSmartsheetsResource,
		NewSourceSnapchatMarketingResource,
		NewSourceSnowflakeResource,
		NewSourceSonarCloudResource,
		NewSourceSpacexAPIResource,
		NewSourceSquareResource,
		NewSourceStravaResource,
		NewSourceStripeResource,
		NewSourceSurveymonkeyResource,
		NewSourceSurveySparrowResource,
		NewSourceTempoResource,
		NewSourceTheGuardianAPIResource,
		NewSourceTiktokMarketingResource,
		NewSourceTodoistResource,
		NewSourceTrelloResource,
		NewSourceTrustpilotResource,
		NewSourceTvmazeScheduleResource,
		NewSourceTwilioResource,
		NewSourceTwilioTaskrouterResource,
		NewSourceTwitterResource,
		NewSourceTypeformResource,
		NewSourceUsCensusResource,
		NewSourceVantageResource,
		NewSourceWebflowResource,
		NewSourceWhiskyHunterResource,
		NewSourceWikipediaPageviewsResource,
		NewSourceWoocommerceResource,
		NewSourceXkcdResource,
		NewSourceYandexMetricaResource,
		NewSourceYotpoResource,
		NewSourceYoutubeAnalyticsResource,
		NewSourceZendeskChatResource,
		NewSourceZendeskSellResource,
		NewSourceZendeskSunshineResource,
		NewSourceZendeskSupportResource,
		NewSourceZendeskTalkResource,
		NewSourceZenloopResource,
		NewSourceZohoCrmResource,
		NewSourceZoomResource,
		NewSourceZuoraResource,
		NewWorkspaceResource,
	}
}

func (p *AirbyteProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewConnectionDataSource,
		NewDestinationAwsDatalakeDataSource,
		NewDestinationAzureBlobStorageDataSource,
		NewDestinationBigqueryDataSource,
		NewDestinationClickhouseDataSource,
		NewDestinationConvexDataSource,
		NewDestinationCumulioDataSource,
		NewDestinationDatabendDataSource,
		NewDestinationDatabricksDataSource,
		NewDestinationDevNullDataSource,
		NewDestinationDuckdbDataSource,
		NewDestinationDynamodbDataSource,
		NewDestinationElasticsearchDataSource,
		NewDestinationFireboltDataSource,
		NewDestinationFirestoreDataSource,
		NewDestinationGcsDataSource,
		NewDestinationGoogleSheetsDataSource,
		NewDestinationKeenDataSource,
		NewDestinationKinesisDataSource,
		NewDestinationLangchainDataSource,
		NewDestinationMilvusDataSource,
		NewDestinationMongodbDataSource,
		NewDestinationMssqlDataSource,
		NewDestinationMysqlDataSource,
		NewDestinationOracleDataSource,
		NewDestinationPineconeDataSource,
		NewDestinationPostgresDataSource,
		NewDestinationPubsubDataSource,
		NewDestinationQdrantDataSource,
		NewDestinationRedisDataSource,
		NewDestinationRedshiftDataSource,
		NewDestinationS3DataSource,
		NewDestinationS3GlueDataSource,
		NewDestinationSftpJSONDataSource,
		NewDestinationSnowflakeDataSource,
		NewDestinationTimeplusDataSource,
		NewDestinationTypesenseDataSource,
		NewDestinationVerticaDataSource,
		NewDestinationWeaviateDataSource,
		NewDestinationXataDataSource,
		NewSourceAhaDataSource,
		NewSourceAircallDataSource,
		NewSourceAirtableDataSource,
		NewSourceAlloydbDataSource,
		NewSourceAmazonAdsDataSource,
		NewSourceAmazonSellerPartnerDataSource,
		NewSourceAmazonSqsDataSource,
		NewSourceAmplitudeDataSource,
		NewSourceApifyDatasetDataSource,
		NewSourceAppfollowDataSource,
		NewSourceAsanaDataSource,
		NewSourceAuth0DataSource,
		NewSourceAwsCloudtrailDataSource,
		NewSourceAzureBlobStorageDataSource,
		NewSourceAzureTableDataSource,
		NewSourceBambooHrDataSource,
		NewSourceBigqueryDataSource,
		NewSourceBingAdsDataSource,
		NewSourceBraintreeDataSource,
		NewSourceBrazeDataSource,
		NewSourceCartDataSource,
		NewSourceChargebeeDataSource,
		NewSourceChartmogulDataSource,
		NewSourceClickhouseDataSource,
		NewSourceClickupAPIDataSource,
		NewSourceClockifyDataSource,
		NewSourceCloseComDataSource,
		NewSourceCodaDataSource,
		NewSourceCoinAPIDataSource,
		NewSourceCoinmarketcapDataSource,
		NewSourceConfigcatDataSource,
		NewSourceConfluenceDataSource,
		NewSourceConvexDataSource,
		NewSourceDatascopeDataSource,
		NewSourceDelightedDataSource,
		NewSourceDixaDataSource,
		NewSourceDockerhubDataSource,
		NewSourceDremioDataSource,
		NewSourceDynamodbDataSource,
		NewSourceEmailoctopusDataSource,
		NewSourceExchangeRatesDataSource,
		NewSourceFacebookMarketingDataSource,
		NewSourceFacebookPagesDataSource,
		NewSourceFakerDataSource,
		NewSourceFaunaDataSource,
		NewSourceFileDataSource,
		NewSourceFireboltDataSource,
		NewSourceFreshcallerDataSource,
		NewSourceFreshdeskDataSource,
		NewSourceFreshsalesDataSource,
		NewSourceGainsightPxDataSource,
		NewSourceGcsDataSource,
		NewSourceGetlagoDataSource,
		NewSourceGithubDataSource,
		NewSourceGitlabDataSource,
		NewSourceGlassfrogDataSource,
		NewSourceGnewsDataSource,
		NewSourceGoogleAdsDataSource,
		NewSourceGoogleAnalyticsDataAPIDataSource,
		NewSourceGoogleDirectoryDataSource,
		NewSourceGoogleDriveDataSource,
		NewSourceGooglePagespeedInsightsDataSource,
		NewSourceGoogleSearchConsoleDataSource,
		NewSourceGoogleSheetsDataSource,
		NewSourceGoogleWebfontsDataSource,
		NewSourceGoogleWorkspaceAdminReportsDataSource,
		NewSourceGreenhouseDataSource,
		NewSourceGridlyDataSource,
		NewSourceHarvestDataSource,
		NewSourceHubplannerDataSource,
		NewSourceHubspotDataSource,
		NewSourceInsightlyDataSource,
		NewSourceInstagramDataSource,
		NewSourceInstatusDataSource,
		NewSourceIntercomDataSource,
		NewSourceIp2whoisDataSource,
		NewSourceIterableDataSource,
		NewSourceJiraDataSource,
		NewSourceK6CloudDataSource,
		NewSourceKlarnaDataSource,
		NewSourceKlaviyoDataSource,
		NewSourceKustomerSingerDataSource,
		NewSourceKyveDataSource,
		NewSourceLaunchdarklyDataSource,
		NewSourceLemlistDataSource,
		NewSourceLeverHiringDataSource,
		NewSourceLinkedinAdsDataSource,
		NewSourceLinkedinPagesDataSource,
		NewSourceLinnworksDataSource,
		NewSourceLokaliseDataSource,
		NewSourceMailchimpDataSource,
		NewSourceMailgunDataSource,
		NewSourceMailjetSmsDataSource,
		NewSourceMarketoDataSource,
		NewSourceMetabaseDataSource,
		NewSourceMicrosoftTeamsDataSource,
		NewSourceMixpanelDataSource,
		NewSourceMondayDataSource,
		NewSourceMongodbInternalPocDataSource,
		NewSourceMongodbV2DataSource,
		NewSourceMssqlDataSource,
		NewSourceMyHoursDataSource,
		NewSourceMysqlDataSource,
		NewSourceNetsuiteDataSource,
		NewSourceNotionDataSource,
		NewSourceNytimesDataSource,
		NewSourceOktaDataSource,
		NewSourceOmnisendDataSource,
		NewSourceOnesignalDataSource,
		NewSourceOracleDataSource,
		NewSourceOrbDataSource,
		NewSourceOrbitDataSource,
		NewSourceOutbrainAmplifyDataSource,
		NewSourceOutreachDataSource,
		NewSourcePaypalTransactionDataSource,
		NewSourcePaystackDataSource,
		NewSourcePendoDataSource,
		NewSourcePersistiqDataSource,
		NewSourcePexelsAPIDataSource,
		NewSourcePinterestDataSource,
		NewSourcePipedriveDataSource,
		NewSourcePocketDataSource,
		NewSourcePokeapiDataSource,
		NewSourcePolygonStockAPIDataSource,
		NewSourcePostgresDataSource,
		NewSourcePosthogDataSource,
		NewSourcePostmarkappDataSource,
		NewSourcePrestashopDataSource,
		NewSourcePunkAPIDataSource,
		NewSourcePypiDataSource,
		NewSourceQualarooDataSource,
		NewSourceQuickbooksDataSource,
		NewSourceRailzDataSource,
		NewSourceRechargeDataSource,
		NewSourceRecreationDataSource,
		NewSourceRecruiteeDataSource,
		NewSourceRecurlyDataSource,
		NewSourceRedshiftDataSource,
		NewSourceRetentlyDataSource,
		NewSourceRkiCovidDataSource,
		NewSourceRssDataSource,
		NewSourceS3DataSource,
		NewSourceSalesforceDataSource,
		NewSourceSalesloftDataSource,
		NewSourceSapFieldglassDataSource,
		NewSourceSecodaDataSource,
		NewSourceSendgridDataSource,
		NewSourceSendinblueDataSource,
		NewSourceSenseforceDataSource,
		NewSourceSentryDataSource,
		NewSourceSftpDataSource,
		NewSourceSftpBulkDataSource,
		NewSourceShopifyDataSource,
		NewSourceShortioDataSource,
		NewSourceSlackDataSource,
		NewSourceSmailyDataSource,
		NewSourceSmartengageDataSource,
		NewSourceSmartsheetsDataSource,
		NewSourceSnapchatMarketingDataSource,
		NewSourceSnowflakeDataSource,
		NewSourceSonarCloudDataSource,
		NewSourceSpacexAPIDataSource,
		NewSourceSquareDataSource,
		NewSourceStravaDataSource,
		NewSourceStripeDataSource,
		NewSourceSurveymonkeyDataSource,
		NewSourceSurveySparrowDataSource,
		NewSourceTempoDataSource,
		NewSourceTheGuardianAPIDataSource,
		NewSourceTiktokMarketingDataSource,
		NewSourceTodoistDataSource,
		NewSourceTrelloDataSource,
		NewSourceTrustpilotDataSource,
		NewSourceTvmazeScheduleDataSource,
		NewSourceTwilioDataSource,
		NewSourceTwilioTaskrouterDataSource,
		NewSourceTwitterDataSource,
		NewSourceTypeformDataSource,
		NewSourceUsCensusDataSource,
		NewSourceVantageDataSource,
		NewSourceWebflowDataSource,
		NewSourceWhiskyHunterDataSource,
		NewSourceWikipediaPageviewsDataSource,
		NewSourceWoocommerceDataSource,
		NewSourceXkcdDataSource,
		NewSourceYandexMetricaDataSource,
		NewSourceYotpoDataSource,
		NewSourceYoutubeAnalyticsDataSource,
		NewSourceZendeskChatDataSource,
		NewSourceZendeskSellDataSource,
		NewSourceZendeskSunshineDataSource,
		NewSourceZendeskSupportDataSource,
		NewSourceZendeskTalkDataSource,
		NewSourceZenloopDataSource,
		NewSourceZohoCrmDataSource,
		NewSourceZoomDataSource,
		NewSourceZuoraDataSource,
		NewWorkspaceDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &AirbyteProvider{
			version: version,
		}
	}
}
