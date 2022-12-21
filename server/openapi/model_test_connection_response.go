/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type TestConnectionResponse struct {
	Successful bool `json:"successful,omitempty"`

	Steps []ConnectionResult `json:"steps,omitempty"`
}

// AssertTestConnectionResponseRequired checks if the required fields are not zero-ed
func AssertTestConnectionResponseRequired(obj TestConnectionResponse) error {
	for _, el := range obj.Steps {
		if err := AssertConnectionResultRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseTestConnectionResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TestConnectionResponse (e.g. [][]TestConnectionResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTestConnectionResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTestConnectionResponse, ok := obj.(TestConnectionResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTestConnectionResponseRequired(aTestConnectionResponse)
	})
}