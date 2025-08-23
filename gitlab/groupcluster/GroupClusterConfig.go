// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupClusterConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The id of the group to add the cluster to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/group_cluster#group GroupCluster#group}
	Group *string `field:"required" json:"group" yaml:"group"`
	// The URL to access the Kubernetes API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/group_cluster#kubernetes_api_url GroupCluster#kubernetes_api_url}
	KubernetesApiUrl *string `field:"required" json:"kubernetesApiUrl" yaml:"kubernetesApiUrl"`
	// The token to authenticate against Kubernetes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/group_cluster#kubernetes_token GroupCluster#kubernetes_token}
	KubernetesToken *string `field:"required" json:"kubernetesToken" yaml:"kubernetesToken"`
	// The name of cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/group_cluster#name GroupCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The base domain of the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/group_cluster#domain GroupCluster#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/group_cluster#enabled GroupCluster#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The associated environment to the cluster. Defaults to `*`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/group_cluster#environment_scope GroupCluster#environment_scope}
	EnvironmentScope *string `field:"optional" json:"environmentScope" yaml:"environmentScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/group_cluster#id GroupCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The cluster authorization type. Valid values are `rbac`, `abac`, `unknown_authorization`. Defaults to `rbac`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/group_cluster#kubernetes_authorization_type GroupCluster#kubernetes_authorization_type}
	KubernetesAuthorizationType *string `field:"optional" json:"kubernetesAuthorizationType" yaml:"kubernetesAuthorizationType"`
	// TLS certificate (needed if API is using a self-signed TLS certificate).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/group_cluster#kubernetes_ca_cert GroupCluster#kubernetes_ca_cert}
	KubernetesCaCert *string `field:"optional" json:"kubernetesCaCert" yaml:"kubernetesCaCert"`
	// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/group_cluster#managed GroupCluster#managed}
	Managed interface{} `field:"optional" json:"managed" yaml:"managed"`
	// The ID of the management project for the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/group_cluster#management_project_id GroupCluster#management_project_id}
	ManagementProjectId *string `field:"optional" json:"managementProjectId" yaml:"managementProjectId"`
}

