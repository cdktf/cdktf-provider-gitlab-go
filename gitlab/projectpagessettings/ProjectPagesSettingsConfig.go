// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectpagessettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectPagesSettingsConfig struct {
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
	// The project ID or path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_pages_settings#project ProjectPagesSettings#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Boolean indicating if the project is set to force https.
	//
	// Requires `external_https` to be configured in the GitLab instance: https://docs.gitlab.com/administration/pages/#custom-domains-with-tls-support.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_pages_settings#force_https ProjectPagesSettings#force_https}
	ForceHttps interface{} `field:"optional" json:"forceHttps" yaml:"forceHttps"`
	// Boolean indicating if a unique domain is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_pages_settings#is_unique_domain_enabled ProjectPagesSettings#is_unique_domain_enabled}
	IsUniqueDomainEnabled interface{} `field:"optional" json:"isUniqueDomainEnabled" yaml:"isUniqueDomainEnabled"`
	// Set to true if the pages settings should not be reset to their pre-terraform defaults on destroy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_pages_settings#keep_settings_on_destroy ProjectPagesSettings#keep_settings_on_destroy}
	KeepSettingsOnDestroy interface{} `field:"optional" json:"keepSettingsOnDestroy" yaml:"keepSettingsOnDestroy"`
}

