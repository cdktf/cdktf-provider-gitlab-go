// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationtelegram

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.integrationTelegram.IntegrationTelegram",
		reflect.TypeOf((*IntegrationTelegram)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "branchesToBeNotified", GoGetter: "BranchesToBeNotified"},
			_jsii_.MemberProperty{JsiiProperty: "branchesToBeNotifiedInput", GoGetter: "BranchesToBeNotifiedInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "confidentialIssuesEvents", GoGetter: "ConfidentialIssuesEvents"},
			_jsii_.MemberProperty{JsiiProperty: "confidentialIssuesEventsInput", GoGetter: "ConfidentialIssuesEventsInput"},
			_jsii_.MemberProperty{JsiiProperty: "confidentialNoteEvents", GoGetter: "ConfidentialNoteEvents"},
			_jsii_.MemberProperty{JsiiProperty: "confidentialNoteEventsInput", GoGetter: "ConfidentialNoteEventsInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
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
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "issuesEvents", GoGetter: "IssuesEvents"},
			_jsii_.MemberProperty{JsiiProperty: "issuesEventsInput", GoGetter: "IssuesEventsInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "mergeRequestsEvents", GoGetter: "MergeRequestsEvents"},
			_jsii_.MemberProperty{JsiiProperty: "mergeRequestsEventsInput", GoGetter: "MergeRequestsEventsInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "noteEvents", GoGetter: "NoteEvents"},
			_jsii_.MemberProperty{JsiiProperty: "noteEventsInput", GoGetter: "NoteEventsInput"},
			_jsii_.MemberProperty{JsiiProperty: "notifyOnlyBrokenPipelines", GoGetter: "NotifyOnlyBrokenPipelines"},
			_jsii_.MemberProperty{JsiiProperty: "notifyOnlyBrokenPipelinesInput", GoGetter: "NotifyOnlyBrokenPipelinesInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pipelineEvents", GoGetter: "PipelineEvents"},
			_jsii_.MemberProperty{JsiiProperty: "pipelineEventsInput", GoGetter: "PipelineEventsInput"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "projectInput", GoGetter: "ProjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "pushEvents", GoGetter: "PushEvents"},
			_jsii_.MemberProperty{JsiiProperty: "pushEventsInput", GoGetter: "PushEventsInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetBranchesToBeNotified", GoMethod: "ResetBranchesToBeNotified"},
			_jsii_.MemberMethod{JsiiMethod: "resetNotifyOnlyBrokenPipelines", GoMethod: "ResetNotifyOnlyBrokenPipelines"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "room", GoGetter: "Room"},
			_jsii_.MemberProperty{JsiiProperty: "roomInput", GoGetter: "RoomInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tagPushEvents", GoGetter: "TagPushEvents"},
			_jsii_.MemberProperty{JsiiProperty: "tagPushEventsInput", GoGetter: "TagPushEventsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "token", GoGetter: "Token"},
			_jsii_.MemberProperty{JsiiProperty: "tokenInput", GoGetter: "TokenInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "wikiPageEvents", GoGetter: "WikiPageEvents"},
			_jsii_.MemberProperty{JsiiProperty: "wikiPageEventsInput", GoGetter: "WikiPageEventsInput"},
		},
		func() interface{} {
			j := jsiiProxy_IntegrationTelegram{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.integrationTelegram.IntegrationTelegramConfig",
		reflect.TypeOf((*IntegrationTelegramConfig)(nil)).Elem(),
	)
}
