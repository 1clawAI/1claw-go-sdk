# \BillingAPI

All URIs are relative to *http://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BillingCreditBalance**](BillingAPI.md#BillingCreditBalance) | **Get** /v1/billing/credits/balance | Get credit balance
[**BillingCreditTopup**](BillingAPI.md#BillingCreditTopup) | **Post** /v1/billing/credits/topup | Top up prepaid credits via Stripe
[**BillingCreditTransactions**](BillingAPI.md#BillingCreditTransactions) | **Get** /v1/billing/credits/transactions | Get credit transaction ledger
[**BillingHistory**](BillingAPI.md#BillingHistory) | **Get** /v1/billing/history | Get usage event history (legacy)
[**BillingOverageMethod**](BillingAPI.md#BillingOverageMethod) | **Patch** /v1/billing/overage-method | Set overage payment method
[**BillingPortal**](BillingAPI.md#BillingPortal) | **Post** /v1/billing/portal | Open Stripe Customer Portal
[**BillingSubscribe**](BillingAPI.md#BillingSubscribe) | **Post** /v1/billing/subscribe | Start a subscription via Stripe Checkout
[**BillingSubscription**](BillingAPI.md#BillingSubscription) | **Get** /v1/billing/subscription | Get subscription, usage, and credit summary
[**BillingUsage**](BillingAPI.md#BillingUsage) | **Get** /v1/billing/usage | Get usage summary (legacy)
[**BillingWebhook**](BillingAPI.md#BillingWebhook) | **Post** /v1/billing/webhooks | Stripe webhook receiver
[**DisableLlmTokenBilling**](BillingAPI.md#DisableLlmTokenBilling) | **Post** /v1/billing/llm-token-billing/disable | Disable LLM token billing
[**GetLlmTokenBilling**](BillingAPI.md#GetLlmTokenBilling) | **Get** /v1/billing/llm-token-billing | Get LLM token billing status
[**SubscribeLlmTokenBilling**](BillingAPI.md#SubscribeLlmTokenBilling) | **Post** /v1/billing/llm-token-billing/subscribe | Subscribe to LLM token billing



## BillingCreditBalance

> CreditBalanceResponse BillingCreditBalance(ctx).Execute()

Get credit balance

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
	resp, r, err := apiClient.BillingAPI.BillingCreditBalance(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingCreditBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingCreditBalance`: CreditBalanceResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingCreditBalance`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBillingCreditBalanceRequest struct via the builder pattern


### Return type

[**CreditBalanceResponse**](CreditBalanceResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingCreditTopup

> CheckoutUrlResponse BillingCreditTopup(ctx).TopupRequest(topupRequest).Execute()

Top up prepaid credits via Stripe

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
	topupRequest := *openapiclient.NewTopupRequest(int32(123)) // TopupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingCreditTopup(context.Background()).TopupRequest(topupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingCreditTopup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingCreditTopup`: CheckoutUrlResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingCreditTopup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingCreditTopupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **topupRequest** | [**TopupRequest**](TopupRequest.md) |  | 

### Return type

[**CheckoutUrlResponse**](CheckoutUrlResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingCreditTransactions

> CreditTransactionsListResponse BillingCreditTransactions(ctx).Page(page).Limit(limit).Execute()

Get credit transaction ledger

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
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingCreditTransactions(context.Background()).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingCreditTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingCreditTransactions`: CreditTransactionsListResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingCreditTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingCreditTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 50]

### Return type

[**CreditTransactionsListResponse**](CreditTransactionsListResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingHistory

> UsageHistoryResponse BillingHistory(ctx).Limit(limit).Execute()

Get usage event history (legacy)

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
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingHistory(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingHistory`: UsageHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 50]

### Return type

[**UsageHistoryResponse**](UsageHistoryResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingOverageMethod

> OverageMethodResponse BillingOverageMethod(ctx).OverageMethodRequest(overageMethodRequest).Execute()

Set overage payment method

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
	overageMethodRequest := *openapiclient.NewOverageMethodRequest("Method_example") // OverageMethodRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingOverageMethod(context.Background()).OverageMethodRequest(overageMethodRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingOverageMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingOverageMethod`: OverageMethodResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingOverageMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingOverageMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **overageMethodRequest** | [**OverageMethodRequest**](OverageMethodRequest.md) |  | 

### Return type

[**OverageMethodResponse**](OverageMethodResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingPortal

> PortalUrlResponse BillingPortal(ctx).Execute()

Open Stripe Customer Portal

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
	resp, r, err := apiClient.BillingAPI.BillingPortal(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingPortal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingPortal`: PortalUrlResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingPortal`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBillingPortalRequest struct via the builder pattern


### Return type

[**PortalUrlResponse**](PortalUrlResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingSubscribe

> CheckoutUrlResponse BillingSubscribe(ctx).SubscribeRequest(subscribeRequest).Execute()

Start a subscription via Stripe Checkout

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
	subscribeRequest := *openapiclient.NewSubscribeRequest("Tier_example", "Interval_example") // SubscribeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingSubscribe(context.Background()).SubscribeRequest(subscribeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingSubscribe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingSubscribe`: CheckoutUrlResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingSubscribe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscribeRequest** | [**SubscribeRequest**](SubscribeRequest.md) |  | 

### Return type

[**CheckoutUrlResponse**](CheckoutUrlResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingSubscription

> SubscriptionResponse BillingSubscription(ctx).Execute()

Get subscription, usage, and credit summary

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
	resp, r, err := apiClient.BillingAPI.BillingSubscription(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingSubscription`: SubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingSubscription`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBillingSubscriptionRequest struct via the builder pattern


### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingUsage

> UsageSummaryResponse BillingUsage(ctx).Execute()

Get usage summary (legacy)

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
	resp, r, err := apiClient.BillingAPI.BillingUsage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingUsage`: UsageSummaryResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingUsage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBillingUsageRequest struct via the builder pattern


### Return type

[**UsageSummaryResponse**](UsageSummaryResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingWebhook

> BillingWebhook(ctx).Execute()

Stripe webhook receiver

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
	r, err := apiClient.BillingAPI.BillingWebhook(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBillingWebhookRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableLlmTokenBilling

> LlmDisableResponse DisableLlmTokenBilling(ctx).Execute()

Disable LLM token billing



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
	resp, r, err := apiClient.BillingAPI.DisableLlmTokenBilling(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.DisableLlmTokenBilling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableLlmTokenBilling`: LlmDisableResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.DisableLlmTokenBilling`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDisableLlmTokenBillingRequest struct via the builder pattern


### Return type

[**LlmDisableResponse**](LlmDisableResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLlmTokenBilling

> LlmTokenBillingStatus GetLlmTokenBilling(ctx).Execute()

Get LLM token billing status



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
	resp, r, err := apiClient.BillingAPI.GetLlmTokenBilling(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetLlmTokenBilling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLlmTokenBilling`: LlmTokenBillingStatus
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetLlmTokenBilling`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLlmTokenBillingRequest struct via the builder pattern


### Return type

[**LlmTokenBillingStatus**](LlmTokenBillingStatus.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscribeLlmTokenBilling

> LlmCheckoutResponse SubscribeLlmTokenBilling(ctx).Execute()

Subscribe to LLM token billing



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
	resp, r, err := apiClient.BillingAPI.SubscribeLlmTokenBilling(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.SubscribeLlmTokenBilling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscribeLlmTokenBilling`: LlmCheckoutResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.SubscribeLlmTokenBilling`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeLlmTokenBillingRequest struct via the builder pattern


### Return type

[**LlmCheckoutResponse**](LlmCheckoutResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

