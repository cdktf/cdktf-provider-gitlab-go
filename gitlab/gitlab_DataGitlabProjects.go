// Prebuilt gitlab Provider for Terraform CDK (cdktf)
package gitlab

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-gitlab-go/gitlab/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-gitlab-go/gitlab/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/gitlab/d/projects gitlab_projects}.
type DataGitlabProjects interface {
	cdktf.TerraformDataSource
	Archived() interface{}
	SetArchived(val interface{})
	ArchivedInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	GroupId() *float64
	SetGroupId(val *float64)
	GroupIdInput() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
	IncludeSubgroups() interface{}
	SetIncludeSubgroups(val interface{})
	IncludeSubgroupsInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxQueryablePages() *float64
	SetMaxQueryablePages(val *float64)
	MaxQueryablePagesInput() *float64
	Membership() interface{}
	SetMembership(val interface{})
	MembershipInput() interface{}
	MinAccessLevel() *float64
	SetMinAccessLevel(val *float64)
	MinAccessLevelInput() *float64
	// The tree node.
	Node() constructs.Node
	OrderBy() *string
	SetOrderBy(val *string)
	OrderByInput() *string
	Owned() interface{}
	SetOwned(val interface{})
	OwnedInput() interface{}
	Page() *float64
	SetPage(val *float64)
	PageInput() *float64
	PerPage() *float64
	SetPerPage(val *float64)
	PerPageInput() *float64
	Projects() DataGitlabProjectsProjectsList
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Search() *string
	SetSearch(val *string)
	SearchInput() *string
	Simple() interface{}
	SetSimple(val interface{})
	SimpleInput() interface{}
	Sort() *string
	SetSort(val *string)
	SortInput() *string
	Starred() interface{}
	SetStarred(val interface{})
	StarredInput() interface{}
	Statistics() interface{}
	SetStatistics(val interface{})
	StatisticsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Visibility() *string
	SetVisibility(val *string)
	VisibilityInput() *string
	WithCustomAttributes() interface{}
	SetWithCustomAttributes(val interface{})
	WithCustomAttributesInput() interface{}
	WithIssuesEnabled() interface{}
	SetWithIssuesEnabled(val interface{})
	WithIssuesEnabledInput() interface{}
	WithMergeRequestsEnabled() interface{}
	SetWithMergeRequestsEnabled(val interface{})
	WithMergeRequestsEnabledInput() interface{}
	WithProgrammingLanguage() *string
	SetWithProgrammingLanguage(val *string)
	WithProgrammingLanguageInput() *string
	WithShared() interface{}
	SetWithShared(val interface{})
	WithSharedInput() interface{}
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
	ResetArchived()
	ResetGroupId()
	ResetId()
	ResetIncludeSubgroups()
	ResetMaxQueryablePages()
	ResetMembership()
	ResetMinAccessLevel()
	ResetOrderBy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOwned()
	ResetPage()
	ResetPerPage()
	ResetSearch()
	ResetSimple()
	ResetSort()
	ResetStarred()
	ResetStatistics()
	ResetVisibility()
	ResetWithCustomAttributes()
	ResetWithIssuesEnabled()
	ResetWithMergeRequestsEnabled()
	ResetWithProgrammingLanguage()
	ResetWithShared()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataGitlabProjects
type jsiiProxy_DataGitlabProjects struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataGitlabProjects) Archived() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archived",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) ArchivedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archivedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) GroupId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) GroupIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) IncludeSubgroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSubgroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) IncludeSubgroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSubgroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) MaxQueryablePages() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxQueryablePages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) MaxQueryablePagesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxQueryablePagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) Membership() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"membership",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) MembershipInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"membershipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) MinAccessLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) MinAccessLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minAccessLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) OrderBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) OrderByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) Owned() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"owned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) OwnedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ownedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) Page() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"page",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) PageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) PerPage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"perPage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) PerPageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"perPageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) Projects() DataGitlabProjectsProjectsList {
	var returns DataGitlabProjectsProjectsList
	_jsii_.Get(
		j,
		"projects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) Search() *string {
	var returns *string
	_jsii_.Get(
		j,
		"search",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) SearchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) Simple() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"simple",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) SimpleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"simpleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) Sort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) SortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) Starred() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"starred",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) StarredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"starredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) Statistics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statistics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) StatisticsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statisticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) Visibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) VisibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) WithCustomAttributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withCustomAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) WithCustomAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withCustomAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) WithIssuesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withIssuesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) WithIssuesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withIssuesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) WithMergeRequestsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withMergeRequestsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) WithMergeRequestsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withMergeRequestsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) WithProgrammingLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"withProgrammingLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) WithProgrammingLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"withProgrammingLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) WithShared() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withShared",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjects) WithSharedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withSharedInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/gitlab/d/projects gitlab_projects} Data Source.
