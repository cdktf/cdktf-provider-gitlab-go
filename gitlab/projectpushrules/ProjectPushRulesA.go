// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectpushrules

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v13/projectpushrules/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.7.0/docs/resources/project_push_rules gitlab_project_push_rules}.
type ProjectPushRulesA interface {
	cdktf.TerraformResource
	AuthorEmailRegex() *string
	SetAuthorEmailRegex(val *string)
	AuthorEmailRegexInput() *string
	BranchNameRegex() *string
	SetBranchNameRegex(val *string)
	BranchNameRegexInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	DenyDeleteTag() interface{}
	SetDenyDeleteTag(val interface{})
	DenyDeleteTagInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	FileNameRegex() *string
	SetFileNameRegex(val *string)
	FileNameRegexInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxFileSize() *float64
	SetMaxFileSize(val *float64)
	MaxFileSizeInput() *float64
	MemberCheck() interface{}
	SetMemberCheck(val interface{})
	MemberCheckInput() interface{}
	// The tree node.
	Node() constructs.Node
	PreventSecrets() interface{}
	SetPreventSecrets(val interface{})
	PreventSecretsInput() interface{}
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
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
	RejectNonDcoCommits() interface{}
	SetRejectNonDcoCommits(val interface{})
	RejectNonDcoCommitsInput() interface{}
	RejectUnsignedCommits() interface{}
	SetRejectUnsignedCommits(val interface{})
	RejectUnsignedCommitsInput() interface{}
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
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreventSecrets()
	ResetRejectNonDcoCommits()
	ResetRejectUnsignedCommits()
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

// The jsii proxy struct for ProjectPushRulesA
type jsiiProxy_ProjectPushRulesA struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ProjectPushRulesA) AuthorEmailRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorEmailRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) AuthorEmailRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorEmailRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) BranchNameRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchNameRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) BranchNameRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchNameRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) CommitCommitterCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commitCommitterCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) CommitCommitterCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commitCommitterCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) CommitCommitterNameCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commitCommitterNameCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) CommitCommitterNameCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commitCommitterNameCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) CommitMessageNegativeRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitMessageNegativeRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) CommitMessageNegativeRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitMessageNegativeRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) CommitMessageRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitMessageRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) CommitMessageRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitMessageRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) DenyDeleteTag() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"denyDeleteTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) DenyDeleteTagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"denyDeleteTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) FileNameRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileNameRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) FileNameRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileNameRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) MaxFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) MaxFileSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) MemberCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memberCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) MemberCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memberCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) PreventSecrets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) PreventSecretsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventSecretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) RejectNonDcoCommits() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rejectNonDcoCommits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) RejectNonDcoCommitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rejectNonDcoCommitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) RejectUnsignedCommits() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rejectUnsignedCommits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) RejectUnsignedCommitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rejectUnsignedCommitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectPushRulesA) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.7.0/docs/resources/project_push_rules gitlab_project_push_rules} Resource.
func NewProjectPushRulesA(scope constructs.Construct, id *string, config *ProjectPushRulesAConfig) ProjectPushRulesA {
	_init_.Initialize()

	if err := validateNewProjectPushRulesAParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProjectPushRulesA{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.projectPushRules.ProjectPushRulesA",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.7.0/docs/resources/project_push_rules gitlab_project_push_rules} Resource.
func NewProjectPushRulesA_Override(p ProjectPushRulesA, scope constructs.Construct, id *string, config *ProjectPushRulesAConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.projectPushRules.ProjectPushRulesA",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_ProjectPushRulesA)SetAuthorEmailRegex(val *string) {
	if err := j.validateSetAuthorEmailRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorEmailRegex",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesA)SetBranchNameRegex(val *string) {
	if err := j.validateSetBranchNameRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branchNameRegex",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesA)SetCommitCommitterCheck(val interface{}) {
	if err := j.validateSetCommitCommitterCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitCommitterCheck",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesA)SetCommitCommitterNameCheck(val interface{}) {
	if err := j.validateSetCommitCommitterNameCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitCommitterNameCheck",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesA)SetCommitMessageNegativeRegex(val *string) {
	if err := j.validateSetCommitMessageNegativeRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitMessageNegativeRegex",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesA)SetCommitMessageRegex(val *string) {
	if err := j.validateSetCommitMessageRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitMessageRegex",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesA)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesA)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesA)SetDenyDeleteTag(val interface{}) {
	if err := j.validateSetDenyDeleteTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"denyDeleteTag",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesA)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesA)SetFileNameRegex(val *string) {
	if err := j.validateSetFileNameRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileNameRegex",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesA)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesA)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesA)SetMaxFileSize(val *float64) {
	if err := j.validateSetMaxFileSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFileSize",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesA)SetMemberCheck(val interface{}) {
	if err := j.validateSetMemberCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memberCheck",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesA)SetPreventSecrets(val interface{}) {
	if err := j.validateSetPreventSecretsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preventSecrets",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesA)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesA)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesA)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesA)SetRejectNonDcoCommits(val interface{}) {
	if err := j.validateSetRejectNonDcoCommitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rejectNonDcoCommits",
		val,
	)
}

