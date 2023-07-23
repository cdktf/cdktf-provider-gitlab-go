package application

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationConfig struct {
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
	// Name of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.2.0/docs/resources/application#name Application#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The URL gitlab should send the user to after authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.2.0/docs/resources/application#redirect_url Application#redirect_url}
	RedirectUrl *string `field:"required" json:"redirectUrl" yaml:"redirectUrl"`
	// Scopes of the application.
	//
	// Use "openid" if you plan to use this as an oidc authentication application. Valid options are: `api`, `read_api`, `read_user`, `read_repository`, `write_repository`, `read_registry`, `write_registry`, `sudo`, `admin_mode`, `openid`, `profile`, `email`.
	// This is only populated when creating a new application. This attribute is not available for imported resources
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.2.0/docs/resources/application#scopes Application#scopes}
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// The application is used where the client secret can be kept confidential.
	//
	// Native mobile apps and Single Page Apps are considered non-confidential. Defaults to true if not supplied
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.2.0/docs/resources/application#confidential Application#confidential}
	Confidential interface{} `field:"optional" json:"confidential" yaml:"confidential"`
}

