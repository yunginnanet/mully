/*
Mullvad App API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mullvad_api

import (
	"encoding/json"
)

// checks if the V1RelaysGet200ResponseOpenvpnPortsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1RelaysGet200ResponseOpenvpnPortsInner{}

// V1RelaysGet200ResponseOpenvpnPortsInner struct for V1RelaysGet200ResponseOpenvpnPortsInner
type V1RelaysGet200ResponseOpenvpnPortsInner struct {
	Port *float32 `json:"port,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
}

// NewV1RelaysGet200ResponseOpenvpnPortsInner instantiates a new V1RelaysGet200ResponseOpenvpnPortsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1RelaysGet200ResponseOpenvpnPortsInner() *V1RelaysGet200ResponseOpenvpnPortsInner {
	this := V1RelaysGet200ResponseOpenvpnPortsInner{}
	return &this
}

// NewV1RelaysGet200ResponseOpenvpnPortsInnerWithDefaults instantiates a new V1RelaysGet200ResponseOpenvpnPortsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1RelaysGet200ResponseOpenvpnPortsInnerWithDefaults() *V1RelaysGet200ResponseOpenvpnPortsInner {
	this := V1RelaysGet200ResponseOpenvpnPortsInner{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *V1RelaysGet200ResponseOpenvpnPortsInner) GetPort() float32 {
	if o == nil || IsNil(o.Port) {
		var ret float32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RelaysGet200ResponseOpenvpnPortsInner) GetPortOk() (*float32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *V1RelaysGet200ResponseOpenvpnPortsInner) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given float32 and assigns it to the Port field.
func (o *V1RelaysGet200ResponseOpenvpnPortsInner) SetPort(v float32) {
	o.Port = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *V1RelaysGet200ResponseOpenvpnPortsInner) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RelaysGet200ResponseOpenvpnPortsInner) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *V1RelaysGet200ResponseOpenvpnPortsInner) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *V1RelaysGet200ResponseOpenvpnPortsInner) SetProtocol(v string) {
	o.Protocol = &v
}

func (o V1RelaysGet200ResponseOpenvpnPortsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1RelaysGet200ResponseOpenvpnPortsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	return toSerialize, nil
}

type NullableV1RelaysGet200ResponseOpenvpnPortsInner struct {
	value *V1RelaysGet200ResponseOpenvpnPortsInner
	isSet bool
}

func (v NullableV1RelaysGet200ResponseOpenvpnPortsInner) Get() *V1RelaysGet200ResponseOpenvpnPortsInner {
	return v.value
}

func (v *NullableV1RelaysGet200ResponseOpenvpnPortsInner) Set(val *V1RelaysGet200ResponseOpenvpnPortsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV1RelaysGet200ResponseOpenvpnPortsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV1RelaysGet200ResponseOpenvpnPortsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1RelaysGet200ResponseOpenvpnPortsInner(val *V1RelaysGet200ResponseOpenvpnPortsInner) *NullableV1RelaysGet200ResponseOpenvpnPortsInner {
	return &NullableV1RelaysGet200ResponseOpenvpnPortsInner{value: val, isSet: true}
}

func (v NullableV1RelaysGet200ResponseOpenvpnPortsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1RelaysGet200ResponseOpenvpnPortsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


