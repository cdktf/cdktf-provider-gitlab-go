package datagitlabprojects

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProjects.DataGitlabProjects",
		reflect.TypeOf((*DataGitlabProjects)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "archived", GoGetter: "Archived"},
			_jsii_.MemberProperty{JsiiProperty: "archivedInput", GoGetter: "ArchivedInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
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
			_jsii_.MemberProperty{JsiiProperty: "groupId", GoGetter: "GroupId"},
			_jsii_.MemberProperty{JsiiProperty: "groupIdInput", GoGetter: "GroupIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "includeSubgroups", GoGetter: "IncludeSubgroups"},
			_jsii_.MemberProperty{JsiiProperty: "includeSubgroupsInput", GoGetter: "IncludeSubgroupsInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "maxQueryablePages", GoGetter: "MaxQueryablePages"},
			_jsii_.MemberProperty{JsiiProperty: "maxQueryablePagesInput", GoGetter: "MaxQueryablePagesInput"},
			_jsii_.MemberProperty{JsiiProperty: "membership", GoGetter: "Membership"},
			_jsii_.MemberProperty{JsiiProperty: "membershipInput", GoGetter: "MembershipInput"},
			_jsii_.MemberProperty{JsiiProperty: "minAccessLevel", GoGetter: "MinAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "minAccessLevelInput", GoGetter: "MinAccessLevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "orderBy", GoGetter: "OrderBy"},
			_jsii_.MemberProperty{JsiiProperty: "orderByInput", GoGetter: "OrderByInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "owned", GoGetter: "Owned"},
			_jsii_.MemberProperty{JsiiProperty: "ownedInput", GoGetter: "OwnedInput"},
			_jsii_.MemberProperty{JsiiProperty: "page", GoGetter: "Page"},
			_jsii_.MemberProperty{JsiiProperty: "pageInput", GoGetter: "PageInput"},
			_jsii_.MemberProperty{JsiiProperty: "perPage", GoGetter: "PerPage"},
			_jsii_.MemberProperty{JsiiProperty: "perPageInput", GoGetter: "PerPageInput"},
			_jsii_.MemberProperty{JsiiProperty: "projects", GoGetter: "Projects"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetArchived", GoMethod: "ResetArchived"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroupId", GoMethod: "ResetGroupId"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeSubgroups", GoMethod: "ResetIncludeSubgroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxQueryablePages", GoMethod: "ResetMaxQueryablePages"},
			_jsii_.MemberMethod{JsiiMethod: "resetMembership", GoMethod: "ResetMembership"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinAccessLevel", GoMethod: "ResetMinAccessLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrderBy", GoMethod: "ResetOrderBy"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOwned", GoMethod: "ResetOwned"},
			_jsii_.MemberMethod{JsiiMethod: "resetPage", GoMethod: "ResetPage"},
			_jsii_.MemberMethod{JsiiMethod: "resetPerPage", GoMethod: "ResetPerPage"},
			_jsii_.MemberMethod{JsiiMethod: "resetSearch", GoMethod: "ResetSearch"},
			_jsii_.MemberMethod{JsiiMethod: "resetSimple", GoMethod: "ResetSimple"},
			_jsii_.MemberMethod{JsiiMethod: "resetSort", GoMethod: "ResetSort"},
			_jsii_.MemberMethod{JsiiMethod: "resetStarred", GoMethod: "ResetStarred"},
			_jsii_.MemberMethod{JsiiMethod: "resetStatistics", GoMethod: "ResetStatistics"},
			_jsii_.MemberMethod{JsiiMethod: "resetVisibility", GoMethod: "ResetVisibility"},
			_jsii_.MemberMethod{JsiiMethod: "resetWithCustomAttributes", GoMethod: "ResetWithCustomAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "resetWithIssuesEnabled", GoMethod: "ResetWithIssuesEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetWithMergeRequestsEnabled", GoMethod: "ResetWithMergeRequestsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetWithProgrammingLanguage", GoMethod: "ResetWithProgrammingLanguage"},
			_jsii_.MemberMethod{JsiiMethod: "resetWithShared", GoMethod: "ResetWithShared"},
			_jsii_.MemberProperty{JsiiProperty: "search", GoGetter: "Search"},
			_jsii_.MemberProperty{JsiiProperty: "searchInput", GoGetter: "SearchInput"},
			_jsii_.MemberProperty{JsiiProperty: "simple", GoGetter: "Simple"},
			_jsii_.MemberProperty{JsiiProperty: "simpleInput", GoGetter: "SimpleInput"},
			_jsii_.MemberProperty{JsiiProperty: "sort", GoGetter: "Sort"},
			_jsii_.MemberProperty{JsiiProperty: "sortInput", GoGetter: "SortInput"},
			_jsii_.MemberProperty{JsiiProperty: "starred", GoGetter: "Starred"},
			_jsii_.MemberProperty{JsiiProperty: "starredInput", GoGetter: "StarredInput"},
			_jsii_.MemberProperty{JsiiProperty: "statistics", GoGetter: "Statistics"},
			_jsii_.MemberProperty{JsiiProperty: "statisticsInput", GoGetter: "StatisticsInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "visibility", GoGetter: "Visibility"},
			_jsii_.MemberProperty{JsiiProperty: "visibilityInput", GoGetter: "VisibilityInput"},
			_jsii_.MemberProperty{JsiiProperty: "withCustomAttributes", GoGetter: "WithCustomAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "withCustomAttributesInput", GoGetter: "WithCustomAttributesInput"},
			_jsii_.MemberProperty{JsiiProperty: "withIssuesEnabled", GoGetter: "WithIssuesEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "withIssuesEnabledInput", GoGetter: "WithIssuesEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "withMergeRequestsEnabled", GoGetter: "WithMergeRequestsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "withMergeRequestsEnabledInput", GoGetter: "WithMergeRequestsEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "withProgrammingLanguage", GoGetter: "WithProgrammingLanguage"},
			_jsii_.MemberProperty{JsiiProperty: "withProgrammingLanguageInput", GoGetter: "WithProgrammingLanguageInput"},
			_jsii_.MemberProperty{JsiiProperty: "withShared", GoGetter: "WithShared"},
			_jsii_.MemberProperty{JsiiProperty: "withSharedInput", GoGetter: "WithSharedInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataGitlabProjects{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.dataGitlabProjects.DataGitlabProjectsConfig",
		reflect.TypeOf((*DataGitlabProjectsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.dataGitlabProjects.DataGitlabProjectsProjects",
		reflect.TypeOf((*DataGitlabProjectsProjects)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.dataGitlabProjects.DataGitlabProjectsProjectsContainerExpirationPolicy",
		reflect.TypeOf((*DataGitlabProjectsProjectsContainerExpirationPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProjects.DataGitlabProjectsProjectsContainerExpirationPolicyList",
		reflect.TypeOf((*DataGitlabProjectsProjectsContainerExpirationPolicyList)(nil)).Elem(),
		[]_jsii_.Member{
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
			j := jsiiProxy_DataGitlabProjectsProjectsContainerExpirationPolicyList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProjects.DataGitlabProjectsProjectsContainerExpirationPolicyOutputReference",
		reflect.TypeOf((*DataGitlabProjectsProjectsContainerExpirationPolicyOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_DataGitlabProjectsProjectsContainerExpirationPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.dataGitlabProjects.DataGitlabProjectsProjectsForkedFromProject",
		reflect.TypeOf((*DataGitlabProjectsProjectsForkedFromProject)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProjects.DataGitlabProjectsProjectsForkedFromProjectList",
		reflect.TypeOf((*DataGitlabProjectsProjectsForkedFromProjectList)(nil)).Elem(),
		[]_jsii_.Member{
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
			j := jsiiProxy_DataGitlabProjectsProjectsForkedFromProjectList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProjects.DataGitlabProjectsProjectsForkedFromProjectOutputReference",
		reflect.TypeOf((*DataGitlabProjectsProjectsForkedFromProjectOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "httpUrlToRepo", GoGetter: "HttpUrlToRepo"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameWithNamespace", GoGetter: "NameWithNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "pathWithNamespace", GoGetter: "PathWithNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "webUrl", GoGetter: "WebUrl"},
		},
		func() interface{} {
			j := jsiiProxy_DataGitlabProjectsProjectsForkedFromProjectOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProjects.DataGitlabProjectsProjectsList",
		reflect.TypeOf((*DataGitlabProjectsProjectsList)(nil)).Elem(),
		[]_jsii_.Member{
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
			j := jsiiProxy_DataGitlabProjectsProjectsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.dataGitlabProjects.DataGitlabProjectsProjectsNamespace",
		reflect.TypeOf((*DataGitlabProjectsProjectsNamespace)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProjects.DataGitlabProjectsProjectsNamespaceList",
		reflect.TypeOf((*DataGitlabProjectsProjectsNamespaceList)(nil)).Elem(),
		[]_jsii_.Member{
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
			j := jsiiProxy_DataGitlabProjectsProjectsNamespaceList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProjects.DataGitlabProjectsProjectsNamespaceOutputReference",
		reflect.TypeOf((*DataGitlabProjectsProjectsNamespaceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "fullPath", GoGetter: "FullPath"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataGitlabProjectsProjectsNamespaceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProjects.DataGitlabProjectsProjectsOutputReference",
		reflect.TypeOf((*DataGitlabProjectsProjectsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowMergeOnSkippedPipeline", GoGetter: "AllowMergeOnSkippedPipeline"},
			_jsii_.MemberProperty{JsiiProperty: "analyticsAccessLevel", GoGetter: "AnalyticsAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "approvalsBeforeMerge", GoGetter: "ApprovalsBeforeMerge"},
			_jsii_.MemberProperty{JsiiProperty: "archived", GoGetter: "Archived"},
			_jsii_.MemberProperty{JsiiProperty: "autoCancelPendingPipelines", GoGetter: "AutoCancelPendingPipelines"},
			_jsii_.MemberProperty{JsiiProperty: "autocloseReferencedIssues", GoGetter: "AutocloseReferencedIssues"},
			_jsii_.MemberProperty{JsiiProperty: "autoDevopsDeployStrategy", GoGetter: "AutoDevopsDeployStrategy"},
			_jsii_.MemberProperty{JsiiProperty: "autoDevopsEnabled", GoGetter: "AutoDevopsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "avatarUrl", GoGetter: "AvatarUrl"},
			_jsii_.MemberProperty{JsiiProperty: "buildCoverageRegex", GoGetter: "BuildCoverageRegex"},
			_jsii_.MemberProperty{JsiiProperty: "buildGitStrategy", GoGetter: "BuildGitStrategy"},
			_jsii_.MemberProperty{JsiiProperty: "buildsAccessLevel", GoGetter: "BuildsAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "buildTimeout", GoGetter: "BuildTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "ciConfigPath", GoGetter: "CiConfigPath"},
			_jsii_.MemberProperty{JsiiProperty: "ciDefaultGitDepth", GoGetter: "CiDefaultGitDepth"},
			_jsii_.MemberProperty{JsiiProperty: "ciForwardDeploymentEnabled", GoGetter: "CiForwardDeploymentEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "containerExpirationPolicy", GoGetter: "ContainerExpirationPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "containerRegistryAccessLevel", GoGetter: "ContainerRegistryAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "containerRegistryEnabled", GoGetter: "ContainerRegistryEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "creatorId", GoGetter: "CreatorId"},
			_jsii_.MemberProperty{JsiiProperty: "customAttributes", GoGetter: "CustomAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "defaultBranch", GoGetter: "DefaultBranch"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "emailsDisabled", GoGetter: "EmailsDisabled"},
			_jsii_.MemberProperty{JsiiProperty: "environmentsAccessLevel", GoGetter: "EnvironmentsAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "externalAuthorizationClassificationLabel", GoGetter: "ExternalAuthorizationClassificationLabel"},
			_jsii_.MemberProperty{JsiiProperty: "featureFlagsAccessLevel", GoGetter: "FeatureFlagsAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "forkedFromProject", GoGetter: "ForkedFromProject"},
			_jsii_.MemberProperty{JsiiProperty: "forkingAccessLevel", GoGetter: "ForkingAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "forksCount", GoGetter: "ForksCount"},
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
			_jsii_.MemberProperty{JsiiProperty: "httpUrlToRepo", GoGetter: "HttpUrlToRepo"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "importError", GoGetter: "ImportError"},
			_jsii_.MemberProperty{JsiiProperty: "importStatus", GoGetter: "ImportStatus"},
			_jsii_.MemberProperty{JsiiProperty: "importUrl", GoGetter: "ImportUrl"},
			_jsii_.MemberProperty{JsiiProperty: "infrastructureAccessLevel", GoGetter: "InfrastructureAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "issuesAccessLevel", GoGetter: "IssuesAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "issuesEnabled", GoGetter: "IssuesEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "jobsEnabled", GoGetter: "JobsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "keepLatestArtifact", GoGetter: "KeepLatestArtifact"},
			_jsii_.MemberProperty{JsiiProperty: "lastActivityAt", GoGetter: "LastActivityAt"},
			_jsii_.MemberProperty{JsiiProperty: "lfsEnabled", GoGetter: "LfsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "links", GoGetter: "Links"},
			_jsii_.MemberProperty{JsiiProperty: "mergeCommitTemplate", GoGetter: "MergeCommitTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "mergeMethod", GoGetter: "MergeMethod"},
			_jsii_.MemberProperty{JsiiProperty: "mergePipelinesEnabled", GoGetter: "MergePipelinesEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "mergeRequestsAccessLevel", GoGetter: "MergeRequestsAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "mergeRequestsEnabled", GoGetter: "MergeRequestsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "mergeTrainsEnabled", GoGetter: "MergeTrainsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "mirror", GoGetter: "Mirror"},
			_jsii_.MemberProperty{JsiiProperty: "mirrorOverwritesDivergedBranches", GoGetter: "MirrorOverwritesDivergedBranches"},
			_jsii_.MemberProperty{JsiiProperty: "mirrorTriggerBuilds", GoGetter: "MirrorTriggerBuilds"},
			_jsii_.MemberProperty{JsiiProperty: "mirrorUserId", GoGetter: "MirrorUserId"},
			_jsii_.MemberProperty{JsiiProperty: "monitorAccessLevel", GoGetter: "MonitorAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "nameWithNamespace", GoGetter: "NameWithNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "onlyAllowMergeIfAllDiscussionsAreResolved", GoGetter: "OnlyAllowMergeIfAllDiscussionsAreResolved"},
			_jsii_.MemberProperty{JsiiProperty: "onlyAllowMergeIfPipelineSucceeds", GoGetter: "OnlyAllowMergeIfPipelineSucceeds"},
			_jsii_.MemberProperty{JsiiProperty: "onlyMirrorProtectedBranches", GoGetter: "OnlyMirrorProtectedBranches"},
			_jsii_.MemberProperty{JsiiProperty: "openIssuesCount", GoGetter: "OpenIssuesCount"},
			_jsii_.MemberProperty{JsiiProperty: "owner", GoGetter: "Owner"},
			_jsii_.MemberProperty{JsiiProperty: "packagesEnabled", GoGetter: "PackagesEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "pathWithNamespace", GoGetter: "PathWithNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
			_jsii_.MemberProperty{JsiiProperty: "public", GoGetter: "Public"},
			_jsii_.MemberProperty{JsiiProperty: "publicBuilds", GoGetter: "PublicBuilds"},
			_jsii_.MemberProperty{JsiiProperty: "readmeUrl", GoGetter: "ReadmeUrl"},
			_jsii_.MemberProperty{JsiiProperty: "releasesAccessLevel", GoGetter: "ReleasesAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryAccessLevel", GoGetter: "RepositoryAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryStorage", GoGetter: "RepositoryStorage"},
			_jsii_.MemberProperty{JsiiProperty: "requestAccessEnabled", GoGetter: "RequestAccessEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "requirementsAccessLevel", GoGetter: "RequirementsAccessLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "resolveOutdatedDiffDiscussions", GoGetter: "ResolveOutdatedDiffDiscussions"},
			_jsii_.MemberProperty{JsiiProperty: "restrictUserDefinedVariables", GoGetter: "RestrictUserDefinedVariables"},
			_jsii_.MemberProperty{JsiiProperty: "runnersToken", GoGetter: "RunnersToken"},
			_jsii_.MemberProperty{JsiiProperty: "securityAndComplianceAccessLevel", GoGetter: "SecurityAndComplianceAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "sharedRunnersEnabled", GoGetter: "SharedRunnersEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "sharedWithGroups", GoGetter: "SharedWithGroups"},
			_jsii_.MemberProperty{JsiiProperty: "snippetsAccessLevel", GoGetter: "SnippetsAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "snippetsEnabled", GoGetter: "SnippetsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "squashCommitTemplate", GoGetter: "SquashCommitTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "sshUrlToRepo", GoGetter: "SshUrlToRepo"},
			_jsii_.MemberProperty{JsiiProperty: "starCount", GoGetter: "StarCount"},
			_jsii_.MemberProperty{JsiiProperty: "statistics", GoGetter: "Statistics"},
			_jsii_.MemberProperty{JsiiProperty: "suggestionCommitMessage", GoGetter: "SuggestionCommitMessage"},
			_jsii_.MemberProperty{JsiiProperty: "tagList", GoGetter: "TagList"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "topics", GoGetter: "Topics"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "visibility", GoGetter: "Visibility"},
			_jsii_.MemberProperty{JsiiProperty: "webUrl", GoGetter: "WebUrl"},
			_jsii_.MemberProperty{JsiiProperty: "wikiAccessLevel", GoGetter: "WikiAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "wikiEnabled", GoGetter: "WikiEnabled"},
		},
		func() interface{} {
			j := jsiiProxy_DataGitlabProjectsProjectsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.dataGitlabProjects.DataGitlabProjectsProjectsOwner",
		reflect.TypeOf((*DataGitlabProjectsProjectsOwner)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProjects.DataGitlabProjectsProjectsOwnerList",
		reflect.TypeOf((*DataGitlabProjectsProjectsOwnerList)(nil)).Elem(),
		[]_jsii_.Member{
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
			j := jsiiProxy_DataGitlabProjectsProjectsOwnerList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProjects.DataGitlabProjectsProjectsOwnerOutputReference",
		reflect.TypeOf((*DataGitlabProjectsProjectsOwnerOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "avatarUrl", GoGetter: "AvatarUrl"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "username", GoGetter: "Username"},
			_jsii_.MemberProperty{JsiiProperty: "websiteUrl", GoGetter: "WebsiteUrl"},
		},
		func() interface{} {
			j := jsiiProxy_DataGitlabProjectsProjectsOwnerOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.dataGitlabProjects.DataGitlabProjectsProjectsPermissions",
		reflect.TypeOf((*DataGitlabProjectsProjectsPermissions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProjects.DataGitlabProjectsProjectsPermissionsList",
		reflect.TypeOf((*DataGitlabProjectsProjectsPermissionsList)(nil)).Elem(),
		[]_jsii_.Member{
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
			j := jsiiProxy_DataGitlabProjectsProjectsPermissionsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProjects.DataGitlabProjectsProjectsPermissionsOutputReference",
		reflect.TypeOf((*DataGitlabProjectsProjectsPermissionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "groupAccess", GoGetter: "GroupAccess"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "projectAccess", GoGetter: "ProjectAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataGitlabProjectsProjectsPermissionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.dataGitlabProjects.DataGitlabProjectsProjectsSharedWithGroups",
		reflect.TypeOf((*DataGitlabProjectsProjectsSharedWithGroups)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProjects.DataGitlabProjectsProjectsSharedWithGroupsList",
		reflect.TypeOf((*DataGitlabProjectsProjectsSharedWithGroupsList)(nil)).Elem(),
		[]_jsii_.Member{
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
			j := jsiiProxy_DataGitlabProjectsProjectsSharedWithGroupsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProjects.DataGitlabProjectsProjectsSharedWithGroupsOutputReference",
		reflect.TypeOf((*DataGitlabProjectsProjectsSharedWithGroupsOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_DataGitlabProjectsProjectsSharedWithGroupsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
