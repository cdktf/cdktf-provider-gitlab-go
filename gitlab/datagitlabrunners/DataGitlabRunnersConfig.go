// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabrunners

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabRunnersConfig struct {
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
	// Filters for runners with the given paused value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.8.0/docs/data-sources/runners#paused DataGitlabRunners#paused}
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
	// Filters for runners with the given status. Valid Values are `online`, `offline`, `stale`, and `never_contacted`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.8.0/docs/data-sources/runners#status DataGitlabRunners#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Filters for runners with all of the given tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.8.0/docs/data-sources/runners#tag_list DataGitlabRunners#tag_list}
	TagList *[]*string `field:"optional" json:"tagList" yaml:"tagList"`
	// The type of runner to return. Valid values are `instance_type`, `group_type` and `project_type`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.8.0/docs/data-sources/runners#type DataGitlabRunners#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

