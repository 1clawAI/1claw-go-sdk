# AddSmartAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | **string** |  | 
**ChainId** | **int32** |  | 
**SafeAddress** | **string** |  | 
**Nonce** | Pointer to **string** |  | [optional] 
**InitData** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAddSmartAccountRequest

`func NewAddSmartAccountRequest(chain string, chainId int32, safeAddress string, ) *AddSmartAccountRequest`

NewAddSmartAccountRequest instantiates a new AddSmartAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSmartAccountRequestWithDefaults

`func NewAddSmartAccountRequestWithDefaults() *AddSmartAccountRequest`

NewAddSmartAccountRequestWithDefaults instantiates a new AddSmartAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *AddSmartAccountRequest) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *AddSmartAccountRequest) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *AddSmartAccountRequest) SetChain(v string)`

SetChain sets Chain field to given value.


### GetChainId

`func (o *AddSmartAccountRequest) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *AddSmartAccountRequest) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *AddSmartAccountRequest) SetChainId(v int32)`

SetChainId sets ChainId field to given value.


### GetSafeAddress

`func (o *AddSmartAccountRequest) GetSafeAddress() string`

GetSafeAddress returns the SafeAddress field if non-nil, zero value otherwise.

### GetSafeAddressOk

`func (o *AddSmartAccountRequest) GetSafeAddressOk() (*string, bool)`

GetSafeAddressOk returns a tuple with the SafeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeAddress

`func (o *AddSmartAccountRequest) SetSafeAddress(v string)`

SetSafeAddress sets SafeAddress field to given value.


### GetNonce

`func (o *AddSmartAccountRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *AddSmartAccountRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *AddSmartAccountRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *AddSmartAccountRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetInitData

`func (o *AddSmartAccountRequest) GetInitData() map[string]interface{}`

GetInitData returns the InitData field if non-nil, zero value otherwise.

### GetInitDataOk

`func (o *AddSmartAccountRequest) GetInitDataOk() (*map[string]interface{}, bool)`

GetInitDataOk returns a tuple with the InitData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitData

`func (o *AddSmartAccountRequest) SetInitData(v map[string]interface{})`

SetInitData sets InitData field to given value.

### HasInitData

`func (o *AddSmartAccountRequest) HasInitData() bool`

HasInitData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


