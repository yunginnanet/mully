/*
Mullvad App API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mullvad_api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the V1ProblemReportPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ProblemReportPostRequest{}

// V1ProblemReportPostRequest struct for V1ProblemReportPostRequest
type V1ProblemReportPostRequest struct {
	Address *string `json:"address,omitempty"`
	Message *string `json:"message,omitempty"`
	Log string `json:"log"`
	Metadata map[string]interface{} `json:"metadata"`
}

type _V1ProblemReportPostRequest V1ProblemReportPostRequest

// NewV1ProblemReportPostRequest instantiates a new V1ProblemReportPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ProblemReportPostRequest(log string, metadata map[string]interface{}) *V1ProblemReportPostRequest {
	this := V1ProblemReportPostRequest{}
	this.Log = log
	this.Metadata = metadata
	return &this
}

// NewV1ProblemReportPostRequestWithDefaults instantiates a new V1ProblemReportPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ProblemReportPostRequestWithDefaults() *V1ProblemReportPostRequest {
	this := V1ProblemReportPostRequest{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *V1ProblemReportPostRequest) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ProblemReportPostRequest) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *V1ProblemReportPostRequest) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *V1ProblemReportPostRequest) SetAddress(v string) {
	o.Address = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *V1ProblemReportPostRequest) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ProblemReportPostRequest) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *V1ProblemReportPostRequest) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *V1ProblemReportPostRequest) SetMessage(v string) {
	o.Message = &v
}

// GetLog returns the Log field value
func (o *V1ProblemReportPostRequest) GetLog() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Log
}

// GetLogOk returns a tuple with the Log field value
// and a boolean to check if the value has been set.
func (o *V1ProblemReportPostRequest) GetLogOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Log, true
}

// SetLog sets field value
func (o *V1ProblemReportPostRequest) SetLog(v string) {
	o.Log = v
}

// GetMetadata returns the Metadata field value
func (o *V1ProblemReportPostRequest) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *V1ProblemReportPostRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *V1ProblemReportPostRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o V1ProblemReportPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ProblemReportPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	toSerialize["log"] = o.Log
	toSerialize["metadata"] = o.Metadata
	return toSerialize, nil
}

func (o *V1ProblemReportPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"log",
		"metadata",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varV1ProblemReportPostRequest := _V1ProblemReportPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1ProblemReportPostRequest)

	if err != nil {
		return err
	}

	*o = V1ProblemReportPostRequest(varV1ProblemReportPostRequest)

	return err
}

type NullableV1ProblemReportPostRequest struct {
	value *V1ProblemReportPostRequest
	isSet bool
}

func (v NullableV1ProblemReportPostRequest) Get() *V1ProblemReportPostRequest {
	return v.value
}

func (v *NullableV1ProblemReportPostRequest) Set(val *V1ProblemReportPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ProblemReportPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ProblemReportPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ProblemReportPostRequest(val *V1ProblemReportPostRequest) *NullableV1ProblemReportPostRequest {
	return &NullableV1ProblemReportPostRequest{value: val, isSet: true}
}

func (v NullableV1ProblemReportPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ProblemReportPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


