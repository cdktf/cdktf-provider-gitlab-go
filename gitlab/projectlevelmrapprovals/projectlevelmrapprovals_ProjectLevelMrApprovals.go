package projectlevelmrapprovals

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-gitlab-go/gitlab/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-gitlab-go/gitlab/v3/projectlevelmrapprovals/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/gitlab/r/project_level_mr_approvals gitlab_project_level_mr_approvals}.
type ProjectLevelMrApprovals interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisableOverridingApproversPerMergeRequest() interface{}
	SetDisableOverridingApproversPerMergeRequest(val interface{})
	DisableOverridingApproversPerMergeRequestInput() interface{}
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MergeRequestsAuthorApproval() interface{}
	SetMergeRequestsAuthorApproval(val interface{})
	MergeRequestsAuthorApprovalInput() interface{}
	MergeRequestsDisableCommittersApproval() interface{}
	SetMergeRequestsDisableCommittersApproval(val interface{})
	MergeRequestsDisableCommittersApprovalInput() interface{}
	// The tree node.
	Node() constructs.Node
	ProjectId() *float64
	SetProjectId(val *float64)
	ProjectIdInput() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RequirePasswordToApprove() interface{}
	SetRequirePasswordToApprove(val interface{})
	RequirePasswordToApproveInput() interface{}
	ResetApprovalsOnPush() interface{}
	SetResetApprovalsOnPush(val interface{})
	ResetApprovalsOnPushInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetDisableOverridingApproversPerMergeRequest()
	ResetId()
	ResetMergeRequestsAuthorApproval()
	ResetMergeRequestsDisableCommittersApproval()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRequirePasswordToApprove()
	ResetResetApprovalsOnPush()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ProjectLevelMrApprovals
type jsiiProxy_ProjectLevelMrApprovals struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ProjectLevelMrApprovals) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) DisableOverridingApproversPerMergeRequest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableOverridingApproversPerMergeRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) DisableOverridingApproversPerMergeRequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableOverridingApproversPerMergeRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) MergeRequestsAuthorApproval() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeRequestsAuthorApproval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) MergeRequestsAuthorApprovalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeRequestsAuthorApprovalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) MergeRequestsDisableCommittersApproval() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeRequestsDisableCommittersApproval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) MergeRequestsDisableCommittersApprovalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeRequestsDisableCommittersApprovalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) ProjectId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) ProjectIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) RequirePasswordToApprove() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requirePasswordToApprove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) RequirePasswordToApproveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requirePasswordToApproveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) ResetApprovalsOnPush() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resetApprovalsOnPush",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) ResetApprovalsOnPushInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resetApprovalsOnPushInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectLevelMrApprovals) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/gitlab/r/project_level_mr_approvals gitlab_project_level_mr_approvals} Resource.
func NewProjectLevelMrApprovals(scope constructs.Construct, id *string, config *ProjectLevelMrApprovalsConfig) ProjectLevelMrApprovals {
	_init_.Initialize()

	if err := validateNewProjectLevelMrApprovalsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProjectLevelMrApprovals{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.projectLevelMrApprovals.ProjectLevelMrApprovals",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/gitlab/r/project_level_mr_approvals gitlab_project_level_mr_approvals} Resource.
func NewProjectLevelMrApprovals_Override(p ProjectLevelMrApprovals, scope constructs.Construct, id *string, config *ProjectLevelMrApprovalsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.projectLevelMrApprovals.ProjectLevelMrApprovals",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_ProjectLevelMrApprovals)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelMrApprovals)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelMrApprovals)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelMrApprovals)SetDisableOverridingApproversPerMergeRequest(val interface{}) {
	if err := j.validateSetDisableOverridingApproversPerMergeRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableOverridingApproversPerMergeRequest",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelMrApprovals)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelMrApprovals)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelMrApprovals)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelMrApprovals)SetMergeRequestsAuthorApproval(val interface{}) {
	if err := j.validateSetMergeRequestsAuthorApprovalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeRequestsAuthorApproval",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelMrApprovals)SetMergeRequestsDisableCommittersApproval(val interface{}) {
	if err := j.validateSetMergeRequestsDisableCommittersApprovalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeRequestsDisableCommittersApproval",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelMrApprovals)SetProjectId(val *float64) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelMrApprovals)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelMrApprovals)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelMrApprovals)SetRequirePasswordToApprove(val interface{}) {
	if err := j.validateSetRequirePasswordToApproveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requirePasswordToApprove",
		val,
	)
}

func (j *jsiiProxy_ProjectLevelMrApprovals)SetResetApprovalsOnPush(val interface{}) {
	if err := j.validateSetResetApprovalsOnPushParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resetApprovalsOnPush",
		val,
	)
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
func ProjectLevelMrApprovals_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjectLevelMrApprovals_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.projectLevelMrApprovals.ProjectLevelMrApprovals",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ProjectLevelMrApprovals_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.projectLevelMrApprovals.ProjectLevelMrApprovals",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_ProjectLevelMrApprovals) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_ProjectLevelMrApprovals) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_ProjectLevelMrApprovals) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_ProjectLevelMrApprovals) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_ProjectLevelMrApprovals) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_ProjectLevelMrApprovals) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_ProjectLevelMrApprovals) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_ProjectLevelMrApprovals) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_ProjectLevelMrApprovals) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_ProjectLevelMrApprovals) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_ProjectLevelMrApprovals) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectLevelMrApprovals) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_ProjectLevelMrApprovals) ResetDisableOverridingApproversPerMergeRequest() {
	_jsii_.InvokeVoid(
		p,
		"resetDisableOverridingApproversPerMergeRequest",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectLevelMrApprovals) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectLevelMrApprovals) ResetMergeRequestsAuthorApproval() {
	_jsii_.InvokeVoid(
		p,
		"resetMergeRequestsAuthorApproval",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectLevelMrApprovals) ResetMergeRequestsDisableCommittersApproval() {
	_jsii_.InvokeVoid(
		p,
		"resetMergeRequestsDisableCommittersApproval",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectLevelMrApprovals) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectLevelMrApprovals) ResetRequirePasswordToApprove() {
	_jsii_.InvokeVoid(
		p,
		"resetRequirePasswordToApprove",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectLevelMrApprovals) ResetResetApprovalsOnPush() {
	_jsii_.InvokeVoid(
		p,
		"resetResetApprovalsOnPush",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectLevelMrApprovals) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectLevelMrApprovals) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectLevelMrApprovals) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectLevelMrApprovals) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

