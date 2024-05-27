// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package globallevelnotifications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v13/globallevelnotifications/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.0.0/docs/resources/global_level_notifications gitlab_global_level_notifications}.
type GlobalLevelNotifications interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloseIssue() interface{}
	SetCloseIssue(val interface{})
	CloseIssueInput() interface{}
	CloseMergeRequest() interface{}
	SetCloseMergeRequest(val interface{})
	CloseMergeRequestInput() interface{}
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
	FailedPipeline() interface{}
	SetFailedPipeline(val interface{})
	FailedPipelineInput() interface{}
	FixedPipeline() interface{}
	SetFixedPipeline(val interface{})
	FixedPipelineInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IssueDue() interface{}
	SetIssueDue(val interface{})
	IssueDueInput() interface{}
	Level() *string
	SetLevel(val *string)
	LevelInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MergeMergeRequest() interface{}
	SetMergeMergeRequest(val interface{})
	MergeMergeRequestInput() interface{}
	MergeWhenPipelineSucceeds() interface{}
	SetMergeWhenPipelineSucceeds(val interface{})
	MergeWhenPipelineSucceedsInput() interface{}
	MovedProject() interface{}
	SetMovedProject(val interface{})
	MovedProjectInput() interface{}
	NewIssue() interface{}
	SetNewIssue(val interface{})
	NewIssueInput() interface{}
	NewMergeRequest() interface{}
	SetNewMergeRequest(val interface{})
	NewMergeRequestInput() interface{}
	NewNote() interface{}
	SetNewNote(val interface{})
	NewNoteInput() interface{}
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PushToMergeRequest() interface{}
	SetPushToMergeRequest(val interface{})
	PushToMergeRequestInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	ReassignIssue() interface{}
	SetReassignIssue(val interface{})
	ReassignIssueInput() interface{}
	ReassignMergeRequest() interface{}
	SetReassignMergeRequest(val interface{})
	ReassignMergeRequestInput() interface{}
	ReopenIssue() interface{}
	SetReopenIssue(val interface{})
	ReopenIssueInput() interface{}
	ReopenMergeRequest() interface{}
	SetReopenMergeRequest(val interface{})
	ReopenMergeRequestInput() interface{}
	SuccessPipeline() interface{}
	SetSuccessPipeline(val interface{})
	SuccessPipelineInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetCloseIssue()
	ResetCloseMergeRequest()
	ResetFailedPipeline()
	ResetFixedPipeline()
	ResetIssueDue()
	ResetLevel()
	ResetMergeMergeRequest()
	ResetMergeWhenPipelineSucceeds()
	ResetMovedProject()
	ResetNewIssue()
	ResetNewMergeRequest()
	ResetNewNote()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPushToMergeRequest()
	ResetReassignIssue()
	ResetReassignMergeRequest()
	ResetReopenIssue()
	ResetReopenMergeRequest()
	ResetSuccessPipeline()
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

