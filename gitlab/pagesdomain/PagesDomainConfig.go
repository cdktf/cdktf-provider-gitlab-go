// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pagesdomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PagesDomainConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The custom domain indicated by the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.11.0/docs/resources/pages_domain#domain PagesDomain#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding) owned by the authenticated user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.11.0/docs/resources/pages_domain#project PagesDomain#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Enables [automatic generation](https://docs.gitlab.com/ee/user/project/pages/custom_domains_ssl_tls_certification/lets_encrypt_integration.html) of SSL certificates issued by Let’s Encrypt for custom domains. When this is set to "true", certificate can't be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.11.0/docs/resources/pages_domain#auto_ssl_enabled PagesDomain#auto_ssl_enabled}
	AutoSslEnabled interface{} `field:"optional" json:"autoSslEnabled" yaml:"autoSslEnabled"`
	// The certificate in PEM format with intermediates following in most specific to least specific order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.11.0/docs/resources/pages_domain#certificate PagesDomain#certificate}
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// Whether the certificate is expired.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.11.0/docs/resources/pages_domain#expired PagesDomain#expired}
	Expired interface{} `field:"optional" json:"expired" yaml:"expired"`
	// The certificate key in PEM format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.11.0/docs/resources/pages_domain#key PagesDomain#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
}

