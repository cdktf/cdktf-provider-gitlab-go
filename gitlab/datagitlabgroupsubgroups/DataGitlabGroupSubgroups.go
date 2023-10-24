// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabgroupsubgroups

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v11/datagitlabgroupsubgroups/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.5.0/docs/data-sources/group_subgroups gitlab_group_subgroups}.
type DataGitlabGroupSubgroups interface {
	cdktf.TerraformDataSource
	AllAvailable() interface{}
	SetAllAvailable(val interface{})
	AllAvailableInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	GroupId() *float64
	SetGroupId(val *float64)
	GroupIdInput() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MinAccessLevel() *string
	SetMinAccessLevel(val *string)
	MinAccessLevelInput() *string
	// The tree node.
	Node() constructs.Node
	OrderBy() *string
	SetOrderBy(val *string)
	OrderByInput() *string
	Owned() interface{}
	SetOwned(val interface{})
	OwnedInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Search() *string
	SetSearch(val *string)
	SearchInput() *string
	SkipGroups() *[]*float64
	SetSkipGroups(val *[]*float64)
	SkipGroupsInput() *[]*float64
	Sort() *string
	SetSort(val *string)
	SortInput() *string
	Statistics() interface{}
	SetStatistics(val interface{})
	StatisticsInput() interface{}
	Subgroups() DataGitlabGroupSubgroupsSubgroupsList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WithCustomAttributes() interface{}
	SetWithCustomAttributes(val interface{})
	WithCustomAttributesInput() interface{}
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
	ResetAllAvailable()
	ResetId()
	ResetMinAccessLevel()
	ResetOrderBy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOwned()
	ResetSearch()
	ResetSkipGroups()
	ResetSort()
	ResetStatistics()
	ResetWithCustomAttributes()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataGitlabGroupSubgroups
type jsiiProxy_DataGitlabGroupSubgroups struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) AllAvailable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) AllAvailableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allAvailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) GroupId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) GroupIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) MinAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) MinAccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minAccessLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) OrderBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) OrderByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) Owned() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"owned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) OwnedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ownedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) Search() *string {
	var returns *string
	_jsii_.Get(
		j,
		"search",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) SearchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) SkipGroups() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"skipGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) SkipGroupsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"skipGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) Sort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) SortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) Statistics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statistics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) StatisticsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statisticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) Subgroups() DataGitlabGroupSubgroupsSubgroupsList {
	var returns DataGitlabGroupSubgroupsSubgroupsList
	_jsii_.Get(
		j,
		"subgroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) WithCustomAttributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withCustomAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabGroupSubgroups) WithCustomAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withCustomAttributesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.5.0/docs/data-sources/group_subgroups gitlab_group_subgroups} Data Source.
func NewDataGitlabGroupSubgroups(scope constructs.Construct, id *string, config *DataGitlabGroupSubgroupsConfig) DataGitlabGroupSubgroups {
	_init_.Initialize()

	if err := validateNewDataGitlabGroupSubgroupsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGitlabGroupSubgroups{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.dataGitlabGroupSubgroups.DataGitlabGroupSubgroups",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.5.0/docs/data-sources/group_subgroups gitlab_group_subgroups} Data Source.
func NewDataGitlabGroupSubgroups_Override(d DataGitlabGroupSubgroups, scope constructs.Construct, id *string, config *DataGitlabGroupSubgroupsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.dataGitlabGroupSubgroups.DataGitlabGroupSubgroups",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataGitlabGroupSubgroups)SetAllAvailable(val interface{}) {
	if err := j.validateSetAllAvailableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allAvailable",
		val,
	)
}

func (j *jsiiProxy_DataGitlabGroupSubgroups)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataGitlabGroupSubgroups)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataGitlabGroupSubgroups)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataGitlabGroupSubgroups)SetGroupId(val *float64) {
	if err := j.validateSetGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupId",
		val,
	)
}

