// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GitlabProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (g *jsiiProxy_GitlabProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateGitlabProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateGitlabProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGitlabProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateGitlabProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_GitlabProvider) validateSetEarlyAuthCheckParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GitlabProvider) validateSetInsecureParameters(val interface{}) error {
	return nil
}

func validateNewGitlabProviderParameters(scope constructs.Construct, id *string, config *GitlabProviderConfig) error {
	return nil
}

