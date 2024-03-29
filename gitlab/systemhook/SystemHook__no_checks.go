// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package systemhook

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SystemHook) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_SystemHook) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_SystemHook) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SystemHook) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SystemHook) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SystemHook) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SystemHook) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SystemHook) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SystemHook) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SystemHook) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SystemHook) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SystemHook) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SystemHook) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SystemHook) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SystemHook) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_SystemHook) validateMoveToIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SystemHook) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateSystemHook_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateSystemHook_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSystemHook_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateSystemHook_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_SystemHook) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SystemHook) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SystemHook) validateSetEnableSslVerificationParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SystemHook) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SystemHook) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_SystemHook) validateSetMergeRequestsEventsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SystemHook) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_SystemHook) validateSetPushEventsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SystemHook) validateSetRepositoryUpdateEventsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SystemHook) validateSetTagPushEventsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SystemHook) validateSetTokenParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SystemHook) validateSetUrlParameters(val *string) error {
	return nil
}

func validateNewSystemHookParameters(scope constructs.Construct, id *string, config *SystemHookConfig) error {
	return nil
}

