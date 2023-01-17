package datagitlabusers

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v5/jsii"

	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v5/datagitlabusers/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabUsersUsersOutputReference interface {
	cdktf.ComplexObject
	AvatarUrl() *string
	Bio() *string
	CanCreateGroup() cdktf.IResolvable
	CanCreateProject() cdktf.IResolvable
	ColorSchemeId() *float64
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
	CreatedAt() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CurrentSignInAt() *string
	Email() *string
	External() cdktf.IResolvable
	ExternUid() *string
	// Experimental.
	Fqn() *string
	Id() *float64
	InternalValue() *DataGitlabUsersUsers
	SetInternalValue(val *DataGitlabUsersUsers)
	IsAdmin() cdktf.IResolvable
	LastSignInAt() *string
	Linkedin() *string
	Location() *string
	Name() *string
	NamespaceId() *float64
	Organization() *string
	ProjectsLimit() *float64
	Provider() *string
	Skype() *string
	State() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThemeId() *float64
	Twitter() *string
	TwoFactorEnabled() cdktf.IResolvable
	Username() *string
	WebsiteUrl() *string
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

// The jsii proxy struct for DataGitlabUsersUsersOutputReference
type jsiiProxy_DataGitlabUsersUsersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) AvatarUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"avatarUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) Bio() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) CanCreateGroup() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"canCreateGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) CanCreateProject() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"canCreateProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) ColorSchemeId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"colorSchemeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) CurrentSignInAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentSignInAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) External() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"external",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) ExternUid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) InternalValue() *DataGitlabUsersUsers {
	var returns *DataGitlabUsersUsers
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) IsAdmin() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isAdmin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) LastSignInAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastSignInAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) Linkedin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkedin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) NamespaceId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"namespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) Organization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) ProjectsLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"projectsLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) Provider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) Skype() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skype",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) ThemeId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"themeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) Twitter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twitter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) TwoFactorEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"twoFactorEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference) WebsiteUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteUrl",
		&returns,
	)
	return returns
}


func NewDataGitlabUsersUsersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGitlabUsersUsersOutputReference {
	_init_.Initialize()

	if err := validateNewDataGitlabUsersUsersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGitlabUsersUsersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.dataGitlabUsers.DataGitlabUsersUsersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGitlabUsersUsersOutputReference_Override(d DataGitlabUsersUsersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.dataGitlabUsers.DataGitlabUsersUsersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference)SetInternalValue(val *DataGitlabUsersUsers) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGitlabUsersUsersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGitlabUsersUsersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabUsersUsersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGitlabUsersUsersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGitlabUsersUsersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGitlabUsersUsersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGitlabUsersUsersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGitlabUsersUsersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGitlabUsersUsersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGitlabUsersUsersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGitlabUsersUsersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGitlabUsersUsersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabUsersUsersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGitlabUsersUsersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataGitlabUsersUsersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

