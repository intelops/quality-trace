/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type OutputInfo struct {
	LogLevel string `json:"logLevel,omitempty"`

	Message string `json:"message,omitempty"`

	OutputName string `json:"outputName,omitempty"`
}

// AssertOutputInfoRequired checks if the required fields are not zero-ed
func AssertOutputInfoRequired(obj OutputInfo) error {
	return nil
}

// AssertRecurseOutputInfoRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of OutputInfo (e.g. [][]OutputInfo), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseOutputInfoRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aOutputInfo, ok := obj.(OutputInfo)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertOutputInfoRequired(aOutputInfo)
	})
}
