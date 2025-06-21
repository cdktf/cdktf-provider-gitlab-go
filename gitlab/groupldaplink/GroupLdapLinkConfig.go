// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupldaplink

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupLdapLinkConfig struct {
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
	// The ID or URL-encoded path of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_ldap_link#group GroupLdapLink#group}
	Group *string `field:"required" json:"group" yaml:"group"`
	// Minimum access level for members of the LDAP group.
	//
	// Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_ldap_link#group_access GroupLdapLink#group_access}
	GroupAccess *string `field:"required" json:"groupAccess" yaml:"groupAccess"`
	// The name of the LDAP provider as stored in the GitLab database.
	//
	// Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/administration/raketasks/ldap/#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_ldap_link#ldap_provider GroupLdapLink#ldap_provider}
	LdapProvider *string `field:"required" json:"ldapProvider" yaml:"ldapProvider"`
	// The CN of the LDAP group to link with. Required if `filter` is not provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_ldap_link#cn GroupLdapLink#cn}
	Cn *string `field:"optional" json:"cn" yaml:"cn"`
	// The LDAP filter for the group. Required if `cn` is not provided. Requires GitLab Premium or above.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_ldap_link#filter GroupLdapLink#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// If true, then delete and replace an existing LDAP link if one exists.
	//
	// Will also remove an LDAP link if the parent group is not found.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_ldap_link#force GroupLdapLink#force}
	Force interface{} `field:"optional" json:"force" yaml:"force"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_ldap_link#id GroupLdapLink#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of a custom member role.
	//
	// Only available for Ultimate instances. When using a custom role, the `group_access` must match the base role used to create the custom role. To remove a custom role and revert to a base role, set this value to `0`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/group_ldap_link#member_role_id GroupLdapLink#member_role_id}
	MemberRoleId *float64 `field:"optional" json:"memberRoleId" yaml:"memberRoleId"`
}

