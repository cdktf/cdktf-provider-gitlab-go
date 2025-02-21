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
	// The Jira API token, password, or personal access token to be used with Jira.
	//
	// When your authentication method is basic (jira_auth_type is 0), use an API token for Jira Cloud or a password for Jira Data Center or Jira Server. When your authentication method is a Jira personal access token (jira_auth_type is 1), use the personal access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/resources/integration_jira#password IntegrationJira#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// ID of the project you want to activate integration on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/resources/integration_jira#project IntegrationJira#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/resources/integration_jira#url IntegrationJira#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// The base URL to the Jira instance API. Web URL value is used if not set. For example, https://jira-api.example.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/resources/integration_jira#api_url IntegrationJira#api_url}
	ApiUrl *string `field:"optional" json:"apiUrl" yaml:"apiUrl"`
	// Enable comments inside Jira issues on each GitLab event (commit / merge request).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/resources/integration_jira#comment_on_event_enabled IntegrationJira#comment_on_event_enabled}
	CommentOnEventEnabled interface{} `field:"optional" json:"commentOnEventEnabled" yaml:"commentOnEventEnabled"`
	// Enable notifications for commit events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/resources/integration_jira#commit_events IntegrationJira#commit_events}
	CommitEvents interface{} `field:"optional" json:"commitEvents" yaml:"commitEvents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/resources/integration_jira#id IntegrationJira#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Enable viewing Jira issues in GitLab.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/resources/integration_jira#issues_enabled IntegrationJira#issues_enabled}
	IssuesEnabled interface{} `field:"optional" json:"issuesEnabled" yaml:"issuesEnabled"`
	// The authentication method to be used with Jira.
	//
	// 0 means Basic Authentication. 1 means Jira personal access token. Defaults to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/resources/integration_jira#jira_auth_type IntegrationJira#jira_auth_type}
	JiraAuthType *float64 `field:"optional" json:"jiraAuthType" yaml:"jiraAuthType"`
	// Prefix to match Jira issue keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/resources/integration_jira#jira_issue_prefix IntegrationJira#jira_issue_prefix}
	JiraIssuePrefix *string `field:"optional" json:"jiraIssuePrefix" yaml:"jiraIssuePrefix"`
	// Regular expression to match Jira issue keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/resources/integration_jira#jira_issue_regex IntegrationJira#jira_issue_regex}
	JiraIssueRegex *string `field:"optional" json:"jiraIssueRegex" yaml:"jiraIssueRegex"`
	// Enable automatic issue transitions.
	//
	// Takes precedence over jira_issue_transition_id if enabled. Defaults to false. This value cannot be imported, and will not perform drift detection if changed outside Terraform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/resources/integration_jira#jira_issue_transition_automatic IntegrationJira#jira_issue_transition_automatic}
	JiraIssueTransitionAutomatic interface{} `field:"optional" json:"jiraIssueTransitionAutomatic" yaml:"jiraIssueTransitionAutomatic"`
	// The ID of a transition that moves issues to a closed state.
	//
	// You can find this number under the JIRA workflow administration (Administration > Issues > Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/resources/integration_jira#jira_issue_transition_id IntegrationJira#jira_issue_transition_id}
	JiraIssueTransitionId *string `field:"optional" json:"jiraIssueTransitionId" yaml:"jiraIssueTransitionId"`
	// Enable notifications for merge request events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/resources/integration_jira#merge_requests_events IntegrationJira#merge_requests_events}
	MergeRequestsEvents interface{} `field:"optional" json:"mergeRequestsEvents" yaml:"mergeRequestsEvents"`
	// The short identifier for your JIRA project, all uppercase, e.g., PROJ.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/resources/integration_jira#project_key IntegrationJira#project_key}
	ProjectKey *string `field:"optional" json:"projectKey" yaml:"projectKey"`
	// Keys of Jira projects.
	//
	// When issues_enabled is true, this setting specifies which Jira projects to view issues from in GitLab.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/resources/integration_jira#project_keys IntegrationJira#project_keys}
	ProjectKeys *[]*string `field:"optional" json:"projectKeys" yaml:"projectKeys"`
	// Indicates whether or not to inherit default settings. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/resources/integration_jira#use_inherited_settings IntegrationJira#use_inherited_settings}
	UseInheritedSettings interface{} `field:"optional" json:"useInheritedSettings" yaml:"useInheritedSettings"`
	// The email or username to be used with Jira.
	//
	// For Jira Cloud use an email, for Jira Data Center and Jira Server use a username. Required when using Basic authentication (jira_auth_type is 0).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/resources/integration_jira#username IntegrationJira#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

