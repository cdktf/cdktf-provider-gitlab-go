// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package instancecluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v14/instancecluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.11.0/docs/resources/instance_cluster gitlab_instance_cluster}.
type InstanceCluster interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterType() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAt() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EnvironmentScope() *string
	SetEnvironmentScope(val *string)
	EnvironmentScopeInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	KubernetesApiUrl() *string
	SetKubernetesApiUrl(val *string)
	KubernetesApiUrlInput() *string
	KubernetesAuthorizationType() *string
	SetKubernetesAuthorizationType(val *string)
	KubernetesAuthorizationTypeInput() *string
	KubernetesCaCert() *string
	SetKubernetesCaCert(val *string)
	KubernetesCaCertInput() *string
	KubernetesNamespace() *string
	SetKubernetesNamespace(val *string)
	KubernetesNamespaceInput() *string
	KubernetesToken() *string
	SetKubernetesToken(val *string)
	KubernetesTokenInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Managed() interface{}
	SetManaged(val interface{})
	ManagedInput() interface{}
	ManagementProjectId() *string
	SetManagementProjectId(val *string)
	ManagementProjectIdInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PlatformType() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	ProviderType() *string
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetDomain()
	ResetEnabled()
	ResetEnvironmentScope()
	ResetId()
	ResetKubernetesAuthorizationType()
	ResetKubernetesCaCert()
	ResetKubernetesNamespace()
	ResetManaged()
	ResetManagementProjectId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for InstanceCluster
type jsiiProxy_InstanceCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_InstanceCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) ClusterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) EnvironmentScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) EnvironmentScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) KubernetesApiUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesApiUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) KubernetesApiUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesApiUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) KubernetesAuthorizationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesAuthorizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) KubernetesAuthorizationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesAuthorizationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) KubernetesCaCert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesCaCert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) KubernetesCaCertInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesCaCertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) KubernetesNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) KubernetesNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) KubernetesToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) KubernetesTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) Managed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) ManagedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) ManagementProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managementProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) ManagementProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managementProjectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) PlatformType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) ProviderType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.11.0/docs/resources/instance_cluster gitlab_instance_cluster} Resource.
func NewInstanceCluster(scope constructs.Construct, id *string, config *InstanceClusterConfig) InstanceCluster {
	_init_.Initialize()

	if err := validateNewInstanceClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_InstanceCluster{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.instanceCluster.InstanceCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.11.0/docs/resources/instance_cluster gitlab_instance_cluster} Resource.
func NewInstanceCluster_Override(i InstanceCluster, scope constructs.Construct, id *string, config *InstanceClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.instanceCluster.InstanceCluster",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_InstanceCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_InstanceCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_InstanceCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_InstanceCluster)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_InstanceCluster)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_InstanceCluster)SetEnvironmentScope(val *string) {
	if err := j.validateSetEnvironmentScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentScope",
		val,
	)
}

func (j *jsiiProxy_InstanceCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_InstanceCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_InstanceCluster)SetKubernetesApiUrl(val *string) {
	if err := j.validateSetKubernetesApiUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kubernetesApiUrl",
		val,
	)
}

func (j *jsiiProxy_InstanceCluster)SetKubernetesAuthorizationType(val *string) {
	if err := j.validateSetKubernetesAuthorizationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kubernetesAuthorizationType",
		val,
	)
}

func (j *jsiiProxy_InstanceCluster)SetKubernetesCaCert(val *string) {
	if err := j.validateSetKubernetesCaCertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kubernetesCaCert",
		val,
	)
}

func (j *jsiiProxy_InstanceCluster)SetKubernetesNamespace(val *string) {
	if err := j.validateSetKubernetesNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kubernetesNamespace",
		val,
	)
}

func (j *jsiiProxy_InstanceCluster)SetKubernetesToken(val *string) {
	if err := j.validateSetKubernetesTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kubernetesToken",
		val,
	)
}

func (j *jsiiProxy_InstanceCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_InstanceCluster)SetManaged(val interface{}) {
	if err := j.validateSetManagedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managed",
		val,
	)
}

func (j *jsiiProxy_InstanceCluster)SetManagementProjectId(val *string) {
	if err := j.validateSetManagementProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managementProjectId",
		val,
	)
}

func (j *jsiiProxy_InstanceCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_InstanceCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_InstanceCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a InstanceCluster resource upon running "cdktf plan <stack-name>".
func InstanceCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateInstanceCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.instanceCluster.InstanceCluster",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func InstanceCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateInstanceCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.instanceCluster.InstanceCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func InstanceCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateInstanceCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.instanceCluster.InstanceCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func InstanceCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateInstanceCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.instanceCluster.InstanceCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func InstanceCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.instanceCluster.InstanceCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_InstanceCluster) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_InstanceCluster) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_InstanceCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_InstanceCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCluster) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_InstanceCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_InstanceCluster) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_InstanceCluster) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_InstanceCluster) ResetDomain() {
	_jsii_.InvokeVoid(
		i,
		"resetDomain",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstanceCluster) ResetEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstanceCluster) ResetEnvironmentScope() {
	_jsii_.InvokeVoid(
		i,
		"resetEnvironmentScope",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstanceCluster) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstanceCluster) ResetKubernetesAuthorizationType() {
	_jsii_.InvokeVoid(
		i,
		"resetKubernetesAuthorizationType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstanceCluster) ResetKubernetesCaCert() {
	_jsii_.InvokeVoid(
		i,
		"resetKubernetesCaCert",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstanceCluster) ResetKubernetesNamespace() {
	_jsii_.InvokeVoid(
		i,
		"resetKubernetesNamespace",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstanceCluster) ResetManaged() {
	_jsii_.InvokeVoid(
		i,
		"resetManaged",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstanceCluster) ResetManagementProjectId() {
	_jsii_.InvokeVoid(
		i,
		"resetManagementProjectId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstanceCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstanceCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

