// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grouphook

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupHookConfig struct {
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
	// The ID or full path of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/resources/group_hook#group GroupHook#group}
	Group *string `field:"required" json:"group" yaml:"group"`
	// The url of the hook to invoke.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/resources/group_hook#url GroupHook#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// Invoke the hook for confidential issues events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/resources/group_hook#confidential_issues_events GroupHook#confidential_issues_events}
	ConfidentialIssuesEvents interface{} `field:"optional" json:"confidentialIssuesEvents" yaml:"confidentialIssuesEvents"`
	// Invoke the hook for confidential notes events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/resources/group_hook#confidential_note_events GroupHook#confidential_note_events}
	ConfidentialNoteEvents interface{} `field:"optional" json:"confidentialNoteEvents" yaml:"confidentialNoteEvents"`
	// Set a custom webhook template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/resources/group_hook#custom_webhook_template GroupHook#custom_webhook_template}
	CustomWebhookTemplate *string `field:"optional" json:"customWebhookTemplate" yaml:"customWebhookTemplate"`
	// Invoke the hook for deployment events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/resources/group_hook#deployment_events GroupHook#deployment_events}
	DeploymentEvents interface{} `field:"optional" json:"deploymentEvents" yaml:"deploymentEvents"`
	// Enable ssl verification when invoking the hook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/resources/group_hook#enable_ssl_verification GroupHook#enable_ssl_verification}
	EnableSslVerification interface{} `field:"optional" json:"enableSslVerification" yaml:"enableSslVerification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/resources/group_hook#id GroupHook#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Invoke the hook for issues events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/resources/group_hook#issues_events GroupHook#issues_events}
	IssuesEvents interface{} `field:"optional" json:"issuesEvents" yaml:"issuesEvents"`
	// Invoke the hook for job events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/resources/group_hook#job_events GroupHook#job_events}
	JobEvents interface{} `field:"optional" json:"jobEvents" yaml:"jobEvents"`
	// Invoke the hook for merge requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/resources/group_hook#merge_requests_events GroupHook#merge_requests_events}
	MergeRequestsEvents interface{} `field:"optional" json:"mergeRequestsEvents" yaml:"mergeRequestsEvents"`
	// Invoke the hook for notes events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/resources/group_hook#note_events GroupHook#note_events}
	NoteEvents interface{} `field:"optional" json:"noteEvents" yaml:"noteEvents"`
	// Invoke the hook for pipeline events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/resources/group_hook#pipeline_events GroupHook#pipeline_events}
	PipelineEvents interface{} `field:"optional" json:"pipelineEvents" yaml:"pipelineEvents"`
	// Invoke the hook for push events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/resources/group_hook#push_events GroupHook#push_events}
	PushEvents interface{} `field:"optional" json:"pushEvents" yaml:"pushEvents"`
	// Invoke the hook for push events on matching branches only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/resources/group_hook#push_events_branch_filter GroupHook#push_events_branch_filter}
	PushEventsBranchFilter *string `field:"optional" json:"pushEventsBranchFilter" yaml:"pushEventsBranchFilter"`
	// Invoke the hook for releases events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/resources/group_hook#releases_events GroupHook#releases_events}
	ReleasesEvents interface{} `field:"optional" json:"releasesEvents" yaml:"releasesEvents"`
	// Invoke the hook for subgroup events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/resources/group_hook#subgroup_events GroupHook#subgroup_events}
	SubgroupEvents interface{} `field:"optional" json:"subgroupEvents" yaml:"subgroupEvents"`
	// Invoke the hook for tag push events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/resources/group_hook#tag_push_events GroupHook#tag_push_events}
	TagPushEvents interface{} `field:"optional" json:"tagPushEvents" yaml:"tagPushEvents"`
	// A token to present when invoking the hook. The token is not available for imported resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/resources/group_hook#token GroupHook#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Invoke the hook for wiki page events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/resources/group_hook#wiki_page_events GroupHook#wiki_page_events}
	WikiPageEvents interface{} `field:"optional" json:"wikiPageEvents" yaml:"wikiPageEvents"`
}

