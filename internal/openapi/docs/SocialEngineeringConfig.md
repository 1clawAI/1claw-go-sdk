# SocialEngineeringConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable social engineering detection | [optional] [default to true]
**Action** | Pointer to **string** | Action when manipulation attempts are detected | [optional] [default to "warn"]
**Sensitivity** | Pointer to **string** | Detection sensitivity level | [optional] [default to "medium"]

## Methods

### NewSocialEngineeringConfig

`func NewSocialEngineeringConfig() *SocialEngineeringConfig`

NewSocialEngineeringConfig instantiates a new SocialEngineeringConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialEngineeringConfigWithDefaults

`func NewSocialEngineeringConfigWithDefaults() *SocialEngineeringConfig`

NewSocialEngineeringConfigWithDefaults instantiates a new SocialEngineeringConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SocialEngineeringConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SocialEngineeringConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SocialEngineeringConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SocialEngineeringConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAction

`func (o *SocialEngineeringConfig) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SocialEngineeringConfig) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SocialEngineeringConfig) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SocialEngineeringConfig) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSensitivity

`func (o *SocialEngineeringConfig) GetSensitivity() string`

GetSensitivity returns the Sensitivity field if non-nil, zero value otherwise.

### GetSensitivityOk

`func (o *SocialEngineeringConfig) GetSensitivityOk() (*string, bool)`

GetSensitivityOk returns a tuple with the Sensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitivity

`func (o *SocialEngineeringConfig) SetSensitivity(v string)`

SetSensitivity sets Sensitivity field to given value.

### HasSensitivity

`func (o *SocialEngineeringConfig) HasSensitivity() bool`

HasSensitivity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


