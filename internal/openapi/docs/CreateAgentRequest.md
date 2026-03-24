# CreateAgentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**AuthMethod** | Pointer to **string** | Authentication method. api_key generates a one-time key; mtls requires client_cert_fingerprint; oidc_client_credentials requires oidc_issuer and oidc_client_id. | [optional] [default to "api_key"]
**Scopes** | Pointer to **[]string** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**IntentsApiEnabled** | Pointer to **bool** |  | [optional] [default to false]
**TxToAllowlist** | Pointer to **[]string** |  | [optional] 
**TxMaxValueEth** | Pointer to **string** |  | [optional] 
**TxDailyLimitEth** | Pointer to **string** |  | [optional] 
**TxAllowedChains** | Pointer to **[]string** |  | [optional] 
**TokenTtlSeconds** | Pointer to **int32** | Per-agent token TTL in seconds (overrides global default) | [optional] 
**VaultIds** | Pointer to **[]string** | Restrict agent to specific vault UUIDs (empty &#x3D; all vaults in org) | [optional] 
**ClientCertFingerprint** | Pointer to **string** | SHA-256 fingerprint of the client certificate (required for mTLS auth) | [optional] 
**OidcIssuer** | Pointer to **string** | OIDC issuer URL (required for oidc_client_credentials auth) | [optional] 
**OidcClientId** | Pointer to **string** | OIDC client ID (required for oidc_client_credentials auth) | [optional] 
**ShroudEnabled** | Pointer to **bool** | Enable Shroud LLM Proxy for this agent | [optional] [default to false]
**ShroudConfig** | Pointer to [**ShroudConfig**](ShroudConfig.md) |  | [optional] 

## Methods

### NewCreateAgentRequest

`func NewCreateAgentRequest(name string, ) *CreateAgentRequest`

NewCreateAgentRequest instantiates a new CreateAgentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAgentRequestWithDefaults

`func NewCreateAgentRequestWithDefaults() *CreateAgentRequest`

NewCreateAgentRequestWithDefaults instantiates a new CreateAgentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAgentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAgentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAgentRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateAgentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAgentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAgentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAgentRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAuthMethod

`func (o *CreateAgentRequest) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *CreateAgentRequest) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *CreateAgentRequest) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *CreateAgentRequest) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### GetScopes

`func (o *CreateAgentRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CreateAgentRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CreateAgentRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *CreateAgentRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CreateAgentRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateAgentRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateAgentRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreateAgentRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetIntentsApiEnabled

`func (o *CreateAgentRequest) GetIntentsApiEnabled() bool`

GetIntentsApiEnabled returns the IntentsApiEnabled field if non-nil, zero value otherwise.

### GetIntentsApiEnabledOk

`func (o *CreateAgentRequest) GetIntentsApiEnabledOk() (*bool, bool)`

GetIntentsApiEnabledOk returns a tuple with the IntentsApiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentsApiEnabled

`func (o *CreateAgentRequest) SetIntentsApiEnabled(v bool)`

SetIntentsApiEnabled sets IntentsApiEnabled field to given value.

### HasIntentsApiEnabled

`func (o *CreateAgentRequest) HasIntentsApiEnabled() bool`

HasIntentsApiEnabled returns a boolean if a field has been set.

### GetTxToAllowlist

`func (o *CreateAgentRequest) GetTxToAllowlist() []string`

GetTxToAllowlist returns the TxToAllowlist field if non-nil, zero value otherwise.

### GetTxToAllowlistOk

`func (o *CreateAgentRequest) GetTxToAllowlistOk() (*[]string, bool)`

GetTxToAllowlistOk returns a tuple with the TxToAllowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxToAllowlist

`func (o *CreateAgentRequest) SetTxToAllowlist(v []string)`

SetTxToAllowlist sets TxToAllowlist field to given value.

### HasTxToAllowlist

`func (o *CreateAgentRequest) HasTxToAllowlist() bool`

HasTxToAllowlist returns a boolean if a field has been set.

### GetTxMaxValueEth

`func (o *CreateAgentRequest) GetTxMaxValueEth() string`

GetTxMaxValueEth returns the TxMaxValueEth field if non-nil, zero value otherwise.

### GetTxMaxValueEthOk

`func (o *CreateAgentRequest) GetTxMaxValueEthOk() (*string, bool)`

GetTxMaxValueEthOk returns a tuple with the TxMaxValueEth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxMaxValueEth

`func (o *CreateAgentRequest) SetTxMaxValueEth(v string)`

SetTxMaxValueEth sets TxMaxValueEth field to given value.

### HasTxMaxValueEth

`func (o *CreateAgentRequest) HasTxMaxValueEth() bool`

HasTxMaxValueEth returns a boolean if a field has been set.

### GetTxDailyLimitEth

`func (o *CreateAgentRequest) GetTxDailyLimitEth() string`

GetTxDailyLimitEth returns the TxDailyLimitEth field if non-nil, zero value otherwise.

### GetTxDailyLimitEthOk

`func (o *CreateAgentRequest) GetTxDailyLimitEthOk() (*string, bool)`

GetTxDailyLimitEthOk returns a tuple with the TxDailyLimitEth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxDailyLimitEth

`func (o *CreateAgentRequest) SetTxDailyLimitEth(v string)`

SetTxDailyLimitEth sets TxDailyLimitEth field to given value.

### HasTxDailyLimitEth

`func (o *CreateAgentRequest) HasTxDailyLimitEth() bool`

HasTxDailyLimitEth returns a boolean if a field has been set.

### GetTxAllowedChains

`func (o *CreateAgentRequest) GetTxAllowedChains() []string`

GetTxAllowedChains returns the TxAllowedChains field if non-nil, zero value otherwise.

### GetTxAllowedChainsOk

`func (o *CreateAgentRequest) GetTxAllowedChainsOk() (*[]string, bool)`

GetTxAllowedChainsOk returns a tuple with the TxAllowedChains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxAllowedChains

`func (o *CreateAgentRequest) SetTxAllowedChains(v []string)`

SetTxAllowedChains sets TxAllowedChains field to given value.

### HasTxAllowedChains

`func (o *CreateAgentRequest) HasTxAllowedChains() bool`

HasTxAllowedChains returns a boolean if a field has been set.

### GetTokenTtlSeconds

`func (o *CreateAgentRequest) GetTokenTtlSeconds() int32`

GetTokenTtlSeconds returns the TokenTtlSeconds field if non-nil, zero value otherwise.

### GetTokenTtlSecondsOk

`func (o *CreateAgentRequest) GetTokenTtlSecondsOk() (*int32, bool)`

GetTokenTtlSecondsOk returns a tuple with the TokenTtlSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTtlSeconds

`func (o *CreateAgentRequest) SetTokenTtlSeconds(v int32)`

SetTokenTtlSeconds sets TokenTtlSeconds field to given value.

### HasTokenTtlSeconds

`func (o *CreateAgentRequest) HasTokenTtlSeconds() bool`

HasTokenTtlSeconds returns a boolean if a field has been set.

### GetVaultIds

`func (o *CreateAgentRequest) GetVaultIds() []string`

GetVaultIds returns the VaultIds field if non-nil, zero value otherwise.

### GetVaultIdsOk

`func (o *CreateAgentRequest) GetVaultIdsOk() (*[]string, bool)`

GetVaultIdsOk returns a tuple with the VaultIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultIds

`func (o *CreateAgentRequest) SetVaultIds(v []string)`

SetVaultIds sets VaultIds field to given value.

### HasVaultIds

`func (o *CreateAgentRequest) HasVaultIds() bool`

HasVaultIds returns a boolean if a field has been set.

### GetClientCertFingerprint

`func (o *CreateAgentRequest) GetClientCertFingerprint() string`

GetClientCertFingerprint returns the ClientCertFingerprint field if non-nil, zero value otherwise.

### GetClientCertFingerprintOk

`func (o *CreateAgentRequest) GetClientCertFingerprintOk() (*string, bool)`

GetClientCertFingerprintOk returns a tuple with the ClientCertFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertFingerprint

`func (o *CreateAgentRequest) SetClientCertFingerprint(v string)`

SetClientCertFingerprint sets ClientCertFingerprint field to given value.

### HasClientCertFingerprint

`func (o *CreateAgentRequest) HasClientCertFingerprint() bool`

HasClientCertFingerprint returns a boolean if a field has been set.

### GetOidcIssuer

`func (o *CreateAgentRequest) GetOidcIssuer() string`

GetOidcIssuer returns the OidcIssuer field if non-nil, zero value otherwise.

### GetOidcIssuerOk

`func (o *CreateAgentRequest) GetOidcIssuerOk() (*string, bool)`

GetOidcIssuerOk returns a tuple with the OidcIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcIssuer

`func (o *CreateAgentRequest) SetOidcIssuer(v string)`

SetOidcIssuer sets OidcIssuer field to given value.

### HasOidcIssuer

`func (o *CreateAgentRequest) HasOidcIssuer() bool`

HasOidcIssuer returns a boolean if a field has been set.

### GetOidcClientId

`func (o *CreateAgentRequest) GetOidcClientId() string`

GetOidcClientId returns the OidcClientId field if non-nil, zero value otherwise.

### GetOidcClientIdOk

`func (o *CreateAgentRequest) GetOidcClientIdOk() (*string, bool)`

GetOidcClientIdOk returns a tuple with the OidcClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientId

`func (o *CreateAgentRequest) SetOidcClientId(v string)`

SetOidcClientId sets OidcClientId field to given value.

### HasOidcClientId

`func (o *CreateAgentRequest) HasOidcClientId() bool`

HasOidcClientId returns a boolean if a field has been set.

### GetShroudEnabled

`func (o *CreateAgentRequest) GetShroudEnabled() bool`

GetShroudEnabled returns the ShroudEnabled field if non-nil, zero value otherwise.

### GetShroudEnabledOk

`func (o *CreateAgentRequest) GetShroudEnabledOk() (*bool, bool)`

GetShroudEnabledOk returns a tuple with the ShroudEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShroudEnabled

`func (o *CreateAgentRequest) SetShroudEnabled(v bool)`

SetShroudEnabled sets ShroudEnabled field to given value.

### HasShroudEnabled

`func (o *CreateAgentRequest) HasShroudEnabled() bool`

HasShroudEnabled returns a boolean if a field has been set.

### GetShroudConfig

`func (o *CreateAgentRequest) GetShroudConfig() ShroudConfig`

GetShroudConfig returns the ShroudConfig field if non-nil, zero value otherwise.

### GetShroudConfigOk

`func (o *CreateAgentRequest) GetShroudConfigOk() (*ShroudConfig, bool)`

GetShroudConfigOk returns a tuple with the ShroudConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShroudConfig

`func (o *CreateAgentRequest) SetShroudConfig(v ShroudConfig)`

SetShroudConfig sets ShroudConfig field to given value.

### HasShroudConfig

`func (o *CreateAgentRequest) HasShroudConfig() bool`

HasShroudConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


