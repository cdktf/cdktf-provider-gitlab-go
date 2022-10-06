package projectlevelmrapprovals

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.projectLevelMrApprovals.ProjectLevelMrApprovals",
		reflect.TypeOf((*ProjectLevelMrApprovals)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "disableOverridingApproversPerMergeRequest", GoGetter: "DisableOverridingApproversPerMergeRequest"},
			_jsii_.MemberProperty{JsiiProperty: "disableOverridingApproversPerMergeRequestInput", GoGetter: "DisableOverridingApproversPerMergeRequestInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "mergeRequestsAuthorApproval", GoGetter: "MergeRequestsAuthorApproval"},
			_jsii_.MemberProperty{JsiiProperty: "mergeRequestsAuthorApprovalInput", GoGetter: "MergeRequestsAuthorApprovalInput"},
			_jsii_.MemberProperty{JsiiProperty: "mergeRequestsDisableCommittersApproval", GoGetter: "MergeRequestsDisableCommittersApproval"},
			_jsii_.MemberProperty{JsiiProperty: "mergeRequestsDisableCommittersApprovalInput", GoGetter: "MergeRequestsDisableCommittersApprovalInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "projectId", GoGetter: "ProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "projectIdInput", GoGetter: "ProjectIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "requirePasswordToApprove", GoGetter: "RequirePasswordToApprove"},
			_jsii_.MemberProperty{JsiiProperty: "requirePasswordToApproveInput", GoGetter: "RequirePasswordToApproveInput"},
			_jsii_.MemberProperty{JsiiProperty: "resetApprovalsOnPush", GoGetter: "ResetApprovalsOnPush"},
			_jsii_.MemberProperty{JsiiProperty: "resetApprovalsOnPushInput", GoGetter: "ResetApprovalsOnPushInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableOverridingApproversPerMergeRequest", GoMethod: "ResetDisableOverridingApproversPerMergeRequest"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMergeRequestsAuthorApproval", GoMethod: "ResetMergeRequestsAuthorApproval"},
			_jsii_.MemberMethod{JsiiMethod: "resetMergeRequestsDisableCommittersApproval", GoMethod: "ResetMergeRequestsDisableCommittersApproval"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequirePasswordToApprove", GoMethod: "ResetRequirePasswordToApprove"},
			_jsii_.MemberMethod{JsiiMethod: "resetResetApprovalsOnPush", GoMethod: "ResetResetApprovalsOnPush"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_ProjectLevelMrApprovals{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.projectLevelMrApprovals.ProjectLevelMrApprovalsConfig",
		reflect.TypeOf((*ProjectLevelMrApprovalsConfig)(nil)).Elem(),
	)
}