// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package branchprotection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BranchProtectionConfig struct {
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
	// Name of the branch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/branch_protection#branch BranchProtection#branch}
	Branch *string `field:"required" json:"branch" yaml:"branch"`
	// The id of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/branch_protection#project BranchProtection#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// allowed_to_merge block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/branch_protection#allowed_to_merge BranchProtection#allowed_to_merge}
	AllowedToMerge interface{} `field:"optional" json:"allowedToMerge" yaml:"allowedToMerge"`
	// allowed_to_push block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/branch_protection#allowed_to_push BranchProtection#allowed_to_push}
	AllowedToPush interface{} `field:"optional" json:"allowedToPush" yaml:"allowedToPush"`
	// allowed_to_unprotect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/branch_protection#allowed_to_unprotect BranchProtection#allowed_to_unprotect}
	AllowedToUnprotect interface{} `field:"optional" json:"allowedToUnprotect" yaml:"allowedToUnprotect"`
	// Can be set to true to allow users with push access to force push.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/branch_protection#allow_force_push BranchProtection#allow_force_push}
	AllowForcePush interface{} `field:"optional" json:"allowForcePush" yaml:"allowForcePush"`
	// Can be set to true to require code owner approval before merging. Only available for Premium and Ultimate instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/branch_protection#code_owner_approval_required BranchProtection#code_owner_approval_required}
	CodeOwnerApprovalRequired interface{} `field:"optional" json:"codeOwnerApprovalRequired" yaml:"codeOwnerApprovalRequired"`
	// Access levels allowed to merge. Valid values are: `no one`, `developer`, `maintainer`, `admin`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/branch_protection#merge_access_level BranchProtection#merge_access_level}
	MergeAccessLevel *string `field:"optional" json:"mergeAccessLevel" yaml:"mergeAccessLevel"`
	// Access levels allowed to push. Valid values are: `no one`, `developer`, `maintainer`, `admin`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/branch_protection#push_access_level BranchProtection#push_access_level}
	PushAccessLevel *string `field:"optional" json:"pushAccessLevel" yaml:"pushAccessLevel"`
	// Access levels allowed to unprotect. Valid values are: `developer`, `maintainer`, `admin`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/branch_protection#unprotect_access_level BranchProtection#unprotect_access_level}
	UnprotectAccessLevel *string `field:"optional" json:"unprotectAccessLevel" yaml:"unprotectAccessLevel"`
}

