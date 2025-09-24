// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package user

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserConfig struct {
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
	// The e-mail address of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/user#email User#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// The name of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/user#name User#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The username of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/user#username User#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Boolean, defaults to false. Whether to allow the user to create groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/user#can_create_group User#can_create_group}
	CanCreateGroup interface{} `field:"optional" json:"canCreateGroup" yaml:"canCreateGroup"`
	// Set user password to a random value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/user#force_random_password User#force_random_password}
	ForceRandomPassword interface{} `field:"optional" json:"forceRandomPassword" yaml:"forceRandomPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/user#id User#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Boolean, defaults to false.  Whether to enable administrative privileges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/user#is_admin User#is_admin}
	IsAdmin interface{} `field:"optional" json:"isAdmin" yaml:"isAdmin"`
	// Boolean, defaults to false.
	//
	// Whether a user has access only to some internal or private projects. External users can only access projects to which they are explicitly granted access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/user#is_external User#is_external}
	IsExternal interface{} `field:"optional" json:"isExternal" yaml:"isExternal"`
	// The ID of the user's namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/user#namespace_id User#namespace_id}
	NamespaceId *float64 `field:"optional" json:"namespaceId" yaml:"namespaceId"`
	// The note associated to the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/user#note User#note}
	Note *string `field:"optional" json:"note" yaml:"note"`
	// The password of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/user#password User#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Integer, defaults to 0.  Number of projects user can create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/user#projects_limit User#projects_limit}
	ProjectsLimit *float64 `field:"optional" json:"projectsLimit" yaml:"projectsLimit"`
	// Boolean, defaults to false. Send user password reset link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/user#reset_password User#reset_password}
	ResetPassword interface{} `field:"optional" json:"resetPassword" yaml:"resetPassword"`
	// Boolean, defaults to true. Whether to skip confirmation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/user#skip_confirmation User#skip_confirmation}
	SkipConfirmation interface{} `field:"optional" json:"skipConfirmation" yaml:"skipConfirmation"`
	// String, defaults to 'active'. The state of the user account. Valid values are `active`, `deactivated`, `blocked`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/user#state User#state}
	State *string `field:"optional" json:"state" yaml:"state"`
}

