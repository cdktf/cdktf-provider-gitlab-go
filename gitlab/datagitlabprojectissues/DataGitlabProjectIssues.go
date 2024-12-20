// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabprojectissues

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v13/datagitlabprojectissues/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.7.0/docs/data-sources/project_issues gitlab_project_issues}.
type DataGitlabProjectIssues interface {
	cdktf.TerraformDataSource
	AssigneeId() *float64
	SetAssigneeId(val *float64)
	AssigneeIdInput() *float64
	AssigneeUsername() *string
	SetAssigneeUsername(val *string)
	AssigneeUsernameInput() *string
	AuthorId() *float64
	SetAuthorId(val *float64)
	AuthorIdInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Confidential() interface{}
	SetConfidential(val interface{})
	ConfidentialInput() interface{}
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAfter() *string
	SetCreatedAfter(val *string)
	CreatedAfterInput() *string
	CreatedBefore() *string
	SetCreatedBefore(val *string)
	CreatedBeforeInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DueDate() *string
	SetDueDate(val *string)
	DueDateInput() *string
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
	Iids() *[]*float64
	SetIids(val *[]*float64)
	IidsInput() *[]*float64
	Issues() DataGitlabProjectIssuesIssuesList
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
	Milestone() *string
	SetMilestone(val *string)
	MilestoneInput() *string
	MyReactionEmoji() *string
	SetMyReactionEmoji(val *string)
	MyReactionEmojiInput() *string
	// The tree node.
	Node() constructs.Node
	NotAssigneeId() *[]*float64
	SetNotAssigneeId(val *[]*float64)
	NotAssigneeIdInput() *[]*float64
	NotAuthorId() *[]*float64
	SetNotAuthorId(val *[]*float64)
	NotAuthorIdInput() *[]*float64
	NotLabels() *[]*string
	SetNotLabels(val *[]*string)
	NotLabelsInput() *[]*string
	NotMilestone() *string
	SetNotMilestone(val *string)
	NotMilestoneInput() *string
	NotMyReactionEmoji() *[]*string
	SetNotMyReactionEmoji(val *[]*string)
	NotMyReactionEmojiInput() *[]*string
	OrderBy() *string
	SetOrderBy(val *string)
	OrderByInput() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Scope() *string
	SetScope(val *string)
	ScopeInput() *string
	Search() *string
	SetSearch(val *string)
	SearchInput() *string
	Sort() *string
	SetSort(val *string)
	SortInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdatedAfter() *string
	SetUpdatedAfter(val *string)
	UpdatedAfterInput() *string
	UpdatedBefore() *string
	SetUpdatedBefore(val *string)
	UpdatedBeforeInput() *string
	Weight() *float64
	SetWeight(val *float64)
	WeightInput() *float64
	WithLabelsDetails() interface{}
	SetWithLabelsDetails(val interface{})
	WithLabelsDetailsInput() interface{}
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAssigneeId()
	ResetAssigneeUsername()
	ResetAuthorId()
	ResetConfidential()
	ResetCreatedAfter()
	ResetCreatedBefore()
	ResetDueDate()
	ResetId()
	ResetIids()
	ResetIssueType()
	ResetLabels()
	ResetMilestone()
	ResetMyReactionEmoji()
	ResetNotAssigneeId()
	ResetNotAuthorId()
	ResetNotLabels()
	ResetNotMilestone()
	ResetNotMyReactionEmoji()
	ResetOrderBy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetScope()
	ResetSearch()
	ResetSort()
	ResetUpdatedAfter()
	ResetUpdatedBefore()
	ResetWeight()
	ResetWithLabelsDetails()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataGitlabProjectIssues
type jsiiProxy_DataGitlabProjectIssues struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataGitlabProjectIssues) AssigneeId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"assigneeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) AssigneeIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"assigneeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) AssigneeUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assigneeUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) AssigneeUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assigneeUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) AuthorId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) AuthorIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) Confidential() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) ConfidentialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) CreatedAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) CreatedAfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) CreatedBefore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) CreatedBeforeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBeforeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) DueDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dueDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) DueDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dueDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) Iids() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"iids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) IidsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"iidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) Issues() DataGitlabProjectIssuesIssuesList {
	var returns DataGitlabProjectIssuesIssuesList
	_jsii_.Get(
		j,
		"issues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) IssueType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issueType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) IssueTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issueTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) Labels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) LabelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) Milestone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"milestone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) MilestoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"milestoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) MyReactionEmoji() *string {
	var returns *string
	_jsii_.Get(
		j,
		"myReactionEmoji",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) MyReactionEmojiInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"myReactionEmojiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) NotAssigneeId() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"notAssigneeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) NotAssigneeIdInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"notAssigneeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) NotAuthorId() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"notAuthorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) NotAuthorIdInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"notAuthorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) NotLabels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) NotLabelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) NotMilestone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notMilestone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) NotMilestoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notMilestoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) NotMyReactionEmoji() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notMyReactionEmoji",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) NotMyReactionEmojiInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notMyReactionEmojiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) OrderBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) OrderByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) Search() *string {
	var returns *string
	_jsii_.Get(
		j,
		"search",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) SearchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) Sort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) SortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) UpdatedAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) UpdatedAfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) UpdatedBefore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) UpdatedBeforeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedBeforeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) Weight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) WeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) WithLabelsDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withLabelsDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssues) WithLabelsDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withLabelsDetailsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.7.0/docs/data-sources/project_issues gitlab_project_issues} Data Source.
