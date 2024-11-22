// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabproject

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProject.DataGitlabProject",
		reflect.TypeOf((*DataGitlabProject)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowPipelineTriggerApproveDeployment", GoGetter: "AllowPipelineTriggerApproveDeployment"},
			_jsii_.MemberProperty{JsiiProperty: "analyticsAccessLevel", GoGetter: "AnalyticsAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "archived", GoGetter: "Archived"},
			_jsii_.MemberProperty{JsiiProperty: "autoCancelPendingPipelines", GoGetter: "AutoCancelPendingPipelines"},
			_jsii_.MemberProperty{JsiiProperty: "autocloseReferencedIssues", GoGetter: "AutocloseReferencedIssues"},
			_jsii_.MemberProperty{JsiiProperty: "autoDevopsDeployStrategy", GoGetter: "AutoDevopsDeployStrategy"},
			_jsii_.MemberProperty{JsiiProperty: "autoDevopsEnabled", GoGetter: "AutoDevopsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "buildGitStrategy", GoGetter: "BuildGitStrategy"},
			_jsii_.MemberProperty{JsiiProperty: "buildsAccessLevel", GoGetter: "BuildsAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "buildTimeout", GoGetter: "BuildTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "ciConfigPath", GoGetter: "CiConfigPath"},
			_jsii_.MemberProperty{JsiiProperty: "ciDefaultGitDepth", GoGetter: "CiDefaultGitDepth"},
			_jsii_.MemberProperty{JsiiProperty: "ciDefaultGitDepthInput", GoGetter: "CiDefaultGitDepthInput"},
			_jsii_.MemberProperty{JsiiProperty: "ciPipelineVariablesMinimumOverrideRole", GoGetter: "CiPipelineVariablesMinimumOverrideRole"},
			_jsii_.MemberProperty{JsiiProperty: "ciRestrictPipelineCancellationRole", GoGetter: "CiRestrictPipelineCancellationRole"},
			_jsii_.MemberProperty{JsiiProperty: "ciSeparatedCaches", GoGetter: "CiSeparatedCaches"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "containerExpirationPolicy", GoGetter: "ContainerExpirationPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "containerRegistryAccessLevel", GoGetter: "ContainerRegistryAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultBranch", GoGetter: "DefaultBranch"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "emailsEnabled", GoGetter: "EmailsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "emptyRepo", GoGetter: "EmptyRepo"},
			_jsii_.MemberProperty{JsiiProperty: "environmentsAccessLevel", GoGetter: "EnvironmentsAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "externalAuthorizationClassificationLabel", GoGetter: "ExternalAuthorizationClassificationLabel"},
			_jsii_.MemberProperty{JsiiProperty: "featureFlagsAccessLevel", GoGetter: "FeatureFlagsAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "forkingAccessLevel", GoGetter: "ForkingAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "httpUrlToRepo", GoGetter: "HttpUrlToRepo"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "importUrl", GoGetter: "ImportUrl"},
			_jsii_.MemberProperty{JsiiProperty: "infrastructureAccessLevel", GoGetter: "InfrastructureAccessLevel"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "issuesAccessLevel", GoGetter: "IssuesAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "issuesEnabled", GoGetter: "IssuesEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "keepLatestArtifact", GoGetter: "KeepLatestArtifact"},
			_jsii_.MemberProperty{JsiiProperty: "lfsEnabled", GoGetter: "LfsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "mergeCommitTemplate", GoGetter: "MergeCommitTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "mergePipelinesEnabled", GoGetter: "MergePipelinesEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "mergeRequestsAccessLevel", GoGetter: "MergeRequestsAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "mergeRequestsEnabled", GoGetter: "MergeRequestsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "mergeTrainsEnabled", GoGetter: "MergeTrainsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "modelExperimentsAccessLevel", GoGetter: "ModelExperimentsAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "modelRegistryAccessLevel", GoGetter: "ModelRegistryAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "monitorAccessLevel", GoGetter: "MonitorAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceId", GoGetter: "NamespaceId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "pathWithNamespace", GoGetter: "PathWithNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "pathWithNamespaceInput", GoGetter: "PathWithNamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "pipelinesEnabled", GoGetter: "PipelinesEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "printingMergeRequestLinkEnabled", GoGetter: "PrintingMergeRequestLinkEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "publicBuilds", GoGetter: "PublicBuilds"},
			_jsii_.MemberProperty{JsiiProperty: "publicBuildsInput", GoGetter: "PublicBuildsInput"},
			_jsii_.MemberProperty{JsiiProperty: "pushRules", GoGetter: "PushRules"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "releasesAccessLevel", GoGetter: "ReleasesAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "removeSourceBranchAfterMerge", GoGetter: "RemoveSourceBranchAfterMerge"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryAccessLevel", GoGetter: "RepositoryAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryStorage", GoGetter: "RepositoryStorage"},
			_jsii_.MemberProperty{JsiiProperty: "requestAccessEnabled", GoGetter: "RequestAccessEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "requirementsAccessLevel", GoGetter: "RequirementsAccessLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetCiDefaultGitDepth", GoMethod: "ResetCiDefaultGitDepth"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPathWithNamespace", GoMethod: "ResetPathWithNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetPublicBuilds", GoMethod: "ResetPublicBuilds"},
			_jsii_.MemberProperty{JsiiProperty: "resolveOutdatedDiffDiscussions", GoGetter: "ResolveOutdatedDiffDiscussions"},
			_jsii_.MemberProperty{JsiiProperty: "restrictUserDefinedVariables", GoGetter: "RestrictUserDefinedVariables"},
			_jsii_.MemberProperty{JsiiProperty: "runnersToken", GoGetter: "RunnersToken"},
			_jsii_.MemberProperty{JsiiProperty: "securityAndComplianceAccessLevel", GoGetter: "SecurityAndComplianceAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "sharedWithGroups", GoGetter: "SharedWithGroups"},
			_jsii_.MemberProperty{JsiiProperty: "snippetsAccessLevel", GoGetter: "SnippetsAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "snippetsEnabled", GoGetter: "SnippetsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "squashCommitTemplate", GoGetter: "SquashCommitTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "sshUrlToRepo", GoGetter: "SshUrlToRepo"},
			_jsii_.MemberProperty{JsiiProperty: "suggestionCommitMessage", GoGetter: "SuggestionCommitMessage"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "topics", GoGetter: "Topics"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "visibilityLevel", GoGetter: "VisibilityLevel"},
			_jsii_.MemberProperty{JsiiProperty: "webUrl", GoGetter: "WebUrl"},
			_jsii_.MemberProperty{JsiiProperty: "wikiAccessLevel", GoGetter: "WikiAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "wikiEnabled", GoGetter: "WikiEnabled"},
		},
		func() interface{} {
			j := jsiiProxy_DataGitlabProject{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.dataGitlabProject.DataGitlabProjectConfig",
		reflect.TypeOf((*DataGitlabProjectConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.dataGitlabProject.DataGitlabProjectContainerExpirationPolicy",
		reflect.TypeOf((*DataGitlabProjectContainerExpirationPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProject.DataGitlabProjectContainerExpirationPolicyList",
		reflect.TypeOf((*DataGitlabProjectContainerExpirationPolicyList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_DataGitlabProjectContainerExpirationPolicyList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProject.DataGitlabProjectContainerExpirationPolicyOutputReference",
		reflect.TypeOf((*DataGitlabProjectContainerExpirationPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cadence", GoGetter: "Cadence"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keepN", GoGetter: "KeepN"},
			_jsii_.MemberProperty{JsiiProperty: "nameRegex", GoGetter: "NameRegex"},
			_jsii_.MemberProperty{JsiiProperty: "nameRegexDelete", GoGetter: "NameRegexDelete"},
			_jsii_.MemberProperty{JsiiProperty: "nameRegexKeep", GoGetter: "NameRegexKeep"},
			_jsii_.MemberProperty{JsiiProperty: "nextRunAt", GoGetter: "NextRunAt"},
			_jsii_.MemberProperty{JsiiProperty: "olderThan", GoGetter: "OlderThan"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataGitlabProjectContainerExpirationPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.dataGitlabProject.DataGitlabProjectPushRules",
		reflect.TypeOf((*DataGitlabProjectPushRules)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProject.DataGitlabProjectPushRulesList",
		reflect.TypeOf((*DataGitlabProjectPushRulesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_DataGitlabProjectPushRulesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProject.DataGitlabProjectPushRulesOutputReference",
		reflect.TypeOf((*DataGitlabProjectPushRulesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authorEmailRegex", GoGetter: "AuthorEmailRegex"},
			_jsii_.MemberProperty{JsiiProperty: "branchNameRegex", GoGetter: "BranchNameRegex"},
			_jsii_.MemberProperty{JsiiProperty: "commitCommitterCheck", GoGetter: "CommitCommitterCheck"},
			_jsii_.MemberProperty{JsiiProperty: "commitCommitterNameCheck", GoGetter: "CommitCommitterNameCheck"},
			_jsii_.MemberProperty{JsiiProperty: "commitMessageNegativeRegex", GoGetter: "CommitMessageNegativeRegex"},
			_jsii_.MemberProperty{JsiiProperty: "commitMessageRegex", GoGetter: "CommitMessageRegex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "denyDeleteTag", GoGetter: "DenyDeleteTag"},
			_jsii_.MemberProperty{JsiiProperty: "fileNameRegex", GoGetter: "FileNameRegex"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maxFileSize", GoGetter: "MaxFileSize"},
			_jsii_.MemberProperty{JsiiProperty: "memberCheck", GoGetter: "MemberCheck"},
			_jsii_.MemberProperty{JsiiProperty: "preventSecrets", GoGetter: "PreventSecrets"},
			_jsii_.MemberProperty{JsiiProperty: "rejectNonDcoCommits", GoGetter: "RejectNonDcoCommits"},
			_jsii_.MemberProperty{JsiiProperty: "rejectUnsignedCommits", GoGetter: "RejectUnsignedCommits"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataGitlabProjectPushRulesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.dataGitlabProject.DataGitlabProjectSharedWithGroups",
		reflect.TypeOf((*DataGitlabProjectSharedWithGroups)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProject.DataGitlabProjectSharedWithGroupsList",
		reflect.TypeOf((*DataGitlabProjectSharedWithGroupsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_DataGitlabProjectSharedWithGroupsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProject.DataGitlabProjectSharedWithGroupsOutputReference",
		reflect.TypeOf((*DataGitlabProjectSharedWithGroupsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "groupAccessLevel", GoGetter: "GroupAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "groupFullPath", GoGetter: "GroupFullPath"},
			_jsii_.MemberProperty{JsiiProperty: "groupId", GoGetter: "GroupId"},
			_jsii_.MemberProperty{JsiiProperty: "groupName", GoGetter: "GroupName"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataGitlabProjectSharedWithGroupsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
