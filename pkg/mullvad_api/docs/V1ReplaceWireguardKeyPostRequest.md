# V1ReplaceWireguardKeyPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Old** | **string** | The WireGuard public key to remove | 
**New** | **string** | The WireGuard public key to add | 

## Methods

### NewV1ReplaceWireguardKeyPostRequest

`func NewV1ReplaceWireguardKeyPostRequest(old string, new string, ) *V1ReplaceWireguardKeyPostRequest`

NewV1ReplaceWireguardKeyPostRequest instantiates a new V1ReplaceWireguardKeyPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ReplaceWireguardKeyPostRequestWithDefaults

`func NewV1ReplaceWireguardKeyPostRequestWithDefaults() *V1ReplaceWireguardKeyPostRequest`

NewV1ReplaceWireguardKeyPostRequestWithDefaults instantiates a new V1ReplaceWireguardKeyPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOld

`func (o *V1ReplaceWireguardKeyPostRequest) GetOld() string`

GetOld returns the Old field if non-nil, zero value otherwise.

### GetOldOk

`func (o *V1ReplaceWireguardKeyPostRequest) GetOldOk() (*string, bool)`

GetOldOk returns a tuple with the Old field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOld

`func (o *V1ReplaceWireguardKeyPostRequest) SetOld(v string)`

SetOld sets Old field to given value.


### GetNew

`func (o *V1ReplaceWireguardKeyPostRequest) GetNew() string`

GetNew returns the New field if non-nil, zero value otherwise.

### GetNewOk

`func (o *V1ReplaceWireguardKeyPostRequest) GetNewOk() (*string, bool)`

GetNewOk returns a tuple with the New field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNew

`func (o *V1ReplaceWireguardKeyPostRequest) SetNew(v string)`

SetNew sets New field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


