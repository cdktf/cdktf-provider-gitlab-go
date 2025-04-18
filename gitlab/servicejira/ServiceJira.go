// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicejira

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v13/servicejira/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.11.0/docs/resources/service_jira gitlab_service_jira}.
type ServiceJira interface {
	cdktf.TerraformResource
	Active() cdktf.IResolvable
	ApiUrl() *string
	SetApiUrl(val *string)
	ApiUrlInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CommentOnEventEnabled() interface{}
	SetCommentOnEventEnabled(val interface{})
	CommentOnEventEnabledInput() interface{}
	CommitEvents() interface{}
	SetCommitEvents(val interface{})
	CommitEventsInput() interface{}
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
	CreatedAt() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IssuesEnabled() interface{}
	SetIssuesEnabled(val interface{})
	IssuesEnabledInput() interface{}
	JiraAuthType() *float64
	SetJiraAuthType(val *float64)
	JiraAuthTypeInput() *float64
	JiraIssuePrefix() *string
	SetJiraIssuePrefix(val *string)
	JiraIssuePrefixInput() *string
	JiraIssueRegex() *string
	SetJiraIssueRegex(val *string)
	JiraIssueRegexInput() *string
	JiraIssueTransitionAutomatic() interface{}
	SetJiraIssueTransitionAutomatic(val interface{})
	JiraIssueTransitionAutomaticInput() interface{}
	JiraIssueTransitionId() *string
	SetJiraIssueTransitionId(val *string)
	JiraIssueTransitionIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MergeRequestsEvents() interface{}
	SetMergeRequestsEvents(val interface{})
	MergeRequestsEventsInput() interface{}
	// The tree node.
	Node() constructs.Node
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	ProjectKey() *string
	SetProjectKey(val *string)
	ProjectKeyInput() *string
	ProjectKeys() *[]*string
	SetProjectKeys(val *[]*string)
	ProjectKeysInput() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Title() *string
	UpdatedAt() *string
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
	UseInheritedSettings() interface{}
	SetUseInheritedSettings(val interface{})
	UseInheritedSettingsInput() interface{}
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	ResetApiUrl()
	ResetCommentOnEventEnabled()
	ResetCommitEvents()
	ResetId()
	ResetIssuesEnabled()
	ResetJiraAuthType()
	ResetJiraIssuePrefix()
	ResetJiraIssueRegex()
	ResetJiraIssueTransitionAutomatic()
	ResetJiraIssueTransitionId()
	ResetMergeRequestsEvents()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProjectKey()
	ResetProjectKeys()
	ResetUseInheritedSettings()
	ResetUsername()
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

// The jsii proxy struct for ServiceJira
type jsiiProxy_ServiceJira struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServiceJira) Active() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"active",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) ApiUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) ApiUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) CommentOnEventEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commentOnEventEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) CommentOnEventEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commentOnEventEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) CommitEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commitEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) CommitEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commitEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) IssuesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issuesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) IssuesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issuesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) JiraAuthType() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jiraAuthType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) JiraAuthTypeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jiraAuthTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) JiraIssuePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraIssuePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) JiraIssuePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraIssuePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) JiraIssueRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraIssueRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) JiraIssueRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraIssueRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) JiraIssueTransitionAutomatic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jiraIssueTransitionAutomatic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) JiraIssueTransitionAutomaticInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jiraIssueTransitionAutomaticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) JiraIssueTransitionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraIssueTransitionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) JiraIssueTransitionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraIssueTransitionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) MergeRequestsEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeRequestsEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) MergeRequestsEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeRequestsEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) ProjectKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) ProjectKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) ProjectKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projectKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) ProjectKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projectKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) UseInheritedSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useInheritedSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) UseInheritedSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useInheritedSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceJira) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.11.0/docs/resources/service_jira gitlab_service_jira} Resource.
func NewServiceJira(scope constructs.Construct, id *string, config *ServiceJiraConfig) ServiceJira {
	_init_.Initialize()

	if err := validateNewServiceJiraParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceJira{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.serviceJira.ServiceJira",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.11.0/docs/resources/service_jira gitlab_service_jira} Resource.
func NewServiceJira_Override(s ServiceJira, scope constructs.Construct, id *string, config *ServiceJiraConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.serviceJira.ServiceJira",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServiceJira)SetApiUrl(val *string) {
	if err := j.validateSetApiUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiUrl",
		val,
	)
}

func (j *jsiiProxy_ServiceJira)SetCommentOnEventEnabled(val interface{}) {
	if err := j.validateSetCommentOnEventEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commentOnEventEnabled",
		val,
	)
}

func (j *jsiiProxy_ServiceJira)SetCommitEvents(val interface{}) {
	if err := j.validateSetCommitEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitEvents",
		val,
	)
}

