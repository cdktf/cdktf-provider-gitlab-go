// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datagitlabrunners

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataGitlabRunnersRunnersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataGitlabRunnersRunnersList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataGitlabRunnersRunnersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataGitlabRunnersRunnersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataGitlabRunnersRunnersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataGitlabRunnersRunnersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataGitlabRunnersRunnersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

