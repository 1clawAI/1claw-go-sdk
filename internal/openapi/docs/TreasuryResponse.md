# TreasuryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SafeAddress** | Pointer to **string** |  | [optional] 
**Chain** | Pointer to **string** |  | [optional] 
**ChainId** | Pointer to **int32** |  | [optional] 
**Threshold** | Pointer to **int32** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Signers** | Pointer to [**[]TreasurySignerResponse**](TreasurySignerResponse.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTreasuryResponse

`func NewTreasuryResponse() *TreasuryResponse`

NewTreasuryResponse instantiates a new TreasuryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTreasuryResponseWithDefaults

`func NewTreasuryResponseWithDefaults() *TreasuryResponse`

NewTreasuryResponseWithDefaults instantiates a new TreasuryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TreasuryResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TreasuryResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TreasuryResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TreasuryResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TreasuryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TreasuryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TreasuryResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TreasuryResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSafeAddress

`func (o *TreasuryResponse) GetSafeAddress() string`

GetSafeAddress returns the SafeAddress field if non-nil, zero value otherwise.

### GetSafeAddressOk

`func (o *TreasuryResponse) GetSafeAddressOk() (*string, bool)`

GetSafeAddressOk returns a tuple with the SafeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeAddress

`func (o *TreasuryResponse) SetSafeAddress(v string)`

SetSafeAddress sets SafeAddress field to given value.

### HasSafeAddress

`func (o *TreasuryResponse) HasSafeAddress() bool`

HasSafeAddress returns a boolean if a field has been set.

### GetChain

`func (o *TreasuryResponse) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *TreasuryResponse) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *TreasuryResponse) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *TreasuryResponse) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetChainId

`func (o *TreasuryResponse) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *TreasuryResponse) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *TreasuryResponse) SetChainId(v int32)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *TreasuryResponse) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetThreshold

`func (o *TreasuryResponse) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *TreasuryResponse) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *TreasuryResponse) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *TreasuryResponse) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetCreatedBy

`func (o *TreasuryResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TreasuryResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TreasuryResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *TreasuryResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetSigners

`func (o *TreasuryResponse) GetSigners() []TreasurySignerResponse`

GetSigners returns the Signers field if non-nil, zero value otherwise.

### GetSignersOk

`func (o *TreasuryResponse) GetSignersOk() (*[]TreasurySignerResponse, bool)`

GetSignersOk returns a tuple with the Signers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigners

`func (o *TreasuryResponse) SetSigners(v []TreasurySignerResponse)`

SetSigners sets Signers field to given value.

### HasSigners

`func (o *TreasuryResponse) HasSigners() bool`

HasSigners returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TreasuryResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TreasuryResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TreasuryResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TreasuryResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


