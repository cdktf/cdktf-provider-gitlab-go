// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package userrunner

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.userRunner.UserRunner",
		reflect.TypeOf((*UserRunner)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessLevel", GoGetter: "AccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "accessLevelInput", GoGetter: "AccessLevelInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "locked", GoGetter: "Locked"},
			_jsii_.MemberProperty{JsiiProperty: "lockedInput", GoGetter: "LockedInput"},
			_jsii_.MemberProperty{JsiiProperty: "maximumTimeout", GoGetter: "MaximumTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "maximumTimeoutInput", GoGetter: "MaximumTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "paused", GoGetter: "Paused"},
			_jsii_.MemberProperty{JsiiProperty: "pausedInput", GoGetter: "PausedInput"},
			_jsii_.MemberProperty{JsiiProperty: "projectId", GoGetter: "ProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "projectIdInput", GoGetter: "ProjectIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessLevel", GoMethod: "ResetAccessLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroupId", GoMethod: "ResetGroupId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocked", GoMethod: "ResetLocked"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaximumTimeout", GoMethod: "ResetMaximumTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPaused", GoMethod: "ResetPaused"},
			_jsii_.MemberMethod{JsiiMethod: "resetProjectId", GoMethod: "ResetProjectId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagList", GoMethod: "ResetTagList"},
			_jsii_.MemberMethod{JsiiMethod: "resetUntagged", GoMethod: "ResetUntagged"},
			_jsii_.MemberProperty{JsiiProperty: "runnerType", GoGetter: "RunnerType"},
			_jsii_.MemberProperty{JsiiProperty: "runnerTypeInput", GoGetter: "RunnerTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tagList", GoGetter: "TagList"},
			_jsii_.MemberProperty{JsiiProperty: "tagListInput", GoGetter: "TagListInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "token", GoGetter: "Token"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "untagged", GoGetter: "Untagged"},
			_jsii_.MemberProperty{JsiiProperty: "untaggedInput", GoGetter: "UntaggedInput"},
		},
		func() interface{} {
			j := jsiiProxy_UserRunner{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.userRunner.UserRunnerConfig",
		reflect.TypeOf((*UserRunnerConfig)(nil)).Elem(),
	)
}
