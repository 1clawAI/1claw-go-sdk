# AdminUsersListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**[]AdminUsersListResponseUsersInner**](AdminUsersListResponseUsersInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewAdminUsersListResponse

`func NewAdminUsersListResponse() *AdminUsersListResponse`

NewAdminUsersListResponse instantiates a new AdminUsersListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminUsersListResponseWithDefaults

`func NewAdminUsersListResponseWithDefaults() *AdminUsersListResponse`

NewAdminUsersListResponseWithDefaults instantiates a new AdminUsersListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *AdminUsersListResponse) GetUsers() []AdminUsersListResponseUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AdminUsersListResponse) GetUsersOk() (*[]AdminUsersListResponseUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AdminUsersListResponse) SetUsers(v []AdminUsersListResponseUsersInner)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *AdminUsersListResponse) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetTotal

`func (o *AdminUsersListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AdminUsersListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AdminUsersListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AdminUsersListResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


