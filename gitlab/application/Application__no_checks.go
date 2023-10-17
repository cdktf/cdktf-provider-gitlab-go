// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package application

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Application) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (a *jsiiProxy_Application) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_Application) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Application) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Application) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Application) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Application) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Application) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Application) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Application) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Application) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Application) validateImportFromParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_Application) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Application) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (a *jsiiProxy_Application) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateApplication_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateApplication_IsConstructParameters(x interface{}) error {
	return nil
}

func validateApplication_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateApplication_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Application) validateSetConfidentialParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Application) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Application) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Application) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Application) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Application) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Application) validateSetRedirectUrlParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Application) validateSetScopesParameters(val *[]*string) error {
	return nil
}

func validateNewApplicationParameters(scope constructs.Construct, id *string, config *ApplicationConfig) error {
	return nil
}

