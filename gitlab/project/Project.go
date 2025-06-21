// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package project

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v15/project/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/project gitlab_project}.
type Project interface {
	cdktf.TerraformResource
	AllowMergeOnSkippedPipeline() interface{}
	SetAllowMergeOnSkippedPipeline(val interface{})
	AllowMergeOnSkippedPipelineInput() interface{}
	AllowPipelineTriggerApproveDeployment() interface{}
	SetAllowPipelineTriggerApproveDeployment(val interface{})
	AllowPipelineTriggerApproveDeploymentInput() interface{}
	AnalyticsAccessLevel() *string
	SetAnalyticsAccessLevel(val *string)
	AnalyticsAccessLevelInput() *string
	ApprovalsBeforeMerge() *float64
	SetApprovalsBeforeMerge(val *float64)
	ApprovalsBeforeMergeInput() *float64
	Archived() interface{}
	SetArchived(val interface{})
	ArchivedInput() interface{}
	ArchiveOnDestroy() interface{}
	SetArchiveOnDestroy(val interface{})
	ArchiveOnDestroyInput() interface{}
	AutoCancelPendingPipelines() *string
	SetAutoCancelPendingPipelines(val *string)
	AutoCancelPendingPipelinesInput() *string
	AutocloseReferencedIssues() interface{}
	SetAutocloseReferencedIssues(val interface{})
	AutocloseReferencedIssuesInput() interface{}
	AutoDevopsDeployStrategy() *string
	SetAutoDevopsDeployStrategy(val *string)
	AutoDevopsDeployStrategyInput() *string
	AutoDevopsEnabled() interface{}
	SetAutoDevopsEnabled(val interface{})
	AutoDevopsEnabledInput() interface{}
	AutoDuoCodeReviewEnabled() interface{}
	SetAutoDuoCodeReviewEnabled(val interface{})
	AutoDuoCodeReviewEnabledInput() interface{}
	Avatar() *string
	SetAvatar(val *string)
	AvatarHash() *string
	SetAvatarHash(val *string)
	AvatarHashInput() *string
	AvatarInput() *string
	AvatarUrl() *string
	Branches() *string
	SetBranches(val *string)
	BranchesInput() *string
	BuildGitStrategy() *string
	SetBuildGitStrategy(val *string)
	BuildGitStrategyInput() *string
	BuildsAccessLevel() *string
	SetBuildsAccessLevel(val *string)
	BuildsAccessLevelInput() *string
	BuildTimeout() *float64
	SetBuildTimeout(val *float64)
	BuildTimeoutInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CiConfigPath() *string
	SetCiConfigPath(val *string)
	CiConfigPathInput() *string
	CiDefaultGitDepth() *float64
	SetCiDefaultGitDepth(val *float64)
	CiDefaultGitDepthInput() *float64
	CiDeletePipelinesInSeconds() *float64
	SetCiDeletePipelinesInSeconds(val *float64)
	CiDeletePipelinesInSecondsInput() *float64
	CiForwardDeploymentEnabled() interface{}
	SetCiForwardDeploymentEnabled(val interface{})
	CiForwardDeploymentEnabledInput() interface{}
	CiForwardDeploymentRollbackAllowed() interface{}
	SetCiForwardDeploymentRollbackAllowed(val interface{})
	CiForwardDeploymentRollbackAllowedInput() interface{}
	CiIdTokenSubClaimComponents() *[]*string
	SetCiIdTokenSubClaimComponents(val *[]*string)
	CiIdTokenSubClaimComponentsInput() *[]*string
	CiPipelineVariablesMinimumOverrideRole() *string
	SetCiPipelineVariablesMinimumOverrideRole(val *string)
	CiPipelineVariablesMinimumOverrideRoleInput() *string
	CiRestrictPipelineCancellationRole() *string
	SetCiRestrictPipelineCancellationRole(val *string)
	CiRestrictPipelineCancellationRoleInput() *string
	CiSeparatedCaches() interface{}
	SetCiSeparatedCaches(val interface{})
	CiSeparatedCachesInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerExpirationPolicy() ProjectContainerExpirationPolicyOutputReference
	ContainerExpirationPolicyInput() *ProjectContainerExpirationPolicy
	ContainerRegistryAccessLevel() *string
	SetContainerRegistryAccessLevel(val *string)
	ContainerRegistryAccessLevelInput() *string
	ContainerRegistryEnabled() interface{}
	SetContainerRegistryEnabled(val interface{})
	ContainerRegistryEnabledInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DefaultBranch() *string
	SetDefaultBranch(val *string)
	DefaultBranchInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EmailsEnabled() interface{}
	SetEmailsEnabled(val interface{})
	EmailsEnabledInput() interface{}
	EmptyRepo() cdktf.IResolvable
	EnvironmentsAccessLevel() *string
	SetEnvironmentsAccessLevel(val *string)
	EnvironmentsAccessLevelInput() *string
	ExternalAuthorizationClassificationLabel() *string
	SetExternalAuthorizationClassificationLabel(val *string)
	ExternalAuthorizationClassificationLabelInput() *string
	FeatureFlagsAccessLevel() *string
	SetFeatureFlagsAccessLevel(val *string)
	FeatureFlagsAccessLevelInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	ForkedFromProjectId() *float64
	SetForkedFromProjectId(val *float64)
	ForkedFromProjectIdInput() *float64
	ForkingAccessLevel() *string
	SetForkingAccessLevel(val *string)
	ForkingAccessLevelInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GroupRunnersEnabled() interface{}
	SetGroupRunnersEnabled(val interface{})
	GroupRunnersEnabledInput() interface{}
	GroupWithProjectTemplatesId() *float64
	SetGroupWithProjectTemplatesId(val *float64)
	GroupWithProjectTemplatesIdInput() *float64
	HttpUrlToRepo() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImportUrl() *string
	SetImportUrl(val *string)
	ImportUrlInput() *string
	ImportUrlPassword() *string
	SetImportUrlPassword(val *string)
	ImportUrlPasswordInput() *string
	ImportUrlUsername() *string
	SetImportUrlUsername(val *string)
	ImportUrlUsernameInput() *string
	InfrastructureAccessLevel() *string
	SetInfrastructureAccessLevel(val *string)
	InfrastructureAccessLevelInput() *string
	InitializeWithReadme() interface{}
	SetInitializeWithReadme(val interface{})
	InitializeWithReadmeInput() interface{}
	IssuesAccessLevel() *string
	SetIssuesAccessLevel(val *string)
	IssuesAccessLevelInput() *string
	IssuesEnabled() interface{}
	SetIssuesEnabled(val interface{})
	IssuesEnabledInput() interface{}
	IssuesTemplate() *string
	SetIssuesTemplate(val *string)
	IssuesTemplateInput() *string
	KeepLatestArtifact() interface{}
	SetKeepLatestArtifact(val interface{})
	KeepLatestArtifactInput() interface{}
	LfsEnabled() interface{}
	SetLfsEnabled(val interface{})
	LfsEnabledInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MergeCommitTemplate() *string
	SetMergeCommitTemplate(val *string)
	MergeCommitTemplateInput() *string
	MergeMethod() *string
	SetMergeMethod(val *string)
	MergeMethodInput() *string
	MergePipelinesEnabled() interface{}
	SetMergePipelinesEnabled(val interface{})
	MergePipelinesEnabledInput() interface{}
	MergeRequestsAccessLevel() *string
	SetMergeRequestsAccessLevel(val *string)
	MergeRequestsAccessLevelInput() *string
	MergeRequestsEnabled() interface{}
	SetMergeRequestsEnabled(val interface{})
	MergeRequestsEnabledInput() interface{}
	MergeRequestsTemplate() *string
	SetMergeRequestsTemplate(val *string)
	MergeRequestsTemplateInput() *string
	MergeTrainsEnabled() interface{}
	SetMergeTrainsEnabled(val interface{})
	MergeTrainsEnabledInput() interface{}
	Mirror() interface{}
	SetMirror(val interface{})
	MirrorInput() interface{}
	MirrorOverwritesDivergedBranches() interface{}
	SetMirrorOverwritesDivergedBranches(val interface{})
	MirrorOverwritesDivergedBranchesInput() interface{}
	MirrorTriggerBuilds() interface{}
	SetMirrorTriggerBuilds(val interface{})
	MirrorTriggerBuildsInput() interface{}
	ModelExperimentsAccessLevel() *string
	SetModelExperimentsAccessLevel(val *string)
	ModelExperimentsAccessLevelInput() *string
	ModelRegistryAccessLevel() *string
	SetModelRegistryAccessLevel(val *string)
	ModelRegistryAccessLevelInput() *string
	MonitorAccessLevel() *string
	SetMonitorAccessLevel(val *string)
	MonitorAccessLevelInput() *string
	MrDefaultTargetSelf() interface{}
	SetMrDefaultTargetSelf(val interface{})
	MrDefaultTargetSelfInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamespaceId() *float64
	SetNamespaceId(val *float64)
	NamespaceIdInput() *float64
	// The tree node.
	Node() constructs.Node
	OnlyAllowMergeIfAllDiscussionsAreResolved() interface{}
	SetOnlyAllowMergeIfAllDiscussionsAreResolved(val interface{})
	OnlyAllowMergeIfAllDiscussionsAreResolvedInput() interface{}
	OnlyAllowMergeIfPipelineSucceeds() interface{}
	SetOnlyAllowMergeIfPipelineSucceeds(val interface{})
	OnlyAllowMergeIfPipelineSucceedsInput() interface{}
	OnlyMirrorProtectedBranches() interface{}
	SetOnlyMirrorProtectedBranches(val interface{})
	OnlyMirrorProtectedBranchesInput() interface{}
	PackagesEnabled() interface{}
	SetPackagesEnabled(val interface{})
	PackagesEnabledInput() interface{}
	PagesAccessLevel() *string
	SetPagesAccessLevel(val *string)
	PagesAccessLevelInput() *string
	Path() *string
	SetPath(val *string)
	PathInput() *string
	PathWithNamespace() *string
	PermanentlyDeleteOnDestroy() interface{}
	SetPermanentlyDeleteOnDestroy(val interface{})
	PermanentlyDeleteOnDestroyInput() interface{}
	PipelinesEnabled() interface{}
	SetPipelinesEnabled(val interface{})
	PipelinesEnabledInput() interface{}
	PreReceiveSecretDetectionEnabled() interface{}
	SetPreReceiveSecretDetectionEnabled(val interface{})
	PreReceiveSecretDetectionEnabledInput() interface{}
	PreventMergeWithoutJiraIssue() interface{}
	SetPreventMergeWithoutJiraIssue(val interface{})
	PreventMergeWithoutJiraIssueInput() interface{}
	PrintingMergeRequestLinkEnabled() interface{}
	SetPrintingMergeRequestLinkEnabled(val interface{})
	PrintingMergeRequestLinkEnabledInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicBuilds() interface{}
	SetPublicBuilds(val interface{})
	PublicBuildsInput() interface{}
	PublicJobs() interface{}
	SetPublicJobs(val interface{})
	PublicJobsInput() interface{}
	PushRules() ProjectPushRulesOutputReference
	PushRulesInput() *ProjectPushRules
	// Experimental.
	RawOverrides() interface{}
	ReleasesAccessLevel() *string
	SetReleasesAccessLevel(val *string)
	ReleasesAccessLevelInput() *string
	RemoveSourceBranchAfterMerge() interface{}
	SetRemoveSourceBranchAfterMerge(val interface{})
	RemoveSourceBranchAfterMergeInput() interface{}
	RepositoryAccessLevel() *string
	SetRepositoryAccessLevel(val *string)
	RepositoryAccessLevelInput() *string
	RepositoryStorage() *string
	SetRepositoryStorage(val *string)
	RepositoryStorageInput() *string
	RequestAccessEnabled() interface{}
	SetRequestAccessEnabled(val interface{})
	RequestAccessEnabledInput() interface{}
	RequirementsAccessLevel() *string
	SetRequirementsAccessLevel(val *string)
	RequirementsAccessLevelInput() *string
	ResolveOutdatedDiffDiscussions() interface{}
	SetResolveOutdatedDiffDiscussions(val interface{})
	ResolveOutdatedDiffDiscussionsInput() interface{}
	RestrictUserDefinedVariables() interface{}
	SetRestrictUserDefinedVariables(val interface{})
	RestrictUserDefinedVariablesInput() interface{}
	RunnersToken() *string
	SecurityAndComplianceAccessLevel() *string
	SetSecurityAndComplianceAccessLevel(val *string)
	SecurityAndComplianceAccessLevelInput() *string
	SharedRunnersEnabled() interface{}
	SetSharedRunnersEnabled(val interface{})
	SharedRunnersEnabledInput() interface{}
	SkipWaitForDefaultBranchProtection() interface{}
	SetSkipWaitForDefaultBranchProtection(val interface{})
	SkipWaitForDefaultBranchProtectionInput() interface{}
	SnippetsAccessLevel() *string
	SetSnippetsAccessLevel(val *string)
	SnippetsAccessLevelInput() *string
	SnippetsEnabled() interface{}
	SetSnippetsEnabled(val interface{})
	SnippetsEnabledInput() interface{}
	SquashCommitTemplate() *string
	SetSquashCommitTemplate(val *string)
	SquashCommitTemplateInput() *string
	SquashOption() *string
	SetSquashOption(val *string)
	SquashOptionInput() *string
	SshUrlToRepo() *string
	SuggestionCommitMessage() *string
	SetSuggestionCommitMessage(val *string)
	SuggestionCommitMessageInput() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	TemplateName() *string
	SetTemplateName(val *string)
	TemplateNameInput() *string
	TemplateProjectId() *float64
	SetTemplateProjectId(val *float64)
	TemplateProjectIdInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ProjectTimeoutsOutputReference
	TimeoutsInput() interface{}
	Topics() *[]*string
	SetTopics(val *[]*string)
	TopicsInput() *[]*string
	UseCustomTemplate() interface{}
	SetUseCustomTemplate(val interface{})
	UseCustomTemplateInput() interface{}
	VisibilityLevel() *string
	SetVisibilityLevel(val *string)
	VisibilityLevelInput() *string
	WebUrl() *string
	WikiAccessLevel() *string
	SetWikiAccessLevel(val *string)
	WikiAccessLevelInput() *string
	WikiEnabled() interface{}
	SetWikiEnabled(val interface{})
	WikiEnabledInput() interface{}
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
	PutContainerExpirationPolicy(value *ProjectContainerExpirationPolicy)
	PutPushRules(value *ProjectPushRules)
	PutTimeouts(value *ProjectTimeouts)
	ResetAllowMergeOnSkippedPipeline()
	ResetAllowPipelineTriggerApproveDeployment()
	ResetAnalyticsAccessLevel()
	ResetApprovalsBeforeMerge()
	ResetArchived()
	ResetArchiveOnDestroy()
	ResetAutoCancelPendingPipelines()
	ResetAutocloseReferencedIssues()
	ResetAutoDevopsDeployStrategy()
	ResetAutoDevopsEnabled()
	ResetAutoDuoCodeReviewEnabled()
	ResetAvatar()
	ResetAvatarHash()
	ResetBranches()
	ResetBuildGitStrategy()
	ResetBuildsAccessLevel()
	ResetBuildTimeout()
	ResetCiConfigPath()
	ResetCiDefaultGitDepth()
	ResetCiDeletePipelinesInSeconds()
	ResetCiForwardDeploymentEnabled()
	ResetCiForwardDeploymentRollbackAllowed()
	ResetCiIdTokenSubClaimComponents()
	ResetCiPipelineVariablesMinimumOverrideRole()
	ResetCiRestrictPipelineCancellationRole()
	ResetCiSeparatedCaches()
	ResetContainerExpirationPolicy()
	ResetContainerRegistryAccessLevel()
	ResetContainerRegistryEnabled()
	ResetDefaultBranch()
	ResetDescription()
	ResetEmailsEnabled()
	ResetEnvironmentsAccessLevel()
	ResetExternalAuthorizationClassificationLabel()
	ResetFeatureFlagsAccessLevel()
	ResetForkedFromProjectId()
	ResetForkingAccessLevel()
	ResetGroupRunnersEnabled()
	ResetGroupWithProjectTemplatesId()
	ResetId()
	ResetImportUrl()
	ResetImportUrlPassword()
	ResetImportUrlUsername()
	ResetInfrastructureAccessLevel()
	ResetInitializeWithReadme()
	ResetIssuesAccessLevel()
	ResetIssuesEnabled()
	ResetIssuesTemplate()
	ResetKeepLatestArtifact()
	ResetLfsEnabled()
	ResetMergeCommitTemplate()
	ResetMergeMethod()
	ResetMergePipelinesEnabled()
	ResetMergeRequestsAccessLevel()
	ResetMergeRequestsEnabled()
	ResetMergeRequestsTemplate()
	ResetMergeTrainsEnabled()
	ResetMirror()
	ResetMirrorOverwritesDivergedBranches()
	ResetMirrorTriggerBuilds()
	ResetModelExperimentsAccessLevel()
	ResetModelRegistryAccessLevel()
	ResetMonitorAccessLevel()
	ResetMrDefaultTargetSelf()
	ResetNamespaceId()
	ResetOnlyAllowMergeIfAllDiscussionsAreResolved()
	ResetOnlyAllowMergeIfPipelineSucceeds()
	ResetOnlyMirrorProtectedBranches()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPackagesEnabled()
	ResetPagesAccessLevel()
	ResetPath()
	ResetPermanentlyDeleteOnDestroy()
	ResetPipelinesEnabled()
	ResetPreReceiveSecretDetectionEnabled()
	ResetPreventMergeWithoutJiraIssue()
	ResetPrintingMergeRequestLinkEnabled()
	ResetPublicBuilds()
	ResetPublicJobs()
	ResetPushRules()
	ResetReleasesAccessLevel()
	ResetRemoveSourceBranchAfterMerge()
	ResetRepositoryAccessLevel()
	ResetRepositoryStorage()
	ResetRequestAccessEnabled()
	ResetRequirementsAccessLevel()
	ResetResolveOutdatedDiffDiscussions()
	ResetRestrictUserDefinedVariables()
	ResetSecurityAndComplianceAccessLevel()
	ResetSharedRunnersEnabled()
	ResetSkipWaitForDefaultBranchProtection()
	ResetSnippetsAccessLevel()
	ResetSnippetsEnabled()
	ResetSquashCommitTemplate()
	ResetSquashOption()
	ResetSuggestionCommitMessage()
	ResetTags()
	ResetTemplateName()
	ResetTemplateProjectId()
	ResetTimeouts()
	ResetTopics()
	ResetUseCustomTemplate()
	ResetVisibilityLevel()
	ResetWikiAccessLevel()
	ResetWikiEnabled()
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

// The jsii proxy struct for Project
type jsiiProxy_Project struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Project) AllowMergeOnSkippedPipeline() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMergeOnSkippedPipeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) AllowMergeOnSkippedPipelineInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMergeOnSkippedPipelineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) AllowPipelineTriggerApproveDeployment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPipelineTriggerApproveDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) AllowPipelineTriggerApproveDeploymentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPipelineTriggerApproveDeploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) AnalyticsAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analyticsAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) AnalyticsAccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analyticsAccessLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ApprovalsBeforeMerge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"approvalsBeforeMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ApprovalsBeforeMergeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"approvalsBeforeMergeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Archived() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archived",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ArchivedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archivedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ArchiveOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archiveOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ArchiveOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archiveOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) AutoCancelPendingPipelines() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoCancelPendingPipelines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) AutoCancelPendingPipelinesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoCancelPendingPipelinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) AutocloseReferencedIssues() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autocloseReferencedIssues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) AutocloseReferencedIssuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autocloseReferencedIssuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) AutoDevopsDeployStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoDevopsDeployStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) AutoDevopsDeployStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoDevopsDeployStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) AutoDevopsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDevopsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) AutoDevopsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDevopsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) AutoDuoCodeReviewEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDuoCodeReviewEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) AutoDuoCodeReviewEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDuoCodeReviewEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Avatar() *string {
	var returns *string
	_jsii_.Get(
		j,
		"avatar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) AvatarHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"avatarHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) AvatarHashInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"avatarHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) AvatarInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"avatarInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) AvatarUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"avatarUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Branches() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) BranchesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) BuildGitStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildGitStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) BuildGitStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildGitStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) BuildsAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildsAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) BuildsAccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildsAccessLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) BuildTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"buildTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) BuildTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"buildTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) CiConfigPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ciConfigPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) CiConfigPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ciConfigPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) CiDefaultGitDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ciDefaultGitDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) CiDefaultGitDepthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ciDefaultGitDepthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) CiDeletePipelinesInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ciDeletePipelinesInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) CiDeletePipelinesInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ciDeletePipelinesInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) CiForwardDeploymentEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ciForwardDeploymentEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) CiForwardDeploymentEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ciForwardDeploymentEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) CiForwardDeploymentRollbackAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ciForwardDeploymentRollbackAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) CiForwardDeploymentRollbackAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ciForwardDeploymentRollbackAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) CiIdTokenSubClaimComponents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ciIdTokenSubClaimComponents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) CiIdTokenSubClaimComponentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ciIdTokenSubClaimComponentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) CiPipelineVariablesMinimumOverrideRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ciPipelineVariablesMinimumOverrideRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) CiPipelineVariablesMinimumOverrideRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ciPipelineVariablesMinimumOverrideRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) CiRestrictPipelineCancellationRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ciRestrictPipelineCancellationRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) CiRestrictPipelineCancellationRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ciRestrictPipelineCancellationRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) CiSeparatedCaches() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ciSeparatedCaches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) CiSeparatedCachesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ciSeparatedCachesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ContainerExpirationPolicy() ProjectContainerExpirationPolicyOutputReference {
	var returns ProjectContainerExpirationPolicyOutputReference
	_jsii_.Get(
		j,
		"containerExpirationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ContainerExpirationPolicyInput() *ProjectContainerExpirationPolicy {
	var returns *ProjectContainerExpirationPolicy
	_jsii_.Get(
		j,
		"containerExpirationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ContainerRegistryAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRegistryAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ContainerRegistryAccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRegistryAccessLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ContainerRegistryEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerRegistryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ContainerRegistryEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerRegistryEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) DefaultBranch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) DefaultBranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBranchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) EmailsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) EmailsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) EmptyRepo() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"emptyRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) EnvironmentsAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentsAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) EnvironmentsAccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentsAccessLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ExternalAuthorizationClassificationLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalAuthorizationClassificationLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ExternalAuthorizationClassificationLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalAuthorizationClassificationLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) FeatureFlagsAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featureFlagsAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) FeatureFlagsAccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featureFlagsAccessLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ForkedFromProjectId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"forkedFromProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ForkedFromProjectIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"forkedFromProjectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ForkingAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forkingAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ForkingAccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forkingAccessLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) GroupRunnersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupRunnersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) GroupRunnersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupRunnersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) GroupWithProjectTemplatesId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupWithProjectTemplatesId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) GroupWithProjectTemplatesIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupWithProjectTemplatesIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) HttpUrlToRepo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpUrlToRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ImportUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ImportUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ImportUrlPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importUrlPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ImportUrlPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importUrlPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ImportUrlUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importUrlUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ImportUrlUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importUrlUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) InfrastructureAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) InfrastructureAccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureAccessLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) InitializeWithReadme() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initializeWithReadme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) InitializeWithReadmeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initializeWithReadmeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) IssuesAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuesAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) IssuesAccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuesAccessLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) IssuesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issuesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) IssuesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issuesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) IssuesTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuesTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) IssuesTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuesTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) KeepLatestArtifact() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepLatestArtifact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) KeepLatestArtifactInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepLatestArtifactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) LfsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lfsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) LfsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lfsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) MergeCommitTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeCommitTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) MergeCommitTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeCommitTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) MergeMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) MergeMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) MergePipelinesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergePipelinesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) MergePipelinesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergePipelinesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) MergeRequestsAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeRequestsAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) MergeRequestsAccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeRequestsAccessLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) MergeRequestsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeRequestsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) MergeRequestsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeRequestsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) MergeRequestsTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeRequestsTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) MergeRequestsTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeRequestsTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) MergeTrainsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeTrainsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) MergeTrainsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeTrainsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Mirror() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mirror",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) MirrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mirrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) MirrorOverwritesDivergedBranches() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mirrorOverwritesDivergedBranches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) MirrorOverwritesDivergedBranchesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mirrorOverwritesDivergedBranchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) MirrorTriggerBuilds() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mirrorTriggerBuilds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) MirrorTriggerBuildsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mirrorTriggerBuildsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ModelExperimentsAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelExperimentsAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ModelExperimentsAccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelExperimentsAccessLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ModelRegistryAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelRegistryAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ModelRegistryAccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelRegistryAccessLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) MonitorAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitorAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) MonitorAccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitorAccessLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) MrDefaultTargetSelf() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mrDefaultTargetSelf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) MrDefaultTargetSelfInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mrDefaultTargetSelfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) NamespaceId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"namespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) NamespaceIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"namespaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) OnlyAllowMergeIfAllDiscussionsAreResolved() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlyAllowMergeIfAllDiscussionsAreResolved",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) OnlyAllowMergeIfAllDiscussionsAreResolvedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlyAllowMergeIfAllDiscussionsAreResolvedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) OnlyAllowMergeIfPipelineSucceeds() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlyAllowMergeIfPipelineSucceeds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) OnlyAllowMergeIfPipelineSucceedsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlyAllowMergeIfPipelineSucceedsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) OnlyMirrorProtectedBranches() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlyMirrorProtectedBranches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) OnlyMirrorProtectedBranchesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlyMirrorProtectedBranchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PackagesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"packagesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PackagesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"packagesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PagesAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pagesAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PagesAccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pagesAccessLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PathWithNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathWithNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PermanentlyDeleteOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permanentlyDeleteOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PermanentlyDeleteOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permanentlyDeleteOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PipelinesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pipelinesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PipelinesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pipelinesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PreReceiveSecretDetectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preReceiveSecretDetectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PreReceiveSecretDetectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preReceiveSecretDetectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PreventMergeWithoutJiraIssue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventMergeWithoutJiraIssue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PreventMergeWithoutJiraIssueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventMergeWithoutJiraIssueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PrintingMergeRequestLinkEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"printingMergeRequestLinkEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PrintingMergeRequestLinkEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"printingMergeRequestLinkEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PublicBuilds() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicBuilds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PublicBuildsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicBuildsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PublicJobs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicJobs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PublicJobsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicJobsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PushRules() ProjectPushRulesOutputReference {
	var returns ProjectPushRulesOutputReference
	_jsii_.Get(
		j,
		"pushRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PushRulesInput() *ProjectPushRules {
	var returns *ProjectPushRules
	_jsii_.Get(
		j,
		"pushRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ReleasesAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releasesAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ReleasesAccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releasesAccessLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) RemoveSourceBranchAfterMerge() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeSourceBranchAfterMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) RemoveSourceBranchAfterMergeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeSourceBranchAfterMergeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) RepositoryAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) RepositoryAccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryAccessLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) RepositoryStorage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) RepositoryStorageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) RequestAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) RequestAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) RequirementsAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) RequirementsAccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsAccessLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ResolveOutdatedDiffDiscussions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resolveOutdatedDiffDiscussions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ResolveOutdatedDiffDiscussionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resolveOutdatedDiffDiscussionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) RestrictUserDefinedVariables() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictUserDefinedVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) RestrictUserDefinedVariablesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictUserDefinedVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) RunnersToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runnersToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) SecurityAndComplianceAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityAndComplianceAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) SecurityAndComplianceAccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityAndComplianceAccessLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) SharedRunnersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharedRunnersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) SharedRunnersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharedRunnersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) SkipWaitForDefaultBranchProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipWaitForDefaultBranchProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) SkipWaitForDefaultBranchProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipWaitForDefaultBranchProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) SnippetsAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snippetsAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) SnippetsAccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snippetsAccessLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) SnippetsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snippetsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) SnippetsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snippetsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) SquashCommitTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"squashCommitTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) SquashCommitTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"squashCommitTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) SquashOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"squashOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) SquashOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"squashOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) SshUrlToRepo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshUrlToRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) SuggestionCommitMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suggestionCommitMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) SuggestionCommitMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suggestionCommitMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) TemplateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) TemplateNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) TemplateProjectId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"templateProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) TemplateProjectIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"templateProjectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Timeouts() ProjectTimeoutsOutputReference {
	var returns ProjectTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Topics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) TopicsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) UseCustomTemplate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCustomTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) UseCustomTemplateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCustomTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) VisibilityLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibilityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) VisibilityLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibilityLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) WebUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) WikiAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wikiAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) WikiAccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wikiAccessLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) WikiEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wikiEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) WikiEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wikiEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/project gitlab_project} Resource.
