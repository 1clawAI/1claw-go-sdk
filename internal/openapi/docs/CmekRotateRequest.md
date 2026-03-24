# CmekRotateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewFingerprint** | **string** | SHA-256 hex fingerprint of the new CMEK key | 

## Methods

### NewCmekRotateRequest

`func NewCmekRotateRequest(newFingerprint string, ) *CmekRotateRequest`

NewCmekRotateRequest instantiates a new CmekRotateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmekRotateRequestWithDefaults

`func NewCmekRotateRequestWithDefaults() *CmekRotateRequest`

NewCmekRotateRequestWithDefaults instantiates a new CmekRotateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewFingerprint

`func (o *CmekRotateRequest) GetNewFingerprint() string`

GetNewFingerprint returns the NewFingerprint field if non-nil, zero value otherwise.

### GetNewFingerprintOk

`func (o *CmekRotateRequest) GetNewFingerprintOk() (*string, bool)`

GetNewFingerprintOk returns a tuple with the NewFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewFingerprint

`func (o *CmekRotateRequest) SetNewFingerprint(v string)`

SetNewFingerprint sets NewFingerprint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


