// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package deploytoken

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeployToken) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (d *jsiiProxy_DeployToken) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (d *jsiiProxy_DeployToken) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeployToken) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeployToken) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeployToken) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeployToken) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeployToken) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeployToken) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeployToken) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeployToken) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeployToken) validateImportFromParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_DeployToken) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeployToken) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_DeployToken) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (d *jsiiProxy_DeployToken) validateMoveToIdParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_DeployToken) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateDeployToken_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateDeployToken_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDeployToken_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateDeployToken_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_DeployToken) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DeployToken) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DeployToken) validateSetExpiresAtParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeployToken) validateSetGroupParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeployToken) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeployToken) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_DeployToken) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeployToken) validateSetProjectParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeployToken) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_DeployToken) validateSetScopesParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_DeployToken) validateSetUsernameParameters(val *string) error {
	return nil
}

func validateNewDeployTokenParameters(scope constructs.Construct, id *string, config *DeployTokenConfig) error {
	return nil
}

