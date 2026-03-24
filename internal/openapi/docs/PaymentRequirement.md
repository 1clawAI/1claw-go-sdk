# PaymentRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X402Version** | Pointer to **int32** |  | [optional] 
**Accepts** | Pointer to [**[]PaymentRequirementAcceptsInner**](PaymentRequirementAcceptsInner.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentRequirement

`func NewPaymentRequirement() *PaymentRequirement`

NewPaymentRequirement instantiates a new PaymentRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequirementWithDefaults

`func NewPaymentRequirementWithDefaults() *PaymentRequirement`

NewPaymentRequirementWithDefaults instantiates a new PaymentRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX402Version

`func (o *PaymentRequirement) GetX402Version() int32`

GetX402Version returns the X402Version field if non-nil, zero value otherwise.

### GetX402VersionOk

`func (o *PaymentRequirement) GetX402VersionOk() (*int32, bool)`

GetX402VersionOk returns a tuple with the X402Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX402Version

`func (o *PaymentRequirement) SetX402Version(v int32)`

SetX402Version sets X402Version field to given value.

### HasX402Version

`func (o *PaymentRequirement) HasX402Version() bool`

HasX402Version returns a boolean if a field has been set.

### GetAccepts

`func (o *PaymentRequirement) GetAccepts() []PaymentRequirementAcceptsInner`

GetAccepts returns the Accepts field if non-nil, zero value otherwise.

### GetAcceptsOk

`func (o *PaymentRequirement) GetAcceptsOk() (*[]PaymentRequirementAcceptsInner, bool)`

GetAcceptsOk returns a tuple with the Accepts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepts

`func (o *PaymentRequirement) SetAccepts(v []PaymentRequirementAcceptsInner)`

SetAccepts sets Accepts field to given value.

### HasAccepts

`func (o *PaymentRequirement) HasAccepts() bool`

HasAccepts returns a boolean if a field has been set.

### GetDescription

`func (o *PaymentRequirement) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentRequirement) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentRequirement) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentRequirement) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


