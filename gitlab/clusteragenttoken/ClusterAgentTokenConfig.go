// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clusteragenttoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClusterAgentTokenConfig struct {
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
	// The ID of the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/cluster_agent_token#agent_id ClusterAgentToken#agent_id}
	AgentId *float64 `field:"required" json:"agentId" yaml:"agentId"`
	// The Name of the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/cluster_agent_token#name ClusterAgentToken#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// ID or full path of the project maintained by the authenticated user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/cluster_agent_token#project ClusterAgentToken#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The Description for the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.1/docs/resources/cluster_agent_token#description ClusterAgentToken#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

