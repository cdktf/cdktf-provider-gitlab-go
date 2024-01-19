// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectlevelnotifications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v12/projectlevelnotifications/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/project_level_notifications gitlab_project_level_notifications}.
type ProjectLevelNotifications interface {
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

// The jsii proxy struct for ProjectLevelNotifications
type jsiiProxy_ProjectLevelNotifications struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ProjectLevelNotifications) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) CloseIssue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"closeIssue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) CloseIssueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"closeIssueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) CloseMergeRequest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"closeMergeRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) CloseMergeRequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"closeMergeRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) FailedPipeline() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failedPipeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) FailedPipelineInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failedPipelineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) FixedPipeline() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fixedPipeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) FixedPipelineInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fixedPipelineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) IssueDue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issueDue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) IssueDueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issueDueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) LevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"levelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) MergeMergeRequest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeMergeRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) MergeMergeRequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeMergeRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) MergeWhenPipelineSucceeds() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeWhenPipelineSucceeds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) MergeWhenPipelineSucceedsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeWhenPipelineSucceedsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) MovedProject() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"movedProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) MovedProjectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"movedProjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) NewIssue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newIssue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) NewIssueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newIssueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) NewMergeRequest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newMergeRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) NewMergeRequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newMergeRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) NewNote() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newNote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) NewNoteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newNoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) PushToMergeRequest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pushToMergeRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) PushToMergeRequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pushToMergeRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) ReassignIssue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reassignIssue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) ReassignIssueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reassignIssueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) ReassignMergeRequest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reassignMergeRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) ReassignMergeRequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reassignMergeRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) ReopenIssue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reopenIssue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) ReopenIssueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reopenIssueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) ReopenMergeRequest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reopenMergeRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) ReopenMergeRequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reopenMergeRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) SuccessPipeline() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"successPipeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) SuccessPipelineInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"successPipelineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelNotifications) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/project_level_notifications gitlab_project_level_notifications} Resource.
func NewProjectLevelNotifications(scope constructs.Construct, id *string, config *ProjectLevelNotificationsConfig) ProjectLevelNotifications {
	_init_.Initialize()

	if err := validateNewProjectLevelNotificationsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProjectLevelNotifications{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.projectLevelNotifications.ProjectLevelNotifications",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/resources/project_level_notifications gitlab_project_level_notifications} Resource.
func NewProjectLevelNotifications_Override(p ProjectLevelNotifications, scope constructs.Construct, id *string, config *ProjectLevelNotificationsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.projectLevelNotifications.ProjectLevelNotifications",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_ProjectLevelNotifications)SetCloseIssue(val interface{}) {
	if err := j.validateSetCloseIssueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"closeIssue",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelNotifications)SetCloseMergeRequest(val interface{}) {
	if err := j.validateSetCloseMergeRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"closeMergeRequest",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelNotifications)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelNotifications)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelNotifications)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelNotifications)SetFailedPipeline(val interface{}) {
	if err := j.validateSetFailedPipelineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failedPipeline",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelNotifications)SetFixedPipeline(val interface{}) {
	if err := j.validateSetFixedPipelineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fixedPipeline",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelNotifications)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelNotifications)SetIssueDue(val interface{}) {
	if err := j.validateSetIssueDueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issueDue",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelNotifications)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelNotifications)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelNotifications)SetMergeMergeRequest(val interface{}) {
	if err := j.validateSetMergeMergeRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeMergeRequest",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelNotifications)SetMergeWhenPipelineSucceeds(val interface{}) {
	if err := j.validateSetMergeWhenPipelineSucceedsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeWhenPipelineSucceeds",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelNotifications)SetMovedProject(val interface{}) {
	if err := j.validateSetMovedProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"movedProject",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelNotifications)SetNewIssue(val interface{}) {
	if err := j.validateSetNewIssueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newIssue",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelNotifications)SetNewMergeRequest(val interface{}) {
	if err := j.validateSetNewMergeRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newMergeRequest",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelNotifications)SetNewNote(val interface{}) {
	if err := j.validateSetNewNoteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newNote",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelNotifications)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelNotifications)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelNotifications)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelNotifications)SetPushToMergeRequest(val interface{}) {
	if err := j.validateSetPushToMergeRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pushToMergeRequest",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelNotifications)SetReassignIssue(val interface{}) {
	if err := j.validateSetReassignIssueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reassignIssue",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelNotifications)SetReassignMergeRequest(val interface{}) {
	if err := j.validateSetReassignMergeRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reassignMergeRequest",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelNotifications)SetReopenIssue(val interface{}) {
	if err := j.validateSetReopenIssueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reopenIssue",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelNotifications)SetReopenMergeRequest(val interface{}) {
	if err := j.validateSetReopenMergeRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reopenMergeRequest",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelNotifications)SetSuccessPipeline(val interface{}) {
	if err := j.validateSetSuccessPipelineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successPipeline",
		val,
	)
}

