// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package personalaccesstoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PersonalAccessTokenConfig struct {
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
	// The name of the personal access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/personal_access_token#name PersonalAccessToken#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The scopes of the personal access token.
	//
	// valid values are: `api`, `read_user`, `read_api`, `read_repository`, `write_repository`, `read_registry`, `write_registry`, `read_virtual_registry`, `write_virtual_registry`, `sudo`, `admin_mode`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `self_rotate`, `read_service_ping`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/personal_access_token#scopes PersonalAccessToken#scopes}
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// The ID of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/personal_access_token#user_id PersonalAccessToken#user_id}
	UserId *float64 `field:"required" json:"userId" yaml:"userId"`
	// The description of the personal access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/personal_access_token#description PersonalAccessToken#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// When the token will expire, YYYY-MM-DD format. Is automatically set when `rotation_configuration` is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/personal_access_token#expires_at PersonalAccessToken#expires_at}
	ExpiresAt *string `field:"optional" json:"expiresAt" yaml:"expiresAt"`
	// The configuration for when to rotate a token automatically. Will not rotate a token until `terraform apply` is run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/personal_access_token#rotation_configuration PersonalAccessToken#rotation_configuration}
	RotationConfiguration *PersonalAccessTokenRotationConfiguration `field:"optional" json:"rotationConfiguration" yaml:"rotationConfiguration"`
}

