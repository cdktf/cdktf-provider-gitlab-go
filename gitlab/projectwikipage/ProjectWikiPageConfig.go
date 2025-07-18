// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectwikipage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectWikiPageConfig struct {
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
	// Content of the wiki page. Must be at least 1 character long.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_wiki_page#content ProjectWikiPage#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The ID or URL-encoded path of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_wiki_page#project ProjectWikiPage#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Title of the wiki page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_wiki_page#title ProjectWikiPage#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// Format of the wiki page (auto-generated if not provided). Valid values are: `markdown`, `rdoc`, `asciidoc`, `org`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/project_wiki_page#format ProjectWikiPage#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
}

