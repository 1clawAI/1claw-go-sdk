# \SharingAPI

All URIs are relative to *http://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptShare**](SharingAPI.md#AcceptShare) | **Post** /v1/shares/{share_id}/accept | Accept an inbound share
[**AccessShare**](SharingAPI.md#AccessShare) | **Get** /v1/share/{share_id} | Access a shared secret
[**CreateShare**](SharingAPI.md#CreateShare) | **Post** /v1/secrets/{secret_id}/share | Create a share link for a secret
[**DeclineShare**](SharingAPI.md#DeclineShare) | **Post** /v1/shares/{share_id}/decline | Decline an inbound share
[**ListInboundShares**](SharingAPI.md#ListInboundShares) | **Get** /v1/shares/inbound | List shares sent to you
[**ListOutboundShares**](SharingAPI.md#ListOutboundShares) | **Get** /v1/shares/outbound | List shares you have sent
[**RevokeShare**](SharingAPI.md#RevokeShare) | **Delete** /v1/share/{share_id} | Revoke a share link



## AcceptShare

> AcceptShare(ctx, shareId).Execute()

Accept an inbound share

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
	shareId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SharingAPI.AcceptShare(context.Background(), shareId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharingAPI.AcceptShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shareId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptShareRequest struct via the builder pattern


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


## AccessShare

> SharedSecretResponse AccessShare(ctx, shareId).Execute()

Access a shared secret

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
	shareId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharingAPI.AccessShare(context.Background(), shareId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharingAPI.AccessShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessShare`: SharedSecretResponse
	fmt.Fprintf(os.Stdout, "Response from `SharingAPI.AccessShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shareId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SharedSecretResponse**](SharedSecretResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateShare

> ShareResponse CreateShare(ctx, secretId).CreateShareRequest(createShareRequest).Execute()

Create a share link for a secret

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	secretId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	createShareRequest := *openapiclient.NewCreateShareRequest("RecipientType_example", time.Now()) // CreateShareRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharingAPI.CreateShare(context.Background(), secretId).CreateShareRequest(createShareRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharingAPI.CreateShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateShare`: ShareResponse
	fmt.Fprintf(os.Stdout, "Response from `SharingAPI.CreateShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createShareRequest** | [**CreateShareRequest**](CreateShareRequest.md) |  | 

### Return type

[**ShareResponse**](ShareResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeclineShare

> DeclineShare(ctx, shareId).Execute()

Decline an inbound share

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
	shareId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SharingAPI.DeclineShare(context.Background(), shareId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharingAPI.DeclineShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shareId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeclineShareRequest struct via the builder pattern


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


## ListInboundShares

> ShareListResponse ListInboundShares(ctx).Execute()

List shares sent to you

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
	resp, r, err := apiClient.SharingAPI.ListInboundShares(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharingAPI.ListInboundShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInboundShares`: ShareListResponse
	fmt.Fprintf(os.Stdout, "Response from `SharingAPI.ListInboundShares`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListInboundSharesRequest struct via the builder pattern


### Return type

[**ShareListResponse**](ShareListResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOutboundShares

> ShareListResponse ListOutboundShares(ctx).Execute()

List shares you have sent

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
	resp, r, err := apiClient.SharingAPI.ListOutboundShares(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharingAPI.ListOutboundShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOutboundShares`: ShareListResponse
	fmt.Fprintf(os.Stdout, "Response from `SharingAPI.ListOutboundShares`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOutboundSharesRequest struct via the builder pattern


### Return type

[**ShareListResponse**](ShareListResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeShare

> RevokeShare(ctx, shareId).Execute()

Revoke a share link

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
	shareId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SharingAPI.RevokeShare(context.Background(), shareId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharingAPI.RevokeShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shareId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeShareRequest struct via the builder pattern


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

