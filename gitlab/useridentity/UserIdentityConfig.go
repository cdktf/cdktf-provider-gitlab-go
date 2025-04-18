// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package useridentity

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserIdentityConfig struct {
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
	// The external provider name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.11.0/docs/resources/user_identity#external_provider UserIdentity#external_provider}
	ExternalProvider *string `field:"required" json:"externalProvider" yaml:"externalProvider"`
	// A specific external authentication provider UID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.11.0/docs/resources/user_identity#external_uid UserIdentity#external_uid}
	ExternalUid *string `field:"required" json:"externalUid" yaml:"externalUid"`
	// The GitLab ID of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.11.0/docs/resources/user_identity#user_id UserIdentity#user_id}
	UserId *float64 `field:"required" json:"userId" yaml:"userId"`
}

