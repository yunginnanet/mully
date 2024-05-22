# V1RelaysGet200ResponseWireguard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortRanges** | Pointer to **[][]int32** | List of port ranges that will accept connections | [optional] 
**Ipv4Gateway** | Pointer to **string** | The ipv4 gateway of the server, you can ping it to check whether a tunnel is connected or not. | [optional] 
**Ipv6Gateway** | Pointer to **string** | The ipv6 gateway of the server, you can ping it to check whether a tunnel is connected or not. | [optional] 
**Relays** | Pointer to [**[]WireGuardRelay**](WireGuardRelay.md) |  | [optional] 

## Methods

### NewV1RelaysGet200ResponseWireguard

`func NewV1RelaysGet200ResponseWireguard() *V1RelaysGet200ResponseWireguard`

NewV1RelaysGet200ResponseWireguard instantiates a new V1RelaysGet200ResponseWireguard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RelaysGet200ResponseWireguardWithDefaults

`func NewV1RelaysGet200ResponseWireguardWithDefaults() *V1RelaysGet200ResponseWireguard`

NewV1RelaysGet200ResponseWireguardWithDefaults instantiates a new V1RelaysGet200ResponseWireguard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortRanges

`func (o *V1RelaysGet200ResponseWireguard) GetPortRanges() [][]int32`

GetPortRanges returns the PortRanges field if non-nil, zero value otherwise.

### GetPortRangesOk

`func (o *V1RelaysGet200ResponseWireguard) GetPortRangesOk() (*[][]int32, bool)`

GetPortRangesOk returns a tuple with the PortRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRanges

`func (o *V1RelaysGet200ResponseWireguard) SetPortRanges(v [][]int32)`

SetPortRanges sets PortRanges field to given value.

### HasPortRanges

`func (o *V1RelaysGet200ResponseWireguard) HasPortRanges() bool`

HasPortRanges returns a boolean if a field has been set.

### GetIpv4Gateway

`func (o *V1RelaysGet200ResponseWireguard) GetIpv4Gateway() string`

GetIpv4Gateway returns the Ipv4Gateway field if non-nil, zero value otherwise.

### GetIpv4GatewayOk

`func (o *V1RelaysGet200ResponseWireguard) GetIpv4GatewayOk() (*string, bool)`

GetIpv4GatewayOk returns a tuple with the Ipv4Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Gateway

`func (o *V1RelaysGet200ResponseWireguard) SetIpv4Gateway(v string)`

SetIpv4Gateway sets Ipv4Gateway field to given value.

### HasIpv4Gateway

`func (o *V1RelaysGet200ResponseWireguard) HasIpv4Gateway() bool`

HasIpv4Gateway returns a boolean if a field has been set.

### GetIpv6Gateway

`func (o *V1RelaysGet200ResponseWireguard) GetIpv6Gateway() string`

GetIpv6Gateway returns the Ipv6Gateway field if non-nil, zero value otherwise.

### GetIpv6GatewayOk

`func (o *V1RelaysGet200ResponseWireguard) GetIpv6GatewayOk() (*string, bool)`

GetIpv6GatewayOk returns a tuple with the Ipv6Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Gateway

`func (o *V1RelaysGet200ResponseWireguard) SetIpv6Gateway(v string)`

SetIpv6Gateway sets Ipv6Gateway field to given value.

### HasIpv6Gateway

`func (o *V1RelaysGet200ResponseWireguard) HasIpv6Gateway() bool`

HasIpv6Gateway returns a boolean if a field has been set.

### GetRelays

`func (o *V1RelaysGet200ResponseWireguard) GetRelays() []WireGuardRelay`

GetRelays returns the Relays field if non-nil, zero value otherwise.

### GetRelaysOk

`func (o *V1RelaysGet200ResponseWireguard) GetRelaysOk() (*[]WireGuardRelay, bool)`

GetRelaysOk returns a tuple with the Relays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelays

`func (o *V1RelaysGet200ResponseWireguard) SetRelays(v []WireGuardRelay)`

SetRelays sets Relays field to given value.

### HasRelays

`func (o *V1RelaysGet200ResponseWireguard) HasRelays() bool`

HasRelays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


