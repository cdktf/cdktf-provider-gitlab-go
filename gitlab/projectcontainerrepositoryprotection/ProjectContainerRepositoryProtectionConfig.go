// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectcontainerrepositoryprotection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectContainerRepositoryProtectionConfig struct {
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
	// ID or URL-encoded path of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_container_repository_protection#project ProjectContainerRepositoryProtection#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Container repository path pattern protected by the protection rule.
	//
	// Wildcard character * allowed. Repository path pattern should start with the project's full path
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_container_repository_protection#repository_path_pattern ProjectContainerRepositoryProtection#repository_path_pattern}
	RepositoryPathPattern *string `field:"required" json:"repositoryPathPattern" yaml:"repositoryPathPattern"`
	// Minimum GitLab access level required to delete container images in the container registry.
	//
	// For example maintainer, owner, admin. Must be provided when `minimum_access_level_for_push` is not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_container_repository_protection#minimum_access_level_for_delete ProjectContainerRepositoryProtection#minimum_access_level_for_delete}
	MinimumAccessLevelForDelete *string `field:"optional" json:"minimumAccessLevelForDelete" yaml:"minimumAccessLevelForDelete"`
	// Minimum GitLab access level required to push container images to the container registry.
	//
	// For example maintainer, owner or admin. Must be provided when `minimum_access_level_for_delete` is not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_container_repository_protection#minimum_access_level_for_push ProjectContainerRepositoryProtection#minimum_access_level_for_push}
	MinimumAccessLevelForPush *string `field:"optional" json:"minimumAccessLevelForPush" yaml:"minimumAccessLevelForPush"`
}

