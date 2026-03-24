# CreditTransactionsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | Pointer to [**[]CreditTransactionsListResponseTransactionsInner**](CreditTransactionsListResponseTransactionsInner.md) |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 

## Methods

### NewCreditTransactionsListResponse

`func NewCreditTransactionsListResponse() *CreditTransactionsListResponse`

NewCreditTransactionsListResponse instantiates a new CreditTransactionsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditTransactionsListResponseWithDefaults

`func NewCreditTransactionsListResponseWithDefaults() *CreditTransactionsListResponse`

NewCreditTransactionsListResponseWithDefaults instantiates a new CreditTransactionsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *CreditTransactionsListResponse) GetTransactions() []CreditTransactionsListResponseTransactionsInner`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *CreditTransactionsListResponse) GetTransactionsOk() (*[]CreditTransactionsListResponseTransactionsInner, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *CreditTransactionsListResponse) SetTransactions(v []CreditTransactionsListResponseTransactionsInner)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *CreditTransactionsListResponse) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetPage

`func (o *CreditTransactionsListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *CreditTransactionsListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *CreditTransactionsListResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *CreditTransactionsListResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLimit

`func (o *CreditTransactionsListResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CreditTransactionsListResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CreditTransactionsListResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CreditTransactionsListResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


