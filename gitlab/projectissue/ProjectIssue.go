// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectissue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v12/projectissue/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.10.0/docs/resources/project_issue gitlab_project_issue}.
type ProjectIssue interface {
	cdktf.TerraformResource
	AssigneeIds() *[]*float64
	SetAssigneeIds(val *[]*float64)
	AssigneeIdsInput() *[]*float64
	AuthorId() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClosedAt() *string
	ClosedByUserId() *float64
	Confidential() interface{}
	SetConfidential(val interface{})
	ConfidentialInput() interface{}
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
	CreatedAt() *string
	SetCreatedAt(val *string)
	CreatedAtInput() *string
	DeleteOnDestroy() interface{}
	SetDeleteOnDestroy(val interface{})
	DeleteOnDestroyInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DiscussionLocked() interface{}
	SetDiscussionLocked(val interface{})
	DiscussionLockedInput() interface{}
	DiscussionToResolve() *string
	SetDiscussionToResolve(val *string)
	DiscussionToResolveInput() *string
	Downvotes() *float64
	DueDate() *string
	SetDueDate(val *string)
	DueDateInput() *string
	EpicId() *float64
	EpicIssueId() *float64
	SetEpicIssueId(val *float64)
	EpicIssueIdInput() *float64
	ExternalId() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HumanTimeEstimate() *string
	HumanTotalTimeSpent() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Iid() *float64
	SetIid(val *float64)
	IidInput() *float64
	IssueId() *float64
	IssueLinkId() *float64
	IssueType() *string
	SetIssueType(val *string)
	IssueTypeInput() *string
	Labels() *[]*string
	SetLabels(val *[]*string)
	LabelsInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Links() cdktf.StringMap
	MergeRequestsCount() *float64
	MergeRequestToResolveDiscussionsOf() *float64
	SetMergeRequestToResolveDiscussionsOf(val *float64)
	MergeRequestToResolveDiscussionsOfInput() *float64
	MilestoneId() *float64
	SetMilestoneId(val *float64)
	MilestoneIdInput() *float64
	MovedToId() *float64
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
	// Experimental.
	RawOverrides() interface{}
	References() cdktf.StringMap
	State() *string
	SetState(val *string)
	StateInput() *string
	Subscribed() cdktf.IResolvable
	TaskCompletionStatus() ProjectIssueTaskCompletionStatusList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeEstimate() *float64
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
	TotalTimeSpent() *float64
	UpdatedAt() *string
	SetUpdatedAt(val *string)
	UpdatedAtInput() *string
	Upvotes() *float64
	UserNotesCount() *float64
	WebUrl() *string
	Weight() *float64
	SetWeight(val *float64)
	WeightInput() *float64
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
	ResetAssigneeIds()
	ResetConfidential()
	ResetCreatedAt()
	ResetDeleteOnDestroy()
	ResetDescription()
	ResetDiscussionLocked()
	ResetDiscussionToResolve()
	ResetDueDate()
	ResetEpicIssueId()
	ResetId()
	ResetIid()
	ResetIssueType()
	ResetLabels()
	ResetMergeRequestToResolveDiscussionsOf()
	ResetMilestoneId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetState()
	ResetUpdatedAt()
	ResetWeight()
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

// The jsii proxy struct for ProjectIssue
type jsiiProxy_ProjectIssue struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ProjectIssue) AssigneeIds() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"assigneeIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) AssigneeIdsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"assigneeIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) AuthorId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) ClosedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"closedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) ClosedByUserId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"closedByUserId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) Confidential() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) ConfidentialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) CreatedAtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) DeleteOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) DeleteOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) DiscussionLocked() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"discussionLocked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) DiscussionLockedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"discussionLockedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) DiscussionToResolve() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discussionToResolve",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) DiscussionToResolveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discussionToResolveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) Downvotes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"downvotes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) DueDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dueDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) DueDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dueDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) EpicId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"epicId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) EpicIssueId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"epicIssueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) EpicIssueIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"epicIssueIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) HumanTimeEstimate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"humanTimeEstimate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) HumanTotalTimeSpent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"humanTotalTimeSpent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) Iid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) IidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) IssueId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"issueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) IssueLinkId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"issueLinkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) IssueType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issueType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) IssueTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issueTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) Labels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) LabelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) Links() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"links",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) MergeRequestsCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mergeRequestsCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) MergeRequestToResolveDiscussionsOf() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mergeRequestToResolveDiscussionsOf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) MergeRequestToResolveDiscussionsOfInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mergeRequestToResolveDiscussionsOfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) MilestoneId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"milestoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) MilestoneIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"milestoneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) MovedToId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"movedToId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) References() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"references",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) Subscribed() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"subscribed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) TaskCompletionStatus() ProjectIssueTaskCompletionStatusList {
	var returns ProjectIssueTaskCompletionStatusList
	_jsii_.Get(
		j,
		"taskCompletionStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) TimeEstimate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeEstimate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) TotalTimeSpent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalTimeSpent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) UpdatedAtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) Upvotes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"upvotes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) UserNotesCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userNotesCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) WebUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) Weight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssue) WeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.10.0/docs/resources/project_issue gitlab_project_issue} Resource.
func NewProjectIssue(scope constructs.Construct, id *string, config *ProjectIssueConfig) ProjectIssue {
	_init_.Initialize()

	if err := validateNewProjectIssueParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProjectIssue{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.projectIssue.ProjectIssue",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.10.0/docs/resources/project_issue gitlab_project_issue} Resource.
func NewProjectIssue_Override(p ProjectIssue, scope constructs.Construct, id *string, config *ProjectIssueConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.projectIssue.ProjectIssue",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_ProjectIssue)SetAssigneeIds(val *[]*float64) {
	if err := j.validateSetAssigneeIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assigneeIds",
		val,
	)
}

func (j *jsiiProxy_ProjectIssue)SetConfidential(val interface{}) {
	if err := j.validateSetConfidentialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidential",
		val,
	)
}

