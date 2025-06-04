// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupdependencyproxy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupDependencyProxyConfig struct {
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
	// The ID or URL-encoded path of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/resources/group_dependency_proxy#group GroupDependencyProxy#group}
	Group *string `field:"required" json:"group" yaml:"group"`
	// Indicates whether the proxy is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/resources/group_dependency_proxy#enabled GroupDependencyProxy#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Identity credential used to authenticate with Docker Hub when pulling images.
	//
	// Can be a username (for password or personal access token (PAT)) or organization name (for organization access token (OAT)).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/resources/group_dependency_proxy#identity GroupDependencyProxy#identity}
	Identity *string `field:"optional" json:"identity" yaml:"identity"`
	// Secret credential used to authenticate with Docker Hub when pulling images.
	//
	// Can be a password, personal access token (PAT), or organization access token (OAT). Cannot be imported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.0.0/docs/resources/group_dependency_proxy#secret GroupDependencyProxy#secret}
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
}

