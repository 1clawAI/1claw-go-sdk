# CreditBalanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceCents** | Pointer to **int32** |  | [optional] 
**BalanceUsd** | Pointer to **string** |  | [optional] 
**ExpiringWithin90Days** | Pointer to [**CreditBalanceResponseExpiringWithin90Days**](CreditBalanceResponseExpiringWithin90Days.md) |  | [optional] 

## Methods

### NewCreditBalanceResponse

`func NewCreditBalanceResponse() *CreditBalanceResponse`

NewCreditBalanceResponse instantiates a new CreditBalanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditBalanceResponseWithDefaults

`func NewCreditBalanceResponseWithDefaults() *CreditBalanceResponse`

NewCreditBalanceResponseWithDefaults instantiates a new CreditBalanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceCents

`func (o *CreditBalanceResponse) GetBalanceCents() int32`

GetBalanceCents returns the BalanceCents field if non-nil, zero value otherwise.

### GetBalanceCentsOk

`func (o *CreditBalanceResponse) GetBalanceCentsOk() (*int32, bool)`

GetBalanceCentsOk returns a tuple with the BalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCents

`func (o *CreditBalanceResponse) SetBalanceCents(v int32)`

SetBalanceCents sets BalanceCents field to given value.

### HasBalanceCents

`func (o *CreditBalanceResponse) HasBalanceCents() bool`

HasBalanceCents returns a boolean if a field has been set.

### GetBalanceUsd

`func (o *CreditBalanceResponse) GetBalanceUsd() string`

GetBalanceUsd returns the BalanceUsd field if non-nil, zero value otherwise.

### GetBalanceUsdOk

`func (o *CreditBalanceResponse) GetBalanceUsdOk() (*string, bool)`

GetBalanceUsdOk returns a tuple with the BalanceUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceUsd

`func (o *CreditBalanceResponse) SetBalanceUsd(v string)`

SetBalanceUsd sets BalanceUsd field to given value.

### HasBalanceUsd

`func (o *CreditBalanceResponse) HasBalanceUsd() bool`

HasBalanceUsd returns a boolean if a field has been set.

### GetExpiringWithin90Days

`func (o *CreditBalanceResponse) GetExpiringWithin90Days() CreditBalanceResponseExpiringWithin90Days`

GetExpiringWithin90Days returns the ExpiringWithin90Days field if non-nil, zero value otherwise.

### GetExpiringWithin90DaysOk

`func (o *CreditBalanceResponse) GetExpiringWithin90DaysOk() (*CreditBalanceResponseExpiringWithin90Days, bool)`

GetExpiringWithin90DaysOk returns a tuple with the ExpiringWithin90Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiringWithin90Days

`func (o *CreditBalanceResponse) SetExpiringWithin90Days(v CreditBalanceResponseExpiringWithin90Days)`

SetExpiringWithin90Days sets ExpiringWithin90Days field to given value.

### HasExpiringWithin90Days

`func (o *CreditBalanceResponse) HasExpiringWithin90Days() bool`

HasExpiringWithin90Days returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


