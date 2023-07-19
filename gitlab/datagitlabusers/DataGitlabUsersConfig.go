package datagitlabusers

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabUsersConfig struct {
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
	// Filter users that are active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.1/docs/data-sources/users#active DataGitlabUsers#active}
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// Filter users that are blocked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.1/docs/data-sources/users#blocked DataGitlabUsers#blocked}
	Blocked interface{} `field:"optional" json:"blocked" yaml:"blocked"`
	// Search for users created after a specific date. (Requires administrator privileges).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.1/docs/data-sources/users#created_after DataGitlabUsers#created_after}
	CreatedAfter *string `field:"optional" json:"createdAfter" yaml:"createdAfter"`
	// Search for users created before a specific date. (Requires administrator privileges).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.1/docs/data-sources/users#created_before DataGitlabUsers#created_before}
	CreatedBefore *string `field:"optional" json:"createdBefore" yaml:"createdBefore"`
	// Lookup users by external provider. (Requires administrator privileges).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.1/docs/data-sources/users#extern_provider DataGitlabUsers#extern_provider}
	ExternProvider *string `field:"optional" json:"externProvider" yaml:"externProvider"`
	// Lookup users by external UID. (Requires administrator privileges).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.1/docs/data-sources/users#extern_uid DataGitlabUsers#extern_uid}
	ExternUid *string `field:"optional" json:"externUid" yaml:"externUid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.1/docs/data-sources/users#id DataGitlabUsers#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Order the users' list by `id`, `name`, `username`, `created_at` or `updated_at`. (Requires administrator privileges).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.1/docs/data-sources/users#order_by DataGitlabUsers#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// Search users by username, name or email.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.1/docs/data-sources/users#search DataGitlabUsers#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
	// Sort users' list in asc or desc order. (Requires administrator privileges).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.1/docs/data-sources/users#sort DataGitlabUsers#sort}
	Sort *string `field:"optional" json:"sort" yaml:"sort"`
}

