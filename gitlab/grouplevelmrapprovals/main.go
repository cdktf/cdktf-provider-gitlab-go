// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grouplevelmrapprovals

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.groupLevelMrApprovals.GroupLevelMrApprovals",
		reflect.TypeOf((*GroupLevelMrApprovals)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowAuthorApproval", GoGetter: "AllowAuthorApproval"},
			_jsii_.MemberProperty{JsiiProperty: "allowAuthorApprovalInput", GoGetter: "AllowAuthorApprovalInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowCommitterApproval", GoGetter: "AllowCommitterApproval"},
			_jsii_.MemberProperty{JsiiProperty: "allowCommitterApprovalInput", GoGetter: "AllowCommitterApprovalInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowOverridesToApproverListPerMergeRequest", GoGetter: "AllowOverridesToApproverListPerMergeRequest"},
			_jsii_.MemberProperty{JsiiProperty: "allowOverridesToApproverListPerMergeRequestInput", GoGetter: "AllowOverridesToApproverListPerMergeRequestInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "group", GoGetter: "Group"},
			_jsii_.MemberProperty{JsiiProperty: "groupInput", GoGetter: "GroupInput"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keepSettingsOnDestroy", GoGetter: "KeepSettingsOnDestroy"},
			_jsii_.MemberProperty{JsiiProperty: "keepSettingsOnDestroyInput", GoGetter: "KeepSettingsOnDestroyInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "requireReauthenticationToApprove", GoGetter: "RequireReauthenticationToApprove"},
			_jsii_.MemberProperty{JsiiProperty: "requireReauthenticationToApproveInput", GoGetter: "RequireReauthenticationToApproveInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowAuthorApproval", GoMethod: "ResetAllowAuthorApproval"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowCommitterApproval", GoMethod: "ResetAllowCommitterApproval"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowOverridesToApproverListPerMergeRequest", GoMethod: "ResetAllowOverridesToApproverListPerMergeRequest"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeepSettingsOnDestroy", GoMethod: "ResetKeepSettingsOnDestroy"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequireReauthenticationToApprove", GoMethod: "ResetRequireReauthenticationToApprove"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetainApprovalsOnPush", GoMethod: "ResetRetainApprovalsOnPush"},
			_jsii_.MemberProperty{JsiiProperty: "retainApprovalsOnPush", GoGetter: "RetainApprovalsOnPush"},
			_jsii_.MemberProperty{JsiiProperty: "retainApprovalsOnPushInput", GoGetter: "RetainApprovalsOnPushInput"},
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
			j := jsiiProxy_GroupLevelMrApprovals{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.groupLevelMrApprovals.GroupLevelMrApprovalsConfig",
		reflect.TypeOf((*GroupLevelMrApprovalsConfig)(nil)).Elem(),
	)
}
