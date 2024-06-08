// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package project


type ProjectTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.0.1/docs/resources/project#create Project#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.0.1/docs/resources/project#delete Project#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

