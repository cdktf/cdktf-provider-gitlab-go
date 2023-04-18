package projectissueboard

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectIssueBoardConfig struct {
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
	// The name of the board.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.10.0/docs/resources/project_issue_board#name ProjectIssueBoard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID or full path of the project maintained by the authenticated user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.10.0/docs/resources/project_issue_board#project ProjectIssueBoard#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The assignee the board should be scoped to. Requires a GitLab EE license.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.10.0/docs/resources/project_issue_board#assignee_id ProjectIssueBoard#assignee_id}
	AssigneeId *float64 `field:"optional" json:"assigneeId" yaml:"assigneeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.10.0/docs/resources/project_issue_board#id ProjectIssueBoard#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The list of label names which the board should be scoped to. Requires a GitLab EE license.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.10.0/docs/resources/project_issue_board#labels ProjectIssueBoard#labels}
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// lists block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.10.0/docs/resources/project_issue_board#lists ProjectIssueBoard#lists}
	Lists interface{} `field:"optional" json:"lists" yaml:"lists"`
	// The milestone the board should be scoped to. Requires a GitLab EE license.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.10.0/docs/resources/project_issue_board#milestone_id ProjectIssueBoard#milestone_id}
	MilestoneId *float64 `field:"optional" json:"milestoneId" yaml:"milestoneId"`
	// The weight range from 0 to 9, to which the board should be scoped to.
	//
	// Requires a GitLab EE license.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.10.0/docs/resources/project_issue_board#weight ProjectIssueBoard#weight}
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

