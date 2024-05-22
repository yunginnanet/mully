# V1WireguardKeysPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pubkey** | **string** | The WireGuard public key to add | 

## Methods

### NewV1WireguardKeysPostRequest

`func NewV1WireguardKeysPostRequest(pubkey string, ) *V1WireguardKeysPostRequest`

NewV1WireguardKeysPostRequest instantiates a new V1WireguardKeysPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1WireguardKeysPostRequestWithDefaults

`func NewV1WireguardKeysPostRequestWithDefaults() *V1WireguardKeysPostRequest`

NewV1WireguardKeysPostRequestWithDefaults instantiates a new V1WireguardKeysPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubkey

`func (o *V1WireguardKeysPostRequest) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *V1WireguardKeysPostRequest) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *V1WireguardKeysPostRequest) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


