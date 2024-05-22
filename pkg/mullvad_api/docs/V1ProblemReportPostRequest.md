# V1ProblemReportPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Log** | **string** |  | 
**Metadata** | **map[string]interface{}** |  | 

## Methods

### NewV1ProblemReportPostRequest

`func NewV1ProblemReportPostRequest(log string, metadata map[string]interface{}, ) *V1ProblemReportPostRequest`

NewV1ProblemReportPostRequest instantiates a new V1ProblemReportPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ProblemReportPostRequestWithDefaults

`func NewV1ProblemReportPostRequestWithDefaults() *V1ProblemReportPostRequest`

NewV1ProblemReportPostRequestWithDefaults instantiates a new V1ProblemReportPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *V1ProblemReportPostRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *V1ProblemReportPostRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *V1ProblemReportPostRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *V1ProblemReportPostRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetMessage

`func (o *V1ProblemReportPostRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1ProblemReportPostRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1ProblemReportPostRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1ProblemReportPostRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetLog

`func (o *V1ProblemReportPostRequest) GetLog() string`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *V1ProblemReportPostRequest) GetLogOk() (*string, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *V1ProblemReportPostRequest) SetLog(v string)`

SetLog sets Log field to given value.


### GetMetadata

`func (o *V1ProblemReportPostRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1ProblemReportPostRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1ProblemReportPostRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


