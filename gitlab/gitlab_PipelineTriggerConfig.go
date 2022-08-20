// Prebuilt gitlab Provider for Terraform CDK (cdktf)
package gitlab

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PipelineTriggerConfig struct {
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
	// The description of the pipeline trigger.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/pipeline_trigger#description PipelineTrigger#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name or id of the project to add the trigger to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/pipeline_trigger#project PipelineTrigger#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/pipeline_trigger#id PipelineTrigger#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}
