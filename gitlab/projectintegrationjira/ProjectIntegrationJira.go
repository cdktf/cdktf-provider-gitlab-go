// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectintegrationjira

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v15/projectintegrationjira/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/project_integration_jira gitlab_project_integration_jira}.
type ProjectIntegrationJira interface {
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

// The jsii proxy struct for ProjectIntegrationJira
type jsiiProxy_ProjectIntegrationJira struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ProjectIntegrationJira) Active() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"active",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) ApiUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) ApiUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) CommentOnEventEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commentOnEventEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) CommentOnEventEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commentOnEventEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) CommitEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commitEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) CommitEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commitEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) IssuesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issuesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) IssuesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issuesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) JiraAuthType() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jiraAuthType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) JiraAuthTypeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jiraAuthTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) JiraIssuePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraIssuePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) JiraIssuePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraIssuePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) JiraIssueRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraIssueRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) JiraIssueRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraIssueRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) JiraIssueTransitionAutomatic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jiraIssueTransitionAutomatic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) JiraIssueTransitionAutomaticInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jiraIssueTransitionAutomaticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) JiraIssueTransitionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraIssueTransitionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) JiraIssueTransitionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraIssueTransitionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) MergeRequestsEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeRequestsEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) MergeRequestsEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeRequestsEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) ProjectKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projectKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) ProjectKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projectKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) UseInheritedSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useInheritedSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) UseInheritedSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useInheritedSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectIntegrationJira) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/project_integration_jira gitlab_project_integration_jira} Resource.
func NewProjectIntegrationJira(scope constructs.Construct, id *string, config *ProjectIntegrationJiraConfig) ProjectIntegrationJira {
	_init_.Initialize()

	if err := validateNewProjectIntegrationJiraParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProjectIntegrationJira{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.projectIntegrationJira.ProjectIntegrationJira",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/project_integration_jira gitlab_project_integration_jira} Resource.
func NewProjectIntegrationJira_Override(p ProjectIntegrationJira, scope constructs.Construct, id *string, config *ProjectIntegrationJiraConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.projectIntegrationJira.ProjectIntegrationJira",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_ProjectIntegrationJira)SetApiUrl(val *string) {
	if err := j.validateSetApiUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiUrl",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationJira)SetCommentOnEventEnabled(val interface{}) {
	if err := j.validateSetCommentOnEventEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commentOnEventEnabled",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationJira)SetCommitEvents(val interface{}) {
	if err := j.validateSetCommitEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitEvents",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationJira)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationJira)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationJira)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationJira)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationJira)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationJira)SetIssuesEnabled(val interface{}) {
	if err := j.validateSetIssuesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuesEnabled",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationJira)SetJiraAuthType(val *float64) {
	if err := j.validateSetJiraAuthTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jiraAuthType",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationJira)SetJiraIssuePrefix(val *string) {
	if err := j.validateSetJiraIssuePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jiraIssuePrefix",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationJira)SetJiraIssueRegex(val *string) {
	if err := j.validateSetJiraIssueRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jiraIssueRegex",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationJira)SetJiraIssueTransitionAutomatic(val interface{}) {
	if err := j.validateSetJiraIssueTransitionAutomaticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jiraIssueTransitionAutomatic",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationJira)SetJiraIssueTransitionId(val *string) {
	if err := j.validateSetJiraIssueTransitionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jiraIssueTransitionId",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationJira)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationJira)SetMergeRequestsEvents(val interface{}) {
	if err := j.validateSetMergeRequestsEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeRequestsEvents",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationJira)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationJira)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationJira)SetProjectKeys(val *[]*string) {
	if err := j.validateSetProjectKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectKeys",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationJira)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationJira)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationJira)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationJira)SetUseInheritedSettings(val interface{}) {
	if err := j.validateSetUseInheritedSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useInheritedSettings",
		val,
	)
}

func (j *jsiiProxy_ProjectIntegrationJira)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

// Generates CDKTF code for importing a ProjectIntegrationJira resource upon running "cdktf plan <stack-name>".
func ProjectIntegrationJira_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateProjectIntegrationJira_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.projectIntegrationJira.ProjectIntegrationJira",
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
func ProjectIntegrationJira_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjectIntegrationJira_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.projectIntegrationJira.ProjectIntegrationJira",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ProjectIntegrationJira_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjectIntegrationJira_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.projectIntegrationJira.ProjectIntegrationJira",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ProjectIntegrationJira_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjectIntegrationJira_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.projectIntegrationJira.ProjectIntegrationJira",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ProjectIntegrationJira_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.projectIntegrationJira.ProjectIntegrationJira",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_ProjectIntegrationJira) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_ProjectIntegrationJira) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_ProjectIntegrationJira) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationJira) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationJira) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationJira) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationJira) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationJira) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationJira) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationJira) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationJira) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationJira) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationJira) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_ProjectIntegrationJira) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationJira) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_ProjectIntegrationJira) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_ProjectIntegrationJira) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_ProjectIntegrationJira) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_ProjectIntegrationJira) ResetApiUrl() {
	_jsii_.InvokeVoid(
		p,
		"resetApiUrl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationJira) ResetCommentOnEventEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetCommentOnEventEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationJira) ResetCommitEvents() {
	_jsii_.InvokeVoid(
		p,
		"resetCommitEvents",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationJira) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationJira) ResetIssuesEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetIssuesEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationJira) ResetJiraAuthType() {
	_jsii_.InvokeVoid(
		p,
		"resetJiraAuthType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationJira) ResetJiraIssuePrefix() {
	_jsii_.InvokeVoid(
		p,
		"resetJiraIssuePrefix",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationJira) ResetJiraIssueRegex() {
	_jsii_.InvokeVoid(
		p,
		"resetJiraIssueRegex",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationJira) ResetJiraIssueTransitionAutomatic() {
	_jsii_.InvokeVoid(
		p,
		"resetJiraIssueTransitionAutomatic",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationJira) ResetJiraIssueTransitionId() {
	_jsii_.InvokeVoid(
		p,
		"resetJiraIssueTransitionId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationJira) ResetMergeRequestsEvents() {
	_jsii_.InvokeVoid(
		p,
		"resetMergeRequestsEvents",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationJira) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationJira) ResetProjectKeys() {
	_jsii_.InvokeVoid(
		p,
		"resetProjectKeys",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationJira) ResetUseInheritedSettings() {
	_jsii_.InvokeVoid(
		p,
		"resetUseInheritedSettings",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationJira) ResetUsername() {
	_jsii_.InvokeVoid(
		p,
		"resetUsername",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectIntegrationJira) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationJira) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationJira) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationJira) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationJira) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectIntegrationJira) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