func NewProject(scope constructs.Construct, id *string, config *ProjectConfig) Project {
	_init_.Initialize()

	if err := validateNewProjectParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Project{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.project.Project",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/project gitlab_project} Resource.
func NewProject_Override(p Project, scope constructs.Construct, id *string, config *ProjectConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.project.Project",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_Project)SetAllowMergeOnSkippedPipeline(val interface{}) {
	if err := j.validateSetAllowMergeOnSkippedPipelineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowMergeOnSkippedPipeline",
		val,
	)
}

func (j *jsiiProxy_Project)SetAllowPipelineTriggerApproveDeployment(val interface{}) {
	if err := j.validateSetAllowPipelineTriggerApproveDeploymentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowPipelineTriggerApproveDeployment",
		val,
	)
}

func (j *jsiiProxy_Project)SetAnalyticsAccessLevel(val *string) {
	if err := j.validateSetAnalyticsAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"analyticsAccessLevel",
		val,
	)
}

func (j *jsiiProxy_Project)SetApprovalsBeforeMerge(val *float64) {
	if err := j.validateSetApprovalsBeforeMergeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approvalsBeforeMerge",
		val,
	)
}

func (j *jsiiProxy_Project)SetArchived(val interface{}) {
	if err := j.validateSetArchivedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"archived",
		val,
	)
}

