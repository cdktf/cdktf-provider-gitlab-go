// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectaccesstoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectAccessTokenConfig struct {
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
	// The name of the project access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_access_token#name ProjectAccessToken#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID or full path of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_access_token#project ProjectAccessToken#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The scopes of the project access token.
	//
	// valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`, `self_rotate`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_access_token#scopes ProjectAccessToken#scopes}
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// The access level for the project access token.
	//
	// Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`. Default is `maintainer`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_access_token#access_level ProjectAccessToken#access_level}
	AccessLevel *string `field:"optional" json:"accessLevel" yaml:"accessLevel"`
	// The description of the project access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_access_token#description ProjectAccessToken#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// When the token will expire, YYYY-MM-DD format. Is automatically set when `rotation_configuration` is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_access_token#expires_at ProjectAccessToken#expires_at}
	ExpiresAt *string `field:"optional" json:"expiresAt" yaml:"expiresAt"`
	// The configuration for when to rotate a token automatically. Will not rotate a token until `terraform apply` is run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_access_token#rotation_configuration ProjectAccessToken#rotation_configuration}
	RotationConfiguration *ProjectAccessTokenRotationConfiguration `field:"optional" json:"rotationConfiguration" yaml:"rotationConfiguration"`
}

