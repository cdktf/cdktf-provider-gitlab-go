// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package userrunner

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserRunner) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (u *jsiiProxy_UserRunner) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (u *jsiiProxy_UserRunner) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (u *jsiiProxy_UserRunner) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (u *jsiiProxy_UserRunner) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (u *jsiiProxy_UserRunner) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (u *jsiiProxy_UserRunner) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (u *jsiiProxy_UserRunner) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (u *jsiiProxy_UserRunner) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (u *jsiiProxy_UserRunner) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (u *jsiiProxy_UserRunner) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (u *jsiiProxy_UserRunner) validateImportFromParameters(id *string) error {
	return nil
}

func (u *jsiiProxy_UserRunner) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (u *jsiiProxy_UserRunner) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (u *jsiiProxy_UserRunner) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (u *jsiiProxy_UserRunner) validateMoveToIdParameters(id *string) error {
	return nil
}

func (u *jsiiProxy_UserRunner) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateUserRunner_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateUserRunner_IsConstructParameters(x interface{}) error {
	return nil
}

func validateUserRunner_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateUserRunner_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_UserRunner) validateSetAccessLevelParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UserRunner) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_UserRunner) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_UserRunner) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UserRunner) validateSetGroupIdParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_UserRunner) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_UserRunner) validateSetLockedParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_UserRunner) validateSetMaximumTimeoutParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_UserRunner) validateSetPausedParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_UserRunner) validateSetProjectIdParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_UserRunner) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_UserRunner) validateSetRunnerTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UserRunner) validateSetTagListParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_UserRunner) validateSetUntaggedParameters(val interface{}) error {
	return nil
}

func validateNewUserRunnerParameters(scope constructs.Construct, id *string, config *UserRunnerConfig) error {
	return nil
}

