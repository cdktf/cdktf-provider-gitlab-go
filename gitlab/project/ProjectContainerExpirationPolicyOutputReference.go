package project

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v9/jsii"

	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v9/project/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectContainerExpirationPolicyOutputReference interface {
	cdktf.ComplexObject
	Cadence() *string
	SetCadence(val *string)
	CadenceInput() *string
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ProjectContainerExpirationPolicy
	SetInternalValue(val *ProjectContainerExpirationPolicy)
	KeepN() *float64
	SetKeepN(val *float64)
	KeepNInput() *float64
	NameRegex() *string
	SetNameRegex(val *string)
	NameRegexDelete() *string
	SetNameRegexDelete(val *string)
	NameRegexDeleteInput() *string
	NameRegexInput() *string
	NameRegexKeep() *string
	SetNameRegexKeep(val *string)
	NameRegexKeepInput() *string
	NextRunAt() *string
	OlderThan() *string
	SetOlderThan(val *string)
	OlderThanInput() *string
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
	ResetCadence()
	ResetEnabled()
	ResetKeepN()
	ResetNameRegex()
	ResetNameRegexDelete()
	ResetNameRegexKeep()
	ResetOlderThan()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ProjectContainerExpirationPolicyOutputReference
type jsiiProxy_ProjectContainerExpirationPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) Cadence() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cadence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) CadenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cadenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) InternalValue() *ProjectContainerExpirationPolicy {
	var returns *ProjectContainerExpirationPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) KeepN() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepN",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) KeepNInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepNInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) NameRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) NameRegexDelete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameRegexDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) NameRegexDeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameRegexDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) NameRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) NameRegexKeep() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameRegexKeep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) NameRegexKeepInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameRegexKeepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) NextRunAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextRunAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) OlderThan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"olderThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) OlderThanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"olderThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewProjectContainerExpirationPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ProjectContainerExpirationPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewProjectContainerExpirationPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProjectContainerExpirationPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.project.ProjectContainerExpirationPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewProjectContainerExpirationPolicyOutputReference_Override(p ProjectContainerExpirationPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.project.ProjectContainerExpirationPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference)SetCadence(val *string) {
	if err := j.validateSetCadenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cadence",
		val,
	)
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference)SetInternalValue(val *ProjectContainerExpirationPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference)SetKeepN(val *float64) {
	if err := j.validateSetKeepNParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepN",
		val,
	)
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference)SetNameRegex(val *string) {
	if err := j.validateSetNameRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nameRegex",
		val,
	)
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference)SetNameRegexDelete(val *string) {
	if err := j.validateSetNameRegexDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nameRegexDelete",
		val,
	)
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference)SetNameRegexKeep(val *string) {
	if err := j.validateSetNameRegexKeepParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nameRegexKeep",
		val,
	)
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference)SetOlderThan(val *string) {
	if err := j.validateSetOlderThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"olderThan",
		val,
	)
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ProjectContainerExpirationPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) ResetCadence() {
	_jsii_.InvokeVoid(
		p,
		"resetCadence",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) ResetKeepN() {
	_jsii_.InvokeVoid(
		p,
		"resetKeepN",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) ResetNameRegex() {
	_jsii_.InvokeVoid(
		p,
		"resetNameRegex",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) ResetNameRegexDelete() {
	_jsii_.InvokeVoid(
		p,
		"resetNameRegexDelete",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) ResetNameRegexKeep() {
	_jsii_.InvokeVoid(
		p,
		"resetNameRegexKeep",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) ResetOlderThan() {
	_jsii_.InvokeVoid(
		p,
		"resetOlderThan",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_ProjectContainerExpirationPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

