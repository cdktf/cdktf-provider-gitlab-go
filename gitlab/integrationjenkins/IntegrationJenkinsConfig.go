// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationjenkins

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationJenkinsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/integration_jenkins#jenkins_url IntegrationJenkins#jenkins_url}
	JenkinsUrl *string `field:"required" json:"jenkinsUrl" yaml:"jenkinsUrl"`
	// ID of the project you want to activate integration on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/integration_jenkins#project IntegrationJenkins#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The URL-friendly project name. Example: `my_project_name`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/integration_jenkins#project_name IntegrationJenkins#project_name}
	ProjectName *string `field:"required" json:"projectName" yaml:"projectName"`
	// Enable SSL verification. Defaults to `true` (enabled).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/integration_jenkins#enable_ssl_verification IntegrationJenkins#enable_ssl_verification}
	EnableSslVerification interface{} `field:"optional" json:"enableSslVerification" yaml:"enableSslVerification"`
	// Enable notifications for merge request events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/integration_jenkins#merge_request_events IntegrationJenkins#merge_request_events}
	MergeRequestEvents interface{} `field:"optional" json:"mergeRequestEvents" yaml:"mergeRequestEvents"`
	// Password for authentication with the Jenkins server, if authentication is required by the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/integration_jenkins#password IntegrationJenkins#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Enable notifications for push events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/integration_jenkins#push_events IntegrationJenkins#push_events}
	PushEvents interface{} `field:"optional" json:"pushEvents" yaml:"pushEvents"`
	// Enable notifications for tag push events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/integration_jenkins#tag_push_events IntegrationJenkins#tag_push_events}
	TagPushEvents interface{} `field:"optional" json:"tagPushEvents" yaml:"tagPushEvents"`
	// Username for authentication with the Jenkins server, if authentication is required by the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.5.0/docs/resources/integration_jenkins#username IntegrationJenkins#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

