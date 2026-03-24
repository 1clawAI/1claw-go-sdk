# UpdateProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**MarketingEmails** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateProfileRequest

`func NewUpdateProfileRequest() *UpdateProfileRequest`

NewUpdateProfileRequest instantiates a new UpdateProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProfileRequestWithDefaults

`func NewUpdateProfileRequestWithDefaults() *UpdateProfileRequest`

NewUpdateProfileRequestWithDefaults instantiates a new UpdateProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *UpdateProfileRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateProfileRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateProfileRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateProfileRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetMarketingEmails

`func (o *UpdateProfileRequest) GetMarketingEmails() bool`

GetMarketingEmails returns the MarketingEmails field if non-nil, zero value otherwise.

### GetMarketingEmailsOk

`func (o *UpdateProfileRequest) GetMarketingEmailsOk() (*bool, bool)`

GetMarketingEmailsOk returns a tuple with the MarketingEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingEmails

`func (o *UpdateProfileRequest) SetMarketingEmails(v bool)`

SetMarketingEmails sets MarketingEmails field to given value.

### HasMarketingEmails

`func (o *UpdateProfileRequest) HasMarketingEmails() bool`

HasMarketingEmails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


