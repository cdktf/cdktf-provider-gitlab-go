// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package release

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Release) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (r *jsiiProxy_Release) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (r *jsiiProxy_Release) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Release) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Release) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Release) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Release) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Release) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Release) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Release) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Release) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Release) validateImportFromParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_Release) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Release) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_Release) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (r *jsiiProxy_Release) validateMoveToIdParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_Release) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (r *jsiiProxy_Release) validatePutAssetsParameters(value *ReleaseAssets) error {
	return nil
}

func validateRelease_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateRelease_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRelease_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateRelease_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Release) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Release) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Release) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Release) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Release) validateSetMilestonesParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Release) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Release) validateSetProjectParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Release) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Release) validateSetRefParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Release) validateSetReleasedAtParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Release) validateSetTagMessageParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Release) validateSetTagNameParameters(val *string) error {
	return nil
}

func validateNewReleaseParameters(scope constructs.Construct, id *string, config *ReleaseConfig) error {
	return nil
}

