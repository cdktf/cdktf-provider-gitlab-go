//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt gitlab Provider for Terraform CDK (cdktf)
package gitlab

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataGitlabProjectProtectedBranchesProtectedBranchesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataGitlabProjectProtectedBranchesProtectedBranchesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

