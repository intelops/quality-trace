/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ConfigurationResource - Represents a configuration structured into the Resources format.
type ConfigurationResource struct {

	// Represents the type of this resource. It should always be set as 'Config'.
	Type string `json:"type,omitempty"`

	Spec ConfigurationResourceSpec `json:"spec,omitempty"`
}

// AssertConfigurationResourceRequired checks if the required fields are not zero-ed
func AssertConfigurationResourceRequired(obj ConfigurationResource) error {
	if err := AssertConfigurationResourceSpecRequired(obj.Spec); err != nil {
		return err
	}
	return nil
}

// AssertRecurseConfigurationResourceRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ConfigurationResource (e.g. [][]ConfigurationResource), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseConfigurationResourceRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aConfigurationResource, ok := obj.(ConfigurationResource)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertConfigurationResourceRequired(aConfigurationResource)
	})
}
