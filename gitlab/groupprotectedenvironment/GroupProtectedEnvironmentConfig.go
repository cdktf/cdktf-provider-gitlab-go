// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupprotectedenvironment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupProtectedEnvironmentConfig struct {
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
	// Array of access levels allowed to deploy, with each described by a hash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.9.1/docs/resources/group_protected_environment#deploy_access_levels GroupProtectedEnvironment#deploy_access_levels}
	DeployAccessLevels interface{} `field:"required" json:"deployAccessLevels" yaml:"deployAccessLevels"`
	// The deployment tier of the environment.  Valid values are `production`, `staging`, `testing`, `development`, `other`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.9.1/docs/resources/group_protected_environment#environment GroupProtectedEnvironment#environment}
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// The ID or full path of the group which the protected environment is created against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.9.1/docs/resources/group_protected_environment#group GroupProtectedEnvironment#group}
	Group *string `field:"required" json:"group" yaml:"group"`
	// Array of approval rules to deploy, with each described by a hash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.9.1/docs/resources/group_protected_environment#approval_rules GroupProtectedEnvironment#approval_rules}
	ApprovalRules interface{} `field:"optional" json:"approvalRules" yaml:"approvalRules"`
	// The number of approvals required to deploy to this environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.9.1/docs/resources/group_protected_environment#required_approval_count GroupProtectedEnvironment#required_approval_count}
	RequiredApprovalCount *float64 `field:"optional" json:"requiredApprovalCount" yaml:"requiredApprovalCount"`
}

