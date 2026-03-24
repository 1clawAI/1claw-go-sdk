# OrgMemberListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to [**[]OrgMemberResponse**](OrgMemberResponse.md) |  | [optional] 

## Methods

### NewOrgMemberListResponse

`func NewOrgMemberListResponse() *OrgMemberListResponse`

NewOrgMemberListResponse instantiates a new OrgMemberListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgMemberListResponseWithDefaults

`func NewOrgMemberListResponseWithDefaults() *OrgMemberListResponse`

NewOrgMemberListResponseWithDefaults instantiates a new OrgMemberListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *OrgMemberListResponse) GetMembers() []OrgMemberResponse`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *OrgMemberListResponse) GetMembersOk() (*[]OrgMemberResponse, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *OrgMemberListResponse) SetMembers(v []OrgMemberResponse)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *OrgMemberListResponse) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


