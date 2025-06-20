// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationappearance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v15/applicationappearance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/resources/application_appearance gitlab_application_appearance}.
type ApplicationAppearance interface {
	cdktf.TerraformResource
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EmailHeaderAndFooterEnabled() interface{}
	SetEmailHeaderAndFooterEnabled(val interface{})
	EmailHeaderAndFooterEnabledInput() interface{}
	FooterMessage() *string
	SetFooterMessage(val *string)
	FooterMessageInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HeaderMessage() *string
	SetHeaderMessage(val *string)
	HeaderMessageInput() *string
	Id() *string
	KeepSettingsOnDestroy() interface{}
	SetKeepSettingsOnDestroy(val interface{})
	KeepSettingsOnDestroyInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MemberGuidelines() *string
	SetMemberGuidelines(val *string)
	MemberGuidelinesInput() *string
	MessageBackgroundColor() *string
	SetMessageBackgroundColor(val *string)
	MessageBackgroundColorInput() *string
	MessageFontColor() *string
	SetMessageFontColor(val *string)
	MessageFontColorInput() *string
	NewProjectGuidelines() *string
	SetNewProjectGuidelines(val *string)
	NewProjectGuidelinesInput() *string
	// The tree node.
	Node() constructs.Node
	ProfileImageGuidelines() *string
	SetProfileImageGuidelines(val *string)
	ProfileImageGuidelinesInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PwaDescription() *string
	SetPwaDescription(val *string)
	PwaDescriptionInput() *string
	PwaName() *string
	SetPwaName(val *string)
	PwaNameInput() *string
	PwaShortName() *string
	SetPwaShortName(val *string)
	PwaShortNameInput() *string
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
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
	ResetDescription()
	ResetEmailHeaderAndFooterEnabled()
	ResetFooterMessage()
	ResetHeaderMessage()
	ResetKeepSettingsOnDestroy()
	ResetMemberGuidelines()
	ResetMessageBackgroundColor()
	ResetMessageFontColor()
	ResetNewProjectGuidelines()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProfileImageGuidelines()
	ResetPwaDescription()
	ResetPwaName()
	ResetPwaShortName()
	ResetTitle()
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

// The jsii proxy struct for ApplicationAppearance
type jsiiProxy_ApplicationAppearance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApplicationAppearance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) EmailHeaderAndFooterEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailHeaderAndFooterEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) EmailHeaderAndFooterEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailHeaderAndFooterEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) FooterMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"footerMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) FooterMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"footerMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) HeaderMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) HeaderMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) KeepSettingsOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepSettingsOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) KeepSettingsOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepSettingsOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) MemberGuidelines() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memberGuidelines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) MemberGuidelinesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memberGuidelinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) MessageBackgroundColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageBackgroundColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) MessageBackgroundColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageBackgroundColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) MessageFontColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageFontColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) MessageFontColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageFontColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) NewProjectGuidelines() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newProjectGuidelines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) NewProjectGuidelinesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newProjectGuidelinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) ProfileImageGuidelines() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileImageGuidelines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) ProfileImageGuidelinesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileImageGuidelinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) PwaDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pwaDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) PwaDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pwaDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) PwaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pwaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) PwaNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pwaNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) PwaShortName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pwaShortName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) PwaShortNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pwaShortNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationAppearance) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/resources/application_appearance gitlab_application_appearance} Resource.
func NewApplicationAppearance(scope constructs.Construct, id *string, config *ApplicationAppearanceConfig) ApplicationAppearance {
	_init_.Initialize()

	if err := validateNewApplicationAppearanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationAppearance{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.applicationAppearance.ApplicationAppearance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.1.0/docs/resources/application_appearance gitlab_application_appearance} Resource.
func NewApplicationAppearance_Override(a ApplicationAppearance, scope constructs.Construct, id *string, config *ApplicationAppearanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.applicationAppearance.ApplicationAppearance",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApplicationAppearance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApplicationAppearance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApplicationAppearance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApplicationAppearance)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ApplicationAppearance)SetEmailHeaderAndFooterEnabled(val interface{}) {
	if err := j.validateSetEmailHeaderAndFooterEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailHeaderAndFooterEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationAppearance)SetFooterMessage(val *string) {
	if err := j.validateSetFooterMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"footerMessage",
		val,
	)
}

func (j *jsiiProxy_ApplicationAppearance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApplicationAppearance)SetHeaderMessage(val *string) {
	if err := j.validateSetHeaderMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headerMessage",
		val,
	)
}

