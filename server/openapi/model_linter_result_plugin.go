/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type LinterResultPlugin struct {
	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Passed bool `json:"passed,omitempty"`

	Score int32 `json:"score,omitempty"`

	Rules []LinterResultPluginRule `json:"rules,omitempty"`
}

// AssertLinterResultPluginRequired checks if the required fields are not zero-ed
func AssertLinterResultPluginRequired(obj LinterResultPlugin) error {
	for _, el := range obj.Rules {
		if err := AssertLinterResultPluginRuleRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseLinterResultPluginRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of LinterResultPlugin (e.g. [][]LinterResultPlugin), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseLinterResultPluginRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aLinterResultPlugin, ok := obj.(LinterResultPlugin)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertLinterResultPluginRequired(aLinterResultPlugin)
	})
}
