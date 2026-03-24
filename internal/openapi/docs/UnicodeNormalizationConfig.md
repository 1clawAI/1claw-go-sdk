# UnicodeNormalizationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable Unicode normalization | [optional] [default to true]
**StripZeroWidth** | Pointer to **bool** | Remove zero-width and invisible Unicode characters | [optional] [default to true]
**NormalizeHomoglyphs** | Pointer to **bool** | Replace look-alike characters (e.g., Cyrillic а → Latin a) | [optional] [default to true]
**NormalizationForm** | Pointer to **string** | Unicode normalization form to apply | [optional] [default to "NFKC"]

## Methods

### NewUnicodeNormalizationConfig

`func NewUnicodeNormalizationConfig() *UnicodeNormalizationConfig`

NewUnicodeNormalizationConfig instantiates a new UnicodeNormalizationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnicodeNormalizationConfigWithDefaults

`func NewUnicodeNormalizationConfigWithDefaults() *UnicodeNormalizationConfig`

NewUnicodeNormalizationConfigWithDefaults instantiates a new UnicodeNormalizationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UnicodeNormalizationConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UnicodeNormalizationConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UnicodeNormalizationConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UnicodeNormalizationConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStripZeroWidth

`func (o *UnicodeNormalizationConfig) GetStripZeroWidth() bool`

GetStripZeroWidth returns the StripZeroWidth field if non-nil, zero value otherwise.

### GetStripZeroWidthOk

`func (o *UnicodeNormalizationConfig) GetStripZeroWidthOk() (*bool, bool)`

GetStripZeroWidthOk returns a tuple with the StripZeroWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripZeroWidth

`func (o *UnicodeNormalizationConfig) SetStripZeroWidth(v bool)`

SetStripZeroWidth sets StripZeroWidth field to given value.

### HasStripZeroWidth

`func (o *UnicodeNormalizationConfig) HasStripZeroWidth() bool`

HasStripZeroWidth returns a boolean if a field has been set.

### GetNormalizeHomoglyphs

`func (o *UnicodeNormalizationConfig) GetNormalizeHomoglyphs() bool`

GetNormalizeHomoglyphs returns the NormalizeHomoglyphs field if non-nil, zero value otherwise.

### GetNormalizeHomoglyphsOk

`func (o *UnicodeNormalizationConfig) GetNormalizeHomoglyphsOk() (*bool, bool)`

GetNormalizeHomoglyphsOk returns a tuple with the NormalizeHomoglyphs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizeHomoglyphs

`func (o *UnicodeNormalizationConfig) SetNormalizeHomoglyphs(v bool)`

SetNormalizeHomoglyphs sets NormalizeHomoglyphs field to given value.

### HasNormalizeHomoglyphs

`func (o *UnicodeNormalizationConfig) HasNormalizeHomoglyphs() bool`

HasNormalizeHomoglyphs returns a boolean if a field has been set.

### GetNormalizationForm

`func (o *UnicodeNormalizationConfig) GetNormalizationForm() string`

GetNormalizationForm returns the NormalizationForm field if non-nil, zero value otherwise.

### GetNormalizationFormOk

`func (o *UnicodeNormalizationConfig) GetNormalizationFormOk() (*string, bool)`

GetNormalizationFormOk returns a tuple with the NormalizationForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizationForm

`func (o *UnicodeNormalizationConfig) SetNormalizationForm(v string)`

SetNormalizationForm sets NormalizationForm field to given value.

### HasNormalizationForm

`func (o *UnicodeNormalizationConfig) HasNormalizationForm() bool`

HasNormalizationForm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


