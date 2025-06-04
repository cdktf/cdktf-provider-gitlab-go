// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package group

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v15/jsii"

	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v15/group/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupPushRulesOutputReference interface {
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
	CommitCommitterNameCheck() interface{}
	SetCommitCommitterNameCheck(val interface{})
	CommitCommitterNameCheckInput() interface{}
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
	InternalValue() *GroupPushRules
	SetInternalValue(val *GroupPushRules)
	MaxFileSize() *float64
	SetMaxFileSize(val *float64)
	MaxFileSizeInput() *float64
	MemberCheck() interface{}
	SetMemberCheck(val interface{})
	MemberCheckInput() interface{}
	PreventSecrets() interface{}
	SetPreventSecrets(val interface{})
	PreventSecretsInput() interface{}
	RejectNonDcoCommits() interface{}
	SetRejectNonDcoCommits(val interface{})
	RejectNonDcoCommitsInput() interface{}
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
	ResetCommitCommitterNameCheck()
	ResetCommitMessageNegativeRegex()
	ResetCommitMessageRegex()
	ResetDenyDeleteTag()
	ResetFileNameRegex()
	ResetMaxFileSize()
	ResetMemberCheck()
	ResetPreventSecrets()
	ResetRejectNonDcoCommits()
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

// The jsii proxy struct for GroupPushRulesOutputReference
type jsiiProxy_GroupPushRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GroupPushRulesOutputReference) AuthorEmailRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorEmailRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) AuthorEmailRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorEmailRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) BranchNameRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchNameRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) BranchNameRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchNameRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) CommitCommitterCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commitCommitterCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) CommitCommitterCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commitCommitterCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) CommitCommitterNameCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commitCommitterNameCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) CommitCommitterNameCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commitCommitterNameCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) CommitMessageNegativeRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitMessageNegativeRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) CommitMessageNegativeRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitMessageNegativeRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) CommitMessageRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitMessageRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) CommitMessageRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitMessageRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) DenyDeleteTag() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"denyDeleteTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) DenyDeleteTagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"denyDeleteTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) FileNameRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileNameRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) FileNameRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileNameRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) InternalValue() *GroupPushRules {
	var returns *GroupPushRules
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) MaxFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) MaxFileSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) MemberCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memberCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) MemberCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memberCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) PreventSecrets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) PreventSecretsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventSecretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) RejectNonDcoCommits() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rejectNonDcoCommits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) RejectNonDcoCommitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rejectNonDcoCommitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) RejectUnsignedCommits() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rejectUnsignedCommits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) RejectUnsignedCommitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rejectUnsignedCommitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupPushRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGroupPushRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GroupPushRulesOutputReference {
	_init_.Initialize()

	if err := validateNewGroupPushRulesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GroupPushRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.group.GroupPushRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGroupPushRulesOutputReference_Override(g GroupPushRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.group.GroupPushRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GroupPushRulesOutputReference)SetAuthorEmailRegex(val *string) {
	if err := j.validateSetAuthorEmailRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorEmailRegex",
		val,
	)
}

func (j *jsiiProxy_GroupPushRulesOutputReference)SetBranchNameRegex(val *string) {
	if err := j.validateSetBranchNameRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branchNameRegex",
		val,
	)
}

func (j *jsiiProxy_GroupPushRulesOutputReference)SetCommitCommitterCheck(val interface{}) {
	if err := j.validateSetCommitCommitterCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitCommitterCheck",
		val,
	)
}

func (j *jsiiProxy_GroupPushRulesOutputReference)SetCommitCommitterNameCheck(val interface{}) {
	if err := j.validateSetCommitCommitterNameCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitCommitterNameCheck",
		val,
	)
}

func (j *jsiiProxy_GroupPushRulesOutputReference)SetCommitMessageNegativeRegex(val *string) {
	if err := j.validateSetCommitMessageNegativeRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitMessageNegativeRegex",
		val,
	)
}

