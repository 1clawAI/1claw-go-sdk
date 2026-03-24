# SimulateTransactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **string** |  | 
**Value** | **string** |  | 
**Chain** | **string** |  | 
**Data** | Pointer to **string** |  | [optional] 
**SigningKeyPath** | Pointer to **string** |  | [optional] 
**GasLimit** | Pointer to **int32** |  | [optional] 

## Methods

### NewSimulateTransactionRequest

`func NewSimulateTransactionRequest(to string, value string, chain string, ) *SimulateTransactionRequest`

NewSimulateTransactionRequest instantiates a new SimulateTransactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulateTransactionRequestWithDefaults

`func NewSimulateTransactionRequestWithDefaults() *SimulateTransactionRequest`

NewSimulateTransactionRequestWithDefaults instantiates a new SimulateTransactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *SimulateTransactionRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SimulateTransactionRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SimulateTransactionRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetValue

`func (o *SimulateTransactionRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SimulateTransactionRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SimulateTransactionRequest) SetValue(v string)`

SetValue sets Value field to given value.


### GetChain

`func (o *SimulateTransactionRequest) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *SimulateTransactionRequest) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *SimulateTransactionRequest) SetChain(v string)`

SetChain sets Chain field to given value.


### GetData

`func (o *SimulateTransactionRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SimulateTransactionRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SimulateTransactionRequest) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *SimulateTransactionRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSigningKeyPath

`func (o *SimulateTransactionRequest) GetSigningKeyPath() string`

GetSigningKeyPath returns the SigningKeyPath field if non-nil, zero value otherwise.

### GetSigningKeyPathOk

`func (o *SimulateTransactionRequest) GetSigningKeyPathOk() (*string, bool)`

GetSigningKeyPathOk returns a tuple with the SigningKeyPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKeyPath

`func (o *SimulateTransactionRequest) SetSigningKeyPath(v string)`

SetSigningKeyPath sets SigningKeyPath field to given value.

### HasSigningKeyPath

`func (o *SimulateTransactionRequest) HasSigningKeyPath() bool`

HasSigningKeyPath returns a boolean if a field has been set.

### GetGasLimit

`func (o *SimulateTransactionRequest) GetGasLimit() int32`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *SimulateTransactionRequest) GetGasLimitOk() (*int32, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *SimulateTransactionRequest) SetGasLimit(v int32)`

SetGasLimit sets GasLimit field to given value.

### HasGasLimit

`func (o *SimulateTransactionRequest) HasGasLimit() bool`

HasGasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


