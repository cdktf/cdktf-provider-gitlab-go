// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package applicationsettings

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationSettings) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationSettings) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_ApplicationSettings) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationSettings) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationSettings) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationSettings) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationSettings) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationSettings) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationSettings) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationSettings) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationSettings) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationSettings) validateImportFromParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationSettings) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationSettings) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationSettings) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (a *jsiiProxy_ApplicationSettings) validateMoveToIdParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationSettings) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationSettings) validatePutDefaultBranchProtectionDefaultsParameters(value *ApplicationSettingsDefaultBranchProtectionDefaults) error {
	return nil
}

func validateApplicationSettings_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateApplicationSettings_IsConstructParameters(x interface{}) error {
	return nil
}

func validateApplicationSettings_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateApplicationSettings_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetAbuseNotificationEmailParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetAdminModeParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetAfterSignOutPathParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetAfterSignUpTextParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetAkismetApiKeyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetAkismetEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetAllowAccountDeletionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetAllowGroupOwnersToManageLdapParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetAllowLocalRequestsFromSystemHooksParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetAllowLocalRequestsFromWebHooksAndServicesParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetAllowProjectCreationForGuestAndBelowParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetAllowRunnerRegistrationTokenParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetArchiveBuildsInHumanReadableParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetAsciidocMaxIncludesParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetAssetProxyAllowlistParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetAssetProxyEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetAssetProxySecretKeyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetAssetProxyUrlParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetAuthorizedKeysEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetAutoBanUserOnExcessiveProjectsDownloadParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetAutoDevopsDomainParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetAutoDevopsEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetAutomaticPurchasedStorageAllocationParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetBulkImportConcurrentPipelineBatchLimitParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetBulkImportEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetBulkImportMaxDownloadFileSizeParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetCanCreateGroupParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetCheckNamespacePlanParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetCiMaxIncludesParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetCiMaxTotalYamlSizeBytesParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetCommitEmailHostnameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetConcurrentBitbucketImportJobsLimitParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetConcurrentBitbucketServerImportJobsLimitParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetConcurrentGithubImportJobsLimitParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetContainerExpirationPoliciesEnableHistoricEntriesParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetContainerRegistryCleanupTagsServiceMaxListSizeParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetContainerRegistryDeleteTagsServiceTimeoutParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetContainerRegistryExpirationPoliciesCachingParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetContainerRegistryExpirationPoliciesWorkerCapacityParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetContainerRegistryTokenExpireDelayParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDeactivateDormantUsersParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDeactivateDormantUsersPeriodParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDecompressArchiveFileTimeoutParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDefaultArtifactsExpireInParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDefaultBranchNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDefaultBranchProtectionParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDefaultCiConfigPathParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDefaultGroupVisibilityParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDefaultPreferredLanguageParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDefaultProjectCreationParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDefaultProjectsLimitParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDefaultProjectVisibilityParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDefaultSnippetVisibilityParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDefaultSyntaxHighlightingThemeParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDeleteInactiveProjectsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDeleteUnconfirmedUsersParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDeletionAdjournedPeriodParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDiagramsnetEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDiagramsnetUrlParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDiffMaxFilesParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDiffMaxLinesParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDiffMaxPatchBytesParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDisableAdminOauthScopesParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDisabledOauthSignInSourcesParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDisableFeedTokenParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDisablePersonalAccessTokensParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDnsRebindingProtectionEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDomainAllowlistParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDomainDenylistParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDomainDenylistEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDownstreamPipelineTriggerLimitPerProjectUserShaParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDsaKeyRestrictionParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetDuoFeaturesEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetEcdsaKeyRestrictionParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetEcdsaSkKeyRestrictionParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetEd25519KeyRestrictionParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetEd25519SkKeyRestrictionParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetEksAccessKeyIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetEksAccountIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetEksIntegrationEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetEksSecretAccessKeyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetElasticsearchAwsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetElasticsearchAwsAccessKeyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetElasticsearchAwsRegionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetElasticsearchAwsSecretAccessKeyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetElasticsearchIndexedFieldLengthLimitParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetElasticsearchIndexedFileSizeLimitKbParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetElasticsearchIndexingParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetElasticsearchLimitIndexingParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetElasticsearchMaxBulkConcurrencyParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetElasticsearchMaxBulkSizeMbParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetElasticsearchNamespaceIdsParameters(val *[]*float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetElasticsearchPasswordParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetElasticsearchProjectIdsParameters(val *[]*float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetElasticsearchSearchParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetElasticsearchUrlParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetElasticsearchUsernameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetEmailAdditionalTextParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetEmailAuthorInBodyParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetEnabledGitAccessProtocolParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetEnforceNamespaceStorageLimitParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetEnforceTermsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetExternalAuthClientCertParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetExternalAuthClientKeyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetExternalAuthClientKeyPassParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetExternalAuthorizationServiceDefaultLabelParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetExternalAuthorizationServiceEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetExternalAuthorizationServiceTimeoutParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetExternalAuthorizationServiceUrlParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetExternalPipelineValidationServiceTimeoutParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetExternalPipelineValidationServiceTokenParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetExternalPipelineValidationServiceUrlParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetFileTemplateProjectIdParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetFirstDayOfWeekParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetGeoNodeAllowedIpsParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetGeoStatusTimeoutParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetGitalyTimeoutDefaultParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetGitalyTimeoutFastParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetGitalyTimeoutMediumParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetGitRateLimitUsersAllowlistParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetGitTwoFactorSessionExpiryParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetGrafanaEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetGrafanaUrlParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetGravatarEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetGroupOwnersCanManageDefaultBranchProtectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetHashedStorageEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetHelpPageHideCommercialContentParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetHelpPageSupportUrlParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetHelpPageTextParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetHelpTextParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetHideThirdPartyOffersParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetHomePageUrlParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetHousekeepingEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetHousekeepingFullRepackPeriodParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetHousekeepingGcPeriodParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetHousekeepingIncrementalRepackPeriodParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetHousekeepingOptimizeRepositoryPeriodParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetHtmlEmailsEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetImportSourcesParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetInactiveProjectsDeleteAfterMonthsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetInactiveProjectsMinSizeMbParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetInactiveProjectsSendWarningEmailAfterMonthsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetInProductMarketingEmailsEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetInvisibleCaptchaEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetIssuesCreateLimitParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetKeepLatestArtifactParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetLocalMarkdownVersionParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetMailgunEventsEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetMailgunSigningKeyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetMaintenanceModeParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetMaintenanceModeMessageParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetMaxArtifactsSizeParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetMaxAttachmentSizeParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetMaxExportSizeParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetMaxImportSizeParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetMaxNumberOfRepositoryDownloadsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetMaxNumberOfRepositoryDownloadsWithinTimePeriodParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetMaxPagesSizeParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetMaxPersonalAccessTokenLifetimeParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetMaxSshKeyLifetimeParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetMaxTerraformStateSizeBytesParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetMetricsMethodCallThresholdParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetMinimumPasswordLengthParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetMirrorAvailableParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetMirrorCapacityThresholdParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetMirrorMaxCapacityParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetMirrorMaxDelayParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetNpmPackageRequestsForwardingParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetOutboundLocalRequestsWhitelistParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetPackageRegistryCleanupPoliciesWorkerCapacityParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetPagesDomainVerificationEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetPasswordAuthenticationEnabledForGitParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetPasswordAuthenticationEnabledForWebParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetPasswordLowercaseRequiredParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetPasswordNumberRequiredParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetPasswordSymbolRequiredParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetPasswordUppercaseRequiredParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetPerformanceBarAllowedGroupPathParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetPersonalAccessTokenPrefixParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetPipelineLimitPerProjectUserShaParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetPlantumlEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetPlantumlUrlParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetPollingIntervalMultiplierParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetProjectExportEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetPrometheusMetricsEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetProtectedCiVariablesParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetPushEventActivitiesLimitParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetPushEventHooksLimitParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetPypiPackageRequestsForwardingParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetRateLimitingResponseTextParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetRawBlobRequestLimitParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetRecaptchaEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetRecaptchaPrivateKeyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetRecaptchaSiteKeyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetReceiveMaxInputSizeParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetRepositoryChecksEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetRepositorySizeLimitParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetRepositoryStoragesParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetRepositoryStoragesWeightedParameters(val *map[string]*float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetRequireAdminApprovalAfterUserSignupParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetRequireTwoFactorAuthenticationParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetRestrictedVisibilityLevelsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetRsaKeyRestrictionParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSearchRateLimitParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSearchRateLimitUnauthenticatedParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSendUserConfirmationEmailParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSessionExpireDelayParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSharedRunnersEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSharedRunnersMinutesParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSharedRunnersTextParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSidekiqJobLimiterCompressionThresholdBytesParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSidekiqJobLimiterLimitBytesParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSidekiqJobLimiterModeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSignInTextParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSignupEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSlackAppEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSlackAppIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSlackAppSecretParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSlackAppSigningSecretParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSlackAppVerificationTokenParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSnippetSizeLimitParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSnowplowAppIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSnowplowCollectorHostnameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSnowplowCookieDomainParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSnowplowEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSourcegraphEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSourcegraphPublicOnlyParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSourcegraphUrlParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSpamCheckApiKeyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSpamCheckEndpointEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSpamCheckEndpointUrlParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetSuggestPipelineEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetTerminalMaxSessionTimeParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetTermsParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetThrottleAuthenticatedApiEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetThrottleAuthenticatedApiPeriodInSecondsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetThrottleAuthenticatedApiRequestsPerPeriodParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetThrottleAuthenticatedPackagesApiEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetThrottleAuthenticatedPackagesApiPeriodInSecondsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetThrottleAuthenticatedPackagesApiRequestsPerPeriodParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetThrottleAuthenticatedWebEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetThrottleAuthenticatedWebPeriodInSecondsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetThrottleAuthenticatedWebRequestsPerPeriodParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetThrottleUnauthenticatedApiEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetThrottleUnauthenticatedApiPeriodInSecondsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetThrottleUnauthenticatedApiRequestsPerPeriodParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetThrottleUnauthenticatedPackagesApiEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetThrottleUnauthenticatedPackagesApiPeriodInSecondsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetThrottleUnauthenticatedPackagesApiRequestsPerPeriodParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetThrottleUnauthenticatedWebEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetThrottleUnauthenticatedWebPeriodInSecondsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetThrottleUnauthenticatedWebRequestsPerPeriodParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetTimeTrackingLimitToHoursParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetTwoFactorGracePeriodParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetUniqueIpsLimitEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetUniqueIpsLimitPerUserParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetUniqueIpsLimitTimeWindowParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetUsagePingEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetUserDeactivationEmailsEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetUserDefaultExternalParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetUserDefaultInternalRegexParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetUserOauthApplicationsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetUserShowAddSshKeyMessageParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetVersionCheckEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetWebIdeClientsidePreviewEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetWhatsNewVariantParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettings) validateSetWikiPageMaxContentBytesParameters(val *float64) error {
	return nil
}

func validateNewApplicationSettingsParameters(scope constructs.Construct, id *string, config *ApplicationSettingsConfig) error {
	return nil
}

