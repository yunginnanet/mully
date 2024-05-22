# V1CreateApplePaymentPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReceiptString** | **string** | An encrypted Base64-encoded App Store receipt | 

## Methods

### NewV1CreateApplePaymentPostRequest

`func NewV1CreateApplePaymentPostRequest(receiptString string, ) *V1CreateApplePaymentPostRequest`

NewV1CreateApplePaymentPostRequest instantiates a new V1CreateApplePaymentPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1CreateApplePaymentPostRequestWithDefaults

`func NewV1CreateApplePaymentPostRequestWithDefaults() *V1CreateApplePaymentPostRequest`

NewV1CreateApplePaymentPostRequestWithDefaults instantiates a new V1CreateApplePaymentPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceiptString

`func (o *V1CreateApplePaymentPostRequest) GetReceiptString() string`

GetReceiptString returns the ReceiptString field if non-nil, zero value otherwise.

### GetReceiptStringOk

`func (o *V1CreateApplePaymentPostRequest) GetReceiptStringOk() (*string, bool)`

GetReceiptStringOk returns a tuple with the ReceiptString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptString

`func (o *V1CreateApplePaymentPostRequest) SetReceiptString(v string)`

SetReceiptString sets ReceiptString field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


