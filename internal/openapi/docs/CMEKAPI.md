# \CMEKAPI

All URIs are relative to *http://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableCmek**](CMEKAPI.md#DisableCmek) | **Delete** /v1/vaults/{vault_id}/cmek | Disable CMEK on a vault
[**EnableCmek**](CMEKAPI.md#EnableCmek) | **Post** /v1/vaults/{vault_id}/cmek | Enable CMEK on a vault
[**GetCmekRotationJob**](CMEKAPI.md#GetCmekRotationJob) | **Get** /v1/vaults/{vault_id}/cmek-rotate/{job_id} | Get CMEK rotation job status
[**RotateCmek**](CMEKAPI.md#RotateCmek) | **Post** /v1/vaults/{vault_id}/cmek-rotate | Start server-assisted CMEK key rotation



## DisableCmek

> VaultResponse DisableCmek(ctx, vaultId).Execute()

Disable CMEK on a vault



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
	resp, r, err := apiClient.CMEKAPI.DisableCmek(context.Background(), vaultId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CMEKAPI.DisableCmek``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableCmek`: VaultResponse
	fmt.Fprintf(os.Stdout, "Response from `CMEKAPI.DisableCmek`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableCmekRequest struct via the builder pattern


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


## EnableCmek

> VaultResponse EnableCmek(ctx, vaultId).EnableCmekRequest(enableCmekRequest).Execute()

Enable CMEK on a vault



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
	enableCmekRequest := *openapiclient.NewEnableCmekRequest("Fingerprint_example") // EnableCmekRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CMEKAPI.EnableCmek(context.Background(), vaultId).EnableCmekRequest(enableCmekRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CMEKAPI.EnableCmek``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableCmek`: VaultResponse
	fmt.Fprintf(os.Stdout, "Response from `CMEKAPI.EnableCmek`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableCmekRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableCmekRequest** | [**EnableCmekRequest**](EnableCmekRequest.md) |  | 

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


## GetCmekRotationJob

> CmekRotationJobResponse GetCmekRotationJob(ctx, vaultId, jobId).Execute()

Get CMEK rotation job status

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
	jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CMEKAPI.GetCmekRotationJob(context.Background(), vaultId, jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CMEKAPI.GetCmekRotationJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmekRotationJob`: CmekRotationJobResponse
	fmt.Fprintf(os.Stdout, "Response from `CMEKAPI.GetCmekRotationJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **string** |  | 
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmekRotationJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CmekRotationJobResponse**](CmekRotationJobResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateCmek

> CmekRotationJobResponse RotateCmek(ctx, vaultId).XCmekOldKey(xCmekOldKey).XCmekNewKey(xCmekNewKey).CmekRotateRequest(cmekRotateRequest).Execute()

Start server-assisted CMEK key rotation



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
	xCmekOldKey := "xCmekOldKey_example" // string | Base64-encoded old CMEK key (32 bytes)
	xCmekNewKey := "xCmekNewKey_example" // string | Base64-encoded new CMEK key (32 bytes)
	cmekRotateRequest := *openapiclient.NewCmekRotateRequest("NewFingerprint_example") // CmekRotateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CMEKAPI.RotateCmek(context.Background(), vaultId).XCmekOldKey(xCmekOldKey).XCmekNewKey(xCmekNewKey).CmekRotateRequest(cmekRotateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CMEKAPI.RotateCmek``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateCmek`: CmekRotationJobResponse
	fmt.Fprintf(os.Stdout, "Response from `CMEKAPI.RotateCmek`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateCmekRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCmekOldKey** | **string** | Base64-encoded old CMEK key (32 bytes) | 
 **xCmekNewKey** | **string** | Base64-encoded new CMEK key (32 bytes) | 
 **cmekRotateRequest** | [**CmekRotateRequest**](CmekRotateRequest.md) |  | 

### Return type

[**CmekRotationJobResponse**](CmekRotationJobResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

