# ShareListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**SecretPath** | Pointer to **string** |  | [optional] 
**RecipientType** | Pointer to **string** |  | [optional] 
**RecipientEmail** | Pointer to **string** |  | [optional] 
**AccessCount** | Pointer to **int32** |  | [optional] 
**MaxAccessCount** | Pointer to **int32** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**IsExpired** | Pointer to **bool** |  | [optional] 
**IsAccepted** | Pointer to **bool** |  | [optional] 

## Methods

### NewShareListItem

`func NewShareListItem() *ShareListItem`

NewShareListItem instantiates a new ShareListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareListItemWithDefaults

`func NewShareListItemWithDefaults() *ShareListItem`

NewShareListItemWithDefaults instantiates a new ShareListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShareListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShareListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShareListItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShareListItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSecretPath

`func (o *ShareListItem) GetSecretPath() string`

GetSecretPath returns the SecretPath field if non-nil, zero value otherwise.

### GetSecretPathOk

`func (o *ShareListItem) GetSecretPathOk() (*string, bool)`

GetSecretPathOk returns a tuple with the SecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretPath

`func (o *ShareListItem) SetSecretPath(v string)`

SetSecretPath sets SecretPath field to given value.

### HasSecretPath

`func (o *ShareListItem) HasSecretPath() bool`

HasSecretPath returns a boolean if a field has been set.

### GetRecipientType

`func (o *ShareListItem) GetRecipientType() string`

GetRecipientType returns the RecipientType field if non-nil, zero value otherwise.

### GetRecipientTypeOk

`func (o *ShareListItem) GetRecipientTypeOk() (*string, bool)`

GetRecipientTypeOk returns a tuple with the RecipientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientType

`func (o *ShareListItem) SetRecipientType(v string)`

SetRecipientType sets RecipientType field to given value.

### HasRecipientType

`func (o *ShareListItem) HasRecipientType() bool`

HasRecipientType returns a boolean if a field has been set.

### GetRecipientEmail

`func (o *ShareListItem) GetRecipientEmail() string`

GetRecipientEmail returns the RecipientEmail field if non-nil, zero value otherwise.

### GetRecipientEmailOk

`func (o *ShareListItem) GetRecipientEmailOk() (*string, bool)`

GetRecipientEmailOk returns a tuple with the RecipientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientEmail

`func (o *ShareListItem) SetRecipientEmail(v string)`

SetRecipientEmail sets RecipientEmail field to given value.

### HasRecipientEmail

`func (o *ShareListItem) HasRecipientEmail() bool`

HasRecipientEmail returns a boolean if a field has been set.

### GetAccessCount

`func (o *ShareListItem) GetAccessCount() int32`

GetAccessCount returns the AccessCount field if non-nil, zero value otherwise.

### GetAccessCountOk

`func (o *ShareListItem) GetAccessCountOk() (*int32, bool)`

GetAccessCountOk returns a tuple with the AccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCount

`func (o *ShareListItem) SetAccessCount(v int32)`

SetAccessCount sets AccessCount field to given value.

### HasAccessCount

`func (o *ShareListItem) HasAccessCount() bool`

HasAccessCount returns a boolean if a field has been set.

### GetMaxAccessCount

`func (o *ShareListItem) GetMaxAccessCount() int32`

GetMaxAccessCount returns the MaxAccessCount field if non-nil, zero value otherwise.

### GetMaxAccessCountOk

`func (o *ShareListItem) GetMaxAccessCountOk() (*int32, bool)`

GetMaxAccessCountOk returns a tuple with the MaxAccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAccessCount

`func (o *ShareListItem) SetMaxAccessCount(v int32)`

SetMaxAccessCount sets MaxAccessCount field to given value.

### HasMaxAccessCount

`func (o *ShareListItem) HasMaxAccessCount() bool`

HasMaxAccessCount returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ShareListItem) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ShareListItem) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ShareListItem) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ShareListItem) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ShareListItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ShareListItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ShareListItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ShareListItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetIsExpired

`func (o *ShareListItem) GetIsExpired() bool`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *ShareListItem) GetIsExpiredOk() (*bool, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *ShareListItem) SetIsExpired(v bool)`

SetIsExpired sets IsExpired field to given value.

### HasIsExpired

`func (o *ShareListItem) HasIsExpired() bool`

HasIsExpired returns a boolean if a field has been set.

### GetIsAccepted

`func (o *ShareListItem) GetIsAccepted() bool`

GetIsAccepted returns the IsAccepted field if non-nil, zero value otherwise.

### GetIsAcceptedOk

`func (o *ShareListItem) GetIsAcceptedOk() (*bool, bool)`

GetIsAcceptedOk returns a tuple with the IsAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccepted

`func (o *ShareListItem) SetIsAccepted(v bool)`

SetIsAccepted sets IsAccepted field to given value.

### HasIsAccepted

`func (o *ShareListItem) HasIsAccepted() bool`

HasIsAccepted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


