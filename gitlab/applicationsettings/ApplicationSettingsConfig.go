package applicationsettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationSettingsConfig struct {
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
	// If set, abuse reports are sent to this address. Abuse reports are always available in the Admin Area.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#abuse_notification_email ApplicationSettings#abuse_notification_email}
	AbuseNotificationEmail *string `field:"optional" json:"abuseNotificationEmail" yaml:"abuseNotificationEmail"`
	// Require administrators to enable Admin Mode by re-authenticating for administrative tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#admin_mode ApplicationSettings#admin_mode}
	AdminMode interface{} `field:"optional" json:"adminMode" yaml:"adminMode"`
	// Where to redirect users after logout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#after_sign_out_path ApplicationSettings#after_sign_out_path}
	AfterSignOutPath *string `field:"optional" json:"afterSignOutPath" yaml:"afterSignOutPath"`
	// Text shown to the user after signing up.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#after_sign_up_text ApplicationSettings#after_sign_up_text}
	AfterSignUpText *string `field:"optional" json:"afterSignUpText" yaml:"afterSignUpText"`
	// API key for Akismet spam protection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#akismet_api_key ApplicationSettings#akismet_api_key}
	AkismetApiKey *string `field:"optional" json:"akismetApiKey" yaml:"akismetApiKey"`
	// (If enabled, requires: akismet_api_key) Enable or disable Akismet spam protection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#akismet_enabled ApplicationSettings#akismet_enabled}
	AkismetEnabled interface{} `field:"optional" json:"akismetEnabled" yaml:"akismetEnabled"`
	// Set to true to allow group owners to manage LDAP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#allow_group_owners_to_manage_ldap ApplicationSettings#allow_group_owners_to_manage_ldap}
	AllowGroupOwnersToManageLdap interface{} `field:"optional" json:"allowGroupOwnersToManageLdap" yaml:"allowGroupOwnersToManageLdap"`
	// Allow requests to the local network from system hooks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#allow_local_requests_from_system_hooks ApplicationSettings#allow_local_requests_from_system_hooks}
	AllowLocalRequestsFromSystemHooks interface{} `field:"optional" json:"allowLocalRequestsFromSystemHooks" yaml:"allowLocalRequestsFromSystemHooks"`
	// Allow requests to the local network from web hooks and services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#allow_local_requests_from_web_hooks_and_services ApplicationSettings#allow_local_requests_from_web_hooks_and_services}
	AllowLocalRequestsFromWebHooksAndServices interface{} `field:"optional" json:"allowLocalRequestsFromWebHooksAndServices" yaml:"allowLocalRequestsFromWebHooksAndServices"`
	// Set the duration for which the jobs are considered as old and expired.
	//
	// After that time passes, the jobs are archived and no longer able to be retried. Make it empty to never expire jobs. It has to be no less than 1 day, for example: 15 days, 1 month, 2 years.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#archive_builds_in_human_readable ApplicationSettings#archive_builds_in_human_readable}
	ArchiveBuildsInHumanReadable *string `field:"optional" json:"archiveBuildsInHumanReadable" yaml:"archiveBuildsInHumanReadable"`
	// Assets that match these domains are not proxied.
	//
	// Wildcards allowed. Your GitLab installation URL is automatically allowlisted. GitLab restart is required to apply changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#asset_proxy_allowlist ApplicationSettings#asset_proxy_allowlist}
	AssetProxyAllowlist *[]*string `field:"optional" json:"assetProxyAllowlist" yaml:"assetProxyAllowlist"`
	// (If enabled, requires: asset_proxy_url) Enable proxying of assets. GitLab restart is required to apply changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#asset_proxy_enabled ApplicationSettings#asset_proxy_enabled}
	AssetProxyEnabled interface{} `field:"optional" json:"assetProxyEnabled" yaml:"assetProxyEnabled"`
	// Shared secret with the asset proxy server. GitLab restart is required to apply changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#asset_proxy_secret_key ApplicationSettings#asset_proxy_secret_key}
	AssetProxySecretKey *string `field:"optional" json:"assetProxySecretKey" yaml:"assetProxySecretKey"`
	// URL of the asset proxy server. GitLab restart is required to apply changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#asset_proxy_url ApplicationSettings#asset_proxy_url}
	AssetProxyUrl *string `field:"optional" json:"assetProxyUrl" yaml:"assetProxyUrl"`
	// By default, we write to the authorized_keys file to support Git over SSH without additional configuration.
	//
	// GitLab can be optimized to authenticate SSH keys via the database file. Only disable this if you have configured your OpenSSH server to use the AuthorizedKeysCommand.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#authorized_keys_enabled ApplicationSettings#authorized_keys_enabled}
	AuthorizedKeysEnabled interface{} `field:"optional" json:"authorizedKeysEnabled" yaml:"authorizedKeysEnabled"`
	// Specify a domain to use by default for every project’s Auto Review Apps and Auto Deploy stages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#auto_devops_domain ApplicationSettings#auto_devops_domain}
	AutoDevopsDomain *string `field:"optional" json:"autoDevopsDomain" yaml:"autoDevopsDomain"`
	// Enable Auto DevOps for projects by default.
	//
	// It automatically builds, tests, and deploys applications based on a predefined CI/CD configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#auto_devops_enabled ApplicationSettings#auto_devops_enabled}
	AutoDevopsEnabled interface{} `field:"optional" json:"autoDevopsEnabled" yaml:"autoDevopsEnabled"`
	// Enabling this permits automatic allocation of purchased storage in a namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#automatic_purchased_storage_allocation ApplicationSettings#automatic_purchased_storage_allocation}
	AutomaticPurchasedStorageAllocation interface{} `field:"optional" json:"automaticPurchasedStorageAllocation" yaml:"automaticPurchasedStorageAllocation"`
	// Indicates whether users can create top-level groups. Introduced in GitLab 15.5.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#can_create_group ApplicationSettings#can_create_group}
	CanCreateGroup interface{} `field:"optional" json:"canCreateGroup" yaml:"canCreateGroup"`
	// Enabling this makes only licensed EE features available to projects if the project namespace’s plan includes the feature or if the project is public.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#check_namespace_plan ApplicationSettings#check_namespace_plan}
	CheckNamespacePlan interface{} `field:"optional" json:"checkNamespacePlan" yaml:"checkNamespacePlan"`
	// Custom hostname (for private commit emails).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#commit_email_hostname ApplicationSettings#commit_email_hostname}
	CommitEmailHostname *string `field:"optional" json:"commitEmailHostname" yaml:"commitEmailHostname"`
	// Enable cleanup policies for all projects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#container_expiration_policies_enable_historic_entries ApplicationSettings#container_expiration_policies_enable_historic_entries}
	ContainerExpirationPoliciesEnableHistoricEntries interface{} `field:"optional" json:"containerExpirationPoliciesEnableHistoricEntries" yaml:"containerExpirationPoliciesEnableHistoricEntries"`
	// The maximum number of tags that can be deleted in a single execution of cleanup policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#container_registry_cleanup_tags_service_max_list_size ApplicationSettings#container_registry_cleanup_tags_service_max_list_size}
	ContainerRegistryCleanupTagsServiceMaxListSize *float64 `field:"optional" json:"containerRegistryCleanupTagsServiceMaxListSize" yaml:"containerRegistryCleanupTagsServiceMaxListSize"`
	// The maximum time, in seconds, that the cleanup process can take to delete a batch of tags for cleanup policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#container_registry_delete_tags_service_timeout ApplicationSettings#container_registry_delete_tags_service_timeout}
	ContainerRegistryDeleteTagsServiceTimeout *float64 `field:"optional" json:"containerRegistryDeleteTagsServiceTimeout" yaml:"containerRegistryDeleteTagsServiceTimeout"`
	// Caching during the execution of cleanup policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#container_registry_expiration_policies_caching ApplicationSettings#container_registry_expiration_policies_caching}
	ContainerRegistryExpirationPoliciesCaching interface{} `field:"optional" json:"containerRegistryExpirationPoliciesCaching" yaml:"containerRegistryExpirationPoliciesCaching"`
	// Number of workers for cleanup policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#container_registry_expiration_policies_worker_capacity ApplicationSettings#container_registry_expiration_policies_worker_capacity}
	ContainerRegistryExpirationPoliciesWorkerCapacity *float64 `field:"optional" json:"containerRegistryExpirationPoliciesWorkerCapacity" yaml:"containerRegistryExpirationPoliciesWorkerCapacity"`
	// Container Registry token duration in minutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#container_registry_token_expire_delay ApplicationSettings#container_registry_token_expire_delay}
	ContainerRegistryTokenExpireDelay *float64 `field:"optional" json:"containerRegistryTokenExpireDelay" yaml:"containerRegistryTokenExpireDelay"`
	// Enable automatic deactivation of dormant users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#deactivate_dormant_users ApplicationSettings#deactivate_dormant_users}
	DeactivateDormantUsers interface{} `field:"optional" json:"deactivateDormantUsers" yaml:"deactivateDormantUsers"`
	// Set the default expiration time for each job’s artifacts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#default_artifacts_expire_in ApplicationSettings#default_artifacts_expire_in}
	DefaultArtifactsExpireIn *string `field:"optional" json:"defaultArtifactsExpireIn" yaml:"defaultArtifactsExpireIn"`
	// Instance-level custom initial branch name (introduced in GitLab 13.2).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#default_branch_name ApplicationSettings#default_branch_name}
	DefaultBranchName *string `field:"optional" json:"defaultBranchName" yaml:"defaultBranchName"`
	// Determine if developers can push to the default branch.
	//
	// Can take: 0 (not protected, both users with the Developer role or Maintainer role can push new commits and force push), 1 (partially protected, users with the Developer role or Maintainer role can push new commits, but cannot force push) or 2 (fully protected, users with the Developer or Maintainer role cannot push new commits, but users with the Developer or Maintainer role can; no one can force push) as a parameter. Default is 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#default_branch_protection ApplicationSettings#default_branch_protection}
	DefaultBranchProtection *float64 `field:"optional" json:"defaultBranchProtection" yaml:"defaultBranchProtection"`
	// Default CI/CD configuration file and path for new projects (.gitlab-ci.yml if not set).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#default_ci_config_path ApplicationSettings#default_ci_config_path}
	DefaultCiConfigPath *string `field:"optional" json:"defaultCiConfigPath" yaml:"defaultCiConfigPath"`
	// What visibility level new groups receive. Can take private, internal and public as a parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#default_group_visibility ApplicationSettings#default_group_visibility}
	DefaultGroupVisibility *string `field:"optional" json:"defaultGroupVisibility" yaml:"defaultGroupVisibility"`
	// Default project creation protection. Can take: 0 (No one), 1 (Maintainers) or 2 (Developers + Maintainers).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#default_project_creation ApplicationSettings#default_project_creation}
	DefaultProjectCreation *float64 `field:"optional" json:"defaultProjectCreation" yaml:"defaultProjectCreation"`
	// Project limit per user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#default_projects_limit ApplicationSettings#default_projects_limit}
	DefaultProjectsLimit *float64 `field:"optional" json:"defaultProjectsLimit" yaml:"defaultProjectsLimit"`
	// What visibility level new projects receive. Can take private, internal and public as a parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#default_project_visibility ApplicationSettings#default_project_visibility}
	DefaultProjectVisibility *string `field:"optional" json:"defaultProjectVisibility" yaml:"defaultProjectVisibility"`
	// What visibility level new snippets receive. Can take private, internal and public as a parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#default_snippet_visibility ApplicationSettings#default_snippet_visibility}
	DefaultSnippetVisibility *string `field:"optional" json:"defaultSnippetVisibility" yaml:"defaultSnippetVisibility"`
	// Enable delayed group deletion.
	//
	// Introduced in GitLab 15.0. From GitLab 15.1, disables and locks the group-level setting for delayed protect deletion when set to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#delayed_group_deletion ApplicationSettings#delayed_group_deletion}
	DelayedGroupDeletion interface{} `field:"optional" json:"delayedGroupDeletion" yaml:"delayedGroupDeletion"`
	// Enable delayed project deletion by default in new groups.
	//
	// From GitLab 15.1, can only be enabled when delayed_group_deletion is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#delayed_project_deletion ApplicationSettings#delayed_project_deletion}
	DelayedProjectDeletion interface{} `field:"optional" json:"delayedProjectDeletion" yaml:"delayedProjectDeletion"`
	// Enable inactive project deletion feature. Introduced in GitLab 14.10. Became operational in GitLab 15.0 (with feature flag inactive_projects_deletion).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#delete_inactive_projects ApplicationSettings#delete_inactive_projects}
	DeleteInactiveProjects interface{} `field:"optional" json:"deleteInactiveProjects" yaml:"deleteInactiveProjects"`
	// The number of days to wait before deleting a project or group that is marked for deletion.
	//
	// Value must be between 1 and 90. From GitLab 15.1, a hook on deletion_adjourned_period sets the period to 1 on every update, and sets both delayed_project_deletion and delayed_group_deletion to false if the period is 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#deletion_adjourned_period ApplicationSettings#deletion_adjourned_period}
	DeletionAdjournedPeriod *float64 `field:"optional" json:"deletionAdjournedPeriod" yaml:"deletionAdjournedPeriod"`
	// Maximum files in a diff.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#diff_max_files ApplicationSettings#diff_max_files}
	DiffMaxFiles *float64 `field:"optional" json:"diffMaxFiles" yaml:"diffMaxFiles"`
	// Maximum lines in a diff.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#diff_max_lines ApplicationSettings#diff_max_lines}
	DiffMaxLines *float64 `field:"optional" json:"diffMaxLines" yaml:"diffMaxLines"`
	// Maximum diff patch size, in bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#diff_max_patch_bytes ApplicationSettings#diff_max_patch_bytes}
	DiffMaxPatchBytes *float64 `field:"optional" json:"diffMaxPatchBytes" yaml:"diffMaxPatchBytes"`
	// Disabled OAuth sign-in sources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#disabled_oauth_sign_in_sources ApplicationSettings#disabled_oauth_sign_in_sources}
	DisabledOauthSignInSources *[]*string `field:"optional" json:"disabledOauthSignInSources" yaml:"disabledOauthSignInSources"`
	// Disable display of RSS/Atom and calendar feed tokens (introduced in GitLab 13.7).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#disable_feed_token ApplicationSettings#disable_feed_token}
	DisableFeedToken interface{} `field:"optional" json:"disableFeedToken" yaml:"disableFeedToken"`
	// Enforce DNS rebinding attack protection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#dns_rebinding_protection_enabled ApplicationSettings#dns_rebinding_protection_enabled}
	DnsRebindingProtectionEnabled interface{} `field:"optional" json:"dnsRebindingProtectionEnabled" yaml:"dnsRebindingProtectionEnabled"`
	// Force people to use only corporate emails for sign-up. Null means there is no restriction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#domain_allowlist ApplicationSettings#domain_allowlist}
	DomainAllowlist *[]*string `field:"optional" json:"domainAllowlist" yaml:"domainAllowlist"`
	// Users with email addresses that match these domains cannot sign up.
	//
	// Wildcards allowed. Use separate lines for multiple entries. Ex: domain.com, *.domain.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#domain_denylist ApplicationSettings#domain_denylist}
	DomainDenylist *[]*string `field:"optional" json:"domainDenylist" yaml:"domainDenylist"`
	// (If enabled, requires: domain_denylist) Allows blocking sign-ups from emails from specific domains.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#domain_denylist_enabled ApplicationSettings#domain_denylist_enabled}
	DomainDenylistEnabled interface{} `field:"optional" json:"domainDenylistEnabled" yaml:"domainDenylistEnabled"`
	// The minimum allowed bit length of an uploaded DSA key. 0 means no restriction. -1 disables DSA keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#dsa_key_restriction ApplicationSettings#dsa_key_restriction}
	DsaKeyRestriction *float64 `field:"optional" json:"dsaKeyRestriction" yaml:"dsaKeyRestriction"`
	// The minimum allowed curve size (in bits) of an uploaded ECDSA key.
	//
	// 0 means no restriction. -1 disables ECDSA keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#ecdsa_key_restriction ApplicationSettings#ecdsa_key_restriction}
	EcdsaKeyRestriction *float64 `field:"optional" json:"ecdsaKeyRestriction" yaml:"ecdsaKeyRestriction"`
	// The minimum allowed curve size (in bits) of an uploaded ECDSA_SK key.
	//
	// 0 means no restriction. -1 disables ECDSA_SK keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#ecdsa_sk_key_restriction ApplicationSettings#ecdsa_sk_key_restriction}
	EcdsaSkKeyRestriction *float64 `field:"optional" json:"ecdsaSkKeyRestriction" yaml:"ecdsaSkKeyRestriction"`
	// The minimum allowed curve size (in bits) of an uploaded ED25519 key.
	//
	// 0 means no restriction. -1 disables ED25519 keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#ed25519_key_restriction ApplicationSettings#ed25519_key_restriction}
	Ed25519KeyRestriction *float64 `field:"optional" json:"ed25519KeyRestriction" yaml:"ed25519KeyRestriction"`
	// The minimum allowed curve size (in bits) of an uploaded ED25519_SK key.
	//
	// 0 means no restriction. -1 disables ED25519_SK keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#ed25519_sk_key_restriction ApplicationSettings#ed25519_sk_key_restriction}
	Ed25519SkKeyRestriction *float64 `field:"optional" json:"ed25519SkKeyRestriction" yaml:"ed25519SkKeyRestriction"`
	// AWS IAM access key ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#eks_access_key_id ApplicationSettings#eks_access_key_id}
	EksAccessKeyId *string `field:"optional" json:"eksAccessKeyId" yaml:"eksAccessKeyId"`
	// Amazon account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#eks_account_id ApplicationSettings#eks_account_id}
	EksAccountId *string `field:"optional" json:"eksAccountId" yaml:"eksAccountId"`
	// Enable integration with Amazon EKS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#eks_integration_enabled ApplicationSettings#eks_integration_enabled}
	EksIntegrationEnabled interface{} `field:"optional" json:"eksIntegrationEnabled" yaml:"eksIntegrationEnabled"`
	// AWS IAM secret access key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#eks_secret_access_key ApplicationSettings#eks_secret_access_key}
	EksSecretAccessKey *string `field:"optional" json:"eksSecretAccessKey" yaml:"eksSecretAccessKey"`
	// Enable the use of AWS hosted Elasticsearch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#elasticsearch_aws ApplicationSettings#elasticsearch_aws}
	ElasticsearchAws interface{} `field:"optional" json:"elasticsearchAws" yaml:"elasticsearchAws"`
	// AWS IAM access key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#elasticsearch_aws_access_key ApplicationSettings#elasticsearch_aws_access_key}
	ElasticsearchAwsAccessKey *string `field:"optional" json:"elasticsearchAwsAccessKey" yaml:"elasticsearchAwsAccessKey"`
	// The AWS region the Elasticsearch domain is configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#elasticsearch_aws_region ApplicationSettings#elasticsearch_aws_region}
	ElasticsearchAwsRegion *string `field:"optional" json:"elasticsearchAwsRegion" yaml:"elasticsearchAwsRegion"`
	// AWS IAM secret access key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#elasticsearch_aws_secret_access_key ApplicationSettings#elasticsearch_aws_secret_access_key}
	ElasticsearchAwsSecretAccessKey *string `field:"optional" json:"elasticsearchAwsSecretAccessKey" yaml:"elasticsearchAwsSecretAccessKey"`
	// Maximum size of text fields to index by Elasticsearch.
	//
	// 0 value means no limit. This does not apply to repository and wiki indexing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#elasticsearch_indexed_field_length_limit ApplicationSettings#elasticsearch_indexed_field_length_limit}
	ElasticsearchIndexedFieldLengthLimit *float64 `field:"optional" json:"elasticsearchIndexedFieldLengthLimit" yaml:"elasticsearchIndexedFieldLengthLimit"`
	// Maximum size of repository and wiki files that are indexed by Elasticsearch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#elasticsearch_indexed_file_size_limit_kb ApplicationSettings#elasticsearch_indexed_file_size_limit_kb}
	ElasticsearchIndexedFileSizeLimitKb *float64 `field:"optional" json:"elasticsearchIndexedFileSizeLimitKb" yaml:"elasticsearchIndexedFileSizeLimitKb"`
	// Enable Elasticsearch indexing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#elasticsearch_indexing ApplicationSettings#elasticsearch_indexing}
	ElasticsearchIndexing interface{} `field:"optional" json:"elasticsearchIndexing" yaml:"elasticsearchIndexing"`
	// Limit Elasticsearch to index certain namespaces and projects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#elasticsearch_limit_indexing ApplicationSettings#elasticsearch_limit_indexing}
	ElasticsearchLimitIndexing interface{} `field:"optional" json:"elasticsearchLimitIndexing" yaml:"elasticsearchLimitIndexing"`
	// Maximum concurrency of Elasticsearch bulk requests per indexing operation. This only applies to repository indexing operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#elasticsearch_max_bulk_concurrency ApplicationSettings#elasticsearch_max_bulk_concurrency}
	ElasticsearchMaxBulkConcurrency *float64 `field:"optional" json:"elasticsearchMaxBulkConcurrency" yaml:"elasticsearchMaxBulkConcurrency"`
	// Maximum size of Elasticsearch bulk indexing requests in MB. This only applies to repository indexing operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#elasticsearch_max_bulk_size_mb ApplicationSettings#elasticsearch_max_bulk_size_mb}
	ElasticsearchMaxBulkSizeMb *float64 `field:"optional" json:"elasticsearchMaxBulkSizeMb" yaml:"elasticsearchMaxBulkSizeMb"`
	// The namespaces to index via Elasticsearch if elasticsearch_limit_indexing is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#elasticsearch_namespace_ids ApplicationSettings#elasticsearch_namespace_ids}
	ElasticsearchNamespaceIds *[]*string `field:"optional" json:"elasticsearchNamespaceIds" yaml:"elasticsearchNamespaceIds"`
	// The password of your Elasticsearch instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#elasticsearch_password ApplicationSettings#elasticsearch_password}
	ElasticsearchPassword *string `field:"optional" json:"elasticsearchPassword" yaml:"elasticsearchPassword"`
	// The projects to index via Elasticsearch if elasticsearch_limit_indexing is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#elasticsearch_project_ids ApplicationSettings#elasticsearch_project_ids}
	ElasticsearchProjectIds *[]*string `field:"optional" json:"elasticsearchProjectIds" yaml:"elasticsearchProjectIds"`
	// Enable Elasticsearch search.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#elasticsearch_search ApplicationSettings#elasticsearch_search}
	ElasticsearchSearch interface{} `field:"optional" json:"elasticsearchSearch" yaml:"elasticsearchSearch"`
	// The URL to use for connecting to Elasticsearch. Use a comma-separated list to support cluster (for example, http://localhost:9200, http://localhost:9201).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#elasticsearch_url ApplicationSettings#elasticsearch_url}
	ElasticsearchUrl *[]*string `field:"optional" json:"elasticsearchUrl" yaml:"elasticsearchUrl"`
	// The username of your Elasticsearch instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#elasticsearch_username ApplicationSettings#elasticsearch_username}
	ElasticsearchUsername *string `field:"optional" json:"elasticsearchUsername" yaml:"elasticsearchUsername"`
	// Additional text added to the bottom of every email for legal/auditing/compliance reasons.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#email_additional_text ApplicationSettings#email_additional_text}
	EmailAdditionalText *string `field:"optional" json:"emailAdditionalText" yaml:"emailAdditionalText"`
	// Some email servers do not support overriding the email sender name.
	//
	// Enable this option to include the name of the author of the issue, merge request or comment in the email body instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#email_author_in_body ApplicationSettings#email_author_in_body}
	EmailAuthorInBody interface{} `field:"optional" json:"emailAuthorInBody" yaml:"emailAuthorInBody"`
	// Enabled protocols for Git access. Allowed values are: ssh, http, and nil to allow both protocols.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#enabled_git_access_protocol ApplicationSettings#enabled_git_access_protocol}
	EnabledGitAccessProtocol *string `field:"optional" json:"enabledGitAccessProtocol" yaml:"enabledGitAccessProtocol"`
	// Enabling this permits enforcement of namespace storage limits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#enforce_namespace_storage_limit ApplicationSettings#enforce_namespace_storage_limit}
	EnforceNamespaceStorageLimit interface{} `field:"optional" json:"enforceNamespaceStorageLimit" yaml:"enforceNamespaceStorageLimit"`
	// (If enabled, requires: terms) Enforce application ToS to all users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#enforce_terms ApplicationSettings#enforce_terms}
	EnforceTerms interface{} `field:"optional" json:"enforceTerms" yaml:"enforceTerms"`
	// (If enabled, requires: external_auth_client_key) The certificate to use to authenticate with the external authorization service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#external_auth_client_cert ApplicationSettings#external_auth_client_cert}
	ExternalAuthClientCert *string `field:"optional" json:"externalAuthClientCert" yaml:"externalAuthClientCert"`
	// Private key for the certificate when authentication is required for the external authorization service, this is encrypted when stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#external_auth_client_key ApplicationSettings#external_auth_client_key}
	ExternalAuthClientKey *string `field:"optional" json:"externalAuthClientKey" yaml:"externalAuthClientKey"`
	// Passphrase to use for the private key when authenticating with the external service this is encrypted when stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#external_auth_client_key_pass ApplicationSettings#external_auth_client_key_pass}
	ExternalAuthClientKeyPass *string `field:"optional" json:"externalAuthClientKeyPass" yaml:"externalAuthClientKeyPass"`
	// The default classification label to use when requesting authorization and no classification label has been specified on the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#external_authorization_service_default_label ApplicationSettings#external_authorization_service_default_label}
	ExternalAuthorizationServiceDefaultLabel *string `field:"optional" json:"externalAuthorizationServiceDefaultLabel" yaml:"externalAuthorizationServiceDefaultLabel"`
	// (If enabled, requires: external_authorization_service_default_label, external_authorization_service_timeout and external_authorization_service_url) Enable using an external authorization service for accessing projects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#external_authorization_service_enabled ApplicationSettings#external_authorization_service_enabled}
	ExternalAuthorizationServiceEnabled interface{} `field:"optional" json:"externalAuthorizationServiceEnabled" yaml:"externalAuthorizationServiceEnabled"`
	// The timeout after which an authorization request is aborted, in seconds.
	//
	// When a request times out, access is denied to the user. (min: 0.001, max: 10, step: 0.001).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#external_authorization_service_timeout ApplicationSettings#external_authorization_service_timeout}
	ExternalAuthorizationServiceTimeout *float64 `field:"optional" json:"externalAuthorizationServiceTimeout" yaml:"externalAuthorizationServiceTimeout"`
	// URL to which authorization requests are directed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#external_authorization_service_url ApplicationSettings#external_authorization_service_url}
	ExternalAuthorizationServiceUrl *string `field:"optional" json:"externalAuthorizationServiceUrl" yaml:"externalAuthorizationServiceUrl"`
	// How long to wait for a response from the pipeline validation service. Assumes OK if it times out.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#external_pipeline_validation_service_timeout ApplicationSettings#external_pipeline_validation_service_timeout}
	ExternalPipelineValidationServiceTimeout *float64 `field:"optional" json:"externalPipelineValidationServiceTimeout" yaml:"externalPipelineValidationServiceTimeout"`
	// Optional. Token to include as the X-Gitlab-Token header in requests to the URL in external_pipeline_validation_service_url.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#external_pipeline_validation_service_token ApplicationSettings#external_pipeline_validation_service_token}
	ExternalPipelineValidationServiceToken *string `field:"optional" json:"externalPipelineValidationServiceToken" yaml:"externalPipelineValidationServiceToken"`
	// URL to use for pipeline validation requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#external_pipeline_validation_service_url ApplicationSettings#external_pipeline_validation_service_url}
	ExternalPipelineValidationServiceUrl *string `field:"optional" json:"externalPipelineValidationServiceUrl" yaml:"externalPipelineValidationServiceUrl"`
	// The ID of a project to load custom file templates from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#file_template_project_id ApplicationSettings#file_template_project_id}
	FileTemplateProjectId *float64 `field:"optional" json:"fileTemplateProjectId" yaml:"fileTemplateProjectId"`
	// Start day of the week for calendar views and date pickers.
	//
	// Valid values are 0 for Sunday, 1 for Monday, and 6 for Saturday.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#first_day_of_week ApplicationSettings#first_day_of_week}
	FirstDayOfWeek *float64 `field:"optional" json:"firstDayOfWeek" yaml:"firstDayOfWeek"`
	// Comma-separated list of IPs and CIDRs of allowed secondary nodes. For example, 1.1.1.1, 2.2.2.0/24.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#geo_node_allowed_ips ApplicationSettings#geo_node_allowed_ips}
	GeoNodeAllowedIps *string `field:"optional" json:"geoNodeAllowedIps" yaml:"geoNodeAllowedIps"`
	// The amount of seconds after which a request to get a secondary node status times out.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#geo_status_timeout ApplicationSettings#geo_status_timeout}
	GeoStatusTimeout *float64 `field:"optional" json:"geoStatusTimeout" yaml:"geoStatusTimeout"`
	// Default Gitaly timeout, in seconds.
	//
	// This timeout is not enforced for Git fetch/push operations or Sidekiq jobs. Set to 0 to disable timeouts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#gitaly_timeout_default ApplicationSettings#gitaly_timeout_default}
	GitalyTimeoutDefault *float64 `field:"optional" json:"gitalyTimeoutDefault" yaml:"gitalyTimeoutDefault"`
	// Gitaly fast operation timeout, in seconds.
	//
	// Some Gitaly operations are expected to be fast. If they exceed this threshold, there may be a problem with a storage shard and ‘failing fast’ can help maintain the stability of the GitLab instance. Set to 0 to disable timeouts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#gitaly_timeout_fast ApplicationSettings#gitaly_timeout_fast}
	GitalyTimeoutFast *float64 `field:"optional" json:"gitalyTimeoutFast" yaml:"gitalyTimeoutFast"`
	// Medium Gitaly timeout, in seconds.
	//
	// This should be a value between the Fast and the Default timeout. Set to 0 to disable timeouts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#gitaly_timeout_medium ApplicationSettings#gitaly_timeout_medium}
	GitalyTimeoutMedium *float64 `field:"optional" json:"gitalyTimeoutMedium" yaml:"gitalyTimeoutMedium"`
	// List of usernames excluded from Git anti-abuse rate limits. Maximum: 100 usernames. Introduced in GitLab 15.2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#git_rate_limit_users_allowlist ApplicationSettings#git_rate_limit_users_allowlist}
	GitRateLimitUsersAllowlist *[]*string `field:"optional" json:"gitRateLimitUsersAllowlist" yaml:"gitRateLimitUsersAllowlist"`
	// Maximum duration (in minutes) of a session for Git operations when 2FA is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#git_two_factor_session_expiry ApplicationSettings#git_two_factor_session_expiry}
	GitTwoFactorSessionExpiry *float64 `field:"optional" json:"gitTwoFactorSessionExpiry" yaml:"gitTwoFactorSessionExpiry"`
	// Enable Grafana.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#grafana_enabled ApplicationSettings#grafana_enabled}
	GrafanaEnabled interface{} `field:"optional" json:"grafanaEnabled" yaml:"grafanaEnabled"`
	// Grafana URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#grafana_url ApplicationSettings#grafana_url}
	GrafanaUrl *string `field:"optional" json:"grafanaUrl" yaml:"grafanaUrl"`
	// Enable Gravatar.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#gravatar_enabled ApplicationSettings#gravatar_enabled}
	GravatarEnabled interface{} `field:"optional" json:"gravatarEnabled" yaml:"gravatarEnabled"`
	// Prevent overrides of default branch protection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#group_owners_can_manage_default_branch_protection ApplicationSettings#group_owners_can_manage_default_branch_protection}
	GroupOwnersCanManageDefaultBranchProtection interface{} `field:"optional" json:"groupOwnersCanManageDefaultBranchProtection" yaml:"groupOwnersCanManageDefaultBranchProtection"`
	// Create new projects using hashed storage paths: Enable immutable, hash-based paths and repository names to store repositories on disk.
	//
	// This prevents repositories from having to be moved or renamed when the Project URL changes and may improve disk I/O performance. (Always enabled in GitLab versions 13.0 and later, configuration is scheduled for removal in 14.0).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#hashed_storage_enabled ApplicationSettings#hashed_storage_enabled}
	HashedStorageEnabled interface{} `field:"optional" json:"hashedStorageEnabled" yaml:"hashedStorageEnabled"`
	// Hide marketing-related entries from help.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#help_page_hide_commercial_content ApplicationSettings#help_page_hide_commercial_content}
	HelpPageHideCommercialContent interface{} `field:"optional" json:"helpPageHideCommercialContent" yaml:"helpPageHideCommercialContent"`
	// Alternate support URL for help page and help dropdown.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#help_page_support_url ApplicationSettings#help_page_support_url}
	HelpPageSupportUrl *string `field:"optional" json:"helpPageSupportUrl" yaml:"helpPageSupportUrl"`
	// Custom text displayed on the help page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#help_page_text ApplicationSettings#help_page_text}
	HelpPageText *string `field:"optional" json:"helpPageText" yaml:"helpPageText"`
	// GitLab server administrator information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#help_text ApplicationSettings#help_text}
	HelpText *string `field:"optional" json:"helpText" yaml:"helpText"`
	// Do not display offers from third parties in GitLab.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#hide_third_party_offers ApplicationSettings#hide_third_party_offers}
	HideThirdPartyOffers interface{} `field:"optional" json:"hideThirdPartyOffers" yaml:"hideThirdPartyOffers"`
	// Redirect to this URL when not logged in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#home_page_url ApplicationSettings#home_page_url}
	HomePageUrl *string `field:"optional" json:"homePageUrl" yaml:"homePageUrl"`
	// (If enabled, requires: housekeeping_bitmaps_enabled, housekeeping_full_repack_period, housekeeping_gc_period, and housekeeping_incremental_repack_period) Enable or disable Git housekeeping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#housekeeping_enabled ApplicationSettings#housekeeping_enabled}
	HousekeepingEnabled interface{} `field:"optional" json:"housekeepingEnabled" yaml:"housekeepingEnabled"`
	// Number of Git pushes after which an incremental git repack is run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#housekeeping_full_repack_period ApplicationSettings#housekeeping_full_repack_period}
	HousekeepingFullRepackPeriod *float64 `field:"optional" json:"housekeepingFullRepackPeriod" yaml:"housekeepingFullRepackPeriod"`
	// Number of Git pushes after which git gc is run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#housekeeping_gc_period ApplicationSettings#housekeeping_gc_period}
	HousekeepingGcPeriod *float64 `field:"optional" json:"housekeepingGcPeriod" yaml:"housekeepingGcPeriod"`
	// Number of Git pushes after which an incremental git repack is run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#housekeeping_incremental_repack_period ApplicationSettings#housekeeping_incremental_repack_period}
	HousekeepingIncrementalRepackPeriod *float64 `field:"optional" json:"housekeepingIncrementalRepackPeriod" yaml:"housekeepingIncrementalRepackPeriod"`
	// Enable HTML emails.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#html_emails_enabled ApplicationSettings#html_emails_enabled}
	HtmlEmailsEnabled interface{} `field:"optional" json:"htmlEmailsEnabled" yaml:"htmlEmailsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#id ApplicationSettings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Sources to allow project import from. Valid values are: `github`, `bitbucket`, `bitbucket_server`, `fogbugz`, `git`, `gitlab_project`, `gitea`, `manifest`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#import_sources ApplicationSettings#import_sources}
	ImportSources *[]*string `field:"optional" json:"importSources" yaml:"importSources"`
	// If delete_inactive_projects is true, the time (in months) to wait before deleting inactive projects.
	//
	// Introduced in GitLab 14.10. Became operational in GitLab 15.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#inactive_projects_delete_after_months ApplicationSettings#inactive_projects_delete_after_months}
	InactiveProjectsDeleteAfterMonths *float64 `field:"optional" json:"inactiveProjectsDeleteAfterMonths" yaml:"inactiveProjectsDeleteAfterMonths"`
	// If delete_inactive_projects is true, the minimum repository size for projects to be checked for inactivity.
	//
	// Introduced in GitLab 14.10. Became operational in GitLab 15.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#inactive_projects_min_size_mb ApplicationSettings#inactive_projects_min_size_mb}
	InactiveProjectsMinSizeMb *float64 `field:"optional" json:"inactiveProjectsMinSizeMb" yaml:"inactiveProjectsMinSizeMb"`
	// If delete_inactive_projects is true, sets the time (in months) to wait before emailing maintainers that the project is scheduled be deleted because it is inactive.
	//
	// Introduced in GitLab 14.10. Became operational in GitLab 15.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#inactive_projects_send_warning_email_after_months ApplicationSettings#inactive_projects_send_warning_email_after_months}
	InactiveProjectsSendWarningEmailAfterMonths *float64 `field:"optional" json:"inactiveProjectsSendWarningEmailAfterMonths" yaml:"inactiveProjectsSendWarningEmailAfterMonths"`
	// Enable in-product marketing emails.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#in_product_marketing_emails_enabled ApplicationSettings#in_product_marketing_emails_enabled}
	InProductMarketingEmailsEnabled interface{} `field:"optional" json:"inProductMarketingEmailsEnabled" yaml:"inProductMarketingEmailsEnabled"`
	// Enable Invisible CAPTCHA spam detection during sign-up.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#invisible_captcha_enabled ApplicationSettings#invisible_captcha_enabled}
	InvisibleCaptchaEnabled interface{} `field:"optional" json:"invisibleCaptchaEnabled" yaml:"invisibleCaptchaEnabled"`
	// Max number of issue creation requests per minute per user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#issues_create_limit ApplicationSettings#issues_create_limit}
	IssuesCreateLimit *float64 `field:"optional" json:"issuesCreateLimit" yaml:"issuesCreateLimit"`
	// Prevent the deletion of the artifacts from the most recent successful jobs, regardless of the expiry time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#keep_latest_artifact ApplicationSettings#keep_latest_artifact}
	KeepLatestArtifact interface{} `field:"optional" json:"keepLatestArtifact" yaml:"keepLatestArtifact"`
	// Increase this value when any cached Markdown should be invalidated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#local_markdown_version ApplicationSettings#local_markdown_version}
	LocalMarkdownVersion *float64 `field:"optional" json:"localMarkdownVersion" yaml:"localMarkdownVersion"`
	// Enable Mailgun event receiver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#mailgun_events_enabled ApplicationSettings#mailgun_events_enabled}
	MailgunEventsEnabled interface{} `field:"optional" json:"mailgunEventsEnabled" yaml:"mailgunEventsEnabled"`
	// The Mailgun HTTP webhook signing key for receiving events from webhook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#mailgun_signing_key ApplicationSettings#mailgun_signing_key}
	MailgunSigningKey *string `field:"optional" json:"mailgunSigningKey" yaml:"mailgunSigningKey"`
	// When instance is in maintenance mode, non-administrative users can sign in with read-only access and make read-only API requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#maintenance_mode ApplicationSettings#maintenance_mode}
	MaintenanceMode interface{} `field:"optional" json:"maintenanceMode" yaml:"maintenanceMode"`
	// Message displayed when instance is in maintenance mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#maintenance_mode_message ApplicationSettings#maintenance_mode_message}
	MaintenanceModeMessage *string `field:"optional" json:"maintenanceModeMessage" yaml:"maintenanceModeMessage"`
	// Maximum artifacts size in MB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#max_artifacts_size ApplicationSettings#max_artifacts_size}
	MaxArtifactsSize *float64 `field:"optional" json:"maxArtifactsSize" yaml:"maxArtifactsSize"`
	// Limit attachment size in MB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#max_attachment_size ApplicationSettings#max_attachment_size}
	MaxAttachmentSize *float64 `field:"optional" json:"maxAttachmentSize" yaml:"maxAttachmentSize"`
	// Maximum export size in MB. 0 for unlimited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#max_export_size ApplicationSettings#max_export_size}
	MaxExportSize *float64 `field:"optional" json:"maxExportSize" yaml:"maxExportSize"`
	// Maximum import size in MB. 0 for unlimited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#max_import_size ApplicationSettings#max_import_size}
	MaxImportSize *float64 `field:"optional" json:"maxImportSize" yaml:"maxImportSize"`
	// Maximum number of unique repositories a user can download in the specified time period before they are banned.
	//
	// Maximum: 10,000 repositories. Introduced in GitLab 15.1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#max_number_of_repository_downloads ApplicationSettings#max_number_of_repository_downloads}
	MaxNumberOfRepositoryDownloads *float64 `field:"optional" json:"maxNumberOfRepositoryDownloads" yaml:"maxNumberOfRepositoryDownloads"`
	// Reporting time period (in seconds). Maximum: 864000 seconds (10 days). Introduced in GitLab 15.1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#max_number_of_repository_downloads_within_time_period ApplicationSettings#max_number_of_repository_downloads_within_time_period}
	MaxNumberOfRepositoryDownloadsWithinTimePeriod *float64 `field:"optional" json:"maxNumberOfRepositoryDownloadsWithinTimePeriod" yaml:"maxNumberOfRepositoryDownloadsWithinTimePeriod"`
	// Maximum size of pages repositories in MB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#max_pages_size ApplicationSettings#max_pages_size}
	MaxPagesSize *float64 `field:"optional" json:"maxPagesSize" yaml:"maxPagesSize"`
	// Maximum allowable lifetime for access tokens in days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#max_personal_access_token_lifetime ApplicationSettings#max_personal_access_token_lifetime}
	MaxPersonalAccessTokenLifetime *float64 `field:"optional" json:"maxPersonalAccessTokenLifetime" yaml:"maxPersonalAccessTokenLifetime"`
	// Maximum allowable lifetime for SSH keys in days. Introduced in GitLab 14.6.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#max_ssh_key_lifetime ApplicationSettings#max_ssh_key_lifetime}
	MaxSshKeyLifetime *float64 `field:"optional" json:"maxSshKeyLifetime" yaml:"maxSshKeyLifetime"`
	// A method call is only tracked when it takes longer than the given amount of milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#metrics_method_call_threshold ApplicationSettings#metrics_method_call_threshold}
	MetricsMethodCallThreshold *float64 `field:"optional" json:"metricsMethodCallThreshold" yaml:"metricsMethodCallThreshold"`
	// Allow repository mirroring to configured by project Maintainers. If disabled, only Administrators can configure repository mirroring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#mirror_available ApplicationSettings#mirror_available}
	MirrorAvailable interface{} `field:"optional" json:"mirrorAvailable" yaml:"mirrorAvailable"`
	// Minimum capacity to be available before scheduling more mirrors preemptively.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#mirror_capacity_threshold ApplicationSettings#mirror_capacity_threshold}
	MirrorCapacityThreshold *float64 `field:"optional" json:"mirrorCapacityThreshold" yaml:"mirrorCapacityThreshold"`
	// Maximum number of mirrors that can be synchronizing at the same time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#mirror_max_capacity ApplicationSettings#mirror_max_capacity}
	MirrorMaxCapacity *float64 `field:"optional" json:"mirrorMaxCapacity" yaml:"mirrorMaxCapacity"`
	// Maximum time (in minutes) between updates that a mirror can have when scheduled to synchronize.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#mirror_max_delay ApplicationSettings#mirror_max_delay}
	MirrorMaxDelay *float64 `field:"optional" json:"mirrorMaxDelay" yaml:"mirrorMaxDelay"`
	// Use npmjs.org as a default remote repository when the package is not found in the GitLab Package Registry for npm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#npm_package_requests_forwarding ApplicationSettings#npm_package_requests_forwarding}
	NpmPackageRequestsForwarding interface{} `field:"optional" json:"npmPackageRequestsForwarding" yaml:"npmPackageRequestsForwarding"`
	// Define a list of trusted domains or IP addresses to which local requests are allowed when local requests for hooks and services are disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#outbound_local_requests_whitelist ApplicationSettings#outbound_local_requests_whitelist}
	OutboundLocalRequestsWhitelist *[]*string `field:"optional" json:"outboundLocalRequestsWhitelist" yaml:"outboundLocalRequestsWhitelist"`
	// Number of workers assigned to the packages cleanup policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#package_registry_cleanup_policies_worker_capacity ApplicationSettings#package_registry_cleanup_policies_worker_capacity}
	PackageRegistryCleanupPoliciesWorkerCapacity *float64 `field:"optional" json:"packageRegistryCleanupPoliciesWorkerCapacity" yaml:"packageRegistryCleanupPoliciesWorkerCapacity"`
	// Require users to prove ownership of custom domains.
	//
	// Domain verification is an essential security measure for public GitLab sites. Users are required to demonstrate they control a domain before it is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#pages_domain_verification_enabled ApplicationSettings#pages_domain_verification_enabled}
	PagesDomainVerificationEnabled interface{} `field:"optional" json:"pagesDomainVerificationEnabled" yaml:"pagesDomainVerificationEnabled"`
	// Enable authentication for Git over HTTP(S) via a GitLab account password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#password_authentication_enabled_for_git ApplicationSettings#password_authentication_enabled_for_git}
	PasswordAuthenticationEnabledForGit interface{} `field:"optional" json:"passwordAuthenticationEnabledForGit" yaml:"passwordAuthenticationEnabledForGit"`
	// Enable authentication for the web interface via a GitLab account password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#password_authentication_enabled_for_web ApplicationSettings#password_authentication_enabled_for_web}
	PasswordAuthenticationEnabledForWeb interface{} `field:"optional" json:"passwordAuthenticationEnabledForWeb" yaml:"passwordAuthenticationEnabledForWeb"`
	// Indicates whether passwords require at least one lowercase letter. Introduced in GitLab 15.1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#password_lowercase_required ApplicationSettings#password_lowercase_required}
	PasswordLowercaseRequired interface{} `field:"optional" json:"passwordLowercaseRequired" yaml:"passwordLowercaseRequired"`
	// Indicates whether passwords require at least one number. Introduced in GitLab 15.1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#password_number_required ApplicationSettings#password_number_required}
	PasswordNumberRequired interface{} `field:"optional" json:"passwordNumberRequired" yaml:"passwordNumberRequired"`
	// Indicates whether passwords require at least one symbol character. Introduced in GitLab 15.1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#password_symbol_required ApplicationSettings#password_symbol_required}
	PasswordSymbolRequired interface{} `field:"optional" json:"passwordSymbolRequired" yaml:"passwordSymbolRequired"`
	// Indicates whether passwords require at least one uppercase letter. Introduced in GitLab 15.1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#password_uppercase_required ApplicationSettings#password_uppercase_required}
	PasswordUppercaseRequired interface{} `field:"optional" json:"passwordUppercaseRequired" yaml:"passwordUppercaseRequired"`
	// Path of the group that is allowed to toggle the performance bar.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#performance_bar_allowed_group_path ApplicationSettings#performance_bar_allowed_group_path}
	PerformanceBarAllowedGroupPath *string `field:"optional" json:"performanceBarAllowedGroupPath" yaml:"performanceBarAllowedGroupPath"`
	// Prefix for all generated personal access tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#personal_access_token_prefix ApplicationSettings#personal_access_token_prefix}
	PersonalAccessTokenPrefix *string `field:"optional" json:"personalAccessTokenPrefix" yaml:"personalAccessTokenPrefix"`
	// Maximum number of pipeline creation requests per minute per user and commit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#pipeline_limit_per_project_user_sha ApplicationSettings#pipeline_limit_per_project_user_sha}
	PipelineLimitPerProjectUserSha *float64 `field:"optional" json:"pipelineLimitPerProjectUserSha" yaml:"pipelineLimitPerProjectUserSha"`
	// (If enabled, requires: plantuml_url) Enable PlantUML integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#plantuml_enabled ApplicationSettings#plantuml_enabled}
	PlantumlEnabled interface{} `field:"optional" json:"plantumlEnabled" yaml:"plantumlEnabled"`
	// The PlantUML instance URL for integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#plantuml_url ApplicationSettings#plantuml_url}
	PlantumlUrl *string `field:"optional" json:"plantumlUrl" yaml:"plantumlUrl"`
	// Interval multiplier used by endpoints that perform polling. Set to 0 to disable polling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#polling_interval_multiplier ApplicationSettings#polling_interval_multiplier}
	PollingIntervalMultiplier *float64 `field:"optional" json:"pollingIntervalMultiplier" yaml:"pollingIntervalMultiplier"`
	// Enable project export.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#project_export_enabled ApplicationSettings#project_export_enabled}
	ProjectExportEnabled interface{} `field:"optional" json:"projectExportEnabled" yaml:"projectExportEnabled"`
	// Enable Prometheus metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#prometheus_metrics_enabled ApplicationSettings#prometheus_metrics_enabled}
	PrometheusMetricsEnabled interface{} `field:"optional" json:"prometheusMetricsEnabled" yaml:"prometheusMetricsEnabled"`
	// CI/CD variables are protected by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#protected_ci_variables ApplicationSettings#protected_ci_variables}
	ProtectedCiVariables interface{} `field:"optional" json:"protectedCiVariables" yaml:"protectedCiVariables"`
	// Number of changes (branches or tags) in a single push to determine whether individual push events or bulk push events are created.
	//
	// Bulk push events are created if it surpasses that value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#push_event_activities_limit ApplicationSettings#push_event_activities_limit}
	PushEventActivitiesLimit *float64 `field:"optional" json:"pushEventActivitiesLimit" yaml:"pushEventActivitiesLimit"`
	// Number of changes (branches or tags) in a single push to determine whether webhooks and services fire or not.
	//
	// Webhooks and services aren’t submitted if it surpasses that value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#push_event_hooks_limit ApplicationSettings#push_event_hooks_limit}
	PushEventHooksLimit *float64 `field:"optional" json:"pushEventHooksLimit" yaml:"pushEventHooksLimit"`
	// Use pypi.org as a default remote repository when the package is not found in the GitLab Package Registry for PyPI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#pypi_package_requests_forwarding ApplicationSettings#pypi_package_requests_forwarding}
	PypiPackageRequestsForwarding interface{} `field:"optional" json:"pypiPackageRequestsForwarding" yaml:"pypiPackageRequestsForwarding"`
	// When rate limiting is enabled via the throttle_* settings, send this plain text response when a rate limit is exceeded.
	//
	// ‘Retry later’ is sent if this is blank.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#rate_limiting_response_text ApplicationSettings#rate_limiting_response_text}
	RateLimitingResponseText *string `field:"optional" json:"rateLimitingResponseText" yaml:"rateLimitingResponseText"`
	// Max number of requests per minute for each raw path. To disable throttling set to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#raw_blob_request_limit ApplicationSettings#raw_blob_request_limit}
	RawBlobRequestLimit *float64 `field:"optional" json:"rawBlobRequestLimit" yaml:"rawBlobRequestLimit"`
	// (If enabled, requires: recaptcha_private_key and recaptcha_site_key) Enable reCAPTCHA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#recaptcha_enabled ApplicationSettings#recaptcha_enabled}
	RecaptchaEnabled interface{} `field:"optional" json:"recaptchaEnabled" yaml:"recaptchaEnabled"`
	// Private key for reCAPTCHA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#recaptcha_private_key ApplicationSettings#recaptcha_private_key}
	RecaptchaPrivateKey *string `field:"optional" json:"recaptchaPrivateKey" yaml:"recaptchaPrivateKey"`
	// Site key for reCAPTCHA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#recaptcha_site_key ApplicationSettings#recaptcha_site_key}
	RecaptchaSiteKey *string `field:"optional" json:"recaptchaSiteKey" yaml:"recaptchaSiteKey"`
	// Maximum push size (MB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#receive_max_input_size ApplicationSettings#receive_max_input_size}
	ReceiveMaxInputSize *float64 `field:"optional" json:"receiveMaxInputSize" yaml:"receiveMaxInputSize"`
	// GitLab periodically runs git fsck in all project and wiki repositories to look for silent disk corruption issues.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#repository_checks_enabled ApplicationSettings#repository_checks_enabled}
	RepositoryChecksEnabled interface{} `field:"optional" json:"repositoryChecksEnabled" yaml:"repositoryChecksEnabled"`
	// Size limit per repository (MB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#repository_size_limit ApplicationSettings#repository_size_limit}
	RepositorySizeLimit *float64 `field:"optional" json:"repositorySizeLimit" yaml:"repositorySizeLimit"`
	// (GitLab 13.0 and earlier) List of names of enabled storage paths, taken from gitlab.yml. New projects are created in one of these stores, chosen at random.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#repository_storages ApplicationSettings#repository_storages}
	RepositoryStorages *[]*string `field:"optional" json:"repositoryStorages" yaml:"repositoryStorages"`
	// (GitLab 13.1 and later) Hash of names of taken from gitlab.yml to weights. New projects are created in one of these stores, chosen by a weighted random selection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#repository_storages_weighted ApplicationSettings#repository_storages_weighted}
	RepositoryStoragesWeighted *map[string]*float64 `field:"optional" json:"repositoryStoragesWeighted" yaml:"repositoryStoragesWeighted"`
	// When enabled, any user that signs up for an account using the registration form is placed under a Pending approval state and has to be explicitly approved by an administrator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#require_admin_approval_after_user_signup ApplicationSettings#require_admin_approval_after_user_signup}
	RequireAdminApprovalAfterUserSignup interface{} `field:"optional" json:"requireAdminApprovalAfterUserSignup" yaml:"requireAdminApprovalAfterUserSignup"`
	// (If enabled, requires: two_factor_grace_period) Require all users to set up Two-factor authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#require_two_factor_authentication ApplicationSettings#require_two_factor_authentication}
	RequireTwoFactorAuthentication interface{} `field:"optional" json:"requireTwoFactorAuthentication" yaml:"requireTwoFactorAuthentication"`
	// Selected levels cannot be used by non-Administrator users for groups, projects or snippets.
	//
	// Can take private, internal and public as a parameter. Null means there is no restriction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#restricted_visibility_levels ApplicationSettings#restricted_visibility_levels}
	RestrictedVisibilityLevels *[]*string `field:"optional" json:"restrictedVisibilityLevels" yaml:"restrictedVisibilityLevels"`
	// The minimum allowed bit length of an uploaded RSA key. 0 means no restriction. -1 disables RSA keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#rsa_key_restriction ApplicationSettings#rsa_key_restriction}
	RsaKeyRestriction *float64 `field:"optional" json:"rsaKeyRestriction" yaml:"rsaKeyRestriction"`
	// Max number of requests per minute for performing a search while authenticated. To disable throttling set to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#search_rate_limit ApplicationSettings#search_rate_limit}
	SearchRateLimit *float64 `field:"optional" json:"searchRateLimit" yaml:"searchRateLimit"`
	// Max number of requests per minute for performing a search while unauthenticated. To disable throttling set to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#search_rate_limit_unauthenticated ApplicationSettings#search_rate_limit_unauthenticated}
	SearchRateLimitUnauthenticated *float64 `field:"optional" json:"searchRateLimitUnauthenticated" yaml:"searchRateLimitUnauthenticated"`
	// Send confirmation email on sign-up.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#send_user_confirmation_email ApplicationSettings#send_user_confirmation_email}
	SendUserConfirmationEmail interface{} `field:"optional" json:"sendUserConfirmationEmail" yaml:"sendUserConfirmationEmail"`
	// Session duration in minutes. GitLab restart is required to apply changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#session_expire_delay ApplicationSettings#session_expire_delay}
	SessionExpireDelay *float64 `field:"optional" json:"sessionExpireDelay" yaml:"sessionExpireDelay"`
	// (If enabled, requires: shared_runners_text and shared_runners_minutes) Enable shared runners for new projects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#shared_runners_enabled ApplicationSettings#shared_runners_enabled}
	SharedRunnersEnabled interface{} `field:"optional" json:"sharedRunnersEnabled" yaml:"sharedRunnersEnabled"`
	// Set the maximum number of CI/CD minutes that a group can use on shared runners per month.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#shared_runners_minutes ApplicationSettings#shared_runners_minutes}
	SharedRunnersMinutes *float64 `field:"optional" json:"sharedRunnersMinutes" yaml:"sharedRunnersMinutes"`
	// Shared runners text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#shared_runners_text ApplicationSettings#shared_runners_text}
	SharedRunnersText *string `field:"optional" json:"sharedRunnersText" yaml:"sharedRunnersText"`
	// The threshold in bytes at which Sidekiq jobs are compressed before being stored in Redis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#sidekiq_job_limiter_compression_threshold_bytes ApplicationSettings#sidekiq_job_limiter_compression_threshold_bytes}
	SidekiqJobLimiterCompressionThresholdBytes *float64 `field:"optional" json:"sidekiqJobLimiterCompressionThresholdBytes" yaml:"sidekiqJobLimiterCompressionThresholdBytes"`
	// The threshold in bytes at which Sidekiq jobs are rejected. 0 means do not reject any job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#sidekiq_job_limiter_limit_bytes ApplicationSettings#sidekiq_job_limiter_limit_bytes}
	SidekiqJobLimiterLimitBytes *float64 `field:"optional" json:"sidekiqJobLimiterLimitBytes" yaml:"sidekiqJobLimiterLimitBytes"`
	// track or compress. Sets the behavior for Sidekiq job size limits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#sidekiq_job_limiter_mode ApplicationSettings#sidekiq_job_limiter_mode}
	SidekiqJobLimiterMode *string `field:"optional" json:"sidekiqJobLimiterMode" yaml:"sidekiqJobLimiterMode"`
	// Text on the login page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#sign_in_text ApplicationSettings#sign_in_text}
	SignInText *string `field:"optional" json:"signInText" yaml:"signInText"`
	// Enable registration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#signup_enabled ApplicationSettings#signup_enabled}
	SignupEnabled interface{} `field:"optional" json:"signupEnabled" yaml:"signupEnabled"`
	// (If enabled, requires: slack_app_id, slack_app_secret and slack_app_secret) Enable Slack app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#slack_app_enabled ApplicationSettings#slack_app_enabled}
	SlackAppEnabled interface{} `field:"optional" json:"slackAppEnabled" yaml:"slackAppEnabled"`
	// The app ID of the Slack-app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#slack_app_id ApplicationSettings#slack_app_id}
	SlackAppId *string `field:"optional" json:"slackAppId" yaml:"slackAppId"`
	// The app secret of the Slack-app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#slack_app_secret ApplicationSettings#slack_app_secret}
	SlackAppSecret *string `field:"optional" json:"slackAppSecret" yaml:"slackAppSecret"`
	// The signing secret of the Slack-app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#slack_app_signing_secret ApplicationSettings#slack_app_signing_secret}
	SlackAppSigningSecret *string `field:"optional" json:"slackAppSigningSecret" yaml:"slackAppSigningSecret"`
	// The verification token of the Slack-app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#slack_app_verification_token ApplicationSettings#slack_app_verification_token}
	SlackAppVerificationToken *string `field:"optional" json:"slackAppVerificationToken" yaml:"slackAppVerificationToken"`
	// Max snippet content size in bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#snippet_size_limit ApplicationSettings#snippet_size_limit}
	SnippetSizeLimit *float64 `field:"optional" json:"snippetSizeLimit" yaml:"snippetSizeLimit"`
	// The Snowplow site name / application ID. (for example, gitlab).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#snowplow_app_id ApplicationSettings#snowplow_app_id}
	SnowplowAppId *string `field:"optional" json:"snowplowAppId" yaml:"snowplowAppId"`
	// The Snowplow collector hostname. (for example, snowplow.trx.gitlab.net).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#snowplow_collector_hostname ApplicationSettings#snowplow_collector_hostname}
	SnowplowCollectorHostname *string `field:"optional" json:"snowplowCollectorHostname" yaml:"snowplowCollectorHostname"`
	// The Snowplow cookie domain. (for example, .gitlab.com).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#snowplow_cookie_domain ApplicationSettings#snowplow_cookie_domain}
	SnowplowCookieDomain *string `field:"optional" json:"snowplowCookieDomain" yaml:"snowplowCookieDomain"`
	// Enable snowplow tracking.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#snowplow_enabled ApplicationSettings#snowplow_enabled}
	SnowplowEnabled interface{} `field:"optional" json:"snowplowEnabled" yaml:"snowplowEnabled"`
	// Enables Sourcegraph integration. If enabled, requires sourcegraph_url.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#sourcegraph_enabled ApplicationSettings#sourcegraph_enabled}
	SourcegraphEnabled interface{} `field:"optional" json:"sourcegraphEnabled" yaml:"sourcegraphEnabled"`
	// Blocks Sourcegraph from being loaded on private and internal projects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#sourcegraph_public_only ApplicationSettings#sourcegraph_public_only}
	SourcegraphPublicOnly interface{} `field:"optional" json:"sourcegraphPublicOnly" yaml:"sourcegraphPublicOnly"`
	// The Sourcegraph instance URL for integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#sourcegraph_url ApplicationSettings#sourcegraph_url}
	SourcegraphUrl *string `field:"optional" json:"sourcegraphUrl" yaml:"sourcegraphUrl"`
	// API key used by GitLab for accessing the Spam Check service endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#spam_check_api_key ApplicationSettings#spam_check_api_key}
	SpamCheckApiKey *string `field:"optional" json:"spamCheckApiKey" yaml:"spamCheckApiKey"`
	// Enables spam checking using external Spam Check API endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#spam_check_endpoint_enabled ApplicationSettings#spam_check_endpoint_enabled}
	SpamCheckEndpointEnabled interface{} `field:"optional" json:"spamCheckEndpointEnabled" yaml:"spamCheckEndpointEnabled"`
	// URL of the external Spamcheck service endpoint.
	//
	// Valid URI schemes are grpc or tls. Specifying tls forces communication to be encrypted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#spam_check_endpoint_url ApplicationSettings#spam_check_endpoint_url}
	SpamCheckEndpointUrl *string `field:"optional" json:"spamCheckEndpointUrl" yaml:"spamCheckEndpointUrl"`
	// Enable pipeline suggestion banner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#suggest_pipeline_enabled ApplicationSettings#suggest_pipeline_enabled}
	SuggestPipelineEnabled interface{} `field:"optional" json:"suggestPipelineEnabled" yaml:"suggestPipelineEnabled"`
	// Maximum time for web terminal websocket connection (in seconds). Set to 0 for unlimited time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#terminal_max_session_time ApplicationSettings#terminal_max_session_time}
	TerminalMaxSessionTime *float64 `field:"optional" json:"terminalMaxSessionTime" yaml:"terminalMaxSessionTime"`
	// (Required by: enforce_terms) Markdown content for the ToS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#terms ApplicationSettings#terms}
	Terms *string `field:"optional" json:"terms" yaml:"terms"`
	// (If enabled, requires: throttle_authenticated_api_period_in_seconds and throttle_authenticated_api_requests_per_period) Enable authenticated API request rate limit.
	//
	// Helps reduce request volume (for example, from crawlers or abusive bots).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#throttle_authenticated_api_enabled ApplicationSettings#throttle_authenticated_api_enabled}
	ThrottleAuthenticatedApiEnabled interface{} `field:"optional" json:"throttleAuthenticatedApiEnabled" yaml:"throttleAuthenticatedApiEnabled"`
	// Rate limit period (in seconds).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#throttle_authenticated_api_period_in_seconds ApplicationSettings#throttle_authenticated_api_period_in_seconds}
	ThrottleAuthenticatedApiPeriodInSeconds *float64 `field:"optional" json:"throttleAuthenticatedApiPeriodInSeconds" yaml:"throttleAuthenticatedApiPeriodInSeconds"`
	// Maximum requests per period per user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#throttle_authenticated_api_requests_per_period ApplicationSettings#throttle_authenticated_api_requests_per_period}
	ThrottleAuthenticatedApiRequestsPerPeriod *float64 `field:"optional" json:"throttleAuthenticatedApiRequestsPerPeriod" yaml:"throttleAuthenticatedApiRequestsPerPeriod"`
	// (If enabled, requires: throttle_authenticated_packages_api_period_in_seconds and throttle_authenticated_packages_api_requests_per_period) Enable authenticated API request rate limit.
	//
	// Helps reduce request volume (for example, from crawlers or abusive bots). View Package Registry rate limits for more details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#throttle_authenticated_packages_api_enabled ApplicationSettings#throttle_authenticated_packages_api_enabled}
	ThrottleAuthenticatedPackagesApiEnabled interface{} `field:"optional" json:"throttleAuthenticatedPackagesApiEnabled" yaml:"throttleAuthenticatedPackagesApiEnabled"`
	// Rate limit period (in seconds). View Package Registry rate limits for more details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#throttle_authenticated_packages_api_period_in_seconds ApplicationSettings#throttle_authenticated_packages_api_period_in_seconds}
	ThrottleAuthenticatedPackagesApiPeriodInSeconds *float64 `field:"optional" json:"throttleAuthenticatedPackagesApiPeriodInSeconds" yaml:"throttleAuthenticatedPackagesApiPeriodInSeconds"`
	// Maximum requests per period per user. View Package Registry rate limits for more details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#throttle_authenticated_packages_api_requests_per_period ApplicationSettings#throttle_authenticated_packages_api_requests_per_period}
	ThrottleAuthenticatedPackagesApiRequestsPerPeriod *float64 `field:"optional" json:"throttleAuthenticatedPackagesApiRequestsPerPeriod" yaml:"throttleAuthenticatedPackagesApiRequestsPerPeriod"`
	// (If enabled, requires: throttle_authenticated_web_period_in_seconds and throttle_authenticated_web_requests_per_period) Enable authenticated web request rate limit.
	//
	// Helps reduce request volume (for example, from crawlers or abusive bots).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#throttle_authenticated_web_enabled ApplicationSettings#throttle_authenticated_web_enabled}
	ThrottleAuthenticatedWebEnabled interface{} `field:"optional" json:"throttleAuthenticatedWebEnabled" yaml:"throttleAuthenticatedWebEnabled"`
	// Rate limit period (in seconds).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#throttle_authenticated_web_period_in_seconds ApplicationSettings#throttle_authenticated_web_period_in_seconds}
	ThrottleAuthenticatedWebPeriodInSeconds *float64 `field:"optional" json:"throttleAuthenticatedWebPeriodInSeconds" yaml:"throttleAuthenticatedWebPeriodInSeconds"`
	// Maximum requests per period per user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#throttle_authenticated_web_requests_per_period ApplicationSettings#throttle_authenticated_web_requests_per_period}
	ThrottleAuthenticatedWebRequestsPerPeriod *float64 `field:"optional" json:"throttleAuthenticatedWebRequestsPerPeriod" yaml:"throttleAuthenticatedWebRequestsPerPeriod"`
	// (If enabled, requires: throttle_unauthenticated_api_period_in_seconds and throttle_unauthenticated_api_requests_per_period) Enable unauthenticated API request rate limit.
	//
	// Helps reduce request volume (for example, from crawlers or abusive bots).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#throttle_unauthenticated_api_enabled ApplicationSettings#throttle_unauthenticated_api_enabled}
	ThrottleUnauthenticatedApiEnabled interface{} `field:"optional" json:"throttleUnauthenticatedApiEnabled" yaml:"throttleUnauthenticatedApiEnabled"`
	// Rate limit period in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#throttle_unauthenticated_api_period_in_seconds ApplicationSettings#throttle_unauthenticated_api_period_in_seconds}
	ThrottleUnauthenticatedApiPeriodInSeconds *float64 `field:"optional" json:"throttleUnauthenticatedApiPeriodInSeconds" yaml:"throttleUnauthenticatedApiPeriodInSeconds"`
	// Max requests per period per IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#throttle_unauthenticated_api_requests_per_period ApplicationSettings#throttle_unauthenticated_api_requests_per_period}
	ThrottleUnauthenticatedApiRequestsPerPeriod *float64 `field:"optional" json:"throttleUnauthenticatedApiRequestsPerPeriod" yaml:"throttleUnauthenticatedApiRequestsPerPeriod"`
	// (If enabled, requires: throttle_unauthenticated_packages_api_period_in_seconds and throttle_unauthenticated_packages_api_requests_per_period) Enable authenticated API request rate limit.
	//
	// Helps reduce request volume (for example, from crawlers or abusive bots). View Package Registry rate limits for more details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#throttle_unauthenticated_packages_api_enabled ApplicationSettings#throttle_unauthenticated_packages_api_enabled}
	ThrottleUnauthenticatedPackagesApiEnabled interface{} `field:"optional" json:"throttleUnauthenticatedPackagesApiEnabled" yaml:"throttleUnauthenticatedPackagesApiEnabled"`
	// Rate limit period (in seconds). View Package Registry rate limits for more details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#throttle_unauthenticated_packages_api_period_in_seconds ApplicationSettings#throttle_unauthenticated_packages_api_period_in_seconds}
	ThrottleUnauthenticatedPackagesApiPeriodInSeconds *float64 `field:"optional" json:"throttleUnauthenticatedPackagesApiPeriodInSeconds" yaml:"throttleUnauthenticatedPackagesApiPeriodInSeconds"`
	// Maximum requests per period per user. View Package Registry rate limits for more details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#throttle_unauthenticated_packages_api_requests_per_period ApplicationSettings#throttle_unauthenticated_packages_api_requests_per_period}
	ThrottleUnauthenticatedPackagesApiRequestsPerPeriod *float64 `field:"optional" json:"throttleUnauthenticatedPackagesApiRequestsPerPeriod" yaml:"throttleUnauthenticatedPackagesApiRequestsPerPeriod"`
	// (If enabled, requires: throttle_unauthenticated_web_period_in_seconds and throttle_unauthenticated_web_requests_per_period) Enable unauthenticated web request rate limit.
	//
	// Helps reduce request volume (for example, from crawlers or abusive bots).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#throttle_unauthenticated_web_enabled ApplicationSettings#throttle_unauthenticated_web_enabled}
	ThrottleUnauthenticatedWebEnabled interface{} `field:"optional" json:"throttleUnauthenticatedWebEnabled" yaml:"throttleUnauthenticatedWebEnabled"`
	// Rate limit period in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#throttle_unauthenticated_web_period_in_seconds ApplicationSettings#throttle_unauthenticated_web_period_in_seconds}
	ThrottleUnauthenticatedWebPeriodInSeconds *float64 `field:"optional" json:"throttleUnauthenticatedWebPeriodInSeconds" yaml:"throttleUnauthenticatedWebPeriodInSeconds"`
	// Max requests per period per IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#throttle_unauthenticated_web_requests_per_period ApplicationSettings#throttle_unauthenticated_web_requests_per_period}
	ThrottleUnauthenticatedWebRequestsPerPeriod *float64 `field:"optional" json:"throttleUnauthenticatedWebRequestsPerPeriod" yaml:"throttleUnauthenticatedWebRequestsPerPeriod"`
	// Limit display of time tracking units to hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#time_tracking_limit_to_hours ApplicationSettings#time_tracking_limit_to_hours}
	TimeTrackingLimitToHours interface{} `field:"optional" json:"timeTrackingLimitToHours" yaml:"timeTrackingLimitToHours"`
	// Amount of time (in hours) that users are allowed to skip forced configuration of two-factor authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#two_factor_grace_period ApplicationSettings#two_factor_grace_period}
	TwoFactorGracePeriod *float64 `field:"optional" json:"twoFactorGracePeriod" yaml:"twoFactorGracePeriod"`
	// (If enabled, requires: unique_ips_limit_per_user and unique_ips_limit_time_window) Limit sign in from multiple IPs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#unique_ips_limit_enabled ApplicationSettings#unique_ips_limit_enabled}
	UniqueIpsLimitEnabled interface{} `field:"optional" json:"uniqueIpsLimitEnabled" yaml:"uniqueIpsLimitEnabled"`
	// Maximum number of IPs per user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#unique_ips_limit_per_user ApplicationSettings#unique_ips_limit_per_user}
	UniqueIpsLimitPerUser *float64 `field:"optional" json:"uniqueIpsLimitPerUser" yaml:"uniqueIpsLimitPerUser"`
	// How many seconds an IP is counted towards the limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#unique_ips_limit_time_window ApplicationSettings#unique_ips_limit_time_window}
	UniqueIpsLimitTimeWindow *float64 `field:"optional" json:"uniqueIpsLimitTimeWindow" yaml:"uniqueIpsLimitTimeWindow"`
	// Every week GitLab reports license usage back to GitLab, Inc.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#usage_ping_enabled ApplicationSettings#usage_ping_enabled}
	UsagePingEnabled interface{} `field:"optional" json:"usagePingEnabled" yaml:"usagePingEnabled"`
	// Send an email to users upon account deactivation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#user_deactivation_emails_enabled ApplicationSettings#user_deactivation_emails_enabled}
	UserDeactivationEmailsEnabled interface{} `field:"optional" json:"userDeactivationEmailsEnabled" yaml:"userDeactivationEmailsEnabled"`
	// Newly registered users are external by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#user_default_external ApplicationSettings#user_default_external}
	UserDefaultExternal interface{} `field:"optional" json:"userDefaultExternal" yaml:"userDefaultExternal"`
	// Specify an email address regex pattern to identify default internal users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#user_default_internal_regex ApplicationSettings#user_default_internal_regex}
	UserDefaultInternalRegex *string `field:"optional" json:"userDefaultInternalRegex" yaml:"userDefaultInternalRegex"`
	// Allow users to register any application to use GitLab as an OAuth provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#user_oauth_applications ApplicationSettings#user_oauth_applications}
	UserOauthApplications interface{} `field:"optional" json:"userOauthApplications" yaml:"userOauthApplications"`
	// When set to false disable the You won't be able to pull or push project code via SSH warning shown to users with no uploaded SSH key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#user_show_add_ssh_key_message ApplicationSettings#user_show_add_ssh_key_message}
	UserShowAddSshKeyMessage interface{} `field:"optional" json:"userShowAddSshKeyMessage" yaml:"userShowAddSshKeyMessage"`
	// Let GitLab inform you when an update is available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#version_check_enabled ApplicationSettings#version_check_enabled}
	VersionCheckEnabled interface{} `field:"optional" json:"versionCheckEnabled" yaml:"versionCheckEnabled"`
	// Live Preview (allow live previews of JavaScript projects in the Web IDE using CodeSandbox Live Preview).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#web_ide_clientside_preview_enabled ApplicationSettings#web_ide_clientside_preview_enabled}
	WebIdeClientsidePreviewEnabled interface{} `field:"optional" json:"webIdeClientsidePreviewEnabled" yaml:"webIdeClientsidePreviewEnabled"`
	// What’s new variant, possible values: all_tiers, current_tier, and disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#whats_new_variant ApplicationSettings#whats_new_variant}
	WhatsNewVariant *string `field:"optional" json:"whatsNewVariant" yaml:"whatsNewVariant"`
	// Maximum wiki page content size in bytes. The minimum value is 1024 bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.3.0/docs/resources/application_settings#wiki_page_max_content_bytes ApplicationSettings#wiki_page_max_content_bytes}
	WikiPageMaxContentBytes *float64 `field:"optional" json:"wikiPageMaxContentBytes" yaml:"wikiPageMaxContentBytes"`
}

