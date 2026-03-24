# MfaDisableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 

## Methods

### NewMfaDisableRequest

`func NewMfaDisableRequest() *MfaDisableRequest`

NewMfaDisableRequest instantiates a new MfaDisableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMfaDisableRequestWithDefaults

`func NewMfaDisableRequestWithDefaults() *MfaDisableRequest`

NewMfaDisableRequestWithDefaults instantiates a new MfaDisableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *MfaDisableRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MfaDisableRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MfaDisableRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *MfaDisableRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetPassword

`func (o *MfaDisableRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MfaDisableRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MfaDisableRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MfaDisableRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


