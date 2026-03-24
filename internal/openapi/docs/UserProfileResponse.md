# UserProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**AuthMethod** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**EmailVerified** | Pointer to **bool** |  | [optional] 
**MarketingEmails** | Pointer to **bool** |  | [optional] 
**TotpEnabled** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUserProfileResponse

`func NewUserProfileResponse() *UserProfileResponse`

NewUserProfileResponse instantiates a new UserProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProfileResponseWithDefaults

`func NewUserProfileResponseWithDefaults() *UserProfileResponse`

NewUserProfileResponseWithDefaults instantiates a new UserProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserProfileResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserProfileResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserProfileResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserProfileResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmail

`func (o *UserProfileResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserProfileResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserProfileResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserProfileResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDisplayName

`func (o *UserProfileResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserProfileResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserProfileResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserProfileResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAuthMethod

`func (o *UserProfileResponse) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *UserProfileResponse) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *UserProfileResponse) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *UserProfileResponse) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### GetRole

`func (o *UserProfileResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserProfileResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserProfileResponse) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *UserProfileResponse) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetEmailVerified

`func (o *UserProfileResponse) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *UserProfileResponse) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *UserProfileResponse) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *UserProfileResponse) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### GetMarketingEmails

`func (o *UserProfileResponse) GetMarketingEmails() bool`

GetMarketingEmails returns the MarketingEmails field if non-nil, zero value otherwise.

### GetMarketingEmailsOk

`func (o *UserProfileResponse) GetMarketingEmailsOk() (*bool, bool)`

GetMarketingEmailsOk returns a tuple with the MarketingEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingEmails

`func (o *UserProfileResponse) SetMarketingEmails(v bool)`

SetMarketingEmails sets MarketingEmails field to given value.

### HasMarketingEmails

`func (o *UserProfileResponse) HasMarketingEmails() bool`

HasMarketingEmails returns a boolean if a field has been set.

### GetTotpEnabled

`func (o *UserProfileResponse) GetTotpEnabled() bool`

GetTotpEnabled returns the TotpEnabled field if non-nil, zero value otherwise.

### GetTotpEnabledOk

`func (o *UserProfileResponse) GetTotpEnabledOk() (*bool, bool)`

GetTotpEnabledOk returns a tuple with the TotpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpEnabled

`func (o *UserProfileResponse) SetTotpEnabled(v bool)`

SetTotpEnabled sets TotpEnabled field to given value.

### HasTotpEnabled

`func (o *UserProfileResponse) HasTotpEnabled() bool`

HasTotpEnabled returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UserProfileResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserProfileResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserProfileResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserProfileResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


