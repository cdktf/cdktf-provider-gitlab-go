// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabprojectmergerequests

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v15/datagitlabprojectmergerequests/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/data-sources/project_merge_requests gitlab_project_merge_requests}.
type DataGitlabProjectMergeRequests interface {
	cdktf.TerraformDataSource
	AuthorId() *float64
	SetAuthorId(val *float64)
	AuthorIdInput() *float64
	AuthorUsername() *string
	SetAuthorUsername(val *string)
	AuthorUsernameInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Iids() *[]*float64
	SetIids(val *[]*float64)
	IidsInput() *[]*float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MergeRequests() DataGitlabProjectMergeRequestsMergeRequestsList
	Milestone() *string
	SetMilestone(val *string)
	MilestoneInput() *string
	MyReactionEmoji() *string
	SetMyReactionEmoji(val *string)
	MyReactionEmojiInput() *string
	// The tree node.
	Node() constructs.Node
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
	ReviewerUsername() *string
	SetReviewerUsername(val *string)
	ReviewerUsernameInput() *string
	Scope() *string
	SetScope(val *string)
	ScopeInput() *string
	Search() *string
	SetSearch(val *string)
	SearchInput() *string
	Sort() *string
	SetSort(val *string)
	SortInput() *string
	SourceBranch() *string
	SetSourceBranch(val *string)
	SourceBranchInput() *string
	State() *string
	SetState(val *string)
	StateInput() *string
	TargetBranch() *string
	SetTargetBranch(val *string)
	TargetBranchInput() *string
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
	Wip() *string
	SetWip(val *string)
	WipInput() *string
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
	ResetAuthorId()
	ResetAuthorUsername()
	ResetCreatedAfter()
	ResetCreatedBefore()
	ResetIids()
	ResetMilestone()
	ResetMyReactionEmoji()
	ResetOrderBy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReviewerUsername()
	ResetScope()
	ResetSearch()
	ResetSort()
	ResetSourceBranch()
	ResetState()
	ResetTargetBranch()
	ResetUpdatedAfter()
	ResetUpdatedBefore()
	ResetWip()
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

// The jsii proxy struct for DataGitlabProjectMergeRequests
type jsiiProxy_DataGitlabProjectMergeRequests struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) AuthorId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) AuthorIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) AuthorUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) AuthorUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) CreatedAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) CreatedAfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) CreatedBefore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) CreatedBeforeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBeforeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) Iids() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"iids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) IidsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"iidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) MergeRequests() DataGitlabProjectMergeRequestsMergeRequestsList {
	var returns DataGitlabProjectMergeRequestsMergeRequestsList
	_jsii_.Get(
		j,
		"mergeRequests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) Milestone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"milestone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) MilestoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"milestoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) MyReactionEmoji() *string {
	var returns *string
	_jsii_.Get(
		j,
		"myReactionEmoji",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) MyReactionEmojiInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"myReactionEmojiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) OrderBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) OrderByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) ReviewerUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reviewerUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) ReviewerUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reviewerUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) Search() *string {
	var returns *string
	_jsii_.Get(
		j,
		"search",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) SearchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) Sort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) SortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) SourceBranch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) SourceBranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceBranchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) TargetBranch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) TargetBranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetBranchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) UpdatedAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) UpdatedAfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) UpdatedBefore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) UpdatedBeforeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedBeforeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) Wip() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests) WipInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wipInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/data-sources/project_merge_requests gitlab_project_merge_requests} Data Source.
