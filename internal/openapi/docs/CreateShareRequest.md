# CreateShareRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecipientType** | **string** |  | 
**RecipientId** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 
**MaxAccessCount** | Pointer to **int32** |  | [optional] 
**ExpiresAt** | **time.Time** |  | 
**Passphrase** | Pointer to **string** |  | [optional] 
**IpAllowlist** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateShareRequest

`func NewCreateShareRequest(recipientType string, expiresAt time.Time, ) *CreateShareRequest`

NewCreateShareRequest instantiates a new CreateShareRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateShareRequestWithDefaults

`func NewCreateShareRequestWithDefaults() *CreateShareRequest`

NewCreateShareRequestWithDefaults instantiates a new CreateShareRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipientType

`func (o *CreateShareRequest) GetRecipientType() string`

GetRecipientType returns the RecipientType field if non-nil, zero value otherwise.

### GetRecipientTypeOk

`func (o *CreateShareRequest) GetRecipientTypeOk() (*string, bool)`

GetRecipientTypeOk returns a tuple with the RecipientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientType

`func (o *CreateShareRequest) SetRecipientType(v string)`

SetRecipientType sets RecipientType field to given value.


### GetRecipientId

`func (o *CreateShareRequest) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *CreateShareRequest) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *CreateShareRequest) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.

### HasRecipientId

`func (o *CreateShareRequest) HasRecipientId() bool`

HasRecipientId returns a boolean if a field has been set.

### GetEmail

`func (o *CreateShareRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateShareRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateShareRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateShareRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPermissions

`func (o *CreateShareRequest) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateShareRequest) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateShareRequest) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreateShareRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetMaxAccessCount

`func (o *CreateShareRequest) GetMaxAccessCount() int32`

GetMaxAccessCount returns the MaxAccessCount field if non-nil, zero value otherwise.

### GetMaxAccessCountOk

`func (o *CreateShareRequest) GetMaxAccessCountOk() (*int32, bool)`

GetMaxAccessCountOk returns a tuple with the MaxAccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAccessCount

`func (o *CreateShareRequest) SetMaxAccessCount(v int32)`

SetMaxAccessCount sets MaxAccessCount field to given value.

### HasMaxAccessCount

`func (o *CreateShareRequest) HasMaxAccessCount() bool`

HasMaxAccessCount returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CreateShareRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateShareRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateShareRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetPassphrase

`func (o *CreateShareRequest) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *CreateShareRequest) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *CreateShareRequest) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *CreateShareRequest) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetIpAllowlist

`func (o *CreateShareRequest) GetIpAllowlist() []string`

GetIpAllowlist returns the IpAllowlist field if non-nil, zero value otherwise.

### GetIpAllowlistOk

`func (o *CreateShareRequest) GetIpAllowlistOk() (*[]string, bool)`

GetIpAllowlistOk returns a tuple with the IpAllowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAllowlist

`func (o *CreateShareRequest) SetIpAllowlist(v []string)`

SetIpAllowlist sets IpAllowlist field to given value.

### HasIpAllowlist

`func (o *CreateShareRequest) HasIpAllowlist() bool`

HasIpAllowlist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


