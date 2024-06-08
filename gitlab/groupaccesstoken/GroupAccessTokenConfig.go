// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupaccesstoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupAccessTokenConfig struct {
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
	// The ID or full path of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.0.1/docs/resources/group_access_token#group GroupAccessToken#group}
	Group *string `field:"required" json:"group" yaml:"group"`
	// The name of the group access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.0.1/docs/resources/group_access_token#name GroupAccessToken#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The scopes of the group access token.
	//
	// Valid values are: `api`, `read_api`, `read_user`, `k8s_proxy`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.0.1/docs/resources/group_access_token#scopes GroupAccessToken#scopes}
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// The access level for the group access token.
	//
	// Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`. Default is `maintainer`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.0.1/docs/resources/group_access_token#access_level GroupAccessToken#access_level}
	AccessLevel *string `field:"optional" json:"accessLevel" yaml:"accessLevel"`
	// When the token will expire, YYYY-MM-DD format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.0.1/docs/resources/group_access_token#expires_at GroupAccessToken#expires_at}
	ExpiresAt *string `field:"optional" json:"expiresAt" yaml:"expiresAt"`
	// The configuration for when to rotate a token automatically. Will not rotate a token until `terraform apply` is run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.0.1/docs/resources/group_access_token#rotation_configuration GroupAccessToken#rotation_configuration}
	RotationConfiguration *GroupAccessTokenRotationConfiguration `field:"optional" json:"rotationConfiguration" yaml:"rotationConfiguration"`
}

