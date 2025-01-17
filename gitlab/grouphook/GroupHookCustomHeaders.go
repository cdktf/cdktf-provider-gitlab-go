// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grouphook


type GroupHookCustomHeaders struct {
	// Key of the custom header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.8.0/docs/resources/group_hook#key GroupHook#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Value of the custom header. This value cannot be imported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.8.0/docs/resources/group_hook#value GroupHook#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

