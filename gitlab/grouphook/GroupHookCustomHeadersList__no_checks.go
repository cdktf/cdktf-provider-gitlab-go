// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package grouphook

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GroupHookCustomHeadersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GroupHookCustomHeadersList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GroupHookCustomHeadersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GroupHookCustomHeadersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GroupHookCustomHeadersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GroupHookCustomHeadersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GroupHookCustomHeadersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGroupHookCustomHeadersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

