// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package projectpagessettings

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProjectPagesSettingsDeploymentsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_ProjectPagesSettingsDeploymentsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProjectPagesSettingsDeploymentsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProjectPagesSettingsDeploymentsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProjectPagesSettingsDeploymentsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProjectPagesSettingsDeploymentsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProjectPagesSettingsDeploymentsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

