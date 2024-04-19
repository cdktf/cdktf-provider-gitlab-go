// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grouphook

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v12/grouphook/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.11.0/docs/resources/group_hook gitlab_group_hook}.
type GroupHook interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	CustomWebhookTemplate() *string
	SetCustomWebhookTemplate(val *string)
	CustomWebhookTemplateInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeploymentEvents() interface{}
	SetDeploymentEvents(val interface{})
	DeploymentEventsInput() interface{}
	EnableSslVerification() interface{}
	SetEnableSslVerification(val interface{})
	EnableSslVerificationInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Group() *string
	SetGroup(val *string)
	GroupId() *float64
	GroupInput() *string
	HookId() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
	IssuesEvents() interface{}
	SetIssuesEvents(val interface{})
	IssuesEventsInput() interface{}
	JobEvents() interface{}
	SetJobEvents(val interface{})
	JobEventsInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MergeRequestsEvents() interface{}
	SetMergeRequestsEvents(val interface{})
	MergeRequestsEventsInput() interface{}
	// The tree node.
	Node() constructs.Node
	NoteEvents() interface{}
	SetNoteEvents(val interface{})
	NoteEventsInput() interface{}
	PipelineEvents() interface{}
	SetPipelineEvents(val interface{})
	PipelineEventsInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PushEvents() interface{}
	SetPushEvents(val interface{})
	PushEventsBranchFilter() *string
	SetPushEventsBranchFilter(val *string)
	PushEventsBranchFilterInput() *string
	PushEventsInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	ReleasesEvents() interface{}
	SetReleasesEvents(val interface{})
	ReleasesEventsInput() interface{}
	SubgroupEvents() interface{}
	SetSubgroupEvents(val interface{})
	SubgroupEventsInput() interface{}
	TagPushEvents() interface{}
	SetTagPushEvents(val interface{})
	TagPushEventsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
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
	ResetConfidentialIssuesEvents()
	ResetConfidentialNoteEvents()
	ResetCustomWebhookTemplate()
	ResetDeploymentEvents()
	ResetEnableSslVerification()
	ResetId()
	ResetIssuesEvents()
	ResetJobEvents()
	ResetMergeRequestsEvents()
	ResetNoteEvents()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPipelineEvents()
	ResetPushEvents()
	ResetPushEventsBranchFilter()
	ResetReleasesEvents()
	ResetSubgroupEvents()
	ResetTagPushEvents()
	ResetToken()
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

// The jsii proxy struct for GroupHook
type jsiiProxy_GroupHook struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GroupHook) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) ConfidentialIssuesEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidentialIssuesEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) ConfidentialIssuesEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidentialIssuesEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) ConfidentialNoteEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidentialNoteEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) ConfidentialNoteEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidentialNoteEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) CustomWebhookTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customWebhookTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) CustomWebhookTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customWebhookTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) DeploymentEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) DeploymentEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) EnableSslVerification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSslVerification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) EnableSslVerificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSslVerificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) Group() *string {
	var returns *string
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) GroupId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) GroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) HookId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hookId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) IssuesEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issuesEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) IssuesEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issuesEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) JobEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) JobEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) MergeRequestsEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeRequestsEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) MergeRequestsEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeRequestsEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) NoteEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) NoteEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) PipelineEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pipelineEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) PipelineEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pipelineEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) PushEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pushEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) PushEventsBranchFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pushEventsBranchFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) PushEventsBranchFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pushEventsBranchFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) PushEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pushEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) ReleasesEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"releasesEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) ReleasesEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"releasesEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) SubgroupEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subgroupEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) SubgroupEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subgroupEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) TagPushEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagPushEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) TagPushEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagPushEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) WikiPageEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wikiPageEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupHook) WikiPageEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wikiPageEventsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.11.0/docs/resources/group_hook gitlab_group_hook} Resource.
func NewGroupHook(scope constructs.Construct, id *string, config *GroupHookConfig) GroupHook {
	_init_.Initialize()

	if err := validateNewGroupHookParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GroupHook{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.groupHook.GroupHook",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.11.0/docs/resources/group_hook gitlab_group_hook} Resource.
func NewGroupHook_Override(g GroupHook, scope constructs.Construct, id *string, config *GroupHookConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.groupHook.GroupHook",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GroupHook)SetConfidentialIssuesEvents(val interface{}) {
	if err := j.validateSetConfidentialIssuesEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidentialIssuesEvents",
		val,
	)
}

func (j *jsiiProxy_GroupHook)SetConfidentialNoteEvents(val interface{}) {
	if err := j.validateSetConfidentialNoteEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidentialNoteEvents",
		val,
	)
}

func (j *jsiiProxy_GroupHook)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GroupHook)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GroupHook)SetCustomWebhookTemplate(val *string) {
	if err := j.validateSetCustomWebhookTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customWebhookTemplate",
		val,
	)
}

func (j *jsiiProxy_GroupHook)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GroupHook)SetDeploymentEvents(val interface{}) {
	if err := j.validateSetDeploymentEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentEvents",
		val,
	)
}

func (j *jsiiProxy_GroupHook)SetEnableSslVerification(val interface{}) {
	if err := j.validateSetEnableSslVerificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSslVerification",
		val,
	)
}

