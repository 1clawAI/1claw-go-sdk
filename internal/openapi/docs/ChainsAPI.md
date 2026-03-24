# \ChainsAPI

All URIs are relative to *http://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminListChains**](ChainsAPI.md#AdminListChains) | **Get** /v1/admin/chains | List all chains including disabled (admin)
[**CreateChain**](ChainsAPI.md#CreateChain) | **Post** /v1/admin/chains | Add a chain (admin)
[**DeleteChain**](ChainsAPI.md#DeleteChain) | **Delete** /v1/admin/chains/{chain_id} | Remove a chain (admin)
[**GetChain**](ChainsAPI.md#GetChain) | **Get** /v1/chains/{identifier} | Get chain by name or ID
[**ListChains**](ChainsAPI.md#ListChains) | **Get** /v1/chains | List enabled chains
[**UpdateChain**](ChainsAPI.md#UpdateChain) | **Put** /v1/admin/chains/{chain_id} | Update a chain (admin)



## AdminListChains

> ChainListResponse AdminListChains(ctx).Execute()

List all chains including disabled (admin)

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
	resp, r, err := apiClient.ChainsAPI.AdminListChains(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChainsAPI.AdminListChains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminListChains`: ChainListResponse
	fmt.Fprintf(os.Stdout, "Response from `ChainsAPI.AdminListChains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminListChainsRequest struct via the builder pattern


### Return type

[**ChainListResponse**](ChainListResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateChain

> ChainResponse CreateChain(ctx).CreateChainRequest(createChainRequest).Execute()

Add a chain (admin)

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
	createChainRequest := *openapiclient.NewCreateChainRequest("Name_example", "DisplayName_example", int32(123)) // CreateChainRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChainsAPI.CreateChain(context.Background()).CreateChainRequest(createChainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChainsAPI.CreateChain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateChain`: ChainResponse
	fmt.Fprintf(os.Stdout, "Response from `ChainsAPI.CreateChain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateChainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createChainRequest** | [**CreateChainRequest**](CreateChainRequest.md) |  | 

### Return type

[**ChainResponse**](ChainResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteChain

> DeleteChain(ctx, chainId).Execute()

Remove a chain (admin)

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
	chainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChainsAPI.DeleteChain(context.Background(), chainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChainsAPI.DeleteChain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteChainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChain

> ChainResponse GetChain(ctx, identifier).Execute()

Get chain by name or ID

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
	identifier := "identifier_example" // string | Chain name (e.g. \"ethereum\") or numeric chain ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChainsAPI.GetChain(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChainsAPI.GetChain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChain`: ChainResponse
	fmt.Fprintf(os.Stdout, "Response from `ChainsAPI.GetChain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | Chain name (e.g. \&quot;ethereum\&quot;) or numeric chain ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChainResponse**](ChainResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChains

> ChainListResponse ListChains(ctx).Execute()

List enabled chains

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
	resp, r, err := apiClient.ChainsAPI.ListChains(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChainsAPI.ListChains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListChains`: ChainListResponse
	fmt.Fprintf(os.Stdout, "Response from `ChainsAPI.ListChains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListChainsRequest struct via the builder pattern


### Return type

[**ChainListResponse**](ChainListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChain

> ChainResponse UpdateChain(ctx, chainId).UpdateChainRequest(updateChainRequest).Execute()

Update a chain (admin)

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
	chainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateChainRequest := *openapiclient.NewUpdateChainRequest() // UpdateChainRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChainsAPI.UpdateChain(context.Background(), chainId).UpdateChainRequest(updateChainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChainsAPI.UpdateChain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateChain`: ChainResponse
	fmt.Fprintf(os.Stdout, "Response from `ChainsAPI.UpdateChain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateChainRequest** | [**UpdateChainRequest**](UpdateChainRequest.md) |  | 

### Return type

[**ChainResponse**](ChainResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

