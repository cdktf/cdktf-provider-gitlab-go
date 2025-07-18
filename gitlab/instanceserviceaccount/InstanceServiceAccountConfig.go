// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package instanceserviceaccount

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type InstanceServiceAccountConfig struct {
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
	// The email of the user account. If not set, generates a no-reply email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/instance_service_account#email InstanceServiceAccount#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
	// The name of the user. If not set, uses Service account user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/instance_service_account#name InstanceServiceAccount#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/instance_service_account#timeouts InstanceServiceAccount#timeouts}.
	Timeouts *InstanceServiceAccountTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The username of the user account. If not set, generates a name prepended with service_account_.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/instance_service_account#username InstanceServiceAccount#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