func NewDataGitlabProjectMergeRequests(scope constructs.Construct, id *string, config *DataGitlabProjectMergeRequestsConfig) DataGitlabProjectMergeRequests {
	_init_.Initialize()

	if err := validateNewDataGitlabProjectMergeRequestsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGitlabProjectMergeRequests{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.dataGitlabProjectMergeRequests.DataGitlabProjectMergeRequests",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/data-sources/project_merge_requests gitlab_project_merge_requests} Data Source.
func NewDataGitlabProjectMergeRequests_Override(d DataGitlabProjectMergeRequests, scope constructs.Construct, id *string, config *DataGitlabProjectMergeRequestsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.dataGitlabProjectMergeRequests.DataGitlabProjectMergeRequests",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests)SetAuthorId(val *float64) {
	if err := j.validateSetAuthorIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorId",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests)SetAuthorUsername(val *string) {
	if err := j.validateSetAuthorUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorUsername",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests)SetCreatedAfter(val *string) {
	if err := j.validateSetCreatedAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdAfter",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests)SetCreatedBefore(val *string) {
	if err := j.validateSetCreatedBeforeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdBefore",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests)SetIids(val *[]*float64) {
	if err := j.validateSetIidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iids",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests)SetMilestone(val *string) {
	if err := j.validateSetMilestoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"milestone",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests)SetMyReactionEmoji(val *string) {
	if err := j.validateSetMyReactionEmojiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"myReactionEmoji",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests)SetOrderBy(val *string) {
	if err := j.validateSetOrderByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orderBy",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests)SetReviewerUsername(val *string) {
	if err := j.validateSetReviewerUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reviewerUsername",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests)SetSearch(val *string) {
	if err := j.validateSetSearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"search",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests)SetSort(val *string) {
	if err := j.validateSetSortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sort",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests)SetSourceBranch(val *string) {
	if err := j.validateSetSourceBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceBranch",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests)SetTargetBranch(val *string) {
	if err := j.validateSetTargetBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetBranch",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests)SetUpdatedAfter(val *string) {
	if err := j.validateSetUpdatedAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedAfter",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests)SetUpdatedBefore(val *string) {
	if err := j.validateSetUpdatedBeforeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedBefore",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectMergeRequests)SetWip(val *string) {
	if err := j.validateSetWipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wip",
		val,
	)
}

// Generates CDKTF code for importing a DataGitlabProjectMergeRequests resource upon running "cdktf plan <stack-name>".
func DataGitlabProjectMergeRequests_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataGitlabProjectMergeRequests_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.dataGitlabProjectMergeRequests.DataGitlabProjectMergeRequests",
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
func DataGitlabProjectMergeRequests_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGitlabProjectMergeRequests_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.dataGitlabProjectMergeRequests.DataGitlabProjectMergeRequests",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGitlabProjectMergeRequests_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGitlabProjectMergeRequests_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.dataGitlabProjectMergeRequests.DataGitlabProjectMergeRequests",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGitlabProjectMergeRequests_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGitlabProjectMergeRequests_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.dataGitlabProjectMergeRequests.DataGitlabProjectMergeRequests",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataGitlabProjectMergeRequests_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.dataGitlabProjectMergeRequests.DataGitlabProjectMergeRequests",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGitlabProjectMergeRequests) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGitlabProjectMergeRequests) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGitlabProjectMergeRequests) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGitlabProjectMergeRequests) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGitlabProjectMergeRequests) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGitlabProjectMergeRequests) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGitlabProjectMergeRequests) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGitlabProjectMergeRequests) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGitlabProjectMergeRequests) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGitlabProjectMergeRequests) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) ResetAuthorId() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthorId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) ResetAuthorUsername() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthorUsername",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) ResetCreatedAfter() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedAfter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) ResetCreatedBefore() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedBefore",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) ResetIids() {
	_jsii_.InvokeVoid(
		d,
		"resetIids",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) ResetMilestone() {
	_jsii_.InvokeVoid(
		d,
		"resetMilestone",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) ResetMyReactionEmoji() {
	_jsii_.InvokeVoid(
		d,
		"resetMyReactionEmoji",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) ResetOrderBy() {
	_jsii_.InvokeVoid(
		d,
		"resetOrderBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) ResetReviewerUsername() {
	_jsii_.InvokeVoid(
		d,
		"resetReviewerUsername",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) ResetScope() {
	_jsii_.InvokeVoid(
		d,
		"resetScope",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) ResetSearch() {
	_jsii_.InvokeVoid(
		d,
		"resetSearch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) ResetSort() {
	_jsii_.InvokeVoid(
		d,
		"resetSort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) ResetSourceBranch() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceBranch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) ResetState() {
	_jsii_.InvokeVoid(
		d,
		"resetState",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) ResetTargetBranch() {
	_jsii_.InvokeVoid(
		d,
		"resetTargetBranch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) ResetUpdatedAfter() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdatedAfter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) ResetUpdatedBefore() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdatedBefore",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) ResetWip() {
	_jsii_.InvokeVoid(
		d,
		"resetWip",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectMergeRequests) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

