// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabgroupprovisionedusers

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabGroupProvisionedUsersConfig struct {
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
	// The ID or URL-encoded path of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/data-sources/group_provisioned_users#id DataGitlabGroupProvisionedUsers#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Return only active provisioned users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/data-sources/group_provisioned_users#active DataGitlabGroupProvisionedUsers#active}
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// Return only blocked provisioned users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/data-sources/group_provisioned_users#blocked DataGitlabGroupProvisionedUsers#blocked}
	Blocked interface{} `field:"optional" json:"blocked" yaml:"blocked"`
	// Return only provisioned users created on or after the specified date. Expected in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/data-sources/group_provisioned_users#created_after DataGitlabGroupProvisionedUsers#created_after}
	CreatedAfter *string `field:"optional" json:"createdAfter" yaml:"createdAfter"`
	// Return only provisioned users created on or before the specified date. Expected in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/data-sources/group_provisioned_users#created_before DataGitlabGroupProvisionedUsers#created_before}
	CreatedBefore *string `field:"optional" json:"createdBefore" yaml:"createdBefore"`
	// provisioned_users block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/data-sources/group_provisioned_users#provisioned_users DataGitlabGroupProvisionedUsers#provisioned_users}
	ProvisionedUsers interface{} `field:"optional" json:"provisionedUsers" yaml:"provisionedUsers"`
	// The search query to filter the provisioned users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/data-sources/group_provisioned_users#search DataGitlabGroupProvisionedUsers#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
	// The username of the provisioned user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/data-sources/group_provisioned_users#username DataGitlabGroupProvisionedUsers#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

