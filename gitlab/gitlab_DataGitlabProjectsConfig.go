// Prebuilt gitlab Provider for Terraform CDK (cdktf)
package gitlab

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabProjectsConfig struct {
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
	// Limit by archived status.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/projects#archived DataGitlabProjects#archived}
	Archived interface{} `field:"optional" json:"archived" yaml:"archived"`
	// The ID of the group owned by the authenticated user to look projects for within.
	//
	// Cannot be used with `min_access_level`, `with_programming_language` or `statistics`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/projects#group_id DataGitlabProjects#group_id}
	GroupId *float64 `field:"optional" json:"groupId" yaml:"groupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/projects#id DataGitlabProjects#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Include projects in subgroups of this group. Default is `false`. Needs `group_id`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/projects#include_subgroups DataGitlabProjects#include_subgroups}
	IncludeSubgroups interface{} `field:"optional" json:"includeSubgroups" yaml:"includeSubgroups"`
	// The maximum number of project results pages that may be queried.
	//
	// Prevents overloading your Gitlab instance in case of a misconfiguration.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/projects#max_queryable_pages DataGitlabProjects#max_queryable_pages}
	MaxQueryablePages *float64 `field:"optional" json:"maxQueryablePages" yaml:"maxQueryablePages"`
	// Limit by projects that the current user is a member of.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/projects#membership DataGitlabProjects#membership}
	Membership interface{} `field:"optional" json:"membership" yaml:"membership"`
	// Limit to projects where current user has at least this access level, refer to the [official documentation](https://docs.gitlab.com/ee/api/members.html) for values. Cannot be used with `group_id`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/projects#min_access_level DataGitlabProjects#min_access_level}
	MinAccessLevel *float64 `field:"optional" json:"minAccessLevel" yaml:"minAccessLevel"`
	// Return projects ordered by `id`, `name`, `path`, `created_at`, `updated_at`, or `last_activity_at` fields. Default is `created_at`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/projects#order_by DataGitlabProjects#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// Limit by projects owned by the current user.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/projects#owned DataGitlabProjects#owned}
	Owned interface{} `field:"optional" json:"owned" yaml:"owned"`
	// The first page to begin the query on.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/projects#page DataGitlabProjects#page}
	Page *float64 `field:"optional" json:"page" yaml:"page"`
	// The number of results to return per page.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/projects#per_page DataGitlabProjects#per_page}
	PerPage *float64 `field:"optional" json:"perPage" yaml:"perPage"`
	// Return list of authorized projects matching the search criteria.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/projects#search DataGitlabProjects#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
	// Return only the ID, URL, name, and path of each project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/projects#simple DataGitlabProjects#simple}
	Simple interface{} `field:"optional" json:"simple" yaml:"simple"`
	// Return projects sorted in `asc` or `desc` order. Default is `desc`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/projects#sort DataGitlabProjects#sort}
	Sort *string `field:"optional" json:"sort" yaml:"sort"`
	// Limit by projects starred by the current user.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/projects#starred DataGitlabProjects#starred}
	Starred interface{} `field:"optional" json:"starred" yaml:"starred"`
	// Include project statistics. Cannot be used with `group_id`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/projects#statistics DataGitlabProjects#statistics}
	Statistics interface{} `field:"optional" json:"statistics" yaml:"statistics"`
	// Limit by visibility `public`, `internal`, or `private`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/projects#visibility DataGitlabProjects#visibility}
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
	// Include custom attributes in response _(admins only)_.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/projects#with_custom_attributes DataGitlabProjects#with_custom_attributes}
	WithCustomAttributes interface{} `field:"optional" json:"withCustomAttributes" yaml:"withCustomAttributes"`
	// Limit by projects with issues feature enabled. Default is `false`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/projects#with_issues_enabled DataGitlabProjects#with_issues_enabled}
	WithIssuesEnabled interface{} `field:"optional" json:"withIssuesEnabled" yaml:"withIssuesEnabled"`
	// Limit by projects with merge requests feature enabled. Default is `false`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/projects#with_merge_requests_enabled DataGitlabProjects#with_merge_requests_enabled}
	WithMergeRequestsEnabled interface{} `field:"optional" json:"withMergeRequestsEnabled" yaml:"withMergeRequestsEnabled"`
	// Limit by projects which use the given programming language. Cannot be used with `group_id`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/projects#with_programming_language DataGitlabProjects#with_programming_language}
	WithProgrammingLanguage *string `field:"optional" json:"withProgrammingLanguage" yaml:"withProgrammingLanguage"`
	// Include projects shared to this group. Default is `true`. Needs `group_id`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/d/projects#with_shared DataGitlabProjects#with_shared}
	WithShared interface{} `field:"optional" json:"withShared" yaml:"withShared"`
}

