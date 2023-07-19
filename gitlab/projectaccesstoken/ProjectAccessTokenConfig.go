package projectaccesstoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectAccessTokenConfig struct {
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
	// A name to describe the project access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.1/docs/resources/project_access_token#name ProjectAccessToken#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The id of the project to add the project access token to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.1/docs/resources/project_access_token#project ProjectAccessToken#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Valid values: `api`, `read_api`, `read_repository`, `write_repository`, `read_registry`, `write_registry`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.1/docs/resources/project_access_token#scopes ProjectAccessToken#scopes}
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// The access level for the project access token.
	//
	// Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`. Default is `maintainer`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.1/docs/resources/project_access_token#access_level ProjectAccessToken#access_level}
	AccessLevel *string `field:"optional" json:"accessLevel" yaml:"accessLevel"`
	// Time the token will expire it, YYYY-MM-DD format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.1/docs/resources/project_access_token#expires_at ProjectAccessToken#expires_at}
	ExpiresAt *string `field:"optional" json:"expiresAt" yaml:"expiresAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.1/docs/resources/project_access_token#id ProjectAccessToken#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

