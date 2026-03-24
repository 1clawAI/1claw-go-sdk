# DeviceCodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** |  | 
**Email** | **string** | Account email; only that user may approve the code in the dashboard. | 

## Methods

### NewDeviceCodeRequest

`func NewDeviceCodeRequest(clientId string, email string, ) *DeviceCodeRequest`

NewDeviceCodeRequest instantiates a new DeviceCodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCodeRequestWithDefaults

`func NewDeviceCodeRequestWithDefaults() *DeviceCodeRequest`

NewDeviceCodeRequestWithDefaults instantiates a new DeviceCodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *DeviceCodeRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *DeviceCodeRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *DeviceCodeRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetEmail

`func (o *DeviceCodeRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DeviceCodeRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DeviceCodeRequest) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


