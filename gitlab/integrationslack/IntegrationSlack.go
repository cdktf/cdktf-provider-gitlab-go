// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationslack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v13/integrationslack/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.7.1/docs/resources/integration_slack gitlab_integration_slack}.
type IntegrationSlack interface {
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
	JobEvents() cdktf.IResolvable
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
	NotifyOnlyDefaultBranch() interface{}
	SetNotifyOnlyDefaultBranch(val interface{})
	NotifyOnlyDefaultBranchInput() interface{}
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
	ResetNotifyOnlyDefaultBranch()
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

// The jsii proxy struct for IntegrationSlack
type jsiiProxy_IntegrationSlack struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IntegrationSlack) BranchesToBeNotified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchesToBeNotified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) BranchesToBeNotifiedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchesToBeNotifiedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) ConfidentialIssueChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confidentialIssueChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) ConfidentialIssueChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confidentialIssueChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) ConfidentialIssuesEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidentialIssuesEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) ConfidentialIssuesEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidentialIssuesEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) ConfidentialNoteChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confidentialNoteChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) ConfidentialNoteChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confidentialNoteChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) ConfidentialNoteEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidentialNoteEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) ConfidentialNoteEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidentialNoteEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) IssueChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issueChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) IssueChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issueChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) IssuesEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issuesEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) IssuesEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issuesEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) JobEvents() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"jobEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) MergeRequestChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeRequestChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) MergeRequestChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeRequestChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) MergeRequestsEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeRequestsEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) MergeRequestsEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeRequestsEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) NoteChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noteChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) NoteChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noteChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) NoteEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) NoteEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) NotifyOnlyBrokenPipelines() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyOnlyBrokenPipelines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) NotifyOnlyBrokenPipelinesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyOnlyBrokenPipelinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) NotifyOnlyDefaultBranch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyOnlyDefaultBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) NotifyOnlyDefaultBranchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyOnlyDefaultBranchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) PipelineChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) PipelineChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) PipelineEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pipelineEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) PipelineEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pipelineEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) PushChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pushChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) PushChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pushChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) PushEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pushEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) PushEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pushEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) TagPushChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagPushChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) TagPushChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagPushChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) TagPushEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagPushEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) TagPushEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagPushEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) Webhook() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) WebhookInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) WikiPageChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wikiPageChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) WikiPageChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wikiPageChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) WikiPageEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wikiPageEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationSlack) WikiPageEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wikiPageEventsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.7.1/docs/resources/integration_slack gitlab_integration_slack} Resource.
func NewIntegrationSlack(scope constructs.Construct, id *string, config *IntegrationSlackConfig) IntegrationSlack {
	_init_.Initialize()

	if err := validateNewIntegrationSlackParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationSlack{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.integrationSlack.IntegrationSlack",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.7.1/docs/resources/integration_slack gitlab_integration_slack} Resource.
func NewIntegrationSlack_Override(i IntegrationSlack, scope constructs.Construct, id *string, config *IntegrationSlackConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.integrationSlack.IntegrationSlack",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetBranchesToBeNotified(val *string) {
	if err := j.validateSetBranchesToBeNotifiedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branchesToBeNotified",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetConfidentialIssueChannel(val *string) {
	if err := j.validateSetConfidentialIssueChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidentialIssueChannel",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetConfidentialIssuesEvents(val interface{}) {
	if err := j.validateSetConfidentialIssuesEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidentialIssuesEvents",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetConfidentialNoteChannel(val *string) {
	if err := j.validateSetConfidentialNoteChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidentialNoteChannel",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetConfidentialNoteEvents(val interface{}) {
	if err := j.validateSetConfidentialNoteEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidentialNoteEvents",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetIssueChannel(val *string) {
	if err := j.validateSetIssueChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issueChannel",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetIssuesEvents(val interface{}) {
	if err := j.validateSetIssuesEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuesEvents",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetMergeRequestChannel(val *string) {
	if err := j.validateSetMergeRequestChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeRequestChannel",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetMergeRequestsEvents(val interface{}) {
	if err := j.validateSetMergeRequestsEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeRequestsEvents",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetNoteChannel(val *string) {
	if err := j.validateSetNoteChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noteChannel",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetNoteEvents(val interface{}) {
	if err := j.validateSetNoteEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noteEvents",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetNotifyOnlyBrokenPipelines(val interface{}) {
	if err := j.validateSetNotifyOnlyBrokenPipelinesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifyOnlyBrokenPipelines",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetNotifyOnlyDefaultBranch(val interface{}) {
	if err := j.validateSetNotifyOnlyDefaultBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifyOnlyDefaultBranch",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetPipelineChannel(val *string) {
	if err := j.validateSetPipelineChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipelineChannel",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetPipelineEvents(val interface{}) {
	if err := j.validateSetPipelineEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipelineEvents",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetPushChannel(val *string) {
	if err := j.validateSetPushChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pushChannel",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetPushEvents(val interface{}) {
	if err := j.validateSetPushEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pushEvents",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetTagPushChannel(val *string) {
	if err := j.validateSetTagPushChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagPushChannel",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetTagPushEvents(val interface{}) {
	if err := j.validateSetTagPushEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagPushEvents",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetWebhook(val *string) {
	if err := j.validateSetWebhookParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webhook",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetWikiPageChannel(val *string) {
	if err := j.validateSetWikiPageChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wikiPageChannel",
		val,
	)
}

func (j *jsiiProxy_IntegrationSlack)SetWikiPageEvents(val interface{}) {
	if err := j.validateSetWikiPageEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wikiPageEvents",
		val,
	)
}

