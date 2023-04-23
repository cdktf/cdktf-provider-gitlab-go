package datagitlabuser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabUserConfig struct {
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
	// The public email address of the user.
	//
	// **Note**: before GitLab 14.8 the lookup was based on the users primary email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.11.0/docs/data-sources/user#email DataGitlabUser#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.11.0/docs/data-sources/user#id DataGitlabUser#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of the user's namespace. Requires admin token to access this field. Available since GitLab 14.10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.11.0/docs/data-sources/user#namespace_id DataGitlabUser#namespace_id}
	NamespaceId *float64 `field:"optional" json:"namespaceId" yaml:"namespaceId"`
	// The ID of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.11.0/docs/data-sources/user#user_id DataGitlabUser#user_id}
	UserId *float64 `field:"optional" json:"userId" yaml:"userId"`
	// The username of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.11.0/docs/data-sources/user#username DataGitlabUser#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

