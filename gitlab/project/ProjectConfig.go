package project

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectConfig struct {
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
	// The name of the project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#name Project#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Set to true if you want to treat skipped pipelines as if they finished with success.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#allow_merge_on_skipped_pipeline Project#allow_merge_on_skipped_pipeline}
	AllowMergeOnSkippedPipeline interface{} `field:"optional" json:"allowMergeOnSkippedPipeline" yaml:"allowMergeOnSkippedPipeline"`
	// Set the analytics access level. Valid values are `disabled`, `private`, `enabled`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#analytics_access_level Project#analytics_access_level}
	AnalyticsAccessLevel *string `field:"optional" json:"analyticsAccessLevel" yaml:"analyticsAccessLevel"`
	// Number of merge request approvals required for merging.
	//
	// Default is 0.
	// This field **does not** work well in combination with the `gitlab_project_approval_rule` resource
	// and is most likely gonna be deprecated in a future GitLab version (see [this upstream epic](https://gitlab.com/groups/gitlab-org/-/epics/7572)).
	// In the meantime we recommend against using this attribute and use `gitlab_project_approval_rule` instead.
	//
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#approvals_before_merge Project#approvals_before_merge}
	ApprovalsBeforeMerge *float64 `field:"optional" json:"approvalsBeforeMerge" yaml:"approvalsBeforeMerge"`
	// Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#archived Project#archived}
	Archived interface{} `field:"optional" json:"archived" yaml:"archived"`
	// Set to `true` to archive the project instead of deleting on destroy.
	//
	// If set to `true` it will entire omit the `DELETE` operation.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#archive_on_destroy Project#archive_on_destroy}
	ArchiveOnDestroy interface{} `field:"optional" json:"archiveOnDestroy" yaml:"archiveOnDestroy"`
	// Auto-cancel pending pipelines. This isn’t a boolean, but enabled/disabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#auto_cancel_pending_pipelines Project#auto_cancel_pending_pipelines}
	AutoCancelPendingPipelines *string `field:"optional" json:"autoCancelPendingPipelines" yaml:"autoCancelPendingPipelines"`
	// Set whether auto-closing referenced issues on default branch.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#autoclose_referenced_issues Project#autoclose_referenced_issues}
	AutocloseReferencedIssues interface{} `field:"optional" json:"autocloseReferencedIssues" yaml:"autocloseReferencedIssues"`
	// Auto Deploy strategy. Valid values are `continuous`, `manual`, `timed_incremental`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#auto_devops_deploy_strategy Project#auto_devops_deploy_strategy}
	AutoDevopsDeployStrategy *string `field:"optional" json:"autoDevopsDeployStrategy" yaml:"autoDevopsDeployStrategy"`
	// Enable Auto DevOps for this project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#auto_devops_enabled Project#auto_devops_enabled}
	AutoDevopsEnabled interface{} `field:"optional" json:"autoDevopsEnabled" yaml:"autoDevopsEnabled"`
	// A local path to the avatar image to upload. **Note**: not available for imported resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#avatar Project#avatar}
	Avatar *string `field:"optional" json:"avatar" yaml:"avatar"`
	// The hash of the avatar image.
	//
	// Use `filesha256("path/to/avatar.png")` whenever possible. **Note**: this is used to trigger an update of the avatar. If it's not given, but an avatar is given, the avatar will be updated each time.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#avatar_hash Project#avatar_hash}
	AvatarHash *string `field:"optional" json:"avatarHash" yaml:"avatarHash"`
	// Test coverage parsing for the project. This is deprecated feature in GitLab 15.0.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#build_coverage_regex Project#build_coverage_regex}
	BuildCoverageRegex *string `field:"optional" json:"buildCoverageRegex" yaml:"buildCoverageRegex"`
	// The Git strategy. Defaults to fetch.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#build_git_strategy Project#build_git_strategy}
	BuildGitStrategy *string `field:"optional" json:"buildGitStrategy" yaml:"buildGitStrategy"`
	// Set the builds access level. Valid values are `disabled`, `private`, `enabled`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#builds_access_level Project#builds_access_level}
	BuildsAccessLevel *string `field:"optional" json:"buildsAccessLevel" yaml:"buildsAccessLevel"`
	// The maximum amount of time, in seconds, that a job can run.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#build_timeout Project#build_timeout}
	BuildTimeout *float64 `field:"optional" json:"buildTimeout" yaml:"buildTimeout"`
	// Custom Path to CI config file.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#ci_config_path Project#ci_config_path}
	CiConfigPath *string `field:"optional" json:"ciConfigPath" yaml:"ciConfigPath"`
	// Default number of revisions for shallow cloning.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#ci_default_git_depth Project#ci_default_git_depth}
	CiDefaultGitDepth *float64 `field:"optional" json:"ciDefaultGitDepth" yaml:"ciDefaultGitDepth"`
	// When a new deployment job starts, skip older deployment jobs that are still pending.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#ci_forward_deployment_enabled Project#ci_forward_deployment_enabled}
	CiForwardDeploymentEnabled interface{} `field:"optional" json:"ciForwardDeploymentEnabled" yaml:"ciForwardDeploymentEnabled"`
	// Use separate caches for protected branches.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#ci_separated_caches Project#ci_separated_caches}
	CiSeparatedCaches interface{} `field:"optional" json:"ciSeparatedCaches" yaml:"ciSeparatedCaches"`
	// container_expiration_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#container_expiration_policy Project#container_expiration_policy}
	ContainerExpirationPolicy *ProjectContainerExpirationPolicy `field:"optional" json:"containerExpirationPolicy" yaml:"containerExpirationPolicy"`
	// Set visibility of container registry, for this project. Valid values are `disabled`, `private`, `enabled`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#container_registry_access_level Project#container_registry_access_level}
	ContainerRegistryAccessLevel *string `field:"optional" json:"containerRegistryAccessLevel" yaml:"containerRegistryAccessLevel"`
	// Enable container registry for the project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#container_registry_enabled Project#container_registry_enabled}
	ContainerRegistryEnabled interface{} `field:"optional" json:"containerRegistryEnabled" yaml:"containerRegistryEnabled"`
	// The default branch for the project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#default_branch Project#default_branch}
	DefaultBranch *string `field:"optional" json:"defaultBranch" yaml:"defaultBranch"`
	// A description of the project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#description Project#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Disable email notifications.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#emails_disabled Project#emails_disabled}
	EmailsDisabled interface{} `field:"optional" json:"emailsDisabled" yaml:"emailsDisabled"`
	// The classification label for the project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#external_authorization_classification_label Project#external_authorization_classification_label}
	ExternalAuthorizationClassificationLabel *string `field:"optional" json:"externalAuthorizationClassificationLabel" yaml:"externalAuthorizationClassificationLabel"`
	// The id of the project to fork.
	//
	// During create the project is forked and during an update the fork relation is changed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#forked_from_project_id Project#forked_from_project_id}
	ForkedFromProjectId *float64 `field:"optional" json:"forkedFromProjectId" yaml:"forkedFromProjectId"`
	// Set the forking access level. Valid values are `disabled`, `private`, `enabled`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#forking_access_level Project#forking_access_level}
	ForkingAccessLevel *string `field:"optional" json:"forkingAccessLevel" yaml:"forkingAccessLevel"`
	// For group-level custom templates, specifies ID of group from which all the custom project templates are sourced.
	//
	// Leave empty for instance-level templates. Requires use_custom_template to be true (enterprise edition).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#group_with_project_templates_id Project#group_with_project_templates_id}
	GroupWithProjectTemplatesId *float64 `field:"optional" json:"groupWithProjectTemplatesId" yaml:"groupWithProjectTemplatesId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#id Project#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Git URL to a repository to be imported.
	//
	// Together with `mirror = true` it will setup a Pull Mirror. This can also be used together with `forked_from_project_id` to setup a Pull Mirror for a fork. The fork takes precedence over the import. Make sure to provide the credentials in `import_url_username` and `import_url_password`. GitLab never returns the credentials, thus the provider cannot detect configuration drift in the credentials. They can also not be imported using `terraform import`. See the examples section for how to properly use it.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#import_url Project#import_url}
	ImportUrl *string `field:"optional" json:"importUrl" yaml:"importUrl"`
	// The password for the `import_url`.
	//
	// The value of this field is used to construct a valid `import_url` and is only related to the provider. This field cannot be imported using `terraform import`. See the examples section for how to properly use it.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#import_url_password Project#import_url_password}
	ImportUrlPassword *string `field:"optional" json:"importUrlPassword" yaml:"importUrlPassword"`
	// The username for the `import_url`.
	//
	// The value of this field is used to construct a valid `import_url` and is only related to the provider. This field cannot be imported using `terraform import`.  See the examples section for how to properly use it.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#import_url_username Project#import_url_username}
	ImportUrlUsername *string `field:"optional" json:"importUrlUsername" yaml:"importUrlUsername"`
	// Create main branch with first commit containing a README.md file.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#initialize_with_readme Project#initialize_with_readme}
	InitializeWithReadme interface{} `field:"optional" json:"initializeWithReadme" yaml:"initializeWithReadme"`
	// Set the issues access level. Valid values are `disabled`, `private`, `enabled`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#issues_access_level Project#issues_access_level}
	IssuesAccessLevel *string `field:"optional" json:"issuesAccessLevel" yaml:"issuesAccessLevel"`
	// Enable issue tracking for the project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#issues_enabled Project#issues_enabled}
	IssuesEnabled interface{} `field:"optional" json:"issuesEnabled" yaml:"issuesEnabled"`
	// Sets the template for new issues in the project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#issues_template Project#issues_template}
	IssuesTemplate *string `field:"optional" json:"issuesTemplate" yaml:"issuesTemplate"`
	// Enable LFS for the project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#lfs_enabled Project#lfs_enabled}
	LfsEnabled interface{} `field:"optional" json:"lfsEnabled" yaml:"lfsEnabled"`
	// Template used to create merge commit message in merge requests. (Introduced in GitLab 14.5.).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#merge_commit_template Project#merge_commit_template}
	MergeCommitTemplate *string `field:"optional" json:"mergeCommitTemplate" yaml:"mergeCommitTemplate"`
	// Set the merge method. Valid values are `merge`, `rebase_merge`, `ff`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#merge_method Project#merge_method}
	MergeMethod *string `field:"optional" json:"mergeMethod" yaml:"mergeMethod"`
	// Enable or disable merge pipelines.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#merge_pipelines_enabled Project#merge_pipelines_enabled}
	MergePipelinesEnabled interface{} `field:"optional" json:"mergePipelinesEnabled" yaml:"mergePipelinesEnabled"`
	// Set the merge requests access level. Valid values are `disabled`, `private`, `enabled`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#merge_requests_access_level Project#merge_requests_access_level}
	MergeRequestsAccessLevel *string `field:"optional" json:"mergeRequestsAccessLevel" yaml:"mergeRequestsAccessLevel"`
	// Enable merge requests for the project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#merge_requests_enabled Project#merge_requests_enabled}
	MergeRequestsEnabled interface{} `field:"optional" json:"mergeRequestsEnabled" yaml:"mergeRequestsEnabled"`
	// Sets the template for new merge requests in the project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#merge_requests_template Project#merge_requests_template}
	MergeRequestsTemplate *string `field:"optional" json:"mergeRequestsTemplate" yaml:"mergeRequestsTemplate"`
	// Enable or disable merge trains. Requires `merge_pipelines_enabled` to be set to `true` to take effect.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#merge_trains_enabled Project#merge_trains_enabled}
	MergeTrainsEnabled interface{} `field:"optional" json:"mergeTrainsEnabled" yaml:"mergeTrainsEnabled"`
	// Enable project pull mirror.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#mirror Project#mirror}
	Mirror interface{} `field:"optional" json:"mirror" yaml:"mirror"`
	// Enable overwrite diverged branches for a mirrored project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#mirror_overwrites_diverged_branches Project#mirror_overwrites_diverged_branches}
	MirrorOverwritesDivergedBranches interface{} `field:"optional" json:"mirrorOverwritesDivergedBranches" yaml:"mirrorOverwritesDivergedBranches"`
	// Enable trigger builds on pushes for a mirrored project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#mirror_trigger_builds Project#mirror_trigger_builds}
	MirrorTriggerBuilds interface{} `field:"optional" json:"mirrorTriggerBuilds" yaml:"mirrorTriggerBuilds"`
	// For forked projects, target merge requests to this project. If false, the target will be the upstream project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#mr_default_target_self Project#mr_default_target_self}
	MrDefaultTargetSelf interface{} `field:"optional" json:"mrDefaultTargetSelf" yaml:"mrDefaultTargetSelf"`
	// The namespace (group or user) of the project. Defaults to your user.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#namespace_id Project#namespace_id}
	NamespaceId *float64 `field:"optional" json:"namespaceId" yaml:"namespaceId"`
	// Set to true if you want allow merges only if all discussions are resolved.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#only_allow_merge_if_all_discussions_are_resolved Project#only_allow_merge_if_all_discussions_are_resolved}
	OnlyAllowMergeIfAllDiscussionsAreResolved interface{} `field:"optional" json:"onlyAllowMergeIfAllDiscussionsAreResolved" yaml:"onlyAllowMergeIfAllDiscussionsAreResolved"`
	// Set to true if you want allow merges only if a pipeline succeeds.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#only_allow_merge_if_pipeline_succeeds Project#only_allow_merge_if_pipeline_succeeds}
	OnlyAllowMergeIfPipelineSucceeds interface{} `field:"optional" json:"onlyAllowMergeIfPipelineSucceeds" yaml:"onlyAllowMergeIfPipelineSucceeds"`
	// Enable only mirror protected branches for a mirrored project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#only_mirror_protected_branches Project#only_mirror_protected_branches}
	OnlyMirrorProtectedBranches interface{} `field:"optional" json:"onlyMirrorProtectedBranches" yaml:"onlyMirrorProtectedBranches"`
	// Set the operations access level. Valid values are `disabled`, `private`, `enabled`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#operations_access_level Project#operations_access_level}
	OperationsAccessLevel *string `field:"optional" json:"operationsAccessLevel" yaml:"operationsAccessLevel"`
	// Enable packages repository for the project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#packages_enabled Project#packages_enabled}
	PackagesEnabled interface{} `field:"optional" json:"packagesEnabled" yaml:"packagesEnabled"`
	// Enable pages access control.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#pages_access_level Project#pages_access_level}
	PagesAccessLevel *string `field:"optional" json:"pagesAccessLevel" yaml:"pagesAccessLevel"`
	// The path of the repository.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#path Project#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Enable pipelines for the project. The `pipelines_enabled` field is being sent as `jobs_enabled` in the GitLab API calls.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#pipelines_enabled Project#pipelines_enabled}
	PipelinesEnabled interface{} `field:"optional" json:"pipelinesEnabled" yaml:"pipelinesEnabled"`
	// Show link to create/view merge request when pushing from the command line.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#printing_merge_request_link_enabled Project#printing_merge_request_link_enabled}
	PrintingMergeRequestLinkEnabled interface{} `field:"optional" json:"printingMergeRequestLinkEnabled" yaml:"printingMergeRequestLinkEnabled"`
	// If true, jobs can be viewed by non-project members.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#public_builds Project#public_builds}
	PublicBuilds interface{} `field:"optional" json:"publicBuilds" yaml:"publicBuilds"`
	// push_rules block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#push_rules Project#push_rules}
	PushRules *ProjectPushRules `field:"optional" json:"pushRules" yaml:"pushRules"`
	// Enable `Delete source branch` option by default for all new merge requests.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#remove_source_branch_after_merge Project#remove_source_branch_after_merge}
	RemoveSourceBranchAfterMerge interface{} `field:"optional" json:"removeSourceBranchAfterMerge" yaml:"removeSourceBranchAfterMerge"`
	// Set the repository access level. Valid values are `disabled`, `private`, `enabled`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#repository_access_level Project#repository_access_level}
	RepositoryAccessLevel *string `field:"optional" json:"repositoryAccessLevel" yaml:"repositoryAccessLevel"`
	// Which storage shard the repository is on. (administrator only).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#repository_storage Project#repository_storage}
	RepositoryStorage *string `field:"optional" json:"repositoryStorage" yaml:"repositoryStorage"`
	// Allow users to request member access.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#request_access_enabled Project#request_access_enabled}
	RequestAccessEnabled interface{} `field:"optional" json:"requestAccessEnabled" yaml:"requestAccessEnabled"`
	// Set the requirements access level. Valid values are `disabled`, `private`, `enabled`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#requirements_access_level Project#requirements_access_level}
	RequirementsAccessLevel *string `field:"optional" json:"requirementsAccessLevel" yaml:"requirementsAccessLevel"`
	// Automatically resolve merge request diffs discussions on lines changed with a push.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#resolve_outdated_diff_discussions Project#resolve_outdated_diff_discussions}
	ResolveOutdatedDiffDiscussions interface{} `field:"optional" json:"resolveOutdatedDiffDiscussions" yaml:"resolveOutdatedDiffDiscussions"`
	// Allow only users with the Maintainer role to pass user-defined variables when triggering a pipeline.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#restrict_user_defined_variables Project#restrict_user_defined_variables}
	RestrictUserDefinedVariables interface{} `field:"optional" json:"restrictUserDefinedVariables" yaml:"restrictUserDefinedVariables"`
	// Set the security and compliance access level. Valid values are `disabled`, `private`, `enabled`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#security_and_compliance_access_level Project#security_and_compliance_access_level}
	SecurityAndComplianceAccessLevel *string `field:"optional" json:"securityAndComplianceAccessLevel" yaml:"securityAndComplianceAccessLevel"`
	// Enable shared runners for this project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#shared_runners_enabled Project#shared_runners_enabled}
	SharedRunnersEnabled interface{} `field:"optional" json:"sharedRunnersEnabled" yaml:"sharedRunnersEnabled"`
	// If `true`, the default behavior to wait for the default branch protection to be created is skipped.
	//
	// This is necessary if the current user is not an admin and the default branch protection is disabled on an instance-level.
	// There is currently no known way to determine if the default branch protection is disabled on an instance-level for non-admin users.
	// This attribute is only used during resource creation, thus changes are suppressed and the attribute cannot be imported.
	//
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#skip_wait_for_default_branch_protection Project#skip_wait_for_default_branch_protection}
	SkipWaitForDefaultBranchProtection interface{} `field:"optional" json:"skipWaitForDefaultBranchProtection" yaml:"skipWaitForDefaultBranchProtection"`
	// Set the snippets access level. Valid values are `disabled`, `private`, `enabled`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#snippets_access_level Project#snippets_access_level}
	SnippetsAccessLevel *string `field:"optional" json:"snippetsAccessLevel" yaml:"snippetsAccessLevel"`
	// Enable snippets for the project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#snippets_enabled Project#snippets_enabled}
	SnippetsEnabled interface{} `field:"optional" json:"snippetsEnabled" yaml:"snippetsEnabled"`
	// Template used to create squash commit message in merge requests. (Introduced in GitLab 14.6.).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#squash_commit_template Project#squash_commit_template}
	SquashCommitTemplate *string `field:"optional" json:"squashCommitTemplate" yaml:"squashCommitTemplate"`
	// Squash commits when merge request.
	//
	// Valid values are `never`, `always`, `default_on`, or `default_off`. The default value is `default_off`. [GitLab >= 14.1]
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#squash_option Project#squash_option}
	SquashOption *string `field:"optional" json:"squashOption" yaml:"squashOption"`
	// The commit message used to apply merge request suggestions.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#suggestion_commit_message Project#suggestion_commit_message}
	SuggestionCommitMessage *string `field:"optional" json:"suggestionCommitMessage" yaml:"suggestionCommitMessage"`
	// The list of tags for a project;
	//
	// put array of tags, that should be finally assigned to a project. Use topics instead.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#tags Project#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// When used without use_custom_template, name of a built-in project template.
	//
	// When used with use_custom_template, name of a custom project template. This option is mutually exclusive with `template_project_id`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#template_name Project#template_name}
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
	// When used with use_custom_template, project ID of a custom project template.
	//
	// This is preferable to using template_name since template_name may be ambiguous (enterprise edition). This option is mutually exclusive with `template_name`. See `gitlab_group_project_file_template` to set a project as a template project. If a project has not been set as a template, using it here will result in an error.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#template_project_id Project#template_project_id}
	TemplateProjectId *float64 `field:"optional" json:"templateProjectId" yaml:"templateProjectId"`
	// The list of topics for the project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#topics Project#topics}
	Topics *[]*string `field:"optional" json:"topics" yaml:"topics"`
	// Use either custom instance or group (with group_with_project_templates_id) project template (enterprise edition).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#use_custom_template Project#use_custom_template}
	UseCustomTemplate interface{} `field:"optional" json:"useCustomTemplate" yaml:"useCustomTemplate"`
	// Set to `public` to create a public project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#visibility_level Project#visibility_level}
	VisibilityLevel *string `field:"optional" json:"visibilityLevel" yaml:"visibilityLevel"`
	// Set the wiki access level. Valid values are `disabled`, `private`, `enabled`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#wiki_access_level Project#wiki_access_level}
	WikiAccessLevel *string `field:"optional" json:"wikiAccessLevel" yaml:"wikiAccessLevel"`
	// Enable wiki for the project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#wiki_enabled Project#wiki_enabled}
	WikiEnabled interface{} `field:"optional" json:"wikiEnabled" yaml:"wikiEnabled"`
}