// The jsii proxy struct for GlobalLevelNotifications
type jsiiProxy_GlobalLevelNotifications struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GlobalLevelNotifications) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) CloseIssue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"closeIssue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) CloseIssueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"closeIssueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) CloseMergeRequest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"closeMergeRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) CloseMergeRequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"closeMergeRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) FailedPipeline() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failedPipeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) FailedPipelineInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failedPipelineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) FixedPipeline() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fixedPipeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) FixedPipelineInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fixedPipelineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) IssueDue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issueDue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) IssueDueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issueDueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) LevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"levelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) MergeMergeRequest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeMergeRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) MergeMergeRequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeMergeRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) MergeWhenPipelineSucceeds() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeWhenPipelineSucceeds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) MergeWhenPipelineSucceedsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeWhenPipelineSucceedsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) MovedProject() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"movedProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) MovedProjectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"movedProjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) NewIssue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newIssue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) NewIssueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newIssueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) NewMergeRequest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newMergeRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) NewMergeRequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newMergeRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) NewNote() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newNote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) NewNoteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newNoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) PushToMergeRequest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pushToMergeRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) PushToMergeRequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pushToMergeRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) ReassignIssue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reassignIssue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) ReassignIssueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reassignIssueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) ReassignMergeRequest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reassignMergeRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) ReassignMergeRequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reassignMergeRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) ReopenIssue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reopenIssue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) ReopenIssueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reopenIssueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) ReopenMergeRequest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reopenMergeRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) ReopenMergeRequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reopenMergeRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) SuccessPipeline() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"successPipeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) SuccessPipelineInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"successPipelineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalLevelNotifications) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.0.0/docs/resources/global_level_notifications gitlab_global_level_notifications} Resource.
func NewGlobalLevelNotifications(scope constructs.Construct, id *string, config *GlobalLevelNotificationsConfig) GlobalLevelNotifications {
	_init_.Initialize()

	if err := validateNewGlobalLevelNotificationsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlobalLevelNotifications{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.globalLevelNotifications.GlobalLevelNotifications",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.0.0/docs/resources/global_level_notifications gitlab_global_level_notifications} Resource.
func NewGlobalLevelNotifications_Override(g GlobalLevelNotifications, scope constructs.Construct, id *string, config *GlobalLevelNotificationsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.globalLevelNotifications.GlobalLevelNotifications",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GlobalLevelNotifications)SetCloseIssue(val interface{}) {
	if err := j.validateSetCloseIssueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"closeIssue",
		val,
	)
}

func (j *jsiiProxy_GlobalLevelNotifications)SetCloseMergeRequest(val interface{}) {
	if err := j.validateSetCloseMergeRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"closeMergeRequest",
		val,
	)
}

func (j *jsiiProxy_GlobalLevelNotifications)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GlobalLevelNotifications)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GlobalLevelNotifications)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GlobalLevelNotifications)SetFailedPipeline(val interface{}) {
	if err := j.validateSetFailedPipelineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failedPipeline",
		val,
	)
}

func (j *jsiiProxy_GlobalLevelNotifications)SetFixedPipeline(val interface{}) {
	if err := j.validateSetFixedPipelineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fixedPipeline",
		val,
	)
}

func (j *jsiiProxy_GlobalLevelNotifications)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GlobalLevelNotifications)SetIssueDue(val interface{}) {
	if err := j.validateSetIssueDueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issueDue",
		val,
	)
}

func (j *jsiiProxy_GlobalLevelNotifications)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_GlobalLevelNotifications)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GlobalLevelNotifications)SetMergeMergeRequest(val interface{}) {
	if err := j.validateSetMergeMergeRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeMergeRequest",
		val,
	)
}

func (j *jsiiProxy_GlobalLevelNotifications)SetMergeWhenPipelineSucceeds(val interface{}) {
	if err := j.validateSetMergeWhenPipelineSucceedsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeWhenPipelineSucceeds",
		val,
	)
}

func (j *jsiiProxy_GlobalLevelNotifications)SetMovedProject(val interface{}) {
	if err := j.validateSetMovedProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"movedProject",
		val,
	)
}

func (j *jsiiProxy_GlobalLevelNotifications)SetNewIssue(val interface{}) {
	if err := j.validateSetNewIssueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newIssue",
		val,
	)
}

func (j *jsiiProxy_GlobalLevelNotifications)SetNewMergeRequest(val interface{}) {
	if err := j.validateSetNewMergeRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newMergeRequest",
		val,
	)
}

func (j *jsiiProxy_GlobalLevelNotifications)SetNewNote(val interface{}) {
	if err := j.validateSetNewNoteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newNote",
		val,
	)
}

func (j *jsiiProxy_GlobalLevelNotifications)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GlobalLevelNotifications)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GlobalLevelNotifications)SetPushToMergeRequest(val interface{}) {
	if err := j.validateSetPushToMergeRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pushToMergeRequest",
		val,
	)
}

func (j *jsiiProxy_GlobalLevelNotifications)SetReassignIssue(val interface{}) {
	if err := j.validateSetReassignIssueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reassignIssue",
		val,
	)
}

func (j *jsiiProxy_GlobalLevelNotifications)SetReassignMergeRequest(val interface{}) {
	if err := j.validateSetReassignMergeRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reassignMergeRequest",
		val,
	)
}

func (j *jsiiProxy_GlobalLevelNotifications)SetReopenIssue(val interface{}) {
	if err := j.validateSetReopenIssueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reopenIssue",
		val,
	)
}