func (j *jsiiProxy_GroupPushRulesOutputReference)SetCommitMessageRegex(val *string) {
	if err := j.validateSetCommitMessageRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitMessageRegex",
		val,
	)
}

func (j *jsiiProxy_GroupPushRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GroupPushRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GroupPushRulesOutputReference)SetDenyDeleteTag(val interface{}) {
	if err := j.validateSetDenyDeleteTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"denyDeleteTag",
		val,
	)
}

func (j *jsiiProxy_GroupPushRulesOutputReference)SetFileNameRegex(val *string) {
	if err := j.validateSetFileNameRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileNameRegex",
		val,
	)
}

func (j *jsiiProxy_GroupPushRulesOutputReference)SetInternalValue(val *GroupPushRules) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GroupPushRulesOutputReference)SetMaxFileSize(val *float64) {
	if err := j.validateSetMaxFileSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFileSize",
		val,
	)
}

func (j *jsiiProxy_GroupPushRulesOutputReference)SetMemberCheck(val interface{}) {
	if err := j.validateSetMemberCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memberCheck",
		val,
	)
}

func (j *jsiiProxy_GroupPushRulesOutputReference)SetPreventSecrets(val interface{}) {
	if err := j.validateSetPreventSecretsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preventSecrets",
		val,
	)
}

func (j *jsiiProxy_GroupPushRulesOutputReference)SetRejectNonDcoCommits(val interface{}) {
	if err := j.validateSetRejectNonDcoCommitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rejectNonDcoCommits",
		val,
	)
}

func (j *jsiiProxy_GroupPushRulesOutputReference)SetRejectUnsignedCommits(val interface{}) {
	if err := j.validateSetRejectUnsignedCommitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rejectUnsignedCommits",
		val,
	)
}

func (j *jsiiProxy_GroupPushRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GroupPushRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GroupPushRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupPushRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GroupPushRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GroupPushRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GroupPushRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GroupPushRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GroupPushRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GroupPushRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GroupPushRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GroupPushRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GroupPushRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupPushRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GroupPushRulesOutputReference) ResetAuthorEmailRegex() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthorEmailRegex",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupPushRulesOutputReference) ResetBranchNameRegex() {
	_jsii_.InvokeVoid(
		g,
		"resetBranchNameRegex",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupPushRulesOutputReference) ResetCommitCommitterCheck() {
	_jsii_.InvokeVoid(
		g,
		"resetCommitCommitterCheck",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupPushRulesOutputReference) ResetCommitCommitterNameCheck() {
	_jsii_.InvokeVoid(
		g,
		"resetCommitCommitterNameCheck",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupPushRulesOutputReference) ResetCommitMessageNegativeRegex() {
	_jsii_.InvokeVoid(
		g,
		"resetCommitMessageNegativeRegex",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupPushRulesOutputReference) ResetCommitMessageRegex() {
	_jsii_.InvokeVoid(
		g,
		"resetCommitMessageRegex",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupPushRulesOutputReference) ResetDenyDeleteTag() {
	_jsii_.InvokeVoid(
		g,
		"resetDenyDeleteTag",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupPushRulesOutputReference) ResetFileNameRegex() {
	_jsii_.InvokeVoid(
		g,
		"resetFileNameRegex",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupPushRulesOutputReference) ResetMaxFileSize() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxFileSize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupPushRulesOutputReference) ResetMemberCheck() {
	_jsii_.InvokeVoid(
		g,
		"resetMemberCheck",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupPushRulesOutputReference) ResetPreventSecrets() {
	_jsii_.InvokeVoid(
		g,
		"resetPreventSecrets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupPushRulesOutputReference) ResetRejectNonDcoCommits() {
	_jsii_.InvokeVoid(
		g,
		"resetRejectNonDcoCommits",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupPushRulesOutputReference) ResetRejectUnsignedCommits() {
	_jsii_.InvokeVoid(
		g,
		"resetRejectUnsignedCommits",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupPushRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GroupPushRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

