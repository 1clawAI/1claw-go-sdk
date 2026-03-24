# SimulateBundleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | [**[]SimulateTransactionRequest**](SimulateTransactionRequest.md) |  | 

## Methods

### NewSimulateBundleRequest

`func NewSimulateBundleRequest(transactions []SimulateTransactionRequest, ) *SimulateBundleRequest`

NewSimulateBundleRequest instantiates a new SimulateBundleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulateBundleRequestWithDefaults

`func NewSimulateBundleRequestWithDefaults() *SimulateBundleRequest`

NewSimulateBundleRequestWithDefaults instantiates a new SimulateBundleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *SimulateBundleRequest) GetTransactions() []SimulateTransactionRequest`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *SimulateBundleRequest) GetTransactionsOk() (*[]SimulateTransactionRequest, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *SimulateBundleRequest) SetTransactions(v []SimulateTransactionRequest)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


