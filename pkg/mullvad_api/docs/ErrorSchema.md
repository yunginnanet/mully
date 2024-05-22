# ErrorSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The error code. | [optional] 
**Error** | Pointer to **string** | The error message. | [optional] 

## Methods

### NewErrorSchema

`func NewErrorSchema() *ErrorSchema`

NewErrorSchema instantiates a new ErrorSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorSchemaWithDefaults

`func NewErrorSchemaWithDefaults() *ErrorSchema`

NewErrorSchemaWithDefaults instantiates a new ErrorSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorSchema) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorSchema) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorSchema) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ErrorSchema) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetError

`func (o *ErrorSchema) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ErrorSchema) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ErrorSchema) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ErrorSchema) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


