# VaultListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vaults** | Pointer to [**[]VaultResponse**](VaultResponse.md) |  | [optional] 

## Methods

### NewVaultListResponse

`func NewVaultListResponse() *VaultListResponse`

NewVaultListResponse instantiates a new VaultListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultListResponseWithDefaults

`func NewVaultListResponseWithDefaults() *VaultListResponse`

NewVaultListResponseWithDefaults instantiates a new VaultListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVaults

`func (o *VaultListResponse) GetVaults() []VaultResponse`

GetVaults returns the Vaults field if non-nil, zero value otherwise.

### GetVaultsOk

`func (o *VaultListResponse) GetVaultsOk() (*[]VaultResponse, bool)`

GetVaultsOk returns a tuple with the Vaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaults

`func (o *VaultListResponse) SetVaults(v []VaultResponse)`

SetVaults sets Vaults field to given value.

### HasVaults

`func (o *VaultListResponse) HasVaults() bool`

HasVaults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


