// Prebuilt gitlab Provider for Terraform CDK (cdktf)
package gitlab

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupLdapLinkConfig struct {
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
	// The CN of the LDAP group to link with.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group_ldap_link#cn GroupLdapLink#cn}
	Cn *string `field:"required" json:"cn" yaml:"cn"`
	// The id of the GitLab group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group_ldap_link#group_id GroupLdapLink#group_id}
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// The name of the LDAP provider as stored in the GitLab database.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group_ldap_link#ldap_provider GroupLdapLink#ldap_provider}
	LdapProvider *string `field:"required" json:"ldapProvider" yaml:"ldapProvider"`
	// Minimum access level for members of the LDAP group.
	//
	// Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group_ldap_link#access_level GroupLdapLink#access_level}
	AccessLevel *string `field:"optional" json:"accessLevel" yaml:"accessLevel"`
	// If true, then delete and replace an existing LDAP link if one exists.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group_ldap_link#force GroupLdapLink#force}
	Force interface{} `field:"optional" json:"force" yaml:"force"`
	// Minimum access level for members of the LDAP group.
	//
	// Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group_ldap_link#group_access GroupLdapLink#group_access}
	GroupAccess *string `field:"optional" json:"groupAccess" yaml:"groupAccess"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/group_ldap_link#id GroupLdapLink#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

