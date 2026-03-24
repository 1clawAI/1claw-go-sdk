# CommandInjectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable command injection detection | [optional] [default to true]
**Action** | Pointer to **string** | Action when command injection is detected | [optional] [default to "block"]
**Patterns** | Pointer to **string** | Pattern strictness level | [optional] [default to "default"]
**CustomPatterns** | Pointer to **[]string** | Custom regex patterns for detection (only used when patterns&#x3D;custom) | [optional] 

## Methods

### NewCommandInjectionConfig

`func NewCommandInjectionConfig() *CommandInjectionConfig`

NewCommandInjectionConfig instantiates a new CommandInjectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandInjectionConfigWithDefaults

`func NewCommandInjectionConfigWithDefaults() *CommandInjectionConfig`

NewCommandInjectionConfigWithDefaults instantiates a new CommandInjectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *CommandInjectionConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CommandInjectionConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CommandInjectionConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CommandInjectionConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAction

`func (o *CommandInjectionConfig) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CommandInjectionConfig) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CommandInjectionConfig) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *CommandInjectionConfig) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetPatterns

`func (o *CommandInjectionConfig) GetPatterns() string`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *CommandInjectionConfig) GetPatternsOk() (*string, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *CommandInjectionConfig) SetPatterns(v string)`

SetPatterns sets Patterns field to given value.

### HasPatterns

`func (o *CommandInjectionConfig) HasPatterns() bool`

HasPatterns returns a boolean if a field has been set.

### GetCustomPatterns

`func (o *CommandInjectionConfig) GetCustomPatterns() []string`

GetCustomPatterns returns the CustomPatterns field if non-nil, zero value otherwise.

### GetCustomPatternsOk

`func (o *CommandInjectionConfig) GetCustomPatternsOk() (*[]string, bool)`

GetCustomPatternsOk returns a tuple with the CustomPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPatterns

`func (o *CommandInjectionConfig) SetCustomPatterns(v []string)`

SetCustomPatterns sets CustomPatterns field to given value.

### HasCustomPatterns

`func (o *CommandInjectionConfig) HasCustomPatterns() bool`

HasCustomPatterns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


