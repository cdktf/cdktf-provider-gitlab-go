package runner

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RunnerConfig struct {
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
	// The registration token used to register the runner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.11.0/docs/resources/runner#registration_token Runner#registration_token}
	RegistrationToken *string `field:"required" json:"registrationToken" yaml:"registrationToken"`
	// The access_level of the runner. Valid values are: `not_protected`, `ref_protected`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.11.0/docs/resources/runner#access_level Runner#access_level}
	AccessLevel *string `field:"optional" json:"accessLevel" yaml:"accessLevel"`
	// The runner's description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.11.0/docs/resources/runner#description Runner#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.11.0/docs/resources/runner#id Runner#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether the runner should be locked for current project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.11.0/docs/resources/runner#locked Runner#locked}
	Locked interface{} `field:"optional" json:"locked" yaml:"locked"`
	// Maximum timeout set when this runner handles the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.11.0/docs/resources/runner#maximum_timeout Runner#maximum_timeout}
	MaximumTimeout *float64 `field:"optional" json:"maximumTimeout" yaml:"maximumTimeout"`
	// Whether the runner should ignore new jobs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.11.0/docs/resources/runner#paused Runner#paused}
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
	// Whether the runner should handle untagged jobs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.11.0/docs/resources/runner#run_untagged Runner#run_untagged}
	RunUntagged interface{} `field:"optional" json:"runUntagged" yaml:"runUntagged"`
	// List of runner’s tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.11.0/docs/resources/runner#tag_list Runner#tag_list}
	TagList *[]*string `field:"optional" json:"tagList" yaml:"tagList"`
}

