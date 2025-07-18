// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package group

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v15/group/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/group gitlab_group}.
type Group interface {
	cdktf.TerraformResource
	AllowedEmailDomainsList() *[]*string
	SetAllowedEmailDomainsList(val *[]*string)
	AllowedEmailDomainsListInput() *[]*string
	AutoDevopsEnabled() interface{}
	SetAutoDevopsEnabled(val interface{})
	AutoDevopsEnabledInput() interface{}
	Avatar() *string
	SetAvatar(val *string)
	AvatarHash() *string
	SetAvatarHash(val *string)
	AvatarHashInput() *string
	AvatarInput() *string
	AvatarUrl() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DefaultBranch() *string
	SetDefaultBranch(val *string)
	DefaultBranchInput() *string
	DefaultBranchProtection() *float64
	SetDefaultBranchProtection(val *float64)
	DefaultBranchProtectionDefaults() GroupDefaultBranchProtectionDefaultsOutputReference
	DefaultBranchProtectionDefaultsInput() *GroupDefaultBranchProtectionDefaults
	DefaultBranchProtectionInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EmailsEnabled() interface{}
	SetEmailsEnabled(val interface{})
	EmailsEnabledInput() interface{}
	ExtraSharedRunnersMinutesLimit() *float64
	SetExtraSharedRunnersMinutesLimit(val *float64)
	ExtraSharedRunnersMinutesLimitInput() *float64
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FullName() *string
	FullPath() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IpRestrictionRanges() *[]*string
	SetIpRestrictionRanges(val *[]*string)
	IpRestrictionRangesInput() *[]*string
	LfsEnabled() interface{}
	SetLfsEnabled(val interface{})
	LfsEnabledInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MembershipLock() interface{}
	SetMembershipLock(val interface{})
	MembershipLockInput() interface{}
	MentionsDisabled() interface{}
	SetMentionsDisabled(val interface{})
	MentionsDisabledInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ParentId() *float64
	SetParentId(val *float64)
	ParentIdInput() *float64
	Path() *string
	SetPath(val *string)
	PathInput() *string
	PermanentlyRemoveOnDelete() interface{}
	SetPermanentlyRemoveOnDelete(val interface{})
	PermanentlyRemoveOnDeleteInput() interface{}
	PreventForkingOutsideGroup() interface{}
	SetPreventForkingOutsideGroup(val interface{})
	PreventForkingOutsideGroupInput() interface{}
	ProjectCreationLevel() *string
	SetProjectCreationLevel(val *string)
	ProjectCreationLevelInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PushRules() GroupPushRulesOutputReference
	PushRulesInput() *GroupPushRules
	// Experimental.
	RawOverrides() interface{}
	RequestAccessEnabled() interface{}
	SetRequestAccessEnabled(val interface{})
	RequestAccessEnabledInput() interface{}
	RequireTwoFactorAuthentication() interface{}
	SetRequireTwoFactorAuthentication(val interface{})
	RequireTwoFactorAuthenticationInput() interface{}
	RunnersToken() *string
	SharedRunnersMinutesLimit() *float64
	SetSharedRunnersMinutesLimit(val *float64)
	SharedRunnersMinutesLimitInput() *float64
	SharedRunnersSetting() *string
	SetSharedRunnersSetting(val *string)
	SharedRunnersSettingInput() *string
	ShareWithGroupLock() interface{}
	SetShareWithGroupLock(val interface{})
	ShareWithGroupLockInput() interface{}
	SubgroupCreationLevel() *string
	SetSubgroupCreationLevel(val *string)
	SubgroupCreationLevelInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TwoFactorGracePeriod() *float64
	SetTwoFactorGracePeriod(val *float64)
	TwoFactorGracePeriodInput() *float64
	VisibilityLevel() *string
	SetVisibilityLevel(val *string)
	VisibilityLevelInput() *string
	WebUrl() *string
	WikiAccessLevel() *string
	SetWikiAccessLevel(val *string)
	WikiAccessLevelInput() *string
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
	PutDefaultBranchProtectionDefaults(value *GroupDefaultBranchProtectionDefaults)
	PutPushRules(value *GroupPushRules)
	ResetAllowedEmailDomainsList()
	ResetAutoDevopsEnabled()
	ResetAvatar()
	ResetAvatarHash()
	ResetDefaultBranch()
	ResetDefaultBranchProtection()
	ResetDefaultBranchProtectionDefaults()
	ResetDescription()
	ResetEmailsEnabled()
	ResetExtraSharedRunnersMinutesLimit()
	ResetId()
	ResetIpRestrictionRanges()
	ResetLfsEnabled()
	ResetMembershipLock()
	ResetMentionsDisabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParentId()
	ResetPermanentlyRemoveOnDelete()
	ResetPreventForkingOutsideGroup()
	ResetProjectCreationLevel()
	ResetPushRules()
	ResetRequestAccessEnabled()
	ResetRequireTwoFactorAuthentication()
	ResetSharedRunnersMinutesLimit()
	ResetSharedRunnersSetting()
	ResetShareWithGroupLock()
	ResetSubgroupCreationLevel()
	ResetTwoFactorGracePeriod()
	ResetVisibilityLevel()
	ResetWikiAccessLevel()
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

// The jsii proxy struct for Group
type jsiiProxy_Group struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Group) AllowedEmailDomainsList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedEmailDomainsList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) AllowedEmailDomainsListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedEmailDomainsListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) AutoDevopsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDevopsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) AutoDevopsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDevopsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Avatar() *string {
	var returns *string
	_jsii_.Get(
		j,
		"avatar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) AvatarHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"avatarHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) AvatarHashInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"avatarHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) AvatarInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"avatarInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) AvatarUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"avatarUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) DefaultBranch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) DefaultBranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBranchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) DefaultBranchProtection() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultBranchProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) DefaultBranchProtectionDefaults() GroupDefaultBranchProtectionDefaultsOutputReference {
	var returns GroupDefaultBranchProtectionDefaultsOutputReference
	_jsii_.Get(
		j,
		"defaultBranchProtectionDefaults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) DefaultBranchProtectionDefaultsInput() *GroupDefaultBranchProtectionDefaults {
	var returns *GroupDefaultBranchProtectionDefaults
	_jsii_.Get(
		j,
		"defaultBranchProtectionDefaultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) DefaultBranchProtectionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultBranchProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) EmailsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) EmailsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) ExtraSharedRunnersMinutesLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"extraSharedRunnersMinutesLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) ExtraSharedRunnersMinutesLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"extraSharedRunnersMinutesLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) FullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) FullPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) IpRestrictionRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipRestrictionRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) IpRestrictionRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipRestrictionRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) LfsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lfsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) LfsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lfsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) MembershipLock() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"membershipLock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) MembershipLockInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"membershipLockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) MentionsDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mentionsDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) MentionsDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mentionsDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) ParentId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) ParentIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) PermanentlyRemoveOnDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permanentlyRemoveOnDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) PermanentlyRemoveOnDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permanentlyRemoveOnDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) PreventForkingOutsideGroup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventForkingOutsideGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) PreventForkingOutsideGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventForkingOutsideGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) ProjectCreationLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectCreationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) ProjectCreationLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectCreationLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) PushRules() GroupPushRulesOutputReference {
	var returns GroupPushRulesOutputReference
	_jsii_.Get(
		j,
		"pushRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) PushRulesInput() *GroupPushRules {
	var returns *GroupPushRules
	_jsii_.Get(
		j,
		"pushRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) RequestAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) RequestAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) RequireTwoFactorAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireTwoFactorAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) RequireTwoFactorAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireTwoFactorAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) RunnersToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runnersToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) SharedRunnersMinutesLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sharedRunnersMinutesLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) SharedRunnersMinutesLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sharedRunnersMinutesLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) SharedRunnersSetting() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedRunnersSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) SharedRunnersSettingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedRunnersSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) ShareWithGroupLock() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shareWithGroupLock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) ShareWithGroupLockInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shareWithGroupLockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) SubgroupCreationLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subgroupCreationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) SubgroupCreationLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subgroupCreationLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) TwoFactorGracePeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"twoFactorGracePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) TwoFactorGracePeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"twoFactorGracePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) VisibilityLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibilityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) VisibilityLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibilityLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) WebUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) WikiAccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wikiAccessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) WikiAccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wikiAccessLevelInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/group gitlab_group} Resource.
func NewGroup(scope constructs.Construct, id *string, config *GroupConfig) Group {
	_init_.Initialize()

	if err := validateNewGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Group{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.group.Group",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/group gitlab_group} Resource.
func NewGroup_Override(g Group, scope constructs.Construct, id *string, config *GroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.group.Group",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_Group)SetAllowedEmailDomainsList(val *[]*string) {
	if err := j.validateSetAllowedEmailDomainsListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedEmailDomainsList",
		val,
	)
}

func (j *jsiiProxy_Group)SetAutoDevopsEnabled(val interface{}) {
	if err := j.validateSetAutoDevopsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoDevopsEnabled",
		val,
	)
}

func (j *jsiiProxy_Group)SetAvatar(val *string) {
	if err := j.validateSetAvatarParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"avatar",
		val,
	)
}

