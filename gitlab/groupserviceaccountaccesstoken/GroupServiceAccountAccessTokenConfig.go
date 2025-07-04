// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupserviceaccountaccesstoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupServiceAccountAccessTokenConfig struct {
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
	// The ID or URL-encoded path of the group containing the service account. Must be a top level group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_service_account_access_token#group GroupServiceAccountAccessToken#group}
	Group *string `field:"required" json:"group" yaml:"group"`
	// The name of the personal access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_service_account_access_token#name GroupServiceAccountAccessToken#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The scopes of the group service account access token.
	//
	// Valid values are: `api`, `read_user`, `read_api`, `read_repository`, `write_repository`, `read_registry`, `write_registry`, `read_virtual_registry`, `write_virtual_registry`, `sudo`, `admin_mode`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `self_rotate`, `read_service_ping`. If `self_rotate` is included, you must also provide either `expires_at` or `rotation_configuration`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_service_account_access_token#scopes GroupServiceAccountAccessToken#scopes}
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// The ID of a service account user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_service_account_access_token#user_id GroupServiceAccountAccessToken#user_id}
	UserId *float64 `field:"required" json:"userId" yaml:"userId"`
	// The service account access token expiry date.
	//
	// When left blank, the token follows the standard rule of expiry for personal access tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_service_account_access_token#expires_at GroupServiceAccountAccessToken#expires_at}
	ExpiresAt *string `field:"optional" json:"expiresAt" yaml:"expiresAt"`
	// The configuration for when to rotate a token automatically. Will not rotate a token until `terraform apply` is run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_service_account_access_token#rotation_configuration GroupServiceAccountAccessToken#rotation_configuration}
	RotationConfiguration *GroupServiceAccountAccessTokenRotationConfiguration `field:"optional" json:"rotationConfiguration" yaml:"rotationConfiguration"`
}

