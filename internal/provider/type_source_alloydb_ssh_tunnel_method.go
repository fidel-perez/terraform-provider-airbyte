// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

type SourceAlloydbSSHTunnelMethod struct {
	SourceAlloydbSSHTunnelMethodNoTunnel                     *DestinationAzureBlobStorageOutputFormatJSONLinesNewlineDelimitedJSON `tfsdk:"source_alloydb_ssh_tunnel_method_no_tunnel"`
	SourceAlloydbSSHTunnelMethodPasswordAuthentication       *DestinationClickhouseSSHTunnelMethodPasswordAuthentication           `tfsdk:"source_alloydb_ssh_tunnel_method_password_authentication"`
	SourceAlloydbSSHTunnelMethodSSHKeyAuthentication         *DestinationClickhouseSSHTunnelMethodSSHKeyAuthentication             `tfsdk:"source_alloydb_ssh_tunnel_method_ssh_key_authentication"`
	SourceAlloydbUpdateSSHTunnelMethodNoTunnel               *DestinationAzureBlobStorageOutputFormatJSONLinesNewlineDelimitedJSON `tfsdk:"source_alloydb_update_ssh_tunnel_method_no_tunnel"`
	SourceAlloydbUpdateSSHTunnelMethodPasswordAuthentication *DestinationClickhouseSSHTunnelMethodPasswordAuthentication           `tfsdk:"source_alloydb_update_ssh_tunnel_method_password_authentication"`
	SourceAlloydbUpdateSSHTunnelMethodSSHKeyAuthentication   *DestinationClickhouseSSHTunnelMethodSSHKeyAuthentication             `tfsdk:"source_alloydb_update_ssh_tunnel_method_ssh_key_authentication"`
}
