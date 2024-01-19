// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabprojectissue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v12/datagitlabprojectissue/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/data-sources/project_issue gitlab_project_issue}.
type DataGitlabProjectIssue interface {
	cdktf.TerraformDataSource
	AssigneeIds() *[]*float64
	AuthorId() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClosedAt() *string
	ClosedByUserId() *float64
	Confidential() cdktf.IResolvable
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAt() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	DiscussionLocked() cdktf.IResolvable
	DiscussionToResolve() *string
	Downvotes() *float64
	DueDate() *string
	EpicId() *float64
	EpicIssueId() *float64
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
	Labels() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Links() cdktf.StringMap
	MergeRequestsCount() *float64
	MergeRequestToResolveDiscussionsOf() *float64
	MilestoneId() *float64
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
	RawOverrides() interface{}
	References() cdktf.StringMap
	State() *string
	Subscribed() cdktf.IResolvable
	TaskCompletionStatus() DataGitlabProjectIssueTaskCompletionStatusList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeEstimate() *float64
	Title() *string
	TotalTimeSpent() *float64
	UpdatedAt() *string
	Upvotes() *float64
	UserNotesCount() *float64
	WebUrl() *string
	Weight() *float64
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
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for DataGitlabProjectIssue
type jsiiProxy_DataGitlabProjectIssue struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataGitlabProjectIssue) AssigneeIds() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"assigneeIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) AuthorId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) ClosedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"closedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) ClosedByUserId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"closedByUserId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) Confidential() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"confidential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) DiscussionLocked() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"discussionLocked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) DiscussionToResolve() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discussionToResolve",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) Downvotes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"downvotes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) DueDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dueDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) EpicId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"epicId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) EpicIssueId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"epicIssueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) HumanTimeEstimate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"humanTimeEstimate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) HumanTotalTimeSpent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"humanTotalTimeSpent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) Iid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) IidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) IssueId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"issueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) IssueLinkId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"issueLinkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) IssueType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issueType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) Labels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) Links() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"links",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) MergeRequestsCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mergeRequestsCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) MergeRequestToResolveDiscussionsOf() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mergeRequestToResolveDiscussionsOf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) MilestoneId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"milestoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) MovedToId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"movedToId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) References() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"references",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) Subscribed() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"subscribed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) TaskCompletionStatus() DataGitlabProjectIssueTaskCompletionStatusList {
	var returns DataGitlabProjectIssueTaskCompletionStatusList
	_jsii_.Get(
		j,
		"taskCompletionStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) TimeEstimate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeEstimate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) TotalTimeSpent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalTimeSpent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) Upvotes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"upvotes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) UserNotesCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userNotesCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) WebUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssue) Weight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weight",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/data-sources/project_issue gitlab_project_issue} Data Source.
func NewDataGitlabProjectIssue(scope constructs.Construct, id *string, config *DataGitlabProjectIssueConfig) DataGitlabProjectIssue {
	_init_.Initialize()

	if err := validateNewDataGitlabProjectIssueParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGitlabProjectIssue{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.dataGitlabProjectIssue.DataGitlabProjectIssue",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.8.0/docs/data-sources/project_issue gitlab_project_issue} Data Source.
func NewDataGitlabProjectIssue_Override(d DataGitlabProjectIssue, scope constructs.Construct, id *string, config *DataGitlabProjectIssueConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.dataGitlabProjectIssue.DataGitlabProjectIssue",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssue)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssue)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssue)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssue)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssue)SetIid(val *float64) {
	if err := j.validateSetIidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iid",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssue)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssue)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssue)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTF code for importing a DataGitlabProjectIssue resource upon running "cdktf plan <stack-name>".
func DataGitlabProjectIssue_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataGitlabProjectIssue_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.dataGitlabProjectIssue.DataGitlabProjectIssue",
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
func DataGitlabProjectIssue_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGitlabProjectIssue_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.dataGitlabProjectIssue.DataGitlabProjectIssue",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGitlabProjectIssue_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGitlabProjectIssue_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.dataGitlabProjectIssue.DataGitlabProjectIssue",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGitlabProjectIssue_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGitlabProjectIssue_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.dataGitlabProjectIssue.DataGitlabProjectIssue",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataGitlabProjectIssue_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.dataGitlabProjectIssue.DataGitlabProjectIssue",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssue) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataGitlabProjectIssue) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGitlabProjectIssue) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGitlabProjectIssue) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGitlabProjectIssue) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGitlabProjectIssue) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGitlabProjectIssue) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGitlabProjectIssue) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGitlabProjectIssue) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGitlabProjectIssue) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGitlabProjectIssue) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGitlabProjectIssue) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataGitlabProjectIssue) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssue) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectIssue) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssue) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssue) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssue) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssue) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssue) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

