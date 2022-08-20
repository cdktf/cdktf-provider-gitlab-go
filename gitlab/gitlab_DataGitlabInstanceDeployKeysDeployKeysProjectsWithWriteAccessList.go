// Prebuilt gitlab Provider for Terraform CDK (cdktf)
package gitlab

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-gitlab-go/gitlab/v2/jsii"

	"github.com/hashicorp/cdktf-provider-gitlab-go/gitlab/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabInstanceDeployKeysDeployKeysProjectsWithWriteAccessList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataGitlabInstanceDeployKeysDeployKeysProjectsWithWriteAccessOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGitlabInstanceDeployKeysDeployKeysProjectsWithWriteAccessList
type jsiiProxy_DataGitlabInstanceDeployKeysDeployKeysProjectsWithWriteAccessList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataGitlabInstanceDeployKeysDeployKeysProjectsWithWriteAccessList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabInstanceDeployKeysDeployKeysProjectsWithWriteAccessList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabInstanceDeployKeysDeployKeysProjectsWithWriteAccessList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabInstanceDeployKeysDeployKeysProjectsWithWriteAccessList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabInstanceDeployKeysDeployKeysProjectsWithWriteAccessList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataGitlabInstanceDeployKeysDeployKeysProjectsWithWriteAccessList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataGitlabInstanceDeployKeysDeployKeysProjectsWithWriteAccessList {
	_init_.Initialize()

	j := jsiiProxy_DataGitlabInstanceDeployKeysDeployKeysProjectsWithWriteAccessList{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.DataGitlabInstanceDeployKeysDeployKeysProjectsWithWriteAccessList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataGitlabInstanceDeployKeysDeployKeysProjectsWithWriteAccessList_Override(d DataGitlabInstanceDeployKeysDeployKeysProjectsWithWriteAccessList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.DataGitlabInstanceDeployKeysDeployKeysProjectsWithWriteAccessList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataGitlabInstanceDeployKeysDeployKeysProjectsWithWriteAccessList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGitlabInstanceDeployKeysDeployKeysProjectsWithWriteAccessList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataGitlabInstanceDeployKeysDeployKeysProjectsWithWriteAccessList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataGitlabInstanceDeployKeysDeployKeysProjectsWithWriteAccessList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabInstanceDeployKeysDeployKeysProjectsWithWriteAccessList) Get(index *float64) DataGitlabInstanceDeployKeysDeployKeysProjectsWithWriteAccessOutputReference {
	var returns DataGitlabInstanceDeployKeysDeployKeysProjectsWithWriteAccessOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabInstanceDeployKeysDeployKeysProjectsWithWriteAccessList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabInstanceDeployKeysDeployKeysProjectsWithWriteAccessList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
