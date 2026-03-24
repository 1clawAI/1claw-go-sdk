# PutSecretRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Secret type (generic, password, api_key, certificate, private_key, ssh_key, env) | [optional] [default to "generic"]
**Value** | **string** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**RotationPolicy** | Pointer to **map[string]interface{}** |  | [optional] 
**MaxAccessCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewPutSecretRequest

`func NewPutSecretRequest(value string, ) *PutSecretRequest`

NewPutSecretRequest instantiates a new PutSecretRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutSecretRequestWithDefaults

`func NewPutSecretRequestWithDefaults() *PutSecretRequest`

NewPutSecretRequestWithDefaults instantiates a new PutSecretRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PutSecretRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PutSecretRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PutSecretRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PutSecretRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *PutSecretRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PutSecretRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PutSecretRequest) SetValue(v string)`

SetValue sets Value field to given value.


### GetMetadata

`func (o *PutSecretRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PutSecretRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PutSecretRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PutSecretRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetExpiresAt

`func (o *PutSecretRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PutSecretRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PutSecretRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PutSecretRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetRotationPolicy

`func (o *PutSecretRequest) GetRotationPolicy() map[string]interface{}`

GetRotationPolicy returns the RotationPolicy field if non-nil, zero value otherwise.

### GetRotationPolicyOk

`func (o *PutSecretRequest) GetRotationPolicyOk() (*map[string]interface{}, bool)`

GetRotationPolicyOk returns a tuple with the RotationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationPolicy

`func (o *PutSecretRequest) SetRotationPolicy(v map[string]interface{})`

SetRotationPolicy sets RotationPolicy field to given value.

### HasRotationPolicy

`func (o *PutSecretRequest) HasRotationPolicy() bool`

HasRotationPolicy returns a boolean if a field has been set.

### GetMaxAccessCount

`func (o *PutSecretRequest) GetMaxAccessCount() int32`

GetMaxAccessCount returns the MaxAccessCount field if non-nil, zero value otherwise.

### GetMaxAccessCountOk

`func (o *PutSecretRequest) GetMaxAccessCountOk() (*int32, bool)`

GetMaxAccessCountOk returns a tuple with the MaxAccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAccessCount

`func (o *PutSecretRequest) SetMaxAccessCount(v int32)`

SetMaxAccessCount sets MaxAccessCount field to given value.

### HasMaxAccessCount

`func (o *PutSecretRequest) HasMaxAccessCount() bool`

HasMaxAccessCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


