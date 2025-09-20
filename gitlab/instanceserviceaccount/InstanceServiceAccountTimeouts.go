// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package instanceserviceaccount


type InstanceServiceAccountTimeouts struct {
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.4.0/docs/resources/instance_service_account#delete InstanceServiceAccount#delete}
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

