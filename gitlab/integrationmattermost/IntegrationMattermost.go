// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationmattermost

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v12/integrationmattermost/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.1/docs/resources/integration_mattermost gitlab_integration_mattermost}.
type IntegrationMattermost interface {
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

// The jsii proxy struct for IntegrationMattermost
type jsiiProxy_IntegrationMattermost struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IntegrationMattermost) BranchesToBeNotified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchesToBeNotified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) BranchesToBeNotifiedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchesToBeNotifiedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) ConfidentialIssueChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confidentialIssueChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) ConfidentialIssueChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confidentialIssueChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) ConfidentialIssuesEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidentialIssuesEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) ConfidentialIssuesEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidentialIssuesEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) ConfidentialNoteChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confidentialNoteChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) ConfidentialNoteChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confidentialNoteChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) ConfidentialNoteEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidentialNoteEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) ConfidentialNoteEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidentialNoteEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) IssueChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issueChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) IssueChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issueChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) IssuesEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issuesEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) IssuesEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issuesEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) MergeRequestChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeRequestChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) MergeRequestChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeRequestChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) MergeRequestsEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeRequestsEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) MergeRequestsEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeRequestsEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) NoteChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noteChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) NoteChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noteChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) NoteEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) NoteEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) NotifyOnlyBrokenPipelines() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyOnlyBrokenPipelines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) NotifyOnlyBrokenPipelinesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyOnlyBrokenPipelinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) PipelineChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) PipelineChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) PipelineEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pipelineEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) PipelineEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pipelineEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) PushChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pushChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) PushChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pushChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) PushEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pushEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) PushEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pushEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) TagPushChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagPushChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) TagPushChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagPushChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) TagPushEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagPushEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) TagPushEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagPushEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) Webhook() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) WebhookInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) WikiPageChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wikiPageChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) WikiPageChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wikiPageChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) WikiPageEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wikiPageEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationMattermost) WikiPageEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wikiPageEventsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.1/docs/resources/integration_mattermost gitlab_integration_mattermost} Resource.
func NewIntegrationMattermost(scope constructs.Construct, id *string, config *IntegrationMattermostConfig) IntegrationMattermost {
	_init_.Initialize()

	if err := validateNewIntegrationMattermostParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationMattermost{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.integrationMattermost.IntegrationMattermost",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.1/docs/resources/integration_mattermost gitlab_integration_mattermost} Resource.
func NewIntegrationMattermost_Override(i IntegrationMattermost, scope constructs.Construct, id *string, config *IntegrationMattermostConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.integrationMattermost.IntegrationMattermost",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetBranchesToBeNotified(val *string) {
	if err := j.validateSetBranchesToBeNotifiedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branchesToBeNotified",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetConfidentialIssueChannel(val *string) {
	if err := j.validateSetConfidentialIssueChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidentialIssueChannel",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetConfidentialIssuesEvents(val interface{}) {
	if err := j.validateSetConfidentialIssuesEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidentialIssuesEvents",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetConfidentialNoteChannel(val *string) {
	if err := j.validateSetConfidentialNoteChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidentialNoteChannel",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetConfidentialNoteEvents(val interface{}) {
	if err := j.validateSetConfidentialNoteEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidentialNoteEvents",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetIssueChannel(val *string) {
	if err := j.validateSetIssueChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issueChannel",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetIssuesEvents(val interface{}) {
	if err := j.validateSetIssuesEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuesEvents",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetMergeRequestChannel(val *string) {
	if err := j.validateSetMergeRequestChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeRequestChannel",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetMergeRequestsEvents(val interface{}) {
	if err := j.validateSetMergeRequestsEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeRequestsEvents",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetNoteChannel(val *string) {
	if err := j.validateSetNoteChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noteChannel",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetNoteEvents(val interface{}) {
	if err := j.validateSetNoteEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noteEvents",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetNotifyOnlyBrokenPipelines(val interface{}) {
	if err := j.validateSetNotifyOnlyBrokenPipelinesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifyOnlyBrokenPipelines",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetPipelineChannel(val *string) {
	if err := j.validateSetPipelineChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipelineChannel",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetPipelineEvents(val interface{}) {
	if err := j.validateSetPipelineEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipelineEvents",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetPushChannel(val *string) {
	if err := j.validateSetPushChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pushChannel",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetPushEvents(val interface{}) {
	if err := j.validateSetPushEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pushEvents",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetTagPushChannel(val *string) {
	if err := j.validateSetTagPushChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagPushChannel",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetTagPushEvents(val interface{}) {
	if err := j.validateSetTagPushEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagPushEvents",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetWebhook(val *string) {
	if err := j.validateSetWebhookParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webhook",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetWikiPageChannel(val *string) {
	if err := j.validateSetWikiPageChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wikiPageChannel",
		val,
	)
}

func (j *jsiiProxy_IntegrationMattermost)SetWikiPageEvents(val interface{}) {
	if err := j.validateSetWikiPageEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wikiPageEvents",
		val,
	)
}

