# AddSignerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignerType** | **string** |  | 
**SignerId** | **string** |  | 
**SignerAddress** | **string** |  | 

## Methods

### NewAddSignerRequest

`func NewAddSignerRequest(signerType string, signerId string, signerAddress string, ) *AddSignerRequest`

NewAddSignerRequest instantiates a new AddSignerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSignerRequestWithDefaults

`func NewAddSignerRequestWithDefaults() *AddSignerRequest`

NewAddSignerRequestWithDefaults instantiates a new AddSignerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignerType

`func (o *AddSignerRequest) GetSignerType() string`

GetSignerType returns the SignerType field if non-nil, zero value otherwise.

### GetSignerTypeOk

`func (o *AddSignerRequest) GetSignerTypeOk() (*string, bool)`

GetSignerTypeOk returns a tuple with the SignerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerType

`func (o *AddSignerRequest) SetSignerType(v string)`

SetSignerType sets SignerType field to given value.


### GetSignerId

`func (o *AddSignerRequest) GetSignerId() string`

GetSignerId returns the SignerId field if non-nil, zero value otherwise.

### GetSignerIdOk

`func (o *AddSignerRequest) GetSignerIdOk() (*string, bool)`

GetSignerIdOk returns a tuple with the SignerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerId

`func (o *AddSignerRequest) SetSignerId(v string)`

SetSignerId sets SignerId field to given value.


### GetSignerAddress

`func (o *AddSignerRequest) GetSignerAddress() string`

GetSignerAddress returns the SignerAddress field if non-nil, zero value otherwise.

### GetSignerAddressOk

`func (o *AddSignerRequest) GetSignerAddressOk() (*string, bool)`

GetSignerAddressOk returns a tuple with the SignerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerAddress

`func (o *AddSignerRequest) SetSignerAddress(v string)`

SetSignerAddress sets SignerAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