func NewDataGitlabProjectIssues(scope constructs.Construct, id *string, config *DataGitlabProjectIssuesConfig) DataGitlabProjectIssues {
	_init_.Initialize()

	if err := validateNewDataGitlabProjectIssuesParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGitlabProjectIssues{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.dataGitlabProjectIssues.DataGitlabProjectIssues",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.7.0/docs/data-sources/project_issues gitlab_project_issues} Data Source.
func NewDataGitlabProjectIssues_Override(d DataGitlabProjectIssues, scope constructs.Construct, id *string, config *DataGitlabProjectIssuesConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.dataGitlabProjectIssues.DataGitlabProjectIssues",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetAssigneeId(val *float64) {
	if err := j.validateSetAssigneeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assigneeId",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetAssigneeUsername(val *string) {
	if err := j.validateSetAssigneeUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assigneeUsername",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetAuthorId(val *float64) {
	if err := j.validateSetAuthorIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorId",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetConfidential(val interface{}) {
	if err := j.validateSetConfidentialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidential",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetCreatedAfter(val *string) {
	if err := j.validateSetCreatedAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdAfter",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetCreatedBefore(val *string) {
	if err := j.validateSetCreatedBeforeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdBefore",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetDueDate(val *string) {
	if err := j.validateSetDueDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dueDate",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetIids(val *[]*float64) {
	if err := j.validateSetIidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iids",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetIssueType(val *string) {
	if err := j.validateSetIssueTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issueType",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetLabels(val *[]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetMilestone(val *string) {
	if err := j.validateSetMilestoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"milestone",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetMyReactionEmoji(val *string) {
	if err := j.validateSetMyReactionEmojiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"myReactionEmoji",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetNotAssigneeId(val *[]*float64) {
	if err := j.validateSetNotAssigneeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notAssigneeId",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetNotAuthorId(val *[]*float64) {
	if err := j.validateSetNotAuthorIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notAuthorId",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetNotLabels(val *[]*string) {
	if err := j.validateSetNotLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notLabels",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetNotMilestone(val *string) {
	if err := j.validateSetNotMilestoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notMilestone",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetNotMyReactionEmoji(val *[]*string) {
	if err := j.validateSetNotMyReactionEmojiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notMyReactionEmoji",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetOrderBy(val *string) {
	if err := j.validateSetOrderByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orderBy",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetSearch(val *string) {
	if err := j.validateSetSearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"search",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetSort(val *string) {
	if err := j.validateSetSortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sort",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetUpdatedAfter(val *string) {
	if err := j.validateSetUpdatedAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedAfter",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetUpdatedBefore(val *string) {
	if err := j.validateSetUpdatedBeforeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedBefore",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetWeight(val *float64) {
	if err := j.validateSetWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weight",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssues)SetWithLabelsDetails(val interface{}) {
	if err := j.validateSetWithLabelsDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"withLabelsDetails",
		val,
	)
}

// Generates CDKTF code for importing a DataGitlabProjectIssues resource upon running "cdktf plan <stack-name>".
func DataGitlabProjectIssues_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataGitlabProjectIssues_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.dataGitlabProjectIssues.DataGitlabProjectIssues",
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
func DataGitlabProjectIssues_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGitlabProjectIssues_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.dataGitlabProjectIssues.DataGitlabProjectIssues",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGitlabProjectIssues_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGitlabProjectIssues_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.dataGitlabProjectIssues.DataGitlabProjectIssues",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGitlabProjectIssues_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGitlabProjectIssues_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.dataGitlabProjectIssues.DataGitlabProjectIssues",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataGitlabProjectIssues_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.dataGitlabProjectIssues.DataGitlabProjectIssues",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssues) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssues) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssues) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssues) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssues) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssues) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssues) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssues) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssues) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssues) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssues) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetAssigneeId() {
	_jsii_.InvokeVoid(
		d,
		"resetAssigneeId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetAssigneeUsername() {
	_jsii_.InvokeVoid(
		d,
		"resetAssigneeUsername",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetAuthorId() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthorId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetConfidential() {
	_jsii_.InvokeVoid(
		d,
		"resetConfidential",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetCreatedAfter() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedAfter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetCreatedBefore() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedBefore",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetDueDate() {
	_jsii_.InvokeVoid(
		d,
		"resetDueDate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetIids() {
	_jsii_.InvokeVoid(
		d,
		"resetIids",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetIssueType() {
	_jsii_.InvokeVoid(
		d,
		"resetIssueType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetLabels() {
	_jsii_.InvokeVoid(
		d,
		"resetLabels",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetMilestone() {
	_jsii_.InvokeVoid(
		d,
		"resetMilestone",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetMyReactionEmoji() {
	_jsii_.InvokeVoid(
		d,
		"resetMyReactionEmoji",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetNotAssigneeId() {
	_jsii_.InvokeVoid(
		d,
		"resetNotAssigneeId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetNotAuthorId() {
	_jsii_.InvokeVoid(
		d,
		"resetNotAuthorId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetNotLabels() {
	_jsii_.InvokeVoid(
		d,
		"resetNotLabels",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetNotMilestone() {
	_jsii_.InvokeVoid(
		d,
		"resetNotMilestone",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetNotMyReactionEmoji() {
	_jsii_.InvokeVoid(
		d,
		"resetNotMyReactionEmoji",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetOrderBy() {
	_jsii_.InvokeVoid(
		d,
		"resetOrderBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetScope() {
	_jsii_.InvokeVoid(
		d,
		"resetScope",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetSearch() {
	_jsii_.InvokeVoid(
		d,
		"resetSearch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetSort() {
	_jsii_.InvokeVoid(
		d,
		"resetSort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetUpdatedAfter() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdatedAfter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetUpdatedBefore() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdatedBefore",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetWeight() {
	_jsii_.InvokeVoid(
		d,
		"resetWeight",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) ResetWithLabelsDetails() {
	_jsii_.InvokeVoid(
		d,
		"resetWithLabelsDetails",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssues) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssues) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssues) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssues) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssues) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssues) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

