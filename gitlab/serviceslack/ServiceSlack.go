// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceslack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v13/serviceslack/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.0.0/docs/resources/service_slack gitlab_service_slack}.
type ServiceSlack interface {
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

// The jsii proxy struct for ServiceSlack
type jsiiProxy_ServiceSlack struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServiceSlack) BranchesToBeNotified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchesToBeNotified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) BranchesToBeNotifiedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchesToBeNotifiedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) ConfidentialIssueChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confidentialIssueChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) ConfidentialIssueChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confidentialIssueChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) ConfidentialIssuesEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidentialIssuesEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) ConfidentialIssuesEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidentialIssuesEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) ConfidentialNoteEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidentialNoteEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) ConfidentialNoteEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidentialNoteEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) IssueChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issueChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) IssueChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issueChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) IssuesEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issuesEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) IssuesEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issuesEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) JobEvents() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"jobEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) MergeRequestChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeRequestChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) MergeRequestChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeRequestChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) MergeRequestsEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeRequestsEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) MergeRequestsEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeRequestsEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) NoteChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noteChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) NoteChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noteChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) NoteEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) NoteEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) NotifyOnlyBrokenPipelines() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyOnlyBrokenPipelines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) NotifyOnlyBrokenPipelinesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyOnlyBrokenPipelinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) NotifyOnlyDefaultBranch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyOnlyDefaultBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) NotifyOnlyDefaultBranchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyOnlyDefaultBranchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) PipelineChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) PipelineChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) PipelineEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pipelineEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) PipelineEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pipelineEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) PushChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pushChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) PushChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pushChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) PushEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pushEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) PushEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pushEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) TagPushChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagPushChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) TagPushChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagPushChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) TagPushEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagPushEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) TagPushEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagPushEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) Webhook() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) WebhookInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) WikiPageChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wikiPageChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) WikiPageChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wikiPageChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) WikiPageEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wikiPageEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSlack) WikiPageEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wikiPageEventsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.0.0/docs/resources/service_slack gitlab_service_slack} Resource.
func NewServiceSlack(scope constructs.Construct, id *string, config *ServiceSlackConfig) ServiceSlack {
	_init_.Initialize()

	if err := validateNewServiceSlackParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceSlack{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.serviceSlack.ServiceSlack",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.0.0/docs/resources/service_slack gitlab_service_slack} Resource.
func NewServiceSlack_Override(s ServiceSlack, scope constructs.Construct, id *string, config *ServiceSlackConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.serviceSlack.ServiceSlack",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServiceSlack)SetBranchesToBeNotified(val *string) {
	if err := j.validateSetBranchesToBeNotifiedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branchesToBeNotified",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetConfidentialIssueChannel(val *string) {
	if err := j.validateSetConfidentialIssueChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidentialIssueChannel",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetConfidentialIssuesEvents(val interface{}) {
	if err := j.validateSetConfidentialIssuesEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidentialIssuesEvents",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetConfidentialNoteEvents(val interface{}) {
	if err := j.validateSetConfidentialNoteEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidentialNoteEvents",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetIssueChannel(val *string) {
	if err := j.validateSetIssueChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issueChannel",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetIssuesEvents(val interface{}) {
	if err := j.validateSetIssuesEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuesEvents",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetMergeRequestChannel(val *string) {
	if err := j.validateSetMergeRequestChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeRequestChannel",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetMergeRequestsEvents(val interface{}) {
	if err := j.validateSetMergeRequestsEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeRequestsEvents",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetNoteChannel(val *string) {
	if err := j.validateSetNoteChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noteChannel",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetNoteEvents(val interface{}) {
	if err := j.validateSetNoteEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noteEvents",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetNotifyOnlyBrokenPipelines(val interface{}) {
	if err := j.validateSetNotifyOnlyBrokenPipelinesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifyOnlyBrokenPipelines",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetNotifyOnlyDefaultBranch(val interface{}) {
	if err := j.validateSetNotifyOnlyDefaultBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifyOnlyDefaultBranch",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetPipelineChannel(val *string) {
	if err := j.validateSetPipelineChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipelineChannel",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetPipelineEvents(val interface{}) {
	if err := j.validateSetPipelineEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipelineEvents",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetPushChannel(val *string) {
	if err := j.validateSetPushChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pushChannel",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetPushEvents(val interface{}) {
	if err := j.validateSetPushEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pushEvents",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetTagPushChannel(val *string) {
	if err := j.validateSetTagPushChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagPushChannel",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetTagPushEvents(val interface{}) {
	if err := j.validateSetTagPushEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagPushEvents",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetWebhook(val *string) {
	if err := j.validateSetWebhookParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webhook",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetWikiPageChannel(val *string) {
	if err := j.validateSetWikiPageChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wikiPageChannel",
		val,
	)
}

func (j *jsiiProxy_ServiceSlack)SetWikiPageEvents(val interface{}) {
	if err := j.validateSetWikiPageEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wikiPageEvents",
		val,
	)
}