func (j *jsiiProxy_Project)SetArchiveOnDestroy(val interface{}) {
	if err := j.validateSetArchiveOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"archiveOnDestroy",
		val,
	)
}

func (j *jsiiProxy_Project)SetAutoCancelPendingPipelines(val *string) {
	if err := j.validateSetAutoCancelPendingPipelinesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoCancelPendingPipelines",
		val,
	)
}

func (j *jsiiProxy_Project)SetAutocloseReferencedIssues(val interface{}) {
	if err := j.validateSetAutocloseReferencedIssuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autocloseReferencedIssues",
		val,
	)
}

func (j *jsiiProxy_Project)SetAutoDevopsDeployStrategy(val *string) {
	if err := j.validateSetAutoDevopsDeployStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoDevopsDeployStrategy",
		val,
	)
}

func (j *jsiiProxy_Project)SetAutoDevopsEnabled(val interface{}) {
	if err := j.validateSetAutoDevopsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoDevopsEnabled",
		val,
	)
}

func (j *jsiiProxy_Project)SetAutoDuoCodeReviewEnabled(val interface{}) {
	if err := j.validateSetAutoDuoCodeReviewEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoDuoCodeReviewEnabled",
		val,
	)
}

func (j *jsiiProxy_Project)SetAvatar(val *string) {
	if err := j.validateSetAvatarParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"avatar",
		val,
	)
}

