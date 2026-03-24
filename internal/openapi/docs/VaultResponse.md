# VaultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedByType** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**CmekEnabled** | Pointer to **bool** | Whether client-managed encryption is enabled | [optional] 
**CmekFingerprint** | Pointer to **string** | SHA-256 fingerprint of the CMEK key (64 hex chars) | [optional] 

## Methods

### NewVaultResponse

`func NewVaultResponse(id string, name string, createdAt time.Time, ) *VaultResponse`

NewVaultResponse instantiates a new VaultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultResponseWithDefaults

`func NewVaultResponseWithDefaults() *VaultResponse`

NewVaultResponseWithDefaults instantiates a new VaultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VaultResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VaultResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VaultResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *VaultResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VaultResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VaultResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VaultResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VaultResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VaultResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VaultResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedBy

`func (o *VaultResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *VaultResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *VaultResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *VaultResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByType

`func (o *VaultResponse) GetCreatedByType() string`

GetCreatedByType returns the CreatedByType field if non-nil, zero value otherwise.

### GetCreatedByTypeOk

`func (o *VaultResponse) GetCreatedByTypeOk() (*string, bool)`

GetCreatedByTypeOk returns a tuple with the CreatedByType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByType

`func (o *VaultResponse) SetCreatedByType(v string)`

SetCreatedByType sets CreatedByType field to given value.

### HasCreatedByType

`func (o *VaultResponse) HasCreatedByType() bool`

HasCreatedByType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VaultResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VaultResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VaultResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCmekEnabled

`func (o *VaultResponse) GetCmekEnabled() bool`

GetCmekEnabled returns the CmekEnabled field if non-nil, zero value otherwise.

### GetCmekEnabledOk

`func (o *VaultResponse) GetCmekEnabledOk() (*bool, bool)`

GetCmekEnabledOk returns a tuple with the CmekEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmekEnabled

`func (o *VaultResponse) SetCmekEnabled(v bool)`

SetCmekEnabled sets CmekEnabled field to given value.

### HasCmekEnabled

`func (o *VaultResponse) HasCmekEnabled() bool`

HasCmekEnabled returns a boolean if a field has been set.

### GetCmekFingerprint

`func (o *VaultResponse) GetCmekFingerprint() string`

GetCmekFingerprint returns the CmekFingerprint field if non-nil, zero value otherwise.

### GetCmekFingerprintOk

`func (o *VaultResponse) GetCmekFingerprintOk() (*string, bool)`

GetCmekFingerprintOk returns a tuple with the CmekFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmekFingerprint

`func (o *VaultResponse) SetCmekFingerprint(v string)`

SetCmekFingerprint sets CmekFingerprint field to given value.

### HasCmekFingerprint

`func (o *VaultResponse) HasCmekFingerprint() bool`

HasCmekFingerprint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


