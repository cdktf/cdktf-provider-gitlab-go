// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package valuestreamanalytics

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ValueStreamAnalyticsConfig struct {
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
	// The name of the value stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/value_stream_analytics#name ValueStreamAnalytics#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Stages of the value stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/value_stream_analytics#stages ValueStreamAnalytics#stages}
	Stages interface{} `field:"required" json:"stages" yaml:"stages"`
	// Full path of the group the value stream is created in. **One of `group_full_path` OR `project_full_path` is required.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/value_stream_analytics#group_full_path ValueStreamAnalytics#group_full_path}
	GroupFullPath *string `field:"optional" json:"groupFullPath" yaml:"groupFullPath"`
	// Full path of the project the value stream is created in. **One of `group_full_path` OR `project_full_path` is required.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/value_stream_analytics#project_full_path ValueStreamAnalytics#project_full_path}
	ProjectFullPath *string `field:"optional" json:"projectFullPath" yaml:"projectFullPath"`
}

