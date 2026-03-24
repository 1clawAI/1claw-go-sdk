# EnrollAgentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **string** | UUID of the created agent (nil UUID when email not found — uniform response) | [optional] 
**Message** | Pointer to **string** | Status message (always generic to prevent email enumeration) | [optional] 

## Methods

### NewEnrollAgentResponse

`func NewEnrollAgentResponse() *EnrollAgentResponse`

NewEnrollAgentResponse instantiates a new EnrollAgentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollAgentResponseWithDefaults

`func NewEnrollAgentResponseWithDefaults() *EnrollAgentResponse`

NewEnrollAgentResponseWithDefaults instantiates a new EnrollAgentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *EnrollAgentResponse) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *EnrollAgentResponse) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *EnrollAgentResponse) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *EnrollAgentResponse) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetMessage

`func (o *EnrollAgentResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EnrollAgentResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EnrollAgentResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EnrollAgentResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