func (j *jsiiProxy_ProjectPushRulesA)SetRejectUnsignedCommits(val interface{}) {
	if err := j.validateSetRejectUnsignedCommitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rejectUnsignedCommits",
		val,
	)
}

// Generates CDKTF code for importing a ProjectPushRulesA resource upon running "cdktf plan <stack-name>".
func ProjectPushRulesA_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateProjectPushRulesA_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.projectPushRules.ProjectPushRulesA",
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
func ProjectPushRulesA_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjectPushRulesA_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.projectPushRules.ProjectPushRulesA",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ProjectPushRulesA_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjectPushRulesA_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.projectPushRules.ProjectPushRulesA",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ProjectPushRulesA_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjectPushRulesA_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.projectPushRules.ProjectPushRulesA",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ProjectPushRulesA_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.projectPushRules.ProjectPushRulesA",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_ProjectPushRulesA) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_ProjectPushRulesA) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_ProjectPushRulesA) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_ProjectPushRulesA) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_ProjectPushRulesA) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_ProjectPushRulesA) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_ProjectPushRulesA) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_ProjectPushRulesA) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_ProjectPushRulesA) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_ProjectPushRulesA) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_ProjectPushRulesA) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_ProjectPushRulesA) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectPushRulesA) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_ProjectPushRulesA) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_ProjectPushRulesA) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_ProjectPushRulesA) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_ProjectPushRulesA) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_ProjectPushRulesA) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_ProjectPushRulesA) ResetAuthorEmailRegex() {
	_jsii_.InvokeVoid(
		p,
		"resetAuthorEmailRegex",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectPushRulesA) ResetBranchNameRegex() {
	_jsii_.InvokeVoid(
		p,
		"resetBranchNameRegex",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectPushRulesA) ResetCommitCommitterCheck() {
	_jsii_.InvokeVoid(
		p,
		"resetCommitCommitterCheck",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectPushRulesA) ResetCommitCommitterNameCheck() {
	_jsii_.InvokeVoid(
		p,
		"resetCommitCommitterNameCheck",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectPushRulesA) ResetCommitMessageNegativeRegex() {
	_jsii_.InvokeVoid(
		p,
		"resetCommitMessageNegativeRegex",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectPushRulesA) ResetCommitMessageRegex() {
	_jsii_.InvokeVoid(
		p,
		"resetCommitMessageRegex",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectPushRulesA) ResetDenyDeleteTag() {
	_jsii_.InvokeVoid(
		p,
		"resetDenyDeleteTag",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectPushRulesA) ResetFileNameRegex() {
	_jsii_.InvokeVoid(
		p,
		"resetFileNameRegex",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectPushRulesA) ResetMaxFileSize() {
	_jsii_.InvokeVoid(
		p,
		"resetMaxFileSize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectPushRulesA) ResetMemberCheck() {
	_jsii_.InvokeVoid(
		p,
		"resetMemberCheck",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectPushRulesA) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectPushRulesA) ResetPreventSecrets() {
	_jsii_.InvokeVoid(
		p,
		"resetPreventSecrets",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectPushRulesA) ResetRejectNonDcoCommits() {
	_jsii_.InvokeVoid(
		p,
		"resetRejectNonDcoCommits",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectPushRulesA) ResetRejectUnsignedCommits() {
	_jsii_.InvokeVoid(
		p,
		"resetRejectUnsignedCommits",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectPushRulesA) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectPushRulesA) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectPushRulesA) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectPushRulesA) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectPushRulesA) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectPushRulesA) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

