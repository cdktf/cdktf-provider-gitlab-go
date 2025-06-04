// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectmirror

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectMirrorConfig struct {
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
	// The id of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/resources/project_mirror#project ProjectMirror#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The URL of the remote repository to be mirrored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/resources/project_mirror#url ProjectMirror#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// Determines the mirror authentication method. Valid values are: `ssh_public_key`, `password`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/resources/project_mirror#auth_method ProjectMirror#auth_method}
	AuthMethod *string `field:"optional" json:"authMethod" yaml:"authMethod"`
	// Determines if the mirror is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/resources/project_mirror#enabled ProjectMirror#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Determines if divergent refs are skipped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/resources/project_mirror#keep_divergent_refs ProjectMirror#keep_divergent_refs}
	KeepDivergentRefs interface{} `field:"optional" json:"keepDivergentRefs" yaml:"keepDivergentRefs"`
	// Contains a regular expression.
	//
	// Only branches with names matching the regex are mirrored. Requires only_protected_branches to be disabled. Premium and Ultimate only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/resources/project_mirror#mirror_branch_regex ProjectMirror#mirror_branch_regex}
	MirrorBranchRegex *string `field:"optional" json:"mirrorBranchRegex" yaml:"mirrorBranchRegex"`
	// Determines if only protected branches are mirrored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/resources/project_mirror#only_protected_branches ProjectMirror#only_protected_branches}
	OnlyProtectedBranches interface{} `field:"optional" json:"onlyProtectedBranches" yaml:"onlyProtectedBranches"`
}

