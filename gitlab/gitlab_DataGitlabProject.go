// Prebuilt gitlab Provider for Terraform CDK (cdktf)
package gitlab

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-gitlab-go/gitlab/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-gitlab-go/gitlab/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/gitlab/d/project gitlab_project}.
type DataGitlabProject interface {
	cdktf.TerraformDataSource
	AnalyticsAccessLevel() *string
	Archived() cdktf.IResolvable
	AutoCancelPendingPipelines() *string
	AutocloseReferencedIssues() cdktf.IResolvable
	AutoDevopsDeployStrategy() *string
	AutoDevopsEnabled() cdktf.IResolvable
	BuildGitStrategy() *string
	BuildsAccessLevel() *string
	BuildTimeout() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CiDefaultGitDepth() *float64
	SetCiDefaultGitDepth(val *float64)
	CiDefaultGitDepthInput() *float64
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerExpirationPolicy() DataGitlabProjectContainerExpirationPolicyList
	ContainerRegistryAccessLevel() *string
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	DefaultBranch() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	EmailsDisabled() cdktf.IResolvable
	ExternalAuthorizationClassificationLabel() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	ForkingAccessLevel() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpUrlToRepo() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IssuesAccessLevel() *string
	IssuesEnabled() cdktf.IResolvable
	LfsEnabled() cdktf.IResolvable
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MergeCommitTemplate() *string
	MergePipelinesEnabled() cdktf.IResolvable
	MergeRequestsAccessLevel() *string
	MergeRequestsEnabled() cdktf.IResolvable
	MergeTrainsEnabled() cdktf.IResolvable
	Name() *string
	NamespaceId() *float64
	// The tree node.
	Node() constructs.Node
	OperationsAccessLevel() *string
	Path() *string
	PathWithNamespace() *string
	SetPathWithNamespace(val *string)
	PathWithNamespaceInput() *string
	PipelinesEnabled() cdktf.IResolvable
	PrintingMergeRequestLinkEnabled() cdktf.IResolvable
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	PublicBuilds() interface{}
	SetPublicBuilds(val interface{})
	PublicBuildsInput() interface{}
	PushRules() DataGitlabProjectPushRulesList
	// Experimental.
	RawOverrides() interface{}
	RemoveSourceBranchAfterMerge() cdktf.IResolvable
	RepositoryAccessLevel() *string
	RepositoryStorage() *string
	RequestAccessEnabled() cdktf.IResolvable
	RequirementsAccessLevel() *string
	ResolveOutdatedDiffDiscussions() cdktf.IResolvable
	RunnersToken() *string
	SecurityAndComplianceAccessLevel() *string
	SnippetsAccessLevel() *string
	SnippetsEnabled() cdktf.IResolvable
	SquashCommitTemplate() *string
	SshUrlToRepo() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Topics() *[]*string
	VisibilityLevel() *string
	WebUrl() *string
	WikiAccessLevel() *string
	WikiEnabled() cdktf.IResolvable
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
	ResetCiDefaultGitDepth()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPathWithNamespace()
	ResetPublicBuilds()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataGitlabProject
type jsiiProxy_DataGitlabProject struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataGitlabProject) AnalyticsAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analyticsAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) Archived() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"archived",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) AutoCancelPendingPipelines() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoCancelPendingPipelines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) AutocloseReferencedIssues() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autocloseReferencedIssues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) AutoDevopsDeployStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoDevopsDeployStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) AutoDevopsEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autoDevopsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) BuildGitStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildGitStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) BuildsAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildsAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) BuildTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"buildTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) CiDefaultGitDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ciDefaultGitDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) CiDefaultGitDepthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ciDefaultGitDepthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) ContainerExpirationPolicy() DataGitlabProjectContainerExpirationPolicyList {
	var returns DataGitlabProjectContainerExpirationPolicyList
	_jsii_.Get(
		j,
		"containerExpirationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) ContainerRegistryAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRegistryAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) DefaultBranch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) EmailsDisabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"emailsDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) ExternalAuthorizationClassificationLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalAuthorizationClassificationLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) ForkingAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forkingAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) HttpUrlToRepo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpUrlToRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) IssuesAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuesAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) IssuesEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"issuesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) LfsEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"lfsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) MergeCommitTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeCommitTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) MergePipelinesEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"mergePipelinesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) MergeRequestsAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeRequestsAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) MergeRequestsEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"mergeRequestsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) MergeTrainsEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"mergeTrainsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) NamespaceId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"namespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) OperationsAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationsAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) PathWithNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathWithNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) PathWithNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathWithNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) PipelinesEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"pipelinesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) PrintingMergeRequestLinkEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"printingMergeRequestLinkEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) PublicBuilds() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicBuilds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) PublicBuildsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicBuildsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) PushRules() DataGitlabProjectPushRulesList {
	var returns DataGitlabProjectPushRulesList
	_jsii_.Get(
		j,
		"pushRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) RemoveSourceBranchAfterMerge() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"removeSourceBranchAfterMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) RepositoryAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) RepositoryStorage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) RequestAccessEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"requestAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) RequirementsAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) ResolveOutdatedDiffDiscussions() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"resolveOutdatedDiffDiscussions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) RunnersToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runnersToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) SecurityAndComplianceAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityAndComplianceAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) SnippetsAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snippetsAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) SnippetsEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"snippetsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) SquashCommitTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"squashCommitTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) SshUrlToRepo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshUrlToRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) Topics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) VisibilityLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibilityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) WebUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) WikiAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wikiAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProject) WikiEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"wikiEnabled",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/gitlab/d/project gitlab_project} Data Source.
func NewDataGitlabProject(scope constructs.Construct, id *string, config *DataGitlabProjectConfig) DataGitlabProject {
	_init_.Initialize()

	j := jsiiProxy_DataGitlabProject{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.DataGitlabProject",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/gitlab/d/project gitlab_project} Data Source.
func NewDataGitlabProject_Override(d DataGitlabProject, scope constructs.Construct, id *string, config *DataGitlabProjectConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.DataGitlabProject",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataGitlabProject) SetCiDefaultGitDepth(val *float64) {
	_jsii_.Set(
		j,
		"ciDefaultGitDepth",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProject) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProject) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProject) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProject) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProject) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProject) SetPathWithNamespace(val *string) {
	_jsii_.Set(
		j,
		"pathWithNamespace",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProject) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProject) SetPublicBuilds(val interface{}) {
	_jsii_.Set(
		j,
		"publicBuilds",
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
func DataGitlabProject_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.DataGitlabProject",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataGitlabProject_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.DataGitlabProject",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataGitlabProject) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataGitlabProject) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProject) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProject) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProject) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProject) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProject) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProject) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProject) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProject) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProject) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProject) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataGitlabProject) ResetCiDefaultGitDepth() {
	_jsii_.InvokeVoid(
		d,
		"resetCiDefaultGitDepth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProject) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProject) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProject) ResetPathWithNamespace() {
	_jsii_.InvokeVoid(
		d,
		"resetPathWithNamespace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProject) ResetPublicBuilds() {
	_jsii_.InvokeVoid(
		d,
		"resetPublicBuilds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGitlabProject) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProject) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProject) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProject) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

