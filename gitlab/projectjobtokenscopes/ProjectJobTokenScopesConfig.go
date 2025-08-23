// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectjobtokenscopes

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectJobTokenScopesConfig struct {
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
	// Enable the given inbound allowlist.
	//
	// If false, will allow any project or group regardless of the values in `target_project_ids` or `target_group_ids`. Deleting the associated `gitlab_project_job_token_scopes` resource will reset `Enabled` on the group to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_job_token_scopes#enabled ProjectJobTokenScopes#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The ID or full path of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_job_token_scopes#project ProjectJobTokenScopes#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The ID of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_job_token_scopes#project_id ProjectJobTokenScopes#project_id}
	ProjectId *float64 `field:"optional" json:"projectId" yaml:"projectId"`
	// A set of group IDs that are in the CI/CD job token inbound allowlist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_job_token_scopes#target_group_ids ProjectJobTokenScopes#target_group_ids}
	TargetGroupIds *[]*float64 `field:"optional" json:"targetGroupIds" yaml:"targetGroupIds"`
	// A set of project IDs that are in the CI/CD job token inbound allowlist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_job_token_scopes#target_project_ids ProjectJobTokenScopes#target_project_ids}
	TargetProjectIds *[]*float64 `field:"optional" json:"targetProjectIds" yaml:"targetProjectIds"`
}

