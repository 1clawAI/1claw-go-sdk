# SignTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignedTx** | Pointer to **string** | Raw signed transaction hex (always included) | [optional] 
**TxHash** | Pointer to **string** |  | [optional] 
**From** | Pointer to **string** | Derived sender address | [optional] 
**To** | Pointer to **string** |  | [optional] 
**Chain** | Pointer to **string** |  | [optional] 
**ChainId** | Pointer to **int32** |  | [optional] 
**Nonce** | Pointer to **int32** |  | [optional] 
**ValueWei** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SimulationId** | Pointer to **string** |  | [optional] 
**SimulationStatus** | Pointer to **string** |  | [optional] 
**MaxFeePerGas** | Pointer to **string** |  | [optional] 
**MaxPriorityFeePerGas** | Pointer to **string** |  | [optional] 

## Methods

### NewSignTransactionResponse

`func NewSignTransactionResponse() *SignTransactionResponse`

NewSignTransactionResponse instantiates a new SignTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignTransactionResponseWithDefaults

`func NewSignTransactionResponseWithDefaults() *SignTransactionResponse`

NewSignTransactionResponseWithDefaults instantiates a new SignTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignedTx

`func (o *SignTransactionResponse) GetSignedTx() string`

GetSignedTx returns the SignedTx field if non-nil, zero value otherwise.

### GetSignedTxOk

`func (o *SignTransactionResponse) GetSignedTxOk() (*string, bool)`

GetSignedTxOk returns a tuple with the SignedTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedTx

`func (o *SignTransactionResponse) SetSignedTx(v string)`

SetSignedTx sets SignedTx field to given value.

### HasSignedTx

`func (o *SignTransactionResponse) HasSignedTx() bool`

HasSignedTx returns a boolean if a field has been set.

### GetTxHash

`func (o *SignTransactionResponse) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *SignTransactionResponse) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *SignTransactionResponse) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *SignTransactionResponse) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.

### GetFrom

`func (o *SignTransactionResponse) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SignTransactionResponse) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SignTransactionResponse) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *SignTransactionResponse) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *SignTransactionResponse) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SignTransactionResponse) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SignTransactionResponse) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *SignTransactionResponse) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetChain

`func (o *SignTransactionResponse) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *SignTransactionResponse) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *SignTransactionResponse) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *SignTransactionResponse) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetChainId

`func (o *SignTransactionResponse) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *SignTransactionResponse) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *SignTransactionResponse) SetChainId(v int32)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *SignTransactionResponse) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetNonce

`func (o *SignTransactionResponse) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *SignTransactionResponse) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *SignTransactionResponse) SetNonce(v int32)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *SignTransactionResponse) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetValueWei

`func (o *SignTransactionResponse) GetValueWei() string`

GetValueWei returns the ValueWei field if non-nil, zero value otherwise.

### GetValueWeiOk

`func (o *SignTransactionResponse) GetValueWeiOk() (*string, bool)`

GetValueWeiOk returns a tuple with the ValueWei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueWei

`func (o *SignTransactionResponse) SetValueWei(v string)`

SetValueWei sets ValueWei field to given value.

### HasValueWei

`func (o *SignTransactionResponse) HasValueWei() bool`

HasValueWei returns a boolean if a field has been set.

### GetStatus

`func (o *SignTransactionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SignTransactionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SignTransactionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SignTransactionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSimulationId

`func (o *SignTransactionResponse) GetSimulationId() string`

GetSimulationId returns the SimulationId field if non-nil, zero value otherwise.

### GetSimulationIdOk

`func (o *SignTransactionResponse) GetSimulationIdOk() (*string, bool)`

GetSimulationIdOk returns a tuple with the SimulationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulationId

`func (o *SignTransactionResponse) SetSimulationId(v string)`

SetSimulationId sets SimulationId field to given value.

### HasSimulationId

`func (o *SignTransactionResponse) HasSimulationId() bool`

HasSimulationId returns a boolean if a field has been set.

### GetSimulationStatus

`func (o *SignTransactionResponse) GetSimulationStatus() string`

GetSimulationStatus returns the SimulationStatus field if non-nil, zero value otherwise.

### GetSimulationStatusOk

`func (o *SignTransactionResponse) GetSimulationStatusOk() (*string, bool)`

GetSimulationStatusOk returns a tuple with the SimulationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulationStatus

`func (o *SignTransactionResponse) SetSimulationStatus(v string)`

SetSimulationStatus sets SimulationStatus field to given value.

### HasSimulationStatus

`func (o *SignTransactionResponse) HasSimulationStatus() bool`

HasSimulationStatus returns a boolean if a field has been set.

### GetMaxFeePerGas

`func (o *SignTransactionResponse) GetMaxFeePerGas() string`

GetMaxFeePerGas returns the MaxFeePerGas field if non-nil, zero value otherwise.

### GetMaxFeePerGasOk

`func (o *SignTransactionResponse) GetMaxFeePerGasOk() (*string, bool)`

GetMaxFeePerGasOk returns a tuple with the MaxFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeePerGas

`func (o *SignTransactionResponse) SetMaxFeePerGas(v string)`

SetMaxFeePerGas sets MaxFeePerGas field to given value.

### HasMaxFeePerGas

`func (o *SignTransactionResponse) HasMaxFeePerGas() bool`

HasMaxFeePerGas returns a boolean if a field has been set.

### GetMaxPriorityFeePerGas

`func (o *SignTransactionResponse) GetMaxPriorityFeePerGas() string`

GetMaxPriorityFeePerGas returns the MaxPriorityFeePerGas field if non-nil, zero value otherwise.

### GetMaxPriorityFeePerGasOk

`func (o *SignTransactionResponse) GetMaxPriorityFeePerGasOk() (*string, bool)`

GetMaxPriorityFeePerGasOk returns a tuple with the MaxPriorityFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriorityFeePerGas

`func (o *SignTransactionResponse) SetMaxPriorityFeePerGas(v string)`

SetMaxPriorityFeePerGas sets MaxPriorityFeePerGas field to given value.

### HasMaxPriorityFeePerGas

`func (o *SignTransactionResponse) HasMaxPriorityFeePerGas() bool`

HasMaxPriorityFeePerGas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


