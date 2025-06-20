// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package label

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_Label) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (l *jsiiProxy_Label) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (l *jsiiProxy_Label) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Label) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Label) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Label) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Label) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Label) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Label) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Label) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Label) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Label) validateImportFromParameters(id *string) error {
	return nil
}

func (l *jsiiProxy_Label) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Label) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (l *jsiiProxy_Label) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (l *jsiiProxy_Label) validateMoveToIdParameters(id *string) error {
	return nil
}

func (l *jsiiProxy_Label) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateLabel_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateLabel_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLabel_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateLabel_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Label) validateSetColorParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Label) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Label) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Label) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Label) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Label) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Label) validateSetProjectParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Label) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewLabelParameters(scope constructs.Construct, id *string, config *LabelConfig) error {
	return nil
}

