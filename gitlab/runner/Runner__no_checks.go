//go:build no_runtime_type_checking

package runner

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Runner) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (r *jsiiProxy_Runner) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Runner) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Runner) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Runner) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Runner) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Runner) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Runner) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Runner) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Runner) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Runner) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Runner) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateRunner_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRunner_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateRunner_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Runner) validateSetAccessLevelParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Runner) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Runner) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Runner) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Runner) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Runner) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Runner) validateSetLockedParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Runner) validateSetMaximumTimeoutParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Runner) validateSetPausedParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Runner) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Runner) validateSetRegistrationTokenParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Runner) validateSetRunUntaggedParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Runner) validateSetTagListParameters(val *[]*string) error {
	return nil
}

func validateNewRunnerParameters(scope constructs.Construct, id *string, config *RunnerConfig) error {
	return nil
}

