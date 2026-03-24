# SubscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tier** | Pointer to **string** |  | [optional] 
**Interval** | Pointer to **string** |  | [optional] 
**PeriodEnd** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CreditBalanceCents** | Pointer to **int32** |  | [optional] 
**CreditBalanceUsd** | Pointer to **string** |  | [optional] 
**OverageMethod** | Pointer to **string** |  | [optional] 
**Usage** | Pointer to [**SubscriptionResponseUsage**](SubscriptionResponseUsage.md) |  | [optional] 

## Methods

### NewSubscriptionResponse

`func NewSubscriptionResponse() *SubscriptionResponse`

NewSubscriptionResponse instantiates a new SubscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionResponseWithDefaults

`func NewSubscriptionResponseWithDefaults() *SubscriptionResponse`

NewSubscriptionResponseWithDefaults instantiates a new SubscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTier

`func (o *SubscriptionResponse) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *SubscriptionResponse) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *SubscriptionResponse) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *SubscriptionResponse) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetInterval

`func (o *SubscriptionResponse) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *SubscriptionResponse) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *SubscriptionResponse) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *SubscriptionResponse) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *SubscriptionResponse) GetPeriodEnd() time.Time`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *SubscriptionResponse) GetPeriodEndOk() (*time.Time, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *SubscriptionResponse) SetPeriodEnd(v time.Time)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *SubscriptionResponse) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetStatus

`func (o *SubscriptionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubscriptionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreditBalanceCents

`func (o *SubscriptionResponse) GetCreditBalanceCents() int32`

GetCreditBalanceCents returns the CreditBalanceCents field if non-nil, zero value otherwise.

### GetCreditBalanceCentsOk

`func (o *SubscriptionResponse) GetCreditBalanceCentsOk() (*int32, bool)`

GetCreditBalanceCentsOk returns a tuple with the CreditBalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditBalanceCents

`func (o *SubscriptionResponse) SetCreditBalanceCents(v int32)`

SetCreditBalanceCents sets CreditBalanceCents field to given value.

### HasCreditBalanceCents

`func (o *SubscriptionResponse) HasCreditBalanceCents() bool`

HasCreditBalanceCents returns a boolean if a field has been set.

### GetCreditBalanceUsd

`func (o *SubscriptionResponse) GetCreditBalanceUsd() string`

GetCreditBalanceUsd returns the CreditBalanceUsd field if non-nil, zero value otherwise.

### GetCreditBalanceUsdOk

`func (o *SubscriptionResponse) GetCreditBalanceUsdOk() (*string, bool)`

GetCreditBalanceUsdOk returns a tuple with the CreditBalanceUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditBalanceUsd

`func (o *SubscriptionResponse) SetCreditBalanceUsd(v string)`

SetCreditBalanceUsd sets CreditBalanceUsd field to given value.

### HasCreditBalanceUsd

`func (o *SubscriptionResponse) HasCreditBalanceUsd() bool`

HasCreditBalanceUsd returns a boolean if a field has been set.

### GetOverageMethod

`func (o *SubscriptionResponse) GetOverageMethod() string`

GetOverageMethod returns the OverageMethod field if non-nil, zero value otherwise.

### GetOverageMethodOk

`func (o *SubscriptionResponse) GetOverageMethodOk() (*string, bool)`

GetOverageMethodOk returns a tuple with the OverageMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverageMethod

`func (o *SubscriptionResponse) SetOverageMethod(v string)`

SetOverageMethod sets OverageMethod field to given value.

### HasOverageMethod

`func (o *SubscriptionResponse) HasOverageMethod() bool`

HasOverageMethod returns a boolean if a field has been set.

### GetUsage

`func (o *SubscriptionResponse) GetUsage() SubscriptionResponseUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *SubscriptionResponse) GetUsageOk() (*SubscriptionResponseUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *SubscriptionResponse) SetUsage(v SubscriptionResponseUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *SubscriptionResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


