// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package deploykey

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeployKey) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (d *jsiiProxy_DeployKey) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (d *jsiiProxy_DeployKey) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeployKey) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeployKey) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeployKey) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeployKey) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeployKey) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeployKey) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeployKey) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeployKey) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeployKey) validateImportFromParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_DeployKey) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeployKey) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_DeployKey) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (d *jsiiProxy_DeployKey) validateMoveToIdParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_DeployKey) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateDeployKey_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateDeployKey_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDeployKey_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateDeployKey_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_DeployKey) validateSetCanPushParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DeployKey) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DeployKey) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DeployKey) validateSetExpiresAtParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeployKey) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeployKey) validateSetKeyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeployKey) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_DeployKey) validateSetProjectParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeployKey) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_DeployKey) validateSetTitleParameters(val *string) error {
	return nil
}

func validateNewDeployKeyParameters(scope constructs.Construct, id *string, config *DeployKeyConfig) error {
	return nil
}