func (j *jsiiProxy_ServiceJira)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ServiceJira)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServiceJira)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServiceJira)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ServiceJira)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ServiceJira)SetIssuesEnabled(val interface{}) {
	if err := j.validateSetIssuesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuesEnabled",
		val,
	)
}

func (j *jsiiProxy_ServiceJira)SetJiraAuthType(val *float64) {
	if err := j.validateSetJiraAuthTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jiraAuthType",
		val,
	)
}

func (j *jsiiProxy_ServiceJira)SetJiraIssuePrefix(val *string) {
	if err := j.validateSetJiraIssuePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jiraIssuePrefix",
		val,
	)
}

func (j *jsiiProxy_ServiceJira)SetJiraIssueRegex(val *string) {
	if err := j.validateSetJiraIssueRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jiraIssueRegex",
		val,
	)
}

func (j *jsiiProxy_ServiceJira)SetJiraIssueTransitionAutomatic(val interface{}) {
	if err := j.validateSetJiraIssueTransitionAutomaticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jiraIssueTransitionAutomatic",
		val,
	)
}

func (j *jsiiProxy_ServiceJira)SetJiraIssueTransitionId(val *string) {
	if err := j.validateSetJiraIssueTransitionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jiraIssueTransitionId",
		val,
	)
}

func (j *jsiiProxy_ServiceJira)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServiceJira)SetMergeRequestsEvents(val interface{}) {
	if err := j.validateSetMergeRequestsEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeRequestsEvents",
		val,
	)
}

func (j *jsiiProxy_ServiceJira)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_ServiceJira)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ServiceJira)SetProjectKey(val *string) {
	if err := j.validateSetProjectKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectKey",
		val,
	)
}

func (j *jsiiProxy_ServiceJira)SetProjectKeys(val *[]*string) {
	if err := j.validateSetProjectKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectKeys",
		val,
	)
}

func (j *jsiiProxy_ServiceJira)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServiceJira)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ServiceJira)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_ServiceJira)SetUseInheritedSettings(val interface{}) {
	if err := j.validateSetUseInheritedSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useInheritedSettings",
		val,
	)
}

func (j *jsiiProxy_ServiceJira)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

// Generates CDKTF code for importing a ServiceJira resource upon running "cdktf plan <stack-name>".
func ServiceJira_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateServiceJira_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.serviceJira.ServiceJira",
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
func ServiceJira_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServiceJira_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.serviceJira.ServiceJira",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServiceJira_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServiceJira_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.serviceJira.ServiceJira",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServiceJira_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServiceJira_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.serviceJira.ServiceJira",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServiceJira_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.serviceJira.ServiceJira",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_ServiceJira) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_ServiceJira) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_ServiceJira) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceJira) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceJira) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceJira) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceJira) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceJira) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceJira) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceJira) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceJira) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceJira) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceJira) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_ServiceJira) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceJira) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_ServiceJira) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_ServiceJira) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_ServiceJira) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServiceJira) ResetApiUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetApiUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceJira) ResetCommentOnEventEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetCommentOnEventEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceJira) ResetCommitEvents() {
	_jsii_.InvokeVoid(
		s,
		"resetCommitEvents",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceJira) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceJira) ResetIssuesEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetIssuesEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceJira) ResetJiraAuthType() {
	_jsii_.InvokeVoid(
		s,
		"resetJiraAuthType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceJira) ResetJiraIssuePrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetJiraIssuePrefix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceJira) ResetJiraIssueRegex() {
	_jsii_.InvokeVoid(
		s,
		"resetJiraIssueRegex",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceJira) ResetJiraIssueTransitionAutomatic() {
	_jsii_.InvokeVoid(
		s,
		"resetJiraIssueTransitionAutomatic",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceJira) ResetJiraIssueTransitionId() {
	_jsii_.InvokeVoid(
		s,
		"resetJiraIssueTransitionId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceJira) ResetMergeRequestsEvents() {
	_jsii_.InvokeVoid(
		s,
		"resetMergeRequestsEvents",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceJira) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceJira) ResetProjectKey() {
	_jsii_.InvokeVoid(
		s,
		"resetProjectKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceJira) ResetProjectKeys() {
	_jsii_.InvokeVoid(
		s,
		"resetProjectKeys",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceJira) ResetUseInheritedSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetUseInheritedSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceJira) ResetUsername() {
	_jsii_.InvokeVoid(
		s,
		"resetUsername",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceJira) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceJira) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceJira) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceJira) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceJira) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceJira) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