// Generates CDKTF code for importing a ServiceSlack resource upon running "cdktf plan <stack-name>".
func ServiceSlack_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateServiceSlack_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.serviceSlack.ServiceSlack",
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
func ServiceSlack_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServiceSlack_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.serviceSlack.ServiceSlack",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServiceSlack_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServiceSlack_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.serviceSlack.ServiceSlack",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServiceSlack_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServiceSlack_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.serviceSlack.ServiceSlack",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServiceSlack_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.serviceSlack.ServiceSlack",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_ServiceSlack) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_ServiceSlack) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_ServiceSlack) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceSlack) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceSlack) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceSlack) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceSlack) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceSlack) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceSlack) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceSlack) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceSlack) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceSlack) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceSlack) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_ServiceSlack) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceSlack) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_ServiceSlack) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_ServiceSlack) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_ServiceSlack) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServiceSlack) ResetBranchesToBeNotified() {
	_jsii_.InvokeVoid(
		s,
		"resetBranchesToBeNotified",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSlack) ResetConfidentialIssueChannel() {
	_jsii_.InvokeVoid(
		s,
		"resetConfidentialIssueChannel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSlack) ResetConfidentialIssuesEvents() {
	_jsii_.InvokeVoid(
		s,
		"resetConfidentialIssuesEvents",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSlack) ResetConfidentialNoteEvents() {
	_jsii_.InvokeVoid(
		s,
		"resetConfidentialNoteEvents",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSlack) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSlack) ResetIssueChannel() {
	_jsii_.InvokeVoid(
		s,
		"resetIssueChannel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSlack) ResetIssuesEvents() {
	_jsii_.InvokeVoid(
		s,
		"resetIssuesEvents",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSlack) ResetMergeRequestChannel() {
	_jsii_.InvokeVoid(
		s,
		"resetMergeRequestChannel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSlack) ResetMergeRequestsEvents() {
	_jsii_.InvokeVoid(
		s,
		"resetMergeRequestsEvents",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSlack) ResetNoteChannel() {
	_jsii_.InvokeVoid(
		s,
		"resetNoteChannel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSlack) ResetNoteEvents() {
	_jsii_.InvokeVoid(
		s,
		"resetNoteEvents",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSlack) ResetNotifyOnlyBrokenPipelines() {
	_jsii_.InvokeVoid(
		s,
		"resetNotifyOnlyBrokenPipelines",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSlack) ResetNotifyOnlyDefaultBranch() {
	_jsii_.InvokeVoid(
		s,
		"resetNotifyOnlyDefaultBranch",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSlack) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSlack) ResetPipelineChannel() {
	_jsii_.InvokeVoid(
		s,
		"resetPipelineChannel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSlack) ResetPipelineEvents() {
	_jsii_.InvokeVoid(
		s,
		"resetPipelineEvents",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSlack) ResetPushChannel() {
	_jsii_.InvokeVoid(
		s,
		"resetPushChannel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSlack) ResetPushEvents() {
	_jsii_.InvokeVoid(
		s,
		"resetPushEvents",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSlack) ResetTagPushChannel() {
	_jsii_.InvokeVoid(
		s,
		"resetTagPushChannel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSlack) ResetTagPushEvents() {
	_jsii_.InvokeVoid(
		s,
		"resetTagPushEvents",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSlack) ResetUsername() {
	_jsii_.InvokeVoid(
		s,
		"resetUsername",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSlack) ResetWikiPageChannel() {
	_jsii_.InvokeVoid(
		s,
		"resetWikiPageChannel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSlack) ResetWikiPageEvents() {
	_jsii_.InvokeVoid(
		s,
		"resetWikiPageEvents",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSlack) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceSlack) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceSlack) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceSlack) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceSlack) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceSlack) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

