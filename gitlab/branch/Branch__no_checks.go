// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package branch

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_Branch) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (b *jsiiProxy_Branch) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (b *jsiiProxy_Branch) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_Branch) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_Branch) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_Branch) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_Branch) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_Branch) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_Branch) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_Branch) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_Branch) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_Branch) validateImportFromParameters(id *string) error {
	return nil
}

func (b *jsiiProxy_Branch) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_Branch) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (b *jsiiProxy_Branch) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (b *jsiiProxy_Branch) validateMoveToIdParameters(id *string) error {
	return nil
}

func (b *jsiiProxy_Branch) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateBranch_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateBranch_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBranch_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateBranch_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Branch) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Branch) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Branch) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Branch) validateSetKeepOnDestroyParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Branch) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Branch) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Branch) validateSetProjectParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Branch) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Branch) validateSetRefParameters(val *string) error {
	return nil
}

func validateNewBranchParameters(scope constructs.Construct, id *string, config *BranchConfig) error {
	return nil
}

