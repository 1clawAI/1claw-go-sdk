# AgentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**AuthMethod** | **string** |  | 
**Scopes** | Pointer to **[]string** |  | [optional] 
**IsActive** | **bool** |  | 
**IntentsApiEnabled** | **bool** |  | 
**TxToAllowlist** | Pointer to **[]string** |  | [optional] 
**TxMaxValueEth** | Pointer to **string** |  | [optional] 
**TxDailyLimitEth** | Pointer to **string** |  | [optional] 
**TxAllowedChains** | Pointer to **[]string** |  | [optional] 
**TokenTtlSeconds** | Pointer to **int32** |  | [optional] 
**VaultIds** | Pointer to **[]string** |  | [optional] 
**ClientCertFingerprint** | Pointer to **string** | SHA-256 fingerprint of the client certificate (mTLS agents) | [optional] 
**OidcIssuer** | Pointer to **string** | OIDC issuer URL (oidc_client_credentials agents) | [optional] 
**OidcClientId** | Pointer to **string** | OIDC client ID (oidc_client_credentials agents) | [optional] 
**SshPublicKey** | Pointer to **string** | Ed25519 SSH public key (base64-encoded, auto-generated at creation) | [optional] 
**EcdhPublicKey** | Pointer to **string** | P-256 ECDH public key (base64 SEC1 uncompressed point, auto-generated at creation) | [optional] 
**ShroudEnabled** | **bool** | Whether this agent routes LLM traffic through the Shroud TEE proxy | 
**ShroudConfig** | Pointer to [**ShroudConfig**](ShroudConfig.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**LastActiveAt** | Pointer to **time.Time** |  | [optional] 
**SmartAccounts** | Pointer to [**[]AgentSmartAccountResponse**](AgentSmartAccountResponse.md) | Multi-chain; one Safe per chain | [optional] 

## Methods

### NewAgentResponse

`func NewAgentResponse(id string, name string, authMethod string, isActive bool, intentsApiEnabled bool, shroudEnabled bool, createdAt time.Time, ) *AgentResponse`

NewAgentResponse instantiates a new AgentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentResponseWithDefaults

`func NewAgentResponseWithDefaults() *AgentResponse`

NewAgentResponseWithDefaults instantiates a new AgentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AgentResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AgentResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAuthMethod

`func (o *AgentResponse) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *AgentResponse) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *AgentResponse) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.


### GetScopes

`func (o *AgentResponse) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AgentResponse) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AgentResponse) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AgentResponse) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetIsActive

