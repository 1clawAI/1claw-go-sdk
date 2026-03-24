# MfaVerifyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**MfaToken** | **string** |  | 

## Methods

### NewMfaVerifyRequest

`func NewMfaVerifyRequest(code string, mfaToken string, ) *MfaVerifyRequest`

NewMfaVerifyRequest instantiates a new MfaVerifyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMfaVerifyRequestWithDefaults

`func NewMfaVerifyRequestWithDefaults() *MfaVerifyRequest`

NewMfaVerifyRequestWithDefaults instantiates a new MfaVerifyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *MfaVerifyRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MfaVerifyRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MfaVerifyRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetMfaToken

`func (o *MfaVerifyRequest) GetMfaToken() string`

GetMfaToken returns the MfaToken field if non-nil, zero value otherwise.

### GetMfaTokenOk

`func (o *MfaVerifyRequest) GetMfaTokenOk() (*string, bool)`

GetMfaTokenOk returns a tuple with the MfaToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaToken

`func (o *MfaVerifyRequest) SetMfaToken(v string)`

SetMfaToken sets MfaToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


