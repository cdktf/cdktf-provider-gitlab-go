//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt gitlab Provider for Terraform CDK (cdktf)
package gitlab

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataGitlabProjectTagsTagsReleaseList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataGitlabProjectTagsTagsReleaseList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataGitlabProjectTagsTagsReleaseList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataGitlabProjectTagsTagsReleaseList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataGitlabProjectTagsTagsReleaseList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataGitlabProjectTagsTagsReleaseListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
