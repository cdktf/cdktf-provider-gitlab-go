// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationappearance

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.applicationAppearance.ApplicationAppearance",
		reflect.TypeOf((*ApplicationAppearance)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "emailHeaderAndFooterEnabled", GoGetter: "EmailHeaderAndFooterEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "emailHeaderAndFooterEnabledInput", GoGetter: "EmailHeaderAndFooterEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "footerMessage", GoGetter: "FooterMessage"},
			_jsii_.MemberProperty{JsiiProperty: "footerMessageInput", GoGetter: "FooterMessageInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "headerMessage", GoGetter: "HeaderMessage"},
			_jsii_.MemberProperty{JsiiProperty: "headerMessageInput", GoGetter: "HeaderMessageInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keepSettingsOnDestroy", GoGetter: "KeepSettingsOnDestroy"},
			_jsii_.MemberProperty{JsiiProperty: "keepSettingsOnDestroyInput", GoGetter: "KeepSettingsOnDestroyInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "memberGuidelines", GoGetter: "MemberGuidelines"},
			_jsii_.MemberProperty{JsiiProperty: "memberGuidelinesInput", GoGetter: "MemberGuidelinesInput"},
			_jsii_.MemberProperty{JsiiProperty: "messageBackgroundColor", GoGetter: "MessageBackgroundColor"},
			_jsii_.MemberProperty{JsiiProperty: "messageBackgroundColorInput", GoGetter: "MessageBackgroundColorInput"},
			_jsii_.MemberProperty{JsiiProperty: "messageFontColor", GoGetter: "MessageFontColor"},
			_jsii_.MemberProperty{JsiiProperty: "messageFontColorInput", GoGetter: "MessageFontColorInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "newProjectGuidelines", GoGetter: "NewProjectGuidelines"},
			_jsii_.MemberProperty{JsiiProperty: "newProjectGuidelinesInput", GoGetter: "NewProjectGuidelinesInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "profileImageGuidelines", GoGetter: "ProfileImageGuidelines"},
			_jsii_.MemberProperty{JsiiProperty: "profileImageGuidelinesInput", GoGetter: "ProfileImageGuidelinesInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "pwaDescription", GoGetter: "PwaDescription"},
			_jsii_.MemberProperty{JsiiProperty: "pwaDescriptionInput", GoGetter: "PwaDescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "pwaName", GoGetter: "PwaName"},
			_jsii_.MemberProperty{JsiiProperty: "pwaNameInput", GoGetter: "PwaNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "pwaShortName", GoGetter: "PwaShortName"},
			_jsii_.MemberProperty{JsiiProperty: "pwaShortNameInput", GoGetter: "PwaShortNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailHeaderAndFooterEnabled", GoMethod: "ResetEmailHeaderAndFooterEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetFooterMessage", GoMethod: "ResetFooterMessage"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderMessage", GoMethod: "ResetHeaderMessage"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeepSettingsOnDestroy", GoMethod: "ResetKeepSettingsOnDestroy"},
			_jsii_.MemberMethod{JsiiMethod: "resetMemberGuidelines", GoMethod: "ResetMemberGuidelines"},
			_jsii_.MemberMethod{JsiiMethod: "resetMessageBackgroundColor", GoMethod: "ResetMessageBackgroundColor"},
			_jsii_.MemberMethod{JsiiMethod: "resetMessageFontColor", GoMethod: "ResetMessageFontColor"},
			_jsii_.MemberMethod{JsiiMethod: "resetNewProjectGuidelines", GoMethod: "ResetNewProjectGuidelines"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetProfileImageGuidelines", GoMethod: "ResetProfileImageGuidelines"},
			_jsii_.MemberMethod{JsiiMethod: "resetPwaDescription", GoMethod: "ResetPwaDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetPwaName", GoMethod: "ResetPwaName"},
			_jsii_.MemberMethod{JsiiMethod: "resetPwaShortName", GoMethod: "ResetPwaShortName"},
			_jsii_.MemberMethod{JsiiMethod: "resetTitle", GoMethod: "ResetTitle"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "title", GoGetter: "Title"},
			_jsii_.MemberProperty{JsiiProperty: "titleInput", GoGetter: "TitleInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_ApplicationAppearance{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.applicationAppearance.ApplicationAppearanceConfig",
		reflect.TypeOf((*ApplicationAppearanceConfig)(nil)).Elem(),
	)
}
