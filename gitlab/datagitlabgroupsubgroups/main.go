// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabgroupsubgroups

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabGroupSubgroups.DataGitlabGroupSubgroups",
		reflect.TypeOf((*DataGitlabGroupSubgroups)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allAvailable", GoGetter: "AllAvailable"},
			_jsii_.MemberProperty{JsiiProperty: "allAvailableInput", GoGetter: "AllAvailableInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "minAccessLevel", GoGetter: "MinAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "minAccessLevelInput", GoGetter: "MinAccessLevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "orderBy", GoGetter: "OrderBy"},
			_jsii_.MemberProperty{JsiiProperty: "orderByInput", GoGetter: "OrderByInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "owned", GoGetter: "Owned"},
			_jsii_.MemberProperty{JsiiProperty: "ownedInput", GoGetter: "OwnedInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllAvailable", GoMethod: "ResetAllAvailable"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinAccessLevel", GoMethod: "ResetMinAccessLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrderBy", GoMethod: "ResetOrderBy"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOwned", GoMethod: "ResetOwned"},
			_jsii_.MemberMethod{JsiiMethod: "resetSearch", GoMethod: "ResetSearch"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipGroups", GoMethod: "ResetSkipGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetSort", GoMethod: "ResetSort"},
			_jsii_.MemberMethod{JsiiMethod: "resetStatistics", GoMethod: "ResetStatistics"},
			_jsii_.MemberMethod{JsiiMethod: "resetWithCustomAttributes", GoMethod: "ResetWithCustomAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "search", GoGetter: "Search"},
			_jsii_.MemberProperty{JsiiProperty: "searchInput", GoGetter: "SearchInput"},
			_jsii_.MemberProperty{JsiiProperty: "skipGroups", GoGetter: "SkipGroups"},
			_jsii_.MemberProperty{JsiiProperty: "skipGroupsInput", GoGetter: "SkipGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "sort", GoGetter: "Sort"},
			_jsii_.MemberProperty{JsiiProperty: "sortInput", GoGetter: "SortInput"},
			_jsii_.MemberProperty{JsiiProperty: "statistics", GoGetter: "Statistics"},
			_jsii_.MemberProperty{JsiiProperty: "statisticsInput", GoGetter: "StatisticsInput"},
			_jsii_.MemberProperty{JsiiProperty: "subgroups", GoGetter: "Subgroups"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "withCustomAttributes", GoGetter: "WithCustomAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "withCustomAttributesInput", GoGetter: "WithCustomAttributesInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataGitlabGroupSubgroups{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.dataGitlabGroupSubgroups.DataGitlabGroupSubgroupsConfig",
		reflect.TypeOf((*DataGitlabGroupSubgroupsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.dataGitlabGroupSubgroups.DataGitlabGroupSubgroupsSubgroups",
		reflect.TypeOf((*DataGitlabGroupSubgroupsSubgroups)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabGroupSubgroups.DataGitlabGroupSubgroupsSubgroupsList",
		reflect.TypeOf((*DataGitlabGroupSubgroupsSubgroupsList)(nil)).Elem(),
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
			j := jsiiProxy_DataGitlabGroupSubgroupsSubgroupsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabGroupSubgroups.DataGitlabGroupSubgroupsSubgroupsOutputReference",
		reflect.TypeOf((*DataGitlabGroupSubgroupsSubgroupsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "autoDevopsEnabled", GoGetter: "AutoDevopsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "avatarUrl", GoGetter: "AvatarUrl"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultBranchProtection", GoGetter: "DefaultBranchProtection"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "emailsEnabled", GoGetter: "EmailsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "fileTemplateProjectId", GoGetter: "FileTemplateProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "fullName", GoGetter: "FullName"},
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
			_jsii_.MemberProperty{JsiiProperty: "groupId", GoGetter: "GroupId"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipRestrictionRanges", GoGetter: "IpRestrictionRanges"},
			_jsii_.MemberProperty{JsiiProperty: "lfsEnabled", GoGetter: "LfsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "mentionsDisabled", GoGetter: "MentionsDisabled"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "parentId", GoGetter: "ParentId"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "projectCreationLevel", GoGetter: "ProjectCreationLevel"},
			_jsii_.MemberProperty{JsiiProperty: "requestAccessEnabled", GoGetter: "RequestAccessEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "requireTwoFactorAuthentication", GoGetter: "RequireTwoFactorAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sharedRunnersSetting", GoGetter: "SharedRunnersSetting"},
			_jsii_.MemberProperty{JsiiProperty: "shareWithGroupLock", GoGetter: "ShareWithGroupLock"},
			_jsii_.MemberProperty{JsiiProperty: "statistics", GoGetter: "Statistics"},
			_jsii_.MemberProperty{JsiiProperty: "subgroupCreationLevel", GoGetter: "SubgroupCreationLevel"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "twoFactorGracePeriod", GoGetter: "TwoFactorGracePeriod"},
			_jsii_.MemberProperty{JsiiProperty: "visibility", GoGetter: "Visibility"},
			_jsii_.MemberProperty{JsiiProperty: "webUrl", GoGetter: "WebUrl"},
			_jsii_.MemberProperty{JsiiProperty: "wikiAccessLevel", GoGetter: "WikiAccessLevel"},
		},
		func() interface{} {
			j := jsiiProxy_DataGitlabGroupSubgroupsSubgroupsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
