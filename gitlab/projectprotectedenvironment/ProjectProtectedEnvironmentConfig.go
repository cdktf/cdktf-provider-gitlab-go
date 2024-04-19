// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectprotectedenvironment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectProtectedEnvironmentConfig struct {
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
	// The name of the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.11.0/docs/resources/project_protected_environment#environment ProjectProtectedEnvironment#environment}
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// The ID or full path of the project which the protected environment is created against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.11.0/docs/resources/project_protected_environment#project ProjectProtectedEnvironment#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Array of approval rules to deploy, with each described by a hash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.11.0/docs/resources/project_protected_environment#approval_rules ProjectProtectedEnvironment#approval_rules}
	ApprovalRules interface{} `field:"optional" json:"approvalRules" yaml:"approvalRules"`
	// deploy_access_levels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.11.0/docs/resources/project_protected_environment#deploy_access_levels ProjectProtectedEnvironment#deploy_access_levels}
	DeployAccessLevels interface{} `field:"optional" json:"deployAccessLevels" yaml:"deployAccessLevels"`
	// The number of approvals required to deploy to this environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.11.0/docs/resources/project_protected_environment#required_approval_count ProjectProtectedEnvironment#required_approval_count}
	RequiredApprovalCount *float64 `field:"optional" json:"requiredApprovalCount" yaml:"requiredApprovalCount"`
}

