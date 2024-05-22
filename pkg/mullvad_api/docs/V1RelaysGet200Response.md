# V1RelaysGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locations** | Pointer to [**map[string]Location**](Location.md) |  | [optional] 
**Openvpn** | Pointer to [**V1RelaysGet200ResponseOpenvpn**](V1RelaysGet200ResponseOpenvpn.md) |  | [optional] 
**Wireguard** | Pointer to [**V1RelaysGet200ResponseWireguard**](V1RelaysGet200ResponseWireguard.md) |  | [optional] 
**Bridge** | Pointer to [**V1RelaysGet200ResponseBridge**](V1RelaysGet200ResponseBridge.md) |  | [optional] 

## Methods

### NewV1RelaysGet200Response

`func NewV1RelaysGet200Response() *V1RelaysGet200Response`

NewV1RelaysGet200Response instantiates a new V1RelaysGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RelaysGet200ResponseWithDefaults

`func NewV1RelaysGet200ResponseWithDefaults() *V1RelaysGet200Response`

NewV1RelaysGet200ResponseWithDefaults instantiates a new V1RelaysGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocations

`func (o *V1RelaysGet200Response) GetLocations() map[string]Location`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *V1RelaysGet200Response) GetLocationsOk() (*map[string]Location, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *V1RelaysGet200Response) SetLocations(v map[string]Location)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *V1RelaysGet200Response) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetOpenvpn

`func (o *V1RelaysGet200Response) GetOpenvpn() V1RelaysGet200ResponseOpenvpn`

GetOpenvpn returns the Openvpn field if non-nil, zero value otherwise.

### GetOpenvpnOk

`func (o *V1RelaysGet200Response) GetOpenvpnOk() (*V1RelaysGet200ResponseOpenvpn, bool)`

GetOpenvpnOk returns a tuple with the Openvpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenvpn

`func (o *V1RelaysGet200Response) SetOpenvpn(v V1RelaysGet200ResponseOpenvpn)`

SetOpenvpn sets Openvpn field to given value.

### HasOpenvpn

`func (o *V1RelaysGet200Response) HasOpenvpn() bool`

HasOpenvpn returns a boolean if a field has been set.

### GetWireguard

`func (o *V1RelaysGet200Response) GetWireguard() V1RelaysGet200ResponseWireguard`

GetWireguard returns the Wireguard field if non-nil, zero value otherwise.

### GetWireguardOk

`func (o *V1RelaysGet200Response) GetWireguardOk() (*V1RelaysGet200ResponseWireguard, bool)`

GetWireguardOk returns a tuple with the Wireguard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireguard

`func (o *V1RelaysGet200Response) SetWireguard(v V1RelaysGet200ResponseWireguard)`

SetWireguard sets Wireguard field to given value.

### HasWireguard

`func (o *V1RelaysGet200Response) HasWireguard() bool`

HasWireguard returns a boolean if a field has been set.

### GetBridge

`func (o *V1RelaysGet200Response) GetBridge() V1RelaysGet200ResponseBridge`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *V1RelaysGet200Response) GetBridgeOk() (*V1RelaysGet200ResponseBridge, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *V1RelaysGet200Response) SetBridge(v V1RelaysGet200ResponseBridge)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *V1RelaysGet200Response) HasBridge() bool`

HasBridge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


