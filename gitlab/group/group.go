package group

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.group.Group",
		reflect.TypeOf((*Group)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "autoDevopsEnabled", GoGetter: "AutoDevopsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "autoDevopsEnabledInput", GoGetter: "AutoDevopsEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultBranchProtection", GoGetter: "DefaultBranchProtection"},
			_jsii_.MemberProperty{JsiiProperty: "defaultBranchProtectionInput", GoGetter: "DefaultBranchProtectionInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "emailsDisabled", GoGetter: "EmailsDisabled"},
			_jsii_.MemberProperty{JsiiProperty: "emailsDisabledInput", GoGetter: "EmailsDisabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lfsEnabled", GoGetter: "LfsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "lfsEnabledInput", GoGetter: "LfsEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "mentionsDisabled", GoGetter: "MentionsDisabled"},
			_jsii_.MemberProperty{JsiiProperty: "mentionsDisabledInput", GoGetter: "MentionsDisabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "parentId", GoGetter: "ParentId"},
			_jsii_.MemberProperty{JsiiProperty: "parentIdInput", GoGetter: "ParentIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "pathInput", GoGetter: "PathInput"},
			_jsii_.MemberProperty{JsiiProperty: "preventForkingOutsideGroup", GoGetter: "PreventForkingOutsideGroup"},
			_jsii_.MemberProperty{JsiiProperty: "preventForkingOutsideGroupInput", GoGetter: "PreventForkingOutsideGroupInput"},
			_jsii_.MemberProperty{JsiiProperty: "projectCreationLevel", GoGetter: "ProjectCreationLevel"},
			_jsii_.MemberProperty{JsiiProperty: "projectCreationLevelInput", GoGetter: "ProjectCreationLevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "requestAccessEnabled", GoGetter: "RequestAccessEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "requestAccessEnabledInput", GoGetter: "RequestAccessEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "requireTwoFactorAuthentication", GoGetter: "RequireTwoFactorAuthentication"},
			_jsii_.MemberProperty{JsiiProperty: "requireTwoFactorAuthenticationInput", GoGetter: "RequireTwoFactorAuthenticationInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoDevopsEnabled", GoMethod: "ResetAutoDevopsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultBranchProtection", GoMethod: "ResetDefaultBranchProtection"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailsDisabled", GoMethod: "ResetEmailsDisabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLfsEnabled", GoMethod: "ResetLfsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetMentionsDisabled", GoMethod: "ResetMentionsDisabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetParentId", GoMethod: "ResetParentId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreventForkingOutsideGroup", GoMethod: "ResetPreventForkingOutsideGroup"},
			_jsii_.MemberMethod{JsiiMethod: "resetProjectCreationLevel", GoMethod: "ResetProjectCreationLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestAccessEnabled", GoMethod: "ResetRequestAccessEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequireTwoFactorAuthentication", GoMethod: "ResetRequireTwoFactorAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "resetShareWithGroupLock", GoMethod: "ResetShareWithGroupLock"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubgroupCreationLevel", GoMethod: "ResetSubgroupCreationLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetTwoFactorGracePeriod", GoMethod: "ResetTwoFactorGracePeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resetVisibilityLevel", GoMethod: "ResetVisibilityLevel"},
			_jsii_.MemberProperty{JsiiProperty: "runnersToken", GoGetter: "RunnersToken"},
			_jsii_.MemberProperty{JsiiProperty: "shareWithGroupLock", GoGetter: "ShareWithGroupLock"},
			_jsii_.MemberProperty{JsiiProperty: "shareWithGroupLockInput", GoGetter: "ShareWithGroupLockInput"},
			_jsii_.MemberProperty{JsiiProperty: "subgroupCreationLevel", GoGetter: "SubgroupCreationLevel"},
			_jsii_.MemberProperty{JsiiProperty: "subgroupCreationLevelInput", GoGetter: "SubgroupCreationLevelInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "twoFactorGracePeriod", GoGetter: "TwoFactorGracePeriod"},
			_jsii_.MemberProperty{JsiiProperty: "twoFactorGracePeriodInput", GoGetter: "TwoFactorGracePeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "visibilityLevel", GoGetter: "VisibilityLevel"},
			_jsii_.MemberProperty{JsiiProperty: "visibilityLevelInput", GoGetter: "VisibilityLevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "webUrl", GoGetter: "WebUrl"},
		},
		func() interface{} {
			j := jsiiProxy_Group{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.group.GroupConfig",
		reflect.TypeOf((*GroupConfig)(nil)).Elem(),
	)
}
