package datagitlabgroupmembership

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabGroupMembershipConfig struct {
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
	// Only return members with the desired access level. Acceptable values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/group_membership#access_level DataGitlabGroupMembership#access_level}
	AccessLevel *string `field:"optional" json:"accessLevel" yaml:"accessLevel"`
	// The full path of the group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/group_membership#full_path DataGitlabGroupMembership#full_path}
	FullPath *string `field:"optional" json:"fullPath" yaml:"fullPath"`
	// The ID of the group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/group_membership#group_id DataGitlabGroupMembership#group_id}
	GroupId *float64 `field:"optional" json:"groupId" yaml:"groupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/group_membership#id DataGitlabGroupMembership#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

