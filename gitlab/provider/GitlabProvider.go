// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-gitlab-go/gitlab/v13/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.3.1/docs gitlab}.
type GitlabProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	BaseUrl() *string
	SetBaseUrl(val *string)
	BaseUrlInput() *string
	CacertFile() *string
	SetCacertFile(val *string)
	CacertFileInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientCert() *string
	SetClientCert(val *string)
	ClientCertInput() *string
	ClientKey() *string
	SetClientKey(val *string)
	ClientKeyInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	EarlyAuthCheck() interface{}
	SetEarlyAuthCheck(val interface{})
	EarlyAuthCheckInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Insecure() interface{}
	SetInsecure(val interface{})
	InsecureInput() interface{}
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetBaseUrl()
	ResetCacertFile()
	ResetClientCert()
	ResetClientKey()
	ResetEarlyAuthCheck()
	ResetInsecure()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetToken()
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

// The jsii proxy struct for GitlabProvider
type jsiiProxy_GitlabProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_GitlabProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabProvider) BaseUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabProvider) BaseUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabProvider) CacertFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacertFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabProvider) CacertFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacertFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabProvider) ClientCert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabProvider) ClientCertInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabProvider) ClientKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabProvider) ClientKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabProvider) EarlyAuthCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"earlyAuthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabProvider) EarlyAuthCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"earlyAuthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabProvider) Insecure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabProvider) InsecureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabProvider) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabProvider) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.3.1/docs gitlab} Resource.
func NewGitlabProvider(scope constructs.Construct, id *string, config *GitlabProviderConfig) GitlabProvider {
	_init_.Initialize()

	if err := validateNewGitlabProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GitlabProvider{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.provider.GitlabProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/gitlabhq/gitlab/17.3.1/docs gitlab} Resource.
func NewGitlabProvider_Override(g GitlabProvider, scope constructs.Construct, id *string, config *GitlabProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.provider.GitlabProvider",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GitlabProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_GitlabProvider)SetBaseUrl(val *string) {
	_jsii_.Set(
		j,
		"baseUrl",
		val,
	)
}

func (j *jsiiProxy_GitlabProvider)SetCacertFile(val *string) {
	_jsii_.Set(
		j,
		"cacertFile",
		val,
	)
}

func (j *jsiiProxy_GitlabProvider)SetClientCert(val *string) {
	_jsii_.Set(
		j,
		"clientCert",
		val,
	)
}

func (j *jsiiProxy_GitlabProvider)SetClientKey(val *string) {
	_jsii_.Set(
		j,
		"clientKey",
		val,
	)
}

func (j *jsiiProxy_GitlabProvider)SetEarlyAuthCheck(val interface{}) {
	if err := j.validateSetEarlyAuthCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"earlyAuthCheck",
		val,
	)
}

func (j *jsiiProxy_GitlabProvider)SetInsecure(val interface{}) {
	if err := j.validateSetInsecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecure",
		val,
	)
}

func (j *jsiiProxy_GitlabProvider)SetToken(val *string) {
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

// Generates CDKTF code for importing a GitlabProvider resource upon running "cdktf plan <stack-name>".
func GitlabProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGitlabProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.provider.GitlabProvider",
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
func GitlabProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGitlabProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.provider.GitlabProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GitlabProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGitlabProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.provider.GitlabProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GitlabProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGitlabProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gitlab.provider.GitlabProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GitlabProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-gitlab.provider.GitlabProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GitlabProvider) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GitlabProvider) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GitlabProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		g,
		"resetAlias",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GitlabProvider) ResetBaseUrl() {
	_jsii_.InvokeVoid(
		g,
		"resetBaseUrl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GitlabProvider) ResetCacertFile() {
	_jsii_.InvokeVoid(
		g,
		"resetCacertFile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GitlabProvider) ResetClientCert() {
	_jsii_.InvokeVoid(
		g,
		"resetClientCert",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GitlabProvider) ResetClientKey() {
	_jsii_.InvokeVoid(
		g,
		"resetClientKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GitlabProvider) ResetEarlyAuthCheck() {
	_jsii_.InvokeVoid(
		g,
		"resetEarlyAuthCheck",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GitlabProvider) ResetInsecure() {
	_jsii_.InvokeVoid(
		g,
		"resetInsecure",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GitlabProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GitlabProvider) ResetToken() {
	_jsii_.InvokeVoid(
		g,
		"resetToken",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GitlabProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitlabProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitlabProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitlabProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitlabProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitlabProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

