# EnableCmekRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fingerprint** | **string** | SHA-256 hex fingerprint of the CMEK key (64 chars) | 

## Methods

### NewEnableCmekRequest

`func NewEnableCmekRequest(fingerprint string, ) *EnableCmekRequest`

NewEnableCmekRequest instantiates a new EnableCmekRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableCmekRequestWithDefaults

`func NewEnableCmekRequestWithDefaults() *EnableCmekRequest`

NewEnableCmekRequestWithDefaults instantiates a new EnableCmekRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFingerprint

`func (o *EnableCmekRequest) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *EnableCmekRequest) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *EnableCmekRequest) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


