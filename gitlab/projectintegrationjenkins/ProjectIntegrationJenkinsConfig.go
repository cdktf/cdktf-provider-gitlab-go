// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectintegrationjenkins

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectIntegrationJenkinsConfig struct {
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
	// Jenkins URL like `http://jenkins.example.com`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_integration_jenkins#jenkins_url ProjectIntegrationJenkins#jenkins_url}
	JenkinsUrl *string `field:"required" json:"jenkinsUrl" yaml:"jenkinsUrl"`
	// ID of the project you want to activate integration on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_integration_jenkins#project ProjectIntegrationJenkins#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The URL-friendly project name. Example: `my_project_name`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_integration_jenkins#project_name ProjectIntegrationJenkins#project_name}
	ProjectName *string `field:"required" json:"projectName" yaml:"projectName"`
	// Enable SSL verification. Defaults to `true` (enabled).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_integration_jenkins#enable_ssl_verification ProjectIntegrationJenkins#enable_ssl_verification}
	EnableSslVerification interface{} `field:"optional" json:"enableSslVerification" yaml:"enableSslVerification"`
	// Enable notifications for merge request events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_integration_jenkins#merge_request_events ProjectIntegrationJenkins#merge_request_events}
	MergeRequestEvents interface{} `field:"optional" json:"mergeRequestEvents" yaml:"mergeRequestEvents"`
	// Password for authentication with the Jenkins server, if authentication is required by the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_integration_jenkins#password ProjectIntegrationJenkins#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Enable notifications for push events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_integration_jenkins#push_events ProjectIntegrationJenkins#push_events}
	PushEvents interface{} `field:"optional" json:"pushEvents" yaml:"pushEvents"`
	// Enable notifications for tag push events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_integration_jenkins#tag_push_events ProjectIntegrationJenkins#tag_push_events}
	TagPushEvents interface{} `field:"optional" json:"tagPushEvents" yaml:"tagPushEvents"`
	// Username for authentication with the Jenkins server, if authentication is required by the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_integration_jenkins#username ProjectIntegrationJenkins#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