func (j *jsiiProxy_GlobalLevelNotifications)SetReopenMergeRequest(val interface{}) {
	if err := j.validateSetReopenMergeRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reopenMergeRequest",
		val,
	)
}

func (j *jsiiProxy_GlobalLevelNotifications)SetSuccessPipeline(val interface{}) {
	if err := j.validateSetSuccessPipelineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successPipeline",
		val,
	)
}

// Generates CDKTF code for importing a GlobalLevelNotifications resource upon running "cdktf plan <stack-name>".
func GlobalLevelNotifications_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGlobalLevelNotifications_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.globalLevelNotifications.GlobalLevelNotifications",
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
func GlobalLevelNotifications_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGlobalLevelNotifications_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.globalLevelNotifications.GlobalLevelNotifications",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GlobalLevelNotifications_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGlobalLevelNotifications_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.globalLevelNotifications.GlobalLevelNotifications",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GlobalLevelNotifications_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGlobalLevelNotifications_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.globalLevelNotifications.GlobalLevelNotifications",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GlobalLevelNotifications_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.globalLevelNotifications.GlobalLevelNotifications",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GlobalLevelNotifications) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GlobalLevelNotifications) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GlobalLevelNotifications) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlobalLevelNotifications) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GlobalLevelNotifications) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlobalLevelNotifications) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlobalLevelNotifications) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlobalLevelNotifications) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlobalLevelNotifications) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlobalLevelNotifications) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlobalLevelNotifications) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlobalLevelNotifications) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalLevelNotifications) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GlobalLevelNotifications) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GlobalLevelNotifications) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GlobalLevelNotifications) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GlobalLevelNotifications) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GlobalLevelNotifications) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GlobalLevelNotifications) ResetCloseIssue() {
	_jsii_.InvokeVoid(
		g,
		"resetCloseIssue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalLevelNotifications) ResetCloseMergeRequest() {
	_jsii_.InvokeVoid(
		g,
		"resetCloseMergeRequest",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalLevelNotifications) ResetFailedPipeline() {
	_jsii_.InvokeVoid(
		g,
		"resetFailedPipeline",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalLevelNotifications) ResetFixedPipeline() {
	_jsii_.InvokeVoid(
		g,
		"resetFixedPipeline",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalLevelNotifications) ResetIssueDue() {
	_jsii_.InvokeVoid(
		g,
		"resetIssueDue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalLevelNotifications) ResetLevel() {
	_jsii_.InvokeVoid(
		g,
		"resetLevel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalLevelNotifications) ResetMergeMergeRequest() {
	_jsii_.InvokeVoid(
		g,
		"resetMergeMergeRequest",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalLevelNotifications) ResetMergeWhenPipelineSucceeds() {
	_jsii_.InvokeVoid(
		g,
		"resetMergeWhenPipelineSucceeds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalLevelNotifications) ResetMovedProject() {
	_jsii_.InvokeVoid(
		g,
		"resetMovedProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalLevelNotifications) ResetNewIssue() {
	_jsii_.InvokeVoid(
		g,
		"resetNewIssue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalLevelNotifications) ResetNewMergeRequest() {
	_jsii_.InvokeVoid(
		g,
		"resetNewMergeRequest",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalLevelNotifications) ResetNewNote() {
	_jsii_.InvokeVoid(
		g,
		"resetNewNote",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalLevelNotifications) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalLevelNotifications) ResetPushToMergeRequest() {
	_jsii_.InvokeVoid(
		g,
		"resetPushToMergeRequest",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalLevelNotifications) ResetReassignIssue() {
	_jsii_.InvokeVoid(
		g,
		"resetReassignIssue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalLevelNotifications) ResetReassignMergeRequest() {
	_jsii_.InvokeVoid(
		g,
		"resetReassignMergeRequest",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalLevelNotifications) ResetReopenIssue() {
	_jsii_.InvokeVoid(
		g,
		"resetReopenIssue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalLevelNotifications) ResetReopenMergeRequest() {
	_jsii_.InvokeVoid(
		g,
		"resetReopenMergeRequest",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalLevelNotifications) ResetSuccessPipeline() {
	_jsii_.InvokeVoid(
		g,
		"resetSuccessPipeline",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalLevelNotifications) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalLevelNotifications) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalLevelNotifications) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalLevelNotifications) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalLevelNotifications) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalLevelNotifications) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