func (j *jsiiProxy_GroupHook)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GroupHook)SetGroup(val *string) {
	if err := j.validateSetGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"group",
		val,
	)
}

func (j *jsiiProxy_GroupHook)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GroupHook)SetIssuesEvents(val interface{}) {
	if err := j.validateSetIssuesEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuesEvents",
		val,
	)
}

func (j *jsiiProxy_GroupHook)SetJobEvents(val interface{}) {
	if err := j.validateSetJobEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobEvents",
		val,
	)
}

func (j *jsiiProxy_GroupHook)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GroupHook)SetMergeRequestsEvents(val interface{}) {
	if err := j.validateSetMergeRequestsEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeRequestsEvents",
		val,
	)
}

func (j *jsiiProxy_GroupHook)SetNoteEvents(val interface{}) {
	if err := j.validateSetNoteEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noteEvents",
		val,
	)
}

func (j *jsiiProxy_GroupHook)SetPipelineEvents(val interface{}) {
	if err := j.validateSetPipelineEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipelineEvents",
		val,
	)
}

func (j *jsiiProxy_GroupHook)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GroupHook)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GroupHook)SetPushEvents(val interface{}) {
	if err := j.validateSetPushEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pushEvents",
		val,
	)
}

func (j *jsiiProxy_GroupHook)SetPushEventsBranchFilter(val *string) {
	if err := j.validateSetPushEventsBranchFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pushEventsBranchFilter",
		val,
	)
}

func (j *jsiiProxy_GroupHook)SetReleasesEvents(val interface{}) {
	if err := j.validateSetReleasesEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releasesEvents",
		val,
	)
}

func (j *jsiiProxy_GroupHook)SetSubgroupEvents(val interface{}) {
	if err := j.validateSetSubgroupEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subgroupEvents",
		val,
	)
}

func (j *jsiiProxy_GroupHook)SetTagPushEvents(val interface{}) {
	if err := j.validateSetTagPushEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagPushEvents",
		val,
	)
}

func (j *jsiiProxy_GroupHook)SetToken(val *string) {
	if err := j.validateSetTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

func (j *jsiiProxy_GroupHook)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_GroupHook)SetWikiPageEvents(val interface{}) {
	if err := j.validateSetWikiPageEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wikiPageEvents",
		val,
	)
}

// Generates CDKTF code for importing a GroupHook resource upon running "cdktf plan <stack-name>".
func GroupHook_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGroupHook_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.groupHook.GroupHook",
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
func GroupHook_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGroupHook_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.groupHook.GroupHook",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GroupHook_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGroupHook_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.groupHook.GroupHook",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GroupHook_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGroupHook_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.groupHook.GroupHook",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GroupHook_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.groupHook.GroupHook",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GroupHook) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GroupHook) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GroupHook) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupHook) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupHook) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupHook) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupHook) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupHook) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupHook) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupHook) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupHook) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupHook) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupHook) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GroupHook) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupHook) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GroupHook) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GroupHook) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GroupHook) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GroupHook) ResetConfidentialIssuesEvents() {
	_jsii_.InvokeVoid(
		g,
		"resetConfidentialIssuesEvents",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupHook) ResetConfidentialNoteEvents() {
	_jsii_.InvokeVoid(
		g,
		"resetConfidentialNoteEvents",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupHook) ResetCustomWebhookTemplate() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomWebhookTemplate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupHook) ResetDeploymentEvents() {
	_jsii_.InvokeVoid(
		g,
		"resetDeploymentEvents",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupHook) ResetEnableSslVerification() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableSslVerification",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupHook) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupHook) ResetIssuesEvents() {
	_jsii_.InvokeVoid(
		g,
		"resetIssuesEvents",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupHook) ResetJobEvents() {
	_jsii_.InvokeVoid(
		g,
		"resetJobEvents",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupHook) ResetMergeRequestsEvents() {
	_jsii_.InvokeVoid(
		g,
		"resetMergeRequestsEvents",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupHook) ResetNoteEvents() {
	_jsii_.InvokeVoid(
		g,
		"resetNoteEvents",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupHook) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupHook) ResetPipelineEvents() {
	_jsii_.InvokeVoid(
		g,
		"resetPipelineEvents",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupHook) ResetPushEvents() {
	_jsii_.InvokeVoid(
		g,
		"resetPushEvents",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupHook) ResetPushEventsBranchFilter() {
	_jsii_.InvokeVoid(
		g,
		"resetPushEventsBranchFilter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupHook) ResetReleasesEvents() {
	_jsii_.InvokeVoid(
		g,
		"resetReleasesEvents",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupHook) ResetSubgroupEvents() {
	_jsii_.InvokeVoid(
		g,
		"resetSubgroupEvents",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupHook) ResetTagPushEvents() {
	_jsii_.InvokeVoid(
		g,
		"resetTagPushEvents",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupHook) ResetToken() {
	_jsii_.InvokeVoid(
		g,
		"resetToken",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupHook) ResetWikiPageEvents() {
	_jsii_.InvokeVoid(
		g,
		"resetWikiPageEvents",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupHook) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupHook) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupHook) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupHook) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupHook) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupHook) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