func (j *jsiiProxy_Project)SetAvatarHash(val *string) {
	if err := j.validateSetAvatarHashParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"avatarHash",
		val,
	)
}

func (j *jsiiProxy_Project)SetBranches(val *string) {
	if err := j.validateSetBranchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branches",
		val,
	)
}

func (j *jsiiProxy_Project)SetBuildGitStrategy(val *string) {
	if err := j.validateSetBuildGitStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildGitStrategy",
		val,
	)
}

func (j *jsiiProxy_Project)SetBuildsAccessLevel(val *string) {
	if err := j.validateSetBuildsAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildsAccessLevel",
		val,
	)
}

func (j *jsiiProxy_Project)SetBuildTimeout(val *float64) {
	if err := j.validateSetBuildTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildTimeout",
		val,
	)
}

func (j *jsiiProxy_Project)SetCiConfigPath(val *string) {
	if err := j.validateSetCiConfigPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ciConfigPath",
		val,
	)
}

func (j *jsiiProxy_Project)SetCiDefaultGitDepth(val *float64) {
	if err := j.validateSetCiDefaultGitDepthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ciDefaultGitDepth",
		val,
	)
}

func (j *jsiiProxy_Project)SetCiDeletePipelinesInSeconds(val *float64) {
	if err := j.validateSetCiDeletePipelinesInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ciDeletePipelinesInSeconds",
		val,
	)
}

