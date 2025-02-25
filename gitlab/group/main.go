// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package group

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.group.Group",
		reflect.TypeOf((*Group)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowedEmailDomainsList", GoGetter: "AllowedEmailDomainsList"},
			_jsii_.MemberProperty{JsiiProperty: "allowedEmailDomainsListInput", GoGetter: "AllowedEmailDomainsListInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoDevopsEnabled", GoGetter: "AutoDevopsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "autoDevopsEnabledInput", GoGetter: "AutoDevopsEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "avatar", GoGetter: "Avatar"},
			_jsii_.MemberProperty{JsiiProperty: "avatarHash", GoGetter: "AvatarHash"},
			_jsii_.MemberProperty{JsiiProperty: "avatarHashInput", GoGetter: "AvatarHashInput"},
			_jsii_.MemberProperty{JsiiProperty: "avatarInput", GoGetter: "AvatarInput"},
			_jsii_.MemberProperty{JsiiProperty: "avatarUrl", GoGetter: "AvatarUrl"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultBranch", GoGetter: "DefaultBranch"},
			_jsii_.MemberProperty{JsiiProperty: "defaultBranchInput", GoGetter: "DefaultBranchInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultBranchProtection", GoGetter: "DefaultBranchProtection"},
			_jsii_.MemberProperty{JsiiProperty: "defaultBranchProtectionDefaults", GoGetter: "DefaultBranchProtectionDefaults"},
			_jsii_.MemberProperty{JsiiProperty: "defaultBranchProtectionDefaultsInput", GoGetter: "DefaultBranchProtectionDefaultsInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultBranchProtectionInput", GoGetter: "DefaultBranchProtectionInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "emailsEnabled", GoGetter: "EmailsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "emailsEnabledInput", GoGetter: "EmailsEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "extraSharedRunnersMinutesLimit", GoGetter: "ExtraSharedRunnersMinutesLimit"},
			_jsii_.MemberProperty{JsiiProperty: "extraSharedRunnersMinutesLimitInput", GoGetter: "ExtraSharedRunnersMinutesLimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "fullName", GoGetter: "FullName"},
			_jsii_.MemberProperty{JsiiProperty: "fullPath", GoGetter: "FullPath"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipRestrictionRanges", GoGetter: "IpRestrictionRanges"},
			_jsii_.MemberProperty{JsiiProperty: "ipRestrictionRangesInput", GoGetter: "IpRestrictionRangesInput"},
			_jsii_.MemberProperty{JsiiProperty: "lfsEnabled", GoGetter: "LfsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "lfsEnabledInput", GoGetter: "LfsEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "membershipLock", GoGetter: "MembershipLock"},
			_jsii_.MemberProperty{JsiiProperty: "membershipLockInput", GoGetter: "MembershipLockInput"},
			_jsii_.MemberProperty{JsiiProperty: "mentionsDisabled", GoGetter: "MentionsDisabled"},
			_jsii_.MemberProperty{JsiiProperty: "mentionsDisabledInput", GoGetter: "MentionsDisabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "parentId", GoGetter: "ParentId"},
			_jsii_.MemberProperty{JsiiProperty: "parentIdInput", GoGetter: "ParentIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "pathInput", GoGetter: "PathInput"},
			_jsii_.MemberProperty{JsiiProperty: "permanentlyRemoveOnDelete", GoGetter: "PermanentlyRemoveOnDelete"},
			_jsii_.MemberProperty{JsiiProperty: "permanentlyRemoveOnDeleteInput", GoGetter: "PermanentlyRemoveOnDeleteInput"},
			_jsii_.MemberProperty{JsiiProperty: "preventForkingOutsideGroup", GoGetter: "PreventForkingOutsideGroup"},
			_jsii_.MemberProperty{JsiiProperty: "preventForkingOutsideGroupInput", GoGetter: "PreventForkingOutsideGroupInput"},
			_jsii_.MemberProperty{JsiiProperty: "projectCreationLevel", GoGetter: "ProjectCreationLevel"},
			_jsii_.MemberProperty{JsiiProperty: "projectCreationLevelInput", GoGetter: "ProjectCreationLevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "pushRules", GoGetter: "PushRules"},
			_jsii_.MemberProperty{JsiiProperty: "pushRulesInput", GoGetter: "PushRulesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultBranchProtectionDefaults", GoMethod: "PutDefaultBranchProtectionDefaults"},
			_jsii_.MemberMethod{JsiiMethod: "putPushRules", GoMethod: "PutPushRules"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "requestAccessEnabled", GoGetter: "RequestAccessEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "requestAccessEnabledInput", GoGetter: "RequestAccessEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "requireTwoFactorAuthentication", GoGetter: "RequireTwoFactorAuthentication"},
			_jsii_.MemberProperty{JsiiProperty: "requireTwoFactorAuthenticationInput", GoGetter: "RequireTwoFactorAuthenticationInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedEmailDomainsList", GoMethod: "ResetAllowedEmailDomainsList"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoDevopsEnabled", GoMethod: "ResetAutoDevopsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetAvatar", GoMethod: "ResetAvatar"},
			_jsii_.MemberMethod{JsiiMethod: "resetAvatarHash", GoMethod: "ResetAvatarHash"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultBranch", GoMethod: "ResetDefaultBranch"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultBranchProtection", GoMethod: "ResetDefaultBranchProtection"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultBranchProtectionDefaults", GoMethod: "ResetDefaultBranchProtectionDefaults"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailsEnabled", GoMethod: "ResetEmailsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetExtraSharedRunnersMinutesLimit", GoMethod: "ResetExtraSharedRunnersMinutesLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpRestrictionRanges", GoMethod: "ResetIpRestrictionRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resetLfsEnabled", GoMethod: "ResetLfsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetMembershipLock", GoMethod: "ResetMembershipLock"},
			_jsii_.MemberMethod{JsiiMethod: "resetMentionsDisabled", GoMethod: "ResetMentionsDisabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetParentId", GoMethod: "ResetParentId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPermanentlyRemoveOnDelete", GoMethod: "ResetPermanentlyRemoveOnDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreventForkingOutsideGroup", GoMethod: "ResetPreventForkingOutsideGroup"},
			_jsii_.MemberMethod{JsiiMethod: "resetProjectCreationLevel", GoMethod: "ResetProjectCreationLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetPushRules", GoMethod: "ResetPushRules"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestAccessEnabled", GoMethod: "ResetRequestAccessEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequireTwoFactorAuthentication", GoMethod: "ResetRequireTwoFactorAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "resetSharedRunnersMinutesLimit", GoMethod: "ResetSharedRunnersMinutesLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetSharedRunnersSetting", GoMethod: "ResetSharedRunnersSetting"},
			_jsii_.MemberMethod{JsiiMethod: "resetShareWithGroupLock", GoMethod: "ResetShareWithGroupLock"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubgroupCreationLevel", GoMethod: "ResetSubgroupCreationLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetTwoFactorGracePeriod", GoMethod: "ResetTwoFactorGracePeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resetVisibilityLevel", GoMethod: "ResetVisibilityLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetWikiAccessLevel", GoMethod: "ResetWikiAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "runnersToken", GoGetter: "RunnersToken"},
			_jsii_.MemberProperty{JsiiProperty: "sharedRunnersMinutesLimit", GoGetter: "SharedRunnersMinutesLimit"},
			_jsii_.MemberProperty{JsiiProperty: "sharedRunnersMinutesLimitInput", GoGetter: "SharedRunnersMinutesLimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "sharedRunnersSetting", GoGetter: "SharedRunnersSetting"},
			_jsii_.MemberProperty{JsiiProperty: "sharedRunnersSettingInput", GoGetter: "SharedRunnersSettingInput"},
			_jsii_.MemberProperty{JsiiProperty: "shareWithGroupLock", GoGetter: "ShareWithGroupLock"},
			_jsii_.MemberProperty{JsiiProperty: "shareWithGroupLockInput", GoGetter: "ShareWithGroupLockInput"},
			_jsii_.MemberProperty{JsiiProperty: "subgroupCreationLevel", GoGetter: "SubgroupCreationLevel"},
			_jsii_.MemberProperty{JsiiProperty: "subgroupCreationLevelInput", GoGetter: "SubgroupCreationLevelInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "twoFactorGracePeriod", GoGetter: "TwoFactorGracePeriod"},
			_jsii_.MemberProperty{JsiiProperty: "twoFactorGracePeriodInput", GoGetter: "TwoFactorGracePeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "visibilityLevel", GoGetter: "VisibilityLevel"},
			_jsii_.MemberProperty{JsiiProperty: "visibilityLevelInput", GoGetter: "VisibilityLevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "webUrl", GoGetter: "WebUrl"},
			_jsii_.MemberProperty{JsiiProperty: "wikiAccessLevel", GoGetter: "WikiAccessLevel"},
			_jsii_.MemberProperty{JsiiProperty: "wikiAccessLevelInput", GoGetter: "WikiAccessLevelInput"},
		},
		func() interface{} {
			j := jsiiProxy_Group{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.group.GroupConfig",
		reflect.TypeOf((*GroupConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.group.GroupDefaultBranchProtectionDefaults",
		reflect.TypeOf((*GroupDefaultBranchProtectionDefaults)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.group.GroupDefaultBranchProtectionDefaultsOutputReference",
		reflect.TypeOf((*GroupDefaultBranchProtectionDefaultsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowedToMerge", GoGetter: "AllowedToMerge"},
			_jsii_.MemberProperty{JsiiProperty: "allowedToMergeInput", GoGetter: "AllowedToMergeInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedToPush", GoGetter: "AllowedToPush"},
			_jsii_.MemberProperty{JsiiProperty: "allowedToPushInput", GoGetter: "AllowedToPushInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowForcePush", GoGetter: "AllowForcePush"},
			_jsii_.MemberProperty{JsiiProperty: "allowForcePushInput", GoGetter: "AllowForcePushInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "developerCanInitialPush", GoGetter: "DeveloperCanInitialPush"},
			_jsii_.MemberProperty{JsiiProperty: "developerCanInitialPushInput", GoGetter: "DeveloperCanInitialPushInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedToMerge", GoMethod: "ResetAllowedToMerge"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedToPush", GoMethod: "ResetAllowedToPush"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowForcePush", GoMethod: "ResetAllowForcePush"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeveloperCanInitialPush", GoMethod: "ResetDeveloperCanInitialPush"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GroupDefaultBranchProtectionDefaultsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.group.GroupPushRules",
		reflect.TypeOf((*GroupPushRules)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.group.GroupPushRulesOutputReference",
		reflect.TypeOf((*GroupPushRulesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authorEmailRegex", GoGetter: "AuthorEmailRegex"},
			_jsii_.MemberProperty{JsiiProperty: "authorEmailRegexInput", GoGetter: "AuthorEmailRegexInput"},
			_jsii_.MemberProperty{JsiiProperty: "branchNameRegex", GoGetter: "BranchNameRegex"},
			_jsii_.MemberProperty{JsiiProperty: "branchNameRegexInput", GoGetter: "BranchNameRegexInput"},
			_jsii_.MemberProperty{JsiiProperty: "commitCommitterCheck", GoGetter: "CommitCommitterCheck"},
			_jsii_.MemberProperty{JsiiProperty: "commitCommitterCheckInput", GoGetter: "CommitCommitterCheckInput"},
			_jsii_.MemberProperty{JsiiProperty: "commitCommitterNameCheck", GoGetter: "CommitCommitterNameCheck"},
			_jsii_.MemberProperty{JsiiProperty: "commitCommitterNameCheckInput", GoGetter: "CommitCommitterNameCheckInput"},
			_jsii_.MemberProperty{JsiiProperty: "commitMessageNegativeRegex", GoGetter: "CommitMessageNegativeRegex"},
			_jsii_.MemberProperty{JsiiProperty: "commitMessageNegativeRegexInput", GoGetter: "CommitMessageNegativeRegexInput"},
			_jsii_.MemberProperty{JsiiProperty: "commitMessageRegex", GoGetter: "CommitMessageRegex"},
			_jsii_.MemberProperty{JsiiProperty: "commitMessageRegexInput", GoGetter: "CommitMessageRegexInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "denyDeleteTag", GoGetter: "DenyDeleteTag"},
			_jsii_.MemberProperty{JsiiProperty: "denyDeleteTagInput", GoGetter: "DenyDeleteTagInput"},
			_jsii_.MemberProperty{JsiiProperty: "fileNameRegex", GoGetter: "FileNameRegex"},
			_jsii_.MemberProperty{JsiiProperty: "fileNameRegexInput", GoGetter: "FileNameRegexInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maxFileSize", GoGetter: "MaxFileSize"},
			_jsii_.MemberProperty{JsiiProperty: "maxFileSizeInput", GoGetter: "MaxFileSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "memberCheck", GoGetter: "MemberCheck"},
			_jsii_.MemberProperty{JsiiProperty: "memberCheckInput", GoGetter: "MemberCheckInput"},
			_jsii_.MemberProperty{JsiiProperty: "preventSecrets", GoGetter: "PreventSecrets"},
			_jsii_.MemberProperty{JsiiProperty: "preventSecretsInput", GoGetter: "PreventSecretsInput"},
			_jsii_.MemberProperty{JsiiProperty: "rejectNonDcoCommits", GoGetter: "RejectNonDcoCommits"},
			_jsii_.MemberProperty{JsiiProperty: "rejectNonDcoCommitsInput", GoGetter: "RejectNonDcoCommitsInput"},
			_jsii_.MemberProperty{JsiiProperty: "rejectUnsignedCommits", GoGetter: "RejectUnsignedCommits"},
			_jsii_.MemberProperty{JsiiProperty: "rejectUnsignedCommitsInput", GoGetter: "RejectUnsignedCommitsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthorEmailRegex", GoMethod: "ResetAuthorEmailRegex"},
			_jsii_.MemberMethod{JsiiMethod: "resetBranchNameRegex", GoMethod: "ResetBranchNameRegex"},
			_jsii_.MemberMethod{JsiiMethod: "resetCommitCommitterCheck", GoMethod: "ResetCommitCommitterCheck"},
			_jsii_.MemberMethod{JsiiMethod: "resetCommitCommitterNameCheck", GoMethod: "ResetCommitCommitterNameCheck"},
			_jsii_.MemberMethod{JsiiMethod: "resetCommitMessageNegativeRegex", GoMethod: "ResetCommitMessageNegativeRegex"},
			_jsii_.MemberMethod{JsiiMethod: "resetCommitMessageRegex", GoMethod: "ResetCommitMessageRegex"},
			_jsii_.MemberMethod{JsiiMethod: "resetDenyDeleteTag", GoMethod: "ResetDenyDeleteTag"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileNameRegex", GoMethod: "ResetFileNameRegex"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxFileSize", GoMethod: "ResetMaxFileSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetMemberCheck", GoMethod: "ResetMemberCheck"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreventSecrets", GoMethod: "ResetPreventSecrets"},
			_jsii_.MemberMethod{JsiiMethod: "resetRejectNonDcoCommits", GoMethod: "ResetRejectNonDcoCommits"},
			_jsii_.MemberMethod{JsiiMethod: "resetRejectUnsignedCommits", GoMethod: "ResetRejectUnsignedCommits"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GroupPushRulesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
