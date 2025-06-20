// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabuser

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.dataGitlabUser.DataGitlabUser",
		reflect.TypeOf((*DataGitlabUser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "avatarUrl", GoGetter: "AvatarUrl"},
			_jsii_.MemberProperty{JsiiProperty: "bio", GoGetter: "Bio"},
			_jsii_.MemberProperty{JsiiProperty: "canCreateGroup", GoGetter: "CanCreateGroup"},
			_jsii_.MemberProperty{JsiiProperty: "canCreateProject", GoGetter: "CanCreateProject"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "colorSchemeId", GoGetter: "ColorSchemeId"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "currentSignInAt", GoGetter: "CurrentSignInAt"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "email", GoGetter: "Email"},
			_jsii_.MemberProperty{JsiiProperty: "emailExactMatch", GoGetter: "EmailExactMatch"},
			_jsii_.MemberProperty{JsiiProperty: "emailExactMatchInput", GoGetter: "EmailExactMatchInput"},
			_jsii_.MemberProperty{JsiiProperty: "emailInput", GoGetter: "EmailInput"},
			_jsii_.MemberProperty{JsiiProperty: "external", GoGetter: "External"},
			_jsii_.MemberProperty{JsiiProperty: "externUid", GoGetter: "ExternUid"},
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
			_jsii_.MemberProperty{JsiiProperty: "isAdmin", GoGetter: "IsAdmin"},
			_jsii_.MemberProperty{JsiiProperty: "isBot", GoGetter: "IsBot"},
			_jsii_.MemberProperty{JsiiProperty: "lastSignInAt", GoGetter: "LastSignInAt"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "linkedin", GoGetter: "Linkedin"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceId", GoGetter: "NamespaceId"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceIdInput", GoGetter: "NamespaceIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "note", GoGetter: "Note"},
			_jsii_.MemberProperty{JsiiProperty: "organization", GoGetter: "Organization"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "projectsLimit", GoGetter: "ProjectsLimit"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmail", GoMethod: "ResetEmail"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailExactMatch", GoMethod: "ResetEmailExactMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespaceId", GoMethod: "ResetNamespaceId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserId", GoMethod: "ResetUserId"},
			_jsii_.MemberMethod{JsiiMethod: "resetUsername", GoMethod: "ResetUsername"},
			_jsii_.MemberProperty{JsiiProperty: "skype", GoGetter: "Skype"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "themeId", GoGetter: "ThemeId"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "twitter", GoGetter: "Twitter"},
			_jsii_.MemberProperty{JsiiProperty: "twoFactorEnabled", GoGetter: "TwoFactorEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "userId", GoGetter: "UserId"},
			_jsii_.MemberProperty{JsiiProperty: "userIdInput", GoGetter: "UserIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "username", GoGetter: "Username"},
			_jsii_.MemberProperty{JsiiProperty: "usernameInput", GoGetter: "UsernameInput"},
			_jsii_.MemberProperty{JsiiProperty: "userProvider", GoGetter: "UserProvider"},
			_jsii_.MemberProperty{JsiiProperty: "websiteUrl", GoGetter: "WebsiteUrl"},
		},
		func() interface{} {
			j := jsiiProxy_DataGitlabUser{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.dataGitlabUser.DataGitlabUserConfig",
		reflect.TypeOf((*DataGitlabUserConfig)(nil)).Elem(),
	)
}