// Generates CDKTF code for importing a IntegrationSlack resource upon running "cdktf plan <stack-name>".
func IntegrationSlack_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIntegrationSlack_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.integrationSlack.IntegrationSlack",
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
func IntegrationSlack_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationSlack_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.integrationSlack.IntegrationSlack",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IntegrationSlack_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationSlack_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.integrationSlack.IntegrationSlack",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IntegrationSlack_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationSlack_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.integrationSlack.IntegrationSlack",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IntegrationSlack_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.integrationSlack.IntegrationSlack",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IntegrationSlack) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IntegrationSlack) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IntegrationSlack) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IntegrationSlack) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IntegrationSlack) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IntegrationSlack) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IntegrationSlack) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IntegrationSlack) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IntegrationSlack) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IntegrationSlack) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IntegrationSlack) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IntegrationSlack) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationSlack) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IntegrationSlack) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IntegrationSlack) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IntegrationSlack) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IntegrationSlack) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IntegrationSlack) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IntegrationSlack) ResetBranchesToBeNotified() {
	_jsii_.InvokeVoid(
		i,
		"resetBranchesToBeNotified",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationSlack) ResetConfidentialIssueChannel() {
	_jsii_.InvokeVoid(
		i,
		"resetConfidentialIssueChannel",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationSlack) ResetConfidentialIssuesEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetConfidentialIssuesEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationSlack) ResetConfidentialNoteChannel() {
	_jsii_.InvokeVoid(
		i,
		"resetConfidentialNoteChannel",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationSlack) ResetConfidentialNoteEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetConfidentialNoteEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationSlack) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationSlack) ResetIssueChannel() {
	_jsii_.InvokeVoid(
		i,
		"resetIssueChannel",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationSlack) ResetIssuesEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetIssuesEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationSlack) ResetMergeRequestChannel() {
	_jsii_.InvokeVoid(
		i,
		"resetMergeRequestChannel",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationSlack) ResetMergeRequestsEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetMergeRequestsEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationSlack) ResetNoteChannel() {
	_jsii_.InvokeVoid(
		i,
		"resetNoteChannel",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationSlack) ResetNoteEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetNoteEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationSlack) ResetNotifyOnlyBrokenPipelines() {
	_jsii_.InvokeVoid(
		i,
		"resetNotifyOnlyBrokenPipelines",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationSlack) ResetNotifyOnlyDefaultBranch() {
	_jsii_.InvokeVoid(
		i,
		"resetNotifyOnlyDefaultBranch",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationSlack) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationSlack) ResetPipelineChannel() {
	_jsii_.InvokeVoid(
		i,
		"resetPipelineChannel",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationSlack) ResetPipelineEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetPipelineEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationSlack) ResetPushChannel() {
	_jsii_.InvokeVoid(
		i,
		"resetPushChannel",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationSlack) ResetPushEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetPushEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationSlack) ResetTagPushChannel() {
	_jsii_.InvokeVoid(
		i,
		"resetTagPushChannel",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationSlack) ResetTagPushEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetTagPushEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationSlack) ResetUsername() {
	_jsii_.InvokeVoid(
		i,
		"resetUsername",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationSlack) ResetWikiPageChannel() {
	_jsii_.InvokeVoid(
		i,
		"resetWikiPageChannel",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationSlack) ResetWikiPageEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetWikiPageEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationSlack) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationSlack) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationSlack) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationSlack) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationSlack) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationSlack) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

