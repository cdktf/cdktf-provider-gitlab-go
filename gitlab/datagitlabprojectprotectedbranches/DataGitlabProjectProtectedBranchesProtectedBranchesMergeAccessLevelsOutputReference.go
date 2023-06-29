package datagitlabprojectprotectedbranches

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v9/jsii"

	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v9/datagitlabprojectprotectedbranches/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference interface {
	cdktf.ComplexObject
	AccessLevel() *string
	AccessLevelDescription() *string
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
	InternalValue() *DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevels
	SetInternalValue(val *DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevels)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserId() *float64
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

// The jsii proxy struct for DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference
type jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference) AccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference) AccessLevelDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessLevelDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference) GroupId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference) InternalValue() *DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevels {
	var returns *DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevels
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference) UserId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}


func NewDataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference {
	_init_.Initialize()

	if err := validateNewDataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.dataGitlabProjectProtectedBranches.DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference_Override(d DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.dataGitlabProjectProtectedBranches.DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference)SetInternalValue(val *DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevels) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesMergeAccessLevelsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

