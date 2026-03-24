# \TreasuryAPI

All URIs are relative to *http://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTreasurySigner**](TreasuryAPI.md#AddTreasurySigner) | **Post** /v1/treasury/{treasury_id}/signers | Add a signer to a treasury
[**ApproveTreasuryAccess**](TreasuryAPI.md#ApproveTreasuryAccess) | **Post** /v1/treasury/{treasury_id}/access-requests/{request_id}/approve | Approve an access request
[**CreateTreasury**](TreasuryAPI.md#CreateTreasury) | **Post** /v1/treasury | Create a treasury (Safe multisig)
[**DeleteTreasury**](TreasuryAPI.md#DeleteTreasury) | **Delete** /v1/treasury/{treasury_id} | Delete a treasury and its signers
[**DenyTreasuryAccess**](TreasuryAPI.md#DenyTreasuryAccess) | **Post** /v1/treasury/{treasury_id}/access-requests/{request_id}/deny | Deny an access request
[**GetTreasury**](TreasuryAPI.md#GetTreasury) | **Get** /v1/treasury/{treasury_id} | Get treasury details
[**ListTreasuries**](TreasuryAPI.md#ListTreasuries) | **Get** /v1/treasury | List treasuries
[**ListTreasuryAccessRequests**](TreasuryAPI.md#ListTreasuryAccessRequests) | **Get** /v1/treasury/{treasury_id}/access-requests | List access requests for a treasury
[**RemoveTreasurySigner**](TreasuryAPI.md#RemoveTreasurySigner) | **Delete** /v1/treasury/{treasury_id}/signers/{signer_id} | Remove a signer from a treasury
[**RequestTreasuryAccess**](TreasuryAPI.md#RequestTreasuryAccess) | **Post** /v1/treasury/{treasury_id}/access-requests | Request access to a treasury (agent-only)
[**UpdateTreasury**](TreasuryAPI.md#UpdateTreasury) | **Patch** /v1/treasury/{treasury_id} | Update treasury name and/or threshold



## AddTreasurySigner

> AddTreasurySigner(ctx, treasuryId).AddSignerRequest(addSignerRequest).Execute()

Add a signer to a treasury

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
	treasuryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	addSignerRequest := *openapiclient.NewAddSignerRequest("SignerType_example", "SignerId_example", "SignerAddress_example") // AddSignerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TreasuryAPI.AddTreasurySigner(context.Background(), treasuryId).AddSignerRequest(addSignerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TreasuryAPI.AddTreasurySigner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**treasuryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTreasurySignerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addSignerRequest** | [**AddSignerRequest**](AddSignerRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApproveTreasuryAccess

> ApproveTreasuryAccess(ctx, treasuryId, requestId).Execute()

Approve an access request

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
	treasuryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	requestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TreasuryAPI.ApproveTreasuryAccess(context.Background(), treasuryId, requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TreasuryAPI.ApproveTreasuryAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**treasuryId** | **string** |  | 
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveTreasuryAccessRequest struct via the builder pattern


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


## CreateTreasury

> TreasuryResponse CreateTreasury(ctx).CreateTreasuryRequest(createTreasuryRequest).Execute()

Create a treasury (Safe multisig)

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
	createTreasuryRequest := *openapiclient.NewCreateTreasuryRequest("Name_example", "SafeAddress_example") // CreateTreasuryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TreasuryAPI.CreateTreasury(context.Background()).CreateTreasuryRequest(createTreasuryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TreasuryAPI.CreateTreasury``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTreasury`: TreasuryResponse
	fmt.Fprintf(os.Stdout, "Response from `TreasuryAPI.CreateTreasury`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTreasuryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTreasuryRequest** | [**CreateTreasuryRequest**](CreateTreasuryRequest.md) |  | 

### Return type

[**TreasuryResponse**](TreasuryResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTreasury

> DeleteTreasury(ctx, treasuryId).Execute()

Delete a treasury and its signers

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
	treasuryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TreasuryAPI.DeleteTreasury(context.Background(), treasuryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TreasuryAPI.DeleteTreasury``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**treasuryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTreasuryRequest struct via the builder pattern


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


## DenyTreasuryAccess

> DenyTreasuryAccess(ctx, treasuryId, requestId).Execute()

Deny an access request

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
	treasuryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	requestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TreasuryAPI.DenyTreasuryAccess(context.Background(), treasuryId, requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TreasuryAPI.DenyTreasuryAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**treasuryId** | **string** |  | 
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDenyTreasuryAccessRequest struct via the builder pattern


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


## GetTreasury

> TreasuryResponse GetTreasury(ctx, treasuryId).Execute()

Get treasury details

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
	treasuryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TreasuryAPI.GetTreasury(context.Background(), treasuryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TreasuryAPI.GetTreasury``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTreasury`: TreasuryResponse
	fmt.Fprintf(os.Stdout, "Response from `TreasuryAPI.GetTreasury`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**treasuryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTreasuryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TreasuryResponse**](TreasuryResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTreasuries

> ListTreasuries200Response ListTreasuries(ctx).Execute()

List treasuries

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
	resp, r, err := apiClient.TreasuryAPI.ListTreasuries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TreasuryAPI.ListTreasuries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTreasuries`: ListTreasuries200Response
	fmt.Fprintf(os.Stdout, "Response from `TreasuryAPI.ListTreasuries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTreasuriesRequest struct via the builder pattern


### Return type

[**ListTreasuries200Response**](ListTreasuries200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTreasuryAccessRequests

> ListTreasuryAccessRequests200Response ListTreasuryAccessRequests(ctx, treasuryId).Execute()

List access requests for a treasury

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
	treasuryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TreasuryAPI.ListTreasuryAccessRequests(context.Background(), treasuryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TreasuryAPI.ListTreasuryAccessRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTreasuryAccessRequests`: ListTreasuryAccessRequests200Response
	fmt.Fprintf(os.Stdout, "Response from `TreasuryAPI.ListTreasuryAccessRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**treasuryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTreasuryAccessRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListTreasuryAccessRequests200Response**](ListTreasuryAccessRequests200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTreasurySigner

> RemoveTreasurySigner(ctx, treasuryId, signerId).Execute()

Remove a signer from a treasury

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
	treasuryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	signerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TreasuryAPI.RemoveTreasurySigner(context.Background(), treasuryId, signerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TreasuryAPI.RemoveTreasurySigner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**treasuryId** | **string** |  | 
**signerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTreasurySignerRequest struct via the builder pattern


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


## RequestTreasuryAccess

> AccessRequestResponse RequestTreasuryAccess(ctx, treasuryId).Execute()

Request access to a treasury (agent-only)

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
	treasuryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TreasuryAPI.RequestTreasuryAccess(context.Background(), treasuryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TreasuryAPI.RequestTreasuryAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestTreasuryAccess`: AccessRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `TreasuryAPI.RequestTreasuryAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**treasuryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestTreasuryAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessRequestResponse**](AccessRequestResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTreasury

> TreasuryResponse UpdateTreasury(ctx, treasuryId).UpdateTreasuryRequest(updateTreasuryRequest).Execute()

Update treasury name and/or threshold

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
	treasuryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateTreasuryRequest := *openapiclient.NewUpdateTreasuryRequest() // UpdateTreasuryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TreasuryAPI.UpdateTreasury(context.Background(), treasuryId).UpdateTreasuryRequest(updateTreasuryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TreasuryAPI.UpdateTreasury``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTreasury`: TreasuryResponse
	fmt.Fprintf(os.Stdout, "Response from `TreasuryAPI.UpdateTreasury`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**treasuryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTreasuryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTreasuryRequest** | [**UpdateTreasuryRequest**](UpdateTreasuryRequest.md) |  | 

### Return type

[**TreasuryResponse**](TreasuryResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

