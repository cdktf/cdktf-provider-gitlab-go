// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectprotectedenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v13/jsii"

	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v13/projectprotectedenvironment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectProtectedEnvironmentApprovalRulesOutputReference interface {
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
	GroupInheritanceType() *float64
	SetGroupInheritanceType(val *float64)
	GroupInheritanceTypeInput() *float64
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
	ResetGroupInheritanceType()
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

// The jsii proxy struct for ProjectProtectedEnvironmentApprovalRulesOutputReference
type jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) AccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) AccessLevelDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessLevelDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) AccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) GroupId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) GroupIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) GroupInheritanceType() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupInheritanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) GroupInheritanceTypeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupInheritanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) RequiredApprovals() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredApprovals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) RequiredApprovalsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredApprovalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) UserId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) UserIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userIdInput",
		&returns,
	)
	return returns
}


func NewProjectProtectedEnvironmentApprovalRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ProjectProtectedEnvironmentApprovalRulesOutputReference {
	_init_.Initialize()

	if err := validateNewProjectProtectedEnvironmentApprovalRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.projectProtectedEnvironment.ProjectProtectedEnvironmentApprovalRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewProjectProtectedEnvironmentApprovalRulesOutputReference_Override(p ProjectProtectedEnvironmentApprovalRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.projectProtectedEnvironment.ProjectProtectedEnvironmentApprovalRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference)SetAccessLevel(val *string) {
	if err := j.validateSetAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessLevel",
		val,
	)
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference)SetGroupId(val *float64) {
	if err := j.validateSetGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupId",
		val,
	)
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference)SetGroupInheritanceType(val *float64) {
	if err := j.validateSetGroupInheritanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupInheritanceType",
		val,
	)
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference)SetRequiredApprovals(val *float64) {
	if err := j.validateSetRequiredApprovalsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredApprovals",
		val,
	)
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference)SetUserId(val *float64) {
	if err := j.validateSetUserIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userId",
		val,
	)
}

func (p *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) ResetAccessLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetAccessLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) ResetGroupId() {
	_jsii_.InvokeVoid(
		p,
		"resetGroupId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) ResetGroupInheritanceType() {
	_jsii_.InvokeVoid(
		p,
		"resetGroupInheritanceType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) ResetRequiredApprovals() {
	_jsii_.InvokeVoid(
		p,
		"resetRequiredApprovals",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) ResetUserId() {
	_jsii_.InvokeVoid(
		p,
		"resetUserId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectProtectedEnvironmentApprovalRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