func (j *jsiiProxy_Group)SetAvatarHash(val *string) {
	if err := j.validateSetAvatarHashParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"avatarHash",
		val,
	)
}

func (j *jsiiProxy_Group)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Group)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Group)SetDefaultBranch(val *string) {
	if err := j.validateSetDefaultBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultBranch",
		val,
	)
}

func (j *jsiiProxy_Group)SetDefaultBranchProtection(val *float64) {
	if err := j.validateSetDefaultBranchProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultBranchProtection",
		val,
	)
}

func (j *jsiiProxy_Group)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Group)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Group)SetEmailsEnabled(val interface{}) {
	if err := j.validateSetEmailsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailsEnabled",
		val,
	)
}

func (j *jsiiProxy_Group)SetExtraSharedRunnersMinutesLimit(val *float64) {
	if err := j.validateSetExtraSharedRunnersMinutesLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extraSharedRunnersMinutesLimit",
		val,
	)
}

func (j *jsiiProxy_Group)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Group)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Group)SetIpRestrictionRanges(val *[]*string) {
	if err := j.validateSetIpRestrictionRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipRestrictionRanges",
		val,
	)
}

func (j *jsiiProxy_Group)SetLfsEnabled(val interface{}) {
	if err := j.validateSetLfsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lfsEnabled",
		val,
	)
}

