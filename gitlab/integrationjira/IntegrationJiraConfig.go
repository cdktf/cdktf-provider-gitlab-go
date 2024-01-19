// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationjira

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationJiraConfig struct {
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
	// The password of the user created to be used with GitLab/JIRA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/integration_jira#password IntegrationJira#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// ID of the project you want to activate integration on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/integration_jira#project IntegrationJira#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/integration_jira#url IntegrationJira#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// The username of the user created to be used with GitLab/JIRA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/integration_jira#username IntegrationJira#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// The base URL to the Jira instance API. Web URL value is used if not set. For example, https://jira-api.example.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/integration_jira#api_url IntegrationJira#api_url}
	ApiUrl *string `field:"optional" json:"apiUrl" yaml:"apiUrl"`
	// Enable comments inside Jira issues on each GitLab event (commit / merge request).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/integration_jira#comment_on_event_enabled IntegrationJira#comment_on_event_enabled}
	CommentOnEventEnabled interface{} `field:"optional" json:"commentOnEventEnabled" yaml:"commentOnEventEnabled"`
	// Enable notifications for commit events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/integration_jira#commit_events IntegrationJira#commit_events}
	CommitEvents interface{} `field:"optional" json:"commitEvents" yaml:"commitEvents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/integration_jira#id IntegrationJira#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Enable notifications for issues events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/integration_jira#issues_events IntegrationJira#issues_events}
	IssuesEvents interface{} `field:"optional" json:"issuesEvents" yaml:"issuesEvents"`
	// The ID of a transition that moves issues to a closed state.
	//
	// You can find this number under the JIRA workflow administration (Administration > Issues > Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2. *Note**: importing this field is only supported since GitLab 15.2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/integration_jira#jira_issue_transition_id IntegrationJira#jira_issue_transition_id}
	JiraIssueTransitionId *string `field:"optional" json:"jiraIssueTransitionId" yaml:"jiraIssueTransitionId"`
	// Enable notifications for job events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/integration_jira#job_events IntegrationJira#job_events}
	JobEvents interface{} `field:"optional" json:"jobEvents" yaml:"jobEvents"`
	// Enable notifications for merge request events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/integration_jira#merge_requests_events IntegrationJira#merge_requests_events}
	MergeRequestsEvents interface{} `field:"optional" json:"mergeRequestsEvents" yaml:"mergeRequestsEvents"`
	// Enable notifications for note events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/integration_jira#note_events IntegrationJira#note_events}
	NoteEvents interface{} `field:"optional" json:"noteEvents" yaml:"noteEvents"`
	// Enable notifications for pipeline events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/integration_jira#pipeline_events IntegrationJira#pipeline_events}
	PipelineEvents interface{} `field:"optional" json:"pipelineEvents" yaml:"pipelineEvents"`
	// The short identifier for your JIRA project, all uppercase, e.g., PROJ.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/integration_jira#project_key IntegrationJira#project_key}
	ProjectKey *string `field:"optional" json:"projectKey" yaml:"projectKey"`
	// Enable notifications for push events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/integration_jira#push_events IntegrationJira#push_events}
	PushEvents interface{} `field:"optional" json:"pushEvents" yaml:"pushEvents"`
	// Enable notifications for tag_push events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/integration_jira#tag_push_events IntegrationJira#tag_push_events}
	TagPushEvents interface{} `field:"optional" json:"tagPushEvents" yaml:"tagPushEvents"`
}

