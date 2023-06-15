package project

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v8/jsii"

	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v8/project/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectPushRulesOutputReference interface {
	cdktf.ComplexObject
	AuthorEmailRegex() *string
	SetAuthorEmailRegex(val *string)
	AuthorEmailRegexInput() *string
	BranchNameRegex() *string
	SetBranchNameRegex(val *string)
	BranchNameRegexInput() *string
	CommitCommitterCheck() interface{}
	SetCommitCommitterCheck(val interface{})
	CommitCommitterCheckInput() interface{}
	CommitMessageNegativeRegex() *string
	SetCommitMessageNegativeRegex(val *string)
	CommitMessageNegativeRegexInput() *string
	CommitMessageRegex() *string
	SetCommitMessageRegex(val *string)
	CommitMessageRegexInput() *string
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
	DenyDeleteTag() interface{}
	SetDenyDeleteTag(val interface{})
	DenyDeleteTagInput() interface{}
	FileNameRegex() *string
	SetFileNameRegex(val *string)
	FileNameRegexInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *ProjectPushRules
	SetInternalValue(val *ProjectPushRules)
	MaxFileSize() *float64
	SetMaxFileSize(val *float64)
	MaxFileSizeInput() *float64
	MemberCheck() interface{}
	SetMemberCheck(val interface{})
	MemberCheckInput() interface{}
	PreventSecrets() interface{}
	SetPreventSecrets(val interface{})
	PreventSecretsInput() interface{}
	RejectUnsignedCommits() interface{}
	SetRejectUnsignedCommits(val interface{})
	RejectUnsignedCommitsInput() interface{}
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
	ResetAuthorEmailRegex()
	ResetBranchNameRegex()
	ResetCommitCommitterCheck()
	ResetCommitMessageNegativeRegex()
	ResetCommitMessageRegex()
	ResetDenyDeleteTag()
	ResetFileNameRegex()
	ResetMaxFileSize()
	ResetMemberCheck()
	ResetPreventSecrets()
	ResetRejectUnsignedCommits()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ProjectPushRulesOutputReference
type jsiiProxy_ProjectPushRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) AuthorEmailRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorEmailRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) AuthorEmailRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorEmailRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) BranchNameRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchNameRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) BranchNameRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchNameRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) CommitCommitterCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commitCommitterCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) CommitCommitterCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commitCommitterCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) CommitMessageNegativeRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitMessageNegativeRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) CommitMessageNegativeRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitMessageNegativeRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) CommitMessageRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitMessageRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) CommitMessageRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitMessageRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) DenyDeleteTag() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"denyDeleteTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) DenyDeleteTagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"denyDeleteTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) FileNameRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileNameRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) FileNameRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileNameRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) InternalValue() *ProjectPushRules {
	var returns *ProjectPushRules
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) MaxFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) MaxFileSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) MemberCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memberCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) MemberCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memberCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) PreventSecrets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) PreventSecretsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventSecretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) RejectUnsignedCommits() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rejectUnsignedCommits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) RejectUnsignedCommitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rejectUnsignedCommitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewProjectPushRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ProjectPushRulesOutputReference {
	_init_.Initialize()

	if err := validateNewProjectPushRulesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProjectPushRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.project.ProjectPushRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewProjectPushRulesOutputReference_Override(p ProjectPushRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.project.ProjectPushRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_ProjectPushRulesOutputReference)SetAuthorEmailRegex(val *string) {
	if err := j.validateSetAuthorEmailRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorEmailRegex",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesOutputReference)SetBranchNameRegex(val *string) {
	if err := j.validateSetBranchNameRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branchNameRegex",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesOutputReference)SetCommitCommitterCheck(val interface{}) {
	if err := j.validateSetCommitCommitterCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitCommitterCheck",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesOutputReference)SetCommitMessageNegativeRegex(val *string) {
	if err := j.validateSetCommitMessageNegativeRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitMessageNegativeRegex",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesOutputReference)SetCommitMessageRegex(val *string) {
	if err := j.validateSetCommitMessageRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitMessageRegex",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesOutputReference)SetDenyDeleteTag(val interface{}) {
	if err := j.validateSetDenyDeleteTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"denyDeleteTag",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesOutputReference)SetFileNameRegex(val *string) {
	if err := j.validateSetFileNameRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileNameRegex",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesOutputReference)SetInternalValue(val *ProjectPushRules) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesOutputReference)SetMaxFileSize(val *float64) {
	if err := j.validateSetMaxFileSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFileSize",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesOutputReference)SetMemberCheck(val interface{}) {
	if err := j.validateSetMemberCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memberCheck",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesOutputReference)SetPreventSecrets(val interface{}) {
	if err := j.validateSetPreventSecretsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preventSecrets",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesOutputReference)SetRejectUnsignedCommits(val interface{}) {
	if err := j.validateSetRejectUnsignedCommitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rejectUnsignedCommits",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_ProjectPushRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectPushRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_ProjectPushRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_ProjectPushRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_ProjectPushRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_ProjectPushRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_ProjectPushRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_ProjectPushRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_ProjectPushRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_ProjectPushRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_ProjectPushRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectPushRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_ProjectPushRulesOutputReference) ResetAuthorEmailRegex() {
	_jsii_.InvokeVoid(
		p,
		"resetAuthorEmailRegex",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectPushRulesOutputReference) ResetBranchNameRegex() {
	_jsii_.InvokeVoid(
		p,
		"resetBranchNameRegex",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectPushRulesOutputReference) ResetCommitCommitterCheck() {
	_jsii_.InvokeVoid(
		p,
		"resetCommitCommitterCheck",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectPushRulesOutputReference) ResetCommitMessageNegativeRegex() {
	_jsii_.InvokeVoid(
		p,
		"resetCommitMessageNegativeRegex",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectPushRulesOutputReference) ResetCommitMessageRegex() {
	_jsii_.InvokeVoid(
		p,
		"resetCommitMessageRegex",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectPushRulesOutputReference) ResetDenyDeleteTag() {
	_jsii_.InvokeVoid(
		p,
		"resetDenyDeleteTag",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectPushRulesOutputReference) ResetFileNameRegex() {
	_jsii_.InvokeVoid(
		p,
		"resetFileNameRegex",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectPushRulesOutputReference) ResetMaxFileSize() {
	_jsii_.InvokeVoid(
		p,
		"resetMaxFileSize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectPushRulesOutputReference) ResetMemberCheck() {
	_jsii_.InvokeVoid(
		p,
		"resetMemberCheck",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectPushRulesOutputReference) ResetPreventSecrets() {
	_jsii_.InvokeVoid(
		p,
		"resetPreventSecrets",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectPushRulesOutputReference) ResetRejectUnsignedCommits() {
	_jsii_.InvokeVoid(
		p,
		"resetRejectUnsignedCommits",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectPushRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_ProjectPushRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