func (j *jsiiProxy_Group)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Group)SetMembershipLock(val interface{}) {
	if err := j.validateSetMembershipLockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"membershipLock",
		val,
	)
}

func (j *jsiiProxy_Group)SetMentionsDisabled(val interface{}) {
	if err := j.validateSetMentionsDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mentionsDisabled",
		val,
	)
}

func (j *jsiiProxy_Group)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Group)SetParentId(val *float64) {
	if err := j.validateSetParentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentId",
		val,
	)
}

func (j *jsiiProxy_Group)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_Group)SetPermanentlyRemoveOnDelete(val interface{}) {
	if err := j.validateSetPermanentlyRemoveOnDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permanentlyRemoveOnDelete",
		val,
	)
}

func (j *jsiiProxy_Group)SetPreventForkingOutsideGroup(val interface{}) {
	if err := j.validateSetPreventForkingOutsideGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preventForkingOutsideGroup",
		val,
	)
}

func (j *jsiiProxy_Group)SetProjectCreationLevel(val *string) {
	if err := j.validateSetProjectCreationLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectCreationLevel",
		val,
	)
}

func (j *jsiiProxy_Group)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Group)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Group)SetRequestAccessEnabled(val interface{}) {
	if err := j.validateSetRequestAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_Group)SetRequireTwoFactorAuthentication(val interface{}) {
	if err := j.validateSetRequireTwoFactorAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireTwoFactorAuthentication",
		val,
	)
}

func (j *jsiiProxy_Group)SetSharedRunnersMinutesLimit(val *float64) {
	if err := j.validateSetSharedRunnersMinutesLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedRunnersMinutesLimit",
		val,
	)
}

func (j *jsiiProxy_Group)SetSharedRunnersSetting(val *string) {
	if err := j.validateSetSharedRunnersSettingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedRunnersSetting",
		val,
	)
}

func (j *jsiiProxy_Group)SetShareWithGroupLock(val interface{}) {
	if err := j.validateSetShareWithGroupLockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareWithGroupLock",
		val,
	)
}

