package datagitlabprojecthooks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v6/jsii"

	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v6/datagitlabprojecthooks/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabProjectHooksHooksOutputReference interface {
	cdktf.ComplexObject
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
	ConfidentialIssuesEvents() cdktf.IResolvable
	ConfidentialNoteEvents() cdktf.IResolvable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeploymentEvents() cdktf.IResolvable
	EnableSslVerification() cdktf.IResolvable
	// Experimental.
	Fqn() *string
	HookId() *float64
	InternalValue() *DataGitlabProjectHooksHooks
	SetInternalValue(val *DataGitlabProjectHooksHooks)
	IssuesEvents() cdktf.IResolvable
	JobEvents() cdktf.IResolvable
	MergeRequestsEvents() cdktf.IResolvable
	NoteEvents() cdktf.IResolvable
	PipelineEvents() cdktf.IResolvable
	Project() *string
	ProjectId() *float64
	PushEvents() cdktf.IResolvable
	PushEventsBranchFilter() *string
	ReleasesEvents() cdktf.IResolvable
	TagPushEvents() cdktf.IResolvable
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Token() *string
	Url() *string
	WikiPageEvents() cdktf.IResolvable
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGitlabProjectHooksHooksOutputReference
type jsiiProxy_DataGitlabProjectHooksHooksOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) ConfidentialIssuesEvents() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"confidentialIssuesEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) ConfidentialNoteEvents() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"confidentialNoteEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) DeploymentEvents() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"deploymentEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) EnableSslVerification() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableSslVerification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) HookId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hookId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) InternalValue() *DataGitlabProjectHooksHooks {
	var returns *DataGitlabProjectHooksHooks
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) IssuesEvents() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"issuesEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) JobEvents() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"jobEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) MergeRequestsEvents() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"mergeRequestsEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) NoteEvents() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"noteEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) PipelineEvents() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"pipelineEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) ProjectId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) PushEvents() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"pushEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) PushEventsBranchFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pushEventsBranchFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) ReleasesEvents() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"releasesEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) TagPushEvents() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"tagPushEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) WikiPageEvents() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"wikiPageEvents",
		&returns,
	)
	return returns
}


func NewDataGitlabProjectHooksHooksOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGitlabProjectHooksHooksOutputReference {
	_init_.Initialize()

	if err := validateNewDataGitlabProjectHooksHooksOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGitlabProjectHooksHooksOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.dataGitlabProjectHooks.DataGitlabProjectHooksHooksOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGitlabProjectHooksHooksOutputReference_Override(d DataGitlabProjectHooksHooksOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.dataGitlabProjectHooks.DataGitlabProjectHooksHooksOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference)SetInternalValue(val *DataGitlabProjectHooksHooks) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataGitlabProjectHooksHooksOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

