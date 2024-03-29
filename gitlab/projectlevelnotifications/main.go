// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectlevelnotifications

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.projectLevelNotifications.ProjectLevelNotifications",
		reflect.TypeOf((*ProjectLevelNotifications)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "closeIssue", GoGetter: "CloseIssue"},
			_jsii_.MemberProperty{JsiiProperty: "closeIssueInput", GoGetter: "CloseIssueInput"},
			_jsii_.MemberProperty{JsiiProperty: "closeMergeRequest", GoGetter: "CloseMergeRequest"},
			_jsii_.MemberProperty{JsiiProperty: "closeMergeRequestInput", GoGetter: "CloseMergeRequestInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "failedPipeline", GoGetter: "FailedPipeline"},
			_jsii_.MemberProperty{JsiiProperty: "failedPipelineInput", GoGetter: "FailedPipelineInput"},
			_jsii_.MemberProperty{JsiiProperty: "fixedPipeline", GoGetter: "FixedPipeline"},
			_jsii_.MemberProperty{JsiiProperty: "fixedPipelineInput", GoGetter: "FixedPipelineInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "issueDue", GoGetter: "IssueDue"},
			_jsii_.MemberProperty{JsiiProperty: "issueDueInput", GoGetter: "IssueDueInput"},
			_jsii_.MemberProperty{JsiiProperty: "level", GoGetter: "Level"},
			_jsii_.MemberProperty{JsiiProperty: "levelInput", GoGetter: "LevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "mergeMergeRequest", GoGetter: "MergeMergeRequest"},
			_jsii_.MemberProperty{JsiiProperty: "mergeMergeRequestInput", GoGetter: "MergeMergeRequestInput"},
			_jsii_.MemberProperty{JsiiProperty: "mergeWhenPipelineSucceeds", GoGetter: "MergeWhenPipelineSucceeds"},
			_jsii_.MemberProperty{JsiiProperty: "mergeWhenPipelineSucceedsInput", GoGetter: "MergeWhenPipelineSucceedsInput"},
			_jsii_.MemberProperty{JsiiProperty: "movedProject", GoGetter: "MovedProject"},
			_jsii_.MemberProperty{JsiiProperty: "movedProjectInput", GoGetter: "MovedProjectInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "newIssue", GoGetter: "NewIssue"},
			_jsii_.MemberProperty{JsiiProperty: "newIssueInput", GoGetter: "NewIssueInput"},
			_jsii_.MemberProperty{JsiiProperty: "newMergeRequest", GoGetter: "NewMergeRequest"},
			_jsii_.MemberProperty{JsiiProperty: "newMergeRequestInput", GoGetter: "NewMergeRequestInput"},
			_jsii_.MemberProperty{JsiiProperty: "newNote", GoGetter: "NewNote"},
			_jsii_.MemberProperty{JsiiProperty: "newNoteInput", GoGetter: "NewNoteInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "projectInput", GoGetter: "ProjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "pushToMergeRequest", GoGetter: "PushToMergeRequest"},
			_jsii_.MemberProperty{JsiiProperty: "pushToMergeRequestInput", GoGetter: "PushToMergeRequestInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "reassignIssue", GoGetter: "ReassignIssue"},
			_jsii_.MemberProperty{JsiiProperty: "reassignIssueInput", GoGetter: "ReassignIssueInput"},
			_jsii_.MemberProperty{JsiiProperty: "reassignMergeRequest", GoGetter: "ReassignMergeRequest"},
			_jsii_.MemberProperty{JsiiProperty: "reassignMergeRequestInput", GoGetter: "ReassignMergeRequestInput"},
			_jsii_.MemberProperty{JsiiProperty: "reopenIssue", GoGetter: "ReopenIssue"},
			_jsii_.MemberProperty{JsiiProperty: "reopenIssueInput", GoGetter: "ReopenIssueInput"},
			_jsii_.MemberProperty{JsiiProperty: "reopenMergeRequest", GoGetter: "ReopenMergeRequest"},
			_jsii_.MemberProperty{JsiiProperty: "reopenMergeRequestInput", GoGetter: "ReopenMergeRequestInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloseIssue", GoMethod: "ResetCloseIssue"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloseMergeRequest", GoMethod: "ResetCloseMergeRequest"},
			_jsii_.MemberMethod{JsiiMethod: "resetFailedPipeline", GoMethod: "ResetFailedPipeline"},
			_jsii_.MemberMethod{JsiiMethod: "resetFixedPipeline", GoMethod: "ResetFixedPipeline"},
			_jsii_.MemberMethod{JsiiMethod: "resetIssueDue", GoMethod: "ResetIssueDue"},
			_jsii_.MemberMethod{JsiiMethod: "resetLevel", GoMethod: "ResetLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetMergeMergeRequest", GoMethod: "ResetMergeMergeRequest"},
			_jsii_.MemberMethod{JsiiMethod: "resetMergeWhenPipelineSucceeds", GoMethod: "ResetMergeWhenPipelineSucceeds"},
			_jsii_.MemberMethod{JsiiMethod: "resetMovedProject", GoMethod: "ResetMovedProject"},
			_jsii_.MemberMethod{JsiiMethod: "resetNewIssue", GoMethod: "ResetNewIssue"},
			_jsii_.MemberMethod{JsiiMethod: "resetNewMergeRequest", GoMethod: "ResetNewMergeRequest"},
			_jsii_.MemberMethod{JsiiMethod: "resetNewNote", GoMethod: "ResetNewNote"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPushToMergeRequest", GoMethod: "ResetPushToMergeRequest"},
			_jsii_.MemberMethod{JsiiMethod: "resetReassignIssue", GoMethod: "ResetReassignIssue"},
			_jsii_.MemberMethod{JsiiMethod: "resetReassignMergeRequest", GoMethod: "ResetReassignMergeRequest"},
			_jsii_.MemberMethod{JsiiMethod: "resetReopenIssue", GoMethod: "ResetReopenIssue"},
			_jsii_.MemberMethod{JsiiMethod: "resetReopenMergeRequest", GoMethod: "ResetReopenMergeRequest"},
			_jsii_.MemberMethod{JsiiMethod: "resetSuccessPipeline", GoMethod: "ResetSuccessPipeline"},
			_jsii_.MemberProperty{JsiiProperty: "successPipeline", GoGetter: "SuccessPipeline"},
			_jsii_.MemberProperty{JsiiProperty: "successPipelineInput", GoGetter: "SuccessPipelineInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_ProjectLevelNotifications{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.projectLevelNotifications.ProjectLevelNotificationsConfig",
		reflect.TypeOf((*ProjectLevelNotificationsConfig)(nil)).Elem(),
	)
}
