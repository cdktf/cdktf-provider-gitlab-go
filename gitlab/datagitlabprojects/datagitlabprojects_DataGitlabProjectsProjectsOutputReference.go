package datagitlabprojects

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v6/jsii"

	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v6/datagitlabprojects/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabProjectsProjectsOutputReference interface {
	cdktf.ComplexObject
	AllowMergeOnSkippedPipeline() cdktf.IResolvable
	AnalyticsAccessLevel() *string
	ApprovalsBeforeMerge() *float64
	Archived() cdktf.IResolvable
	AutoCancelPendingPipelines() *string
	AutocloseReferencedIssues() cdktf.IResolvable
	AutoDevopsDeployStrategy() *string
	AutoDevopsEnabled() cdktf.IResolvable
	AvatarUrl() *string
	BuildCoverageRegex() *string
	BuildGitStrategy() *string
	BuildsAccessLevel() *string
	BuildTimeout() *float64
	CiConfigPath() *string
	CiDefaultGitDepth() *float64
	CiForwardDeploymentEnabled() cdktf.IResolvable
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
	ContainerExpirationPolicy() DataGitlabProjectsProjectsContainerExpirationPolicyList
	ContainerRegistryAccessLevel() *string
	ContainerRegistryEnabled() cdktf.IResolvable
	CreatedAt() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CreatorId() *float64
	CustomAttributes() cdktf.StringMapList
	DefaultBranch() *string
	Description() *string
	EmailsDisabled() cdktf.IResolvable
	ExternalAuthorizationClassificationLabel() *string
	ForkedFromProject() DataGitlabProjectsProjectsForkedFromProjectList
	ForkingAccessLevel() *string
	ForksCount() *float64
	// Experimental.
	Fqn() *string
	HttpUrlToRepo() *string
	Id() *float64
	ImportError() *string
	ImportStatus() *string
	InternalValue() *DataGitlabProjectsProjects
	SetInternalValue(val *DataGitlabProjectsProjects)
	IssuesAccessLevel() *string
	IssuesEnabled() cdktf.IResolvable
	JobsEnabled() cdktf.IResolvable
	LastActivityAt() *string
	LfsEnabled() cdktf.IResolvable
	Links() cdktf.StringMap
	MergeCommitTemplate() *string
	MergeMethod() *string
	MergePipelinesEnabled() cdktf.IResolvable
	MergeRequestsAccessLevel() *string
	MergeRequestsEnabled() cdktf.IResolvable
	MergeTrainsEnabled() cdktf.IResolvable
	Mirror() cdktf.IResolvable
	MirrorOverwritesDivergedBranches() cdktf.IResolvable
	MirrorTriggerBuilds() cdktf.IResolvable
	MirrorUserId() *float64
	Name() *string
	Namespace() DataGitlabProjectsProjectsNamespaceList
	NameWithNamespace() *string
	OnlyAllowMergeIfAllDiscussionsAreResolved() cdktf.IResolvable
	OnlyAllowMergeIfPipelineSucceeds() cdktf.IResolvable
	OnlyMirrorProtectedBranches() cdktf.IResolvable
	OpenIssuesCount() *float64
	OperationsAccessLevel() *string
	Owner() DataGitlabProjectsProjectsOwnerList
	PackagesEnabled() cdktf.IResolvable
	Path() *string
	PathWithNamespace() *string
	Permissions() DataGitlabProjectsProjectsPermissionsList
	Public() cdktf.IResolvable
	PublicBuilds() cdktf.IResolvable
	ReadmeUrl() *string
	RepositoryAccessLevel() *string
	RepositoryStorage() *string
	RequestAccessEnabled() cdktf.IResolvable
	RequirementsAccessLevel() *string
	ResolveOutdatedDiffDiscussions() cdktf.IResolvable
	RestrictUserDefinedVariables() cdktf.IResolvable
	RunnersToken() *string
	SecurityAndComplianceAccessLevel() *string
	SharedRunnersEnabled() cdktf.IResolvable
	SharedWithGroups() DataGitlabProjectsProjectsSharedWithGroupsList
	SnippetsAccessLevel() *string
	SnippetsEnabled() cdktf.IResolvable
	SquashCommitTemplate() *string
	SshUrlToRepo() *string
	StarCount() *float64
	Statistics() cdktf.NumberMap
	SuggestionCommitMessage() *string
	TagList() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Topics() *[]*string
	Visibility() *string
	WebUrl() *string
	WikiAccessLevel() *string
	WikiEnabled() cdktf.IResolvable
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

// The jsii proxy struct for DataGitlabProjectsProjectsOutputReference
type jsiiProxy_DataGitlabProjectsProjectsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) AllowMergeOnSkippedPipeline() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"allowMergeOnSkippedPipeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) AnalyticsAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analyticsAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) ApprovalsBeforeMerge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"approvalsBeforeMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) Archived() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"archived",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) AutoCancelPendingPipelines() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoCancelPendingPipelines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) AutocloseReferencedIssues() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autocloseReferencedIssues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) AutoDevopsDeployStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoDevopsDeployStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) AutoDevopsEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autoDevopsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) AvatarUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"avatarUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) BuildCoverageRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildCoverageRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) BuildGitStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildGitStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) BuildsAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildsAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) BuildTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"buildTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) CiConfigPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ciConfigPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) CiDefaultGitDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ciDefaultGitDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) CiForwardDeploymentEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ciForwardDeploymentEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) ContainerExpirationPolicy() DataGitlabProjectsProjectsContainerExpirationPolicyList {
	var returns DataGitlabProjectsProjectsContainerExpirationPolicyList
	_jsii_.Get(
		j,
		"containerExpirationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) ContainerRegistryAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRegistryAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) ContainerRegistryEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"containerRegistryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) CreatorId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creatorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) CustomAttributes() cdktf.StringMapList {
	var returns cdktf.StringMapList
	_jsii_.Get(
		j,
		"customAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) DefaultBranch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) EmailsDisabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"emailsDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) ExternalAuthorizationClassificationLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalAuthorizationClassificationLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) ForkedFromProject() DataGitlabProjectsProjectsForkedFromProjectList {
	var returns DataGitlabProjectsProjectsForkedFromProjectList
	_jsii_.Get(
		j,
		"forkedFromProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) ForkingAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forkingAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) ForksCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"forksCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) HttpUrlToRepo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpUrlToRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) ImportError() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) ImportStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) InternalValue() *DataGitlabProjectsProjects {
	var returns *DataGitlabProjectsProjects
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) IssuesAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuesAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) IssuesEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"issuesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) JobsEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"jobsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) LastActivityAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastActivityAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) LfsEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"lfsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) Links() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"links",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) MergeCommitTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeCommitTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) MergeMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) MergePipelinesEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"mergePipelinesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) MergeRequestsAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeRequestsAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) MergeRequestsEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"mergeRequestsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) MergeTrainsEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"mergeTrainsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) Mirror() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"mirror",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) MirrorOverwritesDivergedBranches() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"mirrorOverwritesDivergedBranches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) MirrorTriggerBuilds() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"mirrorTriggerBuilds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) MirrorUserId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mirrorUserId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) Namespace() DataGitlabProjectsProjectsNamespaceList {
	var returns DataGitlabProjectsProjectsNamespaceList
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) NameWithNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameWithNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) OnlyAllowMergeIfAllDiscussionsAreResolved() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"onlyAllowMergeIfAllDiscussionsAreResolved",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) OnlyAllowMergeIfPipelineSucceeds() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"onlyAllowMergeIfPipelineSucceeds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) OnlyMirrorProtectedBranches() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"onlyMirrorProtectedBranches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) OpenIssuesCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"openIssuesCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) OperationsAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationsAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) Owner() DataGitlabProjectsProjectsOwnerList {
	var returns DataGitlabProjectsProjectsOwnerList
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) PackagesEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"packagesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) PathWithNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathWithNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) Permissions() DataGitlabProjectsProjectsPermissionsList {
	var returns DataGitlabProjectsProjectsPermissionsList
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) Public() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"public",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) PublicBuilds() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"publicBuilds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) ReadmeUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readmeUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) RepositoryAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) RepositoryStorage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) RequestAccessEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"requestAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) RequirementsAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) ResolveOutdatedDiffDiscussions() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"resolveOutdatedDiffDiscussions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) RestrictUserDefinedVariables() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"restrictUserDefinedVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) RunnersToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runnersToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) SecurityAndComplianceAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityAndComplianceAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) SharedRunnersEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"sharedRunnersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) SharedWithGroups() DataGitlabProjectsProjectsSharedWithGroupsList {
	var returns DataGitlabProjectsProjectsSharedWithGroupsList
	_jsii_.Get(
		j,
		"sharedWithGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) SnippetsAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snippetsAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) SnippetsEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"snippetsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) SquashCommitTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"squashCommitTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) SshUrlToRepo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshUrlToRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) StarCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"starCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) Statistics() cdktf.NumberMap {
	var returns cdktf.NumberMap
	_jsii_.Get(
		j,
		"statistics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) SuggestionCommitMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suggestionCommitMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) TagList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) Topics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) Visibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) WebUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) WikiAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wikiAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference) WikiEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"wikiEnabled",
		&returns,
	)
	return returns
}


func NewDataGitlabProjectsProjectsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGitlabProjectsProjectsOutputReference {
	_init_.Initialize()

	if err := validateNewDataGitlabProjectsProjectsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGitlabProjectsProjectsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.dataGitlabProjects.DataGitlabProjectsProjectsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGitlabProjectsProjectsOutputReference_Override(d DataGitlabProjectsProjectsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.dataGitlabProjects.DataGitlabProjectsProjectsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference)SetInternalValue(val *DataGitlabProjectsProjects) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectsProjectsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGitlabProjectsProjectsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectsProjectsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGitlabProjectsProjectsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGitlabProjectsProjectsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGitlabProjectsProjectsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGitlabProjectsProjectsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGitlabProjectsProjectsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGitlabProjectsProjectsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGitlabProjectsProjectsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGitlabProjectsProjectsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGitlabProjectsProjectsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectsProjectsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGitlabProjectsProjectsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataGitlabProjectsProjectsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

