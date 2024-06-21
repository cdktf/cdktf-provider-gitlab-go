// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectcomplianceframework

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectComplianceFrameworkConfig struct {
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
	// Globally unique ID of the compliance framework to assign to the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.1.0/docs/resources/project_compliance_framework#compliance_framework_id ProjectComplianceFramework#compliance_framework_id}
	ComplianceFrameworkId *string `field:"required" json:"complianceFrameworkId" yaml:"complianceFrameworkId"`
	// The ID or full path of the project to change the compliance framework of.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.1.0/docs/resources/project_compliance_framework#project ProjectComplianceFramework#project}
	Project *string `field:"required" json:"project" yaml:"project"`
}

