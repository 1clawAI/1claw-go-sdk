# EncodingDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable encoding detection | [optional] [default to true]
**Action** | Pointer to **string** | Action when obfuscated content is detected | [optional] [default to "warn"]
**DetectBase64** | Pointer to **bool** | Detect Base64-encoded content | [optional] [default to true]
**DetectHex** | Pointer to **bool** | Detect hex-encoded content (\\\\x41, 0x41) | [optional] [default to true]
**DetectUnicodeEscape** | Pointer to **bool** | Detect Unicode escape sequences (\\\\u0041) | [optional] [default to true]

## Methods

### NewEncodingDetectionConfig

`func NewEncodingDetectionConfig() *EncodingDetectionConfig`

NewEncodingDetectionConfig instantiates a new EncodingDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncodingDetectionConfigWithDefaults

`func NewEncodingDetectionConfigWithDefaults() *EncodingDetectionConfig`

NewEncodingDetectionConfigWithDefaults instantiates a new EncodingDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *EncodingDetectionConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EncodingDetectionConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EncodingDetectionConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *EncodingDetectionConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAction

`func (o *EncodingDetectionConfig) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *EncodingDetectionConfig) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *EncodingDetectionConfig) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *EncodingDetectionConfig) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDetectBase64

`func (o *EncodingDetectionConfig) GetDetectBase64() bool`

GetDetectBase64 returns the DetectBase64 field if non-nil, zero value otherwise.

### GetDetectBase64Ok

`func (o *EncodingDetectionConfig) GetDetectBase64Ok() (*bool, bool)`

GetDetectBase64Ok returns a tuple with the DetectBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectBase64

`func (o *EncodingDetectionConfig) SetDetectBase64(v bool)`

SetDetectBase64 sets DetectBase64 field to given value.

### HasDetectBase64

`func (o *EncodingDetectionConfig) HasDetectBase64() bool`

HasDetectBase64 returns a boolean if a field has been set.

### GetDetectHex

`func (o *EncodingDetectionConfig) GetDetectHex() bool`

GetDetectHex returns the DetectHex field if non-nil, zero value otherwise.

### GetDetectHexOk

`func (o *EncodingDetectionConfig) GetDetectHexOk() (*bool, bool)`

GetDetectHexOk returns a tuple with the DetectHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectHex

`func (o *EncodingDetectionConfig) SetDetectHex(v bool)`

SetDetectHex sets DetectHex field to given value.

### HasDetectHex

`func (o *EncodingDetectionConfig) HasDetectHex() bool`

HasDetectHex returns a boolean if a field has been set.

### GetDetectUnicodeEscape

`func (o *EncodingDetectionConfig) GetDetectUnicodeEscape() bool`

GetDetectUnicodeEscape returns the DetectUnicodeEscape field if non-nil, zero value otherwise.

### GetDetectUnicodeEscapeOk

`func (o *EncodingDetectionConfig) GetDetectUnicodeEscapeOk() (*bool, bool)`

GetDetectUnicodeEscapeOk returns a tuple with the DetectUnicodeEscape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectUnicodeEscape

`func (o *EncodingDetectionConfig) SetDetectUnicodeEscape(v bool)`

SetDetectUnicodeEscape sets DetectUnicodeEscape field to given value.

### HasDetectUnicodeEscape

`func (o *EncodingDetectionConfig) HasDetectUnicodeEscape() bool`

HasDetectUnicodeEscape returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


