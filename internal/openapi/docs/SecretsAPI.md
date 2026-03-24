# \SecretsAPI

All URIs are relative to *http://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSecret**](SecretsAPI.md#DeleteSecret) | **Delete** /v1/vaults/{vault_id}/secrets/{path} | Delete a secret
[**GetSecret**](SecretsAPI.md#GetSecret) | **Get** /v1/vaults/{vault_id}/secrets/{path} | Retrieve a decrypted secret
[**ListSecrets**](SecretsAPI.md#ListSecrets) | **Get** /v1/vaults/{vault_id}/secrets | List secrets in a vault
[**PutSecret**](SecretsAPI.md#PutSecret) | **Put** /v1/vaults/{vault_id}/secrets/{path} | Store or update a secret



## DeleteSecret

> DeleteSecret(ctx, vaultId, path).Execute()

Delete a secret

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
	path := "path_example" // string | Secret path (e.g. \"db/credentials\")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecretsAPI.DeleteSecret(context.Background(), vaultId, path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.DeleteSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **string** |  | 
**path** | **string** | Secret path (e.g. \&quot;db/credentials\&quot;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecretRequest struct via the builder pattern


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


## GetSecret

> SecretResponse GetSecret(ctx, vaultId, path).Execute()

Retrieve a decrypted secret

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
	path := "path_example" // string | Secret path (e.g. \"db/credentials\")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.GetSecret(context.Background(), vaultId, path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.GetSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecret`: SecretResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.GetSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **string** |  | 
**path** | **string** | Secret path (e.g. \&quot;db/credentials\&quot;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SecretResponse**](SecretResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSecrets

> SecretListResponse ListSecrets(ctx, vaultId).Prefix(prefix).Execute()

List secrets in a vault

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
	prefix := "prefix_example" // string | Filter by path prefix (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.ListSecrets(context.Background(), vaultId).Prefix(prefix).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.ListSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecrets`: SecretListResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.ListSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefix** | **string** | Filter by path prefix | 

### Return type

[**SecretListResponse**](SecretListResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSecret

> SecretMetadataResponse PutSecret(ctx, vaultId, path).PutSecretRequest(putSecretRequest).Execute()

Store or update a secret

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
	path := "path_example" // string | Secret path (e.g. \"db/credentials\")
	putSecretRequest := *openapiclient.NewPutSecretRequest("Value_example") // PutSecretRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.PutSecret(context.Background(), vaultId, path).PutSecretRequest(putSecretRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.PutSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSecret`: SecretMetadataResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.PutSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **string** |  | 
**path** | **string** | Secret path (e.g. \&quot;db/credentials\&quot;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putSecretRequest** | [**PutSecretRequest**](PutSecretRequest.md) |  | 

### Return type

[**SecretMetadataResponse**](SecretMetadataResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

