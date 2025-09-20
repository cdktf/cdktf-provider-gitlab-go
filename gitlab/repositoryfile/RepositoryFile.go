// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositoryfile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v15/repositoryfile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/repository_file gitlab_repository_file}.
type RepositoryFile interface {
	cdktf.TerraformResource
	AuthorEmail() *string
	SetAuthorEmail(val *string)
	AuthorEmailInput() *string
	AuthorName() *string
	SetAuthorName(val *string)
	AuthorNameInput() *string
	BlobId() *string
	Branch() *string
	SetBranch(val *string)
	BranchInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CommitId() *string
	CommitMessage() *string
	SetCommitMessage(val *string)
	CommitMessageInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Content() *string
	SetContent(val *string)
	ContentInput() *string
	ContentSha256() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateCommitMessage() *string
	SetCreateCommitMessage(val *string)
	CreateCommitMessageInput() *string
	DeleteCommitMessage() *string
	SetDeleteCommitMessage(val *string)
	DeleteCommitMessageInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Encoding() *string
	SetEncoding(val *string)
	EncodingInput() *string
	ExecuteFilemode() interface{}
	SetExecuteFilemode(val interface{})
	ExecuteFilemodeInput() interface{}
	FileName() *string
	FilePath() *string
	SetFilePath(val *string)
	FilePathInput() *string
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
	LastCommitId() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OverwriteOnCreate() interface{}
	SetOverwriteOnCreate(val interface{})
	OverwriteOnCreateInput() interface{}
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
	Ref() *string
	Size() *float64
	StartBranch() *string
	SetStartBranch(val *string)
	StartBranchInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() RepositoryFileTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateCommitMessage() *string
	SetUpdateCommitMessage(val *string)
	UpdateCommitMessageInput() *string
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
	PutTimeouts(value *RepositoryFileTimeouts)
	ResetAuthorEmail()
	ResetAuthorName()
	ResetCommitMessage()
	ResetCreateCommitMessage()
	ResetDeleteCommitMessage()
	ResetExecuteFilemode()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOverwriteOnCreate()
	ResetStartBranch()
	ResetTimeouts()
	ResetUpdateCommitMessage()
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

// The jsii proxy struct for RepositoryFile
type jsiiProxy_RepositoryFile struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_RepositoryFile) AuthorEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) AuthorEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) AuthorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) AuthorNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) BlobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Branch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) BranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) CommitId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) CommitMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) CommitMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) ContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) ContentSha256() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentSha256",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) CreateCommitMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createCommitMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) CreateCommitMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createCommitMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) DeleteCommitMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteCommitMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) DeleteCommitMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteCommitMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Encoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) EncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) ExecuteFilemode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"executeFilemode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) ExecuteFilemodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"executeFilemodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) FileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) FilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) FilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) LastCommitId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastCommitId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) OverwriteOnCreate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overwriteOnCreate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) OverwriteOnCreateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overwriteOnCreateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Size() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) StartBranch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) StartBranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startBranchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Timeouts() RepositoryFileTimeoutsOutputReference {
	var returns RepositoryFileTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) UpdateCommitMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateCommitMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) UpdateCommitMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateCommitMessageInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/repository_file gitlab_repository_file} Resource.
func NewRepositoryFile(scope constructs.Construct, id *string, config *RepositoryFileConfig) RepositoryFile {
	_init_.Initialize()

	if err := validateNewRepositoryFileParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RepositoryFile{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.repositoryFile.RepositoryFile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/repository_file gitlab_repository_file} Resource.
func NewRepositoryFile_Override(r RepositoryFile, scope constructs.Construct, id *string, config *RepositoryFileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.repositoryFile.RepositoryFile",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RepositoryFile)SetAuthorEmail(val *string) {
	if err := j.validateSetAuthorEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorEmail",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetAuthorName(val *string) {
	if err := j.validateSetAuthorNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorName",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetBranch(val *string) {
	if err := j.validateSetBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branch",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetCommitMessage(val *string) {
	if err := j.validateSetCommitMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitMessage",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetContent(val *string) {
	if err := j.validateSetContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetCreateCommitMessage(val *string) {
	if err := j.validateSetCreateCommitMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createCommitMessage",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetDeleteCommitMessage(val *string) {
	if err := j.validateSetDeleteCommitMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteCommitMessage",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetEncoding(val *string) {
	if err := j.validateSetEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encoding",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetExecuteFilemode(val interface{}) {
	if err := j.validateSetExecuteFilemodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executeFilemode",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetFilePath(val *string) {
	if err := j.validateSetFilePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filePath",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetOverwriteOnCreate(val interface{}) {
	if err := j.validateSetOverwriteOnCreateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overwriteOnCreate",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetStartBranch(val *string) {
	if err := j.validateSetStartBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startBranch",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetUpdateCommitMessage(val *string) {
	if err := j.validateSetUpdateCommitMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateCommitMessage",
		val,
	)
}

// Generates CDKTF code for importing a RepositoryFile resource upon running "cdktf plan <stack-name>".
func RepositoryFile_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateRepositoryFile_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.repositoryFile.RepositoryFile",
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
func RepositoryFile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRepositoryFile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.repositoryFile.RepositoryFile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RepositoryFile_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRepositoryFile_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.repositoryFile.RepositoryFile",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RepositoryFile_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRepositoryFile_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.repositoryFile.RepositoryFile",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RepositoryFile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.repositoryFile.RepositoryFile",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RepositoryFile) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_RepositoryFile) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RepositoryFile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_RepositoryFile) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RepositoryFile) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_RepositoryFile) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RepositoryFile) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RepositoryFile) PutTimeouts(value *RepositoryFileTimeouts) {
	if err := r.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RepositoryFile) ResetAuthorEmail() {
	_jsii_.InvokeVoid(
		r,
		"resetAuthorEmail",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryFile) ResetAuthorName() {
	_jsii_.InvokeVoid(
		r,
		"resetAuthorName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryFile) ResetCommitMessage() {
	_jsii_.InvokeVoid(
		r,
		"resetCommitMessage",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryFile) ResetCreateCommitMessage() {
	_jsii_.InvokeVoid(
		r,
		"resetCreateCommitMessage",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryFile) ResetDeleteCommitMessage() {
	_jsii_.InvokeVoid(
		r,
		"resetDeleteCommitMessage",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryFile) ResetExecuteFilemode() {
	_jsii_.InvokeVoid(
		r,
		"resetExecuteFilemode",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryFile) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryFile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryFile) ResetOverwriteOnCreate() {
	_jsii_.InvokeVoid(
		r,
		"resetOverwriteOnCreate",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryFile) ResetStartBranch() {
	_jsii_.InvokeVoid(
		r,
		"resetStartBranch",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryFile) ResetTimeouts() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryFile) ResetUpdateCommitMessage() {
	_jsii_.InvokeVoid(
		r,
		"resetUpdateCommitMessage",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryFile) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

