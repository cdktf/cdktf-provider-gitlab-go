// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package userimpersonationtoken

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v13/userimpersonationtoken/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/resources/user_impersonation_token gitlab_user_impersonation_token}.
type UserImpersonationToken interface {
	cdktf.TerraformResource
	Active() cdktf.IResolvable
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
	CreatedAt() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExpiresAt() *string
	SetExpiresAt(val *string)
	ExpiresAtInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	Impersonation() cdktf.IResolvable
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	Revoked() cdktf.IResolvable
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Token() *string
	TokenId() *float64
	UserId() *float64
	SetUserId(val *float64)
	UserIdInput() *float64
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
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for UserImpersonationToken
type jsiiProxy_UserImpersonationToken struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_UserImpersonationToken) Active() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"active",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) ExpiresAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiresAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) ExpiresAtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiresAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) Impersonation() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"impersonation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) Revoked() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"revoked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) TokenId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) UserId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserImpersonationToken) UserIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/resources/user_impersonation_token gitlab_user_impersonation_token} Resource.
func NewUserImpersonationToken(scope constructs.Construct, id *string, config *UserImpersonationTokenConfig) UserImpersonationToken {
	_init_.Initialize()

	if err := validateNewUserImpersonationTokenParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_UserImpersonationToken{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.userImpersonationToken.UserImpersonationToken",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.4.0/docs/resources/user_impersonation_token gitlab_user_impersonation_token} Resource.
func NewUserImpersonationToken_Override(u UserImpersonationToken, scope constructs.Construct, id *string, config *UserImpersonationTokenConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.userImpersonationToken.UserImpersonationToken",
		[]interface{}{scope, id, config},
		u,
	)
}

func (j *jsiiProxy_UserImpersonationToken)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_UserImpersonationToken)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_UserImpersonationToken)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_UserImpersonationToken)SetExpiresAt(val *string) {
	if err := j.validateSetExpiresAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expiresAt",
		val,
	)
}

func (j *jsiiProxy_UserImpersonationToken)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_UserImpersonationToken)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_UserImpersonationToken)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_UserImpersonationToken)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_UserImpersonationToken)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_UserImpersonationToken)SetScopes(val *[]*string) {
	if err := j.validateSetScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_UserImpersonationToken)SetUserId(val *float64) {
	if err := j.validateSetUserIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userId",
		val,
	)
}

// Generates CDKTF code for importing a UserImpersonationToken resource upon running "cdktf plan <stack-name>".
func UserImpersonationToken_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateUserImpersonationToken_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.userImpersonationToken.UserImpersonationToken",
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
func UserImpersonationToken_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUserImpersonationToken_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.userImpersonationToken.UserImpersonationToken",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func UserImpersonationToken_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUserImpersonationToken_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.userImpersonationToken.UserImpersonationToken",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func UserImpersonationToken_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUserImpersonationToken_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.userImpersonationToken.UserImpersonationToken",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func UserImpersonationToken_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.userImpersonationToken.UserImpersonationToken",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (u *jsiiProxy_UserImpersonationToken) AddMoveTarget(moveTarget *string) {
	if err := u.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (u *jsiiProxy_UserImpersonationToken) AddOverride(path *string, value interface{}) {
	if err := u.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (u *jsiiProxy_UserImpersonationToken) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := u.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserImpersonationToken) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := u.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserImpersonationToken) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := u.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		u,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserImpersonationToken) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := u.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserImpersonationToken) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := u.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		u,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserImpersonationToken) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := u.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		u,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserImpersonationToken) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := u.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		u,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserImpersonationToken) GetStringAttribute(terraformAttribute *string) *string {
	if err := u.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		u,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserImpersonationToken) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := u.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		u,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserImpersonationToken) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserImpersonationToken) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := u.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (u *jsiiProxy_UserImpersonationToken) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := u.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserImpersonationToken) MoveFromId(id *string) {
	if err := u.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"moveFromId",
		[]interface{}{id},
	)
}

func (u *jsiiProxy_UserImpersonationToken) MoveTo(moveTarget *string, index interface{}) {
	if err := u.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (u *jsiiProxy_UserImpersonationToken) MoveToId(id *string) {
	if err := u.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"moveToId",
		[]interface{}{id},
	)
}

func (u *jsiiProxy_UserImpersonationToken) OverrideLogicalId(newLogicalId *string) {
	if err := u.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (u *jsiiProxy_UserImpersonationToken) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		u,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserImpersonationToken) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserImpersonationToken) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserImpersonationToken) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserImpersonationToken) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserImpersonationToken) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserImpersonationToken) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

