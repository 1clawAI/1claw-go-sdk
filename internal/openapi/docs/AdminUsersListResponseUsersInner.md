# AdminUsersListResponseUsersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**AuthMethod** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**OrgName** | Pointer to **string** |  | [optional] 
**BillingTier** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**FreeTierOverride** | Pointer to **int32** |  | [optional] 
**IsSponsored** | Pointer to **bool** |  | [optional] 
**CurrentMonthRequests** | Pointer to **int32** |  | [optional] 

## Methods

### NewAdminUsersListResponseUsersInner

`func NewAdminUsersListResponseUsersInner() *AdminUsersListResponseUsersInner`

NewAdminUsersListResponseUsersInner instantiates a new AdminUsersListResponseUsersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminUsersListResponseUsersInnerWithDefaults

`func NewAdminUsersListResponseUsersInnerWithDefaults() *AdminUsersListResponseUsersInner`

NewAdminUsersListResponseUsersInnerWithDefaults instantiates a new AdminUsersListResponseUsersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdminUsersListResponseUsersInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdminUsersListResponseUsersInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdminUsersListResponseUsersInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdminUsersListResponseUsersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmail

`func (o *AdminUsersListResponseUsersInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AdminUsersListResponseUsersInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AdminUsersListResponseUsersInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AdminUsersListResponseUsersInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDisplayName

`func (o *AdminUsersListResponseUsersInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AdminUsersListResponseUsersInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AdminUsersListResponseUsersInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AdminUsersListResponseUsersInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetRole

`func (o *AdminUsersListResponseUsersInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AdminUsersListResponseUsersInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AdminUsersListResponseUsersInner) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AdminUsersListResponseUsersInner) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetAuthMethod

`func (o *AdminUsersListResponseUsersInner) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *AdminUsersListResponseUsersInner) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *AdminUsersListResponseUsersInner) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *AdminUsersListResponseUsersInner) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### GetOrgId

`func (o *AdminUsersListResponseUsersInner) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AdminUsersListResponseUsersInner) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AdminUsersListResponseUsersInner) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *AdminUsersListResponseUsersInner) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrgName

`func (o *AdminUsersListResponseUsersInner) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *AdminUsersListResponseUsersInner) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *AdminUsersListResponseUsersInner) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *AdminUsersListResponseUsersInner) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetBillingTier

`func (o *AdminUsersListResponseUsersInner) GetBillingTier() string`

GetBillingTier returns the BillingTier field if non-nil, zero value otherwise.

### GetBillingTierOk

`func (o *AdminUsersListResponseUsersInner) GetBillingTierOk() (*string, bool)`

GetBillingTierOk returns a tuple with the BillingTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTier

`func (o *AdminUsersListResponseUsersInner) SetBillingTier(v string)`

SetBillingTier sets BillingTier field to given value.

### HasBillingTier

`func (o *AdminUsersListResponseUsersInner) HasBillingTier() bool`

HasBillingTier returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AdminUsersListResponseUsersInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdminUsersListResponseUsersInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdminUsersListResponseUsersInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AdminUsersListResponseUsersInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFreeTierOverride

`func (o *AdminUsersListResponseUsersInner) GetFreeTierOverride() int32`

GetFreeTierOverride returns the FreeTierOverride field if non-nil, zero value otherwise.

### GetFreeTierOverrideOk

`func (o *AdminUsersListResponseUsersInner) GetFreeTierOverrideOk() (*int32, bool)`

GetFreeTierOverrideOk returns a tuple with the FreeTierOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeTierOverride

`func (o *AdminUsersListResponseUsersInner) SetFreeTierOverride(v int32)`

SetFreeTierOverride sets FreeTierOverride field to given value.

### HasFreeTierOverride

`func (o *AdminUsersListResponseUsersInner) HasFreeTierOverride() bool`

HasFreeTierOverride returns a boolean if a field has been set.

### GetIsSponsored

`func (o *AdminUsersListResponseUsersInner) GetIsSponsored() bool`

GetIsSponsored returns the IsSponsored field if non-nil, zero value otherwise.

### GetIsSponsoredOk

`func (o *AdminUsersListResponseUsersInner) GetIsSponsoredOk() (*bool, bool)`

GetIsSponsoredOk returns a tuple with the IsSponsored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSponsored

`func (o *AdminUsersListResponseUsersInner) SetIsSponsored(v bool)`

SetIsSponsored sets IsSponsored field to given value.

### HasIsSponsored

`func (o *AdminUsersListResponseUsersInner) HasIsSponsored() bool`

HasIsSponsored returns a boolean if a field has been set.

### GetCurrentMonthRequests

`func (o *AdminUsersListResponseUsersInner) GetCurrentMonthRequests() int32`

GetCurrentMonthRequests returns the CurrentMonthRequests field if non-nil, zero value otherwise.

### GetCurrentMonthRequestsOk

`func (o *AdminUsersListResponseUsersInner) GetCurrentMonthRequestsOk() (*int32, bool)`

GetCurrentMonthRequestsOk returns a tuple with the CurrentMonthRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMonthRequests

`func (o *AdminUsersListResponseUsersInner) SetCurrentMonthRequests(v int32)`

SetCurrentMonthRequests sets CurrentMonthRequests field to given value.

### HasCurrentMonthRequests

`func (o *AdminUsersListResponseUsersInner) HasCurrentMonthRequests() bool`

HasCurrentMonthRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


