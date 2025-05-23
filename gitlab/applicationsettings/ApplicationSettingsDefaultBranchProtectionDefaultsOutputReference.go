// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationsettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v14/jsii"

	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v14/applicationsettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference interface {
	cdktf.ComplexObject
	AllowedToMerge() *[]*float64
	SetAllowedToMerge(val *[]*float64)
	AllowedToMergeInput() *[]*float64
	AllowedToPush() *[]*float64
	SetAllowedToPush(val *[]*float64)
	AllowedToPushInput() *[]*float64
	AllowForcePush() interface{}
	SetAllowForcePush(val interface{})
	AllowForcePushInput() interface{}
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
	DeveloperCanInitialPush() interface{}
	SetDeveloperCanInitialPush(val interface{})
	DeveloperCanInitialPushInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ApplicationSettingsDefaultBranchProtectionDefaults
	SetInternalValue(val *ApplicationSettingsDefaultBranchProtectionDefaults)
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
	ResetAllowedToMerge()
	ResetAllowedToPush()
	ResetAllowForcePush()
	ResetDeveloperCanInitialPush()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference
type jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) AllowedToMerge() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"allowedToMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) AllowedToMergeInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"allowedToMergeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) AllowedToPush() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"allowedToPush",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) AllowedToPushInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"allowedToPushInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) AllowForcePush() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowForcePush",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) AllowForcePushInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowForcePushInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) DeveloperCanInitialPush() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"developerCanInitialPush",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) DeveloperCanInitialPushInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"developerCanInitialPushInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) InternalValue() *ApplicationSettingsDefaultBranchProtectionDefaults {
	var returns *ApplicationSettingsDefaultBranchProtectionDefaults
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApplicationSettingsDefaultBranchProtectionDefaultsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationSettingsDefaultBranchProtectionDefaultsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.applicationSettings.ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApplicationSettingsDefaultBranchProtectionDefaultsOutputReference_Override(a ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.applicationSettings.ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference)SetAllowedToMerge(val *[]*float64) {
	if err := j.validateSetAllowedToMergeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedToMerge",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference)SetAllowedToPush(val *[]*float64) {
	if err := j.validateSetAllowedToPushParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedToPush",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference)SetAllowForcePush(val interface{}) {
	if err := j.validateSetAllowForcePushParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowForcePush",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference)SetDeveloperCanInitialPush(val interface{}) {
	if err := j.validateSetDeveloperCanInitialPushParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"developerCanInitialPush",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference)SetInternalValue(val *ApplicationSettingsDefaultBranchProtectionDefaults) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) ResetAllowedToMerge() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedToMerge",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) ResetAllowedToPush() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedToPush",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) ResetAllowForcePush() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowForcePush",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) ResetDeveloperCanInitialPush() {
	_jsii_.InvokeVoid(
		a,
		"resetDeveloperCanInitialPush",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsDefaultBranchProtectionDefaultsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