func (j *jsiiProxy_Project)SetCiForwardDeploymentEnabled(val interface{}) {
	if err := j.validateSetCiForwardDeploymentEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ciForwardDeploymentEnabled",
		val,
	)
}

func (j *jsiiProxy_Project)SetCiForwardDeploymentRollbackAllowed(val interface{}) {
	if err := j.validateSetCiForwardDeploymentRollbackAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ciForwardDeploymentRollbackAllowed",
		val,
	)
}

func (j *jsiiProxy_Project)SetCiIdTokenSubClaimComponents(val *[]*string) {
	if err := j.validateSetCiIdTokenSubClaimComponentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ciIdTokenSubClaimComponents",
		val,
	)
}

func (j *jsiiProxy_Project)SetCiPipelineVariablesMinimumOverrideRole(val *string) {
	if err := j.validateSetCiPipelineVariablesMinimumOverrideRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ciPipelineVariablesMinimumOverrideRole",
		val,
	)
}

func (j *jsiiProxy_Project)SetCiRestrictPipelineCancellationRole(val *string) {
	if err := j.validateSetCiRestrictPipelineCancellationRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ciRestrictPipelineCancellationRole",
		val,
	)
}

func (j *jsiiProxy_Project)SetCiSeparatedCaches(val interface{}) {
	if err := j.validateSetCiSeparatedCachesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ciSeparatedCaches",
		val,
	)
}

func (j *jsiiProxy_Project)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Project)SetContainerRegistryAccessLevel(val *string) {
	if err := j.validateSetContainerRegistryAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerRegistryAccessLevel",
		val,
	)
}

func (j *jsiiProxy_Project)SetContainerRegistryEnabled(val interface{}) {
	if err := j.validateSetContainerRegistryEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerRegistryEnabled",
		val,
	)
}

func (j *jsiiProxy_Project)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Project)SetDefaultBranch(val *string) {
	if err := j.validateSetDefaultBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultBranch",
		val,
	)
}

func (j *jsiiProxy_Project)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Project)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Project)SetEmailsEnabled(val interface{}) {
	if err := j.validateSetEmailsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailsEnabled",
		val,
	)
}

func (j *jsiiProxy_Project)SetEnvironmentsAccessLevel(val *string) {
	if err := j.validateSetEnvironmentsAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentsAccessLevel",
		val,
	)
}

func (j *jsiiProxy_Project)SetExternalAuthorizationClassificationLabel(val *string) {
	if err := j.validateSetExternalAuthorizationClassificationLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalAuthorizationClassificationLabel",
		val,
	)
}

func (j *jsiiProxy_Project)SetFeatureFlagsAccessLevel(val *string) {
	if err := j.validateSetFeatureFlagsAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"featureFlagsAccessLevel",
		val,
	)
}

func (j *jsiiProxy_Project)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Project)SetForkedFromProjectId(val *float64) {
	if err := j.validateSetForkedFromProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forkedFromProjectId",
		val,
	)
}

func (j *jsiiProxy_Project)SetForkingAccessLevel(val *string) {
	if err := j.validateSetForkingAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forkingAccessLevel",
		val,
	)
}

func (j *jsiiProxy_Project)SetGroupRunnersEnabled(val interface{}) {
	if err := j.validateSetGroupRunnersEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupRunnersEnabled",
		val,
	)
}

