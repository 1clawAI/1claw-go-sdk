# \VaultsAPI

All URIs are relative to *http://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVault**](VaultsAPI.md#CreateVault) | **Post** /v1/vaults | Create a vault
[**DeleteVault**](VaultsAPI.md#DeleteVault) | **Delete** /v1/vaults/{vault_id} | Delete a vault
[**GetVault**](VaultsAPI.md#GetVault) | **Get** /v1/vaults/{vault_id} | Get vault details
[**ListVaults**](VaultsAPI.md#ListVaults) | **Get** /v1/vaults | List vaults



## CreateVault

> VaultResponse CreateVault(ctx).CreateVaultRequest(createVaultRequest).Execute()

Create a vault

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createVaultRequest := *openapiclient.NewCreateVaultRequest("Name_example") // CreateVaultRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultsAPI.CreateVault(context.Background()).CreateVaultRequest(createVaultRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultsAPI.CreateVault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVault`: VaultResponse
	fmt.Fprintf(os.Stdout, "Response from `VaultsAPI.CreateVault`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVaultRequest** | [**CreateVaultRequest**](CreateVaultRequest.md) |  | 

### Return type

[**VaultResponse**](VaultResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVault

> DeleteVault(ctx, vaultId).Execute()

Delete a vault

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	vaultId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VaultsAPI.DeleteVault(context.Background(), vaultId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultsAPI.DeleteVault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVault

> VaultResponse GetVault(ctx, vaultId).Execute()

Get vault details

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	vaultId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultsAPI.GetVault(context.Background(), vaultId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultsAPI.GetVault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVault`: VaultResponse
	fmt.Fprintf(os.Stdout, "Response from `VaultsAPI.GetVault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VaultResponse**](VaultResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVaults

> VaultListResponse ListVaults(ctx).Execute()

List vaults

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultsAPI.ListVaults(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultsAPI.ListVaults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVaults`: VaultListResponse
	fmt.Fprintf(os.Stdout, "Response from `VaultsAPI.ListVaults`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultsRequest struct via the builder pattern


### Return type

[**VaultListResponse**](VaultListResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

