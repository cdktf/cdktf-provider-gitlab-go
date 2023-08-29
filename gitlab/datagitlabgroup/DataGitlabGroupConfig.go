// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabGroupConfig struct {
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
	// The full path of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/data-sources/group#full_path DataGitlabGroup#full_path}
	FullPath *string `field:"optional" json:"fullPath" yaml:"fullPath"`
	// The ID of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/data-sources/group#group_id DataGitlabGroup#group_id}
	GroupId *float64 `field:"optional" json:"groupId" yaml:"groupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/data-sources/group#id DataGitlabGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

