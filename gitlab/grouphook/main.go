// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grouphook

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.groupHook.GroupHook",
		reflect.TypeOf((*GroupHook)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "confidentialIssuesEvents", GoGetter: "ConfidentialIssuesEvents"},
			_jsii_.MemberProperty{JsiiProperty: "confidentialIssuesEventsInput", GoGetter: "ConfidentialIssuesEventsInput"},
			_jsii_.MemberProperty{JsiiProperty: "confidentialNoteEvents", GoGetter: "ConfidentialNoteEvents"},
			_jsii_.MemberProperty{JsiiProperty: "confidentialNoteEventsInput", GoGetter: "ConfidentialNoteEventsInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "customHeaders", GoGetter: "CustomHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "customHeadersInput", GoGetter: "CustomHeadersInput"},
			_jsii_.MemberProperty{JsiiProperty: "customWebhookTemplate", GoGetter: "CustomWebhookTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "customWebhookTemplateInput", GoGetter: "CustomWebhookTemplateInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentEvents", GoGetter: "DeploymentEvents"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentEventsInput", GoGetter: "DeploymentEventsInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableSslVerification", GoGetter: "EnableSslVerification"},
			_jsii_.MemberProperty{JsiiProperty: "enableSslVerificationInput", GoGetter: "EnableSslVerificationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "group", GoGetter: "Group"},
			_jsii_.MemberProperty{JsiiProperty: "groupId", GoGetter: "GroupId"},
			_jsii_.MemberProperty{JsiiProperty: "groupInput", GoGetter: "GroupInput"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "hookId", GoGetter: "HookId"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "issuesEvents", GoGetter: "IssuesEvents"},
			_jsii_.MemberProperty{JsiiProperty: "issuesEventsInput", GoGetter: "IssuesEventsInput"},
			_jsii_.MemberProperty{JsiiProperty: "jobEvents", GoGetter: "JobEvents"},
			_jsii_.MemberProperty{JsiiProperty: "jobEventsInput", GoGetter: "JobEventsInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "mergeRequestsEvents", GoGetter: "MergeRequestsEvents"},
			_jsii_.MemberProperty{JsiiProperty: "mergeRequestsEventsInput", GoGetter: "MergeRequestsEventsInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "noteEvents", GoGetter: "NoteEvents"},
			_jsii_.MemberProperty{JsiiProperty: "noteEventsInput", GoGetter: "NoteEventsInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pipelineEvents", GoGetter: "PipelineEvents"},
			_jsii_.MemberProperty{JsiiProperty: "pipelineEventsInput", GoGetter: "PipelineEventsInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "pushEvents", GoGetter: "PushEvents"},
			_jsii_.MemberProperty{JsiiProperty: "pushEventsBranchFilter", GoGetter: "PushEventsBranchFilter"},
			_jsii_.MemberProperty{JsiiProperty: "pushEventsBranchFilterInput", GoGetter: "PushEventsBranchFilterInput"},
			_jsii_.MemberProperty{JsiiProperty: "pushEventsInput", GoGetter: "PushEventsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomHeaders", GoMethod: "PutCustomHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "releasesEvents", GoGetter: "ReleasesEvents"},
			_jsii_.MemberProperty{JsiiProperty: "releasesEventsInput", GoGetter: "ReleasesEventsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfidentialIssuesEvents", GoMethod: "ResetConfidentialIssuesEvents"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfidentialNoteEvents", GoMethod: "ResetConfidentialNoteEvents"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomHeaders", GoMethod: "ResetCustomHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomWebhookTemplate", GoMethod: "ResetCustomWebhookTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeploymentEvents", GoMethod: "ResetDeploymentEvents"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableSslVerification", GoMethod: "ResetEnableSslVerification"},
			_jsii_.MemberMethod{JsiiMethod: "resetIssuesEvents", GoMethod: "ResetIssuesEvents"},
			_jsii_.MemberMethod{JsiiMethod: "resetJobEvents", GoMethod: "ResetJobEvents"},
			_jsii_.MemberMethod{JsiiMethod: "resetMergeRequestsEvents", GoMethod: "ResetMergeRequestsEvents"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoteEvents", GoMethod: "ResetNoteEvents"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPipelineEvents", GoMethod: "ResetPipelineEvents"},
			_jsii_.MemberMethod{JsiiMethod: "resetPushEvents", GoMethod: "ResetPushEvents"},
			_jsii_.MemberMethod{JsiiMethod: "resetPushEventsBranchFilter", GoMethod: "ResetPushEventsBranchFilter"},
			_jsii_.MemberMethod{JsiiMethod: "resetReleasesEvents", GoMethod: "ResetReleasesEvents"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubgroupEvents", GoMethod: "ResetSubgroupEvents"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagPushEvents", GoMethod: "ResetTagPushEvents"},
			_jsii_.MemberMethod{JsiiMethod: "resetToken", GoMethod: "ResetToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetWikiPageEvents", GoMethod: "ResetWikiPageEvents"},
			_jsii_.MemberProperty{JsiiProperty: "subgroupEvents", GoGetter: "SubgroupEvents"},
			_jsii_.MemberProperty{JsiiProperty: "subgroupEventsInput", GoGetter: "SubgroupEventsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "urlInput", GoGetter: "UrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "wikiPageEvents", GoGetter: "WikiPageEvents"},
			_jsii_.MemberProperty{JsiiProperty: "wikiPageEventsInput", GoGetter: "WikiPageEventsInput"},
		},
		func() interface{} {
			j := jsiiProxy_GroupHook{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.groupHook.GroupHookConfig",
		reflect.TypeOf((*GroupHookConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.groupHook.GroupHookCustomHeaders",
		reflect.TypeOf((*GroupHookCustomHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.groupHook.GroupHookCustomHeadersList",
		reflect.TypeOf((*GroupHookCustomHeadersList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_GroupHookCustomHeadersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.groupHook.GroupHookCustomHeadersOutputReference",
		reflect.TypeOf((*GroupHookCustomHeadersOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "keyInput", GoGetter: "KeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_GroupHookCustomHeadersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
