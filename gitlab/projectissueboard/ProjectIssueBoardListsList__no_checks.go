// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package projectissueboard

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProjectIssueBoardListsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_ProjectIssueBoardListsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProjectIssueBoardListsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProjectIssueBoardListsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ProjectIssueBoardListsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProjectIssueBoardListsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProjectIssueBoardListsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProjectIssueBoardListsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

