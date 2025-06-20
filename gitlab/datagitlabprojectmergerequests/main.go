// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabprojectmergerequests

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProjectMergeRequests.DataGitlabProjectMergeRequests",
		reflect.TypeOf((*DataGitlabProjectMergeRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "authorId", GoGetter: "AuthorId"},
			_jsii_.MemberProperty{JsiiProperty: "authorIdInput", GoGetter: "AuthorIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "authorUsername", GoGetter: "AuthorUsername"},
			_jsii_.MemberProperty{JsiiProperty: "authorUsernameInput", GoGetter: "AuthorUsernameInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createdAfter", GoGetter: "CreatedAfter"},
			_jsii_.MemberProperty{JsiiProperty: "createdAfterInput", GoGetter: "CreatedAfterInput"},
			_jsii_.MemberProperty{JsiiProperty: "createdBefore", GoGetter: "CreatedBefore"},
			_jsii_.MemberProperty{JsiiProperty: "createdBeforeInput", GoGetter: "CreatedBeforeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "iids", GoGetter: "Iids"},
			_jsii_.MemberProperty{JsiiProperty: "iidsInput", GoGetter: "IidsInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "mergeRequests", GoGetter: "MergeRequests"},
			_jsii_.MemberProperty{JsiiProperty: "milestone", GoGetter: "Milestone"},
			_jsii_.MemberProperty{JsiiProperty: "milestoneInput", GoGetter: "MilestoneInput"},
			_jsii_.MemberProperty{JsiiProperty: "myReactionEmoji", GoGetter: "MyReactionEmoji"},
			_jsii_.MemberProperty{JsiiProperty: "myReactionEmojiInput", GoGetter: "MyReactionEmojiInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "orderBy", GoGetter: "OrderBy"},
			_jsii_.MemberProperty{JsiiProperty: "orderByInput", GoGetter: "OrderByInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "projectInput", GoGetter: "ProjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthorId", GoMethod: "ResetAuthorId"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthorUsername", GoMethod: "ResetAuthorUsername"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreatedAfter", GoMethod: "ResetCreatedAfter"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreatedBefore", GoMethod: "ResetCreatedBefore"},
			_jsii_.MemberMethod{JsiiMethod: "resetIids", GoMethod: "ResetIids"},
			_jsii_.MemberMethod{JsiiMethod: "resetMilestone", GoMethod: "ResetMilestone"},
			_jsii_.MemberMethod{JsiiMethod: "resetMyReactionEmoji", GoMethod: "ResetMyReactionEmoji"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrderBy", GoMethod: "ResetOrderBy"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetReviewerUsername", GoMethod: "ResetReviewerUsername"},
			_jsii_.MemberMethod{JsiiMethod: "resetScope", GoMethod: "ResetScope"},
			_jsii_.MemberMethod{JsiiMethod: "resetSearch", GoMethod: "ResetSearch"},
			_jsii_.MemberMethod{JsiiMethod: "resetSort", GoMethod: "ResetSort"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceBranch", GoMethod: "ResetSourceBranch"},
			_jsii_.MemberMethod{JsiiMethod: "resetState", GoMethod: "ResetState"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetBranch", GoMethod: "ResetTargetBranch"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdatedAfter", GoMethod: "ResetUpdatedAfter"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdatedBefore", GoMethod: "ResetUpdatedBefore"},
			_jsii_.MemberMethod{JsiiMethod: "resetWip", GoMethod: "ResetWip"},
			_jsii_.MemberProperty{JsiiProperty: "reviewerUsername", GoGetter: "ReviewerUsername"},
			_jsii_.MemberProperty{JsiiProperty: "reviewerUsernameInput", GoGetter: "ReviewerUsernameInput"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberProperty{JsiiProperty: "scopeInput", GoGetter: "ScopeInput"},
			_jsii_.MemberProperty{JsiiProperty: "search", GoGetter: "Search"},
			_jsii_.MemberProperty{JsiiProperty: "searchInput", GoGetter: "SearchInput"},
			_jsii_.MemberProperty{JsiiProperty: "sort", GoGetter: "Sort"},
			_jsii_.MemberProperty{JsiiProperty: "sortInput", GoGetter: "SortInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceBranch", GoGetter: "SourceBranch"},
			_jsii_.MemberProperty{JsiiProperty: "sourceBranchInput", GoGetter: "SourceBranchInput"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberProperty{JsiiProperty: "stateInput", GoGetter: "StateInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "targetBranch", GoGetter: "TargetBranch"},
			_jsii_.MemberProperty{JsiiProperty: "targetBranchInput", GoGetter: "TargetBranchInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAfter", GoGetter: "UpdatedAfter"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAfterInput", GoGetter: "UpdatedAfterInput"},
			_jsii_.MemberProperty{JsiiProperty: "updatedBefore", GoGetter: "UpdatedBefore"},
			_jsii_.MemberProperty{JsiiProperty: "updatedBeforeInput", GoGetter: "UpdatedBeforeInput"},
			_jsii_.MemberProperty{JsiiProperty: "wip", GoGetter: "Wip"},
			_jsii_.MemberProperty{JsiiProperty: "wipInput", GoGetter: "WipInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataGitlabProjectMergeRequests{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.dataGitlabProjectMergeRequests.DataGitlabProjectMergeRequestsConfig",
		reflect.TypeOf((*DataGitlabProjectMergeRequestsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.dataGitlabProjectMergeRequests.DataGitlabProjectMergeRequestsMergeRequests",
		reflect.TypeOf((*DataGitlabProjectMergeRequestsMergeRequests)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.dataGitlabProjectMergeRequests.DataGitlabProjectMergeRequestsMergeRequestsAssignee",
		reflect.TypeOf((*DataGitlabProjectMergeRequestsMergeRequestsAssignee)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProjectMergeRequests.DataGitlabProjectMergeRequestsMergeRequestsAssigneeOutputReference",
		reflect.TypeOf((*DataGitlabProjectMergeRequestsMergeRequestsAssigneeOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "webUrl", GoGetter: "WebUrl"},
		},
		func() interface{} {
			j := jsiiProxy_DataGitlabProjectMergeRequestsMergeRequestsAssigneeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.dataGitlabProjectMergeRequests.DataGitlabProjectMergeRequestsMergeRequestsAssignees",
		reflect.TypeOf((*DataGitlabProjectMergeRequestsMergeRequestsAssignees)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProjectMergeRequests.DataGitlabProjectMergeRequestsMergeRequestsAssigneesList",
		reflect.TypeOf((*DataGitlabProjectMergeRequestsMergeRequestsAssigneesList)(nil)).Elem(),
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
			j := jsiiProxy_DataGitlabProjectMergeRequestsMergeRequestsAssigneesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProjectMergeRequests.DataGitlabProjectMergeRequestsMergeRequestsAssigneesOutputReference",
		reflect.TypeOf((*DataGitlabProjectMergeRequestsMergeRequestsAssigneesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "webUrl", GoGetter: "WebUrl"},
		},
		func() interface{} {
			j := jsiiProxy_DataGitlabProjectMergeRequestsMergeRequestsAssigneesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.dataGitlabProjectMergeRequests.DataGitlabProjectMergeRequestsMergeRequestsAuthor",
		reflect.TypeOf((*DataGitlabProjectMergeRequestsMergeRequestsAuthor)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProjectMergeRequests.DataGitlabProjectMergeRequestsMergeRequestsAuthorOutputReference",
		reflect.TypeOf((*DataGitlabProjectMergeRequestsMergeRequestsAuthorOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "webUrl", GoGetter: "WebUrl"},
		},
		func() interface{} {
			j := jsiiProxy_DataGitlabProjectMergeRequestsMergeRequestsAuthorOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.dataGitlabProjectMergeRequests.DataGitlabProjectMergeRequestsMergeRequestsClosedBy",
		reflect.TypeOf((*DataGitlabProjectMergeRequestsMergeRequestsClosedBy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProjectMergeRequests.DataGitlabProjectMergeRequestsMergeRequestsClosedByOutputReference",
		reflect.TypeOf((*DataGitlabProjectMergeRequestsMergeRequestsClosedByOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "webUrl", GoGetter: "WebUrl"},
		},
		func() interface{} {
			j := jsiiProxy_DataGitlabProjectMergeRequestsMergeRequestsClosedByOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProjectMergeRequests.DataGitlabProjectMergeRequestsMergeRequestsList",
		reflect.TypeOf((*DataGitlabProjectMergeRequestsMergeRequestsList)(nil)).Elem(),
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
			j := jsiiProxy_DataGitlabProjectMergeRequestsMergeRequestsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProjectMergeRequests.DataGitlabProjectMergeRequestsMergeRequestsOutputReference",
		reflect.TypeOf((*DataGitlabProjectMergeRequestsMergeRequestsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "assignee", GoGetter: "Assignee"},
			_jsii_.MemberProperty{JsiiProperty: "assignees", GoGetter: "Assignees"},
			_jsii_.MemberProperty{JsiiProperty: "author", GoGetter: "Author"},
			_jsii_.MemberProperty{JsiiProperty: "blockingDiscussionsResolved", GoGetter: "BlockingDiscussionsResolved"},
			_jsii_.MemberProperty{JsiiProperty: "closedAt", GoGetter: "ClosedAt"},
			_jsii_.MemberProperty{JsiiProperty: "closedBy", GoGetter: "ClosedBy"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
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
			_jsii_.MemberProperty{JsiiProperty: "iid", GoGetter: "Iid"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataGitlabProjectMergeRequestsMergeRequestsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
