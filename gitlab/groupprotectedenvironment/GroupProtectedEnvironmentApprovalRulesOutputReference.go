// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupprotectedenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v11/jsii"

	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v11/groupprotectedenvironment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupProtectedEnvironmentApprovalRulesOutputReference interface {
	cdktf.ComplexObject
	AccessLevel() *string
	SetAccessLevel(val *string)
	AccessLevelDescription() *string
	AccessLevelInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GroupId() *float64
	SetGroupId(val *float64)
	GroupIdInput() *float64
	Id() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RequiredApprovals() *float64
	SetRequiredApprovals(val *float64)
	RequiredApprovalsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserId() *float64
	SetUserId(val *float64)
	UserIdInput() *float64
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAccessLevel()
	ResetGroupId()
	ResetRequiredApprovals()
	ResetUserId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GroupProtectedEnvironmentApprovalRulesOutputReference
type jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) AccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) AccessLevelDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessLevelDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) AccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) GroupId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) GroupIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) RequiredApprovals() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredApprovals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) RequiredApprovalsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredApprovalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) UserId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) UserIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userIdInput",
		&returns,
	)
	return returns
}


func NewGroupProtectedEnvironmentApprovalRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GroupProtectedEnvironmentApprovalRulesOutputReference {
	_init_.Initialize()

	if err := validateNewGroupProtectedEnvironmentApprovalRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.groupProtectedEnvironment.GroupProtectedEnvironmentApprovalRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGroupProtectedEnvironmentApprovalRulesOutputReference_Override(g GroupProtectedEnvironmentApprovalRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.groupProtectedEnvironment.GroupProtectedEnvironmentApprovalRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference)SetAccessLevel(val *string) {
	if err := j.validateSetAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessLevel",
		val,
	)
}

func (j *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference)SetGroupId(val *float64) {
	if err := j.validateSetGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupId",
		val,
	)
}

func (j *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference)SetRequiredApprovals(val *float64) {
	if err := j.validateSetRequiredApprovalsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredApprovals",
		val,
	)
}

func (j *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference)SetUserId(val *float64) {
	if err := j.validateSetUserIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userId",
		val,
	)
}

func (g *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) ResetAccessLevel() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessLevel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) ResetGroupId() {
	_jsii_.InvokeVoid(
		g,
		"resetGroupId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) ResetRequiredApprovals() {
	_jsii_.InvokeVoid(
		g,
		"resetRequiredApprovals",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) ResetUserId() {
	_jsii_.InvokeVoid(
		g,
		"resetUserId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupProtectedEnvironmentApprovalRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

