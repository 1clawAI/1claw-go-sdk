# CreditBalanceResponseExpiringWithin90Days

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int32** |  | [optional] 
**EarliestExpiry** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCreditBalanceResponseExpiringWithin90Days

`func NewCreditBalanceResponseExpiringWithin90Days() *CreditBalanceResponseExpiringWithin90Days`

NewCreditBalanceResponseExpiringWithin90Days instantiates a new CreditBalanceResponseExpiringWithin90Days object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditBalanceResponseExpiringWithin90DaysWithDefaults

`func NewCreditBalanceResponseExpiringWithin90DaysWithDefaults() *CreditBalanceResponseExpiringWithin90Days`

NewCreditBalanceResponseExpiringWithin90DaysWithDefaults instantiates a new CreditBalanceResponseExpiringWithin90Days object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *CreditBalanceResponseExpiringWithin90Days) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *CreditBalanceResponseExpiringWithin90Days) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *CreditBalanceResponseExpiringWithin90Days) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *CreditBalanceResponseExpiringWithin90Days) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetEarliestExpiry

`func (o *CreditBalanceResponseExpiringWithin90Days) GetEarliestExpiry() time.Time`

GetEarliestExpiry returns the EarliestExpiry field if non-nil, zero value otherwise.

### GetEarliestExpiryOk

`func (o *CreditBalanceResponseExpiringWithin90Days) GetEarliestExpiryOk() (*time.Time, bool)`

GetEarliestExpiryOk returns a tuple with the EarliestExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestExpiry

`func (o *CreditBalanceResponseExpiringWithin90Days) SetEarliestExpiry(v time.Time)`

SetEarliestExpiry sets EarliestExpiry field to given value.

### HasEarliestExpiry

`func (o *CreditBalanceResponseExpiringWithin90Days) HasEarliestExpiry() bool`

HasEarliestExpiry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


