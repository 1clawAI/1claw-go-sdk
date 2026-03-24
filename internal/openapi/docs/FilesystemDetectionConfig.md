# FilesystemDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable filesystem path detection (disabled by default as it can be noisy) | [optional] [default to false]
**Action** | Pointer to **string** | Action when filesystem paths are detected | [optional] [default to "log"]
**BlockedPaths** | Pointer to **[]string** | Path patterns to block (e.g., /etc/passwd, ~/.ssh) | [optional] 

## Methods

### NewFilesystemDetectionConfig

`func NewFilesystemDetectionConfig() *FilesystemDetectionConfig`

NewFilesystemDetectionConfig instantiates a new FilesystemDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesystemDetectionConfigWithDefaults

`func NewFilesystemDetectionConfigWithDefaults() *FilesystemDetectionConfig`

NewFilesystemDetectionConfigWithDefaults instantiates a new FilesystemDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *FilesystemDetectionConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FilesystemDetectionConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FilesystemDetectionConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FilesystemDetectionConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAction

`func (o *FilesystemDetectionConfig) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FilesystemDetectionConfig) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FilesystemDetectionConfig) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *FilesystemDetectionConfig) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetBlockedPaths

`func (o *FilesystemDetectionConfig) GetBlockedPaths() []string`

GetBlockedPaths returns the BlockedPaths field if non-nil, zero value otherwise.

### GetBlockedPathsOk

`func (o *FilesystemDetectionConfig) GetBlockedPathsOk() (*[]string, bool)`

GetBlockedPathsOk returns a tuple with the BlockedPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedPaths

`func (o *FilesystemDetectionConfig) SetBlockedPaths(v []string)`

SetBlockedPaths sets BlockedPaths field to given value.

### HasBlockedPaths

`func (o *FilesystemDetectionConfig) HasBlockedPaths() bool`

HasBlockedPaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


