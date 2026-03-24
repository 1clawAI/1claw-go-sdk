# ShareResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ShareUrl** | Pointer to **string** |  | [optional] 
**RecipientType** | Pointer to **string** |  | [optional] 
**RecipientEmail** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**MaxAccessCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewShareResponse

`func NewShareResponse() *ShareResponse`

NewShareResponse instantiates a new ShareResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareResponseWithDefaults

`func NewShareResponseWithDefaults() *ShareResponse`

NewShareResponseWithDefaults instantiates a new ShareResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShareResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShareResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShareResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShareResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetShareUrl

`func (o *ShareResponse) GetShareUrl() string`

GetShareUrl returns the ShareUrl field if non-nil, zero value otherwise.

### GetShareUrlOk

`func (o *ShareResponse) GetShareUrlOk() (*string, bool)`

GetShareUrlOk returns a tuple with the ShareUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareUrl

`func (o *ShareResponse) SetShareUrl(v string)`

SetShareUrl sets ShareUrl field to given value.

### HasShareUrl

`func (o *ShareResponse) HasShareUrl() bool`

HasShareUrl returns a boolean if a field has been set.

### GetRecipientType

`func (o *ShareResponse) GetRecipientType() string`

GetRecipientType returns the RecipientType field if non-nil, zero value otherwise.

### GetRecipientTypeOk

`func (o *ShareResponse) GetRecipientTypeOk() (*string, bool)`

GetRecipientTypeOk returns a tuple with the RecipientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientType

`func (o *ShareResponse) SetRecipientType(v string)`

SetRecipientType sets RecipientType field to given value.

### HasRecipientType

`func (o *ShareResponse) HasRecipientType() bool`

HasRecipientType returns a boolean if a field has been set.

### GetRecipientEmail

`func (o *ShareResponse) GetRecipientEmail() string`

GetRecipientEmail returns the RecipientEmail field if non-nil, zero value otherwise.

### GetRecipientEmailOk

`func (o *ShareResponse) GetRecipientEmailOk() (*string, bool)`

GetRecipientEmailOk returns a tuple with the RecipientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientEmail

`func (o *ShareResponse) SetRecipientEmail(v string)`

SetRecipientEmail sets RecipientEmail field to given value.

### HasRecipientEmail

`func (o *ShareResponse) HasRecipientEmail() bool`

HasRecipientEmail returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ShareResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ShareResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ShareResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ShareResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetMaxAccessCount

`func (o *ShareResponse) GetMaxAccessCount() int32`

GetMaxAccessCount returns the MaxAccessCount field if non-nil, zero value otherwise.

### GetMaxAccessCountOk

`func (o *ShareResponse) GetMaxAccessCountOk() (*int32, bool)`

GetMaxAccessCountOk returns a tuple with the MaxAccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAccessCount

`func (o *ShareResponse) SetMaxAccessCount(v int32)`

SetMaxAccessCount sets MaxAccessCount field to given value.

### HasMaxAccessCount

`func (o *ShareResponse) HasMaxAccessCount() bool`

HasMaxAccessCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


