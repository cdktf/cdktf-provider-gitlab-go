// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package groupissueboard

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GroupIssueBoardListsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GroupIssueBoardListsList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GroupIssueBoardListsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GroupIssueBoardListsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GroupIssueBoardListsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GroupIssueBoardListsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GroupIssueBoardListsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGroupIssueBoardListsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

