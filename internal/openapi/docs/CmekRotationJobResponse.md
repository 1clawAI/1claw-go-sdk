# CmekRotationJobResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**VaultId** | **string** |  | 
**OldFingerprint** | Pointer to **string** |  | [optional] 
**NewFingerprint** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**TotalSecrets** | **int32** |  | 
**Processed** | **int32** |  | 
**Error** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewCmekRotationJobResponse

`func NewCmekRotationJobResponse(id string, vaultId string, status string, totalSecrets int32, processed int32, createdAt time.Time, ) *CmekRotationJobResponse`

NewCmekRotationJobResponse instantiates a new CmekRotationJobResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmekRotationJobResponseWithDefaults

`func NewCmekRotationJobResponseWithDefaults() *CmekRotationJobResponse`

NewCmekRotationJobResponseWithDefaults instantiates a new CmekRotationJobResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CmekRotationJobResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CmekRotationJobResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CmekRotationJobResponse) SetId(v string)`

SetId sets Id field to given value.


### GetVaultId

`func (o *CmekRotationJobResponse) GetVaultId() string`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *CmekRotationJobResponse) GetVaultIdOk() (*string, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *CmekRotationJobResponse) SetVaultId(v string)`

SetVaultId sets VaultId field to given value.


### GetOldFingerprint

`func (o *CmekRotationJobResponse) GetOldFingerprint() string`

GetOldFingerprint returns the OldFingerprint field if non-nil, zero value otherwise.

### GetOldFingerprintOk

`func (o *CmekRotationJobResponse) GetOldFingerprintOk() (*string, bool)`

GetOldFingerprintOk returns a tuple with the OldFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldFingerprint

`func (o *CmekRotationJobResponse) SetOldFingerprint(v string)`

SetOldFingerprint sets OldFingerprint field to given value.

### HasOldFingerprint

`func (o *CmekRotationJobResponse) HasOldFingerprint() bool`

HasOldFingerprint returns a boolean if a field has been set.

### GetNewFingerprint

`func (o *CmekRotationJobResponse) GetNewFingerprint() string`

GetNewFingerprint returns the NewFingerprint field if non-nil, zero value otherwise.

### GetNewFingerprintOk

`func (o *CmekRotationJobResponse) GetNewFingerprintOk() (*string, bool)`

GetNewFingerprintOk returns a tuple with the NewFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewFingerprint

`func (o *CmekRotationJobResponse) SetNewFingerprint(v string)`

SetNewFingerprint sets NewFingerprint field to given value.

### HasNewFingerprint

`func (o *CmekRotationJobResponse) HasNewFingerprint() bool`

HasNewFingerprint returns a boolean if a field has been set.

### GetStatus

`func (o *CmekRotationJobResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CmekRotationJobResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CmekRotationJobResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTotalSecrets

`func (o *CmekRotationJobResponse) GetTotalSecrets() int32`

GetTotalSecrets returns the TotalSecrets field if non-nil, zero value otherwise.

### GetTotalSecretsOk

`func (o *CmekRotationJobResponse) GetTotalSecretsOk() (*int32, bool)`

GetTotalSecretsOk returns a tuple with the TotalSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSecrets

`func (o *CmekRotationJobResponse) SetTotalSecrets(v int32)`

SetTotalSecrets sets TotalSecrets field to given value.


### GetProcessed

`func (o *CmekRotationJobResponse) GetProcessed() int32`

GetProcessed returns the Processed field if non-nil, zero value otherwise.

### GetProcessedOk

`func (o *CmekRotationJobResponse) GetProcessedOk() (*int32, bool)`

GetProcessedOk returns a tuple with the Processed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessed

`func (o *CmekRotationJobResponse) SetProcessed(v int32)`

SetProcessed sets Processed field to given value.


### GetError

`func (o *CmekRotationJobResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CmekRotationJobResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CmekRotationJobResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CmekRotationJobResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStartedAt

`func (o *CmekRotationJobResponse) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *CmekRotationJobResponse) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *CmekRotationJobResponse) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *CmekRotationJobResponse) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *CmekRotationJobResponse) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *CmekRotationJobResponse) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *CmekRotationJobResponse) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *CmekRotationJobResponse) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CmekRotationJobResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CmekRotationJobResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CmekRotationJobResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


