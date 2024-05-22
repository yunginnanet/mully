# V1RelaysGet200ResponseBridge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shadowsocks** | Pointer to [**[]V1RelaysGet200ResponseBridgeShadowsocksInner**](V1RelaysGet200ResponseBridgeShadowsocksInner.md) | List of connection specs for Shadowsocks ports | [optional] 
**Relays** | Pointer to [**[]BridgeRelay**](BridgeRelay.md) |  | [optional] 

## Methods

### NewV1RelaysGet200ResponseBridge

`func NewV1RelaysGet200ResponseBridge() *V1RelaysGet200ResponseBridge`

NewV1RelaysGet200ResponseBridge instantiates a new V1RelaysGet200ResponseBridge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RelaysGet200ResponseBridgeWithDefaults

`func NewV1RelaysGet200ResponseBridgeWithDefaults() *V1RelaysGet200ResponseBridge`

NewV1RelaysGet200ResponseBridgeWithDefaults instantiates a new V1RelaysGet200ResponseBridge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShadowsocks

`func (o *V1RelaysGet200ResponseBridge) GetShadowsocks() []V1RelaysGet200ResponseBridgeShadowsocksInner`

GetShadowsocks returns the Shadowsocks field if non-nil, zero value otherwise.

### GetShadowsocksOk

`func (o *V1RelaysGet200ResponseBridge) GetShadowsocksOk() (*[]V1RelaysGet200ResponseBridgeShadowsocksInner, bool)`

GetShadowsocksOk returns a tuple with the Shadowsocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadowsocks

`func (o *V1RelaysGet200ResponseBridge) SetShadowsocks(v []V1RelaysGet200ResponseBridgeShadowsocksInner)`

SetShadowsocks sets Shadowsocks field to given value.

### HasShadowsocks

`func (o *V1RelaysGet200ResponseBridge) HasShadowsocks() bool`

HasShadowsocks returns a boolean if a field has been set.

### GetRelays

`func (o *V1RelaysGet200ResponseBridge) GetRelays() []BridgeRelay`

GetRelays returns the Relays field if non-nil, zero value otherwise.

### GetRelaysOk

`func (o *V1RelaysGet200ResponseBridge) GetRelaysOk() (*[]BridgeRelay, bool)`

GetRelaysOk returns a tuple with the Relays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelays

`func (o *V1RelaysGet200ResponseBridge) SetRelays(v []BridgeRelay)`

SetRelays sets Relays field to given value.

### HasRelays

`func (o *V1RelaysGet200ResponseBridge) HasRelays() bool`

HasRelays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


