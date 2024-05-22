# WireGuardRelay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Owned** | Pointer to **bool** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Ipv4AddrIn** | Pointer to **string** |  | [optional] 
**IncludeInCountry** | Pointer to **bool** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**Ipv6AddrIn** | Pointer to **string** |  | [optional] 
**SameIp** | Pointer to **bool** | Whether the server supports clients using the same IP address. This flag will be removed once all servers support the feature, so clients should default to True if it&#39;s missing.  | [optional] [default to true]
**Daita** | Pointer to **bool** | If true, clients can expect DAITA to be available on this relay. Note that this field is not guaranteed to be present. A missing &#x60;daita&#x60; property is semantically equivalent to &#x60;\&quot;daita\&quot;: false&#x60;.  | [optional] [default to false]

## Methods

### NewWireGuardRelay

`func NewWireGuardRelay() *WireGuardRelay`

NewWireGuardRelay instantiates a new WireGuardRelay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireGuardRelayWithDefaults

`func NewWireGuardRelayWithDefaults() *WireGuardRelay`

NewWireGuardRelayWithDefaults instantiates a new WireGuardRelay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *WireGuardRelay) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *WireGuardRelay) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *WireGuardRelay) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *WireGuardRelay) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetLocation

`func (o *WireGuardRelay) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *WireGuardRelay) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *WireGuardRelay) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *WireGuardRelay) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetActive

`func (o *WireGuardRelay) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WireGuardRelay) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WireGuardRelay) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WireGuardRelay) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetOwned

`func (o *WireGuardRelay) GetOwned() bool`

GetOwned returns the Owned field if non-nil, zero value otherwise.

### GetOwnedOk

`func (o *WireGuardRelay) GetOwnedOk() (*bool, bool)`

GetOwnedOk returns a tuple with the Owned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwned

`func (o *WireGuardRelay) SetOwned(v bool)`

SetOwned sets Owned field to given value.

### HasOwned

`func (o *WireGuardRelay) HasOwned() bool`

HasOwned returns a boolean if a field has been set.

### GetProvider

`func (o *WireGuardRelay) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *WireGuardRelay) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *WireGuardRelay) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *WireGuardRelay) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetIpv4AddrIn

`func (o *WireGuardRelay) GetIpv4AddrIn() string`

GetIpv4AddrIn returns the Ipv4AddrIn field if non-nil, zero value otherwise.

### GetIpv4AddrInOk

`func (o *WireGuardRelay) GetIpv4AddrInOk() (*string, bool)`

GetIpv4AddrInOk returns a tuple with the Ipv4AddrIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4AddrIn

`func (o *WireGuardRelay) SetIpv4AddrIn(v string)`

SetIpv4AddrIn sets Ipv4AddrIn field to given value.

### HasIpv4AddrIn

`func (o *WireGuardRelay) HasIpv4AddrIn() bool`

HasIpv4AddrIn returns a boolean if a field has been set.

### GetIncludeInCountry

`func (o *WireGuardRelay) GetIncludeInCountry() bool`

GetIncludeInCountry returns the IncludeInCountry field if non-nil, zero value otherwise.

### GetIncludeInCountryOk

`func (o *WireGuardRelay) GetIncludeInCountryOk() (*bool, bool)`

GetIncludeInCountryOk returns a tuple with the IncludeInCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInCountry

`func (o *WireGuardRelay) SetIncludeInCountry(v bool)`

SetIncludeInCountry sets IncludeInCountry field to given value.

### HasIncludeInCountry

`func (o *WireGuardRelay) HasIncludeInCountry() bool`

HasIncludeInCountry returns a boolean if a field has been set.

### GetWeight

`func (o *WireGuardRelay) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *WireGuardRelay) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *WireGuardRelay) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *WireGuardRelay) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetPublicKey

`func (o *WireGuardRelay) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *WireGuardRelay) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *WireGuardRelay) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *WireGuardRelay) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetIpv6AddrIn

`func (o *WireGuardRelay) GetIpv6AddrIn() string`

GetIpv6AddrIn returns the Ipv6AddrIn field if non-nil, zero value otherwise.

### GetIpv6AddrInOk

`func (o *WireGuardRelay) GetIpv6AddrInOk() (*string, bool)`

GetIpv6AddrInOk returns a tuple with the Ipv6AddrIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6AddrIn

`func (o *WireGuardRelay) SetIpv6AddrIn(v string)`

SetIpv6AddrIn sets Ipv6AddrIn field to given value.

### HasIpv6AddrIn

`func (o *WireGuardRelay) HasIpv6AddrIn() bool`

HasIpv6AddrIn returns a boolean if a field has been set.

### GetSameIp

`func (o *WireGuardRelay) GetSameIp() bool`

GetSameIp returns the SameIp field if non-nil, zero value otherwise.

### GetSameIpOk

`func (o *WireGuardRelay) GetSameIpOk() (*bool, bool)`

GetSameIpOk returns a tuple with the SameIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSameIp

`func (o *WireGuardRelay) SetSameIp(v bool)`

SetSameIp sets SameIp field to given value.

### HasSameIp

`func (o *WireGuardRelay) HasSameIp() bool`

HasSameIp returns a boolean if a field has been set.

### GetDaita

`func (o *WireGuardRelay) GetDaita() bool`

GetDaita returns the Daita field if non-nil, zero value otherwise.

### GetDaitaOk

`func (o *WireGuardRelay) GetDaitaOk() (*bool, bool)`

GetDaitaOk returns a tuple with the Daita field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaita

`func (o *WireGuardRelay) SetDaita(v bool)`

SetDaita sets Daita field to given value.

### HasDaita

`func (o *WireGuardRelay) HasDaita() bool`

HasDaita returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


