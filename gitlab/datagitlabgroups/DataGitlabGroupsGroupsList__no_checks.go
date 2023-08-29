// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datagitlabgroups

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataGitlabGroupsGroupsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataGitlabGroupsGroupsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataGitlabGroupsGroupsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataGitlabGroupsGroupsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataGitlabGroupsGroupsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataGitlabGroupsGroupsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

