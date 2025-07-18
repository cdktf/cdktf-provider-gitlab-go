// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectintegrationmattermost

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v15/projectintegrationmattermost/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_integration_mattermost gitlab_project_integration_mattermost}.
type ProjectIntegrationMattermost interface {
	cdktf.TerraformResource
	BranchesToBeNotified() *string
	SetBranchesToBeNotified(val *string)
	BranchesToBeNotifiedInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConfidentialIssueChannel() *string
	SetConfidentialIssueChannel(val *string)
	ConfidentialIssueChannelInput() *string
	ConfidentialIssuesEvents() interface{}
	SetConfidentialIssuesEvents(val interface{})
	ConfidentialIssuesEventsInput() interface{}
	ConfidentialNoteChannel() *string
	SetConfidentialNoteChannel(val *string)
	ConfidentialNoteChannelInput() *string
	ConfidentialNoteEvents() interface{}
	SetConfidentialNoteEvents(val interface{})
	ConfidentialNoteEventsInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IssueChannel() *string
	SetIssueChannel(val *string)
	IssueChannelInput() *string
	IssuesEvents() interface{}
	SetIssuesEvents(val interface{})
	IssuesEventsInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MergeRequestChannel() *string
	SetMergeRequestChannel(val *string)
	MergeRequestChannelInput() *string
	MergeRequestsEvents() interface{}
	SetMergeRequestsEvents(val interface{})
	MergeRequestsEventsInput() interface{}
	// The tree node.
	Node() constructs.Node
	NoteChannel() *string
	SetNoteChannel(val *string)
	NoteChannelInput() *string
	NoteEvents() interface{}
	SetNoteEvents(val interface{})
	NoteEventsInput() interface{}
	NotifyOnlyBrokenPipelines() interface{}
	SetNotifyOnlyBrokenPipelines(val interface{})
	NotifyOnlyBrokenPipelinesInput() interface{}
	PipelineChannel() *string
	SetPipelineChannel(val *string)
	PipelineChannelInput() *string
	PipelineEvents() interface{}
	SetPipelineEvents(val interface{})
	PipelineEventsInput() interface{}
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PushChannel() *string
	SetPushChannel(val *string)
	PushChannelInput() *string
	PushEvents() interface{}
	SetPushEvents(val interface{})
	PushEventsInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	TagPushChannel() *string
	SetTagPushChannel(val *string)
	TagPushChannelInput() *string
	TagPushEvents() interface{}
	SetTagPushEvents(val interface{})
	TagPushEventsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	Webhook() *string
	SetWebhook(val *string)
	WebhookInput() *string
	WikiPageChannel() *string
	SetWikiPageChannel(val *string)
	WikiPageChannelInput() *string
	WikiPageEvents() interface{}
	SetWikiPageEvents(val interface{})
	WikiPageEventsInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetBranchesToBeNotified()
	ResetConfidentialIssueChannel()
	ResetConfidentialIssuesEvents()
	ResetConfidentialNoteChannel()
	ResetConfidentialNoteEvents()
	ResetId()
	ResetIssueChannel()
	ResetIssuesEvents()
	ResetMergeRequestChannel()
	ResetMergeRequestsEvents()
	ResetNoteChannel()
	ResetNoteEvents()
	ResetNotifyOnlyBrokenPipelines()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPipelineChannel()
	ResetPipelineEvents()
	ResetPushChannel()
	ResetPushEvents()
	ResetTagPushChannel()
	ResetTagPushEvents()
	ResetUsername()
	ResetWikiPageChannel()
	ResetWikiPageEvents()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ProjectIntegrationMattermost
type jsiiProxy_ProjectIntegrationMattermost struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ProjectIntegrationMattermost) BranchesToBeNotified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchesToBeNotified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) BranchesToBeNotifiedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchesToBeNotifiedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) ConfidentialIssueChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confidentialIssueChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) ConfidentialIssueChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confidentialIssueChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) ConfidentialIssuesEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidentialIssuesEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) ConfidentialIssuesEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidentialIssuesEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) ConfidentialNoteChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confidentialNoteChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) ConfidentialNoteChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confidentialNoteChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) ConfidentialNoteEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidentialNoteEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) ConfidentialNoteEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidentialNoteEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) IssueChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issueChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) IssueChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issueChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) IssuesEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issuesEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) IssuesEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issuesEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) MergeRequestChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeRequestChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) MergeRequestChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeRequestChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) MergeRequestsEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeRequestsEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) MergeRequestsEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeRequestsEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) NoteChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noteChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) NoteChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noteChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) NoteEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) NoteEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) NotifyOnlyBrokenPipelines() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyOnlyBrokenPipelines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) NotifyOnlyBrokenPipelinesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyOnlyBrokenPipelinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) PipelineChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) PipelineChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) PipelineEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pipelineEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) PipelineEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pipelineEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) PushChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pushChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) PushChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pushChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) PushEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pushEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) PushEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pushEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) TagPushChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagPushChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) TagPushChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagPushChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) TagPushEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagPushEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) TagPushEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagPushEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) Webhook() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) WebhookInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) WikiPageChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wikiPageChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) WikiPageChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wikiPageChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) WikiPageEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wikiPageEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationMattermost) WikiPageEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wikiPageEventsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_integration_mattermost gitlab_project_integration_mattermost} Resource.
func NewProjectIntegrationMattermost(scope constructs.Construct, id *string, config *ProjectIntegrationMattermostConfig) ProjectIntegrationMattermost {
	_init_.Initialize()

	if err := validateNewProjectIntegrationMattermostParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProjectIntegrationMattermost{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.projectIntegrationMattermost.ProjectIntegrationMattermost",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_integration_mattermost gitlab_project_integration_mattermost} Resource.
func NewProjectIntegrationMattermost_Override(p ProjectIntegrationMattermost, scope constructs.Construct, id *string, config *ProjectIntegrationMattermostConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.projectIntegrationMattermost.ProjectIntegrationMattermost",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetBranchesToBeNotified(val *string) {
	if err := j.validateSetBranchesToBeNotifiedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branchesToBeNotified",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetConfidentialIssueChannel(val *string) {
	if err := j.validateSetConfidentialIssueChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidentialIssueChannel",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetConfidentialIssuesEvents(val interface{}) {
	if err := j.validateSetConfidentialIssuesEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidentialIssuesEvents",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetConfidentialNoteChannel(val *string) {
	if err := j.validateSetConfidentialNoteChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidentialNoteChannel",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetConfidentialNoteEvents(val interface{}) {
	if err := j.validateSetConfidentialNoteEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidentialNoteEvents",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetIssueChannel(val *string) {
	if err := j.validateSetIssueChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issueChannel",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetIssuesEvents(val interface{}) {
	if err := j.validateSetIssuesEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuesEvents",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetMergeRequestChannel(val *string) {
	if err := j.validateSetMergeRequestChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeRequestChannel",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetMergeRequestsEvents(val interface{}) {
	if err := j.validateSetMergeRequestsEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeRequestsEvents",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetNoteChannel(val *string) {
	if err := j.validateSetNoteChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noteChannel",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetNoteEvents(val interface{}) {
	if err := j.validateSetNoteEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noteEvents",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetNotifyOnlyBrokenPipelines(val interface{}) {
	if err := j.validateSetNotifyOnlyBrokenPipelinesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifyOnlyBrokenPipelines",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetPipelineChannel(val *string) {
	if err := j.validateSetPipelineChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipelineChannel",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetPipelineEvents(val interface{}) {
	if err := j.validateSetPipelineEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipelineEvents",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetPushChannel(val *string) {
	if err := j.validateSetPushChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pushChannel",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetPushEvents(val interface{}) {
	if err := j.validateSetPushEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pushEvents",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetTagPushChannel(val *string) {
	if err := j.validateSetTagPushChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagPushChannel",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetTagPushEvents(val interface{}) {
	if err := j.validateSetTagPushEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagPushEvents",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetWebhook(val *string) {
	if err := j.validateSetWebhookParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webhook",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetWikiPageChannel(val *string) {
	if err := j.validateSetWikiPageChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wikiPageChannel",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationMattermost)SetWikiPageEvents(val interface{}) {
	if err := j.validateSetWikiPageEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wikiPageEvents",
		val,
	)
}

// Generates CDKTF code for importing a ProjectIntegrationMattermost resource upon running "cdktf plan <stack-name>".
func ProjectIntegrationMattermost_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateProjectIntegrationMattermost_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.projectIntegrationMattermost.ProjectIntegrationMattermost",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ProjectIntegrationMattermost_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjectIntegrationMattermost_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.projectIntegrationMattermost.ProjectIntegrationMattermost",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ProjectIntegrationMattermost_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjectIntegrationMattermost_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.projectIntegrationMattermost.ProjectIntegrationMattermost",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ProjectIntegrationMattermost_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjectIntegrationMattermost_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.projectIntegrationMattermost.ProjectIntegrationMattermost",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ProjectIntegrationMattermost_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.projectIntegrationMattermost.ProjectIntegrationMattermost",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_ProjectIntegrationMattermost) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationMattermost) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationMattermost) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationMattermost) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationMattermost) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationMattermost) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationMattermost) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationMattermost) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationMattermost) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationMattermost) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationMattermost) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ResetBranchesToBeNotified() {
	_jsii_.InvokeVoid(
		p,
		"resetBranchesToBeNotified",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ResetConfidentialIssueChannel() {
	_jsii_.InvokeVoid(
		p,
		"resetConfidentialIssueChannel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ResetConfidentialIssuesEvents() {
	_jsii_.InvokeVoid(
		p,
		"resetConfidentialIssuesEvents",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ResetConfidentialNoteChannel() {
	_jsii_.InvokeVoid(
		p,
		"resetConfidentialNoteChannel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ResetConfidentialNoteEvents() {
	_jsii_.InvokeVoid(
		p,
		"resetConfidentialNoteEvents",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ResetIssueChannel() {
	_jsii_.InvokeVoid(
		p,
		"resetIssueChannel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ResetIssuesEvents() {
	_jsii_.InvokeVoid(
		p,
		"resetIssuesEvents",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ResetMergeRequestChannel() {
	_jsii_.InvokeVoid(
		p,
		"resetMergeRequestChannel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ResetMergeRequestsEvents() {
	_jsii_.InvokeVoid(
		p,
		"resetMergeRequestsEvents",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ResetNoteChannel() {
	_jsii_.InvokeVoid(
		p,
		"resetNoteChannel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ResetNoteEvents() {
	_jsii_.InvokeVoid(
		p,
		"resetNoteEvents",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ResetNotifyOnlyBrokenPipelines() {
	_jsii_.InvokeVoid(
		p,
		"resetNotifyOnlyBrokenPipelines",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ResetPipelineChannel() {
	_jsii_.InvokeVoid(
		p,
		"resetPipelineChannel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ResetPipelineEvents() {
	_jsii_.InvokeVoid(
		p,
		"resetPipelineEvents",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ResetPushChannel() {
	_jsii_.InvokeVoid(
		p,
		"resetPushChannel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ResetPushEvents() {
	_jsii_.InvokeVoid(
		p,
		"resetPushEvents",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ResetTagPushChannel() {
	_jsii_.InvokeVoid(
		p,
		"resetTagPushChannel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ResetTagPushEvents() {
	_jsii_.InvokeVoid(
		p,
		"resetTagPushEvents",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ResetUsername() {
	_jsii_.InvokeVoid(
		p,
		"resetUsername",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ResetWikiPageChannel() {
	_jsii_.InvokeVoid(
		p,
		"resetWikiPageChannel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ResetWikiPageEvents() {
	_jsii_.InvokeVoid(
		p,
		"resetWikiPageEvents",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationMattermost) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationMattermost) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationMattermost) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

