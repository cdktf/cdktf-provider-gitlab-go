// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectcomplianceframeworks

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectComplianceFrameworksConfig struct {
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
	// Globally unique IDs of the compliance frameworks to assign to the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_compliance_frameworks#compliance_framework_ids ProjectComplianceFrameworks#compliance_framework_ids}
	ComplianceFrameworkIds *[]*string `field:"required" json:"complianceFrameworkIds" yaml:"complianceFrameworkIds"`
	// The ID or full path of the project to change the compliance frameworks of.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_compliance_frameworks#project ProjectComplianceFrameworks#project}
	Project *string `field:"required" json:"project" yaml:"project"`
}

