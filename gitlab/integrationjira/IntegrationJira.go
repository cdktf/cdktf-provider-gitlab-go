// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationjira

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v15/integrationjira/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/integration_jira gitlab_integration_jira}.
type IntegrationJira interface {
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

// The jsii proxy struct for IntegrationJira
type jsiiProxy_IntegrationJira struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IntegrationJira) Active() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"active",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) ApiUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) ApiUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) CommentOnEventEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commentOnEventEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) CommentOnEventEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commentOnEventEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) CommitEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commitEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) CommitEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commitEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) IssuesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issuesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) IssuesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issuesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) JiraAuthType() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jiraAuthType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) JiraAuthTypeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jiraAuthTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) JiraIssuePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraIssuePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) JiraIssuePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraIssuePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) JiraIssueRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraIssueRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) JiraIssueRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraIssueRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) JiraIssueTransitionAutomatic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jiraIssueTransitionAutomatic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) JiraIssueTransitionAutomaticInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jiraIssueTransitionAutomaticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) JiraIssueTransitionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraIssueTransitionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) JiraIssueTransitionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraIssueTransitionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) MergeRequestsEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeRequestsEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) MergeRequestsEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mergeRequestsEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) ProjectKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projectKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) ProjectKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projectKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) UseInheritedSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useInheritedSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) UseInheritedSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useInheritedSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationJira) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/integration_jira gitlab_integration_jira} Resource.
func NewIntegrationJira(scope constructs.Construct, id *string, config *IntegrationJiraConfig) IntegrationJira {
	_init_.Initialize()

	if err := validateNewIntegrationJiraParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationJira{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.integrationJira.IntegrationJira",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.1/docs/resources/integration_jira gitlab_integration_jira} Resource.
func NewIntegrationJira_Override(i IntegrationJira, scope constructs.Construct, id *string, config *IntegrationJiraConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.integrationJira.IntegrationJira",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IntegrationJira)SetApiUrl(val *string) {
	if err := j.validateSetApiUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiUrl",
		val,
	)
}

func (j *jsiiProxy_IntegrationJira)SetCommentOnEventEnabled(val interface{}) {
	if err := j.validateSetCommentOnEventEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commentOnEventEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationJira)SetCommitEvents(val interface{}) {
	if err := j.validateSetCommitEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitEvents",
		val,
	)
}

func (j *jsiiProxy_IntegrationJira)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IntegrationJira)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IntegrationJira)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IntegrationJira)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IntegrationJira)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IntegrationJira)SetIssuesEnabled(val interface{}) {
	if err := j.validateSetIssuesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuesEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationJira)SetJiraAuthType(val *float64) {
	if err := j.validateSetJiraAuthTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jiraAuthType",
		val,
	)
}

func (j *jsiiProxy_IntegrationJira)SetJiraIssuePrefix(val *string) {
	if err := j.validateSetJiraIssuePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jiraIssuePrefix",
		val,
	)
}

func (j *jsiiProxy_IntegrationJira)SetJiraIssueRegex(val *string) {
	if err := j.validateSetJiraIssueRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jiraIssueRegex",
		val,
	)
}

func (j *jsiiProxy_IntegrationJira)SetJiraIssueTransitionAutomatic(val interface{}) {
	if err := j.validateSetJiraIssueTransitionAutomaticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jiraIssueTransitionAutomatic",
		val,
	)
}

func (j *jsiiProxy_IntegrationJira)SetJiraIssueTransitionId(val *string) {
	if err := j.validateSetJiraIssueTransitionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jiraIssueTransitionId",
		val,
	)
}

func (j *jsiiProxy_IntegrationJira)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IntegrationJira)SetMergeRequestsEvents(val interface{}) {
	if err := j.validateSetMergeRequestsEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeRequestsEvents",
		val,
	)
}

func (j *jsiiProxy_IntegrationJira)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_IntegrationJira)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_IntegrationJira)SetProjectKeys(val *[]*string) {
	if err := j.validateSetProjectKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectKeys",
		val,
	)
}

func (j *jsiiProxy_IntegrationJira)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IntegrationJira)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IntegrationJira)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_IntegrationJira)SetUseInheritedSettings(val interface{}) {
	if err := j.validateSetUseInheritedSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useInheritedSettings",
		val,
	)
}

func (j *jsiiProxy_IntegrationJira)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

// Generates CDKTF code for importing a IntegrationJira resource upon running "cdktf plan <stack-name>".
func IntegrationJira_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIntegrationJira_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.integrationJira.IntegrationJira",
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
func IntegrationJira_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationJira_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.integrationJira.IntegrationJira",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IntegrationJira_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationJira_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.integrationJira.IntegrationJira",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IntegrationJira_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationJira_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.integrationJira.IntegrationJira",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IntegrationJira_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.integrationJira.IntegrationJira",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IntegrationJira) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IntegrationJira) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IntegrationJira) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationJira) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationJira) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationJira) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationJira) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationJira) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationJira) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationJira) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationJira) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationJira) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationJira) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IntegrationJira) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationJira) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IntegrationJira) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IntegrationJira) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IntegrationJira) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IntegrationJira) ResetApiUrl() {
	_jsii_.InvokeVoid(
		i,
		"resetApiUrl",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationJira) ResetCommentOnEventEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetCommentOnEventEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationJira) ResetCommitEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetCommitEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationJira) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationJira) ResetIssuesEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetIssuesEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationJira) ResetJiraAuthType() {
	_jsii_.InvokeVoid(
		i,
		"resetJiraAuthType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationJira) ResetJiraIssuePrefix() {
	_jsii_.InvokeVoid(
		i,
		"resetJiraIssuePrefix",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationJira) ResetJiraIssueRegex() {
	_jsii_.InvokeVoid(
		i,
		"resetJiraIssueRegex",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationJira) ResetJiraIssueTransitionAutomatic() {
	_jsii_.InvokeVoid(
		i,
		"resetJiraIssueTransitionAutomatic",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationJira) ResetJiraIssueTransitionId() {
	_jsii_.InvokeVoid(
		i,
		"resetJiraIssueTransitionId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationJira) ResetMergeRequestsEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetMergeRequestsEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationJira) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationJira) ResetProjectKeys() {
	_jsii_.InvokeVoid(
		i,
		"resetProjectKeys",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationJira) ResetUseInheritedSettings() {
	_jsii_.InvokeVoid(
		i,
		"resetUseInheritedSettings",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationJira) ResetUsername() {
	_jsii_.InvokeVoid(
		i,
		"resetUsername",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationJira) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationJira) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationJira) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationJira) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationJira) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationJira) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

