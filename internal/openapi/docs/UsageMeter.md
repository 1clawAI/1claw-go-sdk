# UsageMeter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Used** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 

## Methods

### NewUsageMeter

`func NewUsageMeter() *UsageMeter`

NewUsageMeter instantiates a new UsageMeter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageMeterWithDefaults

`func NewUsageMeterWithDefaults() *UsageMeter`

NewUsageMeterWithDefaults instantiates a new UsageMeter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsed

`func (o *UsageMeter) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *UsageMeter) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *UsageMeter) SetUsed(v int32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *UsageMeter) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetLimit

`func (o *UsageMeter) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *UsageMeter) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *UsageMeter) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *UsageMeter) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


