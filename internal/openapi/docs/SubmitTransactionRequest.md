# SubmitTransactionRequest

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
**Mode** | Pointer to **string** | Transaction mode | [optional] [default to "eoa"]

## Methods

### NewSubmitTransactionRequest

`func NewSubmitTransactionRequest(to string, value string, chain string, ) *SubmitTransactionRequest`

NewSubmitTransactionRequest instantiates a new SubmitTransactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitTransactionRequestWithDefaults

`func NewSubmitTransactionRequestWithDefaults() *SubmitTransactionRequest`

NewSubmitTransactionRequestWithDefaults instantiates a new SubmitTransactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *SubmitTransactionRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SubmitTransactionRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SubmitTransactionRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetValue

`func (o *SubmitTransactionRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SubmitTransactionRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SubmitTransactionRequest) SetValue(v string)`

SetValue sets Value field to given value.


### GetChain

`func (o *SubmitTransactionRequest) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *SubmitTransactionRequest) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *SubmitTransactionRequest) SetChain(v string)`

SetChain sets Chain field to given value.


### GetData

`func (o *SubmitTransactionRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubmitTransactionRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubmitTransactionRequest) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *SubmitTransactionRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSigningKeyPath

`func (o *SubmitTransactionRequest) GetSigningKeyPath() string`

GetSigningKeyPath returns the SigningKeyPath field if non-nil, zero value otherwise.

### GetSigningKeyPathOk

`func (o *SubmitTransactionRequest) GetSigningKeyPathOk() (*string, bool)`

GetSigningKeyPathOk returns a tuple with the SigningKeyPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKeyPath

`func (o *SubmitTransactionRequest) SetSigningKeyPath(v string)`

SetSigningKeyPath sets SigningKeyPath field to given value.

### HasSigningKeyPath

`func (o *SubmitTransactionRequest) HasSigningKeyPath() bool`

HasSigningKeyPath returns a boolean if a field has been set.

### GetNonce

`func (o *SubmitTransactionRequest) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *SubmitTransactionRequest) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *SubmitTransactionRequest) SetNonce(v int32)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *SubmitTransactionRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetGasPrice

`func (o *SubmitTransactionRequest) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *SubmitTransactionRequest) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *SubmitTransactionRequest) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *SubmitTransactionRequest) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetGasLimit

`func (o *SubmitTransactionRequest) GetGasLimit() int32`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *SubmitTransactionRequest) GetGasLimitOk() (*int32, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *SubmitTransactionRequest) SetGasLimit(v int32)`

SetGasLimit sets GasLimit field to given value.

### HasGasLimit

`func (o *SubmitTransactionRequest) HasGasLimit() bool`

HasGasLimit returns a boolean if a field has been set.

### GetMaxFeePerGas

`func (o *SubmitTransactionRequest) GetMaxFeePerGas() string`

GetMaxFeePerGas returns the MaxFeePerGas field if non-nil, zero value otherwise.

### GetMaxFeePerGasOk

`func (o *SubmitTransactionRequest) GetMaxFeePerGasOk() (*string, bool)`

GetMaxFeePerGasOk returns a tuple with the MaxFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeePerGas

`func (o *SubmitTransactionRequest) SetMaxFeePerGas(v string)`

SetMaxFeePerGas sets MaxFeePerGas field to given value.

### HasMaxFeePerGas

`func (o *SubmitTransactionRequest) HasMaxFeePerGas() bool`

HasMaxFeePerGas returns a boolean if a field has been set.

### GetMaxPriorityFeePerGas

`func (o *SubmitTransactionRequest) GetMaxPriorityFeePerGas() string`

GetMaxPriorityFeePerGas returns the MaxPriorityFeePerGas field if non-nil, zero value otherwise.

### GetMaxPriorityFeePerGasOk

`func (o *SubmitTransactionRequest) GetMaxPriorityFeePerGasOk() (*string, bool)`

GetMaxPriorityFeePerGasOk returns a tuple with the MaxPriorityFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriorityFeePerGas

`func (o *SubmitTransactionRequest) SetMaxPriorityFeePerGas(v string)`

SetMaxPriorityFeePerGas sets MaxPriorityFeePerGas field to given value.

### HasMaxPriorityFeePerGas

`func (o *SubmitTransactionRequest) HasMaxPriorityFeePerGas() bool`

HasMaxPriorityFeePerGas returns a boolean if a field has been set.

### GetSimulateFirst

`func (o *SubmitTransactionRequest) GetSimulateFirst() bool`

GetSimulateFirst returns the SimulateFirst field if non-nil, zero value otherwise.

### GetSimulateFirstOk

`func (o *SubmitTransactionRequest) GetSimulateFirstOk() (*bool, bool)`

GetSimulateFirstOk returns a tuple with the SimulateFirst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulateFirst

`func (o *SubmitTransactionRequest) SetSimulateFirst(v bool)`

SetSimulateFirst sets SimulateFirst field to given value.

### HasSimulateFirst

`func (o *SubmitTransactionRequest) HasSimulateFirst() bool`

HasSimulateFirst returns a boolean if a field has been set.

### GetMode

`func (o *SubmitTransactionRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SubmitTransactionRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SubmitTransactionRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *SubmitTransactionRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