// Generates CDKTF code for importing a IntegrationMattermost resource upon running "cdktf plan <stack-name>".
func IntegrationMattermost_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIntegrationMattermost_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.integrationMattermost.IntegrationMattermost",
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
func IntegrationMattermost_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationMattermost_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.integrationMattermost.IntegrationMattermost",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IntegrationMattermost_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationMattermost_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.integrationMattermost.IntegrationMattermost",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IntegrationMattermost_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationMattermost_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.integrationMattermost.IntegrationMattermost",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IntegrationMattermost_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.integrationMattermost.IntegrationMattermost",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IntegrationMattermost) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IntegrationMattermost) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IntegrationMattermost) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationMattermost) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationMattermost) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationMattermost) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationMattermost) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationMattermost) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationMattermost) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationMattermost) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationMattermost) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationMattermost) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationMattermost) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IntegrationMattermost) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationMattermost) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IntegrationMattermost) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IntegrationMattermost) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IntegrationMattermost) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IntegrationMattermost) ResetBranchesToBeNotified() {
	_jsii_.InvokeVoid(
		i,
		"resetBranchesToBeNotified",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationMattermost) ResetConfidentialIssueChannel() {
	_jsii_.InvokeVoid(
		i,
		"resetConfidentialIssueChannel",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationMattermost) ResetConfidentialIssuesEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetConfidentialIssuesEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationMattermost) ResetConfidentialNoteChannel() {
	_jsii_.InvokeVoid(
		i,
		"resetConfidentialNoteChannel",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationMattermost) ResetConfidentialNoteEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetConfidentialNoteEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationMattermost) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationMattermost) ResetIssueChannel() {
	_jsii_.InvokeVoid(
		i,
		"resetIssueChannel",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationMattermost) ResetIssuesEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetIssuesEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationMattermost) ResetMergeRequestChannel() {
	_jsii_.InvokeVoid(
		i,
		"resetMergeRequestChannel",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationMattermost) ResetMergeRequestsEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetMergeRequestsEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationMattermost) ResetNoteChannel() {
	_jsii_.InvokeVoid(
		i,
		"resetNoteChannel",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationMattermost) ResetNoteEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetNoteEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationMattermost) ResetNotifyOnlyBrokenPipelines() {
	_jsii_.InvokeVoid(
		i,
		"resetNotifyOnlyBrokenPipelines",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationMattermost) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationMattermost) ResetPipelineChannel() {
	_jsii_.InvokeVoid(
		i,
		"resetPipelineChannel",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationMattermost) ResetPipelineEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetPipelineEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationMattermost) ResetPushChannel() {
	_jsii_.InvokeVoid(
		i,
		"resetPushChannel",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationMattermost) ResetPushEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetPushEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationMattermost) ResetTagPushChannel() {
	_jsii_.InvokeVoid(
		i,
		"resetTagPushChannel",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationMattermost) ResetTagPushEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetTagPushEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationMattermost) ResetUsername() {
	_jsii_.InvokeVoid(
		i,
		"resetUsername",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationMattermost) ResetWikiPageChannel() {
	_jsii_.InvokeVoid(
		i,
		"resetWikiPageChannel",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationMattermost) ResetWikiPageEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetWikiPageEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationMattermost) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationMattermost) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationMattermost) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationMattermost) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationMattermost) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationMattermost) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

