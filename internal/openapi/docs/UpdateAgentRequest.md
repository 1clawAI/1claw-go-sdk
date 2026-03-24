# UpdateAgentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**IntentsApiEnabled** | Pointer to **bool** |  | [optional] 
**TxToAllowlist** | Pointer to **[]string** |  | [optional] 
**TxMaxValueEth** | Pointer to **string** |  | [optional] 
**TxDailyLimitEth** | Pointer to **string** |  | [optional] 
**TxAllowedChains** | Pointer to **[]string** |  | [optional] 
**TokenTtlSeconds** | Pointer to **int32** |  | [optional] 
**VaultIds** | Pointer to **[]string** |  | [optional] 
**ShroudEnabled** | Pointer to **bool** | Enable/disable Shroud LLM Proxy | [optional] 
**ShroudConfig** | Pointer to [**ShroudConfig**](ShroudConfig.md) |  | [optional] 

## Methods

### NewUpdateAgentRequest

`func NewUpdateAgentRequest() *UpdateAgentRequest`

NewUpdateAgentRequest instantiates a new UpdateAgentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAgentRequestWithDefaults

`func NewUpdateAgentRequestWithDefaults() *UpdateAgentRequest`

NewUpdateAgentRequestWithDefaults instantiates a new UpdateAgentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateAgentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAgentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAgentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAgentRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScopes

