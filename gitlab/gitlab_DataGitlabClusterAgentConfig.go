// Prebuilt gitlab Provider for Terraform CDK (cdktf)
package gitlab

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabClusterAgentConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/cluster_agent#agent_id DataGitlabClusterAgent#agent_id}
	AgentId *float64 `field:"required" json:"agentId" yaml:"agentId"`
	// ID or full path of the project maintained by the authenticated user.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/cluster_agent#project DataGitlabClusterAgent#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/cluster_agent#id DataGitlabClusterAgent#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}
