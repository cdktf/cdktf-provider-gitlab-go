// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationsettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v12/applicationsettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.9.1/docs/resources/application_settings gitlab_application_settings}.
type ApplicationSettings interface {
	cdktf.TerraformResource
	AbuseNotificationEmail() *string
	SetAbuseNotificationEmail(val *string)
	AbuseNotificationEmailInput() *string
	AdminMode() interface{}
	SetAdminMode(val interface{})
	AdminModeInput() interface{}
	AfterSignOutPath() *string
	SetAfterSignOutPath(val *string)
	AfterSignOutPathInput() *string
	AfterSignUpText() *string
	SetAfterSignUpText(val *string)
	AfterSignUpTextInput() *string
	AkismetApiKey() *string
	SetAkismetApiKey(val *string)
	AkismetApiKeyInput() *string
	AkismetEnabled() interface{}
	SetAkismetEnabled(val interface{})
	AkismetEnabledInput() interface{}
	AllowGroupOwnersToManageLdap() interface{}
	SetAllowGroupOwnersToManageLdap(val interface{})
	AllowGroupOwnersToManageLdapInput() interface{}
	AllowLocalRequestsFromSystemHooks() interface{}
	SetAllowLocalRequestsFromSystemHooks(val interface{})
	AllowLocalRequestsFromSystemHooksInput() interface{}
	AllowLocalRequestsFromWebHooksAndServices() interface{}
	SetAllowLocalRequestsFromWebHooksAndServices(val interface{})
	AllowLocalRequestsFromWebHooksAndServicesInput() interface{}
	ArchiveBuildsInHumanReadable() *string
	SetArchiveBuildsInHumanReadable(val *string)
	ArchiveBuildsInHumanReadableInput() *string
	AssetProxyAllowlist() *[]*string
	SetAssetProxyAllowlist(val *[]*string)
	AssetProxyAllowlistInput() *[]*string
	AssetProxyEnabled() interface{}
	SetAssetProxyEnabled(val interface{})
	AssetProxyEnabledInput() interface{}
	AssetProxySecretKey() *string
	SetAssetProxySecretKey(val *string)
	AssetProxySecretKeyInput() *string
	AssetProxyUrl() *string
	SetAssetProxyUrl(val *string)
	AssetProxyUrlInput() *string
	AuthorizedKeysEnabled() interface{}
	SetAuthorizedKeysEnabled(val interface{})
	AuthorizedKeysEnabledInput() interface{}
	AutoDevopsDomain() *string
	SetAutoDevopsDomain(val *string)
	AutoDevopsDomainInput() *string
	AutoDevopsEnabled() interface{}
	SetAutoDevopsEnabled(val interface{})
	AutoDevopsEnabledInput() interface{}
	AutomaticPurchasedStorageAllocation() interface{}
	SetAutomaticPurchasedStorageAllocation(val interface{})
	AutomaticPurchasedStorageAllocationInput() interface{}
	CanCreateGroup() interface{}
	SetCanCreateGroup(val interface{})
	CanCreateGroupInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CheckNamespacePlan() interface{}
	SetCheckNamespacePlan(val interface{})
	CheckNamespacePlanInput() interface{}
	CommitEmailHostname() *string
	SetCommitEmailHostname(val *string)
	CommitEmailHostnameInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerExpirationPoliciesEnableHistoricEntries() interface{}
	SetContainerExpirationPoliciesEnableHistoricEntries(val interface{})
	ContainerExpirationPoliciesEnableHistoricEntriesInput() interface{}
	ContainerRegistryCleanupTagsServiceMaxListSize() *float64
	SetContainerRegistryCleanupTagsServiceMaxListSize(val *float64)
	ContainerRegistryCleanupTagsServiceMaxListSizeInput() *float64
	ContainerRegistryDeleteTagsServiceTimeout() *float64
	SetContainerRegistryDeleteTagsServiceTimeout(val *float64)
	ContainerRegistryDeleteTagsServiceTimeoutInput() *float64
	ContainerRegistryExpirationPoliciesCaching() interface{}
	SetContainerRegistryExpirationPoliciesCaching(val interface{})
	ContainerRegistryExpirationPoliciesCachingInput() interface{}
	ContainerRegistryExpirationPoliciesWorkerCapacity() *float64
	SetContainerRegistryExpirationPoliciesWorkerCapacity(val *float64)
	ContainerRegistryExpirationPoliciesWorkerCapacityInput() *float64
	ContainerRegistryTokenExpireDelay() *float64
	SetContainerRegistryTokenExpireDelay(val *float64)
	ContainerRegistryTokenExpireDelayInput() *float64
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DeactivateDormantUsers() interface{}
	SetDeactivateDormantUsers(val interface{})
	DeactivateDormantUsersInput() interface{}
	DefaultArtifactsExpireIn() *string
	SetDefaultArtifactsExpireIn(val *string)
	DefaultArtifactsExpireInInput() *string
	DefaultBranchName() *string
	SetDefaultBranchName(val *string)
	DefaultBranchNameInput() *string
	DefaultBranchProtection() *float64
	SetDefaultBranchProtection(val *float64)
	DefaultBranchProtectionInput() *float64
	DefaultCiConfigPath() *string
	SetDefaultCiConfigPath(val *string)
	DefaultCiConfigPathInput() *string
	DefaultGroupVisibility() *string
	SetDefaultGroupVisibility(val *string)
	DefaultGroupVisibilityInput() *string
	DefaultProjectCreation() *float64
	SetDefaultProjectCreation(val *float64)
	DefaultProjectCreationInput() *float64
	DefaultProjectsLimit() *float64
	SetDefaultProjectsLimit(val *float64)
	DefaultProjectsLimitInput() *float64
	DefaultProjectVisibility() *string
	SetDefaultProjectVisibility(val *string)
	DefaultProjectVisibilityInput() *string
	DefaultSnippetVisibility() *string
	SetDefaultSnippetVisibility(val *string)
	DefaultSnippetVisibilityInput() *string
	DeleteInactiveProjects() interface{}
	SetDeleteInactiveProjects(val interface{})
	DeleteInactiveProjectsInput() interface{}
	DeletionAdjournedPeriod() *float64
	SetDeletionAdjournedPeriod(val *float64)
	DeletionAdjournedPeriodInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DiffMaxFiles() *float64
	SetDiffMaxFiles(val *float64)
	DiffMaxFilesInput() *float64
	DiffMaxLines() *float64
	SetDiffMaxLines(val *float64)
	DiffMaxLinesInput() *float64
	DiffMaxPatchBytes() *float64
	SetDiffMaxPatchBytes(val *float64)
	DiffMaxPatchBytesInput() *float64
	DisabledOauthSignInSources() *[]*string
	SetDisabledOauthSignInSources(val *[]*string)
	DisabledOauthSignInSourcesInput() *[]*string
	DisableFeedToken() interface{}
	SetDisableFeedToken(val interface{})
	DisableFeedTokenInput() interface{}
	DnsRebindingProtectionEnabled() interface{}
	SetDnsRebindingProtectionEnabled(val interface{})
	DnsRebindingProtectionEnabledInput() interface{}
	DomainAllowlist() *[]*string
	SetDomainAllowlist(val *[]*string)
	DomainAllowlistInput() *[]*string
	DomainDenylist() *[]*string
	SetDomainDenylist(val *[]*string)
	DomainDenylistEnabled() interface{}
	SetDomainDenylistEnabled(val interface{})
	DomainDenylistEnabledInput() interface{}
	DomainDenylistInput() *[]*string
	DsaKeyRestriction() *float64
	SetDsaKeyRestriction(val *float64)
	DsaKeyRestrictionInput() *float64
	EcdsaKeyRestriction() *float64
	SetEcdsaKeyRestriction(val *float64)
	EcdsaKeyRestrictionInput() *float64
	EcdsaSkKeyRestriction() *float64
	SetEcdsaSkKeyRestriction(val *float64)
	EcdsaSkKeyRestrictionInput() *float64
	Ed25519KeyRestriction() *float64
	SetEd25519KeyRestriction(val *float64)
	Ed25519KeyRestrictionInput() *float64
	Ed25519SkKeyRestriction() *float64
	SetEd25519SkKeyRestriction(val *float64)
	Ed25519SkKeyRestrictionInput() *float64
	EksAccessKeyId() *string
	SetEksAccessKeyId(val *string)
	EksAccessKeyIdInput() *string
	EksAccountId() *string
	SetEksAccountId(val *string)
	EksAccountIdInput() *string
	EksIntegrationEnabled() interface{}
	SetEksIntegrationEnabled(val interface{})
	EksIntegrationEnabledInput() interface{}
	EksSecretAccessKey() *string
	SetEksSecretAccessKey(val *string)
	EksSecretAccessKeyInput() *string
	ElasticsearchAws() interface{}
	SetElasticsearchAws(val interface{})
	ElasticsearchAwsAccessKey() *string
	SetElasticsearchAwsAccessKey(val *string)
	ElasticsearchAwsAccessKeyInput() *string
	ElasticsearchAwsInput() interface{}
	ElasticsearchAwsRegion() *string
	SetElasticsearchAwsRegion(val *string)
	ElasticsearchAwsRegionInput() *string
	ElasticsearchAwsSecretAccessKey() *string
	SetElasticsearchAwsSecretAccessKey(val *string)
	ElasticsearchAwsSecretAccessKeyInput() *string
	ElasticsearchIndexedFieldLengthLimit() *float64
	SetElasticsearchIndexedFieldLengthLimit(val *float64)
	ElasticsearchIndexedFieldLengthLimitInput() *float64
	ElasticsearchIndexedFileSizeLimitKb() *float64
	SetElasticsearchIndexedFileSizeLimitKb(val *float64)
	ElasticsearchIndexedFileSizeLimitKbInput() *float64
	ElasticsearchIndexing() interface{}
	SetElasticsearchIndexing(val interface{})
	ElasticsearchIndexingInput() interface{}
	ElasticsearchLimitIndexing() interface{}
	SetElasticsearchLimitIndexing(val interface{})
	ElasticsearchLimitIndexingInput() interface{}
	ElasticsearchMaxBulkConcurrency() *float64
	SetElasticsearchMaxBulkConcurrency(val *float64)
	ElasticsearchMaxBulkConcurrencyInput() *float64
	ElasticsearchMaxBulkSizeMb() *float64
	SetElasticsearchMaxBulkSizeMb(val *float64)
	ElasticsearchMaxBulkSizeMbInput() *float64
	ElasticsearchNamespaceIds() *[]*float64
	SetElasticsearchNamespaceIds(val *[]*float64)
	ElasticsearchNamespaceIdsInput() *[]*float64
	ElasticsearchPassword() *string
	SetElasticsearchPassword(val *string)
	ElasticsearchPasswordInput() *string
	ElasticsearchProjectIds() *[]*float64
	SetElasticsearchProjectIds(val *[]*float64)
	ElasticsearchProjectIdsInput() *[]*float64
	ElasticsearchSearch() interface{}
	SetElasticsearchSearch(val interface{})
	ElasticsearchSearchInput() interface{}
	ElasticsearchUrl() *[]*string
	SetElasticsearchUrl(val *[]*string)
	ElasticsearchUrlInput() *[]*string
	ElasticsearchUsername() *string
	SetElasticsearchUsername(val *string)
	ElasticsearchUsernameInput() *string
	EmailAdditionalText() *string
	SetEmailAdditionalText(val *string)
	EmailAdditionalTextInput() *string
	EmailAuthorInBody() interface{}
	SetEmailAuthorInBody(val interface{})
	EmailAuthorInBodyInput() interface{}
	EnabledGitAccessProtocol() *string
	SetEnabledGitAccessProtocol(val *string)
	EnabledGitAccessProtocolInput() *string
	EnforceNamespaceStorageLimit() interface{}
	SetEnforceNamespaceStorageLimit(val interface{})
	EnforceNamespaceStorageLimitInput() interface{}
	EnforceTerms() interface{}
	SetEnforceTerms(val interface{})
	EnforceTermsInput() interface{}
	ExternalAuthClientCert() *string
	SetExternalAuthClientCert(val *string)
	ExternalAuthClientCertInput() *string
	ExternalAuthClientKey() *string
	SetExternalAuthClientKey(val *string)
	ExternalAuthClientKeyInput() *string
	ExternalAuthClientKeyPass() *string
	SetExternalAuthClientKeyPass(val *string)
	ExternalAuthClientKeyPassInput() *string
	ExternalAuthorizationServiceDefaultLabel() *string
	SetExternalAuthorizationServiceDefaultLabel(val *string)
	ExternalAuthorizationServiceDefaultLabelInput() *string
	ExternalAuthorizationServiceEnabled() interface{}
	SetExternalAuthorizationServiceEnabled(val interface{})
	ExternalAuthorizationServiceEnabledInput() interface{}
	ExternalAuthorizationServiceTimeout() *float64
	SetExternalAuthorizationServiceTimeout(val *float64)
	ExternalAuthorizationServiceTimeoutInput() *float64
	ExternalAuthorizationServiceUrl() *string
	SetExternalAuthorizationServiceUrl(val *string)
	ExternalAuthorizationServiceUrlInput() *string
	ExternalPipelineValidationServiceTimeout() *float64
	SetExternalPipelineValidationServiceTimeout(val *float64)
	ExternalPipelineValidationServiceTimeoutInput() *float64
	ExternalPipelineValidationServiceToken() *string
	SetExternalPipelineValidationServiceToken(val *string)
	ExternalPipelineValidationServiceTokenInput() *string
	ExternalPipelineValidationServiceUrl() *string
	SetExternalPipelineValidationServiceUrl(val *string)
	ExternalPipelineValidationServiceUrlInput() *string
	FileTemplateProjectId() *float64
	SetFileTemplateProjectId(val *float64)
	FileTemplateProjectIdInput() *float64
	FirstDayOfWeek() *float64
	SetFirstDayOfWeek(val *float64)
	FirstDayOfWeekInput() *float64
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GeoNodeAllowedIps() *string
	SetGeoNodeAllowedIps(val *string)
	GeoNodeAllowedIpsInput() *string
	GeoStatusTimeout() *float64
	SetGeoStatusTimeout(val *float64)
	GeoStatusTimeoutInput() *float64
	GitalyTimeoutDefault() *float64
	SetGitalyTimeoutDefault(val *float64)
	GitalyTimeoutDefaultInput() *float64
	GitalyTimeoutFast() *float64
	SetGitalyTimeoutFast(val *float64)
	GitalyTimeoutFastInput() *float64
	GitalyTimeoutMedium() *float64
	SetGitalyTimeoutMedium(val *float64)
	GitalyTimeoutMediumInput() *float64
	GitRateLimitUsersAllowlist() *[]*string
	SetGitRateLimitUsersAllowlist(val *[]*string)
	GitRateLimitUsersAllowlistInput() *[]*string
	GitTwoFactorSessionExpiry() *float64
	SetGitTwoFactorSessionExpiry(val *float64)
	GitTwoFactorSessionExpiryInput() *float64
	GrafanaEnabled() interface{}
	SetGrafanaEnabled(val interface{})
	GrafanaEnabledInput() interface{}
	GrafanaUrl() *string
	SetGrafanaUrl(val *string)
	GrafanaUrlInput() *string
	GravatarEnabled() interface{}
	SetGravatarEnabled(val interface{})
	GravatarEnabledInput() interface{}
	GroupOwnersCanManageDefaultBranchProtection() interface{}
	SetGroupOwnersCanManageDefaultBranchProtection(val interface{})
	GroupOwnersCanManageDefaultBranchProtectionInput() interface{}
	HashedStorageEnabled() interface{}
	SetHashedStorageEnabled(val interface{})
	HashedStorageEnabledInput() interface{}
	HelpPageHideCommercialContent() interface{}
	SetHelpPageHideCommercialContent(val interface{})
	HelpPageHideCommercialContentInput() interface{}
	HelpPageSupportUrl() *string
	SetHelpPageSupportUrl(val *string)
	HelpPageSupportUrlInput() *string
	HelpPageText() *string
	SetHelpPageText(val *string)
	HelpPageTextInput() *string
	HelpText() *string
	SetHelpText(val *string)
	HelpTextInput() *string
	HideThirdPartyOffers() interface{}
	SetHideThirdPartyOffers(val interface{})
	HideThirdPartyOffersInput() interface{}
	HomePageUrl() *string
	SetHomePageUrl(val *string)
	HomePageUrlInput() *string
	HousekeepingEnabled() interface{}
	SetHousekeepingEnabled(val interface{})
	HousekeepingEnabledInput() interface{}
	HousekeepingFullRepackPeriod() *float64
	SetHousekeepingFullRepackPeriod(val *float64)
	HousekeepingFullRepackPeriodInput() *float64
	HousekeepingGcPeriod() *float64
	SetHousekeepingGcPeriod(val *float64)
	HousekeepingGcPeriodInput() *float64
	HousekeepingIncrementalRepackPeriod() *float64
	SetHousekeepingIncrementalRepackPeriod(val *float64)
	HousekeepingIncrementalRepackPeriodInput() *float64
	HousekeepingOptimizeRepositoryPeriod() *float64
	SetHousekeepingOptimizeRepositoryPeriod(val *float64)
	HousekeepingOptimizeRepositoryPeriodInput() *float64
	HtmlEmailsEnabled() interface{}
	SetHtmlEmailsEnabled(val interface{})
	HtmlEmailsEnabledInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImportSources() *[]*string
	SetImportSources(val *[]*string)
	ImportSourcesInput() *[]*string
	InactiveProjectsDeleteAfterMonths() *float64
	SetInactiveProjectsDeleteAfterMonths(val *float64)
	InactiveProjectsDeleteAfterMonthsInput() *float64
	InactiveProjectsMinSizeMb() *float64
	SetInactiveProjectsMinSizeMb(val *float64)
	InactiveProjectsMinSizeMbInput() *float64
	InactiveProjectsSendWarningEmailAfterMonths() *float64
	SetInactiveProjectsSendWarningEmailAfterMonths(val *float64)
	InactiveProjectsSendWarningEmailAfterMonthsInput() *float64
	InProductMarketingEmailsEnabled() interface{}
	SetInProductMarketingEmailsEnabled(val interface{})
	InProductMarketingEmailsEnabledInput() interface{}
	InvisibleCaptchaEnabled() interface{}
	SetInvisibleCaptchaEnabled(val interface{})
	InvisibleCaptchaEnabledInput() interface{}
	IssuesCreateLimit() *float64
	SetIssuesCreateLimit(val *float64)
	IssuesCreateLimitInput() *float64
	KeepLatestArtifact() interface{}
	SetKeepLatestArtifact(val interface{})
	KeepLatestArtifactInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalMarkdownVersion() *float64
	SetLocalMarkdownVersion(val *float64)
	LocalMarkdownVersionInput() *float64
	MailgunEventsEnabled() interface{}
	SetMailgunEventsEnabled(val interface{})
	MailgunEventsEnabledInput() interface{}
	MailgunSigningKey() *string
	SetMailgunSigningKey(val *string)
	MailgunSigningKeyInput() *string
	MaintenanceMode() interface{}
	SetMaintenanceMode(val interface{})
	MaintenanceModeInput() interface{}
	MaintenanceModeMessage() *string
	SetMaintenanceModeMessage(val *string)
	MaintenanceModeMessageInput() *string
	MaxArtifactsSize() *float64
	SetMaxArtifactsSize(val *float64)
	MaxArtifactsSizeInput() *float64
	MaxAttachmentSize() *float64
	SetMaxAttachmentSize(val *float64)
	MaxAttachmentSizeInput() *float64
	MaxExportSize() *float64
	SetMaxExportSize(val *float64)
	MaxExportSizeInput() *float64
	MaxImportSize() *float64
	SetMaxImportSize(val *float64)
	MaxImportSizeInput() *float64
	MaxNumberOfRepositoryDownloads() *float64
	SetMaxNumberOfRepositoryDownloads(val *float64)
	MaxNumberOfRepositoryDownloadsInput() *float64
	MaxNumberOfRepositoryDownloadsWithinTimePeriod() *float64
	SetMaxNumberOfRepositoryDownloadsWithinTimePeriod(val *float64)
	MaxNumberOfRepositoryDownloadsWithinTimePeriodInput() *float64
	MaxPagesSize() *float64
	SetMaxPagesSize(val *float64)
	MaxPagesSizeInput() *float64
	MaxPersonalAccessTokenLifetime() *float64
	SetMaxPersonalAccessTokenLifetime(val *float64)
	MaxPersonalAccessTokenLifetimeInput() *float64
	MaxSshKeyLifetime() *float64
	SetMaxSshKeyLifetime(val *float64)
	MaxSshKeyLifetimeInput() *float64
	MetricsMethodCallThreshold() *float64
	SetMetricsMethodCallThreshold(val *float64)
	MetricsMethodCallThresholdInput() *float64
	MirrorAvailable() interface{}
	SetMirrorAvailable(val interface{})
	MirrorAvailableInput() interface{}
	MirrorCapacityThreshold() *float64
	SetMirrorCapacityThreshold(val *float64)
	MirrorCapacityThresholdInput() *float64
	MirrorMaxCapacity() *float64
	SetMirrorMaxCapacity(val *float64)
	MirrorMaxCapacityInput() *float64
	MirrorMaxDelay() *float64
	SetMirrorMaxDelay(val *float64)
	MirrorMaxDelayInput() *float64
	// The tree node.
	Node() constructs.Node
	NpmPackageRequestsForwarding() interface{}
	SetNpmPackageRequestsForwarding(val interface{})
	NpmPackageRequestsForwardingInput() interface{}
	OutboundLocalRequestsWhitelist() *[]*string
	SetOutboundLocalRequestsWhitelist(val *[]*string)
	OutboundLocalRequestsWhitelistInput() *[]*string
	PackageRegistryCleanupPoliciesWorkerCapacity() *float64
	SetPackageRegistryCleanupPoliciesWorkerCapacity(val *float64)
	PackageRegistryCleanupPoliciesWorkerCapacityInput() *float64
	PagesDomainVerificationEnabled() interface{}
	SetPagesDomainVerificationEnabled(val interface{})
	PagesDomainVerificationEnabledInput() interface{}
	PasswordAuthenticationEnabledForGit() interface{}
	SetPasswordAuthenticationEnabledForGit(val interface{})
	PasswordAuthenticationEnabledForGitInput() interface{}
	PasswordAuthenticationEnabledForWeb() interface{}
	SetPasswordAuthenticationEnabledForWeb(val interface{})
	PasswordAuthenticationEnabledForWebInput() interface{}
	PasswordLowercaseRequired() interface{}
	SetPasswordLowercaseRequired(val interface{})
	PasswordLowercaseRequiredInput() interface{}
	PasswordNumberRequired() interface{}
	SetPasswordNumberRequired(val interface{})
	PasswordNumberRequiredInput() interface{}
	PasswordSymbolRequired() interface{}
	SetPasswordSymbolRequired(val interface{})
	PasswordSymbolRequiredInput() interface{}
	PasswordUppercaseRequired() interface{}
	SetPasswordUppercaseRequired(val interface{})
	PasswordUppercaseRequiredInput() interface{}
	PerformanceBarAllowedGroupPath() *string
	SetPerformanceBarAllowedGroupPath(val *string)
	PerformanceBarAllowedGroupPathInput() *string
	PersonalAccessTokenPrefix() *string
	SetPersonalAccessTokenPrefix(val *string)
	PersonalAccessTokenPrefixInput() *string
	PipelineLimitPerProjectUserSha() *float64
	SetPipelineLimitPerProjectUserSha(val *float64)
	PipelineLimitPerProjectUserShaInput() *float64
	PlantumlEnabled() interface{}
	SetPlantumlEnabled(val interface{})
	PlantumlEnabledInput() interface{}
	PlantumlUrl() *string
	SetPlantumlUrl(val *string)
	PlantumlUrlInput() *string
	PollingIntervalMultiplier() *float64
	SetPollingIntervalMultiplier(val *float64)
	PollingIntervalMultiplierInput() *float64
	ProjectExportEnabled() interface{}
	SetProjectExportEnabled(val interface{})
	ProjectExportEnabledInput() interface{}
	PrometheusMetricsEnabled() interface{}
	SetPrometheusMetricsEnabled(val interface{})
	PrometheusMetricsEnabledInput() interface{}
	ProtectedCiVariables() interface{}
	SetProtectedCiVariables(val interface{})
	ProtectedCiVariablesInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PushEventActivitiesLimit() *float64
	SetPushEventActivitiesLimit(val *float64)
	PushEventActivitiesLimitInput() *float64
	PushEventHooksLimit() *float64
	SetPushEventHooksLimit(val *float64)
	PushEventHooksLimitInput() *float64
	PypiPackageRequestsForwarding() interface{}
	SetPypiPackageRequestsForwarding(val interface{})
	PypiPackageRequestsForwardingInput() interface{}
	RateLimitingResponseText() *string
	SetRateLimitingResponseText(val *string)
	RateLimitingResponseTextInput() *string
	RawBlobRequestLimit() *float64
	SetRawBlobRequestLimit(val *float64)
	RawBlobRequestLimitInput() *float64
	// Experimental.
	RawOverrides() interface{}
	RecaptchaEnabled() interface{}
	SetRecaptchaEnabled(val interface{})
	RecaptchaEnabledInput() interface{}
	RecaptchaPrivateKey() *string
	SetRecaptchaPrivateKey(val *string)
	RecaptchaPrivateKeyInput() *string
	RecaptchaSiteKey() *string
	SetRecaptchaSiteKey(val *string)
	RecaptchaSiteKeyInput() *string
	ReceiveMaxInputSize() *float64
	SetReceiveMaxInputSize(val *float64)
	ReceiveMaxInputSizeInput() *float64
	RepositoryChecksEnabled() interface{}
	SetRepositoryChecksEnabled(val interface{})
	RepositoryChecksEnabledInput() interface{}
	RepositorySizeLimit() *float64
	SetRepositorySizeLimit(val *float64)
	RepositorySizeLimitInput() *float64
	RepositoryStorages() *[]*string
	SetRepositoryStorages(val *[]*string)
	RepositoryStoragesInput() *[]*string
	RepositoryStoragesWeighted() *map[string]*float64
	SetRepositoryStoragesWeighted(val *map[string]*float64)
	RepositoryStoragesWeightedInput() *map[string]*float64
	RequireAdminApprovalAfterUserSignup() interface{}
	SetRequireAdminApprovalAfterUserSignup(val interface{})
	RequireAdminApprovalAfterUserSignupInput() interface{}
	RequireTwoFactorAuthentication() interface{}
	SetRequireTwoFactorAuthentication(val interface{})
	RequireTwoFactorAuthenticationInput() interface{}
	RestrictedVisibilityLevels() *[]*string
	SetRestrictedVisibilityLevels(val *[]*string)
	RestrictedVisibilityLevelsInput() *[]*string
	RsaKeyRestriction() *float64
	SetRsaKeyRestriction(val *float64)
	RsaKeyRestrictionInput() *float64
	SearchRateLimit() *float64
	SetSearchRateLimit(val *float64)
	SearchRateLimitInput() *float64
	SearchRateLimitUnauthenticated() *float64
	SetSearchRateLimitUnauthenticated(val *float64)
	SearchRateLimitUnauthenticatedInput() *float64
	SendUserConfirmationEmail() interface{}
	SetSendUserConfirmationEmail(val interface{})
	SendUserConfirmationEmailInput() interface{}
	SessionExpireDelay() *float64
	SetSessionExpireDelay(val *float64)
	SessionExpireDelayInput() *float64
	SharedRunnersEnabled() interface{}
	SetSharedRunnersEnabled(val interface{})
	SharedRunnersEnabledInput() interface{}
	SharedRunnersMinutes() *float64
	SetSharedRunnersMinutes(val *float64)
	SharedRunnersMinutesInput() *float64
	SharedRunnersText() *string
	SetSharedRunnersText(val *string)
	SharedRunnersTextInput() *string
	SidekiqJobLimiterCompressionThresholdBytes() *float64
	SetSidekiqJobLimiterCompressionThresholdBytes(val *float64)
	SidekiqJobLimiterCompressionThresholdBytesInput() *float64
	SidekiqJobLimiterLimitBytes() *float64
	SetSidekiqJobLimiterLimitBytes(val *float64)
	SidekiqJobLimiterLimitBytesInput() *float64
	SidekiqJobLimiterMode() *string
	SetSidekiqJobLimiterMode(val *string)
	SidekiqJobLimiterModeInput() *string
	SignInText() *string
	SetSignInText(val *string)
	SignInTextInput() *string
	SignupEnabled() interface{}
	SetSignupEnabled(val interface{})
	SignupEnabledInput() interface{}
	SlackAppEnabled() interface{}
	SetSlackAppEnabled(val interface{})
	SlackAppEnabledInput() interface{}
	SlackAppId() *string
	SetSlackAppId(val *string)
	SlackAppIdInput() *string
	SlackAppSecret() *string
	SetSlackAppSecret(val *string)
	SlackAppSecretInput() *string
	SlackAppSigningSecret() *string
	SetSlackAppSigningSecret(val *string)
	SlackAppSigningSecretInput() *string
	SlackAppVerificationToken() *string
	SetSlackAppVerificationToken(val *string)
	SlackAppVerificationTokenInput() *string
	SnippetSizeLimit() *float64
	SetSnippetSizeLimit(val *float64)
	SnippetSizeLimitInput() *float64
	SnowplowAppId() *string
	SetSnowplowAppId(val *string)
	SnowplowAppIdInput() *string
	SnowplowCollectorHostname() *string
	SetSnowplowCollectorHostname(val *string)
	SnowplowCollectorHostnameInput() *string
	SnowplowCookieDomain() *string
	SetSnowplowCookieDomain(val *string)
	SnowplowCookieDomainInput() *string
	SnowplowEnabled() interface{}
	SetSnowplowEnabled(val interface{})
	SnowplowEnabledInput() interface{}
	SourcegraphEnabled() interface{}
	SetSourcegraphEnabled(val interface{})
	SourcegraphEnabledInput() interface{}
	SourcegraphPublicOnly() interface{}
	SetSourcegraphPublicOnly(val interface{})
	SourcegraphPublicOnlyInput() interface{}
	SourcegraphUrl() *string
	SetSourcegraphUrl(val *string)
	SourcegraphUrlInput() *string
	SpamCheckApiKey() *string
	SetSpamCheckApiKey(val *string)
	SpamCheckApiKeyInput() *string
	SpamCheckEndpointEnabled() interface{}
	SetSpamCheckEndpointEnabled(val interface{})
	SpamCheckEndpointEnabledInput() interface{}
	SpamCheckEndpointUrl() *string
	SetSpamCheckEndpointUrl(val *string)
	SpamCheckEndpointUrlInput() *string
	SuggestPipelineEnabled() interface{}
	SetSuggestPipelineEnabled(val interface{})
	SuggestPipelineEnabledInput() interface{}
	TerminalMaxSessionTime() *float64
	SetTerminalMaxSessionTime(val *float64)
	TerminalMaxSessionTimeInput() *float64
	Terms() *string
	SetTerms(val *string)
	TermsInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ThrottleAuthenticatedApiEnabled() interface{}
	SetThrottleAuthenticatedApiEnabled(val interface{})
	ThrottleAuthenticatedApiEnabledInput() interface{}
	ThrottleAuthenticatedApiPeriodInSeconds() *float64
	SetThrottleAuthenticatedApiPeriodInSeconds(val *float64)
	ThrottleAuthenticatedApiPeriodInSecondsInput() *float64
	ThrottleAuthenticatedApiRequestsPerPeriod() *float64
	SetThrottleAuthenticatedApiRequestsPerPeriod(val *float64)
	ThrottleAuthenticatedApiRequestsPerPeriodInput() *float64
	ThrottleAuthenticatedPackagesApiEnabled() interface{}
	SetThrottleAuthenticatedPackagesApiEnabled(val interface{})
	ThrottleAuthenticatedPackagesApiEnabledInput() interface{}
	ThrottleAuthenticatedPackagesApiPeriodInSeconds() *float64
	SetThrottleAuthenticatedPackagesApiPeriodInSeconds(val *float64)
	ThrottleAuthenticatedPackagesApiPeriodInSecondsInput() *float64
	ThrottleAuthenticatedPackagesApiRequestsPerPeriod() *float64
	SetThrottleAuthenticatedPackagesApiRequestsPerPeriod(val *float64)
	ThrottleAuthenticatedPackagesApiRequestsPerPeriodInput() *float64
	ThrottleAuthenticatedWebEnabled() interface{}
	SetThrottleAuthenticatedWebEnabled(val interface{})
	ThrottleAuthenticatedWebEnabledInput() interface{}
	ThrottleAuthenticatedWebPeriodInSeconds() *float64
	SetThrottleAuthenticatedWebPeriodInSeconds(val *float64)
	ThrottleAuthenticatedWebPeriodInSecondsInput() *float64
	ThrottleAuthenticatedWebRequestsPerPeriod() *float64
	SetThrottleAuthenticatedWebRequestsPerPeriod(val *float64)
	ThrottleAuthenticatedWebRequestsPerPeriodInput() *float64
	ThrottleUnauthenticatedApiEnabled() interface{}
	SetThrottleUnauthenticatedApiEnabled(val interface{})
	ThrottleUnauthenticatedApiEnabledInput() interface{}
	ThrottleUnauthenticatedApiPeriodInSeconds() *float64
	SetThrottleUnauthenticatedApiPeriodInSeconds(val *float64)
	ThrottleUnauthenticatedApiPeriodInSecondsInput() *float64
	ThrottleUnauthenticatedApiRequestsPerPeriod() *float64
	SetThrottleUnauthenticatedApiRequestsPerPeriod(val *float64)
	ThrottleUnauthenticatedApiRequestsPerPeriodInput() *float64
	ThrottleUnauthenticatedPackagesApiEnabled() interface{}
	SetThrottleUnauthenticatedPackagesApiEnabled(val interface{})
	ThrottleUnauthenticatedPackagesApiEnabledInput() interface{}
	ThrottleUnauthenticatedPackagesApiPeriodInSeconds() *float64
	SetThrottleUnauthenticatedPackagesApiPeriodInSeconds(val *float64)
	ThrottleUnauthenticatedPackagesApiPeriodInSecondsInput() *float64
	ThrottleUnauthenticatedPackagesApiRequestsPerPeriod() *float64
	SetThrottleUnauthenticatedPackagesApiRequestsPerPeriod(val *float64)
	ThrottleUnauthenticatedPackagesApiRequestsPerPeriodInput() *float64
	ThrottleUnauthenticatedWebEnabled() interface{}
	SetThrottleUnauthenticatedWebEnabled(val interface{})
	ThrottleUnauthenticatedWebEnabledInput() interface{}
	ThrottleUnauthenticatedWebPeriodInSeconds() *float64
	SetThrottleUnauthenticatedWebPeriodInSeconds(val *float64)
	ThrottleUnauthenticatedWebPeriodInSecondsInput() *float64
	ThrottleUnauthenticatedWebRequestsPerPeriod() *float64
	SetThrottleUnauthenticatedWebRequestsPerPeriod(val *float64)
	ThrottleUnauthenticatedWebRequestsPerPeriodInput() *float64
	TimeTrackingLimitToHours() interface{}
	SetTimeTrackingLimitToHours(val interface{})
	TimeTrackingLimitToHoursInput() interface{}
	TwoFactorGracePeriod() *float64
	SetTwoFactorGracePeriod(val *float64)
	TwoFactorGracePeriodInput() *float64
	UniqueIpsLimitEnabled() interface{}
	SetUniqueIpsLimitEnabled(val interface{})
	UniqueIpsLimitEnabledInput() interface{}
	UniqueIpsLimitPerUser() *float64
	SetUniqueIpsLimitPerUser(val *float64)
	UniqueIpsLimitPerUserInput() *float64
	UniqueIpsLimitTimeWindow() *float64
	SetUniqueIpsLimitTimeWindow(val *float64)
	UniqueIpsLimitTimeWindowInput() *float64
	UsagePingEnabled() interface{}
	SetUsagePingEnabled(val interface{})
	UsagePingEnabledInput() interface{}
	UserDeactivationEmailsEnabled() interface{}
	SetUserDeactivationEmailsEnabled(val interface{})
	UserDeactivationEmailsEnabledInput() interface{}
	UserDefaultExternal() interface{}
	SetUserDefaultExternal(val interface{})
	UserDefaultExternalInput() interface{}
	UserDefaultInternalRegex() *string
	SetUserDefaultInternalRegex(val *string)
	UserDefaultInternalRegexInput() *string
	UserOauthApplications() interface{}
	SetUserOauthApplications(val interface{})
	UserOauthApplicationsInput() interface{}
	UserShowAddSshKeyMessage() interface{}
	SetUserShowAddSshKeyMessage(val interface{})
	UserShowAddSshKeyMessageInput() interface{}
	VersionCheckEnabled() interface{}
	SetVersionCheckEnabled(val interface{})
	VersionCheckEnabledInput() interface{}
	WebIdeClientsidePreviewEnabled() interface{}
	SetWebIdeClientsidePreviewEnabled(val interface{})
	WebIdeClientsidePreviewEnabledInput() interface{}
	WhatsNewVariant() *string
	SetWhatsNewVariant(val *string)
	WhatsNewVariantInput() *string
	WikiPageMaxContentBytes() *float64
	SetWikiPageMaxContentBytes(val *float64)
	WikiPageMaxContentBytesInput() *float64
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAbuseNotificationEmail()
	ResetAdminMode()
	ResetAfterSignOutPath()
	ResetAfterSignUpText()
	ResetAkismetApiKey()
	ResetAkismetEnabled()
	ResetAllowGroupOwnersToManageLdap()
	ResetAllowLocalRequestsFromSystemHooks()
	ResetAllowLocalRequestsFromWebHooksAndServices()
	ResetArchiveBuildsInHumanReadable()
	ResetAssetProxyAllowlist()
	ResetAssetProxyEnabled()
	ResetAssetProxySecretKey()
	ResetAssetProxyUrl()
	ResetAuthorizedKeysEnabled()
	ResetAutoDevopsDomain()
	ResetAutoDevopsEnabled()
	ResetAutomaticPurchasedStorageAllocation()
	ResetCanCreateGroup()
	ResetCheckNamespacePlan()
	ResetCommitEmailHostname()
	ResetContainerExpirationPoliciesEnableHistoricEntries()
	ResetContainerRegistryCleanupTagsServiceMaxListSize()
	ResetContainerRegistryDeleteTagsServiceTimeout()
	ResetContainerRegistryExpirationPoliciesCaching()
	ResetContainerRegistryExpirationPoliciesWorkerCapacity()
	ResetContainerRegistryTokenExpireDelay()
	ResetDeactivateDormantUsers()
	ResetDefaultArtifactsExpireIn()
	ResetDefaultBranchName()
	ResetDefaultBranchProtection()
	ResetDefaultCiConfigPath()
	ResetDefaultGroupVisibility()
	ResetDefaultProjectCreation()
	ResetDefaultProjectsLimit()
	ResetDefaultProjectVisibility()
	ResetDefaultSnippetVisibility()
	ResetDeleteInactiveProjects()
	ResetDeletionAdjournedPeriod()
	ResetDiffMaxFiles()
	ResetDiffMaxLines()
	ResetDiffMaxPatchBytes()
	ResetDisabledOauthSignInSources()
	ResetDisableFeedToken()
	ResetDnsRebindingProtectionEnabled()
	ResetDomainAllowlist()
	ResetDomainDenylist()
	ResetDomainDenylistEnabled()
	ResetDsaKeyRestriction()
	ResetEcdsaKeyRestriction()
	ResetEcdsaSkKeyRestriction()
	ResetEd25519KeyRestriction()
	ResetEd25519SkKeyRestriction()
	ResetEksAccessKeyId()
	ResetEksAccountId()
	ResetEksIntegrationEnabled()
	ResetEksSecretAccessKey()
	ResetElasticsearchAws()
	ResetElasticsearchAwsAccessKey()
	ResetElasticsearchAwsRegion()
	ResetElasticsearchAwsSecretAccessKey()
	ResetElasticsearchIndexedFieldLengthLimit()
	ResetElasticsearchIndexedFileSizeLimitKb()
	ResetElasticsearchIndexing()
	ResetElasticsearchLimitIndexing()
	ResetElasticsearchMaxBulkConcurrency()
	ResetElasticsearchMaxBulkSizeMb()
	ResetElasticsearchNamespaceIds()
	ResetElasticsearchPassword()
	ResetElasticsearchProjectIds()
	ResetElasticsearchSearch()
	ResetElasticsearchUrl()
	ResetElasticsearchUsername()
	ResetEmailAdditionalText()
	ResetEmailAuthorInBody()
	ResetEnabledGitAccessProtocol()
	ResetEnforceNamespaceStorageLimit()
	ResetEnforceTerms()
	ResetExternalAuthClientCert()
	ResetExternalAuthClientKey()
	ResetExternalAuthClientKeyPass()
	ResetExternalAuthorizationServiceDefaultLabel()
	ResetExternalAuthorizationServiceEnabled()
	ResetExternalAuthorizationServiceTimeout()
	ResetExternalAuthorizationServiceUrl()
	ResetExternalPipelineValidationServiceTimeout()
	ResetExternalPipelineValidationServiceToken()
	ResetExternalPipelineValidationServiceUrl()
	ResetFileTemplateProjectId()
	ResetFirstDayOfWeek()
	ResetGeoNodeAllowedIps()
	ResetGeoStatusTimeout()
	ResetGitalyTimeoutDefault()
	ResetGitalyTimeoutFast()
	ResetGitalyTimeoutMedium()
	ResetGitRateLimitUsersAllowlist()
	ResetGitTwoFactorSessionExpiry()
	ResetGrafanaEnabled()
	ResetGrafanaUrl()
	ResetGravatarEnabled()
	ResetGroupOwnersCanManageDefaultBranchProtection()
	ResetHashedStorageEnabled()
	ResetHelpPageHideCommercialContent()
	ResetHelpPageSupportUrl()
	ResetHelpPageText()
	ResetHelpText()
	ResetHideThirdPartyOffers()
	ResetHomePageUrl()
	ResetHousekeepingEnabled()
	ResetHousekeepingFullRepackPeriod()
	ResetHousekeepingGcPeriod()
	ResetHousekeepingIncrementalRepackPeriod()
	ResetHousekeepingOptimizeRepositoryPeriod()
	ResetHtmlEmailsEnabled()
	ResetId()
	ResetImportSources()
	ResetInactiveProjectsDeleteAfterMonths()
	ResetInactiveProjectsMinSizeMb()
	ResetInactiveProjectsSendWarningEmailAfterMonths()
	ResetInProductMarketingEmailsEnabled()
	ResetInvisibleCaptchaEnabled()
	ResetIssuesCreateLimit()
	ResetKeepLatestArtifact()
	ResetLocalMarkdownVersion()
	ResetMailgunEventsEnabled()
	ResetMailgunSigningKey()
	ResetMaintenanceMode()
	ResetMaintenanceModeMessage()
	ResetMaxArtifactsSize()
	ResetMaxAttachmentSize()
	ResetMaxExportSize()
	ResetMaxImportSize()
	ResetMaxNumberOfRepositoryDownloads()
	ResetMaxNumberOfRepositoryDownloadsWithinTimePeriod()
	ResetMaxPagesSize()
	ResetMaxPersonalAccessTokenLifetime()
	ResetMaxSshKeyLifetime()
	ResetMetricsMethodCallThreshold()
	ResetMirrorAvailable()
	ResetMirrorCapacityThreshold()
	ResetMirrorMaxCapacity()
	ResetMirrorMaxDelay()
	ResetNpmPackageRequestsForwarding()
	ResetOutboundLocalRequestsWhitelist()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPackageRegistryCleanupPoliciesWorkerCapacity()
	ResetPagesDomainVerificationEnabled()
	ResetPasswordAuthenticationEnabledForGit()
	ResetPasswordAuthenticationEnabledForWeb()
	ResetPasswordLowercaseRequired()
	ResetPasswordNumberRequired()
	ResetPasswordSymbolRequired()
	ResetPasswordUppercaseRequired()
	ResetPerformanceBarAllowedGroupPath()
	ResetPersonalAccessTokenPrefix()
	ResetPipelineLimitPerProjectUserSha()
	ResetPlantumlEnabled()
	ResetPlantumlUrl()
	ResetPollingIntervalMultiplier()
	ResetProjectExportEnabled()
	ResetPrometheusMetricsEnabled()
	ResetProtectedCiVariables()
	ResetPushEventActivitiesLimit()
	ResetPushEventHooksLimit()
	ResetPypiPackageRequestsForwarding()
	ResetRateLimitingResponseText()
	ResetRawBlobRequestLimit()
	ResetRecaptchaEnabled()
	ResetRecaptchaPrivateKey()
	ResetRecaptchaSiteKey()
	ResetReceiveMaxInputSize()
	ResetRepositoryChecksEnabled()
	ResetRepositorySizeLimit()
	ResetRepositoryStorages()
	ResetRepositoryStoragesWeighted()
	ResetRequireAdminApprovalAfterUserSignup()
	ResetRequireTwoFactorAuthentication()
	ResetRestrictedVisibilityLevels()
	ResetRsaKeyRestriction()
	ResetSearchRateLimit()
	ResetSearchRateLimitUnauthenticated()
	ResetSendUserConfirmationEmail()
	ResetSessionExpireDelay()
	ResetSharedRunnersEnabled()
	ResetSharedRunnersMinutes()
	ResetSharedRunnersText()
	ResetSidekiqJobLimiterCompressionThresholdBytes()
	ResetSidekiqJobLimiterLimitBytes()
	ResetSidekiqJobLimiterMode()
	ResetSignInText()
	ResetSignupEnabled()
	ResetSlackAppEnabled()
	ResetSlackAppId()
	ResetSlackAppSecret()
	ResetSlackAppSigningSecret()
	ResetSlackAppVerificationToken()
	ResetSnippetSizeLimit()
	ResetSnowplowAppId()
	ResetSnowplowCollectorHostname()
	ResetSnowplowCookieDomain()
	ResetSnowplowEnabled()
	ResetSourcegraphEnabled()
	ResetSourcegraphPublicOnly()
	ResetSourcegraphUrl()
	ResetSpamCheckApiKey()
	ResetSpamCheckEndpointEnabled()
	ResetSpamCheckEndpointUrl()
	ResetSuggestPipelineEnabled()
	ResetTerminalMaxSessionTime()
	ResetTerms()
	ResetThrottleAuthenticatedApiEnabled()
	ResetThrottleAuthenticatedApiPeriodInSeconds()
	ResetThrottleAuthenticatedApiRequestsPerPeriod()
	ResetThrottleAuthenticatedPackagesApiEnabled()
	ResetThrottleAuthenticatedPackagesApiPeriodInSeconds()
	ResetThrottleAuthenticatedPackagesApiRequestsPerPeriod()
	ResetThrottleAuthenticatedWebEnabled()
	ResetThrottleAuthenticatedWebPeriodInSeconds()
	ResetThrottleAuthenticatedWebRequestsPerPeriod()
	ResetThrottleUnauthenticatedApiEnabled()
	ResetThrottleUnauthenticatedApiPeriodInSeconds()
	ResetThrottleUnauthenticatedApiRequestsPerPeriod()
	ResetThrottleUnauthenticatedPackagesApiEnabled()
	ResetThrottleUnauthenticatedPackagesApiPeriodInSeconds()
	ResetThrottleUnauthenticatedPackagesApiRequestsPerPeriod()
	ResetThrottleUnauthenticatedWebEnabled()
	ResetThrottleUnauthenticatedWebPeriodInSeconds()
	ResetThrottleUnauthenticatedWebRequestsPerPeriod()
	ResetTimeTrackingLimitToHours()
	ResetTwoFactorGracePeriod()
	ResetUniqueIpsLimitEnabled()
	ResetUniqueIpsLimitPerUser()
	ResetUniqueIpsLimitTimeWindow()
	ResetUsagePingEnabled()
	ResetUserDeactivationEmailsEnabled()
	ResetUserDefaultExternal()
	ResetUserDefaultInternalRegex()
	ResetUserOauthApplications()
	ResetUserShowAddSshKeyMessage()
	ResetVersionCheckEnabled()
	ResetWebIdeClientsidePreviewEnabled()
	ResetWhatsNewVariant()
	ResetWikiPageMaxContentBytes()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ApplicationSettings
type jsiiProxy_ApplicationSettings struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApplicationSettings) AbuseNotificationEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"abuseNotificationEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AbuseNotificationEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"abuseNotificationEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AdminMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AdminModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AfterSignOutPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterSignOutPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AfterSignOutPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterSignOutPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AfterSignUpText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterSignUpText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AfterSignUpTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterSignUpTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AkismetApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"akismetApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AkismetApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"akismetApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AkismetEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"akismetEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AkismetEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"akismetEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AllowGroupOwnersToManageLdap() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowGroupOwnersToManageLdap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AllowGroupOwnersToManageLdapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowGroupOwnersToManageLdapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AllowLocalRequestsFromSystemHooks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowLocalRequestsFromSystemHooks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AllowLocalRequestsFromSystemHooksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowLocalRequestsFromSystemHooksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AllowLocalRequestsFromWebHooksAndServices() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowLocalRequestsFromWebHooksAndServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AllowLocalRequestsFromWebHooksAndServicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowLocalRequestsFromWebHooksAndServicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ArchiveBuildsInHumanReadable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"archiveBuildsInHumanReadable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ArchiveBuildsInHumanReadableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"archiveBuildsInHumanReadableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AssetProxyAllowlist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"assetProxyAllowlist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AssetProxyAllowlistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"assetProxyAllowlistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AssetProxyEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assetProxyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AssetProxyEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assetProxyEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AssetProxySecretKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetProxySecretKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AssetProxySecretKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetProxySecretKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AssetProxyUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetProxyUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AssetProxyUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetProxyUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AuthorizedKeysEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authorizedKeysEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AuthorizedKeysEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authorizedKeysEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AutoDevopsDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoDevopsDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AutoDevopsDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoDevopsDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AutoDevopsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDevopsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AutoDevopsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDevopsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AutomaticPurchasedStorageAllocation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticPurchasedStorageAllocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AutomaticPurchasedStorageAllocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticPurchasedStorageAllocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) CanCreateGroup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canCreateGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) CanCreateGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canCreateGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) CheckNamespacePlan() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkNamespacePlan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) CheckNamespacePlanInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkNamespacePlanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) CommitEmailHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitEmailHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) CommitEmailHostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitEmailHostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ContainerExpirationPoliciesEnableHistoricEntries() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerExpirationPoliciesEnableHistoricEntries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ContainerExpirationPoliciesEnableHistoricEntriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerExpirationPoliciesEnableHistoricEntriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ContainerRegistryCleanupTagsServiceMaxListSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerRegistryCleanupTagsServiceMaxListSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ContainerRegistryCleanupTagsServiceMaxListSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerRegistryCleanupTagsServiceMaxListSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ContainerRegistryDeleteTagsServiceTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerRegistryDeleteTagsServiceTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ContainerRegistryDeleteTagsServiceTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerRegistryDeleteTagsServiceTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ContainerRegistryExpirationPoliciesCaching() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerRegistryExpirationPoliciesCaching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ContainerRegistryExpirationPoliciesCachingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerRegistryExpirationPoliciesCachingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ContainerRegistryExpirationPoliciesWorkerCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerRegistryExpirationPoliciesWorkerCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ContainerRegistryExpirationPoliciesWorkerCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerRegistryExpirationPoliciesWorkerCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ContainerRegistryTokenExpireDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerRegistryTokenExpireDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ContainerRegistryTokenExpireDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerRegistryTokenExpireDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DeactivateDormantUsers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deactivateDormantUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DeactivateDormantUsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deactivateDormantUsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DefaultArtifactsExpireIn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultArtifactsExpireIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DefaultArtifactsExpireInInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultArtifactsExpireInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DefaultBranchName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBranchName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DefaultBranchNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBranchNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DefaultBranchProtection() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultBranchProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DefaultBranchProtectionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultBranchProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DefaultCiConfigPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultCiConfigPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DefaultCiConfigPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultCiConfigPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DefaultGroupVisibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultGroupVisibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DefaultGroupVisibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultGroupVisibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DefaultProjectCreation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultProjectCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DefaultProjectCreationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultProjectCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DefaultProjectsLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultProjectsLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DefaultProjectsLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultProjectsLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DefaultProjectVisibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultProjectVisibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DefaultProjectVisibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultProjectVisibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DefaultSnippetVisibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSnippetVisibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DefaultSnippetVisibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSnippetVisibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DeleteInactiveProjects() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteInactiveProjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DeleteInactiveProjectsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteInactiveProjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DeletionAdjournedPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deletionAdjournedPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DeletionAdjournedPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deletionAdjournedPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DiffMaxFiles() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diffMaxFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DiffMaxFilesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diffMaxFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DiffMaxLines() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diffMaxLines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DiffMaxLinesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diffMaxLinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DiffMaxPatchBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diffMaxPatchBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DiffMaxPatchBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diffMaxPatchBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DisabledOauthSignInSources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disabledOauthSignInSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DisabledOauthSignInSourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disabledOauthSignInSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DisableFeedToken() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableFeedToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DisableFeedTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableFeedTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DnsRebindingProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dnsRebindingProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DnsRebindingProtectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dnsRebindingProtectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DomainAllowlist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainAllowlist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DomainAllowlistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainAllowlistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DomainDenylist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainDenylist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DomainDenylistEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"domainDenylistEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DomainDenylistEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"domainDenylistEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DomainDenylistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainDenylistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DsaKeyRestriction() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dsaKeyRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DsaKeyRestrictionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dsaKeyRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EcdsaKeyRestriction() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ecdsaKeyRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EcdsaKeyRestrictionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ecdsaKeyRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EcdsaSkKeyRestriction() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ecdsaSkKeyRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EcdsaSkKeyRestrictionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ecdsaSkKeyRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) Ed25519KeyRestriction() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ed25519KeyRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) Ed25519KeyRestrictionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ed25519KeyRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) Ed25519SkKeyRestriction() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ed25519SkKeyRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) Ed25519SkKeyRestrictionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ed25519SkKeyRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EksAccessKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eksAccessKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EksAccessKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eksAccessKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EksAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eksAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EksAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eksAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EksIntegrationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eksIntegrationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EksIntegrationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eksIntegrationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EksSecretAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eksSecretAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EksSecretAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eksSecretAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchAws() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchAws",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchAwsAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchAwsAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchAwsAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchAwsAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchAwsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchAwsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchAwsRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchAwsRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchAwsRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchAwsRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchAwsSecretAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchAwsSecretAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchAwsSecretAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchAwsSecretAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchIndexedFieldLengthLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"elasticsearchIndexedFieldLengthLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchIndexedFieldLengthLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"elasticsearchIndexedFieldLengthLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchIndexedFileSizeLimitKb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"elasticsearchIndexedFileSizeLimitKb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchIndexedFileSizeLimitKbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"elasticsearchIndexedFileSizeLimitKbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchIndexing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchIndexing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchIndexingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchIndexingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchLimitIndexing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchLimitIndexing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchLimitIndexingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchLimitIndexingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchMaxBulkConcurrency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"elasticsearchMaxBulkConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchMaxBulkConcurrencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"elasticsearchMaxBulkConcurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchMaxBulkSizeMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"elasticsearchMaxBulkSizeMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchMaxBulkSizeMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"elasticsearchMaxBulkSizeMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchNamespaceIds() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"elasticsearchNamespaceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchNamespaceIdsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"elasticsearchNamespaceIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchProjectIds() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"elasticsearchProjectIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchProjectIdsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"elasticsearchProjectIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchSearch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchSearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchSearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchSearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchUrl() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"elasticsearchUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchUrlInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"elasticsearchUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ElasticsearchUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EmailAdditionalText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAdditionalText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EmailAdditionalTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAdditionalTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EmailAuthorInBody() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailAuthorInBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EmailAuthorInBodyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailAuthorInBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EnabledGitAccessProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enabledGitAccessProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EnabledGitAccessProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enabledGitAccessProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EnforceNamespaceStorageLimit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceNamespaceStorageLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EnforceNamespaceStorageLimitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceNamespaceStorageLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EnforceTerms() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceTerms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EnforceTermsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceTermsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ExternalAuthClientCert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalAuthClientCert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ExternalAuthClientCertInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalAuthClientCertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ExternalAuthClientKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalAuthClientKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ExternalAuthClientKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalAuthClientKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ExternalAuthClientKeyPass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalAuthClientKeyPass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ExternalAuthClientKeyPassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalAuthClientKeyPassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ExternalAuthorizationServiceDefaultLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalAuthorizationServiceDefaultLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ExternalAuthorizationServiceDefaultLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalAuthorizationServiceDefaultLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ExternalAuthorizationServiceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalAuthorizationServiceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ExternalAuthorizationServiceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalAuthorizationServiceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ExternalAuthorizationServiceTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"externalAuthorizationServiceTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ExternalAuthorizationServiceTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"externalAuthorizationServiceTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ExternalAuthorizationServiceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalAuthorizationServiceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ExternalAuthorizationServiceUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalAuthorizationServiceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ExternalPipelineValidationServiceTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"externalPipelineValidationServiceTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ExternalPipelineValidationServiceTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"externalPipelineValidationServiceTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ExternalPipelineValidationServiceToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalPipelineValidationServiceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ExternalPipelineValidationServiceTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalPipelineValidationServiceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ExternalPipelineValidationServiceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalPipelineValidationServiceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ExternalPipelineValidationServiceUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalPipelineValidationServiceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) FileTemplateProjectId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fileTemplateProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) FileTemplateProjectIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fileTemplateProjectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) FirstDayOfWeek() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firstDayOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) FirstDayOfWeekInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firstDayOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) GeoNodeAllowedIps() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geoNodeAllowedIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) GeoNodeAllowedIpsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geoNodeAllowedIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) GeoStatusTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"geoStatusTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) GeoStatusTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"geoStatusTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) GitalyTimeoutDefault() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gitalyTimeoutDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) GitalyTimeoutDefaultInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gitalyTimeoutDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) GitalyTimeoutFast() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gitalyTimeoutFast",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) GitalyTimeoutFastInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gitalyTimeoutFastInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) GitalyTimeoutMedium() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gitalyTimeoutMedium",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) GitalyTimeoutMediumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gitalyTimeoutMediumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) GitRateLimitUsersAllowlist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gitRateLimitUsersAllowlist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) GitRateLimitUsersAllowlistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gitRateLimitUsersAllowlistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) GitTwoFactorSessionExpiry() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gitTwoFactorSessionExpiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) GitTwoFactorSessionExpiryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gitTwoFactorSessionExpiryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) GrafanaEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"grafanaEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) GrafanaEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"grafanaEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) GrafanaUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) GrafanaUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) GravatarEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gravatarEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) GravatarEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gravatarEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) GroupOwnersCanManageDefaultBranchProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupOwnersCanManageDefaultBranchProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) GroupOwnersCanManageDefaultBranchProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupOwnersCanManageDefaultBranchProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) HashedStorageEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hashedStorageEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) HashedStorageEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hashedStorageEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) HelpPageHideCommercialContent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"helpPageHideCommercialContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) HelpPageHideCommercialContentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"helpPageHideCommercialContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) HelpPageSupportUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"helpPageSupportUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) HelpPageSupportUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"helpPageSupportUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) HelpPageText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"helpPageText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) HelpPageTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"helpPageTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) HelpText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"helpText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) HelpTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"helpTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) HideThirdPartyOffers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideThirdPartyOffers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) HideThirdPartyOffersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideThirdPartyOffersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) HomePageUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homePageUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) HomePageUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homePageUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) HousekeepingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"housekeepingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) HousekeepingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"housekeepingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) HousekeepingFullRepackPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"housekeepingFullRepackPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) HousekeepingFullRepackPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"housekeepingFullRepackPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) HousekeepingGcPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"housekeepingGcPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) HousekeepingGcPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"housekeepingGcPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) HousekeepingIncrementalRepackPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"housekeepingIncrementalRepackPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) HousekeepingIncrementalRepackPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"housekeepingIncrementalRepackPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) HousekeepingOptimizeRepositoryPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"housekeepingOptimizeRepositoryPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) HousekeepingOptimizeRepositoryPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"housekeepingOptimizeRepositoryPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) HtmlEmailsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"htmlEmailsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) HtmlEmailsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"htmlEmailsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ImportSources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"importSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ImportSourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"importSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) InactiveProjectsDeleteAfterMonths() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"inactiveProjectsDeleteAfterMonths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) InactiveProjectsDeleteAfterMonthsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"inactiveProjectsDeleteAfterMonthsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) InactiveProjectsMinSizeMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"inactiveProjectsMinSizeMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) InactiveProjectsMinSizeMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"inactiveProjectsMinSizeMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) InactiveProjectsSendWarningEmailAfterMonths() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"inactiveProjectsSendWarningEmailAfterMonths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) InactiveProjectsSendWarningEmailAfterMonthsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"inactiveProjectsSendWarningEmailAfterMonthsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) InProductMarketingEmailsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inProductMarketingEmailsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) InProductMarketingEmailsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inProductMarketingEmailsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) InvisibleCaptchaEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invisibleCaptchaEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) InvisibleCaptchaEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invisibleCaptchaEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) IssuesCreateLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"issuesCreateLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) IssuesCreateLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"issuesCreateLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) KeepLatestArtifact() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepLatestArtifact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) KeepLatestArtifactInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepLatestArtifactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) LocalMarkdownVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localMarkdownVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) LocalMarkdownVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localMarkdownVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MailgunEventsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mailgunEventsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MailgunEventsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mailgunEventsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MailgunSigningKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailgunSigningKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MailgunSigningKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailgunSigningKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MaintenanceMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maintenanceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MaintenanceModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maintenanceModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MaintenanceModeMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceModeMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MaintenanceModeMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceModeMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MaxArtifactsSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxArtifactsSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MaxArtifactsSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxArtifactsSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MaxAttachmentSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAttachmentSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MaxAttachmentSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAttachmentSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MaxExportSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxExportSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MaxExportSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxExportSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MaxImportSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxImportSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MaxImportSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxImportSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MaxNumberOfRepositoryDownloads() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNumberOfRepositoryDownloads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MaxNumberOfRepositoryDownloadsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNumberOfRepositoryDownloadsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MaxNumberOfRepositoryDownloadsWithinTimePeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNumberOfRepositoryDownloadsWithinTimePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MaxNumberOfRepositoryDownloadsWithinTimePeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNumberOfRepositoryDownloadsWithinTimePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MaxPagesSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPagesSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MaxPagesSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPagesSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MaxPersonalAccessTokenLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPersonalAccessTokenLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MaxPersonalAccessTokenLifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPersonalAccessTokenLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MaxSshKeyLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSshKeyLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MaxSshKeyLifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSshKeyLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MetricsMethodCallThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsMethodCallThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MetricsMethodCallThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsMethodCallThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MirrorAvailable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mirrorAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MirrorAvailableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mirrorAvailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MirrorCapacityThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mirrorCapacityThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MirrorCapacityThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mirrorCapacityThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MirrorMaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mirrorMaxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MirrorMaxCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mirrorMaxCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MirrorMaxDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mirrorMaxDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) MirrorMaxDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mirrorMaxDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) NpmPackageRequestsForwarding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"npmPackageRequestsForwarding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) NpmPackageRequestsForwardingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"npmPackageRequestsForwardingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) OutboundLocalRequestsWhitelist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outboundLocalRequestsWhitelist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) OutboundLocalRequestsWhitelistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outboundLocalRequestsWhitelistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PackageRegistryCleanupPoliciesWorkerCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"packageRegistryCleanupPoliciesWorkerCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PackageRegistryCleanupPoliciesWorkerCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"packageRegistryCleanupPoliciesWorkerCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PagesDomainVerificationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pagesDomainVerificationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PagesDomainVerificationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pagesDomainVerificationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PasswordAuthenticationEnabledForGit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordAuthenticationEnabledForGit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PasswordAuthenticationEnabledForGitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordAuthenticationEnabledForGitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PasswordAuthenticationEnabledForWeb() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordAuthenticationEnabledForWeb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PasswordAuthenticationEnabledForWebInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordAuthenticationEnabledForWebInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PasswordLowercaseRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordLowercaseRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PasswordLowercaseRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordLowercaseRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PasswordNumberRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordNumberRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PasswordNumberRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordNumberRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PasswordSymbolRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordSymbolRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PasswordSymbolRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordSymbolRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PasswordUppercaseRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordUppercaseRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PasswordUppercaseRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordUppercaseRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PerformanceBarAllowedGroupPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceBarAllowedGroupPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PerformanceBarAllowedGroupPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceBarAllowedGroupPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PersonalAccessTokenPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"personalAccessTokenPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PersonalAccessTokenPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"personalAccessTokenPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PipelineLimitPerProjectUserSha() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pipelineLimitPerProjectUserSha",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PipelineLimitPerProjectUserShaInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pipelineLimitPerProjectUserShaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PlantumlEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"plantumlEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PlantumlEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"plantumlEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PlantumlUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"plantumlUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PlantumlUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"plantumlUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PollingIntervalMultiplier() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pollingIntervalMultiplier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PollingIntervalMultiplierInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pollingIntervalMultiplierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ProjectExportEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"projectExportEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ProjectExportEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"projectExportEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PrometheusMetricsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prometheusMetricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PrometheusMetricsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prometheusMetricsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ProtectedCiVariables() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protectedCiVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ProtectedCiVariablesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protectedCiVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PushEventActivitiesLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pushEventActivitiesLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PushEventActivitiesLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pushEventActivitiesLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PushEventHooksLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pushEventHooksLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PushEventHooksLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pushEventHooksLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PypiPackageRequestsForwarding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pypiPackageRequestsForwarding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) PypiPackageRequestsForwardingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pypiPackageRequestsForwardingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RateLimitingResponseText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateLimitingResponseText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RateLimitingResponseTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateLimitingResponseTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RawBlobRequestLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rawBlobRequestLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RawBlobRequestLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rawBlobRequestLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RecaptchaEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recaptchaEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RecaptchaEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recaptchaEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RecaptchaPrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recaptchaPrivateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RecaptchaPrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recaptchaPrivateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RecaptchaSiteKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recaptchaSiteKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RecaptchaSiteKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recaptchaSiteKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ReceiveMaxInputSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"receiveMaxInputSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ReceiveMaxInputSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"receiveMaxInputSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RepositoryChecksEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"repositoryChecksEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RepositoryChecksEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"repositoryChecksEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RepositorySizeLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"repositorySizeLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RepositorySizeLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"repositorySizeLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RepositoryStorages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"repositoryStorages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RepositoryStoragesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"repositoryStoragesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RepositoryStoragesWeighted() *map[string]*float64 {
	var returns *map[string]*float64
	_jsii_.Get(
		j,
		"repositoryStoragesWeighted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RepositoryStoragesWeightedInput() *map[string]*float64 {
	var returns *map[string]*float64
	_jsii_.Get(
		j,
		"repositoryStoragesWeightedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RequireAdminApprovalAfterUserSignup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireAdminApprovalAfterUserSignup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RequireAdminApprovalAfterUserSignupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireAdminApprovalAfterUserSignupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RequireTwoFactorAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireTwoFactorAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RequireTwoFactorAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireTwoFactorAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RestrictedVisibilityLevels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"restrictedVisibilityLevels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RestrictedVisibilityLevelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"restrictedVisibilityLevelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RsaKeyRestriction() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rsaKeyRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RsaKeyRestrictionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rsaKeyRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SearchRateLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"searchRateLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SearchRateLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"searchRateLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SearchRateLimitUnauthenticated() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"searchRateLimitUnauthenticated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SearchRateLimitUnauthenticatedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"searchRateLimitUnauthenticatedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SendUserConfirmationEmail() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendUserConfirmationEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SendUserConfirmationEmailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendUserConfirmationEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SessionExpireDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionExpireDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SessionExpireDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionExpireDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SharedRunnersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharedRunnersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SharedRunnersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharedRunnersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SharedRunnersMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sharedRunnersMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SharedRunnersMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sharedRunnersMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SharedRunnersText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedRunnersText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SharedRunnersTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedRunnersTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SidekiqJobLimiterCompressionThresholdBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sidekiqJobLimiterCompressionThresholdBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SidekiqJobLimiterCompressionThresholdBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sidekiqJobLimiterCompressionThresholdBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SidekiqJobLimiterLimitBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sidekiqJobLimiterLimitBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SidekiqJobLimiterLimitBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sidekiqJobLimiterLimitBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SidekiqJobLimiterMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sidekiqJobLimiterMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SidekiqJobLimiterModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sidekiqJobLimiterModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SignInText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signInText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SignInTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signInTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SignupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"signupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SignupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"signupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SlackAppEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slackAppEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SlackAppEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slackAppEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SlackAppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackAppId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SlackAppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackAppIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SlackAppSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackAppSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SlackAppSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackAppSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SlackAppSigningSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackAppSigningSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SlackAppSigningSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackAppSigningSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SlackAppVerificationToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackAppVerificationToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SlackAppVerificationTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackAppVerificationTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SnippetSizeLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snippetSizeLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SnippetSizeLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snippetSizeLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SnowplowAppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snowplowAppId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SnowplowAppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snowplowAppIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SnowplowCollectorHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snowplowCollectorHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SnowplowCollectorHostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snowplowCollectorHostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SnowplowCookieDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snowplowCookieDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SnowplowCookieDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snowplowCookieDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SnowplowEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snowplowEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SnowplowEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snowplowEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SourcegraphEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourcegraphEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SourcegraphEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourcegraphEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SourcegraphPublicOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourcegraphPublicOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SourcegraphPublicOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourcegraphPublicOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SourcegraphUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcegraphUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SourcegraphUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcegraphUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SpamCheckApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spamCheckApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SpamCheckApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spamCheckApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SpamCheckEndpointEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spamCheckEndpointEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SpamCheckEndpointEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spamCheckEndpointEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SpamCheckEndpointUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spamCheckEndpointUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SpamCheckEndpointUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spamCheckEndpointUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SuggestPipelineEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suggestPipelineEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) SuggestPipelineEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suggestPipelineEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) TerminalMaxSessionTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminalMaxSessionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) TerminalMaxSessionTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminalMaxSessionTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) Terms() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) TermsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"termsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleAuthenticatedApiEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"throttleAuthenticatedApiEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleAuthenticatedApiEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"throttleAuthenticatedApiEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleAuthenticatedApiPeriodInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttleAuthenticatedApiPeriodInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleAuthenticatedApiPeriodInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttleAuthenticatedApiPeriodInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleAuthenticatedApiRequestsPerPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttleAuthenticatedApiRequestsPerPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleAuthenticatedApiRequestsPerPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttleAuthenticatedApiRequestsPerPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleAuthenticatedPackagesApiEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"throttleAuthenticatedPackagesApiEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleAuthenticatedPackagesApiEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"throttleAuthenticatedPackagesApiEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleAuthenticatedPackagesApiPeriodInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttleAuthenticatedPackagesApiPeriodInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleAuthenticatedPackagesApiPeriodInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttleAuthenticatedPackagesApiPeriodInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleAuthenticatedPackagesApiRequestsPerPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttleAuthenticatedPackagesApiRequestsPerPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleAuthenticatedPackagesApiRequestsPerPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttleAuthenticatedPackagesApiRequestsPerPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleAuthenticatedWebEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"throttleAuthenticatedWebEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleAuthenticatedWebEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"throttleAuthenticatedWebEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleAuthenticatedWebPeriodInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttleAuthenticatedWebPeriodInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleAuthenticatedWebPeriodInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttleAuthenticatedWebPeriodInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleAuthenticatedWebRequestsPerPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttleAuthenticatedWebRequestsPerPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleAuthenticatedWebRequestsPerPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttleAuthenticatedWebRequestsPerPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleUnauthenticatedApiEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"throttleUnauthenticatedApiEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleUnauthenticatedApiEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"throttleUnauthenticatedApiEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleUnauthenticatedApiPeriodInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttleUnauthenticatedApiPeriodInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleUnauthenticatedApiPeriodInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttleUnauthenticatedApiPeriodInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleUnauthenticatedApiRequestsPerPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttleUnauthenticatedApiRequestsPerPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleUnauthenticatedApiRequestsPerPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttleUnauthenticatedApiRequestsPerPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleUnauthenticatedPackagesApiEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"throttleUnauthenticatedPackagesApiEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleUnauthenticatedPackagesApiEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"throttleUnauthenticatedPackagesApiEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleUnauthenticatedPackagesApiPeriodInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttleUnauthenticatedPackagesApiPeriodInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleUnauthenticatedPackagesApiPeriodInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttleUnauthenticatedPackagesApiPeriodInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleUnauthenticatedPackagesApiRequestsPerPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttleUnauthenticatedPackagesApiRequestsPerPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleUnauthenticatedPackagesApiRequestsPerPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttleUnauthenticatedPackagesApiRequestsPerPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleUnauthenticatedWebEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"throttleUnauthenticatedWebEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleUnauthenticatedWebEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"throttleUnauthenticatedWebEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleUnauthenticatedWebPeriodInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttleUnauthenticatedWebPeriodInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleUnauthenticatedWebPeriodInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttleUnauthenticatedWebPeriodInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleUnauthenticatedWebRequestsPerPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttleUnauthenticatedWebRequestsPerPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ThrottleUnauthenticatedWebRequestsPerPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttleUnauthenticatedWebRequestsPerPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) TimeTrackingLimitToHours() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeTrackingLimitToHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) TimeTrackingLimitToHoursInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeTrackingLimitToHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) TwoFactorGracePeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"twoFactorGracePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) TwoFactorGracePeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"twoFactorGracePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) UniqueIpsLimitEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uniqueIpsLimitEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) UniqueIpsLimitEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uniqueIpsLimitEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) UniqueIpsLimitPerUser() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uniqueIpsLimitPerUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) UniqueIpsLimitPerUserInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uniqueIpsLimitPerUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) UniqueIpsLimitTimeWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uniqueIpsLimitTimeWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) UniqueIpsLimitTimeWindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uniqueIpsLimitTimeWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) UsagePingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usagePingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) UsagePingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usagePingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) UserDeactivationEmailsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDeactivationEmailsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) UserDeactivationEmailsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDeactivationEmailsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) UserDefaultExternal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDefaultExternal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) UserDefaultExternalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDefaultExternalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) UserDefaultInternalRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDefaultInternalRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) UserDefaultInternalRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDefaultInternalRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) UserOauthApplications() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userOauthApplications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) UserOauthApplicationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userOauthApplicationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) UserShowAddSshKeyMessage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userShowAddSshKeyMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) UserShowAddSshKeyMessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userShowAddSshKeyMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) VersionCheckEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versionCheckEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) VersionCheckEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versionCheckEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) WebIdeClientsidePreviewEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webIdeClientsidePreviewEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) WebIdeClientsidePreviewEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webIdeClientsidePreviewEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) WhatsNewVariant() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whatsNewVariant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) WhatsNewVariantInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whatsNewVariantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) WikiPageMaxContentBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"wikiPageMaxContentBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) WikiPageMaxContentBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"wikiPageMaxContentBytesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.9.1/docs/resources/application_settings gitlab_application_settings} Resource.
