// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package projecttag

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProjectTagReleaseList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_ProjectTagReleaseList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProjectTagReleaseList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProjectTagReleaseList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProjectTagReleaseList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProjectTagReleaseList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProjectTagReleaseListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

