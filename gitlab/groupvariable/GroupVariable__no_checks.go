// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package groupvariable

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GroupVariable) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (g *jsiiProxy_GroupVariable) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (g *jsiiProxy_GroupVariable) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GroupVariable) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GroupVariable) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GroupVariable) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GroupVariable) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GroupVariable) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GroupVariable) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GroupVariable) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GroupVariable) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GroupVariable) validateImportFromParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GroupVariable) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GroupVariable) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GroupVariable) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (g *jsiiProxy_GroupVariable) validateMoveToIdParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GroupVariable) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateGroupVariable_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateGroupVariable_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGroupVariable_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateGroupVariable_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_GroupVariable) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GroupVariable) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GroupVariable) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GroupVariable) validateSetEnvironmentScopeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GroupVariable) validateSetGroupParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GroupVariable) validateSetHiddenParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GroupVariable) validateSetKeyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GroupVariable) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_GroupVariable) validateSetMaskedParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GroupVariable) validateSetProtectedParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GroupVariable) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_GroupVariable) validateSetRawParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GroupVariable) validateSetValueParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GroupVariable) validateSetVariableTypeParameters(val *string) error {
	return nil
}

func validateNewGroupVariableParameters(scope constructs.Construct, id *string, config *GroupVariableConfig) error {
	return nil
}

