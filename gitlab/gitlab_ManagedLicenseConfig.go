// Prebuilt gitlab Provider for Terraform CDK (cdktf)
package gitlab

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedLicenseConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The approval status of the license.
	//
	// Valid values are: `approved`, `blacklisted`, `allowed`, `denied`. "approved" and "blacklisted"
	// have been deprecated in favor of "allowed" and "denied"; use "allowed" and "denied" for GitLab versions 15.0 and higher.
	// Prior to version 15.0 and after 14.6, the values are equivalent.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/managed_license#approval_status ManagedLicense#approval_status}
	ApprovalStatus *string `field:"required" json:"approvalStatus" yaml:"approvalStatus"`
	// The name of the managed license (I.e., 'Apache License 2.0' or 'MIT license').
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/managed_license#name ManagedLicense#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the project under which the managed license will be created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/managed_license#project ManagedLicense#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/managed_license#id ManagedLicense#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}
