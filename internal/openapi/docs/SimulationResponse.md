# SimulationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SimulationId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**GasUsed** | Pointer to **int32** |  | [optional] 
**GasEstimateUsd** | Pointer to **string** |  | [optional] 
**BalanceChanges** | Pointer to [**[]BalanceChange**](BalanceChange.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**ErrorHumanReadable** | Pointer to **string** |  | [optional] 
**RevertReason** | Pointer to **string** |  | [optional] 
**TenderlyDashboardUrl** | Pointer to **string** |  | [optional] 
**SimulatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSimulationResponse

`func NewSimulationResponse() *SimulationResponse`

NewSimulationResponse instantiates a new SimulationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulationResponseWithDefaults

`func NewSimulationResponseWithDefaults() *SimulationResponse`

NewSimulationResponseWithDefaults instantiates a new SimulationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSimulationId

`func (o *SimulationResponse) GetSimulationId() string`

GetSimulationId returns the SimulationId field if non-nil, zero value otherwise.

### GetSimulationIdOk

`func (o *SimulationResponse) GetSimulationIdOk() (*string, bool)`

GetSimulationIdOk returns a tuple with the SimulationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulationId

`func (o *SimulationResponse) SetSimulationId(v string)`

SetSimulationId sets SimulationId field to given value.

### HasSimulationId

`func (o *SimulationResponse) HasSimulationId() bool`

HasSimulationId returns a boolean if a field has been set.

### GetStatus

`func (o *SimulationResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SimulationResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SimulationResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SimulationResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetGasUsed

`func (o *SimulationResponse) GetGasUsed() int32`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *SimulationResponse) GetGasUsedOk() (*int32, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *SimulationResponse) SetGasUsed(v int32)`

SetGasUsed sets GasUsed field to given value.

### HasGasUsed

`func (o *SimulationResponse) HasGasUsed() bool`

HasGasUsed returns a boolean if a field has been set.

### GetGasEstimateUsd

`func (o *SimulationResponse) GetGasEstimateUsd() string`

GetGasEstimateUsd returns the GasEstimateUsd field if non-nil, zero value otherwise.

### GetGasEstimateUsdOk

`func (o *SimulationResponse) GetGasEstimateUsdOk() (*string, bool)`

GetGasEstimateUsdOk returns a tuple with the GasEstimateUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasEstimateUsd

`func (o *SimulationResponse) SetGasEstimateUsd(v string)`

SetGasEstimateUsd sets GasEstimateUsd field to given value.

### HasGasEstimateUsd

`func (o *SimulationResponse) HasGasEstimateUsd() bool`

HasGasEstimateUsd returns a boolean if a field has been set.

### GetBalanceChanges

`func (o *SimulationResponse) GetBalanceChanges() []BalanceChange`

GetBalanceChanges returns the BalanceChanges field if non-nil, zero value otherwise.

### GetBalanceChangesOk

`func (o *SimulationResponse) GetBalanceChangesOk() (*[]BalanceChange, bool)`

GetBalanceChangesOk returns a tuple with the BalanceChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceChanges

`func (o *SimulationResponse) SetBalanceChanges(v []BalanceChange)`

SetBalanceChanges sets BalanceChanges field to given value.

### HasBalanceChanges

`func (o *SimulationResponse) HasBalanceChanges() bool`

HasBalanceChanges returns a boolean if a field has been set.

### GetError

`func (o *SimulationResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SimulationResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SimulationResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SimulationResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorCode

`func (o *SimulationResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *SimulationResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *SimulationResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *SimulationResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorHumanReadable

`func (o *SimulationResponse) GetErrorHumanReadable() string`

GetErrorHumanReadable returns the ErrorHumanReadable field if non-nil, zero value otherwise.

### GetErrorHumanReadableOk

`func (o *SimulationResponse) GetErrorHumanReadableOk() (*string, bool)`

GetErrorHumanReadableOk returns a tuple with the ErrorHumanReadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorHumanReadable

`func (o *SimulationResponse) SetErrorHumanReadable(v string)`

SetErrorHumanReadable sets ErrorHumanReadable field to given value.

### HasErrorHumanReadable

`func (o *SimulationResponse) HasErrorHumanReadable() bool`

HasErrorHumanReadable returns a boolean if a field has been set.

### GetRevertReason

`func (o *SimulationResponse) GetRevertReason() string`

GetRevertReason returns the RevertReason field if non-nil, zero value otherwise.

### GetRevertReasonOk

`func (o *SimulationResponse) GetRevertReasonOk() (*string, bool)`

GetRevertReasonOk returns a tuple with the RevertReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevertReason

`func (o *SimulationResponse) SetRevertReason(v string)`

SetRevertReason sets RevertReason field to given value.

### HasRevertReason

`func (o *SimulationResponse) HasRevertReason() bool`

HasRevertReason returns a boolean if a field has been set.

### GetTenderlyDashboardUrl

`func (o *SimulationResponse) GetTenderlyDashboardUrl() string`

GetTenderlyDashboardUrl returns the TenderlyDashboardUrl field if non-nil, zero value otherwise.

### GetTenderlyDashboardUrlOk

`func (o *SimulationResponse) GetTenderlyDashboardUrlOk() (*string, bool)`

GetTenderlyDashboardUrlOk returns a tuple with the TenderlyDashboardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenderlyDashboardUrl

`func (o *SimulationResponse) SetTenderlyDashboardUrl(v string)`

SetTenderlyDashboardUrl sets TenderlyDashboardUrl field to given value.

### HasTenderlyDashboardUrl

`func (o *SimulationResponse) HasTenderlyDashboardUrl() bool`

HasTenderlyDashboardUrl returns a boolean if a field has been set.

### GetSimulatedAt

`func (o *SimulationResponse) GetSimulatedAt() time.Time`

GetSimulatedAt returns the SimulatedAt field if non-nil, zero value otherwise.

### GetSimulatedAtOk

`func (o *SimulationResponse) GetSimulatedAtOk() (*time.Time, bool)`

GetSimulatedAtOk returns a tuple with the SimulatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulatedAt

`func (o *SimulationResponse) SetSimulatedAt(v time.Time)`

SetSimulatedAt sets SimulatedAt field to given value.

### HasSimulatedAt

`func (o *SimulationResponse) HasSimulatedAt() bool`

HasSimulatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


