# MfaSetupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OtpauthUri** | Pointer to **string** |  | [optional] 
**Secret** | Pointer to **string** |  | [optional] 

## Methods

### NewMfaSetupResponse

`func NewMfaSetupResponse() *MfaSetupResponse`

NewMfaSetupResponse instantiates a new MfaSetupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMfaSetupResponseWithDefaults

`func NewMfaSetupResponseWithDefaults() *MfaSetupResponse`

NewMfaSetupResponseWithDefaults instantiates a new MfaSetupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtpauthUri

`func (o *MfaSetupResponse) GetOtpauthUri() string`

GetOtpauthUri returns the OtpauthUri field if non-nil, zero value otherwise.

### GetOtpauthUriOk

`func (o *MfaSetupResponse) GetOtpauthUriOk() (*string, bool)`

GetOtpauthUriOk returns a tuple with the OtpauthUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpauthUri

`func (o *MfaSetupResponse) SetOtpauthUri(v string)`

SetOtpauthUri sets OtpauthUri field to given value.

### HasOtpauthUri

`func (o *MfaSetupResponse) HasOtpauthUri() bool`

HasOtpauthUri returns a boolean if a field has been set.

### GetSecret

`func (o *MfaSetupResponse) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *MfaSetupResponse) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *MfaSetupResponse) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *MfaSetupResponse) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


