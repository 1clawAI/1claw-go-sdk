# SetBillingTierRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tier** | **string** |  | 
**DurationDays** | Pointer to **int32** | How many days the tier lasts (default 365). Use 90 for a 3-month trial. | [optional] 

## Methods

### NewSetBillingTierRequest

`func NewSetBillingTierRequest(tier string, ) *SetBillingTierRequest`

NewSetBillingTierRequest instantiates a new SetBillingTierRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetBillingTierRequestWithDefaults

`func NewSetBillingTierRequestWithDefaults() *SetBillingTierRequest`

NewSetBillingTierRequestWithDefaults instantiates a new SetBillingTierRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTier

`func (o *SetBillingTierRequest) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *SetBillingTierRequest) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *SetBillingTierRequest) SetTier(v string)`

SetTier sets Tier field to given value.


### GetDurationDays

`func (o *SetBillingTierRequest) GetDurationDays() int32`

GetDurationDays returns the DurationDays field if non-nil, zero value otherwise.

### GetDurationDaysOk

`func (o *SetBillingTierRequest) GetDurationDaysOk() (*int32, bool)`

GetDurationDaysOk returns a tuple with the DurationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationDays

`func (o *SetBillingTierRequest) SetDurationDays(v int32)`

SetDurationDays sets DurationDays field to given value.

### HasDurationDays

`func (o *SetBillingTierRequest) HasDurationDays() bool`

HasDurationDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


