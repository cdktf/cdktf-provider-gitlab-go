// Prebuilt gitlab Provider for Terraform CDK (cdktf)
package gitlab

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabGroupHookConfig struct {
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
	// The ID or full path of the group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/group_hook#group DataGitlabGroupHook#group}
	Group *string `field:"required" json:"group" yaml:"group"`
	// The id of the group hook.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/group_hook#hook_id DataGitlabGroupHook#hook_id}
	HookId *float64 `field:"required" json:"hookId" yaml:"hookId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/group_hook#id DataGitlabGroupHook#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}
