# ShareListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shares** | Pointer to [**[]ShareListItem**](ShareListItem.md) |  | [optional] 

## Methods

### NewShareListResponse

`func NewShareListResponse() *ShareListResponse`

NewShareListResponse instantiates a new ShareListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareListResponseWithDefaults

`func NewShareListResponseWithDefaults() *ShareListResponse`

NewShareListResponseWithDefaults instantiates a new ShareListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShares

`func (o *ShareListResponse) GetShares() []ShareListItem`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *ShareListResponse) GetSharesOk() (*[]ShareListItem, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *ShareListResponse) SetShares(v []ShareListItem)`

SetShares sets Shares field to given value.

### HasShares

`func (o *ShareListResponse) HasShares() bool`

HasShares returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


