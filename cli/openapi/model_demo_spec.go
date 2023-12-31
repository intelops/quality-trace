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

// checks if the DemoSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DemoSpec{}

// DemoSpec Represents the attributes of a Demonstration API.
type DemoSpec struct {
	Id *string `json:"id,omitempty"`
	// String defining that this demo is a Open Telemetry Store demo.
	Type *string `json:"type,omitempty"`
	// Name of the demo
	Name *string `json:"name,omitempty"`
	// Flag telling if this API is enabled on Tracetest.
	Enabled            bool                    `json:"enabled"`
	Pokeshop           *DemoPokeshop           `json:"pokeshop,omitempty"`
	OpentelemetryStore *DemoOpenTelemetryStore `json:"opentelemetryStore,omitempty"`
}

// NewDemoSpec instantiates a new DemoSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDemoSpec(enabled bool) *DemoSpec {
	this := DemoSpec{}
	this.Enabled = enabled
	return &this
}

// NewDemoSpecWithDefaults instantiates a new DemoSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDemoSpecWithDefaults() *DemoSpec {
	this := DemoSpec{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DemoSpec) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DemoSpec) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DemoSpec) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DemoSpec) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DemoSpec) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DemoSpec) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DemoSpec) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DemoSpec) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DemoSpec) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DemoSpec) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DemoSpec) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DemoSpec) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value
func (o *DemoSpec) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DemoSpec) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DemoSpec) SetEnabled(v bool) {
	o.Enabled = v
}

// GetPokeshop returns the Pokeshop field value if set, zero value otherwise.
func (o *DemoSpec) GetPokeshop() DemoPokeshop {
	if o == nil || isNil(o.Pokeshop) {
		var ret DemoPokeshop
		return ret
	}
	return *o.Pokeshop
}

// GetPokeshopOk returns a tuple with the Pokeshop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DemoSpec) GetPokeshopOk() (*DemoPokeshop, bool) {
	if o == nil || isNil(o.Pokeshop) {
		return nil, false
	}
	return o.Pokeshop, true
}

// HasPokeshop returns a boolean if a field has been set.
func (o *DemoSpec) HasPokeshop() bool {
	if o != nil && !isNil(o.Pokeshop) {
		return true
	}

	return false
}

// SetPokeshop gets a reference to the given DemoPokeshop and assigns it to the Pokeshop field.
func (o *DemoSpec) SetPokeshop(v DemoPokeshop) {
	o.Pokeshop = &v
}

// GetOpentelemetryStore returns the OpentelemetryStore field value if set, zero value otherwise.
func (o *DemoSpec) GetOpentelemetryStore() DemoOpenTelemetryStore {
	if o == nil || isNil(o.OpentelemetryStore) {
		var ret DemoOpenTelemetryStore
		return ret
	}
	return *o.OpentelemetryStore
}

// GetOpentelemetryStoreOk returns a tuple with the OpentelemetryStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DemoSpec) GetOpentelemetryStoreOk() (*DemoOpenTelemetryStore, bool) {
	if o == nil || isNil(o.OpentelemetryStore) {
		return nil, false
	}
	return o.OpentelemetryStore, true
}

// HasOpentelemetryStore returns a boolean if a field has been set.
func (o *DemoSpec) HasOpentelemetryStore() bool {
	if o != nil && !isNil(o.OpentelemetryStore) {
		return true
	}

	return false
}

// SetOpentelemetryStore gets a reference to the given DemoOpenTelemetryStore and assigns it to the OpentelemetryStore field.
func (o *DemoSpec) SetOpentelemetryStore(v DemoOpenTelemetryStore) {
	o.OpentelemetryStore = &v
}

func (o DemoSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DemoSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["enabled"] = o.Enabled
	if !isNil(o.Pokeshop) {
		toSerialize["pokeshop"] = o.Pokeshop
	}
	if !isNil(o.OpentelemetryStore) {
		toSerialize["opentelemetryStore"] = o.OpentelemetryStore
	}
	return toSerialize, nil
}

type NullableDemoSpec struct {
	value *DemoSpec
	isSet bool
}

func (v NullableDemoSpec) Get() *DemoSpec {
	return v.value
}

func (v *NullableDemoSpec) Set(val *DemoSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableDemoSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableDemoSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDemoSpec(val *DemoSpec) *NullableDemoSpec {
	return &NullableDemoSpec{value: val, isSet: true}
}

func (v NullableDemoSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDemoSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
