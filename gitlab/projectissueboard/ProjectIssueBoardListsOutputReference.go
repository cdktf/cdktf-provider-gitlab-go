// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectissueboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v10/jsii"

	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v10/projectissueboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectIssueBoardListsOutputReference interface {
	cdktf.ComplexObject
	AssigneeId() *float64
	SetAssigneeId(val *float64)
	AssigneeIdInput() *float64
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
	Id() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IterationId() *float64
	SetIterationId(val *float64)
	IterationIdInput() *float64
	LabelId() *float64
	SetLabelId(val *float64)
	LabelIdInput() *float64
	MilestoneId() *float64
	SetMilestoneId(val *float64)
	MilestoneIdInput() *float64
	Position() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetAssigneeId()
	ResetIterationId()
	ResetLabelId()
	ResetMilestoneId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ProjectIssueBoardListsOutputReference
type jsiiProxy_ProjectIssueBoardListsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ProjectIssueBoardListsOutputReference) AssigneeId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"assigneeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssueBoardListsOutputReference) AssigneeIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"assigneeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssueBoardListsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssueBoardListsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssueBoardListsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssueBoardListsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssueBoardListsOutputReference) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssueBoardListsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssueBoardListsOutputReference) IterationId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iterationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssueBoardListsOutputReference) IterationIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iterationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssueBoardListsOutputReference) LabelId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"labelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssueBoardListsOutputReference) LabelIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"labelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssueBoardListsOutputReference) MilestoneId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"milestoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssueBoardListsOutputReference) MilestoneIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"milestoneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssueBoardListsOutputReference) Position() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"position",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssueBoardListsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIssueBoardListsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewProjectIssueBoardListsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ProjectIssueBoardListsOutputReference {
	_init_.Initialize()

	if err := validateNewProjectIssueBoardListsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProjectIssueBoardListsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.projectIssueBoard.ProjectIssueBoardListsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewProjectIssueBoardListsOutputReference_Override(p ProjectIssueBoardListsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.projectIssueBoard.ProjectIssueBoardListsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_ProjectIssueBoardListsOutputReference)SetAssigneeId(val *float64) {
	if err := j.validateSetAssigneeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assigneeId",
		val,
	)
}

func (j *jsiiProxy_ProjectIssueBoardListsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ProjectIssueBoardListsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ProjectIssueBoardListsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ProjectIssueBoardListsOutputReference)SetIterationId(val *float64) {
	if err := j.validateSetIterationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iterationId",
		val,
	)
}

func (j *jsiiProxy_ProjectIssueBoardListsOutputReference)SetLabelId(val *float64) {
	if err := j.validateSetLabelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labelId",
		val,
	)
}

func (j *jsiiProxy_ProjectIssueBoardListsOutputReference)SetMilestoneId(val *float64) {
	if err := j.validateSetMilestoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"milestoneId",
		val,
	)
}

func (j *jsiiProxy_ProjectIssueBoardListsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ProjectIssueBoardListsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_ProjectIssueBoardListsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIssueBoardListsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_ProjectIssueBoardListsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_ProjectIssueBoardListsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_ProjectIssueBoardListsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_ProjectIssueBoardListsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_ProjectIssueBoardListsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_ProjectIssueBoardListsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_ProjectIssueBoardListsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_ProjectIssueBoardListsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_ProjectIssueBoardListsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIssueBoardListsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_ProjectIssueBoardListsOutputReference) ResetAssigneeId() {
	_jsii_.InvokeVoid(
		p,
		"resetAssigneeId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIssueBoardListsOutputReference) ResetIterationId() {
	_jsii_.InvokeVoid(
		p,
		"resetIterationId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIssueBoardListsOutputReference) ResetLabelId() {
	_jsii_.InvokeVoid(
		p,
		"resetLabelId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIssueBoardListsOutputReference) ResetMilestoneId() {
	_jsii_.InvokeVoid(
		p,
		"resetMilestoneId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIssueBoardListsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_ProjectIssueBoardListsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

