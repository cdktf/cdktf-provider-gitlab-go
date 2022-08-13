// Prebuilt gitlab Provider for Terraform CDK (cdktf)
package gitlab

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabGroupVariablesConfig struct {
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
	// The name or id of the group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/group_variables#group DataGitlabGroupVariables#group}
	Group *string `field:"required" json:"group" yaml:"group"`
	// The environment scope of the variable. Defaults to all environment (`*`).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/group_variables#environment_scope DataGitlabGroupVariables#environment_scope}
	EnvironmentScope *string `field:"optional" json:"environmentScope" yaml:"environmentScope"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/group_variables#id DataGitlabGroupVariables#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