func NewDataGitlabProjects(scope constructs.Construct, id *string, config *DataGitlabProjectsConfig) DataGitlabProjects {
	_init_.Initialize()

	j := jsiiProxy_DataGitlabProjects{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.DataGitlabProjects",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/gitlab/d/projects gitlab_projects} Data Source.
func NewDataGitlabProjects_Override(d DataGitlabProjects, scope constructs.Construct, id *string, config *DataGitlabProjectsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.DataGitlabProjects",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetArchived(val interface{}) {
	_jsii_.Set(
		j,
		"archived",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetGroupId(val *float64) {
	_jsii_.Set(
		j,
		"groupId",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetIncludeSubgroups(val interface{}) {
	_jsii_.Set(
		j,
		"includeSubgroups",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetMaxQueryablePages(val *float64) {
	_jsii_.Set(
		j,
		"maxQueryablePages",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetMembership(val interface{}) {
	_jsii_.Set(
		j,
		"membership",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetMinAccessLevel(val *float64) {
	_jsii_.Set(
		j,
		"minAccessLevel",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetOrderBy(val *string) {
	_jsii_.Set(
		j,
		"orderBy",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetOwned(val interface{}) {
	_jsii_.Set(
		j,
		"owned",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetPage(val *float64) {
	_jsii_.Set(
		j,
		"page",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetPerPage(val *float64) {
	_jsii_.Set(
		j,
		"perPage",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetSearch(val *string) {
	_jsii_.Set(
		j,
		"search",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetSimple(val interface{}) {
	_jsii_.Set(
		j,
		"simple",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetSort(val *string) {
	_jsii_.Set(
		j,
		"sort",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetStarred(val interface{}) {
	_jsii_.Set(
		j,
		"starred",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetStatistics(val interface{}) {
	_jsii_.Set(
		j,
		"statistics",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetVisibility(val *string) {
	_jsii_.Set(
		j,
		"visibility",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetWithCustomAttributes(val interface{}) {
	_jsii_.Set(
		j,
		"withCustomAttributes",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetWithIssuesEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"withIssuesEnabled",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetWithMergeRequestsEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"withMergeRequestsEnabled",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetWithProgrammingLanguage(val *string) {
	_jsii_.Set(
		j,
		"withProgrammingLanguage",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjects) SetWithShared(val interface{}) {
	_jsii_.Set(
		j,
		"withShared",
		val,
	)
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
func DataGitlabProjects_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.DataGitlabProjects",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataGitlabProjects_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.DataGitlabProjects",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataGitlabProjects) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataGitlabProjects) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjects) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjects) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjects) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjects) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjects) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjects) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjects) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjects) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjects) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjects) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataGitlabProjects) ResetArchived() {
	_jsii_.InvokeVoid(
		d,
		"resetArchived",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjects) ResetGroupId() {
	_jsii_.InvokeVoid(
		d,
		"resetGroupId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjects) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjects) ResetIncludeSubgroups() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeSubgroups",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjects) ResetMaxQueryablePages() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxQueryablePages",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjects) ResetMembership() {
	_jsii_.InvokeVoid(
		d,
		"resetMembership",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjects) ResetMinAccessLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetMinAccessLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjects) ResetOrderBy() {
	_jsii_.InvokeVoid(
		d,
		"resetOrderBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjects) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjects) ResetOwned() {
	_jsii_.InvokeVoid(
		d,
		"resetOwned",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjects) ResetPage() {
	_jsii_.InvokeVoid(
		d,
		"resetPage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjects) ResetPerPage() {
	_jsii_.InvokeVoid(
		d,
		"resetPerPage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjects) ResetSearch() {
	_jsii_.InvokeVoid(
		d,
		"resetSearch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjects) ResetSimple() {
	_jsii_.InvokeVoid(
		d,
		"resetSimple",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjects) ResetSort() {
	_jsii_.InvokeVoid(
		d,
		"resetSort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjects) ResetStarred() {
	_jsii_.InvokeVoid(
		d,
		"resetStarred",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjects) ResetStatistics() {
	_jsii_.InvokeVoid(
		d,
		"resetStatistics",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjects) ResetVisibility() {
	_jsii_.InvokeVoid(
		d,
		"resetVisibility",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjects) ResetWithCustomAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetWithCustomAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjects) ResetWithIssuesEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetWithIssuesEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjects) ResetWithMergeRequestsEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetWithMergeRequestsEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjects) ResetWithProgrammingLanguage() {
	_jsii_.InvokeVoid(
		d,
		"resetWithProgrammingLanguage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjects) ResetWithShared() {
	_jsii_.InvokeVoid(
		d,
		"resetWithShared",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjects) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjects) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjects) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjects) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
