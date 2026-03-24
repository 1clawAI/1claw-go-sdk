# SignTransactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **string** | Destination address (0x-prefixed) | 
**Value** | **string** | Value in ETH | 
**Chain** | **string** | Chain name or numeric ID | 
**Data** | Pointer to **string** | Hex-encoded calldata | [optional] 
**SigningKeyPath** | Pointer to **string** |  | [optional] 
**Nonce** | Pointer to **int32** |  | [optional] 
**GasPrice** | Pointer to **string** |  | [optional] 
**GasLimit** | Pointer to **int32** |  | [optional] 
**MaxFeePerGas** | Pointer to **string** |  | [optional] 
**MaxPriorityFeePerGas** | Pointer to **string** |  | [optional] 
**SimulateFirst** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSignTransactionRequest

`func NewSignTransactionRequest(to string, value string, chain string, ) *SignTransactionRequest`

NewSignTransactionRequest instantiates a new SignTransactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignTransactionRequestWithDefaults

`func NewSignTransactionRequestWithDefaults() *SignTransactionRequest`

NewSignTransactionRequestWithDefaults instantiates a new SignTransactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *SignTransactionRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SignTransactionRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SignTransactionRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetValue

`func (o *SignTransactionRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SignTransactionRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SignTransactionRequest) SetValue(v string)`

SetValue sets Value field to given value.


### GetChain

`func (o *SignTransactionRequest) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *SignTransactionRequest) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *SignTransactionRequest) SetChain(v string)`

SetChain sets Chain field to given value.


### GetData

`func (o *SignTransactionRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SignTransactionRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SignTransactionRequest) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *SignTransactionRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSigningKeyPath

`func (o *SignTransactionRequest) GetSigningKeyPath() string`

GetSigningKeyPath returns the SigningKeyPath field if non-nil, zero value otherwise.

### GetSigningKeyPathOk

`func (o *SignTransactionRequest) GetSigningKeyPathOk() (*string, bool)`

GetSigningKeyPathOk returns a tuple with the SigningKeyPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKeyPath

`func (o *SignTransactionRequest) SetSigningKeyPath(v string)`

SetSigningKeyPath sets SigningKeyPath field to given value.

### HasSigningKeyPath

`func (o *SignTransactionRequest) HasSigningKeyPath() bool`

HasSigningKeyPath returns a boolean if a field has been set.

### GetNonce

`func (o *SignTransactionRequest) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *SignTransactionRequest) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *SignTransactionRequest) SetNonce(v int32)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *SignTransactionRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetGasPrice

`func (o *SignTransactionRequest) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *SignTransactionRequest) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *SignTransactionRequest) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *SignTransactionRequest) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetGasLimit

`func (o *SignTransactionRequest) GetGasLimit() int32`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *SignTransactionRequest) GetGasLimitOk() (*int32, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *SignTransactionRequest) SetGasLimit(v int32)`

SetGasLimit sets GasLimit field to given value.

### HasGasLimit

`func (o *SignTransactionRequest) HasGasLimit() bool`

HasGasLimit returns a boolean if a field has been set.

### GetMaxFeePerGas

`func (o *SignTransactionRequest) GetMaxFeePerGas() string`

GetMaxFeePerGas returns the MaxFeePerGas field if non-nil, zero value otherwise.

### GetMaxFeePerGasOk

`func (o *SignTransactionRequest) GetMaxFeePerGasOk() (*string, bool)`

GetMaxFeePerGasOk returns a tuple with the MaxFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeePerGas

`func (o *SignTransactionRequest) SetMaxFeePerGas(v string)`

SetMaxFeePerGas sets MaxFeePerGas field to given value.

### HasMaxFeePerGas

`func (o *SignTransactionRequest) HasMaxFeePerGas() bool`

HasMaxFeePerGas returns a boolean if a field has been set.

### GetMaxPriorityFeePerGas

`func (o *SignTransactionRequest) GetMaxPriorityFeePerGas() string`

GetMaxPriorityFeePerGas returns the MaxPriorityFeePerGas field if non-nil, zero value otherwise.

### GetMaxPriorityFeePerGasOk

`func (o *SignTransactionRequest) GetMaxPriorityFeePerGasOk() (*string, bool)`

GetMaxPriorityFeePerGasOk returns a tuple with the MaxPriorityFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriorityFeePerGas

`func (o *SignTransactionRequest) SetMaxPriorityFeePerGas(v string)`

SetMaxPriorityFeePerGas sets MaxPriorityFeePerGas field to given value.

### HasMaxPriorityFeePerGas

`func (o *SignTransactionRequest) HasMaxPriorityFeePerGas() bool`

HasMaxPriorityFeePerGas returns a boolean if a field has been set.

### GetSimulateFirst

`func (o *SignTransactionRequest) GetSimulateFirst() bool`

GetSimulateFirst returns the SimulateFirst field if non-nil, zero value otherwise.

### GetSimulateFirstOk

`func (o *SignTransactionRequest) GetSimulateFirstOk() (*bool, bool)`

GetSimulateFirstOk returns a tuple with the SimulateFirst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulateFirst

`func (o *SignTransactionRequest) SetSimulateFirst(v bool)`

SetSimulateFirst sets SimulateFirst field to given value.

### HasSimulateFirst

`func (o *SignTransactionRequest) HasSimulateFirst() bool`

HasSimulateFirst returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


