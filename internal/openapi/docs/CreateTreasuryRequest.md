# CreateTreasuryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name (1–128 characters) | 
**SafeAddress** | **string** | Deployed Safe contract address (0x-prefixed, 42 characters) | 
**Chain** | Pointer to **string** | Chain name (default base) | [optional] 
**ChainId** | Pointer to **int32** | EVM chain ID (default 8453 for Base) | [optional] 
**Threshold** | Pointer to **int32** | Safe threshold (default 1) | [optional] 
**Signers** | Pointer to [**[]CreateTreasurySignerEntry**](CreateTreasurySignerEntry.md) |  | [optional] 

## Methods

### NewCreateTreasuryRequest

`func NewCreateTreasuryRequest(name string, safeAddress string, ) *CreateTreasuryRequest`

NewCreateTreasuryRequest instantiates a new CreateTreasuryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTreasuryRequestWithDefaults

`func NewCreateTreasuryRequestWithDefaults() *CreateTreasuryRequest`

NewCreateTreasuryRequestWithDefaults instantiates a new CreateTreasuryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTreasuryRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTreasuryRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTreasuryRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSafeAddress

`func (o *CreateTreasuryRequest) GetSafeAddress() string`

GetSafeAddress returns the SafeAddress field if non-nil, zero value otherwise.

### GetSafeAddressOk

`func (o *CreateTreasuryRequest) GetSafeAddressOk() (*string, bool)`

GetSafeAddressOk returns a tuple with the SafeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeAddress

`func (o *CreateTreasuryRequest) SetSafeAddress(v string)`

SetSafeAddress sets SafeAddress field to given value.


### GetChain

`func (o *CreateTreasuryRequest) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *CreateTreasuryRequest) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *CreateTreasuryRequest) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *CreateTreasuryRequest) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetChainId

`func (o *CreateTreasuryRequest) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *CreateTreasuryRequest) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *CreateTreasuryRequest) SetChainId(v int32)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *CreateTreasuryRequest) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetThreshold

`func (o *CreateTreasuryRequest) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *CreateTreasuryRequest) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *CreateTreasuryRequest) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *CreateTreasuryRequest) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetSigners

`func (o *CreateTreasuryRequest) GetSigners() []CreateTreasurySignerEntry`

GetSigners returns the Signers field if non-nil, zero value otherwise.

### GetSignersOk

`func (o *CreateTreasuryRequest) GetSignersOk() (*[]CreateTreasurySignerEntry, bool)`

GetSignersOk returns a tuple with the Signers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigners

`func (o *CreateTreasuryRequest) SetSigners(v []CreateTreasurySignerEntry)`

SetSigners sets Signers field to given value.

### HasSigners

`func (o *CreateTreasuryRequest) HasSigners() bool`

HasSigners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


