# ChainResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ChainId** | Pointer to **int32** |  | [optional] 
**RpcUrl** | Pointer to **string** |  | [optional] 
**WsUrl** | Pointer to **string** |  | [optional] 
**ExplorerUrl** | Pointer to **string** |  | [optional] 
**NativeCurrency** | Pointer to **string** |  | [optional] 
**IsTestnet** | Pointer to **bool** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewChainResponse

`func NewChainResponse() *ChainResponse`

NewChainResponse instantiates a new ChainResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChainResponseWithDefaults

`func NewChainResponseWithDefaults() *ChainResponse`

NewChainResponseWithDefaults instantiates a new ChainResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChainResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChainResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChainResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChainResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ChainResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChainResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChainResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChainResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *ChainResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ChainResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ChainResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ChainResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetChainId

`func (o *ChainResponse) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ChainResponse) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ChainResponse) SetChainId(v int32)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *ChainResponse) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetRpcUrl

`func (o *ChainResponse) GetRpcUrl() string`

GetRpcUrl returns the RpcUrl field if non-nil, zero value otherwise.

### GetRpcUrlOk

`func (o *ChainResponse) GetRpcUrlOk() (*string, bool)`

GetRpcUrlOk returns a tuple with the RpcUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcUrl

`func (o *ChainResponse) SetRpcUrl(v string)`

SetRpcUrl sets RpcUrl field to given value.

### HasRpcUrl

`func (o *ChainResponse) HasRpcUrl() bool`

HasRpcUrl returns a boolean if a field has been set.

### GetWsUrl

`func (o *ChainResponse) GetWsUrl() string`

GetWsUrl returns the WsUrl field if non-nil, zero value otherwise.

### GetWsUrlOk

`func (o *ChainResponse) GetWsUrlOk() (*string, bool)`

GetWsUrlOk returns a tuple with the WsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWsUrl

`func (o *ChainResponse) SetWsUrl(v string)`

SetWsUrl sets WsUrl field to given value.

### HasWsUrl

`func (o *ChainResponse) HasWsUrl() bool`

HasWsUrl returns a boolean if a field has been set.

### GetExplorerUrl

`func (o *ChainResponse) GetExplorerUrl() string`

GetExplorerUrl returns the ExplorerUrl field if non-nil, zero value otherwise.

### GetExplorerUrlOk

`func (o *ChainResponse) GetExplorerUrlOk() (*string, bool)`

GetExplorerUrlOk returns a tuple with the ExplorerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplorerUrl

`func (o *ChainResponse) SetExplorerUrl(v string)`

SetExplorerUrl sets ExplorerUrl field to given value.

### HasExplorerUrl

`func (o *ChainResponse) HasExplorerUrl() bool`

HasExplorerUrl returns a boolean if a field has been set.

### GetNativeCurrency

`func (o *ChainResponse) GetNativeCurrency() string`

GetNativeCurrency returns the NativeCurrency field if non-nil, zero value otherwise.

### GetNativeCurrencyOk

`func (o *ChainResponse) GetNativeCurrencyOk() (*string, bool)`

GetNativeCurrencyOk returns a tuple with the NativeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeCurrency

`func (o *ChainResponse) SetNativeCurrency(v string)`

SetNativeCurrency sets NativeCurrency field to given value.

### HasNativeCurrency

`func (o *ChainResponse) HasNativeCurrency() bool`

HasNativeCurrency returns a boolean if a field has been set.

### GetIsTestnet

`func (o *ChainResponse) GetIsTestnet() bool`

GetIsTestnet returns the IsTestnet field if non-nil, zero value otherwise.

### GetIsTestnetOk

`func (o *ChainResponse) GetIsTestnetOk() (*bool, bool)`

GetIsTestnetOk returns a tuple with the IsTestnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTestnet

`func (o *ChainResponse) SetIsTestnet(v bool)`

SetIsTestnet sets IsTestnet field to given value.

### HasIsTestnet

`func (o *ChainResponse) HasIsTestnet() bool`

HasIsTestnet returns a boolean if a field has been set.

### GetIsEnabled

`func (o *ChainResponse) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ChainResponse) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ChainResponse) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ChainResponse) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ChainResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChainResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChainResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ChainResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ChainResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ChainResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ChainResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ChainResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


