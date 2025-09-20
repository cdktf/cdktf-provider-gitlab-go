// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grouplevelmrapprovals

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v15/grouplevelmrapprovals/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/group_level_mr_approvals gitlab_group_level_mr_approvals}.
type GroupLevelMrApprovals interface {
	cdktf.TerraformResource
	AllowAuthorApproval() interface{}
	SetAllowAuthorApproval(val interface{})
	AllowAuthorApprovalInput() interface{}
	AllowCommitterApproval() interface{}
	SetAllowCommitterApproval(val interface{})
	AllowCommitterApprovalInput() interface{}
	AllowOverridesToApproverListPerMergeRequest() interface{}
	SetAllowOverridesToApproverListPerMergeRequest(val interface{})
	AllowOverridesToApproverListPerMergeRequestInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	Group() *string
	SetGroup(val *string)
	GroupInput() *string
	Id() *string
	KeepSettingsOnDestroy() interface{}
	SetKeepSettingsOnDestroy(val interface{})
	KeepSettingsOnDestroyInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	// Experimental.
	RawOverrides() interface{}
	RequireReauthenticationToApprove() interface{}
	SetRequireReauthenticationToApprove(val interface{})
	RequireReauthenticationToApproveInput() interface{}
	RetainApprovalsOnPush() interface{}
	SetRetainApprovalsOnPush(val interface{})
	RetainApprovalsOnPushInput() interface{}
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
	ResetAllowAuthorApproval()
	ResetAllowCommitterApproval()
	ResetAllowOverridesToApproverListPerMergeRequest()
	ResetKeepSettingsOnDestroy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRequireReauthenticationToApprove()
	ResetRetainApprovalsOnPush()
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

// The jsii proxy struct for GroupLevelMrApprovals
type jsiiProxy_GroupLevelMrApprovals struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GroupLevelMrApprovals) AllowAuthorApproval() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAuthorApproval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) AllowAuthorApprovalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAuthorApprovalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) AllowCommitterApproval() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCommitterApproval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) AllowCommitterApprovalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCommitterApprovalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) AllowOverridesToApproverListPerMergeRequest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowOverridesToApproverListPerMergeRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) AllowOverridesToApproverListPerMergeRequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowOverridesToApproverListPerMergeRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) Group() *string {
	var returns *string
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) GroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) KeepSettingsOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepSettingsOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) KeepSettingsOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepSettingsOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) RequireReauthenticationToApprove() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireReauthenticationToApprove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) RequireReauthenticationToApproveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireReauthenticationToApproveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) RetainApprovalsOnPush() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainApprovalsOnPush",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) RetainApprovalsOnPushInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainApprovalsOnPushInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupLevelMrApprovals) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/group_level_mr_approvals gitlab_group_level_mr_approvals} Resource.
func NewGroupLevelMrApprovals(scope constructs.Construct, id *string, config *GroupLevelMrApprovalsConfig) GroupLevelMrApprovals {
	_init_.Initialize()

	if err := validateNewGroupLevelMrApprovalsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GroupLevelMrApprovals{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.groupLevelMrApprovals.GroupLevelMrApprovals",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/group_level_mr_approvals gitlab_group_level_mr_approvals} Resource.
func NewGroupLevelMrApprovals_Override(g GroupLevelMrApprovals, scope constructs.Construct, id *string, config *GroupLevelMrApprovalsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.groupLevelMrApprovals.GroupLevelMrApprovals",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GroupLevelMrApprovals)SetAllowAuthorApproval(val interface{}) {
	if err := j.validateSetAllowAuthorApprovalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowAuthorApproval",
		val,
	)
}

func (j *jsiiProxy_GroupLevelMrApprovals)SetAllowCommitterApproval(val interface{}) {
	if err := j.validateSetAllowCommitterApprovalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowCommitterApproval",
		val,
	)
}

func (j *jsiiProxy_GroupLevelMrApprovals)SetAllowOverridesToApproverListPerMergeRequest(val interface{}) {
	if err := j.validateSetAllowOverridesToApproverListPerMergeRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowOverridesToApproverListPerMergeRequest",
		val,
	)
}

func (j *jsiiProxy_GroupLevelMrApprovals)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GroupLevelMrApprovals)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GroupLevelMrApprovals)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GroupLevelMrApprovals)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GroupLevelMrApprovals)SetGroup(val *string) {
	if err := j.validateSetGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"group",
		val,
	)
}

func (j *jsiiProxy_GroupLevelMrApprovals)SetKeepSettingsOnDestroy(val interface{}) {
	if err := j.validateSetKeepSettingsOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepSettingsOnDestroy",
		val,
	)
}

func (j *jsiiProxy_GroupLevelMrApprovals)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GroupLevelMrApprovals)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GroupLevelMrApprovals)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GroupLevelMrApprovals)SetRequireReauthenticationToApprove(val interface{}) {
	if err := j.validateSetRequireReauthenticationToApproveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireReauthenticationToApprove",
		val,
	)
}

func (j *jsiiProxy_GroupLevelMrApprovals)SetRetainApprovalsOnPush(val interface{}) {
	if err := j.validateSetRetainApprovalsOnPushParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retainApprovalsOnPush",
		val,
	)
}

// Generates CDKTF code for importing a GroupLevelMrApprovals resource upon running "cdktf plan <stack-name>".
func GroupLevelMrApprovals_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGroupLevelMrApprovals_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.groupLevelMrApprovals.GroupLevelMrApprovals",
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
func GroupLevelMrApprovals_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGroupLevelMrApprovals_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.groupLevelMrApprovals.GroupLevelMrApprovals",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GroupLevelMrApprovals_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGroupLevelMrApprovals_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.groupLevelMrApprovals.GroupLevelMrApprovals",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GroupLevelMrApprovals_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGroupLevelMrApprovals_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.groupLevelMrApprovals.GroupLevelMrApprovals",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GroupLevelMrApprovals_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.groupLevelMrApprovals.GroupLevelMrApprovals",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GroupLevelMrApprovals) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GroupLevelMrApprovals) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GroupLevelMrApprovals) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GroupLevelMrApprovals) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GroupLevelMrApprovals) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GroupLevelMrApprovals) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GroupLevelMrApprovals) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GroupLevelMrApprovals) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GroupLevelMrApprovals) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GroupLevelMrApprovals) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GroupLevelMrApprovals) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GroupLevelMrApprovals) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupLevelMrApprovals) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GroupLevelMrApprovals) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GroupLevelMrApprovals) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GroupLevelMrApprovals) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GroupLevelMrApprovals) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GroupLevelMrApprovals) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GroupLevelMrApprovals) ResetAllowAuthorApproval() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowAuthorApproval",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupLevelMrApprovals) ResetAllowCommitterApproval() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowCommitterApproval",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupLevelMrApprovals) ResetAllowOverridesToApproverListPerMergeRequest() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowOverridesToApproverListPerMergeRequest",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupLevelMrApprovals) ResetKeepSettingsOnDestroy() {
	_jsii_.InvokeVoid(
		g,
		"resetKeepSettingsOnDestroy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupLevelMrApprovals) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupLevelMrApprovals) ResetRequireReauthenticationToApprove() {
	_jsii_.InvokeVoid(
		g,
		"resetRequireReauthenticationToApprove",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupLevelMrApprovals) ResetRetainApprovalsOnPush() {
	_jsii_.InvokeVoid(
		g,
		"resetRetainApprovalsOnPush",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupLevelMrApprovals) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupLevelMrApprovals) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupLevelMrApprovals) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupLevelMrApprovals) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupLevelMrApprovals) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupLevelMrApprovals) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