func (j *jsiiProxy_ProjectIssue)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ProjectIssue)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ProjectIssue)SetCreatedAt(val *string) {
	if err := j.validateSetCreatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdAt",
		val,
	)
}

func (j *jsiiProxy_ProjectIssue)SetDeleteOnDestroy(val interface{}) {
	if err := j.validateSetDeleteOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteOnDestroy",
		val,
	)
}

func (j *jsiiProxy_ProjectIssue)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ProjectIssue)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ProjectIssue)SetDiscussionLocked(val interface{}) {
	if err := j.validateSetDiscussionLockedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"discussionLocked",
		val,
	)
}

func (j *jsiiProxy_ProjectIssue)SetDiscussionToResolve(val *string) {
	if err := j.validateSetDiscussionToResolveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"discussionToResolve",
		val,
	)
}

func (j *jsiiProxy_ProjectIssue)SetDueDate(val *string) {
	if err := j.validateSetDueDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dueDate",
		val,
	)
}

func (j *jsiiProxy_ProjectIssue)SetEpicIssueId(val *float64) {
	if err := j.validateSetEpicIssueIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"epicIssueId",
		val,
	)
}

func (j *jsiiProxy_ProjectIssue)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ProjectIssue)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ProjectIssue)SetIid(val *float64) {
	if err := j.validateSetIidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iid",
		val,
	)
}

func (j *jsiiProxy_ProjectIssue)SetIssueType(val *string) {
	if err := j.validateSetIssueTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issueType",
		val,
	)
}

func (j *jsiiProxy_ProjectIssue)SetLabels(val *[]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_ProjectIssue)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ProjectIssue)SetMergeRequestToResolveDiscussionsOf(val *float64) {
	if err := j.validateSetMergeRequestToResolveDiscussionsOfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeRequestToResolveDiscussionsOf",
		val,
	)
}

func (j *jsiiProxy_ProjectIssue)SetMilestoneId(val *float64) {
	if err := j.validateSetMilestoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"milestoneId",
		val,
	)
}

func (j *jsiiProxy_ProjectIssue)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ProjectIssue)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ProjectIssue)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ProjectIssue)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_ProjectIssue)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_ProjectIssue)SetUpdatedAt(val *string) {
	if err := j.validateSetUpdatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedAt",
		val,
	)
}

func (j *jsiiProxy_ProjectIssue)SetWeight(val *float64) {
	if err := j.validateSetWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weight",
		val,
	)
}

// Generates CDKTF code for importing a ProjectIssue resource upon running "cdktf plan <stack-name>".
func ProjectIssue_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateProjectIssue_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.projectIssue.ProjectIssue",
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
func ProjectIssue_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjectIssue_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.projectIssue.ProjectIssue",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ProjectIssue_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjectIssue_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.projectIssue.ProjectIssue",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ProjectIssue_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjectIssue_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.projectIssue.ProjectIssue",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ProjectIssue_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.projectIssue.ProjectIssue",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_ProjectIssue) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_ProjectIssue) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_ProjectIssue) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_ProjectIssue) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_ProjectIssue) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_ProjectIssue) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_ProjectIssue) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_ProjectIssue) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_ProjectIssue) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_ProjectIssue) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_ProjectIssue) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_ProjectIssue) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIssue) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_ProjectIssue) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_ProjectIssue) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_ProjectIssue) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_ProjectIssue) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_ProjectIssue) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_ProjectIssue) ResetAssigneeIds() {
	_jsii_.InvokeVoid(
		p,
		"resetAssigneeIds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIssue) ResetConfidential() {
	_jsii_.InvokeVoid(
		p,
		"resetConfidential",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIssue) ResetCreatedAt() {
	_jsii_.InvokeVoid(
		p,
		"resetCreatedAt",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIssue) ResetDeleteOnDestroy() {
	_jsii_.InvokeVoid(
		p,
		"resetDeleteOnDestroy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIssue) ResetDescription() {
	_jsii_.InvokeVoid(
		p,
		"resetDescription",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIssue) ResetDiscussionLocked() {
	_jsii_.InvokeVoid(
		p,
		"resetDiscussionLocked",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIssue) ResetDiscussionToResolve() {
	_jsii_.InvokeVoid(
		p,
		"resetDiscussionToResolve",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIssue) ResetDueDate() {
	_jsii_.InvokeVoid(
		p,
		"resetDueDate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIssue) ResetEpicIssueId() {
	_jsii_.InvokeVoid(
		p,
		"resetEpicIssueId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIssue) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIssue) ResetIid() {
	_jsii_.InvokeVoid(
		p,
		"resetIid",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIssue) ResetIssueType() {
	_jsii_.InvokeVoid(
		p,
		"resetIssueType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIssue) ResetLabels() {
	_jsii_.InvokeVoid(
		p,
		"resetLabels",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIssue) ResetMergeRequestToResolveDiscussionsOf() {
	_jsii_.InvokeVoid(
		p,
		"resetMergeRequestToResolveDiscussionsOf",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIssue) ResetMilestoneId() {
	_jsii_.InvokeVoid(
		p,
		"resetMilestoneId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIssue) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIssue) ResetState() {
	_jsii_.InvokeVoid(
		p,
		"resetState",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIssue) ResetUpdatedAt() {
	_jsii_.InvokeVoid(
		p,
		"resetUpdatedAt",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIssue) ResetWeight() {
	_jsii_.InvokeVoid(
		p,
		"resetWeight",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIssue) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIssue) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIssue) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIssue) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIssue) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIssue) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

