// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.provider.GitlabProvider",
		reflect.TypeOf((*GitlabProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "baseUrl", GoGetter: "BaseUrl"},
			_jsii_.MemberProperty{JsiiProperty: "baseUrlInput", GoGetter: "BaseUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "cacertFile", GoGetter: "CacertFile"},
			_jsii_.MemberProperty{JsiiProperty: "cacertFileInput", GoGetter: "CacertFileInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clientCert", GoGetter: "ClientCert"},
			_jsii_.MemberProperty{JsiiProperty: "clientCertInput", GoGetter: "ClientCertInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientKey", GoGetter: "ClientKey"},
			_jsii_.MemberProperty{JsiiProperty: "clientKeyInput", GoGetter: "ClientKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "configFile", GoGetter: "ConfigFile"},
			_jsii_.MemberProperty{JsiiProperty: "configFileInput", GoGetter: "ConfigFileInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "context", GoGetter: "Context"},
			_jsii_.MemberProperty{JsiiProperty: "contextInput", GoGetter: "ContextInput"},
			_jsii_.MemberProperty{JsiiProperty: "earlyAuthCheck", GoGetter: "EarlyAuthCheck"},
			_jsii_.MemberProperty{JsiiProperty: "earlyAuthCheckInput", GoGetter: "EarlyAuthCheckInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableAutoCiSupport", GoGetter: "EnableAutoCiSupport"},
			_jsii_.MemberProperty{JsiiProperty: "enableAutoCiSupportInput", GoGetter: "EnableAutoCiSupportInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "headers", GoGetter: "Headers"},
			_jsii_.MemberProperty{JsiiProperty: "headersInput", GoGetter: "HeadersInput"},
			_jsii_.MemberProperty{JsiiProperty: "insecure", GoGetter: "Insecure"},
			_jsii_.MemberProperty{JsiiProperty: "insecureInput", GoGetter: "InsecureInput"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetBaseUrl", GoMethod: "ResetBaseUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetCacertFile", GoMethod: "ResetCacertFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientCert", GoMethod: "ResetClientCert"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientKey", GoMethod: "ResetClientKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfigFile", GoMethod: "ResetConfigFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetContext", GoMethod: "ResetContext"},
			_jsii_.MemberMethod{JsiiMethod: "resetEarlyAuthCheck", GoMethod: "ResetEarlyAuthCheck"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableAutoCiSupport", GoMethod: "ResetEnableAutoCiSupport"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaders", GoMethod: "ResetHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetInsecure", GoMethod: "ResetInsecure"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetries", GoMethod: "ResetRetries"},
			_jsii_.MemberMethod{JsiiMethod: "resetToken", GoMethod: "ResetToken"},
			_jsii_.MemberProperty{JsiiProperty: "retries", GoGetter: "Retries"},
			_jsii_.MemberProperty{JsiiProperty: "retriesInput", GoGetter: "RetriesInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "token", GoGetter: "Token"},
			_jsii_.MemberProperty{JsiiProperty: "tokenInput", GoGetter: "TokenInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_GitlabProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.provider.GitlabProviderConfig",
		reflect.TypeOf((*GitlabProviderConfig)(nil)).Elem(),
	)
}
