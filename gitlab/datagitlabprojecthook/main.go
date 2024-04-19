// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabprojecthook

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabProjectHook.DataGitlabProjectHook",
		reflect.TypeOf((*DataGitlabProjectHook)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "confidentialIssuesEvents", GoGetter: "ConfidentialIssuesEvents"},
			_jsii_.MemberProperty{JsiiProperty: "confidentialNoteEvents", GoGetter: "ConfidentialNoteEvents"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "customWebhookTemplate", GoGetter: "CustomWebhookTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentEvents", GoGetter: "DeploymentEvents"},
			_jsii_.MemberProperty{JsiiProperty: "enableSslVerification", GoGetter: "EnableSslVerification"},
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
			_jsii_.MemberProperty{JsiiProperty: "hookId", GoGetter: "HookId"},
			_jsii_.MemberProperty{JsiiProperty: "hookIdInput", GoGetter: "HookIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "issuesEvents", GoGetter: "IssuesEvents"},
			_jsii_.MemberProperty{JsiiProperty: "jobEvents", GoGetter: "JobEvents"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "mergeRequestsEvents", GoGetter: "MergeRequestsEvents"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "noteEvents", GoGetter: "NoteEvents"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pipelineEvents", GoGetter: "PipelineEvents"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "projectId", GoGetter: "ProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "projectInput", GoGetter: "ProjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "pushEvents", GoGetter: "PushEvents"},
			_jsii_.MemberProperty{JsiiProperty: "pushEventsBranchFilter", GoGetter: "PushEventsBranchFilter"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "releasesEvents", GoGetter: "ReleasesEvents"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tagPushEvents", GoGetter: "TagPushEvents"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "token", GoGetter: "Token"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "wikiPageEvents", GoGetter: "WikiPageEvents"},
		},
		func() interface{} {
			j := jsiiProxy_DataGitlabProjectHook{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.dataGitlabProjectHook.DataGitlabProjectHookConfig",
		reflect.TypeOf((*DataGitlabProjectHookConfig)(nil)).Elem(),
	)
}
