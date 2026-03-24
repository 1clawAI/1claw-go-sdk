# PolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**VaultId** | **string** |  | 
**SecretPathPattern** | **string** |  | 
**PrincipalType** | **string** |  | 
**PrincipalId** | **string** |  | 
**Permissions** | **[]string** |  | 
**Conditions** | Pointer to **map[string]interface{}** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedByType** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewPolicyResponse

`func NewPolicyResponse(id string, vaultId string, secretPathPattern string, principalType string, principalId string, permissions []string, createdAt time.Time, ) *PolicyResponse`

NewPolicyResponse instantiates a new PolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyResponseWithDefaults

`func NewPolicyResponseWithDefaults() *PolicyResponse`

NewPolicyResponseWithDefaults instantiates a new PolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PolicyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetVaultId

`func (o *PolicyResponse) GetVaultId() string`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *PolicyResponse) GetVaultIdOk() (*string, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *PolicyResponse) SetVaultId(v string)`

SetVaultId sets VaultId field to given value.


### GetSecretPathPattern

`func (o *PolicyResponse) GetSecretPathPattern() string`

GetSecretPathPattern returns the SecretPathPattern field if non-nil, zero value otherwise.

### GetSecretPathPatternOk

`func (o *PolicyResponse) GetSecretPathPatternOk() (*string, bool)`

GetSecretPathPatternOk returns a tuple with the SecretPathPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretPathPattern

`func (o *PolicyResponse) SetSecretPathPattern(v string)`

SetSecretPathPattern sets SecretPathPattern field to given value.


### GetPrincipalType

`func (o *PolicyResponse) GetPrincipalType() string`

GetPrincipalType returns the PrincipalType field if non-nil, zero value otherwise.

### GetPrincipalTypeOk

`func (o *PolicyResponse) GetPrincipalTypeOk() (*string, bool)`

GetPrincipalTypeOk returns a tuple with the PrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalType

`func (o *PolicyResponse) SetPrincipalType(v string)`

SetPrincipalType sets PrincipalType field to given value.


### GetPrincipalId

`func (o *PolicyResponse) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *PolicyResponse) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *PolicyResponse) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.


### GetPermissions

`func (o *PolicyResponse) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PolicyResponse) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PolicyResponse) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetConditions

`func (o *PolicyResponse) GetConditions() map[string]interface{}`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *PolicyResponse) GetConditionsOk() (*map[string]interface{}, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *PolicyResponse) SetConditions(v map[string]interface{})`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *PolicyResponse) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetExpiresAt

`func (o *PolicyResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PolicyResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PolicyResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PolicyResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PolicyResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PolicyResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PolicyResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PolicyResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByType

`func (o *PolicyResponse) GetCreatedByType() string`

GetCreatedByType returns the CreatedByType field if non-nil, zero value otherwise.

### GetCreatedByTypeOk

`func (o *PolicyResponse) GetCreatedByTypeOk() (*string, bool)`

GetCreatedByTypeOk returns a tuple with the CreatedByType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByType

`func (o *PolicyResponse) SetCreatedByType(v string)`

SetCreatedByType sets CreatedByType field to given value.

### HasCreatedByType

`func (o *PolicyResponse) HasCreatedByType() bool`

HasCreatedByType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PolicyResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PolicyResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PolicyResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


