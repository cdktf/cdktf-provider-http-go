// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (h *jsiiProxy_HttpProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateHttpProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateHttpProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateHttpProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateHttpProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func validateNewHttpProviderParameters(scope constructs.Construct, id *string, config *HttpProviderConfig) error {
	return nil
}

