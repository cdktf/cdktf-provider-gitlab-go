// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectenvironment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectEnvironmentConfig struct {
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
	// The name of the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_environment#name ProjectEnvironment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID or full path of the project to environment is created for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_environment#project ProjectEnvironment#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The auto stop setting for the environment.
	//
	// Allowed values are `always`, `with_action`. If this is set to `with_action` and `stop_before_destroy` is `true`, the environment will be force-stopped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_environment#auto_stop_setting ProjectEnvironment#auto_stop_setting}
	AutoStopSetting *string `field:"optional" json:"autoStopSetting" yaml:"autoStopSetting"`
	// The cluster agent to associate with this environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_environment#cluster_agent_id ProjectEnvironment#cluster_agent_id}
	ClusterAgentId *float64 `field:"optional" json:"clusterAgentId" yaml:"clusterAgentId"`
	// The description of the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_environment#description ProjectEnvironment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Place to link to for this environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_environment#external_url ProjectEnvironment#external_url}
	ExternalUrl *string `field:"optional" json:"externalUrl" yaml:"externalUrl"`
	// The Flux resource path to associate with this environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_environment#flux_resource_path ProjectEnvironment#flux_resource_path}
	FluxResourcePath *string `field:"optional" json:"fluxResourcePath" yaml:"fluxResourcePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_environment#id ProjectEnvironment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The Kubernetes namespace to associate with this environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_environment#kubernetes_namespace ProjectEnvironment#kubernetes_namespace}
	KubernetesNamespace *string `field:"optional" json:"kubernetesNamespace" yaml:"kubernetesNamespace"`
	// Determines whether the environment is attempted to be stopped before the environment is deleted.
	//
	// If `auto_stop_setting` is set to `with_action`, this will perform a force stop.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_environment#stop_before_destroy ProjectEnvironment#stop_before_destroy}
	StopBeforeDestroy interface{} `field:"optional" json:"stopBeforeDestroy" yaml:"stopBeforeDestroy"`
	// The tier of the new environment. Valid values are `production`, `staging`, `testing`, `development`, `other`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_environment#tier ProjectEnvironment#tier}
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
}

