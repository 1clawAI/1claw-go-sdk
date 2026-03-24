# AgentCreatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | [**AgentResponse**](AgentResponse.md) |  | 
**ApiKey** | Pointer to **string** | One-time API key (only present for api_key auth method) | [optional] 

## Methods

### NewAgentCreatedResponse

`func NewAgentCreatedResponse(agent AgentResponse, ) *AgentCreatedResponse`

NewAgentCreatedResponse instantiates a new AgentCreatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentCreatedResponseWithDefaults

`func NewAgentCreatedResponseWithDefaults() *AgentCreatedResponse`

NewAgentCreatedResponseWithDefaults instantiates a new AgentCreatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *AgentCreatedResponse) GetAgent() AgentResponse`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *AgentCreatedResponse) GetAgentOk() (*AgentResponse, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *AgentCreatedResponse) SetAgent(v AgentResponse)`

SetAgent sets Agent field to given value.


### GetApiKey

`func (o *AgentCreatedResponse) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *AgentCreatedResponse) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *AgentCreatedResponse) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *AgentCreatedResponse) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


