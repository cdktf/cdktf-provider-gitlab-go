// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package userimpersonationtoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserImpersonationTokenConfig struct {
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
	// Expiration date of the impersonation token in ISO format (YYYY-MM-DD).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.8.0/docs/resources/user_impersonation_token#expires_at UserImpersonationToken#expires_at}
	ExpiresAt *string `field:"required" json:"expiresAt" yaml:"expiresAt"`
	// The name of the impersonation token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.8.0/docs/resources/user_impersonation_token#name UserImpersonationToken#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Array of scopes of the impersonation token.
	//
	// valid values are: `api`, `read_user`, `read_api`, `read_repository`, `write_repository`, `read_registry`, `write_registry`, `sudo`, `admin_mode`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_service_ping`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.8.0/docs/resources/user_impersonation_token#scopes UserImpersonationToken#scopes}
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// The ID of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.8.0/docs/resources/user_impersonation_token#user_id UserImpersonationToken#user_id}
	UserId *float64 `field:"required" json:"userId" yaml:"userId"`
}

