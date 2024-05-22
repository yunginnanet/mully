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

// checks if the V1MeGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1MeGet200Response{}

// V1MeGet200Response struct for V1MeGet200Response
type V1MeGet200Response struct {
	Token *string `json:"token,omitempty"`
	Expires *string `json:"expires,omitempty"`
}

// NewV1MeGet200Response instantiates a new V1MeGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1MeGet200Response() *V1MeGet200Response {
	this := V1MeGet200Response{}
	return &this
}

// NewV1MeGet200ResponseWithDefaults instantiates a new V1MeGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1MeGet200ResponseWithDefaults() *V1MeGet200Response {
	this := V1MeGet200Response{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *V1MeGet200Response) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1MeGet200Response) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *V1MeGet200Response) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *V1MeGet200Response) SetToken(v string) {
	o.Token = &v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *V1MeGet200Response) GetExpires() string {
	if o == nil || IsNil(o.Expires) {
		var ret string
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1MeGet200Response) GetExpiresOk() (*string, bool) {
	if o == nil || IsNil(o.Expires) {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *V1MeGet200Response) HasExpires() bool {
	if o != nil && !IsNil(o.Expires) {
		return true
	}

	return false
}

// SetExpires gets a reference to the given string and assigns it to the Expires field.
func (o *V1MeGet200Response) SetExpires(v string) {
	o.Expires = &v
}

func (o V1MeGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1MeGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Expires) {
		toSerialize["expires"] = o.Expires
	}
	return toSerialize, nil
}

type NullableV1MeGet200Response struct {
	value *V1MeGet200Response
	isSet bool
}

func (v NullableV1MeGet200Response) Get() *V1MeGet200Response {
	return v.value
}

func (v *NullableV1MeGet200Response) Set(val *V1MeGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV1MeGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV1MeGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1MeGet200Response(val *V1MeGet200Response) *NullableV1MeGet200Response {
	return &NullableV1MeGet200Response{value: val, isSet: true}
}

func (v NullableV1MeGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1MeGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