func (j *jsiiProxy_Project)SetGroupWithProjectTemplatesId(val *float64) {
	if err := j.validateSetGroupWithProjectTemplatesIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupWithProjectTemplatesId",
		val,
	)
}

func (j *jsiiProxy_Project)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Project)SetImportUrl(val *string) {
	if err := j.validateSetImportUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importUrl",
		val,
	)
}

func (j *jsiiProxy_Project)SetImportUrlPassword(val *string) {
	if err := j.validateSetImportUrlPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importUrlPassword",
		val,
	)
}

func (j *jsiiProxy_Project)SetImportUrlUsername(val *string) {
	if err := j.validateSetImportUrlUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importUrlUsername",
		val,
	)
}

func (j *jsiiProxy_Project)SetInfrastructureAccessLevel(val *string) {
	if err := j.validateSetInfrastructureAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"infrastructureAccessLevel",
		val,
	)
}

func (j *jsiiProxy_Project)SetInitializeWithReadme(val interface{}) {
	if err := j.validateSetInitializeWithReadmeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initializeWithReadme",
		val,
	)
}

func (j *jsiiProxy_Project)SetIssuesAccessLevel(val *string) {
	if err := j.validateSetIssuesAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuesAccessLevel",
		val,
	)
}

func (j *jsiiProxy_Project)SetIssuesEnabled(val interface{}) {
	if err := j.validateSetIssuesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuesEnabled",
		val,
	)
}

func (j *jsiiProxy_Project)SetIssuesTemplate(val *string) {
	if err := j.validateSetIssuesTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuesTemplate",
		val,
	)
}

func (j *jsiiProxy_Project)SetKeepLatestArtifact(val interface{}) {
	if err := j.validateSetKeepLatestArtifactParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepLatestArtifact",
		val,
	)
}

func (j *jsiiProxy_Project)SetLfsEnabled(val interface{}) {
	if err := j.validateSetLfsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lfsEnabled",
		val,
	)
}

func (j *jsiiProxy_Project)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Project)SetMergeCommitTemplate(val *string) {
	if err := j.validateSetMergeCommitTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeCommitTemplate",
		val,
	)
}

func (j *jsiiProxy_Project)SetMergeMethod(val *string) {
	if err := j.validateSetMergeMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeMethod",
		val,
	)
}

func (j *jsiiProxy_Project)SetMergePipelinesEnabled(val interface{}) {
	if err := j.validateSetMergePipelinesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergePipelinesEnabled",
		val,
	)
}

func (j *jsiiProxy_Project)SetMergeRequestsAccessLevel(val *string) {
	if err := j.validateSetMergeRequestsAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeRequestsAccessLevel",
		val,
	)
}

func (j *jsiiProxy_Project)SetMergeRequestsEnabled(val interface{}) {
	if err := j.validateSetMergeRequestsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeRequestsEnabled",
		val,
	)
}

func (j *jsiiProxy_Project)SetMergeRequestsTemplate(val *string) {
	if err := j.validateSetMergeRequestsTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeRequestsTemplate",
		val,
	)
}

func (j *jsiiProxy_Project)SetMergeTrainsEnabled(val interface{}) {
	if err := j.validateSetMergeTrainsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeTrainsEnabled",
		val,
	)
}

func (j *jsiiProxy_Project)SetMirror(val interface{}) {
	if err := j.validateSetMirrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mirror",
		val,
	)
}

func (j *jsiiProxy_Project)SetMirrorOverwritesDivergedBranches(val interface{}) {
	if err := j.validateSetMirrorOverwritesDivergedBranchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mirrorOverwritesDivergedBranches",
		val,
	)
}

func (j *jsiiProxy_Project)SetMirrorTriggerBuilds(val interface{}) {
	if err := j.validateSetMirrorTriggerBuildsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mirrorTriggerBuilds",
		val,
	)
}

func (j *jsiiProxy_Project)SetModelExperimentsAccessLevel(val *string) {
	if err := j.validateSetModelExperimentsAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelExperimentsAccessLevel",
		val,
	)
}

func (j *jsiiProxy_Project)SetModelRegistryAccessLevel(val *string) {
	if err := j.validateSetModelRegistryAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelRegistryAccessLevel",
		val,
	)
}

func (j *jsiiProxy_Project)SetMonitorAccessLevel(val *string) {
	if err := j.validateSetMonitorAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitorAccessLevel",
		val,
	)
}

func (j *jsiiProxy_Project)SetMrDefaultTargetSelf(val interface{}) {
	if err := j.validateSetMrDefaultTargetSelfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mrDefaultTargetSelf",
		val,
	)
}

func (j *jsiiProxy_Project)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Project)SetNamespaceId(val *float64) {
	if err := j.validateSetNamespaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespaceId",
		val,
	)
}

func (j *jsiiProxy_Project)SetOnlyAllowMergeIfAllDiscussionsAreResolved(val interface{}) {
	if err := j.validateSetOnlyAllowMergeIfAllDiscussionsAreResolvedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onlyAllowMergeIfAllDiscussionsAreResolved",
		val,
	)
}

func (j *jsiiProxy_Project)SetOnlyAllowMergeIfPipelineSucceeds(val interface{}) {
	if err := j.validateSetOnlyAllowMergeIfPipelineSucceedsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onlyAllowMergeIfPipelineSucceeds",
		val,
	)
}

func (j *jsiiProxy_Project)SetOnlyMirrorProtectedBranches(val interface{}) {
	if err := j.validateSetOnlyMirrorProtectedBranchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onlyMirrorProtectedBranches",
		val,
	)
}

func (j *jsiiProxy_Project)SetPackagesEnabled(val interface{}) {
	if err := j.validateSetPackagesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packagesEnabled",
		val,
	)
}

func (j *jsiiProxy_Project)SetPagesAccessLevel(val *string) {
	if err := j.validateSetPagesAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pagesAccessLevel",
		val,
	)
}

func (j *jsiiProxy_Project)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_Project)SetPermanentlyDeleteOnDestroy(val interface{}) {
	if err := j.validateSetPermanentlyDeleteOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permanentlyDeleteOnDestroy",
		val,
	)
}

func (j *jsiiProxy_Project)SetPipelinesEnabled(val interface{}) {
	if err := j.validateSetPipelinesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipelinesEnabled",
		val,
	)
}

func (j *jsiiProxy_Project)SetPreReceiveSecretDetectionEnabled(val interface{}) {
	if err := j.validateSetPreReceiveSecretDetectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preReceiveSecretDetectionEnabled",
		val,
	)
}

func (j *jsiiProxy_Project)SetPreventMergeWithoutJiraIssue(val interface{}) {
	if err := j.validateSetPreventMergeWithoutJiraIssueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preventMergeWithoutJiraIssue",
		val,
	)
}

func (j *jsiiProxy_Project)SetPrintingMergeRequestLinkEnabled(val interface{}) {
	if err := j.validateSetPrintingMergeRequestLinkEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"printingMergeRequestLinkEnabled",
		val,
	)
}

func (j *jsiiProxy_Project)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Project)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Project)SetPublicBuilds(val interface{}) {
	if err := j.validateSetPublicBuildsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicBuilds",
		val,
	)
}

func (j *jsiiProxy_Project)SetPublicJobs(val interface{}) {
	if err := j.validateSetPublicJobsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicJobs",
		val,
	)
}

func (j *jsiiProxy_Project)SetReleasesAccessLevel(val *string) {
	if err := j.validateSetReleasesAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releasesAccessLevel",
		val,
	)
}

