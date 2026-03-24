# NetworkDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable network/URL detection | [optional] [default to true]
**Action** | Pointer to **string** | Action when suspicious URLs are detected | [optional] [default to "warn"]
**BlockedDomains** | Pointer to **[]string** | Domains to always block (e.g., pastebin.com, ngrok.io) | [optional] 
**AllowedDomains** | Pointer to **[]string** | Domains to always allow (allowlist mode when non-empty) | [optional] 

## Methods

### NewNetworkDetectionConfig

`func NewNetworkDetectionConfig() *NetworkDetectionConfig`

NewNetworkDetectionConfig instantiates a new NetworkDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDetectionConfigWithDefaults

`func NewNetworkDetectionConfigWithDefaults() *NetworkDetectionConfig`

NewNetworkDetectionConfigWithDefaults instantiates a new NetworkDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *NetworkDetectionConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkDetectionConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkDetectionConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworkDetectionConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAction

`func (o *NetworkDetectionConfig) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *NetworkDetectionConfig) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *NetworkDetectionConfig) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *NetworkDetectionConfig) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetBlockedDomains

`func (o *NetworkDetectionConfig) GetBlockedDomains() []string`

GetBlockedDomains returns the BlockedDomains field if non-nil, zero value otherwise.

### GetBlockedDomainsOk

`func (o *NetworkDetectionConfig) GetBlockedDomainsOk() (*[]string, bool)`

GetBlockedDomainsOk returns a tuple with the BlockedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedDomains

`func (o *NetworkDetectionConfig) SetBlockedDomains(v []string)`

SetBlockedDomains sets BlockedDomains field to given value.

### HasBlockedDomains

`func (o *NetworkDetectionConfig) HasBlockedDomains() bool`

HasBlockedDomains returns a boolean if a field has been set.

### GetAllowedDomains

`func (o *NetworkDetectionConfig) GetAllowedDomains() []string`

GetAllowedDomains returns the AllowedDomains field if non-nil, zero value otherwise.

### GetAllowedDomainsOk

`func (o *NetworkDetectionConfig) GetAllowedDomainsOk() (*[]string, bool)`

GetAllowedDomainsOk returns a tuple with the AllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomains

`func (o *NetworkDetectionConfig) SetAllowedDomains(v []string)`

SetAllowedDomains sets AllowedDomains field to given value.

### HasAllowedDomains

`func (o *NetworkDetectionConfig) HasAllowedDomains() bool`

HasAllowedDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