func NewApplicationSettings(scope constructs.Construct, id *string, config *ApplicationSettingsConfig) ApplicationSettings {
	_init_.Initialize()

	if err := validateNewApplicationSettingsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationSettings{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.applicationSettings.ApplicationSettings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.9.1/docs/resources/application_settings gitlab_application_settings} Resource.
func NewApplicationSettings_Override(a ApplicationSettings, scope constructs.Construct, id *string, config *ApplicationSettingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.applicationSettings.ApplicationSettings",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetAbuseNotificationEmail(val *string) {
	if err := j.validateSetAbuseNotificationEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"abuseNotificationEmail",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetAdminMode(val interface{}) {
	if err := j.validateSetAdminModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminMode",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetAfterSignOutPath(val *string) {
	if err := j.validateSetAfterSignOutPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"afterSignOutPath",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetAfterSignUpText(val *string) {
	if err := j.validateSetAfterSignUpTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"afterSignUpText",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetAkismetApiKey(val *string) {
	if err := j.validateSetAkismetApiKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"akismetApiKey",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetAkismetEnabled(val interface{}) {
	if err := j.validateSetAkismetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"akismetEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetAllowGroupOwnersToManageLdap(val interface{}) {
	if err := j.validateSetAllowGroupOwnersToManageLdapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowGroupOwnersToManageLdap",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetAllowLocalRequestsFromSystemHooks(val interface{}) {
	if err := j.validateSetAllowLocalRequestsFromSystemHooksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowLocalRequestsFromSystemHooks",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetAllowLocalRequestsFromWebHooksAndServices(val interface{}) {
	if err := j.validateSetAllowLocalRequestsFromWebHooksAndServicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowLocalRequestsFromWebHooksAndServices",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetArchiveBuildsInHumanReadable(val *string) {
	if err := j.validateSetArchiveBuildsInHumanReadableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"archiveBuildsInHumanReadable",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetAssetProxyAllowlist(val *[]*string) {
	if err := j.validateSetAssetProxyAllowlistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetProxyAllowlist",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetAssetProxyEnabled(val interface{}) {
	if err := j.validateSetAssetProxyEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetProxyEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetAssetProxySecretKey(val *string) {
	if err := j.validateSetAssetProxySecretKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetProxySecretKey",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetAssetProxyUrl(val *string) {
	if err := j.validateSetAssetProxyUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetProxyUrl",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetAuthorizedKeysEnabled(val interface{}) {
	if err := j.validateSetAuthorizedKeysEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizedKeysEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetAutoDevopsDomain(val *string) {
	if err := j.validateSetAutoDevopsDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoDevopsDomain",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetAutoDevopsEnabled(val interface{}) {
	if err := j.validateSetAutoDevopsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoDevopsEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetAutomaticPurchasedStorageAllocation(val interface{}) {
	if err := j.validateSetAutomaticPurchasedStorageAllocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticPurchasedStorageAllocation",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetCanCreateGroup(val interface{}) {
	if err := j.validateSetCanCreateGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"canCreateGroup",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetCheckNamespacePlan(val interface{}) {
	if err := j.validateSetCheckNamespacePlanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkNamespacePlan",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetCommitEmailHostname(val *string) {
	if err := j.validateSetCommitEmailHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitEmailHostname",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetContainerExpirationPoliciesEnableHistoricEntries(val interface{}) {
	if err := j.validateSetContainerExpirationPoliciesEnableHistoricEntriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerExpirationPoliciesEnableHistoricEntries",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetContainerRegistryCleanupTagsServiceMaxListSize(val *float64) {
	if err := j.validateSetContainerRegistryCleanupTagsServiceMaxListSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerRegistryCleanupTagsServiceMaxListSize",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetContainerRegistryDeleteTagsServiceTimeout(val *float64) {
	if err := j.validateSetContainerRegistryDeleteTagsServiceTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerRegistryDeleteTagsServiceTimeout",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetContainerRegistryExpirationPoliciesCaching(val interface{}) {
	if err := j.validateSetContainerRegistryExpirationPoliciesCachingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerRegistryExpirationPoliciesCaching",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetContainerRegistryExpirationPoliciesWorkerCapacity(val *float64) {
	if err := j.validateSetContainerRegistryExpirationPoliciesWorkerCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerRegistryExpirationPoliciesWorkerCapacity",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetContainerRegistryTokenExpireDelay(val *float64) {
	if err := j.validateSetContainerRegistryTokenExpireDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerRegistryTokenExpireDelay",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetDeactivateDormantUsers(val interface{}) {
	if err := j.validateSetDeactivateDormantUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deactivateDormantUsers",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetDefaultArtifactsExpireIn(val *string) {
	if err := j.validateSetDefaultArtifactsExpireInParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultArtifactsExpireIn",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetDefaultBranchName(val *string) {
	if err := j.validateSetDefaultBranchNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultBranchName",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetDefaultBranchProtection(val *float64) {
	if err := j.validateSetDefaultBranchProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultBranchProtection",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetDefaultCiConfigPath(val *string) {
	if err := j.validateSetDefaultCiConfigPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultCiConfigPath",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetDefaultGroupVisibility(val *string) {
	if err := j.validateSetDefaultGroupVisibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultGroupVisibility",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetDefaultProjectCreation(val *float64) {
	if err := j.validateSetDefaultProjectCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultProjectCreation",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetDefaultProjectsLimit(val *float64) {
	if err := j.validateSetDefaultProjectsLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultProjectsLimit",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetDefaultProjectVisibility(val *string) {
	if err := j.validateSetDefaultProjectVisibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultProjectVisibility",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetDefaultSnippetVisibility(val *string) {
	if err := j.validateSetDefaultSnippetVisibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSnippetVisibility",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetDeleteInactiveProjects(val interface{}) {
	if err := j.validateSetDeleteInactiveProjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteInactiveProjects",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetDeletionAdjournedPeriod(val *float64) {
	if err := j.validateSetDeletionAdjournedPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionAdjournedPeriod",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetDiffMaxFiles(val *float64) {
	if err := j.validateSetDiffMaxFilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diffMaxFiles",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetDiffMaxLines(val *float64) {
	if err := j.validateSetDiffMaxLinesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diffMaxLines",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetDiffMaxPatchBytes(val *float64) {
	if err := j.validateSetDiffMaxPatchBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diffMaxPatchBytes",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetDisabledOauthSignInSources(val *[]*string) {
	if err := j.validateSetDisabledOauthSignInSourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabledOauthSignInSources",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetDisableFeedToken(val interface{}) {
	if err := j.validateSetDisableFeedTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableFeedToken",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetDnsRebindingProtectionEnabled(val interface{}) {
	if err := j.validateSetDnsRebindingProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsRebindingProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetDomainAllowlist(val *[]*string) {
	if err := j.validateSetDomainAllowlistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainAllowlist",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetDomainDenylist(val *[]*string) {
	if err := j.validateSetDomainDenylistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainDenylist",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetDomainDenylistEnabled(val interface{}) {
	if err := j.validateSetDomainDenylistEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainDenylistEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetDsaKeyRestriction(val *float64) {
	if err := j.validateSetDsaKeyRestrictionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dsaKeyRestriction",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetEcdsaKeyRestriction(val *float64) {
	if err := j.validateSetEcdsaKeyRestrictionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ecdsaKeyRestriction",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetEcdsaSkKeyRestriction(val *float64) {
	if err := j.validateSetEcdsaSkKeyRestrictionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ecdsaSkKeyRestriction",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetEd25519KeyRestriction(val *float64) {
	if err := j.validateSetEd25519KeyRestrictionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ed25519KeyRestriction",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetEd25519SkKeyRestriction(val *float64) {
	if err := j.validateSetEd25519SkKeyRestrictionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ed25519SkKeyRestriction",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetEksAccessKeyId(val *string) {
	if err := j.validateSetEksAccessKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eksAccessKeyId",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetEksAccountId(val *string) {
	if err := j.validateSetEksAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eksAccountId",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetEksIntegrationEnabled(val interface{}) {
	if err := j.validateSetEksIntegrationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eksIntegrationEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetEksSecretAccessKey(val *string) {
	if err := j.validateSetEksSecretAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eksSecretAccessKey",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetElasticsearchAws(val interface{}) {
	if err := j.validateSetElasticsearchAwsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticsearchAws",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetElasticsearchAwsAccessKey(val *string) {
	if err := j.validateSetElasticsearchAwsAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticsearchAwsAccessKey",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetElasticsearchAwsRegion(val *string) {
	if err := j.validateSetElasticsearchAwsRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticsearchAwsRegion",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetElasticsearchAwsSecretAccessKey(val *string) {
	if err := j.validateSetElasticsearchAwsSecretAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticsearchAwsSecretAccessKey",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetElasticsearchIndexedFieldLengthLimit(val *float64) {
	if err := j.validateSetElasticsearchIndexedFieldLengthLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticsearchIndexedFieldLengthLimit",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetElasticsearchIndexedFileSizeLimitKb(val *float64) {
	if err := j.validateSetElasticsearchIndexedFileSizeLimitKbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticsearchIndexedFileSizeLimitKb",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetElasticsearchIndexing(val interface{}) {
	if err := j.validateSetElasticsearchIndexingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticsearchIndexing",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetElasticsearchLimitIndexing(val interface{}) {
	if err := j.validateSetElasticsearchLimitIndexingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticsearchLimitIndexing",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetElasticsearchMaxBulkConcurrency(val *float64) {
	if err := j.validateSetElasticsearchMaxBulkConcurrencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticsearchMaxBulkConcurrency",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetElasticsearchMaxBulkSizeMb(val *float64) {
	if err := j.validateSetElasticsearchMaxBulkSizeMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticsearchMaxBulkSizeMb",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetElasticsearchNamespaceIds(val *[]*float64) {
	if err := j.validateSetElasticsearchNamespaceIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticsearchNamespaceIds",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetElasticsearchPassword(val *string) {
	if err := j.validateSetElasticsearchPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticsearchPassword",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetElasticsearchProjectIds(val *[]*float64) {
	if err := j.validateSetElasticsearchProjectIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticsearchProjectIds",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetElasticsearchSearch(val interface{}) {
	if err := j.validateSetElasticsearchSearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticsearchSearch",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetElasticsearchUrl(val *[]*string) {
	if err := j.validateSetElasticsearchUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticsearchUrl",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetElasticsearchUsername(val *string) {
	if err := j.validateSetElasticsearchUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticsearchUsername",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetEmailAdditionalText(val *string) {
	if err := j.validateSetEmailAdditionalTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailAdditionalText",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetEmailAuthorInBody(val interface{}) {
	if err := j.validateSetEmailAuthorInBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailAuthorInBody",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetEnabledGitAccessProtocol(val *string) {
	if err := j.validateSetEnabledGitAccessProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledGitAccessProtocol",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetEnforceNamespaceStorageLimit(val interface{}) {
	if err := j.validateSetEnforceNamespaceStorageLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceNamespaceStorageLimit",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetEnforceTerms(val interface{}) {
	if err := j.validateSetEnforceTermsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceTerms",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetExternalAuthClientCert(val *string) {
	if err := j.validateSetExternalAuthClientCertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalAuthClientCert",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetExternalAuthClientKey(val *string) {
	if err := j.validateSetExternalAuthClientKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalAuthClientKey",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetExternalAuthClientKeyPass(val *string) {
	if err := j.validateSetExternalAuthClientKeyPassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalAuthClientKeyPass",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetExternalAuthorizationServiceDefaultLabel(val *string) {
	if err := j.validateSetExternalAuthorizationServiceDefaultLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalAuthorizationServiceDefaultLabel",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetExternalAuthorizationServiceEnabled(val interface{}) {
	if err := j.validateSetExternalAuthorizationServiceEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalAuthorizationServiceEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetExternalAuthorizationServiceTimeout(val *float64) {
	if err := j.validateSetExternalAuthorizationServiceTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalAuthorizationServiceTimeout",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetExternalAuthorizationServiceUrl(val *string) {
	if err := j.validateSetExternalAuthorizationServiceUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalAuthorizationServiceUrl",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetExternalPipelineValidationServiceTimeout(val *float64) {
	if err := j.validateSetExternalPipelineValidationServiceTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalPipelineValidationServiceTimeout",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetExternalPipelineValidationServiceToken(val *string) {
	if err := j.validateSetExternalPipelineValidationServiceTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalPipelineValidationServiceToken",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetExternalPipelineValidationServiceUrl(val *string) {
	if err := j.validateSetExternalPipelineValidationServiceUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalPipelineValidationServiceUrl",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetFileTemplateProjectId(val *float64) {
	if err := j.validateSetFileTemplateProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileTemplateProjectId",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetFirstDayOfWeek(val *float64) {
	if err := j.validateSetFirstDayOfWeekParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstDayOfWeek",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetGeoNodeAllowedIps(val *string) {
	if err := j.validateSetGeoNodeAllowedIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geoNodeAllowedIps",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetGeoStatusTimeout(val *float64) {
	if err := j.validateSetGeoStatusTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geoStatusTimeout",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetGitalyTimeoutDefault(val *float64) {
	if err := j.validateSetGitalyTimeoutDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gitalyTimeoutDefault",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetGitalyTimeoutFast(val *float64) {
	if err := j.validateSetGitalyTimeoutFastParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gitalyTimeoutFast",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetGitalyTimeoutMedium(val *float64) {
	if err := j.validateSetGitalyTimeoutMediumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gitalyTimeoutMedium",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetGitRateLimitUsersAllowlist(val *[]*string) {
	if err := j.validateSetGitRateLimitUsersAllowlistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gitRateLimitUsersAllowlist",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetGitTwoFactorSessionExpiry(val *float64) {
	if err := j.validateSetGitTwoFactorSessionExpiryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gitTwoFactorSessionExpiry",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetGrafanaEnabled(val interface{}) {
	if err := j.validateSetGrafanaEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grafanaEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetGrafanaUrl(val *string) {
	if err := j.validateSetGrafanaUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grafanaUrl",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetGravatarEnabled(val interface{}) {
	if err := j.validateSetGravatarEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gravatarEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetGroupOwnersCanManageDefaultBranchProtection(val interface{}) {
	if err := j.validateSetGroupOwnersCanManageDefaultBranchProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupOwnersCanManageDefaultBranchProtection",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetHashedStorageEnabled(val interface{}) {
	if err := j.validateSetHashedStorageEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hashedStorageEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetHelpPageHideCommercialContent(val interface{}) {
	if err := j.validateSetHelpPageHideCommercialContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"helpPageHideCommercialContent",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetHelpPageSupportUrl(val *string) {
	if err := j.validateSetHelpPageSupportUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"helpPageSupportUrl",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetHelpPageText(val *string) {
	if err := j.validateSetHelpPageTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"helpPageText",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetHelpText(val *string) {
	if err := j.validateSetHelpTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"helpText",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetHideThirdPartyOffers(val interface{}) {
	if err := j.validateSetHideThirdPartyOffersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideThirdPartyOffers",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetHomePageUrl(val *string) {
	if err := j.validateSetHomePageUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"homePageUrl",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetHousekeepingEnabled(val interface{}) {
	if err := j.validateSetHousekeepingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"housekeepingEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetHousekeepingFullRepackPeriod(val *float64) {
	if err := j.validateSetHousekeepingFullRepackPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"housekeepingFullRepackPeriod",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetHousekeepingGcPeriod(val *float64) {
	if err := j.validateSetHousekeepingGcPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"housekeepingGcPeriod",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetHousekeepingIncrementalRepackPeriod(val *float64) {
	if err := j.validateSetHousekeepingIncrementalRepackPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"housekeepingIncrementalRepackPeriod",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetHousekeepingOptimizeRepositoryPeriod(val *float64) {
	if err := j.validateSetHousekeepingOptimizeRepositoryPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"housekeepingOptimizeRepositoryPeriod",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetHtmlEmailsEnabled(val interface{}) {
	if err := j.validateSetHtmlEmailsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"htmlEmailsEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetImportSources(val *[]*string) {
	if err := j.validateSetImportSourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importSources",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetInactiveProjectsDeleteAfterMonths(val *float64) {
	if err := j.validateSetInactiveProjectsDeleteAfterMonthsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inactiveProjectsDeleteAfterMonths",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetInactiveProjectsMinSizeMb(val *float64) {
	if err := j.validateSetInactiveProjectsMinSizeMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inactiveProjectsMinSizeMb",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetInactiveProjectsSendWarningEmailAfterMonths(val *float64) {
	if err := j.validateSetInactiveProjectsSendWarningEmailAfterMonthsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inactiveProjectsSendWarningEmailAfterMonths",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetInProductMarketingEmailsEnabled(val interface{}) {
	if err := j.validateSetInProductMarketingEmailsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inProductMarketingEmailsEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetInvisibleCaptchaEnabled(val interface{}) {
	if err := j.validateSetInvisibleCaptchaEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"invisibleCaptchaEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetIssuesCreateLimit(val *float64) {
	if err := j.validateSetIssuesCreateLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuesCreateLimit",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetKeepLatestArtifact(val interface{}) {
	if err := j.validateSetKeepLatestArtifactParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepLatestArtifact",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetLocalMarkdownVersion(val *float64) {
	if err := j.validateSetLocalMarkdownVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localMarkdownVersion",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetMailgunEventsEnabled(val interface{}) {
	if err := j.validateSetMailgunEventsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mailgunEventsEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetMailgunSigningKey(val *string) {
	if err := j.validateSetMailgunSigningKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mailgunSigningKey",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetMaintenanceMode(val interface{}) {
	if err := j.validateSetMaintenanceModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceMode",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetMaintenanceModeMessage(val *string) {
	if err := j.validateSetMaintenanceModeMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceModeMessage",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetMaxArtifactsSize(val *float64) {
	if err := j.validateSetMaxArtifactsSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxArtifactsSize",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetMaxAttachmentSize(val *float64) {
	if err := j.validateSetMaxAttachmentSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAttachmentSize",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetMaxExportSize(val *float64) {
	if err := j.validateSetMaxExportSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxExportSize",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetMaxImportSize(val *float64) {
	if err := j.validateSetMaxImportSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxImportSize",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetMaxNumberOfRepositoryDownloads(val *float64) {
	if err := j.validateSetMaxNumberOfRepositoryDownloadsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxNumberOfRepositoryDownloads",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetMaxNumberOfRepositoryDownloadsWithinTimePeriod(val *float64) {
	if err := j.validateSetMaxNumberOfRepositoryDownloadsWithinTimePeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxNumberOfRepositoryDownloadsWithinTimePeriod",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetMaxPagesSize(val *float64) {
	if err := j.validateSetMaxPagesSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPagesSize",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetMaxPersonalAccessTokenLifetime(val *float64) {
	if err := j.validateSetMaxPersonalAccessTokenLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPersonalAccessTokenLifetime",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetMaxSshKeyLifetime(val *float64) {
	if err := j.validateSetMaxSshKeyLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSshKeyLifetime",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetMetricsMethodCallThreshold(val *float64) {
	if err := j.validateSetMetricsMethodCallThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsMethodCallThreshold",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetMirrorAvailable(val interface{}) {
	if err := j.validateSetMirrorAvailableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mirrorAvailable",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetMirrorCapacityThreshold(val *float64) {
	if err := j.validateSetMirrorCapacityThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mirrorCapacityThreshold",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetMirrorMaxCapacity(val *float64) {
	if err := j.validateSetMirrorMaxCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mirrorMaxCapacity",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetMirrorMaxDelay(val *float64) {
	if err := j.validateSetMirrorMaxDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mirrorMaxDelay",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetNpmPackageRequestsForwarding(val interface{}) {
	if err := j.validateSetNpmPackageRequestsForwardingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"npmPackageRequestsForwarding",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetOutboundLocalRequestsWhitelist(val *[]*string) {
	if err := j.validateSetOutboundLocalRequestsWhitelistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outboundLocalRequestsWhitelist",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetPackageRegistryCleanupPoliciesWorkerCapacity(val *float64) {
	if err := j.validateSetPackageRegistryCleanupPoliciesWorkerCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packageRegistryCleanupPoliciesWorkerCapacity",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetPagesDomainVerificationEnabled(val interface{}) {
	if err := j.validateSetPagesDomainVerificationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pagesDomainVerificationEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetPasswordAuthenticationEnabledForGit(val interface{}) {
	if err := j.validateSetPasswordAuthenticationEnabledForGitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordAuthenticationEnabledForGit",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetPasswordAuthenticationEnabledForWeb(val interface{}) {
	if err := j.validateSetPasswordAuthenticationEnabledForWebParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordAuthenticationEnabledForWeb",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetPasswordLowercaseRequired(val interface{}) {
	if err := j.validateSetPasswordLowercaseRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordLowercaseRequired",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetPasswordNumberRequired(val interface{}) {
	if err := j.validateSetPasswordNumberRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordNumberRequired",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetPasswordSymbolRequired(val interface{}) {
	if err := j.validateSetPasswordSymbolRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordSymbolRequired",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetPasswordUppercaseRequired(val interface{}) {
	if err := j.validateSetPasswordUppercaseRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordUppercaseRequired",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetPerformanceBarAllowedGroupPath(val *string) {
	if err := j.validateSetPerformanceBarAllowedGroupPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceBarAllowedGroupPath",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetPersonalAccessTokenPrefix(val *string) {
	if err := j.validateSetPersonalAccessTokenPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"personalAccessTokenPrefix",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetPipelineLimitPerProjectUserSha(val *float64) {
	if err := j.validateSetPipelineLimitPerProjectUserShaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipelineLimitPerProjectUserSha",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetPlantumlEnabled(val interface{}) {
	if err := j.validateSetPlantumlEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"plantumlEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetPlantumlUrl(val *string) {
	if err := j.validateSetPlantumlUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"plantumlUrl",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetPollingIntervalMultiplier(val *float64) {
	if err := j.validateSetPollingIntervalMultiplierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pollingIntervalMultiplier",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetProjectExportEnabled(val interface{}) {
	if err := j.validateSetProjectExportEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectExportEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetPrometheusMetricsEnabled(val interface{}) {
	if err := j.validateSetPrometheusMetricsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prometheusMetricsEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetProtectedCiVariables(val interface{}) {
	if err := j.validateSetProtectedCiVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protectedCiVariables",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetPushEventActivitiesLimit(val *float64) {
	if err := j.validateSetPushEventActivitiesLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pushEventActivitiesLimit",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetPushEventHooksLimit(val *float64) {
	if err := j.validateSetPushEventHooksLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pushEventHooksLimit",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetPypiPackageRequestsForwarding(val interface{}) {
	if err := j.validateSetPypiPackageRequestsForwardingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pypiPackageRequestsForwarding",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetRateLimitingResponseText(val *string) {
	if err := j.validateSetRateLimitingResponseTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rateLimitingResponseText",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetRawBlobRequestLimit(val *float64) {
	if err := j.validateSetRawBlobRequestLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rawBlobRequestLimit",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetRecaptchaEnabled(val interface{}) {
	if err := j.validateSetRecaptchaEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recaptchaEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetRecaptchaPrivateKey(val *string) {
	if err := j.validateSetRecaptchaPrivateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recaptchaPrivateKey",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetRecaptchaSiteKey(val *string) {
	if err := j.validateSetRecaptchaSiteKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recaptchaSiteKey",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetReceiveMaxInputSize(val *float64) {
	if err := j.validateSetReceiveMaxInputSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"receiveMaxInputSize",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetRepositoryChecksEnabled(val interface{}) {
	if err := j.validateSetRepositoryChecksEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repositoryChecksEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetRepositorySizeLimit(val *float64) {
	if err := j.validateSetRepositorySizeLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repositorySizeLimit",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetRepositoryStorages(val *[]*string) {
	if err := j.validateSetRepositoryStoragesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repositoryStorages",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetRepositoryStoragesWeighted(val *map[string]*float64) {
	if err := j.validateSetRepositoryStoragesWeightedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repositoryStoragesWeighted",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetRequireAdminApprovalAfterUserSignup(val interface{}) {
	if err := j.validateSetRequireAdminApprovalAfterUserSignupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireAdminApprovalAfterUserSignup",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetRequireTwoFactorAuthentication(val interface{}) {
	if err := j.validateSetRequireTwoFactorAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireTwoFactorAuthentication",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetRestrictedVisibilityLevels(val *[]*string) {
	if err := j.validateSetRestrictedVisibilityLevelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictedVisibilityLevels",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetRsaKeyRestriction(val *float64) {
	if err := j.validateSetRsaKeyRestrictionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rsaKeyRestriction",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSearchRateLimit(val *float64) {
	if err := j.validateSetSearchRateLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchRateLimit",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSearchRateLimitUnauthenticated(val *float64) {
	if err := j.validateSetSearchRateLimitUnauthenticatedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchRateLimitUnauthenticated",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSendUserConfirmationEmail(val interface{}) {
	if err := j.validateSetSendUserConfirmationEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendUserConfirmationEmail",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSessionExpireDelay(val *float64) {
	if err := j.validateSetSessionExpireDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionExpireDelay",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSharedRunnersEnabled(val interface{}) {
	if err := j.validateSetSharedRunnersEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedRunnersEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSharedRunnersMinutes(val *float64) {
	if err := j.validateSetSharedRunnersMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedRunnersMinutes",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSharedRunnersText(val *string) {
	if err := j.validateSetSharedRunnersTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedRunnersText",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSidekiqJobLimiterCompressionThresholdBytes(val *float64) {
	if err := j.validateSetSidekiqJobLimiterCompressionThresholdBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sidekiqJobLimiterCompressionThresholdBytes",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSidekiqJobLimiterLimitBytes(val *float64) {
	if err := j.validateSetSidekiqJobLimiterLimitBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sidekiqJobLimiterLimitBytes",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSidekiqJobLimiterMode(val *string) {
	if err := j.validateSetSidekiqJobLimiterModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sidekiqJobLimiterMode",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSignInText(val *string) {
	if err := j.validateSetSignInTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signInText",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSignupEnabled(val interface{}) {
	if err := j.validateSetSignupEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signupEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSlackAppEnabled(val interface{}) {
	if err := j.validateSetSlackAppEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slackAppEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSlackAppId(val *string) {
	if err := j.validateSetSlackAppIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slackAppId",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSlackAppSecret(val *string) {
	if err := j.validateSetSlackAppSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slackAppSecret",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSlackAppSigningSecret(val *string) {
	if err := j.validateSetSlackAppSigningSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slackAppSigningSecret",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSlackAppVerificationToken(val *string) {
	if err := j.validateSetSlackAppVerificationTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slackAppVerificationToken",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSnippetSizeLimit(val *float64) {
	if err := j.validateSetSnippetSizeLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snippetSizeLimit",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSnowplowAppId(val *string) {
	if err := j.validateSetSnowplowAppIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snowplowAppId",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSnowplowCollectorHostname(val *string) {
	if err := j.validateSetSnowplowCollectorHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snowplowCollectorHostname",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSnowplowCookieDomain(val *string) {
	if err := j.validateSetSnowplowCookieDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snowplowCookieDomain",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSnowplowEnabled(val interface{}) {
	if err := j.validateSetSnowplowEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snowplowEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSourcegraphEnabled(val interface{}) {
	if err := j.validateSetSourcegraphEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcegraphEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSourcegraphPublicOnly(val interface{}) {
	if err := j.validateSetSourcegraphPublicOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcegraphPublicOnly",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSourcegraphUrl(val *string) {
	if err := j.validateSetSourcegraphUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcegraphUrl",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSpamCheckApiKey(val *string) {
	if err := j.validateSetSpamCheckApiKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spamCheckApiKey",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSpamCheckEndpointEnabled(val interface{}) {
	if err := j.validateSetSpamCheckEndpointEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spamCheckEndpointEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSpamCheckEndpointUrl(val *string) {
	if err := j.validateSetSpamCheckEndpointUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spamCheckEndpointUrl",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetSuggestPipelineEnabled(val interface{}) {
	if err := j.validateSetSuggestPipelineEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suggestPipelineEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetTerminalMaxSessionTime(val *float64) {
	if err := j.validateSetTerminalMaxSessionTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminalMaxSessionTime",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetTerms(val *string) {
	if err := j.validateSetTermsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terms",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetThrottleAuthenticatedApiEnabled(val interface{}) {
	if err := j.validateSetThrottleAuthenticatedApiEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttleAuthenticatedApiEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetThrottleAuthenticatedApiPeriodInSeconds(val *float64) {
	if err := j.validateSetThrottleAuthenticatedApiPeriodInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttleAuthenticatedApiPeriodInSeconds",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetThrottleAuthenticatedApiRequestsPerPeriod(val *float64) {
	if err := j.validateSetThrottleAuthenticatedApiRequestsPerPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttleAuthenticatedApiRequestsPerPeriod",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetThrottleAuthenticatedPackagesApiEnabled(val interface{}) {
	if err := j.validateSetThrottleAuthenticatedPackagesApiEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttleAuthenticatedPackagesApiEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetThrottleAuthenticatedPackagesApiPeriodInSeconds(val *float64) {
	if err := j.validateSetThrottleAuthenticatedPackagesApiPeriodInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttleAuthenticatedPackagesApiPeriodInSeconds",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetThrottleAuthenticatedPackagesApiRequestsPerPeriod(val *float64) {
	if err := j.validateSetThrottleAuthenticatedPackagesApiRequestsPerPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttleAuthenticatedPackagesApiRequestsPerPeriod",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetThrottleAuthenticatedWebEnabled(val interface{}) {
	if err := j.validateSetThrottleAuthenticatedWebEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttleAuthenticatedWebEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetThrottleAuthenticatedWebPeriodInSeconds(val *float64) {
	if err := j.validateSetThrottleAuthenticatedWebPeriodInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttleAuthenticatedWebPeriodInSeconds",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetThrottleAuthenticatedWebRequestsPerPeriod(val *float64) {
	if err := j.validateSetThrottleAuthenticatedWebRequestsPerPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttleAuthenticatedWebRequestsPerPeriod",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetThrottleUnauthenticatedApiEnabled(val interface{}) {
	if err := j.validateSetThrottleUnauthenticatedApiEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttleUnauthenticatedApiEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetThrottleUnauthenticatedApiPeriodInSeconds(val *float64) {
	if err := j.validateSetThrottleUnauthenticatedApiPeriodInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttleUnauthenticatedApiPeriodInSeconds",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetThrottleUnauthenticatedApiRequestsPerPeriod(val *float64) {
	if err := j.validateSetThrottleUnauthenticatedApiRequestsPerPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttleUnauthenticatedApiRequestsPerPeriod",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetThrottleUnauthenticatedPackagesApiEnabled(val interface{}) {
	if err := j.validateSetThrottleUnauthenticatedPackagesApiEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttleUnauthenticatedPackagesApiEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetThrottleUnauthenticatedPackagesApiPeriodInSeconds(val *float64) {
	if err := j.validateSetThrottleUnauthenticatedPackagesApiPeriodInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttleUnauthenticatedPackagesApiPeriodInSeconds",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetThrottleUnauthenticatedPackagesApiRequestsPerPeriod(val *float64) {
	if err := j.validateSetThrottleUnauthenticatedPackagesApiRequestsPerPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttleUnauthenticatedPackagesApiRequestsPerPeriod",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetThrottleUnauthenticatedWebEnabled(val interface{}) {
	if err := j.validateSetThrottleUnauthenticatedWebEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttleUnauthenticatedWebEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetThrottleUnauthenticatedWebPeriodInSeconds(val *float64) {
	if err := j.validateSetThrottleUnauthenticatedWebPeriodInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttleUnauthenticatedWebPeriodInSeconds",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetThrottleUnauthenticatedWebRequestsPerPeriod(val *float64) {
	if err := j.validateSetThrottleUnauthenticatedWebRequestsPerPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttleUnauthenticatedWebRequestsPerPeriod",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetTimeTrackingLimitToHours(val interface{}) {
	if err := j.validateSetTimeTrackingLimitToHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeTrackingLimitToHours",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetTwoFactorGracePeriod(val *float64) {
	if err := j.validateSetTwoFactorGracePeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoFactorGracePeriod",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetUniqueIpsLimitEnabled(val interface{}) {
	if err := j.validateSetUniqueIpsLimitEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uniqueIpsLimitEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetUniqueIpsLimitPerUser(val *float64) {
	if err := j.validateSetUniqueIpsLimitPerUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uniqueIpsLimitPerUser",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetUniqueIpsLimitTimeWindow(val *float64) {
	if err := j.validateSetUniqueIpsLimitTimeWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uniqueIpsLimitTimeWindow",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetUsagePingEnabled(val interface{}) {
	if err := j.validateSetUsagePingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usagePingEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetUserDeactivationEmailsEnabled(val interface{}) {
	if err := j.validateSetUserDeactivationEmailsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userDeactivationEmailsEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetUserDefaultExternal(val interface{}) {
	if err := j.validateSetUserDefaultExternalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userDefaultExternal",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetUserDefaultInternalRegex(val *string) {
	if err := j.validateSetUserDefaultInternalRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userDefaultInternalRegex",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetUserOauthApplications(val interface{}) {
	if err := j.validateSetUserOauthApplicationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userOauthApplications",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetUserShowAddSshKeyMessage(val interface{}) {
	if err := j.validateSetUserShowAddSshKeyMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userShowAddSshKeyMessage",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetVersionCheckEnabled(val interface{}) {
	if err := j.validateSetVersionCheckEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionCheckEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetWebIdeClientsidePreviewEnabled(val interface{}) {
	if err := j.validateSetWebIdeClientsidePreviewEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webIdeClientsidePreviewEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetWhatsNewVariant(val *string) {
	if err := j.validateSetWhatsNewVariantParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whatsNewVariant",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetWikiPageMaxContentBytes(val *float64) {
	if err := j.validateSetWikiPageMaxContentBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wikiPageMaxContentBytes",
		val,
	)
}

// Generates CDKTF code for importing a ApplicationSettings resource upon running "cdktf plan <stack-name>".
func ApplicationSettings_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateApplicationSettings_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.applicationSettings.ApplicationSettings",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ApplicationSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationSettings_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.applicationSettings.ApplicationSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApplicationSettings_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationSettings_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.applicationSettings.ApplicationSettings",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApplicationSettings_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationSettings_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.applicationSettings.ApplicationSettings",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApplicationSettings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.applicationSettings.ApplicationSettings",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApplicationSettings) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ApplicationSettings) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApplicationSettings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ApplicationSettings) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApplicationSettings) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ApplicationSettings) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApplicationSettings) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetAbuseNotificationEmail() {
	_jsii_.InvokeVoid(
		a,
		"resetAbuseNotificationEmail",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetAdminMode() {
	_jsii_.InvokeVoid(
		a,
		"resetAdminMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetAfterSignOutPath() {
	_jsii_.InvokeVoid(
		a,
		"resetAfterSignOutPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetAfterSignUpText() {
	_jsii_.InvokeVoid(
		a,
		"resetAfterSignUpText",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetAkismetApiKey() {
	_jsii_.InvokeVoid(
		a,
		"resetAkismetApiKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetAkismetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetAkismetEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetAllowGroupOwnersToManageLdap() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowGroupOwnersToManageLdap",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetAllowLocalRequestsFromSystemHooks() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowLocalRequestsFromSystemHooks",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetAllowLocalRequestsFromWebHooksAndServices() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowLocalRequestsFromWebHooksAndServices",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetArchiveBuildsInHumanReadable() {
	_jsii_.InvokeVoid(
		a,
		"resetArchiveBuildsInHumanReadable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetAssetProxyAllowlist() {
	_jsii_.InvokeVoid(
		a,
		"resetAssetProxyAllowlist",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetAssetProxyEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetAssetProxyEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetAssetProxySecretKey() {
	_jsii_.InvokeVoid(
		a,
		"resetAssetProxySecretKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetAssetProxyUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetAssetProxyUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetAuthorizedKeysEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizedKeysEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetAutoDevopsDomain() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoDevopsDomain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetAutoDevopsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoDevopsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetAutomaticPurchasedStorageAllocation() {
	_jsii_.InvokeVoid(
		a,
		"resetAutomaticPurchasedStorageAllocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetCanCreateGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetCanCreateGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetCheckNamespacePlan() {
	_jsii_.InvokeVoid(
		a,
		"resetCheckNamespacePlan",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetCommitEmailHostname() {
	_jsii_.InvokeVoid(
		a,
		"resetCommitEmailHostname",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetContainerExpirationPoliciesEnableHistoricEntries() {
	_jsii_.InvokeVoid(
		a,
		"resetContainerExpirationPoliciesEnableHistoricEntries",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetContainerRegistryCleanupTagsServiceMaxListSize() {
	_jsii_.InvokeVoid(
		a,
		"resetContainerRegistryCleanupTagsServiceMaxListSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetContainerRegistryDeleteTagsServiceTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetContainerRegistryDeleteTagsServiceTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetContainerRegistryExpirationPoliciesCaching() {
	_jsii_.InvokeVoid(
		a,
		"resetContainerRegistryExpirationPoliciesCaching",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetContainerRegistryExpirationPoliciesWorkerCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetContainerRegistryExpirationPoliciesWorkerCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetContainerRegistryTokenExpireDelay() {
	_jsii_.InvokeVoid(
		a,
		"resetContainerRegistryTokenExpireDelay",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetDeactivateDormantUsers() {
	_jsii_.InvokeVoid(
		a,
		"resetDeactivateDormantUsers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetDefaultArtifactsExpireIn() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultArtifactsExpireIn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetDefaultBranchName() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultBranchName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetDefaultBranchProtection() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultBranchProtection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetDefaultCiConfigPath() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultCiConfigPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetDefaultGroupVisibility() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultGroupVisibility",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetDefaultProjectCreation() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultProjectCreation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetDefaultProjectsLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultProjectsLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetDefaultProjectVisibility() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultProjectVisibility",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetDefaultSnippetVisibility() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultSnippetVisibility",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetDeleteInactiveProjects() {
	_jsii_.InvokeVoid(
		a,
		"resetDeleteInactiveProjects",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetDeletionAdjournedPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetDeletionAdjournedPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetDiffMaxFiles() {
	_jsii_.InvokeVoid(
		a,
		"resetDiffMaxFiles",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetDiffMaxLines() {
	_jsii_.InvokeVoid(
		a,
		"resetDiffMaxLines",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetDiffMaxPatchBytes() {
	_jsii_.InvokeVoid(
		a,
		"resetDiffMaxPatchBytes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetDisabledOauthSignInSources() {
	_jsii_.InvokeVoid(
		a,
		"resetDisabledOauthSignInSources",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetDisableFeedToken() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableFeedToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetDnsRebindingProtectionEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetDnsRebindingProtectionEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetDomainAllowlist() {
	_jsii_.InvokeVoid(
		a,
		"resetDomainAllowlist",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetDomainDenylist() {
	_jsii_.InvokeVoid(
		a,
		"resetDomainDenylist",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetDomainDenylistEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetDomainDenylistEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetDsaKeyRestriction() {
	_jsii_.InvokeVoid(
		a,
		"resetDsaKeyRestriction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetEcdsaKeyRestriction() {
	_jsii_.InvokeVoid(
		a,
		"resetEcdsaKeyRestriction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetEcdsaSkKeyRestriction() {
	_jsii_.InvokeVoid(
		a,
		"resetEcdsaSkKeyRestriction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetEd25519KeyRestriction() {
	_jsii_.InvokeVoid(
		a,
		"resetEd25519KeyRestriction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetEd25519SkKeyRestriction() {
	_jsii_.InvokeVoid(
		a,
		"resetEd25519SkKeyRestriction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetEksAccessKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetEksAccessKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetEksAccountId() {
	_jsii_.InvokeVoid(
		a,
		"resetEksAccountId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetEksIntegrationEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEksIntegrationEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetEksSecretAccessKey() {
	_jsii_.InvokeVoid(
		a,
		"resetEksSecretAccessKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetElasticsearchAws() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticsearchAws",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetElasticsearchAwsAccessKey() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticsearchAwsAccessKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetElasticsearchAwsRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticsearchAwsRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetElasticsearchAwsSecretAccessKey() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticsearchAwsSecretAccessKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetElasticsearchIndexedFieldLengthLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticsearchIndexedFieldLengthLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetElasticsearchIndexedFileSizeLimitKb() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticsearchIndexedFileSizeLimitKb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetElasticsearchIndexing() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticsearchIndexing",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetElasticsearchLimitIndexing() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticsearchLimitIndexing",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetElasticsearchMaxBulkConcurrency() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticsearchMaxBulkConcurrency",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetElasticsearchMaxBulkSizeMb() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticsearchMaxBulkSizeMb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetElasticsearchNamespaceIds() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticsearchNamespaceIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetElasticsearchPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticsearchPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetElasticsearchProjectIds() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticsearchProjectIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetElasticsearchSearch() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticsearchSearch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetElasticsearchUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticsearchUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetElasticsearchUsername() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticsearchUsername",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetEmailAdditionalText() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailAdditionalText",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetEmailAuthorInBody() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailAuthorInBody",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetEnabledGitAccessProtocol() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabledGitAccessProtocol",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetEnforceNamespaceStorageLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetEnforceNamespaceStorageLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetEnforceTerms() {
	_jsii_.InvokeVoid(
		a,
		"resetEnforceTerms",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetExternalAuthClientCert() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalAuthClientCert",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetExternalAuthClientKey() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalAuthClientKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetExternalAuthClientKeyPass() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalAuthClientKeyPass",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetExternalAuthorizationServiceDefaultLabel() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalAuthorizationServiceDefaultLabel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetExternalAuthorizationServiceEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalAuthorizationServiceEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetExternalAuthorizationServiceTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalAuthorizationServiceTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetExternalAuthorizationServiceUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalAuthorizationServiceUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetExternalPipelineValidationServiceTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalPipelineValidationServiceTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetExternalPipelineValidationServiceToken() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalPipelineValidationServiceToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetExternalPipelineValidationServiceUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalPipelineValidationServiceUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetFileTemplateProjectId() {
	_jsii_.InvokeVoid(
		a,
		"resetFileTemplateProjectId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetFirstDayOfWeek() {
	_jsii_.InvokeVoid(
		a,
		"resetFirstDayOfWeek",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetGeoNodeAllowedIps() {
	_jsii_.InvokeVoid(
		a,
		"resetGeoNodeAllowedIps",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetGeoStatusTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetGeoStatusTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetGitalyTimeoutDefault() {
	_jsii_.InvokeVoid(
		a,
		"resetGitalyTimeoutDefault",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetGitalyTimeoutFast() {
	_jsii_.InvokeVoid(
		a,
		"resetGitalyTimeoutFast",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetGitalyTimeoutMedium() {
	_jsii_.InvokeVoid(
		a,
		"resetGitalyTimeoutMedium",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetGitRateLimitUsersAllowlist() {
	_jsii_.InvokeVoid(
		a,
		"resetGitRateLimitUsersAllowlist",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetGitTwoFactorSessionExpiry() {
	_jsii_.InvokeVoid(
		a,
		"resetGitTwoFactorSessionExpiry",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetGrafanaEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetGrafanaEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetGrafanaUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetGrafanaUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetGravatarEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetGravatarEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetGroupOwnersCanManageDefaultBranchProtection() {
	_jsii_.InvokeVoid(
		a,
		"resetGroupOwnersCanManageDefaultBranchProtection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetHashedStorageEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetHashedStorageEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetHelpPageHideCommercialContent() {
	_jsii_.InvokeVoid(
		a,
		"resetHelpPageHideCommercialContent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetHelpPageSupportUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetHelpPageSupportUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetHelpPageText() {
	_jsii_.InvokeVoid(
		a,
		"resetHelpPageText",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetHelpText() {
	_jsii_.InvokeVoid(
		a,
		"resetHelpText",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetHideThirdPartyOffers() {
	_jsii_.InvokeVoid(
		a,
		"resetHideThirdPartyOffers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetHomePageUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetHomePageUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetHousekeepingEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetHousekeepingEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetHousekeepingFullRepackPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetHousekeepingFullRepackPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetHousekeepingGcPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetHousekeepingGcPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetHousekeepingIncrementalRepackPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetHousekeepingIncrementalRepackPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetHousekeepingOptimizeRepositoryPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetHousekeepingOptimizeRepositoryPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetHtmlEmailsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetHtmlEmailsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetImportSources() {
	_jsii_.InvokeVoid(
		a,
		"resetImportSources",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetInactiveProjectsDeleteAfterMonths() {
	_jsii_.InvokeVoid(
		a,
		"resetInactiveProjectsDeleteAfterMonths",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetInactiveProjectsMinSizeMb() {
	_jsii_.InvokeVoid(
		a,
		"resetInactiveProjectsMinSizeMb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetInactiveProjectsSendWarningEmailAfterMonths() {
	_jsii_.InvokeVoid(
		a,
		"resetInactiveProjectsSendWarningEmailAfterMonths",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetInProductMarketingEmailsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetInProductMarketingEmailsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetInvisibleCaptchaEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetInvisibleCaptchaEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetIssuesCreateLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetIssuesCreateLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetKeepLatestArtifact() {
	_jsii_.InvokeVoid(
		a,
		"resetKeepLatestArtifact",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetLocalMarkdownVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetLocalMarkdownVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetMailgunEventsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetMailgunEventsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetMailgunSigningKey() {
	_jsii_.InvokeVoid(
		a,
		"resetMailgunSigningKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetMaintenanceMode() {
	_jsii_.InvokeVoid(
		a,
		"resetMaintenanceMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetMaintenanceModeMessage() {
	_jsii_.InvokeVoid(
		a,
		"resetMaintenanceModeMessage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetMaxArtifactsSize() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxArtifactsSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetMaxAttachmentSize() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxAttachmentSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetMaxExportSize() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxExportSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetMaxImportSize() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxImportSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetMaxNumberOfRepositoryDownloads() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxNumberOfRepositoryDownloads",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetMaxNumberOfRepositoryDownloadsWithinTimePeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxNumberOfRepositoryDownloadsWithinTimePeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetMaxPagesSize() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxPagesSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetMaxPersonalAccessTokenLifetime() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxPersonalAccessTokenLifetime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetMaxSshKeyLifetime() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxSshKeyLifetime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetMetricsMethodCallThreshold() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricsMethodCallThreshold",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetMirrorAvailable() {
	_jsii_.InvokeVoid(
		a,
		"resetMirrorAvailable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetMirrorCapacityThreshold() {
	_jsii_.InvokeVoid(
		a,
		"resetMirrorCapacityThreshold",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetMirrorMaxCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetMirrorMaxCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetMirrorMaxDelay() {
	_jsii_.InvokeVoid(
		a,
		"resetMirrorMaxDelay",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetNpmPackageRequestsForwarding() {
	_jsii_.InvokeVoid(
		a,
		"resetNpmPackageRequestsForwarding",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetOutboundLocalRequestsWhitelist() {
	_jsii_.InvokeVoid(
		a,
		"resetOutboundLocalRequestsWhitelist",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetPackageRegistryCleanupPoliciesWorkerCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetPackageRegistryCleanupPoliciesWorkerCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetPagesDomainVerificationEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetPagesDomainVerificationEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetPasswordAuthenticationEnabledForGit() {
	_jsii_.InvokeVoid(
		a,
		"resetPasswordAuthenticationEnabledForGit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetPasswordAuthenticationEnabledForWeb() {
	_jsii_.InvokeVoid(
		a,
		"resetPasswordAuthenticationEnabledForWeb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetPasswordLowercaseRequired() {
	_jsii_.InvokeVoid(
		a,
		"resetPasswordLowercaseRequired",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetPasswordNumberRequired() {
	_jsii_.InvokeVoid(
		a,
		"resetPasswordNumberRequired",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetPasswordSymbolRequired() {
	_jsii_.InvokeVoid(
		a,
		"resetPasswordSymbolRequired",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetPasswordUppercaseRequired() {
	_jsii_.InvokeVoid(
		a,
		"resetPasswordUppercaseRequired",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetPerformanceBarAllowedGroupPath() {
	_jsii_.InvokeVoid(
		a,
		"resetPerformanceBarAllowedGroupPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetPersonalAccessTokenPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetPersonalAccessTokenPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetPipelineLimitPerProjectUserSha() {
	_jsii_.InvokeVoid(
		a,
		"resetPipelineLimitPerProjectUserSha",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetPlantumlEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetPlantumlEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetPlantumlUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetPlantumlUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetPollingIntervalMultiplier() {
	_jsii_.InvokeVoid(
		a,
		"resetPollingIntervalMultiplier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetProjectExportEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetProjectExportEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetPrometheusMetricsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetPrometheusMetricsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetProtectedCiVariables() {
	_jsii_.InvokeVoid(
		a,
		"resetProtectedCiVariables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetPushEventActivitiesLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetPushEventActivitiesLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetPushEventHooksLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetPushEventHooksLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetPypiPackageRequestsForwarding() {
	_jsii_.InvokeVoid(
		a,
		"resetPypiPackageRequestsForwarding",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetRateLimitingResponseText() {
	_jsii_.InvokeVoid(
		a,
		"resetRateLimitingResponseText",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetRawBlobRequestLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetRawBlobRequestLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetRecaptchaEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetRecaptchaEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetRecaptchaPrivateKey() {
	_jsii_.InvokeVoid(
		a,
		"resetRecaptchaPrivateKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetRecaptchaSiteKey() {
	_jsii_.InvokeVoid(
		a,
		"resetRecaptchaSiteKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetReceiveMaxInputSize() {
	_jsii_.InvokeVoid(
		a,
		"resetReceiveMaxInputSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetRepositoryChecksEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetRepositoryChecksEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetRepositorySizeLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetRepositorySizeLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetRepositoryStorages() {
	_jsii_.InvokeVoid(
		a,
		"resetRepositoryStorages",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetRepositoryStoragesWeighted() {
	_jsii_.InvokeVoid(
		a,
		"resetRepositoryStoragesWeighted",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetRequireAdminApprovalAfterUserSignup() {
	_jsii_.InvokeVoid(
		a,
		"resetRequireAdminApprovalAfterUserSignup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetRequireTwoFactorAuthentication() {
	_jsii_.InvokeVoid(
		a,
		"resetRequireTwoFactorAuthentication",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetRestrictedVisibilityLevels() {
	_jsii_.InvokeVoid(
		a,
		"resetRestrictedVisibilityLevels",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetRsaKeyRestriction() {
	_jsii_.InvokeVoid(
		a,
		"resetRsaKeyRestriction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSearchRateLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetSearchRateLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSearchRateLimitUnauthenticated() {
	_jsii_.InvokeVoid(
		a,
		"resetSearchRateLimitUnauthenticated",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSendUserConfirmationEmail() {
	_jsii_.InvokeVoid(
		a,
		"resetSendUserConfirmationEmail",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSessionExpireDelay() {
	_jsii_.InvokeVoid(
		a,
		"resetSessionExpireDelay",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSharedRunnersEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetSharedRunnersEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSharedRunnersMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetSharedRunnersMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSharedRunnersText() {
	_jsii_.InvokeVoid(
		a,
		"resetSharedRunnersText",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSidekiqJobLimiterCompressionThresholdBytes() {
	_jsii_.InvokeVoid(
		a,
		"resetSidekiqJobLimiterCompressionThresholdBytes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSidekiqJobLimiterLimitBytes() {
	_jsii_.InvokeVoid(
		a,
		"resetSidekiqJobLimiterLimitBytes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSidekiqJobLimiterMode() {
	_jsii_.InvokeVoid(
		a,
		"resetSidekiqJobLimiterMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSignInText() {
	_jsii_.InvokeVoid(
		a,
		"resetSignInText",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSignupEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetSignupEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSlackAppEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetSlackAppEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSlackAppId() {
	_jsii_.InvokeVoid(
		a,
		"resetSlackAppId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSlackAppSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetSlackAppSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSlackAppSigningSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetSlackAppSigningSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSlackAppVerificationToken() {
	_jsii_.InvokeVoid(
		a,
		"resetSlackAppVerificationToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSnippetSizeLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetSnippetSizeLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSnowplowAppId() {
	_jsii_.InvokeVoid(
		a,
		"resetSnowplowAppId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSnowplowCollectorHostname() {
	_jsii_.InvokeVoid(
		a,
		"resetSnowplowCollectorHostname",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSnowplowCookieDomain() {
	_jsii_.InvokeVoid(
		a,
		"resetSnowplowCookieDomain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSnowplowEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetSnowplowEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSourcegraphEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetSourcegraphEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSourcegraphPublicOnly() {
	_jsii_.InvokeVoid(
		a,
		"resetSourcegraphPublicOnly",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSourcegraphUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetSourcegraphUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSpamCheckApiKey() {
	_jsii_.InvokeVoid(
		a,
		"resetSpamCheckApiKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSpamCheckEndpointEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetSpamCheckEndpointEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSpamCheckEndpointUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetSpamCheckEndpointUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetSuggestPipelineEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetSuggestPipelineEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetTerminalMaxSessionTime() {
	_jsii_.InvokeVoid(
		a,
		"resetTerminalMaxSessionTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetTerms() {
	_jsii_.InvokeVoid(
		a,
		"resetTerms",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetThrottleAuthenticatedApiEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottleAuthenticatedApiEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetThrottleAuthenticatedApiPeriodInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottleAuthenticatedApiPeriodInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetThrottleAuthenticatedApiRequestsPerPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottleAuthenticatedApiRequestsPerPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetThrottleAuthenticatedPackagesApiEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottleAuthenticatedPackagesApiEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetThrottleAuthenticatedPackagesApiPeriodInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottleAuthenticatedPackagesApiPeriodInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetThrottleAuthenticatedPackagesApiRequestsPerPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottleAuthenticatedPackagesApiRequestsPerPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetThrottleAuthenticatedWebEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottleAuthenticatedWebEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetThrottleAuthenticatedWebPeriodInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottleAuthenticatedWebPeriodInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetThrottleAuthenticatedWebRequestsPerPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottleAuthenticatedWebRequestsPerPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetThrottleUnauthenticatedApiEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottleUnauthenticatedApiEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetThrottleUnauthenticatedApiPeriodInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottleUnauthenticatedApiPeriodInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetThrottleUnauthenticatedApiRequestsPerPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottleUnauthenticatedApiRequestsPerPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetThrottleUnauthenticatedPackagesApiEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottleUnauthenticatedPackagesApiEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetThrottleUnauthenticatedPackagesApiPeriodInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottleUnauthenticatedPackagesApiPeriodInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetThrottleUnauthenticatedPackagesApiRequestsPerPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottleUnauthenticatedPackagesApiRequestsPerPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetThrottleUnauthenticatedWebEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottleUnauthenticatedWebEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetThrottleUnauthenticatedWebPeriodInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottleUnauthenticatedWebPeriodInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetThrottleUnauthenticatedWebRequestsPerPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottleUnauthenticatedWebRequestsPerPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetTimeTrackingLimitToHours() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeTrackingLimitToHours",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetTwoFactorGracePeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetTwoFactorGracePeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetUniqueIpsLimitEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetUniqueIpsLimitEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetUniqueIpsLimitPerUser() {
	_jsii_.InvokeVoid(
		a,
		"resetUniqueIpsLimitPerUser",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetUniqueIpsLimitTimeWindow() {
	_jsii_.InvokeVoid(
		a,
		"resetUniqueIpsLimitTimeWindow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetUsagePingEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetUsagePingEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetUserDeactivationEmailsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetUserDeactivationEmailsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetUserDefaultExternal() {
	_jsii_.InvokeVoid(
		a,
		"resetUserDefaultExternal",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetUserDefaultInternalRegex() {
	_jsii_.InvokeVoid(
		a,
		"resetUserDefaultInternalRegex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetUserOauthApplications() {
	_jsii_.InvokeVoid(
		a,
		"resetUserOauthApplications",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetUserShowAddSshKeyMessage() {
	_jsii_.InvokeVoid(
		a,
		"resetUserShowAddSshKeyMessage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetVersionCheckEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetVersionCheckEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetWebIdeClientsidePreviewEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetWebIdeClientsidePreviewEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetWhatsNewVariant() {
	_jsii_.InvokeVoid(
		a,
		"resetWhatsNewVariant",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetWikiPageMaxContentBytes() {
	_jsii_.InvokeVoid(
		a,
		"resetWikiPageMaxContentBytes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

