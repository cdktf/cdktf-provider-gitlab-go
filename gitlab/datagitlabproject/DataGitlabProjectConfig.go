// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabproject

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabProjectConfig struct {
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
	// Default number of revisions for shallow cloning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.11.0/docs/data-sources/project#ci_default_git_depth DataGitlabProject#ci_default_git_depth}
	CiDefaultGitDepth *float64 `field:"optional" json:"ciDefaultGitDepth" yaml:"ciDefaultGitDepth"`
	// The integer that uniquely identifies the project within the gitlab install.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.11.0/docs/data-sources/project#id DataGitlabProject#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The path of the repository with namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.11.0/docs/data-sources/project#path_with_namespace DataGitlabProject#path_with_namespace}
	PathWithNamespace *string `field:"optional" json:"pathWithNamespace" yaml:"pathWithNamespace"`
	// If true, jobs can be viewed by non-project members.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.11.0/docs/data-sources/project#public_builds DataGitlabProject#public_builds}
	PublicBuilds interface{} `field:"optional" json:"publicBuilds" yaml:"publicBuilds"`
}