`func (o *UpdateAgentRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *UpdateAgentRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *UpdateAgentRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *UpdateAgentRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetIsActive

`func (o *UpdateAgentRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateAgentRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateAgentRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UpdateAgentRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetExpiresAt

`func (o *UpdateAgentRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *UpdateAgentRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *UpdateAgentRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *UpdateAgentRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetIntentsApiEnabled

`func (o *UpdateAgentRequest) GetIntentsApiEnabled() bool`

GetIntentsApiEnabled returns the IntentsApiEnabled field if non-nil, zero value otherwise.

### GetIntentsApiEnabledOk

`func (o *UpdateAgentRequest) GetIntentsApiEnabledOk() (*bool, bool)`

GetIntentsApiEnabledOk returns a tuple with the IntentsApiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentsApiEnabled

`func (o *UpdateAgentRequest) SetIntentsApiEnabled(v bool)`

SetIntentsApiEnabled sets IntentsApiEnabled field to given value.

### HasIntentsApiEnabled

`func (o *UpdateAgentRequest) HasIntentsApiEnabled() bool`

HasIntentsApiEnabled returns a boolean if a field has been set.

### GetTxToAllowlist

`func (o *UpdateAgentRequest) GetTxToAllowlist() []string`

GetTxToAllowlist returns the TxToAllowlist field if non-nil, zero value otherwise.

### GetTxToAllowlistOk

`func (o *UpdateAgentRequest) GetTxToAllowlistOk() (*[]string, bool)`

GetTxToAllowlistOk returns a tuple with the TxToAllowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxToAllowlist

`func (o *UpdateAgentRequest) SetTxToAllowlist(v []string)`

SetTxToAllowlist sets TxToAllowlist field to given value.

### HasTxToAllowlist

`func (o *UpdateAgentRequest) HasTxToAllowlist() bool`

HasTxToAllowlist returns a boolean if a field has been set.

### GetTxMaxValueEth

`func (o *UpdateAgentRequest) GetTxMaxValueEth() string`

GetTxMaxValueEth returns the TxMaxValueEth field if non-nil, zero value otherwise.

### GetTxMaxValueEthOk

`func (o *UpdateAgentRequest) GetTxMaxValueEthOk() (*string, bool)`

GetTxMaxValueEthOk returns a tuple with the TxMaxValueEth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxMaxValueEth

`func (o *UpdateAgentRequest) SetTxMaxValueEth(v string)`

SetTxMaxValueEth sets TxMaxValueEth field to given value.

### HasTxMaxValueEth

`func (o *UpdateAgentRequest) HasTxMaxValueEth() bool`

HasTxMaxValueEth returns a boolean if a field has been set.

### GetTxDailyLimitEth

`func (o *UpdateAgentRequest) GetTxDailyLimitEth() string`

GetTxDailyLimitEth returns the TxDailyLimitEth field if non-nil, zero value otherwise.

### GetTxDailyLimitEthOk

`func (o *UpdateAgentRequest) GetTxDailyLimitEthOk() (*string, bool)`

GetTxDailyLimitEthOk returns a tuple with the TxDailyLimitEth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxDailyLimitEth

`func (o *UpdateAgentRequest) SetTxDailyLimitEth(v string)`

SetTxDailyLimitEth sets TxDailyLimitEth field to given value.

### HasTxDailyLimitEth

`func (o *UpdateAgentRequest) HasTxDailyLimitEth() bool`

HasTxDailyLimitEth returns a boolean if a field has been set.

### GetTxAllowedChains

`func (o *UpdateAgentRequest) GetTxAllowedChains() []string`

GetTxAllowedChains returns the TxAllowedChains field if non-nil, zero value otherwise.

### GetTxAllowedChainsOk

`func (o *UpdateAgentRequest) GetTxAllowedChainsOk() (*[]string, bool)`

GetTxAllowedChainsOk returns a tuple with the TxAllowedChains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxAllowedChains

`func (o *UpdateAgentRequest) SetTxAllowedChains(v []string)`

SetTxAllowedChains sets TxAllowedChains field to given value.

### HasTxAllowedChains

`func (o *UpdateAgentRequest) HasTxAllowedChains() bool`

HasTxAllowedChains returns a boolean if a field has been set.

### GetTokenTtlSeconds

`func (o *UpdateAgentRequest) GetTokenTtlSeconds() int32`

GetTokenTtlSeconds returns the TokenTtlSeconds field if non-nil, zero value otherwise.

### GetTokenTtlSecondsOk

`func (o *UpdateAgentRequest) GetTokenTtlSecondsOk() (*int32, bool)`

GetTokenTtlSecondsOk returns a tuple with the TokenTtlSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTtlSeconds

`func (o *UpdateAgentRequest) SetTokenTtlSeconds(v int32)`

SetTokenTtlSeconds sets TokenTtlSeconds field to given value.

### HasTokenTtlSeconds

`func (o *UpdateAgentRequest) HasTokenTtlSeconds() bool`

HasTokenTtlSeconds returns a boolean if a field has been set.

### GetVaultIds

`func (o *UpdateAgentRequest) GetVaultIds() []string`

GetVaultIds returns the VaultIds field if non-nil, zero value otherwise.

### GetVaultIdsOk

`func (o *UpdateAgentRequest) GetVaultIdsOk() (*[]string, bool)`

GetVaultIdsOk returns a tuple with the VaultIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultIds

`func (o *UpdateAgentRequest) SetVaultIds(v []string)`

SetVaultIds sets VaultIds field to given value.

### HasVaultIds

`func (o *UpdateAgentRequest) HasVaultIds() bool`

HasVaultIds returns a boolean if a field has been set.

### GetShroudEnabled

`func (o *UpdateAgentRequest) GetShroudEnabled() bool`

GetShroudEnabled returns the ShroudEnabled field if non-nil, zero value otherwise.

### GetShroudEnabledOk

`func (o *UpdateAgentRequest) GetShroudEnabledOk() (*bool, bool)`

GetShroudEnabledOk returns a tuple with the ShroudEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShroudEnabled

`func (o *UpdateAgentRequest) SetShroudEnabled(v bool)`

SetShroudEnabled sets ShroudEnabled field to given value.

### HasShroudEnabled

`func (o *UpdateAgentRequest) HasShroudEnabled() bool`

HasShroudEnabled returns a boolean if a field has been set.

### GetShroudConfig

`func (o *UpdateAgentRequest) GetShroudConfig() ShroudConfig`

GetShroudConfig returns the ShroudConfig field if non-nil, zero value otherwise.

### GetShroudConfigOk

`func (o *UpdateAgentRequest) GetShroudConfigOk() (*ShroudConfig, bool)`

GetShroudConfigOk returns a tuple with the ShroudConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShroudConfig

`func (o *UpdateAgentRequest) SetShroudConfig(v ShroudConfig)`

SetShroudConfig sets ShroudConfig field to given value.

### HasShroudConfig

`func (o *UpdateAgentRequest) HasShroudConfig() bool`

HasShroudConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


