// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datagitlabprojecthooks

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataGitlabProjectHooksHooksList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataGitlabProjectHooksHooksList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataGitlabProjectHooksHooksList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataGitlabProjectHooksHooksListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