func (j *jsiiProxy_ApplicationAppearance)SetKeepSettingsOnDestroy(val interface{}) {
	if err := j.validateSetKeepSettingsOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepSettingsOnDestroy",
		val,
	)
}

func (j *jsiiProxy_ApplicationAppearance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApplicationAppearance)SetMemberGuidelines(val *string) {
	if err := j.validateSetMemberGuidelinesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memberGuidelines",
		val,
	)
}

func (j *jsiiProxy_ApplicationAppearance)SetMessageBackgroundColor(val *string) {
	if err := j.validateSetMessageBackgroundColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageBackgroundColor",
		val,
	)
}

func (j *jsiiProxy_ApplicationAppearance)SetMessageFontColor(val *string) {
	if err := j.validateSetMessageFontColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageFontColor",
		val,
	)
}

func (j *jsiiProxy_ApplicationAppearance)SetNewProjectGuidelines(val *string) {
	if err := j.validateSetNewProjectGuidelinesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newProjectGuidelines",
		val,
	)
}

func (j *jsiiProxy_ApplicationAppearance)SetProfileImageGuidelines(val *string) {
	if err := j.validateSetProfileImageGuidelinesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profileImageGuidelines",
		val,
	)
}

func (j *jsiiProxy_ApplicationAppearance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApplicationAppearance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApplicationAppearance)SetPwaDescription(val *string) {
	if err := j.validateSetPwaDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pwaDescription",
		val,
	)
}

func (j *jsiiProxy_ApplicationAppearance)SetPwaName(val *string) {
	if err := j.validateSetPwaNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pwaName",
		val,
	)
}

func (j *jsiiProxy_ApplicationAppearance)SetPwaShortName(val *string) {
	if err := j.validateSetPwaShortNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pwaShortName",
		val,
	)
}

func (j *jsiiProxy_ApplicationAppearance)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

// Generates CDKTF code for importing a ApplicationAppearance resource upon running "cdktf plan <stack-name>".
func ApplicationAppearance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateApplicationAppearance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.applicationAppearance.ApplicationAppearance",
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
func ApplicationAppearance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationAppearance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.applicationAppearance.ApplicationAppearance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApplicationAppearance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationAppearance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.applicationAppearance.ApplicationAppearance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApplicationAppearance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationAppearance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.applicationAppearance.ApplicationAppearance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApplicationAppearance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.applicationAppearance.ApplicationAppearance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApplicationAppearance) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ApplicationAppearance) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApplicationAppearance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApplicationAppearance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationAppearance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApplicationAppearance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApplicationAppearance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApplicationAppearance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApplicationAppearance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApplicationAppearance) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApplicationAppearance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApplicationAppearance) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationAppearance) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ApplicationAppearance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationAppearance) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApplicationAppearance) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ApplicationAppearance) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApplicationAppearance) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApplicationAppearance) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationAppearance) ResetEmailHeaderAndFooterEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailHeaderAndFooterEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationAppearance) ResetFooterMessage() {
	_jsii_.InvokeVoid(
		a,
		"resetFooterMessage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationAppearance) ResetHeaderMessage() {
	_jsii_.InvokeVoid(
		a,
		"resetHeaderMessage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationAppearance) ResetKeepSettingsOnDestroy() {
	_jsii_.InvokeVoid(
		a,
		"resetKeepSettingsOnDestroy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationAppearance) ResetMemberGuidelines() {
	_jsii_.InvokeVoid(
		a,
		"resetMemberGuidelines",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationAppearance) ResetMessageBackgroundColor() {
	_jsii_.InvokeVoid(
		a,
		"resetMessageBackgroundColor",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationAppearance) ResetMessageFontColor() {
	_jsii_.InvokeVoid(
		a,
		"resetMessageFontColor",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationAppearance) ResetNewProjectGuidelines() {
	_jsii_.InvokeVoid(
		a,
		"resetNewProjectGuidelines",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationAppearance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationAppearance) ResetProfileImageGuidelines() {
	_jsii_.InvokeVoid(
		a,
		"resetProfileImageGuidelines",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationAppearance) ResetPwaDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetPwaDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationAppearance) ResetPwaName() {
	_jsii_.InvokeVoid(
		a,
		"resetPwaName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationAppearance) ResetPwaShortName() {
	_jsii_.InvokeVoid(
		a,
		"resetPwaShortName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationAppearance) ResetTitle() {
	_jsii_.InvokeVoid(
		a,
		"resetTitle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationAppearance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationAppearance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationAppearance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationAppearance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationAppearance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationAppearance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

