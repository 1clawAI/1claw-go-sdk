# EnrollAgentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name for the new agent | 
**HumanEmail** | **string** | Email of the human who will receive the agent credentials | 
**Description** | Pointer to **string** | Optional agent description | [optional] 

## Methods

### NewEnrollAgentRequest

`func NewEnrollAgentRequest(name string, humanEmail string, ) *EnrollAgentRequest`

NewEnrollAgentRequest instantiates a new EnrollAgentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollAgentRequestWithDefaults

`func NewEnrollAgentRequestWithDefaults() *EnrollAgentRequest`

NewEnrollAgentRequestWithDefaults instantiates a new EnrollAgentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnrollAgentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnrollAgentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnrollAgentRequest) SetName(v string)`

SetName sets Name field to given value.


### GetHumanEmail

`func (o *EnrollAgentRequest) GetHumanEmail() string`

GetHumanEmail returns the HumanEmail field if non-nil, zero value otherwise.

### GetHumanEmailOk

`func (o *EnrollAgentRequest) GetHumanEmailOk() (*string, bool)`

GetHumanEmailOk returns a tuple with the HumanEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumanEmail

`func (o *EnrollAgentRequest) SetHumanEmail(v string)`

SetHumanEmail sets HumanEmail field to given value.


### GetDescription

`func (o *EnrollAgentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnrollAgentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnrollAgentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnrollAgentRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


