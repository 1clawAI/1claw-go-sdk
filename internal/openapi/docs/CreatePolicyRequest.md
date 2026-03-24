# CreatePolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretPathPattern** | **string** |  | 
**PrincipalType** | **string** |  | 
**PrincipalId** | **string** |  | 
**Permissions** | **[]string** |  | 
**Conditions** | Pointer to **map[string]interface{}** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCreatePolicyRequest

`func NewCreatePolicyRequest(secretPathPattern string, principalType string, principalId string, permissions []string, ) *CreatePolicyRequest`

NewCreatePolicyRequest instantiates a new CreatePolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePolicyRequestWithDefaults

`func NewCreatePolicyRequestWithDefaults() *CreatePolicyRequest`

NewCreatePolicyRequestWithDefaults instantiates a new CreatePolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretPathPattern

`func (o *CreatePolicyRequest) GetSecretPathPattern() string`

GetSecretPathPattern returns the SecretPathPattern field if non-nil, zero value otherwise.

### GetSecretPathPatternOk

`func (o *CreatePolicyRequest) GetSecretPathPatternOk() (*string, bool)`

GetSecretPathPatternOk returns a tuple with the SecretPathPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretPathPattern

`func (o *CreatePolicyRequest) SetSecretPathPattern(v string)`

SetSecretPathPattern sets SecretPathPattern field to given value.


### GetPrincipalType

`func (o *CreatePolicyRequest) GetPrincipalType() string`

GetPrincipalType returns the PrincipalType field if non-nil, zero value otherwise.

### GetPrincipalTypeOk

`func (o *CreatePolicyRequest) GetPrincipalTypeOk() (*string, bool)`

GetPrincipalTypeOk returns a tuple with the PrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalType

`func (o *CreatePolicyRequest) SetPrincipalType(v string)`

SetPrincipalType sets PrincipalType field to given value.


### GetPrincipalId

`func (o *CreatePolicyRequest) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *CreatePolicyRequest) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *CreatePolicyRequest) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.


### GetPermissions

`func (o *CreatePolicyRequest) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreatePolicyRequest) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreatePolicyRequest) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetConditions

`func (o *CreatePolicyRequest) GetConditions() map[string]interface{}`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *CreatePolicyRequest) GetConditionsOk() (*map[string]interface{}, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *CreatePolicyRequest) SetConditions(v map[string]interface{})`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *CreatePolicyRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CreatePolicyRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreatePolicyRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreatePolicyRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreatePolicyRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


