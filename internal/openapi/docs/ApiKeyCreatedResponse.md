# ApiKeyCreatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to [**ApiKeyResponse**](ApiKeyResponse.md) |  | [optional] 
**ApiKey** | Pointer to **string** | Full key (shown once) | [optional] 

## Methods

### NewApiKeyCreatedResponse

`func NewApiKeyCreatedResponse() *ApiKeyCreatedResponse`

NewApiKeyCreatedResponse instantiates a new ApiKeyCreatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyCreatedResponseWithDefaults

`func NewApiKeyCreatedResponseWithDefaults() *ApiKeyCreatedResponse`

NewApiKeyCreatedResponseWithDefaults instantiates a new ApiKeyCreatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ApiKeyCreatedResponse) GetKey() ApiKeyResponse`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApiKeyCreatedResponse) GetKeyOk() (*ApiKeyResponse, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApiKeyCreatedResponse) SetKey(v ApiKeyResponse)`

SetKey sets Key field to given value.

### HasKey

`func (o *ApiKeyCreatedResponse) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetApiKey

`func (o *ApiKeyCreatedResponse) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ApiKeyCreatedResponse) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ApiKeyCreatedResponse) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *ApiKeyCreatedResponse) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


