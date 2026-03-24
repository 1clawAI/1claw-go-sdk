# DeviceTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceCode** | **string** |  | 
**GrantType** | **string** |  | 

## Methods

### NewDeviceTokenRequest

`func NewDeviceTokenRequest(deviceCode string, grantType string, ) *DeviceTokenRequest`

NewDeviceTokenRequest instantiates a new DeviceTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceTokenRequestWithDefaults

`func NewDeviceTokenRequestWithDefaults() *DeviceTokenRequest`

NewDeviceTokenRequestWithDefaults instantiates a new DeviceTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceCode

`func (o *DeviceTokenRequest) GetDeviceCode() string`

GetDeviceCode returns the DeviceCode field if non-nil, zero value otherwise.

### GetDeviceCodeOk

`func (o *DeviceTokenRequest) GetDeviceCodeOk() (*string, bool)`

GetDeviceCodeOk returns a tuple with the DeviceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCode

`func (o *DeviceTokenRequest) SetDeviceCode(v string)`

SetDeviceCode sets DeviceCode field to given value.


### GetGrantType

`func (o *DeviceTokenRequest) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *DeviceTokenRequest) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *DeviceTokenRequest) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


