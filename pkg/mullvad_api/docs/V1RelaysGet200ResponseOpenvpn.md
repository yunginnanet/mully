# V1RelaysGet200ResponseOpenvpn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ports** | Pointer to [**[]V1RelaysGet200ResponseOpenvpnPortsInner**](V1RelaysGet200ResponseOpenvpnPortsInner.md) | Pairs of port/protocol that will accept connections | [optional] 
**Relays** | Pointer to [**[]OpenVpnRelay**](OpenVpnRelay.md) |  | [optional] 

## Methods

### NewV1RelaysGet200ResponseOpenvpn

`func NewV1RelaysGet200ResponseOpenvpn() *V1RelaysGet200ResponseOpenvpn`

NewV1RelaysGet200ResponseOpenvpn instantiates a new V1RelaysGet200ResponseOpenvpn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RelaysGet200ResponseOpenvpnWithDefaults

`func NewV1RelaysGet200ResponseOpenvpnWithDefaults() *V1RelaysGet200ResponseOpenvpn`

NewV1RelaysGet200ResponseOpenvpnWithDefaults instantiates a new V1RelaysGet200ResponseOpenvpn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPorts

`func (o *V1RelaysGet200ResponseOpenvpn) GetPorts() []V1RelaysGet200ResponseOpenvpnPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *V1RelaysGet200ResponseOpenvpn) GetPortsOk() (*[]V1RelaysGet200ResponseOpenvpnPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *V1RelaysGet200ResponseOpenvpn) SetPorts(v []V1RelaysGet200ResponseOpenvpnPortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *V1RelaysGet200ResponseOpenvpn) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetRelays

`func (o *V1RelaysGet200ResponseOpenvpn) GetRelays() []OpenVpnRelay`

GetRelays returns the Relays field if non-nil, zero value otherwise.

### GetRelaysOk

`func (o *V1RelaysGet200ResponseOpenvpn) GetRelaysOk() (*[]OpenVpnRelay, bool)`

GetRelaysOk returns a tuple with the Relays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelays

`func (o *V1RelaysGet200ResponseOpenvpn) SetRelays(v []OpenVpnRelay)`

SetRelays sets Relays field to given value.

### HasRelays

`func (o *V1RelaysGet200ResponseOpenvpn) HasRelays() bool`

HasRelays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


