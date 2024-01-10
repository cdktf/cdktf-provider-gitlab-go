// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package projectissue

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProjectIssueTaskCompletionStatusList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_ProjectIssueTaskCompletionStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProjectIssueTaskCompletionStatusList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProjectIssueTaskCompletionStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProjectIssueTaskCompletionStatusList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProjectIssueTaskCompletionStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProjectIssueTaskCompletionStatusListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

