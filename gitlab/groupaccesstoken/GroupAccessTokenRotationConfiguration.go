// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupaccesstoken


type GroupAccessTokenRotationConfiguration struct {
	// The duration (in days) the new token should be valid for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/resources/group_access_token#expiration_days GroupAccessToken#expiration_days}
	ExpirationDays *float64 `field:"required" json:"expirationDays" yaml:"expirationDays"`
	// The duration (in days) before the expiration when the token should be rotated.
	//
	// As an example, if set to 7 days, the token will rotate 7 days before the expiration date, but only when `terraform apply` is run in that timeframe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.9.0/docs/resources/group_access_token#rotate_before_days GroupAccessToken#rotate_before_days}
	RotateBeforeDays *float64 `field:"required" json:"rotateBeforeDays" yaml:"rotateBeforeDays"`
}

