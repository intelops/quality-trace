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

// checks if the Resource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Resource{}

// Resource struct for Resource
type Resource struct {
	Type string      `json:"type"`
	Item interface{} `json:"item"`
}

// NewResource instantiates a new Resource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResource(type_ string, item interface{}) *Resource {
	this := Resource{}
	this.Type = type_
	this.Item = item
	return &this
}

// NewResourceWithDefaults instantiates a new Resource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceWithDefaults() *Resource {
	this := Resource{}
	return &this
}

// GetType returns the Type field value
func (o *Resource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Resource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Resource) SetType(v string) {
	o.Type = v
}

// GetItem returns the Item field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *Resource) GetItem() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Resource) GetItemOk() (*interface{}, bool) {
	if o == nil || isNil(o.Item) {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *Resource) SetItem(v interface{}) {
	o.Item = v
}

func (o Resource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Resource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: type is readOnly
	if o.Item != nil {
		toSerialize["item"] = o.Item
	}
	return toSerialize, nil
}

type NullableResource struct {
	value *Resource
	isSet bool
}

func (v NullableResource) Get() *Resource {
	return v.value
}

func (v *NullableResource) Set(val *Resource) {
	v.value = val
	v.isSet = true
}

func (v NullableResource) IsSet() bool {
	return v.isSet
}

func (v *NullableResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResource(val *Resource) *NullableResource {
	return &NullableResource{value: val, isSet: true}
}

func (v NullableResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
