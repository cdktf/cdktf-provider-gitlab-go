// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectexternalstatuscheck

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectExternalStatusCheckConfig struct {
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
	// The URL of the external status check service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/project_external_status_check#external_url ProjectExternalStatusCheck#external_url}
	ExternalUrl *string `field:"required" json:"externalUrl" yaml:"externalUrl"`
	// The display name of the external status check service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/project_external_status_check#name ProjectExternalStatusCheck#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/project_external_status_check#project_id ProjectExternalStatusCheck#project_id}
	ProjectId *float64 `field:"required" json:"projectId" yaml:"projectId"`
	// The list of IDs of protected branches to scope the rule by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/project_external_status_check#protected_branch_ids ProjectExternalStatusCheck#protected_branch_ids}
	ProtectedBranchIds *[]*float64 `field:"optional" json:"protectedBranchIds" yaml:"protectedBranchIds"`
	// The HMAC secret for the external status check.
	//
	// If this is set, then removed from the config, the value will get set to empty in the state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/project_external_status_check#shared_secret ProjectExternalStatusCheck#shared_secret}
	SharedSecret *string `field:"optional" json:"sharedSecret" yaml:"sharedSecret"`
}

