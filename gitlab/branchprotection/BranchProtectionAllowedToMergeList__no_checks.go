// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package branchprotection

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BranchProtectionAllowedToMergeList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BranchProtectionAllowedToMergeList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BranchProtectionAllowedToMergeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BranchProtectionAllowedToMergeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BranchProtectionAllowedToMergeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BranchProtectionAllowedToMergeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BranchProtectionAllowedToMergeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBranchProtectionAllowedToMergeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

