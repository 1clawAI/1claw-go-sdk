# AgentTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **string** | Optional when using key-only auth (ocv_ keys auto-resolve agent) | [optional] 
**ApiKey** | **string** |  | 

## Methods

### NewAgentTokenRequest

`func NewAgentTokenRequest(apiKey string, ) *AgentTokenRequest`

NewAgentTokenRequest instantiates a new AgentTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentTokenRequestWithDefaults

`func NewAgentTokenRequestWithDefaults() *AgentTokenRequest`

NewAgentTokenRequestWithDefaults instantiates a new AgentTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *AgentTokenRequest) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *AgentTokenRequest) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *AgentTokenRequest) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *AgentTokenRequest) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetApiKey

`func (o *AgentTokenRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *AgentTokenRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *AgentTokenRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


