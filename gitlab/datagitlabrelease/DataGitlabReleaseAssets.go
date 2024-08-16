// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabrelease


type DataGitlabReleaseAssets struct {
	// links block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.3.0/docs/data-sources/release#links DataGitlabRelease#links}
	Links interface{} `field:"optional" json:"links" yaml:"links"`
	// sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.3.0/docs/data-sources/release#sources DataGitlabRelease#sources}
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
}

