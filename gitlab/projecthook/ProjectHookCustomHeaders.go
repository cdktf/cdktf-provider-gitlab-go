// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projecthook


type ProjectHookCustomHeaders struct {
	// Key of the custom header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_hook#key ProjectHook#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Value of the custom header. This value cannot be imported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.3.0/docs/resources/project_hook#value ProjectHook#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