func (j *jsiiProxy_Project)SetRemoveSourceBranchAfterMerge(val interface{}) {
	if err := j.validateSetRemoveSourceBranchAfterMergeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"removeSourceBranchAfterMerge",
		val,
	)
}

func (j *jsiiProxy_Project)SetRepositoryAccessLevel(val *string) {
	if err := j.validateSetRepositoryAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repositoryAccessLevel",
		val,
	)
}

func (j *jsiiProxy_Project)SetRepositoryStorage(val *string) {
	if err := j.validateSetRepositoryStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repositoryStorage",
		val,
	)
}

func (j *jsiiProxy_Project)SetRequestAccessEnabled(val interface{}) {
	if err := j.validateSetRequestAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_Project)SetRequirementsAccessLevel(val *string) {
	if err := j.validateSetRequirementsAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requirementsAccessLevel",
		val,
	)
}

func (j *jsiiProxy_Project)SetResolveOutdatedDiffDiscussions(val interface{}) {
	if err := j.validateSetResolveOutdatedDiffDiscussionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resolveOutdatedDiffDiscussions",
		val,
	)
}

func (j *jsiiProxy_Project)SetRestrictUserDefinedVariables(val interface{}) {
	if err := j.validateSetRestrictUserDefinedVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictUserDefinedVariables",
		val,
	)
}

func (j *jsiiProxy_Project)SetSecurityAndComplianceAccessLevel(val *string) {
	if err := j.validateSetSecurityAndComplianceAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityAndComplianceAccessLevel",
		val,
	)
}

func (j *jsiiProxy_Project)SetSharedRunnersEnabled(val interface{}) {
	if err := j.validateSetSharedRunnersEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedRunnersEnabled",
		val,
	)
}

func (j *jsiiProxy_Project)SetSkipWaitForDefaultBranchProtection(val interface{}) {
	if err := j.validateSetSkipWaitForDefaultBranchProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipWaitForDefaultBranchProtection",
		val,
	)
}

func (j *jsiiProxy_Project)SetSnippetsAccessLevel(val *string) {
	if err := j.validateSetSnippetsAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snippetsAccessLevel",
		val,
	)
}

func (j *jsiiProxy_Project)SetSnippetsEnabled(val interface{}) {
	if err := j.validateSetSnippetsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snippetsEnabled",
		val,
	)
}

func (j *jsiiProxy_Project)SetSquashCommitTemplate(val *string) {
	if err := j.validateSetSquashCommitTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"squashCommitTemplate",
		val,
	)
}

func (j *jsiiProxy_Project)SetSquashOption(val *string) {
	if err := j.validateSetSquashOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"squashOption",
		val,
	)
}

func (j *jsiiProxy_Project)SetSuggestionCommitMessage(val *string) {
	if err := j.validateSetSuggestionCommitMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suggestionCommitMessage",
		val,
	)
}

func (j *jsiiProxy_Project)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Project)SetTemplateName(val *string) {
	if err := j.validateSetTemplateNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateName",
		val,
	)
}

func (j *jsiiProxy_Project)SetTemplateProjectId(val *float64) {
	if err := j.validateSetTemplateProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateProjectId",
		val,
	)
}

func (j *jsiiProxy_Project)SetTopics(val *[]*string) {
	if err := j.validateSetTopicsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topics",
		val,
	)
}

func (j *jsiiProxy_Project)SetUseCustomTemplate(val interface{}) {
	if err := j.validateSetUseCustomTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCustomTemplate",
		val,
	)
}

func (j *jsiiProxy_Project)SetVisibilityLevel(val *string) {
	if err := j.validateSetVisibilityLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visibilityLevel",
		val,
	)
}

func (j *jsiiProxy_Project)SetWikiAccessLevel(val *string) {
	if err := j.validateSetWikiAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wikiAccessLevel",
		val,
	)
}

func (j *jsiiProxy_Project)SetWikiEnabled(val interface{}) {
	if err := j.validateSetWikiEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wikiEnabled",
		val,
	)
}

