// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabprojectprotectedbranches

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v12/jsii"

	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v12/datagitlabprojectprotectedbranches/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference interface {
	cdktf.ComplexObject
	AllowForcePush() cdktf.IResolvable
	CodeOwnerApprovalRequired() cdktf.IResolvable
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
	MergeAccessLevels() DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsList
	MergeAccessLevelsInput() interface{}
	Name() *string
	PushAccessLevels() DataGitlabProjectProtectedBranchesProtectedBranchesPushAccessLevelsList
	PushAccessLevelsInput() interface{}
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
	PutMergeAccessLevels(value interface{})
	PutPushAccessLevels(value interface{})
	ResetMergeAccessLevels()
	ResetPushAccessLevels()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference
type jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) AllowForcePush() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"allowForcePush",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) CodeOwnerApprovalRequired() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"codeOwnerApprovalRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) MergeAccessLevels() DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsList {
	var returns DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsList
	_jsii_.Get(
		j,
		"mergeAccessLevels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) MergeAccessLevelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeAccessLevelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) PushAccessLevels() DataGitlabProjectProtectedBranchesProtectedBranchesPushAccessLevelsList {
	var returns DataGitlabProjectProtectedBranchesProtectedBranchesPushAccessLevelsList
	_jsii_.Get(
		j,
		"pushAccessLevels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) PushAccessLevelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pushAccessLevelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataGitlabProjectProtectedBranchesProtectedBranchesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference {
	_init_.Initialize()

	if err := validateNewDataGitlabProjectProtectedBranchesProtectedBranchesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.dataGitlabProjectProtectedBranches.DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGitlabProjectProtectedBranchesProtectedBranchesOutputReference_Override(d DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.dataGitlabProjectProtectedBranches.DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) PutMergeAccessLevels(value interface{}) {
	if err := d.validatePutMergeAccessLevelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMergeAccessLevels",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) PutPushAccessLevels(value interface{}) {
	if err := d.validatePutPushAccessLevelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPushAccessLevels",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) ResetMergeAccessLevels() {
	_jsii_.InvokeVoid(
		d,
		"resetMergeAccessLevels",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) ResetPushAccessLevels() {
	_jsii_.InvokeVoid(
		d,
		"resetPushAccessLevels",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

