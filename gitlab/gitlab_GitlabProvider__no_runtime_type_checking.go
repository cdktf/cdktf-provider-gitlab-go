//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt gitlab Provider for Terraform CDK (cdktf)
package gitlab

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GitlabProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (g *jsiiProxy_GitlabProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateGitlabProvider_IsConstructParameters(x interface{}) error {
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

