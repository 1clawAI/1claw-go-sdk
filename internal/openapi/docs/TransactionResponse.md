# TransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**AgentId** | Pointer to **string** |  | [optional] 
**Chain** | Pointer to **string** |  | [optional] 
**ChainId** | Pointer to **int32** |  | [optional] 
**To** | Pointer to **string** |  | [optional] 
**ValueWei** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SignedTx** | Pointer to **string** | Raw signed transaction hex. On GET list and GET by id, this property is omitted by default (absent from the response). Pass &#x60;include_signed_tx&#x3D;true&#x60; on those endpoints to include it. Always present on the initial POST submit response.  | [optional] 
**TxHash** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**SignedAt** | Pointer to **time.Time** |  | [optional] 
**SimulationId** | Pointer to **string** |  | [optional] 
**SimulationStatus** | Pointer to **string** |  | [optional] 
**TenderlyDashboardUrl** | Pointer to **string** |  | [optional] 
**MaxFeePerGas** | Pointer to **string** |  | [optional] 
**MaxPriorityFeePerGas** | Pointer to **string** |  | [optional] 

## Methods

### NewTransactionResponse

`func NewTransactionResponse() *TransactionResponse`

NewTransactionResponse instantiates a new TransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionResponseWithDefaults

`func NewTransactionResponseWithDefaults() *TransactionResponse`

NewTransactionResponseWithDefaults instantiates a new TransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransactionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransactionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAgentId

`func (o *TransactionResponse) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *TransactionResponse) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *TransactionResponse) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *TransactionResponse) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetChain

`func (o *TransactionResponse) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *TransactionResponse) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *TransactionResponse) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *TransactionResponse) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetChainId

`func (o *TransactionResponse) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *TransactionResponse) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *TransactionResponse) SetChainId(v int32)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *TransactionResponse) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetTo

`func (o *TransactionResponse) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *TransactionResponse) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *TransactionResponse) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *TransactionResponse) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetValueWei

`func (o *TransactionResponse) GetValueWei() string`

GetValueWei returns the ValueWei field if non-nil, zero value otherwise.

### GetValueWeiOk

`func (o *TransactionResponse) GetValueWeiOk() (*string, bool)`

GetValueWeiOk returns a tuple with the ValueWei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueWei

`func (o *TransactionResponse) SetValueWei(v string)`

SetValueWei sets ValueWei field to given value.

### HasValueWei

`func (o *TransactionResponse) HasValueWei() bool`

HasValueWei returns a boolean if a field has been set.

### GetStatus

`func (o *TransactionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransactionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSignedTx

`func (o *TransactionResponse) GetSignedTx() string`

GetSignedTx returns the SignedTx field if non-nil, zero value otherwise.

### GetSignedTxOk

`func (o *TransactionResponse) GetSignedTxOk() (*string, bool)`

GetSignedTxOk returns a tuple with the SignedTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedTx

`func (o *TransactionResponse) SetSignedTx(v string)`

SetSignedTx sets SignedTx field to given value.

### HasSignedTx

`func (o *TransactionResponse) HasSignedTx() bool`

HasSignedTx returns a boolean if a field has been set.

### GetTxHash

`func (o *TransactionResponse) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *TransactionResponse) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *TransactionResponse) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *TransactionResponse) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.

### GetErrorMessage

`func (o *TransactionResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TransactionResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TransactionResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *TransactionResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TransactionResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TransactionResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TransactionResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TransactionResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetSignedAt

`func (o *TransactionResponse) GetSignedAt() time.Time`

GetSignedAt returns the SignedAt field if non-nil, zero value otherwise.

### GetSignedAtOk

`func (o *TransactionResponse) GetSignedAtOk() (*time.Time, bool)`

GetSignedAtOk returns a tuple with the SignedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedAt

`func (o *TransactionResponse) SetSignedAt(v time.Time)`

SetSignedAt sets SignedAt field to given value.

### HasSignedAt

`func (o *TransactionResponse) HasSignedAt() bool`

HasSignedAt returns a boolean if a field has been set.

### GetSimulationId

`func (o *TransactionResponse) GetSimulationId() string`

GetSimulationId returns the SimulationId field if non-nil, zero value otherwise.

### GetSimulationIdOk

`func (o *TransactionResponse) GetSimulationIdOk() (*string, bool)`

GetSimulationIdOk returns a tuple with the SimulationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulationId

`func (o *TransactionResponse) SetSimulationId(v string)`

SetSimulationId sets SimulationId field to given value.

### HasSimulationId

`func (o *TransactionResponse) HasSimulationId() bool`

HasSimulationId returns a boolean if a field has been set.

### GetSimulationStatus

`func (o *TransactionResponse) GetSimulationStatus() string`

GetSimulationStatus returns the SimulationStatus field if non-nil, zero value otherwise.

### GetSimulationStatusOk

`func (o *TransactionResponse) GetSimulationStatusOk() (*string, bool)`

GetSimulationStatusOk returns a tuple with the SimulationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulationStatus

`func (o *TransactionResponse) SetSimulationStatus(v string)`

SetSimulationStatus sets SimulationStatus field to given value.

### HasSimulationStatus

`func (o *TransactionResponse) HasSimulationStatus() bool`

HasSimulationStatus returns a boolean if a field has been set.

### GetTenderlyDashboardUrl

`func (o *TransactionResponse) GetTenderlyDashboardUrl() string`

GetTenderlyDashboardUrl returns the TenderlyDashboardUrl field if non-nil, zero value otherwise.

### GetTenderlyDashboardUrlOk

`func (o *TransactionResponse) GetTenderlyDashboardUrlOk() (*string, bool)`

GetTenderlyDashboardUrlOk returns a tuple with the TenderlyDashboardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenderlyDashboardUrl

`func (o *TransactionResponse) SetTenderlyDashboardUrl(v string)`

SetTenderlyDashboardUrl sets TenderlyDashboardUrl field to given value.

### HasTenderlyDashboardUrl

`func (o *TransactionResponse) HasTenderlyDashboardUrl() bool`

HasTenderlyDashboardUrl returns a boolean if a field has been set.

### GetMaxFeePerGas

`func (o *TransactionResponse) GetMaxFeePerGas() string`

GetMaxFeePerGas returns the MaxFeePerGas field if non-nil, zero value otherwise.

### GetMaxFeePerGasOk

`func (o *TransactionResponse) GetMaxFeePerGasOk() (*string, bool)`

GetMaxFeePerGasOk returns a tuple with the MaxFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeePerGas

`func (o *TransactionResponse) SetMaxFeePerGas(v string)`

SetMaxFeePerGas sets MaxFeePerGas field to given value.

### HasMaxFeePerGas

`func (o *TransactionResponse) HasMaxFeePerGas() bool`

HasMaxFeePerGas returns a boolean if a field has been set.

### GetMaxPriorityFeePerGas

`func (o *TransactionResponse) GetMaxPriorityFeePerGas() string`

GetMaxPriorityFeePerGas returns the MaxPriorityFeePerGas field if non-nil, zero value otherwise.

### GetMaxPriorityFeePerGasOk

`func (o *TransactionResponse) GetMaxPriorityFeePerGasOk() (*string, bool)`

GetMaxPriorityFeePerGasOk returns a tuple with the MaxPriorityFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriorityFeePerGas

`func (o *TransactionResponse) SetMaxPriorityFeePerGas(v string)`

SetMaxPriorityFeePerGas sets MaxPriorityFeePerGas field to given value.

### HasMaxPriorityFeePerGas

`func (o *TransactionResponse) HasMaxPriorityFeePerGas() bool`

HasMaxPriorityFeePerGas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


