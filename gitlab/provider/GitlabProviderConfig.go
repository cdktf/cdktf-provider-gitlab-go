// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type GitlabProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs#alias GitlabProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// This is the target GitLab base API endpoint.
	//
	// Providing a value is a requirement when working with GitLab CE or GitLab Enterprise e.g. `https://my.gitlab.server/api/v4/`. It is optional to provide this value and it can also be sourced from the `GITLAB_BASE_URL` environment variable. The value must end with a slash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs#base_url GitlabProvider#base_url}
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// This is a file containing the ca cert to verify the gitlab instance.
	//
	// This is available for use when working with GitLab CE or Gitlab Enterprise with a locally-issued or self-signed certificate chain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs#cacert_file GitlabProvider#cacert_file}
	CacertFile *string `field:"optional" json:"cacertFile" yaml:"cacertFile"`
	// File path to client certificate when GitLab instance is behind company proxy. File must contain PEM encoded data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs#client_cert GitlabProvider#client_cert}
	ClientCert *string `field:"optional" json:"clientCert" yaml:"clientCert"`
	// File path to client key when GitLab instance is behind company proxy.
	//
	// File must contain PEM encoded data. Required when `client_cert` is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs#client_key GitlabProvider#client_key}
	ClientKey *string `field:"optional" json:"clientKey" yaml:"clientKey"`
	// (Experimental) By default the provider does a dummy request to get the current user in order to verify that the provider configuration is correct and the GitLab API is reachable.
	//
	// Set this to `false` to skip this check. This may be useful if the GitLab instance does not yet exist and is created within the same terraform module. It may be sourced from the `GITLAB_EARLY_AUTH_CHECK`. This is an experimental feature and may change in the future. Please make sure to always keep backups of your state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs#early_auth_check GitlabProvider#early_auth_check}
	EarlyAuthCheck interface{} `field:"optional" json:"earlyAuthCheck" yaml:"earlyAuthCheck"`
	// A map of headers to append to all API request to the GitLab instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs#headers GitlabProvider#headers}
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
	// When set to true this disables SSL verification of the connection to the GitLab instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs#insecure GitlabProvider#insecure}
	Insecure interface{} `field:"optional" json:"insecure" yaml:"insecure"`
	// The number of retries to execute when receiving a 429 Rate Limit error. Each retry will exponentially back off.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs#retries GitlabProvider#retries}
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// The OAuth2 Token, Project, Group, Personal Access Token or CI Job Token used to connect to GitLab.
	//
	// The OAuth method is used in this provider for authentication (using Bearer authorization token). See https://docs.gitlab.com/ee/api/#authentication for details. It may be sourced from the `GITLAB_TOKEN` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.0/docs#token GitlabProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
}

