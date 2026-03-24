# AgentSelfResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**IntentsApiEnabled** | Pointer to **bool** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**LastActiveAt** | Pointer to **time.Time** |  | [optional] 
**SshPublicKey** | Pointer to **string** | Ed25519 SSH public key (base64-encoded) | [optional] 
**EcdhPublicKey** | Pointer to **string** | P-256 ECDH public key (base64 SEC1 uncompressed point) | [optional] 
**ShroudEnabled** | Pointer to **bool** | Whether this agent routes LLM traffic through the Shroud TEE proxy | [optional] 
**ShroudConfig** | Pointer to [**ShroudConfig**](ShroudConfig.md) |  | [optional] 

## Methods

### NewAgentSelfResponse

`func NewAgentSelfResponse() *AgentSelfResponse`

NewAgentSelfResponse instantiates a new AgentSelfResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentSelfResponseWithDefaults

`func NewAgentSelfResponseWithDefaults() *AgentSelfResponse`

NewAgentSelfResponseWithDefaults instantiates a new AgentSelfResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentSelfResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentSelfResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentSelfResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentSelfResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AgentSelfResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentSelfResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentSelfResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentSelfResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AgentSelfResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentSelfResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentSelfResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentSelfResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrgId

`func (o *AgentSelfResponse) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AgentSelfResponse) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AgentSelfResponse) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *AgentSelfResponse) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetScopes

`func (o *AgentSelfResponse) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AgentSelfResponse) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AgentSelfResponse) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AgentSelfResponse) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetIsActive

`func (o *AgentSelfResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AgentSelfResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AgentSelfResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *AgentSelfResponse) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIntentsApiEnabled

`func (o *AgentSelfResponse) GetIntentsApiEnabled() bool`

GetIntentsApiEnabled returns the IntentsApiEnabled field if non-nil, zero value otherwise.

### GetIntentsApiEnabledOk

`func (o *AgentSelfResponse) GetIntentsApiEnabledOk() (*bool, bool)`

GetIntentsApiEnabledOk returns a tuple with the IntentsApiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentsApiEnabled

`func (o *AgentSelfResponse) SetIntentsApiEnabled(v bool)`

SetIntentsApiEnabled sets IntentsApiEnabled field to given value.

### HasIntentsApiEnabled

`func (o *AgentSelfResponse) HasIntentsApiEnabled() bool`

HasIntentsApiEnabled returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AgentSelfResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AgentSelfResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AgentSelfResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AgentSelfResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AgentSelfResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentSelfResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentSelfResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AgentSelfResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *AgentSelfResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AgentSelfResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AgentSelfResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *AgentSelfResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetLastActiveAt

`func (o *AgentSelfResponse) GetLastActiveAt() time.Time`

GetLastActiveAt returns the LastActiveAt field if non-nil, zero value otherwise.

### GetLastActiveAtOk

`func (o *AgentSelfResponse) GetLastActiveAtOk() (*time.Time, bool)`

GetLastActiveAtOk returns a tuple with the LastActiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveAt

`func (o *AgentSelfResponse) SetLastActiveAt(v time.Time)`

SetLastActiveAt sets LastActiveAt field to given value.

### HasLastActiveAt

`func (o *AgentSelfResponse) HasLastActiveAt() bool`

HasLastActiveAt returns a boolean if a field has been set.

### GetSshPublicKey

`func (o *AgentSelfResponse) GetSshPublicKey() string`

GetSshPublicKey returns the SshPublicKey field if non-nil, zero value otherwise.

### GetSshPublicKeyOk

`func (o *AgentSelfResponse) GetSshPublicKeyOk() (*string, bool)`

GetSshPublicKeyOk returns a tuple with the SshPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPublicKey

`func (o *AgentSelfResponse) SetSshPublicKey(v string)`

SetSshPublicKey sets SshPublicKey field to given value.

### HasSshPublicKey

`func (o *AgentSelfResponse) HasSshPublicKey() bool`

HasSshPublicKey returns a boolean if a field has been set.

### GetEcdhPublicKey

`func (o *AgentSelfResponse) GetEcdhPublicKey() string`

GetEcdhPublicKey returns the EcdhPublicKey field if non-nil, zero value otherwise.

### GetEcdhPublicKeyOk

`func (o *AgentSelfResponse) GetEcdhPublicKeyOk() (*string, bool)`

GetEcdhPublicKeyOk returns a tuple with the EcdhPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdhPublicKey

`func (o *AgentSelfResponse) SetEcdhPublicKey(v string)`

SetEcdhPublicKey sets EcdhPublicKey field to given value.

### HasEcdhPublicKey

`func (o *AgentSelfResponse) HasEcdhPublicKey() bool`

HasEcdhPublicKey returns a boolean if a field has been set.

### GetShroudEnabled

`func (o *AgentSelfResponse) GetShroudEnabled() bool`

GetShroudEnabled returns the ShroudEnabled field if non-nil, zero value otherwise.

### GetShroudEnabledOk

`func (o *AgentSelfResponse) GetShroudEnabledOk() (*bool, bool)`

GetShroudEnabledOk returns a tuple with the ShroudEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShroudEnabled

`func (o *AgentSelfResponse) SetShroudEnabled(v bool)`

SetShroudEnabled sets ShroudEnabled field to given value.

### HasShroudEnabled

`func (o *AgentSelfResponse) HasShroudEnabled() bool`

HasShroudEnabled returns a boolean if a field has been set.

### GetShroudConfig

`func (o *AgentSelfResponse) GetShroudConfig() ShroudConfig`

GetShroudConfig returns the ShroudConfig field if non-nil, zero value otherwise.

### GetShroudConfigOk

`func (o *AgentSelfResponse) GetShroudConfigOk() (*ShroudConfig, bool)`

GetShroudConfigOk returns a tuple with the ShroudConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShroudConfig

`func (o *AgentSelfResponse) SetShroudConfig(v ShroudConfig)`

SetShroudConfig sets ShroudConfig field to given value.

### HasShroudConfig

`func (o *AgentSelfResponse) HasShroudConfig() bool`

HasShroudConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


