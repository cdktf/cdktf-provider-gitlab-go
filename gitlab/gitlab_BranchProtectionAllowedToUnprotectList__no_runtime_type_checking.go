//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt gitlab Provider for Terraform CDK (cdktf)
package gitlab

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BranchProtectionAllowedToUnprotectList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BranchProtectionAllowedToUnprotectList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BranchProtectionAllowedToUnprotectList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BranchProtectionAllowedToUnprotectList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BranchProtectionAllowedToUnprotectList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BranchProtectionAllowedToUnprotectList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBranchProtectionAllowedToUnprotectListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

