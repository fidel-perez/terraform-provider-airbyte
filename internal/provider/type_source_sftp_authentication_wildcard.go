// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

type SourceSftpAuthenticationWildcard struct {
	PasswordAuthentication *SourceSftpPasswordAuthentication `tfsdk:"password_authentication"`
	SSHKeyAuthentication   *SourceSftpSSHKeyAuthentication   `tfsdk:"ssh_key_authentication"`
}
