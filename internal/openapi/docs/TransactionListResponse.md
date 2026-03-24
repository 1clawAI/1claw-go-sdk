# TransactionListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | Pointer to [**[]TransactionResponse**](TransactionResponse.md) |  | [optional] 

## Methods

### NewTransactionListResponse

`func NewTransactionListResponse() *TransactionListResponse`

NewTransactionListResponse instantiates a new TransactionListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionListResponseWithDefaults

`func NewTransactionListResponseWithDefaults() *TransactionListResponse`

NewTransactionListResponseWithDefaults instantiates a new TransactionListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *TransactionListResponse) GetTransactions() []TransactionResponse`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *TransactionListResponse) GetTransactionsOk() (*[]TransactionResponse, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *TransactionListResponse) SetTransactions(v []TransactionResponse)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *TransactionListResponse) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


