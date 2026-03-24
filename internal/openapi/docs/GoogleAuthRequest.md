# GoogleAuthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdToken** | **string** |  | 

## Methods

### NewGoogleAuthRequest

`func NewGoogleAuthRequest(idToken string, ) *GoogleAuthRequest`

NewGoogleAuthRequest instantiates a new GoogleAuthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleAuthRequestWithDefaults

`func NewGoogleAuthRequestWithDefaults() *GoogleAuthRequest`

NewGoogleAuthRequestWithDefaults instantiates a new GoogleAuthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdToken

`func (o *GoogleAuthRequest) GetIdToken() string`

GetIdToken returns the IdToken field if non-nil, zero value otherwise.

### GetIdTokenOk

`func (o *GoogleAuthRequest) GetIdTokenOk() (*string, bool)`

GetIdTokenOk returns a tuple with the IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdToken

`func (o *GoogleAuthRequest) SetIdToken(v string)`

SetIdToken sets IdToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


