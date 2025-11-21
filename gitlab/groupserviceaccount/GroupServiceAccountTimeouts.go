// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupserviceaccount


type GroupServiceAccountTimeouts struct {
	// How long to wait for the service account to be fully deleted. Defaults to 10 minutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.6.0/docs/resources/group_service_account#delete GroupServiceAccount#delete}
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

