# SecretResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Path** | **string** |  | 
**Type** | **string** |  | 
**Value** | **string** |  | 
**Version** | **int32** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**CmekEncrypted** | Pointer to **bool** | Whether this secret value is CMEK-encrypted (requires client-side decryption) | [optional] 

## Methods

### NewSecretResponse

`func NewSecretResponse(id string, path string, type_ string, value string, version int32, createdAt time.Time, ) *SecretResponse`

NewSecretResponse instantiates a new SecretResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretResponseWithDefaults

`func NewSecretResponseWithDefaults() *SecretResponse`

NewSecretResponseWithDefaults instantiates a new SecretResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecretResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecretResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecretResponse) SetId(v string)`

SetId sets Id field to given value.


### GetPath

`func (o *SecretResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SecretResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SecretResponse) SetPath(v string)`

SetPath sets Path field to given value.


### GetType

`func (o *SecretResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecretResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecretResponse) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *SecretResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SecretResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SecretResponse) SetValue(v string)`

SetValue sets Value field to given value.


### GetVersion

`func (o *SecretResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SecretResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SecretResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetMetadata

`func (o *SecretResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SecretResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SecretResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SecretResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedBy

`func (o *SecretResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SecretResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SecretResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SecretResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SecretResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SecretResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SecretResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiresAt

`func (o *SecretResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SecretResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SecretResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *SecretResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetCmekEncrypted

`func (o *SecretResponse) GetCmekEncrypted() bool`

GetCmekEncrypted returns the CmekEncrypted field if non-nil, zero value otherwise.

### GetCmekEncryptedOk

`func (o *SecretResponse) GetCmekEncryptedOk() (*bool, bool)`

GetCmekEncryptedOk returns a tuple with the CmekEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmekEncrypted

`func (o *SecretResponse) SetCmekEncrypted(v bool)`

SetCmekEncrypted sets CmekEncrypted field to given value.

### HasCmekEncrypted

`func (o *SecretResponse) HasCmekEncrypted() bool`

HasCmekEncrypted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


