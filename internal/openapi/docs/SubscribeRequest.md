# SubscribeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tier** | **string** |  | 
**Interval** | **string** |  | 
**Trial** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewSubscribeRequest

`func NewSubscribeRequest(tier string, interval string, ) *SubscribeRequest`

NewSubscribeRequest instantiates a new SubscribeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscribeRequestWithDefaults

`func NewSubscribeRequestWithDefaults() *SubscribeRequest`

NewSubscribeRequestWithDefaults instantiates a new SubscribeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTier

`func (o *SubscribeRequest) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *SubscribeRequest) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *SubscribeRequest) SetTier(v string)`

SetTier sets Tier field to given value.


### GetInterval

`func (o *SubscribeRequest) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *SubscribeRequest) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *SubscribeRequest) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetTrial

`func (o *SubscribeRequest) GetTrial() bool`

GetTrial returns the Trial field if non-nil, zero value otherwise.

### GetTrialOk

`func (o *SubscribeRequest) GetTrialOk() (*bool, bool)`

GetTrialOk returns a tuple with the Trial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrial

`func (o *SubscribeRequest) SetTrial(v bool)`

SetTrial sets Trial field to given value.

### HasTrial

`func (o *SubscribeRequest) HasTrial() bool`

HasTrial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


