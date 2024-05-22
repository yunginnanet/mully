# WireGuardPubkey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Convenience field with the public key encoded to use with the API (eg when deleting it), same as the pubkey field but urlencoded. | [optional] 
**Pubkey** | Pointer to **string** | The WireGuard public key. | [optional] 
**Ipv4Address** | Pointer to **string** | The ipv4 peer adress for WireGuard. Note that the mask may be bigger then a single IP. | [optional] 
**Ipv6Address** | Pointer to **string** | The ipv6 peer address for WireGuard. Note that the mask may be bigger then a single IP. | [optional] 

## Methods

### NewWireGuardPubkey

`func NewWireGuardPubkey() *WireGuardPubkey`

NewWireGuardPubkey instantiates a new WireGuardPubkey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireGuardPubkeyWithDefaults

`func NewWireGuardPubkeyWithDefaults() *WireGuardPubkey`

NewWireGuardPubkeyWithDefaults instantiates a new WireGuardPubkey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WireGuardPubkey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WireGuardPubkey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WireGuardPubkey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WireGuardPubkey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPubkey

`func (o *WireGuardPubkey) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *WireGuardPubkey) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *WireGuardPubkey) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.

### HasPubkey

`func (o *WireGuardPubkey) HasPubkey() bool`

HasPubkey returns a boolean if a field has been set.

### GetIpv4Address

`func (o *WireGuardPubkey) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *WireGuardPubkey) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *WireGuardPubkey) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.

### HasIpv4Address

`func (o *WireGuardPubkey) HasIpv4Address() bool`

HasIpv4Address returns a boolean if a field has been set.

### GetIpv6Address

`func (o *WireGuardPubkey) GetIpv6Address() string`

GetIpv6Address returns the Ipv6Address field if non-nil, zero value otherwise.

### GetIpv6AddressOk

`func (o *WireGuardPubkey) GetIpv6AddressOk() (*string, bool)`

GetIpv6AddressOk returns a tuple with the Ipv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Address

`func (o *WireGuardPubkey) SetIpv6Address(v string)`

SetIpv6Address sets Ipv6Address field to given value.

### HasIpv6Address

`func (o *WireGuardPubkey) HasIpv6Address() bool`

HasIpv6Address returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


