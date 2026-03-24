# ApiKeyListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | Pointer to [**[]ApiKeyResponse**](ApiKeyResponse.md) |  | [optional] 

## Methods

### NewApiKeyListResponse

`func NewApiKeyListResponse() *ApiKeyListResponse`

NewApiKeyListResponse instantiates a new ApiKeyListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyListResponseWithDefaults

`func NewApiKeyListResponseWithDefaults() *ApiKeyListResponse`

NewApiKeyListResponseWithDefaults instantiates a new ApiKeyListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *ApiKeyListResponse) GetKeys() []ApiKeyResponse`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *ApiKeyListResponse) GetKeysOk() (*[]ApiKeyResponse, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *ApiKeyListResponse) SetKeys(v []ApiKeyResponse)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *ApiKeyListResponse) HasKeys() bool`

HasKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