// Generates CDKTF code for importing a ProjectLevelNotifications resource upon running "cdktf plan <stack-name>".
func ProjectLevelNotifications_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateProjectLevelNotifications_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.projectLevelNotifications.ProjectLevelNotifications",
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
func ProjectLevelNotifications_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjectLevelNotifications_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.projectLevelNotifications.ProjectLevelNotifications",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ProjectLevelNotifications_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjectLevelNotifications_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.projectLevelNotifications.ProjectLevelNotifications",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ProjectLevelNotifications_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjectLevelNotifications_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.projectLevelNotifications.ProjectLevelNotifications",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ProjectLevelNotifications_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.projectLevelNotifications.ProjectLevelNotifications",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_ProjectLevelNotifications) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_ProjectLevelNotifications) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_ProjectLevelNotifications) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_ProjectLevelNotifications) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_ProjectLevelNotifications) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_ProjectLevelNotifications) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_ProjectLevelNotifications) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_ProjectLevelNotifications) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_ProjectLevelNotifications) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_ProjectLevelNotifications) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_ProjectLevelNotifications) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_ProjectLevelNotifications) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectLevelNotifications) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_ProjectLevelNotifications) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_ProjectLevelNotifications) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_ProjectLevelNotifications) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_ProjectLevelNotifications) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_ProjectLevelNotifications) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_ProjectLevelNotifications) ResetCloseIssue() {
	_jsii_.InvokeVoid(
		p,
		"resetCloseIssue",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectLevelNotifications) ResetCloseMergeRequest() {
	_jsii_.InvokeVoid(
		p,
		"resetCloseMergeRequest",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectLevelNotifications) ResetFailedPipeline() {
	_jsii_.InvokeVoid(
		p,
		"resetFailedPipeline",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectLevelNotifications) ResetFixedPipeline() {
	_jsii_.InvokeVoid(
		p,
		"resetFixedPipeline",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectLevelNotifications) ResetIssueDue() {
	_jsii_.InvokeVoid(
		p,
		"resetIssueDue",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectLevelNotifications) ResetLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectLevelNotifications) ResetMergeMergeRequest() {
	_jsii_.InvokeVoid(
		p,
		"resetMergeMergeRequest",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectLevelNotifications) ResetMergeWhenPipelineSucceeds() {
	_jsii_.InvokeVoid(
		p,
		"resetMergeWhenPipelineSucceeds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectLevelNotifications) ResetMovedProject() {
	_jsii_.InvokeVoid(
		p,
		"resetMovedProject",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectLevelNotifications) ResetNewIssue() {
	_jsii_.InvokeVoid(
		p,
		"resetNewIssue",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectLevelNotifications) ResetNewMergeRequest() {
	_jsii_.InvokeVoid(
		p,
		"resetNewMergeRequest",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectLevelNotifications) ResetNewNote() {
	_jsii_.InvokeVoid(
		p,
		"resetNewNote",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectLevelNotifications) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectLevelNotifications) ResetPushToMergeRequest() {
	_jsii_.InvokeVoid(
		p,
		"resetPushToMergeRequest",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectLevelNotifications) ResetReassignIssue() {
	_jsii_.InvokeVoid(
		p,
		"resetReassignIssue",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectLevelNotifications) ResetReassignMergeRequest() {
	_jsii_.InvokeVoid(
		p,
		"resetReassignMergeRequest",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectLevelNotifications) ResetReopenIssue() {
	_jsii_.InvokeVoid(
		p,
		"resetReopenIssue",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectLevelNotifications) ResetReopenMergeRequest() {
	_jsii_.InvokeVoid(
		p,
		"resetReopenMergeRequest",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectLevelNotifications) ResetSuccessPipeline() {
	_jsii_.InvokeVoid(
		p,
		"resetSuccessPipeline",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectLevelNotifications) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectLevelNotifications) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectLevelNotifications) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectLevelNotifications) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectLevelNotifications) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectLevelNotifications) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

