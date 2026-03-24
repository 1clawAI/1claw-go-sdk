# CreateIpRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleType** | **string** |  | 
**Cidr** | **string** |  | 
**Label** | Pointer to **string** |  | [optional] 
**AppliesTo** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateIpRuleRequest

`func NewCreateIpRuleRequest(ruleType string, cidr string, ) *CreateIpRuleRequest`

NewCreateIpRuleRequest instantiates a new CreateIpRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIpRuleRequestWithDefaults

`func NewCreateIpRuleRequestWithDefaults() *CreateIpRuleRequest`

NewCreateIpRuleRequestWithDefaults instantiates a new CreateIpRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleType

`func (o *CreateIpRuleRequest) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *CreateIpRuleRequest) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *CreateIpRuleRequest) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.


### GetCidr

`func (o *CreateIpRuleRequest) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *CreateIpRuleRequest) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *CreateIpRuleRequest) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetLabel

`func (o *CreateIpRuleRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateIpRuleRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateIpRuleRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateIpRuleRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetAppliesTo

`func (o *CreateIpRuleRequest) GetAppliesTo() string`

GetAppliesTo returns the AppliesTo field if non-nil, zero value otherwise.

### GetAppliesToOk

`func (o *CreateIpRuleRequest) GetAppliesToOk() (*string, bool)`

GetAppliesToOk returns a tuple with the AppliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesTo

`func (o *CreateIpRuleRequest) SetAppliesTo(v string)`

SetAppliesTo sets AppliesTo field to given value.

### HasAppliesTo

`func (o *CreateIpRuleRequest) HasAppliesTo() bool`

HasAppliesTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


