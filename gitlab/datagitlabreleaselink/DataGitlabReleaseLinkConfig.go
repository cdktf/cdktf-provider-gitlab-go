// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabreleaselink

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabReleaseLinkConfig struct {
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
	// The ID of the link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/data-sources/release_link#link_id DataGitlabReleaseLink#link_id}
	LinkId *float64 `field:"required" json:"linkId" yaml:"linkId"`
	// The ID or [URL-encoded path of the project](https://docs.gitlab.com/api/index/#namespaced-path-encoding).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/data-sources/release_link#project DataGitlabReleaseLink#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The tag associated with the Release.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/data-sources/release_link#tag_name DataGitlabReleaseLink#tag_name}
	TagName *string `field:"required" json:"tagName" yaml:"tagName"`
}

