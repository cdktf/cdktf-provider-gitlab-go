package usersshkey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserSshkeyConfig struct {
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
	// The ssh key.
	//
	// The SSH key `comment` (trailing part) is optional and ignored for diffing, because GitLab overrides it with the username and GitLab hostname.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/user_sshkey#key UserSshkey#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The title of the ssh key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/user_sshkey#title UserSshkey#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// The ID of the user to add the ssh key to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/user_sshkey#user_id UserSshkey#user_id}
	UserId *float64 `field:"required" json:"userId" yaml:"userId"`
	// The expiration date of the SSH key in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/user_sshkey#expires_at UserSshkey#expires_at}
	ExpiresAt *string `field:"optional" json:"expiresAt" yaml:"expiresAt"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/user_sshkey#id UserSshkey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}
