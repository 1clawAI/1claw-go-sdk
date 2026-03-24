# UsageHistoryResponseEventsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**PrincipalType** | Pointer to **string** |  | [optional] 
**PrincipalId** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 
**PriceUsd** | Pointer to **string** |  | [optional] 
**IsPaid** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUsageHistoryResponseEventsInner

`func NewUsageHistoryResponseEventsInner() *UsageHistoryResponseEventsInner`

NewUsageHistoryResponseEventsInner instantiates a new UsageHistoryResponseEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageHistoryResponseEventsInnerWithDefaults

`func NewUsageHistoryResponseEventsInnerWithDefaults() *UsageHistoryResponseEventsInner`

NewUsageHistoryResponseEventsInnerWithDefaults instantiates a new UsageHistoryResponseEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UsageHistoryResponseEventsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsageHistoryResponseEventsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsageHistoryResponseEventsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UsageHistoryResponseEventsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrincipalType

`func (o *UsageHistoryResponseEventsInner) GetPrincipalType() string`

GetPrincipalType returns the PrincipalType field if non-nil, zero value otherwise.

### GetPrincipalTypeOk

`func (o *UsageHistoryResponseEventsInner) GetPrincipalTypeOk() (*string, bool)`

GetPrincipalTypeOk returns a tuple with the PrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalType

`func (o *UsageHistoryResponseEventsInner) SetPrincipalType(v string)`

SetPrincipalType sets PrincipalType field to given value.

### HasPrincipalType

`func (o *UsageHistoryResponseEventsInner) HasPrincipalType() bool`

HasPrincipalType returns a boolean if a field has been set.

### GetPrincipalId

`func (o *UsageHistoryResponseEventsInner) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *UsageHistoryResponseEventsInner) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *UsageHistoryResponseEventsInner) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.

### HasPrincipalId

`func (o *UsageHistoryResponseEventsInner) HasPrincipalId() bool`

HasPrincipalId returns a boolean if a field has been set.

### GetMethod

`func (o *UsageHistoryResponseEventsInner) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *UsageHistoryResponseEventsInner) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *UsageHistoryResponseEventsInner) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *UsageHistoryResponseEventsInner) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetEndpoint

`func (o *UsageHistoryResponseEventsInner) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *UsageHistoryResponseEventsInner) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *UsageHistoryResponseEventsInner) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *UsageHistoryResponseEventsInner) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetStatusCode

`func (o *UsageHistoryResponseEventsInner) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *UsageHistoryResponseEventsInner) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *UsageHistoryResponseEventsInner) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *UsageHistoryResponseEventsInner) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetPriceUsd

`func (o *UsageHistoryResponseEventsInner) GetPriceUsd() string`

GetPriceUsd returns the PriceUsd field if non-nil, zero value otherwise.

### GetPriceUsdOk

`func (o *UsageHistoryResponseEventsInner) GetPriceUsdOk() (*string, bool)`

GetPriceUsdOk returns a tuple with the PriceUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUsd

`func (o *UsageHistoryResponseEventsInner) SetPriceUsd(v string)`

SetPriceUsd sets PriceUsd field to given value.

### HasPriceUsd

`func (o *UsageHistoryResponseEventsInner) HasPriceUsd() bool`

HasPriceUsd returns a boolean if a field has been set.

### GetIsPaid

`func (o *UsageHistoryResponseEventsInner) GetIsPaid() bool`

GetIsPaid returns the IsPaid field if non-nil, zero value otherwise.

### GetIsPaidOk

`func (o *UsageHistoryResponseEventsInner) GetIsPaidOk() (*bool, bool)`

GetIsPaidOk returns a tuple with the IsPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaid

`func (o *UsageHistoryResponseEventsInner) SetIsPaid(v bool)`

SetIsPaid sets IsPaid field to given value.

### HasIsPaid

`func (o *UsageHistoryResponseEventsInner) HasIsPaid() bool`

HasIsPaid returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UsageHistoryResponseEventsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UsageHistoryResponseEventsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UsageHistoryResponseEventsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UsageHistoryResponseEventsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


