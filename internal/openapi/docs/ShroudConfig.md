# ShroudConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PiiPolicy** | Pointer to **string** | How PII detections are handled | [optional] [default to "redact"]
**InjectionThreshold** | Pointer to **float32** | Prompt injection score threshold (0.0–1.0). Requests above are blocked | [optional] [default to 0.7]
**ContextInjectionThreshold** | Pointer to **float32** | Context injection score threshold (0.0–1.0) | [optional] [default to 0.7]
**AllowedProviders** | Pointer to **[]string** | LLM providers this agent may use (empty &#x3D; all) | [optional] 
**AllowedModels** | Pointer to **[]string** | Specific models allowed (empty &#x3D; all) | [optional] 
**DeniedModels** | Pointer to **[]string** | Models explicitly blocked | [optional] 
**MaxTokensPerRequest** | Pointer to **int32** | Maximum input tokens per request | [optional] 
**MaxRequestsPerMinute** | Pointer to **int32** | Rate limit (requests per minute) | [optional] 
**MaxRequestsPerDay** | Pointer to **int32** | Rate limit (requests per day) | [optional] 
**DailyBudgetUsd** | Pointer to **float32** | Daily LLM spend cap in USD (0 &#x3D; unlimited) | [optional] 
**EnableSecretRedaction** | Pointer to **bool** | Whether vault secrets are redacted from prompts/responses | [optional] [default to true]
**EnableResponseFiltering** | Pointer to **bool** | Whether response credential scanning is active | [optional] [default to true]
**UnicodeNormalization** | Pointer to [**UnicodeNormalizationConfig**](UnicodeNormalizationConfig.md) |  | [optional] 
**CommandInjectionDetection** | Pointer to [**CommandInjectionConfig**](CommandInjectionConfig.md) |  | [optional] 
**SocialEngineeringDetection** | Pointer to [**SocialEngineeringConfig**](SocialEngineeringConfig.md) |  | [optional] 
**EncodingDetection** | Pointer to [**EncodingDetectionConfig**](EncodingDetectionConfig.md) |  | [optional] 
**NetworkDetection** | Pointer to [**NetworkDetectionConfig**](NetworkDetectionConfig.md) |  | [optional] 
**FilesystemDetection** | Pointer to [**FilesystemDetectionConfig**](FilesystemDetectionConfig.md) |  | [optional] 
**SanitizationMode** | Pointer to **string** | Global behavior when threats are detected (block&#x3D;reject, surgical&#x3D;remove only malicious parts, log_only&#x3D;audit without action) | [optional] [default to "block"]
**ThreatLogging** | Pointer to **bool** | Whether to log all detected threats to audit (even when action is allow/warn) | [optional] [default to true]

## Methods

### NewShroudConfig

`func NewShroudConfig() *ShroudConfig`

NewShroudConfig instantiates a new ShroudConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShroudConfigWithDefaults

`func NewShroudConfigWithDefaults() *ShroudConfig`

NewShroudConfigWithDefaults instantiates a new ShroudConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPiiPolicy

`func (o *ShroudConfig) GetPiiPolicy() string`

GetPiiPolicy returns the PiiPolicy field if non-nil, zero value otherwise.

### GetPiiPolicyOk

`func (o *ShroudConfig) GetPiiPolicyOk() (*string, bool)`

GetPiiPolicyOk returns a tuple with the PiiPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiiPolicy

`func (o *ShroudConfig) SetPiiPolicy(v string)`

SetPiiPolicy sets PiiPolicy field to given value.

### HasPiiPolicy

`func (o *ShroudConfig) HasPiiPolicy() bool`

HasPiiPolicy returns a boolean if a field has been set.

### GetInjectionThreshold

`func (o *ShroudConfig) GetInjectionThreshold() float32`

GetInjectionThreshold returns the InjectionThreshold field if non-nil, zero value otherwise.

### GetInjectionThresholdOk

`func (o *ShroudConfig) GetInjectionThresholdOk() (*float32, bool)`

GetInjectionThresholdOk returns a tuple with the InjectionThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectionThreshold

`func (o *ShroudConfig) SetInjectionThreshold(v float32)`

SetInjectionThreshold sets InjectionThreshold field to given value.

### HasInjectionThreshold

`func (o *ShroudConfig) HasInjectionThreshold() bool`

HasInjectionThreshold returns a boolean if a field has been set.

### GetContextInjectionThreshold

`func (o *ShroudConfig) GetContextInjectionThreshold() float32`

GetContextInjectionThreshold returns the ContextInjectionThreshold field if non-nil, zero value otherwise.

### GetContextInjectionThresholdOk

`func (o *ShroudConfig) GetContextInjectionThresholdOk() (*float32, bool)`

GetContextInjectionThresholdOk returns a tuple with the ContextInjectionThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextInjectionThreshold

`func (o *ShroudConfig) SetContextInjectionThreshold(v float32)`

SetContextInjectionThreshold sets ContextInjectionThreshold field to given value.

### HasContextInjectionThreshold

`func (o *ShroudConfig) HasContextInjectionThreshold() bool`

HasContextInjectionThreshold returns a boolean if a field has been set.

### GetAllowedProviders

`func (o *ShroudConfig) GetAllowedProviders() []string`

GetAllowedProviders returns the AllowedProviders field if non-nil, zero value otherwise.

### GetAllowedProvidersOk

`func (o *ShroudConfig) GetAllowedProvidersOk() (*[]string, bool)`

GetAllowedProvidersOk returns a tuple with the AllowedProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedProviders

`func (o *ShroudConfig) SetAllowedProviders(v []string)`

SetAllowedProviders sets AllowedProviders field to given value.

### HasAllowedProviders

`func (o *ShroudConfig) HasAllowedProviders() bool`

HasAllowedProviders returns a boolean if a field has been set.

### GetAllowedModels

`func (o *ShroudConfig) GetAllowedModels() []string`

GetAllowedModels returns the AllowedModels field if non-nil, zero value otherwise.

### GetAllowedModelsOk

`func (o *ShroudConfig) GetAllowedModelsOk() (*[]string, bool)`

GetAllowedModelsOk returns a tuple with the AllowedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedModels

`func (o *ShroudConfig) SetAllowedModels(v []string)`

SetAllowedModels sets AllowedModels field to given value.

### HasAllowedModels

`func (o *ShroudConfig) HasAllowedModels() bool`

HasAllowedModels returns a boolean if a field has been set.

### GetDeniedModels

`func (o *ShroudConfig) GetDeniedModels() []string`

GetDeniedModels returns the DeniedModels field if non-nil, zero value otherwise.

### GetDeniedModelsOk

`func (o *ShroudConfig) GetDeniedModelsOk() (*[]string, bool)`

GetDeniedModelsOk returns a tuple with the DeniedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedModels

`func (o *ShroudConfig) SetDeniedModels(v []string)`

SetDeniedModels sets DeniedModels field to given value.

### HasDeniedModels

`func (o *ShroudConfig) HasDeniedModels() bool`

HasDeniedModels returns a boolean if a field has been set.

### GetMaxTokensPerRequest

`func (o *ShroudConfig) GetMaxTokensPerRequest() int32`

GetMaxTokensPerRequest returns the MaxTokensPerRequest field if non-nil, zero value otherwise.

### GetMaxTokensPerRequestOk

`func (o *ShroudConfig) GetMaxTokensPerRequestOk() (*int32, bool)`

GetMaxTokensPerRequestOk returns a tuple with the MaxTokensPerRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokensPerRequest

`func (o *ShroudConfig) SetMaxTokensPerRequest(v int32)`

SetMaxTokensPerRequest sets MaxTokensPerRequest field to given value.

### HasMaxTokensPerRequest

`func (o *ShroudConfig) HasMaxTokensPerRequest() bool`

HasMaxTokensPerRequest returns a boolean if a field has been set.

### GetMaxRequestsPerMinute

`func (o *ShroudConfig) GetMaxRequestsPerMinute() int32`

GetMaxRequestsPerMinute returns the MaxRequestsPerMinute field if non-nil, zero value otherwise.

### GetMaxRequestsPerMinuteOk

`func (o *ShroudConfig) GetMaxRequestsPerMinuteOk() (*int32, bool)`

GetMaxRequestsPerMinuteOk returns a tuple with the MaxRequestsPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRequestsPerMinute

`func (o *ShroudConfig) SetMaxRequestsPerMinute(v int32)`

SetMaxRequestsPerMinute sets MaxRequestsPerMinute field to given value.

### HasMaxRequestsPerMinute

`func (o *ShroudConfig) HasMaxRequestsPerMinute() bool`

HasMaxRequestsPerMinute returns a boolean if a field has been set.

### GetMaxRequestsPerDay

`func (o *ShroudConfig) GetMaxRequestsPerDay() int32`

GetMaxRequestsPerDay returns the MaxRequestsPerDay field if non-nil, zero value otherwise.

### GetMaxRequestsPerDayOk

`func (o *ShroudConfig) GetMaxRequestsPerDayOk() (*int32, bool)`

GetMaxRequestsPerDayOk returns a tuple with the MaxRequestsPerDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRequestsPerDay

`func (o *ShroudConfig) SetMaxRequestsPerDay(v int32)`

SetMaxRequestsPerDay sets MaxRequestsPerDay field to given value.

### HasMaxRequestsPerDay

`func (o *ShroudConfig) HasMaxRequestsPerDay() bool`

HasMaxRequestsPerDay returns a boolean if a field has been set.

### GetDailyBudgetUsd

`func (o *ShroudConfig) GetDailyBudgetUsd() float32`

GetDailyBudgetUsd returns the DailyBudgetUsd field if non-nil, zero value otherwise.

### GetDailyBudgetUsdOk

`func (o *ShroudConfig) GetDailyBudgetUsdOk() (*float32, bool)`

GetDailyBudgetUsdOk returns a tuple with the DailyBudgetUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyBudgetUsd

`func (o *ShroudConfig) SetDailyBudgetUsd(v float32)`

SetDailyBudgetUsd sets DailyBudgetUsd field to given value.

### HasDailyBudgetUsd

`func (o *ShroudConfig) HasDailyBudgetUsd() bool`

HasDailyBudgetUsd returns a boolean if a field has been set.

### GetEnableSecretRedaction

`func (o *ShroudConfig) GetEnableSecretRedaction() bool`

GetEnableSecretRedaction returns the EnableSecretRedaction field if non-nil, zero value otherwise.

### GetEnableSecretRedactionOk

`func (o *ShroudConfig) GetEnableSecretRedactionOk() (*bool, bool)`

GetEnableSecretRedactionOk returns a tuple with the EnableSecretRedaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSecretRedaction

`func (o *ShroudConfig) SetEnableSecretRedaction(v bool)`

SetEnableSecretRedaction sets EnableSecretRedaction field to given value.

### HasEnableSecretRedaction

`func (o *ShroudConfig) HasEnableSecretRedaction() bool`

HasEnableSecretRedaction returns a boolean if a field has been set.

### GetEnableResponseFiltering

`func (o *ShroudConfig) GetEnableResponseFiltering() bool`

GetEnableResponseFiltering returns the EnableResponseFiltering field if non-nil, zero value otherwise.

### GetEnableResponseFilteringOk

`func (o *ShroudConfig) GetEnableResponseFilteringOk() (*bool, bool)`

GetEnableResponseFilteringOk returns a tuple with the EnableResponseFiltering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableResponseFiltering

`func (o *ShroudConfig) SetEnableResponseFiltering(v bool)`

SetEnableResponseFiltering sets EnableResponseFiltering field to given value.

### HasEnableResponseFiltering

`func (o *ShroudConfig) HasEnableResponseFiltering() bool`

HasEnableResponseFiltering returns a boolean if a field has been set.

### GetUnicodeNormalization

`func (o *ShroudConfig) GetUnicodeNormalization() UnicodeNormalizationConfig`

GetUnicodeNormalization returns the UnicodeNormalization field if non-nil, zero value otherwise.

### GetUnicodeNormalizationOk

`func (o *ShroudConfig) GetUnicodeNormalizationOk() (*UnicodeNormalizationConfig, bool)`

GetUnicodeNormalizationOk returns a tuple with the UnicodeNormalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnicodeNormalization

`func (o *ShroudConfig) SetUnicodeNormalization(v UnicodeNormalizationConfig)`

SetUnicodeNormalization sets UnicodeNormalization field to given value.

### HasUnicodeNormalization

`func (o *ShroudConfig) HasUnicodeNormalization() bool`

HasUnicodeNormalization returns a boolean if a field has been set.

### GetCommandInjectionDetection

`func (o *ShroudConfig) GetCommandInjectionDetection() CommandInjectionConfig`

GetCommandInjectionDetection returns the CommandInjectionDetection field if non-nil, zero value otherwise.

### GetCommandInjectionDetectionOk

`func (o *ShroudConfig) GetCommandInjectionDetectionOk() (*CommandInjectionConfig, bool)`

GetCommandInjectionDetectionOk returns a tuple with the CommandInjectionDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandInjectionDetection

`func (o *ShroudConfig) SetCommandInjectionDetection(v CommandInjectionConfig)`

SetCommandInjectionDetection sets CommandInjectionDetection field to given value.

### HasCommandInjectionDetection

`func (o *ShroudConfig) HasCommandInjectionDetection() bool`

HasCommandInjectionDetection returns a boolean if a field has been set.

### GetSocialEngineeringDetection

`func (o *ShroudConfig) GetSocialEngineeringDetection() SocialEngineeringConfig`

GetSocialEngineeringDetection returns the SocialEngineeringDetection field if non-nil, zero value otherwise.

### GetSocialEngineeringDetectionOk

`func (o *ShroudConfig) GetSocialEngineeringDetectionOk() (*SocialEngineeringConfig, bool)`

GetSocialEngineeringDetectionOk returns a tuple with the SocialEngineeringDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialEngineeringDetection

`func (o *ShroudConfig) SetSocialEngineeringDetection(v SocialEngineeringConfig)`

SetSocialEngineeringDetection sets SocialEngineeringDetection field to given value.

### HasSocialEngineeringDetection

`func (o *ShroudConfig) HasSocialEngineeringDetection() bool`

HasSocialEngineeringDetection returns a boolean if a field has been set.

### GetEncodingDetection

`func (o *ShroudConfig) GetEncodingDetection() EncodingDetectionConfig`

GetEncodingDetection returns the EncodingDetection field if non-nil, zero value otherwise.

### GetEncodingDetectionOk

`func (o *ShroudConfig) GetEncodingDetectionOk() (*EncodingDetectionConfig, bool)`

GetEncodingDetectionOk returns a tuple with the EncodingDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingDetection

`func (o *ShroudConfig) SetEncodingDetection(v EncodingDetectionConfig)`

SetEncodingDetection sets EncodingDetection field to given value.

### HasEncodingDetection

`func (o *ShroudConfig) HasEncodingDetection() bool`

HasEncodingDetection returns a boolean if a field has been set.

### GetNetworkDetection

`func (o *ShroudConfig) GetNetworkDetection() NetworkDetectionConfig`

GetNetworkDetection returns the NetworkDetection field if non-nil, zero value otherwise.

### GetNetworkDetectionOk

`func (o *ShroudConfig) GetNetworkDetectionOk() (*NetworkDetectionConfig, bool)`

GetNetworkDetectionOk returns a tuple with the NetworkDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDetection

`func (o *ShroudConfig) SetNetworkDetection(v NetworkDetectionConfig)`

SetNetworkDetection sets NetworkDetection field to given value.

### HasNetworkDetection

`func (o *ShroudConfig) HasNetworkDetection() bool`

HasNetworkDetection returns a boolean if a field has been set.

### GetFilesystemDetection

`func (o *ShroudConfig) GetFilesystemDetection() FilesystemDetectionConfig`

GetFilesystemDetection returns the FilesystemDetection field if non-nil, zero value otherwise.

### GetFilesystemDetectionOk

`func (o *ShroudConfig) GetFilesystemDetectionOk() (*FilesystemDetectionConfig, bool)`

GetFilesystemDetectionOk returns a tuple with the FilesystemDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemDetection

`func (o *ShroudConfig) SetFilesystemDetection(v FilesystemDetectionConfig)`

SetFilesystemDetection sets FilesystemDetection field to given value.

### HasFilesystemDetection

`func (o *ShroudConfig) HasFilesystemDetection() bool`

HasFilesystemDetection returns a boolean if a field has been set.

### GetSanitizationMode

`func (o *ShroudConfig) GetSanitizationMode() string`

GetSanitizationMode returns the SanitizationMode field if non-nil, zero value otherwise.

### GetSanitizationModeOk

`func (o *ShroudConfig) GetSanitizationModeOk() (*string, bool)`

GetSanitizationModeOk returns a tuple with the SanitizationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanitizationMode

`func (o *ShroudConfig) SetSanitizationMode(v string)`

SetSanitizationMode sets SanitizationMode field to given value.

### HasSanitizationMode

`func (o *ShroudConfig) HasSanitizationMode() bool`

HasSanitizationMode returns a boolean if a field has been set.

### GetThreatLogging

`func (o *ShroudConfig) GetThreatLogging() bool`

GetThreatLogging returns the ThreatLogging field if non-nil, zero value otherwise.

### GetThreatLoggingOk

`func (o *ShroudConfig) GetThreatLoggingOk() (*bool, bool)`

GetThreatLoggingOk returns a tuple with the ThreatLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLogging

`func (o *ShroudConfig) SetThreatLogging(v bool)`

SetThreatLogging sets ThreatLogging field to given value.

### HasThreatLogging

`func (o *ShroudConfig) HasThreatLogging() bool`

HasThreatLogging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