func (j *jsiiProxy_Group)SetSubgroupCreationLevel(val *string) {
	if err := j.validateSetSubgroupCreationLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subgroupCreationLevel",
		val,
	)
}

func (j *jsiiProxy_Group)SetTwoFactorGracePeriod(val *float64) {
	if err := j.validateSetTwoFactorGracePeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoFactorGracePeriod",
		val,
	)
}

func (j *jsiiProxy_Group)SetVisibilityLevel(val *string) {
	if err := j.validateSetVisibilityLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visibilityLevel",
		val,
	)
}

func (j *jsiiProxy_Group)SetWikiAccessLevel(val *string) {
	if err := j.validateSetWikiAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wikiAccessLevel",
		val,
	)
}

// Generates CDKTF code for importing a Group resource upon running "cdktf plan <stack-name>".
func Group_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.group.Group",
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
func Group_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.group.Group",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Group_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.group.Group",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Group_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.group.Group",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Group_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.group.Group",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_Group) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_Group) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_Group) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_Group) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_Group) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_Group) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_Group) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_Group) PutDefaultBranchProtectionDefaults(value *GroupDefaultBranchProtectionDefaults) {
	if err := g.validatePutDefaultBranchProtectionDefaultsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDefaultBranchProtectionDefaults",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_Group) PutPushRules(value *GroupPushRules) {
	if err := g.validatePutPushRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPushRules",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_Group) ResetAllowedEmailDomainsList() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowedEmailDomainsList",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetAutoDevopsEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoDevopsEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetAvatar() {
	_jsii_.InvokeVoid(
		g,
		"resetAvatar",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetAvatarHash() {
	_jsii_.InvokeVoid(
		g,
		"resetAvatarHash",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetDefaultBranch() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultBranch",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetDefaultBranchProtection() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultBranchProtection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetDefaultBranchProtectionDefaults() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultBranchProtectionDefaults",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetEmailsEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetEmailsEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetExtraSharedRunnersMinutesLimit() {
	_jsii_.InvokeVoid(
		g,
		"resetExtraSharedRunnersMinutesLimit",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetIpRestrictionRanges() {
	_jsii_.InvokeVoid(
		g,
		"resetIpRestrictionRanges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetLfsEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetLfsEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetMembershipLock() {
	_jsii_.InvokeVoid(
		g,
		"resetMembershipLock",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetMentionsDisabled() {
	_jsii_.InvokeVoid(
		g,
		"resetMentionsDisabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetParentId() {
	_jsii_.InvokeVoid(
		g,
		"resetParentId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetPermanentlyRemoveOnDelete() {
	_jsii_.InvokeVoid(
		g,
		"resetPermanentlyRemoveOnDelete",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetPreventForkingOutsideGroup() {
	_jsii_.InvokeVoid(
		g,
		"resetPreventForkingOutsideGroup",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetProjectCreationLevel() {
	_jsii_.InvokeVoid(
		g,
		"resetProjectCreationLevel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetPushRules() {
	_jsii_.InvokeVoid(
		g,
		"resetPushRules",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetRequestAccessEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestAccessEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetRequireTwoFactorAuthentication() {
	_jsii_.InvokeVoid(
		g,
		"resetRequireTwoFactorAuthentication",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetSharedRunnersMinutesLimit() {
	_jsii_.InvokeVoid(
		g,
		"resetSharedRunnersMinutesLimit",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetSharedRunnersSetting() {
	_jsii_.InvokeVoid(
		g,
		"resetSharedRunnersSetting",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetShareWithGroupLock() {
	_jsii_.InvokeVoid(
		g,
		"resetShareWithGroupLock",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetSubgroupCreationLevel() {
	_jsii_.InvokeVoid(
		g,
		"resetSubgroupCreationLevel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetTwoFactorGracePeriod() {
	_jsii_.InvokeVoid(
		g,
		"resetTwoFactorGracePeriod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetVisibilityLevel() {
	_jsii_.InvokeVoid(
		g,
		"resetVisibilityLevel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetWikiAccessLevel() {
	_jsii_.InvokeVoid(
		g,
		"resetWikiAccessLevel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

