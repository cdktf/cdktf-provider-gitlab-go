// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectcluster

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.projectCluster.ProjectCluster",
		reflect.TypeOf((*ProjectCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clusterType", GoGetter: "ClusterType"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "domain", GoGetter: "Domain"},
			_jsii_.MemberProperty{JsiiProperty: "domainInput", GoGetter: "DomainInput"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "environmentScope", GoGetter: "EnvironmentScope"},
			_jsii_.MemberProperty{JsiiProperty: "environmentScopeInput", GoGetter: "EnvironmentScopeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "kubernetesApiUrl", GoGetter: "KubernetesApiUrl"},
			_jsii_.MemberProperty{JsiiProperty: "kubernetesApiUrlInput", GoGetter: "KubernetesApiUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "kubernetesAuthorizationType", GoGetter: "KubernetesAuthorizationType"},
			_jsii_.MemberProperty{JsiiProperty: "kubernetesAuthorizationTypeInput", GoGetter: "KubernetesAuthorizationTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "kubernetesCaCert", GoGetter: "KubernetesCaCert"},
			_jsii_.MemberProperty{JsiiProperty: "kubernetesCaCertInput", GoGetter: "KubernetesCaCertInput"},
			_jsii_.MemberProperty{JsiiProperty: "kubernetesNamespace", GoGetter: "KubernetesNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "kubernetesNamespaceInput", GoGetter: "KubernetesNamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "kubernetesToken", GoGetter: "KubernetesToken"},
			_jsii_.MemberProperty{JsiiProperty: "kubernetesTokenInput", GoGetter: "KubernetesTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "managed", GoGetter: "Managed"},
			_jsii_.MemberProperty{JsiiProperty: "managedInput", GoGetter: "ManagedInput"},
			_jsii_.MemberProperty{JsiiProperty: "managementProjectId", GoGetter: "ManagementProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "managementProjectIdInput", GoGetter: "ManagementProjectIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "platformType", GoGetter: "PlatformType"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "projectInput", GoGetter: "ProjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "providerType", GoGetter: "ProviderType"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDomain", GoMethod: "ResetDomain"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnvironmentScope", GoMethod: "ResetEnvironmentScope"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetKubernetesAuthorizationType", GoMethod: "ResetKubernetesAuthorizationType"},
			_jsii_.MemberMethod{JsiiMethod: "resetKubernetesCaCert", GoMethod: "ResetKubernetesCaCert"},
			_jsii_.MemberMethod{JsiiMethod: "resetKubernetesNamespace", GoMethod: "ResetKubernetesNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetManaged", GoMethod: "ResetManaged"},
			_jsii_.MemberMethod{JsiiMethod: "resetManagementProjectId", GoMethod: "ResetManagementProjectId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_ProjectCluster{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.projectCluster.ProjectClusterConfig",
		reflect.TypeOf((*ProjectClusterConfig)(nil)).Elem(),
	)
}
