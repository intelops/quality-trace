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

// checks if the SignalFX type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignalFX{}

// SignalFX struct for SignalFX
type SignalFX struct {
	Realm *string `json:"realm,omitempty"`
	Token *string `json:"token,omitempty"`
}

// NewSignalFX instantiates a new SignalFX object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignalFX() *SignalFX {
	this := SignalFX{}
	return &this
}

// NewSignalFXWithDefaults instantiates a new SignalFX object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignalFXWithDefaults() *SignalFX {
	this := SignalFX{}
	return &this
}

// GetRealm returns the Realm field value if set, zero value otherwise.
func (o *SignalFX) GetRealm() string {
	if o == nil || isNil(o.Realm) {
		var ret string
		return ret
	}
	return *o.Realm
}

// GetRealmOk returns a tuple with the Realm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalFX) GetRealmOk() (*string, bool) {
	if o == nil || isNil(o.Realm) {
		return nil, false
	}
	return o.Realm, true
}

// HasRealm returns a boolean if a field has been set.
func (o *SignalFX) HasRealm() bool {
	if o != nil && !isNil(o.Realm) {
		return true
	}

	return false
}

// SetRealm gets a reference to the given string and assigns it to the Realm field.
func (o *SignalFX) SetRealm(v string) {
	o.Realm = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *SignalFX) GetToken() string {
	if o == nil || isNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalFX) GetTokenOk() (*string, bool) {
	if o == nil || isNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *SignalFX) HasToken() bool {
	if o != nil && !isNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *SignalFX) SetToken(v string) {
	o.Token = &v
}

func (o SignalFX) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignalFX) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Realm) {
		toSerialize["realm"] = o.Realm
	}
	if !isNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableSignalFX struct {
	value *SignalFX
	isSet bool
}

func (v NullableSignalFX) Get() *SignalFX {
	return v.value
}

func (v *NullableSignalFX) Set(val *SignalFX) {
	v.value = val
	v.isSet = true
}

func (v NullableSignalFX) IsSet() bool {
	return v.isSet
}

func (v *NullableSignalFX) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignalFX(val *SignalFX) *NullableSignalFX {
	return &NullableSignalFX{value: val, isSet: true}
}

func (v NullableSignalFX) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignalFX) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
