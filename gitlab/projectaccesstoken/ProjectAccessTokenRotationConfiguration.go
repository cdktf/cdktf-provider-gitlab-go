// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectaccesstoken


type ProjectAccessTokenRotationConfiguration struct {
	// The duration (in days) the new token should be valid for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/project_access_token#expiration_days ProjectAccessToken#expiration_days}
	ExpirationDays *float64 `field:"required" json:"expirationDays" yaml:"expirationDays"`
	// The duration (in days) before the expiration when the token should be rotated.
	//
	// As an example, if set to 7 days, the token will rotate 7 days before the expiration date, but only when `terraform apply` is run in that timeframe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.6.1/docs/resources/project_access_token#rotate_before_days ProjectAccessToken#rotate_before_days}
	RotateBeforeDays *float64 `field:"required" json:"rotateBeforeDays" yaml:"rotateBeforeDays"`
}

