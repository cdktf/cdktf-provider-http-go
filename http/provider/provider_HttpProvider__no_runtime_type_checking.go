//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (h *jsiiProxy_HttpProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateHttpProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewHttpProviderParameters(scope constructs.Construct, id *string, config *HttpProviderConfig) error {
	return nil
}

