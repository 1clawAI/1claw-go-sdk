# CreateChainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DisplayName** | **string** |  | 
**ChainId** | **int32** |  | 
**RpcUrl** | Pointer to **string** |  | [optional] 
**WsUrl** | Pointer to **string** |  | [optional] 
**ExplorerUrl** | Pointer to **string** |  | [optional] 
**NativeCurrency** | Pointer to **string** |  | [optional] [default to "ETH"]
**IsTestnet** | Pointer to **bool** |  | [optional] [default to false]
**IsEnabled** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewCreateChainRequest

`func NewCreateChainRequest(name string, displayName string, chainId int32, ) *CreateChainRequest`

NewCreateChainRequest instantiates a new CreateChainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateChainRequestWithDefaults

`func NewCreateChainRequestWithDefaults() *CreateChainRequest`

NewCreateChainRequestWithDefaults instantiates a new CreateChainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateChainRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateChainRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateChainRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *CreateChainRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateChainRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateChainRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetChainId

`func (o *CreateChainRequest) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *CreateChainRequest) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *CreateChainRequest) SetChainId(v int32)`

SetChainId sets ChainId field to given value.


### GetRpcUrl

`func (o *CreateChainRequest) GetRpcUrl() string`

GetRpcUrl returns the RpcUrl field if non-nil, zero value otherwise.

### GetRpcUrlOk

`func (o *CreateChainRequest) GetRpcUrlOk() (*string, bool)`

GetRpcUrlOk returns a tuple with the RpcUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcUrl

`func (o *CreateChainRequest) SetRpcUrl(v string)`

SetRpcUrl sets RpcUrl field to given value.

### HasRpcUrl

`func (o *CreateChainRequest) HasRpcUrl() bool`

HasRpcUrl returns a boolean if a field has been set.

### GetWsUrl

`func (o *CreateChainRequest) GetWsUrl() string`

GetWsUrl returns the WsUrl field if non-nil, zero value otherwise.

### GetWsUrlOk

`func (o *CreateChainRequest) GetWsUrlOk() (*string, bool)`

GetWsUrlOk returns a tuple with the WsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWsUrl

`func (o *CreateChainRequest) SetWsUrl(v string)`

SetWsUrl sets WsUrl field to given value.

### HasWsUrl

`func (o *CreateChainRequest) HasWsUrl() bool`

HasWsUrl returns a boolean if a field has been set.

### GetExplorerUrl

`func (o *CreateChainRequest) GetExplorerUrl() string`

GetExplorerUrl returns the ExplorerUrl field if non-nil, zero value otherwise.

### GetExplorerUrlOk

`func (o *CreateChainRequest) GetExplorerUrlOk() (*string, bool)`

GetExplorerUrlOk returns a tuple with the ExplorerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplorerUrl

`func (o *CreateChainRequest) SetExplorerUrl(v string)`

SetExplorerUrl sets ExplorerUrl field to given value.

### HasExplorerUrl

`func (o *CreateChainRequest) HasExplorerUrl() bool`

HasExplorerUrl returns a boolean if a field has been set.

### GetNativeCurrency

`func (o *CreateChainRequest) GetNativeCurrency() string`

GetNativeCurrency returns the NativeCurrency field if non-nil, zero value otherwise.

### GetNativeCurrencyOk

`func (o *CreateChainRequest) GetNativeCurrencyOk() (*string, bool)`

GetNativeCurrencyOk returns a tuple with the NativeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeCurrency

`func (o *CreateChainRequest) SetNativeCurrency(v string)`

SetNativeCurrency sets NativeCurrency field to given value.

### HasNativeCurrency

`func (o *CreateChainRequest) HasNativeCurrency() bool`

HasNativeCurrency returns a boolean if a field has been set.

### GetIsTestnet

`func (o *CreateChainRequest) GetIsTestnet() bool`

GetIsTestnet returns the IsTestnet field if non-nil, zero value otherwise.

### GetIsTestnetOk

`func (o *CreateChainRequest) GetIsTestnetOk() (*bool, bool)`

GetIsTestnetOk returns a tuple with the IsTestnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTestnet

`func (o *CreateChainRequest) SetIsTestnet(v bool)`

SetIsTestnet sets IsTestnet field to given value.

### HasIsTestnet

`func (o *CreateChainRequest) HasIsTestnet() bool`

HasIsTestnet returns a boolean if a field has been set.

### GetIsEnabled

`func (o *CreateChainRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *CreateChainRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *CreateChainRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *CreateChainRequest) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