// Generates CDKTF code for importing a Project resource upon running "cdktf plan <stack-name>".
func Project_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateProject_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.project.Project",
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
func Project_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProject_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.project.Project",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Project_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProject_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.project.Project",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Project_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProject_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.project.Project",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Project_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.project.Project",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_Project) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_Project) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_Project) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_Project) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_Project) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_Project) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_Project) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_Project) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_Project) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_Project) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_Project) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_Project) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_Project) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_Project) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_Project) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_Project) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_Project) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_Project) PutContainerExpirationPolicy(value *ProjectContainerExpirationPolicy) {
	if err := p.validatePutContainerExpirationPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putContainerExpirationPolicy",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_Project) PutPushRules(value *ProjectPushRules) {
	if err := p.validatePutPushRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPushRules",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_Project) PutTimeouts(value *ProjectTimeouts) {
	if err := p.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_Project) ResetAllowMergeOnSkippedPipeline() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowMergeOnSkippedPipeline",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetAllowPipelineTriggerApproveDeployment() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowPipelineTriggerApproveDeployment",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetAnalyticsAccessLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetAnalyticsAccessLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetApprovalsBeforeMerge() {
	_jsii_.InvokeVoid(
		p,
		"resetApprovalsBeforeMerge",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetArchived() {
	_jsii_.InvokeVoid(
		p,
		"resetArchived",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetArchiveOnDestroy() {
	_jsii_.InvokeVoid(
		p,
		"resetArchiveOnDestroy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetAutoCancelPendingPipelines() {
	_jsii_.InvokeVoid(
		p,
		"resetAutoCancelPendingPipelines",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetAutocloseReferencedIssues() {
	_jsii_.InvokeVoid(
		p,
		"resetAutocloseReferencedIssues",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetAutoDevopsDeployStrategy() {
	_jsii_.InvokeVoid(
		p,
		"resetAutoDevopsDeployStrategy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetAutoDevopsEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetAutoDevopsEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetAutoDuoCodeReviewEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetAutoDuoCodeReviewEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetAvatar() {
	_jsii_.InvokeVoid(
		p,
		"resetAvatar",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetAvatarHash() {
	_jsii_.InvokeVoid(
		p,
		"resetAvatarHash",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetBranches() {
	_jsii_.InvokeVoid(
		p,
		"resetBranches",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetBuildGitStrategy() {
	_jsii_.InvokeVoid(
		p,
		"resetBuildGitStrategy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetBuildsAccessLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetBuildsAccessLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetBuildTimeout() {
	_jsii_.InvokeVoid(
		p,
		"resetBuildTimeout",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetCiConfigPath() {
	_jsii_.InvokeVoid(
		p,
		"resetCiConfigPath",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetCiDefaultGitDepth() {
	_jsii_.InvokeVoid(
		p,
		"resetCiDefaultGitDepth",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetCiDeletePipelinesInSeconds() {
	_jsii_.InvokeVoid(
		p,
		"resetCiDeletePipelinesInSeconds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetCiForwardDeploymentEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetCiForwardDeploymentEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetCiForwardDeploymentRollbackAllowed() {
	_jsii_.InvokeVoid(
		p,
		"resetCiForwardDeploymentRollbackAllowed",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetCiIdTokenSubClaimComponents() {
	_jsii_.InvokeVoid(
		p,
		"resetCiIdTokenSubClaimComponents",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetCiPipelineVariablesMinimumOverrideRole() {
	_jsii_.InvokeVoid(
		p,
		"resetCiPipelineVariablesMinimumOverrideRole",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetCiRestrictPipelineCancellationRole() {
	_jsii_.InvokeVoid(
		p,
		"resetCiRestrictPipelineCancellationRole",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetCiSeparatedCaches() {
	_jsii_.InvokeVoid(
		p,
		"resetCiSeparatedCaches",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetContainerExpirationPolicy() {
	_jsii_.InvokeVoid(
		p,
		"resetContainerExpirationPolicy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetContainerRegistryAccessLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetContainerRegistryAccessLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetContainerRegistryEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetContainerRegistryEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetDefaultBranch() {
	_jsii_.InvokeVoid(
		p,
		"resetDefaultBranch",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetDescription() {
	_jsii_.InvokeVoid(
		p,
		"resetDescription",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetEmailsEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetEmailsEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetEnvironmentsAccessLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetEnvironmentsAccessLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetExternalAuthorizationClassificationLabel() {
	_jsii_.InvokeVoid(
		p,
		"resetExternalAuthorizationClassificationLabel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetFeatureFlagsAccessLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetFeatureFlagsAccessLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetForkedFromProjectId() {
	_jsii_.InvokeVoid(
		p,
		"resetForkedFromProjectId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetForkingAccessLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetForkingAccessLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetGroupRunnersEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetGroupRunnersEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetGroupWithProjectTemplatesId() {
	_jsii_.InvokeVoid(
		p,
		"resetGroupWithProjectTemplatesId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetImportUrl() {
	_jsii_.InvokeVoid(
		p,
		"resetImportUrl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetImportUrlPassword() {
	_jsii_.InvokeVoid(
		p,
		"resetImportUrlPassword",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetImportUrlUsername() {
	_jsii_.InvokeVoid(
		p,
		"resetImportUrlUsername",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetInfrastructureAccessLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetInfrastructureAccessLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetInitializeWithReadme() {
	_jsii_.InvokeVoid(
		p,
		"resetInitializeWithReadme",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetIssuesAccessLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetIssuesAccessLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetIssuesEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetIssuesEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetIssuesTemplate() {
	_jsii_.InvokeVoid(
		p,
		"resetIssuesTemplate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetKeepLatestArtifact() {
	_jsii_.InvokeVoid(
		p,
		"resetKeepLatestArtifact",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetLfsEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetLfsEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetMergeCommitTemplate() {
	_jsii_.InvokeVoid(
		p,
		"resetMergeCommitTemplate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetMergeMethod() {
	_jsii_.InvokeVoid(
		p,
		"resetMergeMethod",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetMergePipelinesEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetMergePipelinesEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetMergeRequestsAccessLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetMergeRequestsAccessLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetMergeRequestsEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetMergeRequestsEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetMergeRequestsTemplate() {
	_jsii_.InvokeVoid(
		p,
		"resetMergeRequestsTemplate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetMergeTrainsEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetMergeTrainsEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetMirror() {
	_jsii_.InvokeVoid(
		p,
		"resetMirror",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetMirrorOverwritesDivergedBranches() {
	_jsii_.InvokeVoid(
		p,
		"resetMirrorOverwritesDivergedBranches",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetMirrorTriggerBuilds() {
	_jsii_.InvokeVoid(
		p,
		"resetMirrorTriggerBuilds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetModelExperimentsAccessLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetModelExperimentsAccessLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetModelRegistryAccessLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetModelRegistryAccessLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetMonitorAccessLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetMonitorAccessLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetMrDefaultTargetSelf() {
	_jsii_.InvokeVoid(
		p,
		"resetMrDefaultTargetSelf",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetNamespaceId() {
	_jsii_.InvokeVoid(
		p,
		"resetNamespaceId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetOnlyAllowMergeIfAllDiscussionsAreResolved() {
	_jsii_.InvokeVoid(
		p,
		"resetOnlyAllowMergeIfAllDiscussionsAreResolved",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetOnlyAllowMergeIfPipelineSucceeds() {
	_jsii_.InvokeVoid(
		p,
		"resetOnlyAllowMergeIfPipelineSucceeds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetOnlyMirrorProtectedBranches() {
	_jsii_.InvokeVoid(
		p,
		"resetOnlyMirrorProtectedBranches",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetPackagesEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetPackagesEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetPagesAccessLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetPagesAccessLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetPath() {
	_jsii_.InvokeVoid(
		p,
		"resetPath",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetPermanentlyDeleteOnDestroy() {
	_jsii_.InvokeVoid(
		p,
		"resetPermanentlyDeleteOnDestroy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetPipelinesEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetPipelinesEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetPreReceiveSecretDetectionEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetPreReceiveSecretDetectionEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetPreventMergeWithoutJiraIssue() {
	_jsii_.InvokeVoid(
		p,
		"resetPreventMergeWithoutJiraIssue",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetPrintingMergeRequestLinkEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetPrintingMergeRequestLinkEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetPublicBuilds() {
	_jsii_.InvokeVoid(
		p,
		"resetPublicBuilds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetPublicJobs() {
	_jsii_.InvokeVoid(
		p,
		"resetPublicJobs",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetPushRules() {
	_jsii_.InvokeVoid(
		p,
		"resetPushRules",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetReleasesAccessLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetReleasesAccessLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetRemoveSourceBranchAfterMerge() {
	_jsii_.InvokeVoid(
		p,
		"resetRemoveSourceBranchAfterMerge",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetRepositoryAccessLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetRepositoryAccessLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetRepositoryStorage() {
	_jsii_.InvokeVoid(
		p,
		"resetRepositoryStorage",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetRequestAccessEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetRequestAccessEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetRequirementsAccessLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetRequirementsAccessLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetResolveOutdatedDiffDiscussions() {
	_jsii_.InvokeVoid(
		p,
		"resetResolveOutdatedDiffDiscussions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetRestrictUserDefinedVariables() {
	_jsii_.InvokeVoid(
		p,
		"resetRestrictUserDefinedVariables",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetSecurityAndComplianceAccessLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetSecurityAndComplianceAccessLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetSharedRunnersEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetSharedRunnersEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetSkipWaitForDefaultBranchProtection() {
	_jsii_.InvokeVoid(
		p,
		"resetSkipWaitForDefaultBranchProtection",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetSnippetsAccessLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetSnippetsAccessLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetSnippetsEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetSnippetsEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetSquashCommitTemplate() {
	_jsii_.InvokeVoid(
		p,
		"resetSquashCommitTemplate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetSquashOption() {
	_jsii_.InvokeVoid(
		p,
		"resetSquashOption",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetSuggestionCommitMessage() {
	_jsii_.InvokeVoid(
		p,
		"resetSuggestionCommitMessage",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetTags() {
	_jsii_.InvokeVoid(
		p,
		"resetTags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetTemplateName() {
	_jsii_.InvokeVoid(
		p,
		"resetTemplateName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetTemplateProjectId() {
	_jsii_.InvokeVoid(
		p,
		"resetTemplateProjectId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetTimeouts() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetTopics() {
	_jsii_.InvokeVoid(
		p,
		"resetTopics",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetUseCustomTemplate() {
	_jsii_.InvokeVoid(
		p,
		"resetUseCustomTemplate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetVisibilityLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetVisibilityLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetWikiAccessLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetWikiAccessLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetWikiEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetWikiEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

