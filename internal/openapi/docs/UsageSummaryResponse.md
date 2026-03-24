# UsageSummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingTier** | Pointer to **string** |  | [optional] 
**FreeTierLimit** | Pointer to **int32** |  | [optional] 
**CurrentMonth** | Pointer to [**UsageSummaryResponseCurrentMonth**](UsageSummaryResponseCurrentMonth.md) |  | [optional] 

## Methods

### NewUsageSummaryResponse

`func NewUsageSummaryResponse() *UsageSummaryResponse`

NewUsageSummaryResponse instantiates a new UsageSummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageSummaryResponseWithDefaults

`func NewUsageSummaryResponseWithDefaults() *UsageSummaryResponse`

NewUsageSummaryResponseWithDefaults instantiates a new UsageSummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingTier

`func (o *UsageSummaryResponse) GetBillingTier() string`

GetBillingTier returns the BillingTier field if non-nil, zero value otherwise.

### GetBillingTierOk

`func (o *UsageSummaryResponse) GetBillingTierOk() (*string, bool)`

GetBillingTierOk returns a tuple with the BillingTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTier

`func (o *UsageSummaryResponse) SetBillingTier(v string)`

SetBillingTier sets BillingTier field to given value.

### HasBillingTier

`func (o *UsageSummaryResponse) HasBillingTier() bool`

HasBillingTier returns a boolean if a field has been set.

### GetFreeTierLimit

`func (o *UsageSummaryResponse) GetFreeTierLimit() int32`

GetFreeTierLimit returns the FreeTierLimit field if non-nil, zero value otherwise.

### GetFreeTierLimitOk

`func (o *UsageSummaryResponse) GetFreeTierLimitOk() (*int32, bool)`

GetFreeTierLimitOk returns a tuple with the FreeTierLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeTierLimit

`func (o *UsageSummaryResponse) SetFreeTierLimit(v int32)`

SetFreeTierLimit sets FreeTierLimit field to given value.

### HasFreeTierLimit

`func (o *UsageSummaryResponse) HasFreeTierLimit() bool`

HasFreeTierLimit returns a boolean if a field has been set.

### GetCurrentMonth

`func (o *UsageSummaryResponse) GetCurrentMonth() UsageSummaryResponseCurrentMonth`

GetCurrentMonth returns the CurrentMonth field if non-nil, zero value otherwise.

### GetCurrentMonthOk

`func (o *UsageSummaryResponse) GetCurrentMonthOk() (*UsageSummaryResponseCurrentMonth, bool)`

GetCurrentMonthOk returns a tuple with the CurrentMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMonth

`func (o *UsageSummaryResponse) SetCurrentMonth(v UsageSummaryResponseCurrentMonth)`

SetCurrentMonth sets CurrentMonth field to given value.

### HasCurrentMonth

`func (o *UsageSummaryResponse) HasCurrentMonth() bool`

HasCurrentMonth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


