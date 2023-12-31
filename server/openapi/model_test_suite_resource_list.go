/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type TestSuiteResourceList struct {
	Count int32 `json:"count,omitempty"`

	Items []TestSuiteResource `json:"items,omitempty"`
}

// AssertTestSuiteResourceListRequired checks if the required fields are not zero-ed
func AssertTestSuiteResourceListRequired(obj TestSuiteResourceList) error {
	for _, el := range obj.Items {
		if err := AssertTestSuiteResourceRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseTestSuiteResourceListRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TestSuiteResourceList (e.g. [][]TestSuiteResourceList), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTestSuiteResourceListRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTestSuiteResourceList, ok := obj.(TestSuiteResourceList)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTestSuiteResourceListRequired(aTestSuiteResourceList)
	})
}
