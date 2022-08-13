// Prebuilt gitlab Provider for Terraform CDK (cdktf)
package gitlab

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectClusterConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The URL to access the Kubernetes API.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project_cluster#kubernetes_api_url ProjectCluster#kubernetes_api_url}
	KubernetesApiUrl *string `field:"required" json:"kubernetesApiUrl" yaml:"kubernetesApiUrl"`
	// The token to authenticate against Kubernetes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project_cluster#kubernetes_token ProjectCluster#kubernetes_token}
	KubernetesToken *string `field:"required" json:"kubernetesToken" yaml:"kubernetesToken"`
	// The name of cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project_cluster#name ProjectCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The id of the project to add the cluster to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project_cluster#project ProjectCluster#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The base domain of the cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project_cluster#domain ProjectCluster#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project_cluster#enabled ProjectCluster#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The associated environment to the cluster. Defaults to `*`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project_cluster#environment_scope ProjectCluster#environment_scope}
	EnvironmentScope *string `field:"optional" json:"environmentScope" yaml:"environmentScope"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project_cluster#id ProjectCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The cluster authorization type. Valid values are `rbac`, `abac`, `unknown_authorization`. Defaults to `rbac`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project_cluster#kubernetes_authorization_type ProjectCluster#kubernetes_authorization_type}
	KubernetesAuthorizationType *string `field:"optional" json:"kubernetesAuthorizationType" yaml:"kubernetesAuthorizationType"`
	// TLS certificate (needed if API is using a self-signed TLS certificate).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project_cluster#kubernetes_ca_cert ProjectCluster#kubernetes_ca_cert}
	KubernetesCaCert *string `field:"optional" json:"kubernetesCaCert" yaml:"kubernetesCaCert"`
	// The unique namespace related to the project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project_cluster#kubernetes_namespace ProjectCluster#kubernetes_namespace}
	KubernetesNamespace *string `field:"optional" json:"kubernetesNamespace" yaml:"kubernetesNamespace"`
	// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project_cluster#managed ProjectCluster#managed}
	Managed interface{} `field:"optional" json:"managed" yaml:"managed"`
	// The ID of the management project for the cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project_cluster#management_project_id ProjectCluster#management_project_id}
	ManagementProjectId *string `field:"optional" json:"managementProjectId" yaml:"managementProjectId"`
}

