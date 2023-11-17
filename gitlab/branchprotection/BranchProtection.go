// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package branchprotection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v11/branchprotection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.6.0/docs/resources/branch_protection gitlab_branch_protection}.
type BranchProtection interface {
	cdktf.TerraformResource
	AllowedToMerge() BranchProtectionAllowedToMergeList
	AllowedToMergeInput() interface{}
	AllowedToPush() BranchProtectionAllowedToPushList
	AllowedToPushInput() interface{}
	AllowedToUnprotect() BranchProtectionAllowedToUnprotectList
	AllowedToUnprotectInput() interface{}
	AllowForcePush() interface{}
	SetAllowForcePush(val interface{})
	AllowForcePushInput() interface{}
	Branch() *string
	SetBranch(val *string)
	BranchInput() *string
	BranchProtectionId() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CodeOwnerApprovalRequired() interface{}
	SetCodeOwnerApprovalRequired(val interface{})
	CodeOwnerApprovalRequiredInput() interface{}
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MergeAccessLevel() *string
	SetMergeAccessLevel(val *string)
	MergeAccessLevelInput() *string
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
	PushAccessLevel() *string
	SetPushAccessLevel(val *string)
	PushAccessLevelInput() *string
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UnprotectAccessLevel() *string
	SetUnprotectAccessLevel(val *string)
	UnprotectAccessLevelInput() *string
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAllowedToMerge(value interface{})
	PutAllowedToPush(value interface{})
	PutAllowedToUnprotect(value interface{})
	ResetAllowedToMerge()
	ResetAllowedToPush()
	ResetAllowedToUnprotect()
	ResetAllowForcePush()
	ResetCodeOwnerApprovalRequired()
	ResetId()
	ResetMergeAccessLevel()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPushAccessLevel()
	ResetUnprotectAccessLevel()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for BranchProtection
type jsiiProxy_BranchProtection struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BranchProtection) AllowedToMerge() BranchProtectionAllowedToMergeList {
	var returns BranchProtectionAllowedToMergeList
	_jsii_.Get(
		j,
		"allowedToMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) AllowedToMergeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedToMergeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) AllowedToPush() BranchProtectionAllowedToPushList {
	var returns BranchProtectionAllowedToPushList
	_jsii_.Get(
		j,
		"allowedToPush",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) AllowedToPushInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedToPushInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) AllowedToUnprotect() BranchProtectionAllowedToUnprotectList {
	var returns BranchProtectionAllowedToUnprotectList
	_jsii_.Get(
		j,
		"allowedToUnprotect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) AllowedToUnprotectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedToUnprotectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) AllowForcePush() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowForcePush",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) AllowForcePushInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowForcePushInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) Branch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) BranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) BranchProtectionId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"branchProtectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) CodeOwnerApprovalRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeOwnerApprovalRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) CodeOwnerApprovalRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeOwnerApprovalRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) MergeAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) MergeAccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeAccessLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) PushAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pushAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) PushAccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pushAccessLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) UnprotectAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unprotectAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) UnprotectAccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unprotectAccessLevelInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.6.0/docs/resources/branch_protection gitlab_branch_protection} Resource.
func NewBranchProtection(scope constructs.Construct, id *string, config *BranchProtectionConfig) BranchProtection {
	_init_.Initialize()

	if err := validateNewBranchProtectionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BranchProtection{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.branchProtection.BranchProtection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.6.0/docs/resources/branch_protection gitlab_branch_protection} Resource.
func NewBranchProtection_Override(b BranchProtection, scope constructs.Construct, id *string, config *BranchProtectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.branchProtection.BranchProtection",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BranchProtection)SetAllowForcePush(val interface{}) {
	if err := j.validateSetAllowForcePushParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowForcePush",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetBranch(val *string) {
	if err := j.validateSetBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branch",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetCodeOwnerApprovalRequired(val interface{}) {
	if err := j.validateSetCodeOwnerApprovalRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codeOwnerApprovalRequired",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetMergeAccessLevel(val *string) {
	if err := j.validateSetMergeAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeAccessLevel",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetPushAccessLevel(val *string) {
	if err := j.validateSetPushAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pushAccessLevel",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetUnprotectAccessLevel(val *string) {
	if err := j.validateSetUnprotectAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unprotectAccessLevel",
		val,
	)
}

// Generates CDKTF code for importing a BranchProtection resource upon running "cdktf plan <stack-name>".
func BranchProtection_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateBranchProtection_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.branchProtection.BranchProtection",
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
func BranchProtection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBranchProtection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.branchProtection.BranchProtection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BranchProtection_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBranchProtection_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.branchProtection.BranchProtection",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BranchProtection_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBranchProtection_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.branchProtection.BranchProtection",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BranchProtection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.branchProtection.BranchProtection",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BranchProtection) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BranchProtection) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BranchProtection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtection) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtection) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtection) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtection) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtection) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtection) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtection) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtection) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtection) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BranchProtection) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtection) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BranchProtection) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BranchProtection) PutAllowedToMerge(value interface{}) {
	if err := b.validatePutAllowedToMergeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAllowedToMerge",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BranchProtection) PutAllowedToPush(value interface{}) {
	if err := b.validatePutAllowedToPushParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAllowedToPush",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BranchProtection) PutAllowedToUnprotect(value interface{}) {
	if err := b.validatePutAllowedToUnprotectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAllowedToUnprotect",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BranchProtection) ResetAllowedToMerge() {
	_jsii_.InvokeVoid(
		b,
		"resetAllowedToMerge",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtection) ResetAllowedToPush() {
	_jsii_.InvokeVoid(
		b,
		"resetAllowedToPush",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtection) ResetAllowedToUnprotect() {
	_jsii_.InvokeVoid(
		b,
		"resetAllowedToUnprotect",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtection) ResetAllowForcePush() {
	_jsii_.InvokeVoid(
		b,
		"resetAllowForcePush",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtection) ResetCodeOwnerApprovalRequired() {
	_jsii_.InvokeVoid(
		b,
		"resetCodeOwnerApprovalRequired",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtection) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtection) ResetMergeAccessLevel() {
	_jsii_.InvokeVoid(
		b,
		"resetMergeAccessLevel",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtection) ResetPushAccessLevel() {
	_jsii_.InvokeVoid(
		b,
		"resetPushAccessLevel",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtection) ResetUnprotectAccessLevel() {
	_jsii_.InvokeVoid(
		b,
		"resetUnprotectAccessLevel",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtection) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtection) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

