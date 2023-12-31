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

// checks if the TRACEIDResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TRACEIDResponse{}

// TRACEIDResponse struct for TRACEIDResponse
type TRACEIDResponse struct {
	Id *string `json:"id,omitempty"`
}

// NewTRACEIDResponse instantiates a new TRACEIDResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTRACEIDResponse() *TRACEIDResponse {
	this := TRACEIDResponse{}
	return &this
}

// NewTRACEIDResponseWithDefaults instantiates a new TRACEIDResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTRACEIDResponseWithDefaults() *TRACEIDResponse {
	this := TRACEIDResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TRACEIDResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TRACEIDResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TRACEIDResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TRACEIDResponse) SetId(v string) {
	o.Id = &v
}

func (o TRACEIDResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TRACEIDResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableTRACEIDResponse struct {
	value *TRACEIDResponse
	isSet bool
}

func (v NullableTRACEIDResponse) Get() *TRACEIDResponse {
	return v.value
}

func (v *NullableTRACEIDResponse) Set(val *TRACEIDResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTRACEIDResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTRACEIDResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTRACEIDResponse(val *TRACEIDResponse) *NullableTRACEIDResponse {
	return &NullableTRACEIDResponse{value: val, isSet: true}
}

func (v NullableTRACEIDResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTRACEIDResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
