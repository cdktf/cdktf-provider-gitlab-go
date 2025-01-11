// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package personalaccesstoken


type PersonalAccessTokenRotationConfiguration struct {
	// The duration (in days) the new token should be valid for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.7.1/docs/resources/personal_access_token#expiration_days PersonalAccessToken#expiration_days}
	ExpirationDays *float64 `field:"required" json:"expirationDays" yaml:"expirationDays"`
	// The duration (in days) before the expiration when the token should be rotated.
	//
	// As an example, if set to 7 days, the token will rotate 7 days before the expiration date, but only when `terraform apply` is run in that timeframe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.7.1/docs/resources/personal_access_token#rotate_before_days PersonalAccessToken#rotate_before_days}
	RotateBeforeDays *float64 `field:"required" json:"rotateBeforeDays" yaml:"rotateBeforeDays"`
}