func (j *jsiiProxy_DataGitlabGroupSubgroups)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataGitlabGroupSubgroups)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataGitlabGroupSubgroups)SetMinAccessLevel(val *string) {
	if err := j.validateSetMinAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minAccessLevel",
		val,
	)
}

func (j *jsiiProxy_DataGitlabGroupSubgroups)SetOrderBy(val *string) {
	if err := j.validateSetOrderByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orderBy",
		val,
	)
}

func (j *jsiiProxy_DataGitlabGroupSubgroups)SetOwned(val interface{}) {
	if err := j.validateSetOwnedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"owned",
		val,
	)
}

func (j *jsiiProxy_DataGitlabGroupSubgroups)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataGitlabGroupSubgroups)SetSearch(val *string) {
	if err := j.validateSetSearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"search",
		val,
	)
}

func (j *jsiiProxy_DataGitlabGroupSubgroups)SetSkipGroups(val *[]*float64) {
	if err := j.validateSetSkipGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipGroups",
		val,
	)
}

func (j *jsiiProxy_DataGitlabGroupSubgroups)SetSort(val *string) {
	if err := j.validateSetSortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sort",
		val,
	)
}

func (j *jsiiProxy_DataGitlabGroupSubgroups)SetStatistics(val interface{}) {
	if err := j.validateSetStatisticsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statistics",
		val,
	)
}

func (j *jsiiProxy_DataGitlabGroupSubgroups)SetWithCustomAttributes(val interface{}) {
	if err := j.validateSetWithCustomAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"withCustomAttributes",
		val,
	)
}

// Generates CDKTF code for importing a DataGitlabGroupSubgroups resource upon running "cdktf plan <stack-name>".
func DataGitlabGroupSubgroups_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataGitlabGroupSubgroups_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.dataGitlabGroupSubgroups.DataGitlabGroupSubgroups",
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
func DataGitlabGroupSubgroups_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGitlabGroupSubgroups_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.dataGitlabGroupSubgroups.DataGitlabGroupSubgroups",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGitlabGroupSubgroups_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGitlabGroupSubgroups_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.dataGitlabGroupSubgroups.DataGitlabGroupSubgroups",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGitlabGroupSubgroups_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGitlabGroupSubgroups_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.dataGitlabGroupSubgroups.DataGitlabGroupSubgroups",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataGitlabGroupSubgroups_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.dataGitlabGroupSubgroups.DataGitlabGroupSubgroups",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataGitlabGroupSubgroups) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataGitlabGroupSubgroups) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGitlabGroupSubgroups) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGitlabGroupSubgroups) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGitlabGroupSubgroups) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGitlabGroupSubgroups) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGitlabGroupSubgroups) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGitlabGroupSubgroups) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGitlabGroupSubgroups) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGitlabGroupSubgroups) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGitlabGroupSubgroups) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGitlabGroupSubgroups) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataGitlabGroupSubgroups) ResetAllAvailable() {
	_jsii_.InvokeVoid(
		d,
		"resetAllAvailable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabGroupSubgroups) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabGroupSubgroups) ResetMinAccessLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetMinAccessLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabGroupSubgroups) ResetOrderBy() {
	_jsii_.InvokeVoid(
		d,
		"resetOrderBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabGroupSubgroups) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabGroupSubgroups) ResetOwned() {
	_jsii_.InvokeVoid(
		d,
		"resetOwned",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabGroupSubgroups) ResetSearch() {
	_jsii_.InvokeVoid(
		d,
		"resetSearch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabGroupSubgroups) ResetSkipGroups() {
	_jsii_.InvokeVoid(
		d,
		"resetSkipGroups",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabGroupSubgroups) ResetSort() {
	_jsii_.InvokeVoid(
		d,
		"resetSort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabGroupSubgroups) ResetStatistics() {
	_jsii_.InvokeVoid(
		d,
		"resetStatistics",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabGroupSubgroups) ResetWithCustomAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetWithCustomAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabGroupSubgroups) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabGroupSubgroups) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabGroupSubgroups) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabGroupSubgroups) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