`func (o *AgentResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AgentResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AgentResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetIntentsApiEnabled

`func (o *AgentResponse) GetIntentsApiEnabled() bool`

GetIntentsApiEnabled returns the IntentsApiEnabled field if non-nil, zero value otherwise.

### GetIntentsApiEnabledOk

`func (o *AgentResponse) GetIntentsApiEnabledOk() (*bool, bool)`

GetIntentsApiEnabledOk returns a tuple with the IntentsApiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentsApiEnabled

`func (o *AgentResponse) SetIntentsApiEnabled(v bool)`

SetIntentsApiEnabled sets IntentsApiEnabled field to given value.


### GetTxToAllowlist

`func (o *AgentResponse) GetTxToAllowlist() []string`

GetTxToAllowlist returns the TxToAllowlist field if non-nil, zero value otherwise.

### GetTxToAllowlistOk

`func (o *AgentResponse) GetTxToAllowlistOk() (*[]string, bool)`

GetTxToAllowlistOk returns a tuple with the TxToAllowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxToAllowlist

`func (o *AgentResponse) SetTxToAllowlist(v []string)`

SetTxToAllowlist sets TxToAllowlist field to given value.

### HasTxToAllowlist

`func (o *AgentResponse) HasTxToAllowlist() bool`

HasTxToAllowlist returns a boolean if a field has been set.

### GetTxMaxValueEth

`func (o *AgentResponse) GetTxMaxValueEth() string`

GetTxMaxValueEth returns the TxMaxValueEth field if non-nil, zero value otherwise.

### GetTxMaxValueEthOk

`func (o *AgentResponse) GetTxMaxValueEthOk() (*string, bool)`

GetTxMaxValueEthOk returns a tuple with the TxMaxValueEth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxMaxValueEth

`func (o *AgentResponse) SetTxMaxValueEth(v string)`

SetTxMaxValueEth sets TxMaxValueEth field to given value.

### HasTxMaxValueEth

`func (o *AgentResponse) HasTxMaxValueEth() bool`

HasTxMaxValueEth returns a boolean if a field has been set.

### GetTxDailyLimitEth

`func (o *AgentResponse) GetTxDailyLimitEth() string`

GetTxDailyLimitEth returns the TxDailyLimitEth field if non-nil, zero value otherwise.

### GetTxDailyLimitEthOk

`func (o *AgentResponse) GetTxDailyLimitEthOk() (*string, bool)`

GetTxDailyLimitEthOk returns a tuple with the TxDailyLimitEth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxDailyLimitEth

`func (o *AgentResponse) SetTxDailyLimitEth(v string)`

SetTxDailyLimitEth sets TxDailyLimitEth field to given value.

### HasTxDailyLimitEth

`func (o *AgentResponse) HasTxDailyLimitEth() bool`

HasTxDailyLimitEth returns a boolean if a field has been set.

### GetTxAllowedChains

`func (o *AgentResponse) GetTxAllowedChains() []string`

GetTxAllowedChains returns the TxAllowedChains field if non-nil, zero value otherwise.

### GetTxAllowedChainsOk

`func (o *AgentResponse) GetTxAllowedChainsOk() (*[]string, bool)`

GetTxAllowedChainsOk returns a tuple with the TxAllowedChains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxAllowedChains

`func (o *AgentResponse) SetTxAllowedChains(v []string)`

SetTxAllowedChains sets TxAllowedChains field to given value.

### HasTxAllowedChains

`func (o *AgentResponse) HasTxAllowedChains() bool`

HasTxAllowedChains returns a boolean if a field has been set.

### GetTokenTtlSeconds

`func (o *AgentResponse) GetTokenTtlSeconds() int32`

GetTokenTtlSeconds returns the TokenTtlSeconds field if non-nil, zero value otherwise.

### GetTokenTtlSecondsOk

`func (o *AgentResponse) GetTokenTtlSecondsOk() (*int32, bool)`

GetTokenTtlSecondsOk returns a tuple with the TokenTtlSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTtlSeconds

`func (o *AgentResponse) SetTokenTtlSeconds(v int32)`

SetTokenTtlSeconds sets TokenTtlSeconds field to given value.

### HasTokenTtlSeconds

`func (o *AgentResponse) HasTokenTtlSeconds() bool`

HasTokenTtlSeconds returns a boolean if a field has been set.

### GetVaultIds

`func (o *AgentResponse) GetVaultIds() []string`

GetVaultIds returns the VaultIds field if non-nil, zero value otherwise.

### GetVaultIdsOk

`func (o *AgentResponse) GetVaultIdsOk() (*[]string, bool)`

GetVaultIdsOk returns a tuple with the VaultIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultIds

`func (o *AgentResponse) SetVaultIds(v []string)`

SetVaultIds sets VaultIds field to given value.

### HasVaultIds

`func (o *AgentResponse) HasVaultIds() bool`

HasVaultIds returns a boolean if a field has been set.

### GetClientCertFingerprint

`func (o *AgentResponse) GetClientCertFingerprint() string`

GetClientCertFingerprint returns the ClientCertFingerprint field if non-nil, zero value otherwise.

### GetClientCertFingerprintOk

`func (o *AgentResponse) GetClientCertFingerprintOk() (*string, bool)`

GetClientCertFingerprintOk returns a tuple with the ClientCertFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertFingerprint

`func (o *AgentResponse) SetClientCertFingerprint(v string)`

SetClientCertFingerprint sets ClientCertFingerprint field to given value.

### HasClientCertFingerprint

`func (o *AgentResponse) HasClientCertFingerprint() bool`

HasClientCertFingerprint returns a boolean if a field has been set.

### GetOidcIssuer

`func (o *AgentResponse) GetOidcIssuer() string`

GetOidcIssuer returns the OidcIssuer field if non-nil, zero value otherwise.

### GetOidcIssuerOk

`func (o *AgentResponse) GetOidcIssuerOk() (*string, bool)`

GetOidcIssuerOk returns a tuple with the OidcIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcIssuer

`func (o *AgentResponse) SetOidcIssuer(v string)`

SetOidcIssuer sets OidcIssuer field to given value.

### HasOidcIssuer

`func (o *AgentResponse) HasOidcIssuer() bool`

HasOidcIssuer returns a boolean if a field has been set.

### GetOidcClientId

`func (o *AgentResponse) GetOidcClientId() string`

GetOidcClientId returns the OidcClientId field if non-nil, zero value otherwise.

### GetOidcClientIdOk

`func (o *AgentResponse) GetOidcClientIdOk() (*string, bool)`

GetOidcClientIdOk returns a tuple with the OidcClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientId

`func (o *AgentResponse) SetOidcClientId(v string)`

SetOidcClientId sets OidcClientId field to given value.

### HasOidcClientId

`func (o *AgentResponse) HasOidcClientId() bool`

HasOidcClientId returns a boolean if a field has been set.

### GetSshPublicKey

`func (o *AgentResponse) GetSshPublicKey() string`

GetSshPublicKey returns the SshPublicKey field if non-nil, zero value otherwise.

### GetSshPublicKeyOk

`func (o *AgentResponse) GetSshPublicKeyOk() (*string, bool)`

GetSshPublicKeyOk returns a tuple with the SshPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPublicKey

`func (o *AgentResponse) SetSshPublicKey(v string)`

SetSshPublicKey sets SshPublicKey field to given value.

### HasSshPublicKey

`func (o *AgentResponse) HasSshPublicKey() bool`

HasSshPublicKey returns a boolean if a field has been set.

### GetEcdhPublicKey

`func (o *AgentResponse) GetEcdhPublicKey() string`

GetEcdhPublicKey returns the EcdhPublicKey field if non-nil, zero value otherwise.

### GetEcdhPublicKeyOk

`func (o *AgentResponse) GetEcdhPublicKeyOk() (*string, bool)`

GetEcdhPublicKeyOk returns a tuple with the EcdhPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdhPublicKey

`func (o *AgentResponse) SetEcdhPublicKey(v string)`

SetEcdhPublicKey sets EcdhPublicKey field to given value.

### HasEcdhPublicKey

`func (o *AgentResponse) HasEcdhPublicKey() bool`

HasEcdhPublicKey returns a boolean if a field has been set.

### GetShroudEnabled

`func (o *AgentResponse) GetShroudEnabled() bool`

GetShroudEnabled returns the ShroudEnabled field if non-nil, zero value otherwise.

### GetShroudEnabledOk

`func (o *AgentResponse) GetShroudEnabledOk() (*bool, bool)`

GetShroudEnabledOk returns a tuple with the ShroudEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShroudEnabled

`func (o *AgentResponse) SetShroudEnabled(v bool)`

SetShroudEnabled sets ShroudEnabled field to given value.


### GetShroudConfig

`func (o *AgentResponse) GetShroudConfig() ShroudConfig`

GetShroudConfig returns the ShroudConfig field if non-nil, zero value otherwise.

### GetShroudConfigOk

`func (o *AgentResponse) GetShroudConfigOk() (*ShroudConfig, bool)`

GetShroudConfigOk returns a tuple with the ShroudConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShroudConfig

`func (o *AgentResponse) SetShroudConfig(v ShroudConfig)`

SetShroudConfig sets ShroudConfig field to given value.

### HasShroudConfig

`func (o *AgentResponse) HasShroudConfig() bool`

HasShroudConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AgentResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiresAt

`func (o *AgentResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AgentResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AgentResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *AgentResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetLastActiveAt

`func (o *AgentResponse) GetLastActiveAt() time.Time`

GetLastActiveAt returns the LastActiveAt field if non-nil, zero value otherwise.

### GetLastActiveAtOk

`func (o *AgentResponse) GetLastActiveAtOk() (*time.Time, bool)`

GetLastActiveAtOk returns a tuple with the LastActiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveAt

`func (o *AgentResponse) SetLastActiveAt(v time.Time)`

SetLastActiveAt sets LastActiveAt field to given value.

### HasLastActiveAt

`func (o *AgentResponse) HasLastActiveAt() bool`

HasLastActiveAt returns a boolean if a field has been set.

### GetSmartAccounts

`func (o *AgentResponse) GetSmartAccounts() []AgentSmartAccountResponse`

GetSmartAccounts returns the SmartAccounts field if non-nil, zero value otherwise.

### GetSmartAccountsOk

`func (o *AgentResponse) GetSmartAccountsOk() (*[]AgentSmartAccountResponse, bool)`

GetSmartAccountsOk returns a tuple with the SmartAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartAccounts

`func (o *AgentResponse) SetSmartAccounts(v []AgentSmartAccountResponse)`

SetSmartAccounts sets SmartAccounts field to given value.

### HasSmartAccounts

`func (o *AgentResponse) HasSmartAccounts() bool`

HasSmartAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


