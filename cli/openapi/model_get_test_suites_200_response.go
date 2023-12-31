/*
TraceTest

OpenAPI definition for TraceTest endpoint and resources

API version: 0.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GetTestSuites200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTestSuites200Response{}

// GetTestSuites200Response struct for GetTestSuites200Response
type GetTestSuites200Response struct {
	Count *int32              `json:"count,omitempty"`
	Items []TestSuiteResource `json:"items,omitempty"`
}

// NewGetTestSuites200Response instantiates a new GetTestSuites200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTestSuites200Response() *GetTestSuites200Response {
	this := GetTestSuites200Response{}
	return &this
}

// NewGetTestSuites200ResponseWithDefaults instantiates a new GetTestSuites200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTestSuites200ResponseWithDefaults() *GetTestSuites200Response {
	this := GetTestSuites200Response{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GetTestSuites200Response) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTestSuites200Response) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GetTestSuites200Response) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *GetTestSuites200Response) SetCount(v int32) {
	o.Count = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *GetTestSuites200Response) GetItems() []TestSuiteResource {
	if o == nil || isNil(o.Items) {
		var ret []TestSuiteResource
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTestSuites200Response) GetItemsOk() ([]TestSuiteResource, bool) {
	if o == nil || isNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *GetTestSuites200Response) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []TestSuiteResource and assigns it to the Items field.
func (o *GetTestSuites200Response) SetItems(v []TestSuiteResource) {
	o.Items = v
}

func (o GetTestSuites200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTestSuites200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableGetTestSuites200Response struct {
	value *GetTestSuites200Response
	isSet bool
}

func (v NullableGetTestSuites200Response) Get() *GetTestSuites200Response {
	return v.value
}

func (v *NullableGetTestSuites200Response) Set(val *GetTestSuites200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTestSuites200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTestSuites200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTestSuites200Response(val *GetTestSuites200Response) *NullableGetTestSuites200Response {
	return &NullableGetTestSuites200Response{value: val, isSet: true}
}

func (v NullableGetTestSuites200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTestSuites200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
