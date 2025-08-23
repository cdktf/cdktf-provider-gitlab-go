// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabuser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabUserConfig struct {
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
	// The public email address of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/data-sources/user#email DataGitlabUser#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
	// (Experimental) If true, returns only an exact match.
	//
	// Otherwise, fuzzy matching might return the closest result. If no exact match is available, the data source returns an error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/data-sources/user#email_exact_match DataGitlabUser#email_exact_match}
	EmailExactMatch interface{} `field:"optional" json:"emailExactMatch" yaml:"emailExactMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/data-sources/user#id DataGitlabUser#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of the user's namespace. Requires admin token to access this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/data-sources/user#namespace_id DataGitlabUser#namespace_id}
	NamespaceId *float64 `field:"optional" json:"namespaceId" yaml:"namespaceId"`
	// The ID of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/data-sources/user#user_id DataGitlabUser#user_id}
	UserId *float64 `field:"optional" json:"userId" yaml:"userId"`
	// The username of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/data-sources/user#username DataGitlabUser#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

