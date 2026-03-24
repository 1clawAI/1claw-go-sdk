# \TransactionsAPI

All URIs are relative to *http://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTransaction**](TransactionsAPI.md#GetTransaction) | **Get** /v1/agents/{agent_id}/transactions/{tx_id} | Get a transaction by ID
[**ListTransactions**](TransactionsAPI.md#ListTransactions) | **Get** /v1/agents/{agent_id}/transactions | List agent transactions
[**SignTransaction**](TransactionsAPI.md#SignTransaction) | **Post** /v1/agents/{agent_id}/transactions/sign | Sign a transaction without broadcasting
[**SimulateBundle**](TransactionsAPI.md#SimulateBundle) | **Post** /v1/agents/{agent_id}/transactions/simulate-bundle | Simulate a bundle of transactions
[**SimulateTransaction**](TransactionsAPI.md#SimulateTransaction) | **Post** /v1/agents/{agent_id}/transactions/simulate | Simulate a transaction via Tenderly
[**SubmitTransaction**](TransactionsAPI.md#SubmitTransaction) | **Post** /v1/agents/{agent_id}/transactions | Submit a transaction for signing



## GetTransaction

> TransactionResponse GetTransaction(ctx, agentId, txId).IncludeSignedTx(includeSignedTx).Execute()

Get a transaction by ID

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
	agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	txId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	includeSignedTx := false // bool | Set to `true` or `1` to include the raw signed transaction hex in the response. Omitted by default to reduce key exfiltration risk. Only the literal values \"true\" or \"1\" enable inclusion; any other value or omission returns responses without signed_tx. Applies to GET /v1/agents/{agent_id}/transactions and GET /v1/agents/{agent_id}/transactions/{tx_id}.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTransaction(context.Background(), agentId, txId).IncludeSignedTx(includeSignedTx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransaction`: TransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 
**txId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeSignedTx** | **bool** | Set to &#x60;true&#x60; or &#x60;1&#x60; to include the raw signed transaction hex in the response. Omitted by default to reduce key exfiltration risk. Only the literal values \&quot;true\&quot; or \&quot;1\&quot; enable inclusion; any other value or omission returns responses without signed_tx. Applies to GET /v1/agents/{agent_id}/transactions and GET /v1/agents/{agent_id}/transactions/{tx_id}.  | [default to false]

### Return type

[**TransactionResponse**](TransactionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactions

> TransactionListResponse ListTransactions(ctx, agentId).IncludeSignedTx(includeSignedTx).Execute()

List agent transactions

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
	agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	includeSignedTx := false // bool | Set to `true` or `1` to include the raw signed transaction hex in the response. Omitted by default to reduce key exfiltration risk. Only the literal values \"true\" or \"1\" enable inclusion; any other value or omission returns responses without signed_tx. Applies to GET /v1/agents/{agent_id}/transactions and GET /v1/agents/{agent_id}/transactions/{tx_id}.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.ListTransactions(context.Background(), agentId).IncludeSignedTx(includeSignedTx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.ListTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTransactions`: TransactionListResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.ListTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeSignedTx** | **bool** | Set to &#x60;true&#x60; or &#x60;1&#x60; to include the raw signed transaction hex in the response. Omitted by default to reduce key exfiltration risk. Only the literal values \&quot;true\&quot; or \&quot;1\&quot; enable inclusion; any other value or omission returns responses without signed_tx. Applies to GET /v1/agents/{agent_id}/transactions and GET /v1/agents/{agent_id}/transactions/{tx_id}.  | [default to false]

### Return type

[**TransactionListResponse**](TransactionListResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignTransaction

> SignTransactionResponse SignTransaction(ctx, agentId).SignTransactionRequest(signTransactionRequest).Execute()

Sign a transaction without broadcasting



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
	agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	signTransactionRequest := *openapiclient.NewSignTransactionRequest("To_example", "Value_example", "Chain_example") // SignTransactionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.SignTransaction(context.Background(), agentId).SignTransactionRequest(signTransactionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.SignTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignTransaction`: SignTransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.SignTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSignTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signTransactionRequest** | [**SignTransactionRequest**](SignTransactionRequest.md) |  | 

### Return type

[**SignTransactionResponse**](SignTransactionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SimulateBundle

> BundleSimulationResponse SimulateBundle(ctx, agentId).SimulateBundleRequest(simulateBundleRequest).Execute()

Simulate a bundle of transactions

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
	agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	simulateBundleRequest := *openapiclient.NewSimulateBundleRequest([]openapiclient.SimulateTransactionRequest{*openapiclient.NewSimulateTransactionRequest("To_example", "Value_example", "Chain_example")}) // SimulateBundleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.SimulateBundle(context.Background(), agentId).SimulateBundleRequest(simulateBundleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.SimulateBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SimulateBundle`: BundleSimulationResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.SimulateBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSimulateBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **simulateBundleRequest** | [**SimulateBundleRequest**](SimulateBundleRequest.md) |  | 

### Return type

[**BundleSimulationResponse**](BundleSimulationResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SimulateTransaction

> SimulationResponse SimulateTransaction(ctx, agentId).SimulateTransactionRequest(simulateTransactionRequest).Execute()

Simulate a transaction via Tenderly

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
	agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	simulateTransactionRequest := *openapiclient.NewSimulateTransactionRequest("To_example", "Value_example", "Chain_example") // SimulateTransactionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.SimulateTransaction(context.Background(), agentId).SimulateTransactionRequest(simulateTransactionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.SimulateTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SimulateTransaction`: SimulationResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.SimulateTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSimulateTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **simulateTransactionRequest** | [**SimulateTransactionRequest**](SimulateTransactionRequest.md) |  | 

### Return type

[**SimulationResponse**](SimulationResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitTransaction

> TransactionResponse SubmitTransaction(ctx, agentId).SubmitTransactionRequest(submitTransactionRequest).IdempotencyKey(idempotencyKey).Execute()

Submit a transaction for signing



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
	agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	submitTransactionRequest := *openapiclient.NewSubmitTransactionRequest("To_example", "Value_example", "Chain_example") // SubmitTransactionRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Optional key for replay protection; duplicate requests return cached response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.SubmitTransaction(context.Background(), agentId).SubmitTransactionRequest(submitTransactionRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.SubmitTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitTransaction`: TransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.SubmitTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **submitTransactionRequest** | [**SubmitTransactionRequest**](SubmitTransactionRequest.md) |  | 
 **idempotencyKey** | **string** | Optional key for replay protection; duplicate requests return cached response. | 

### Return type

[**TransactionResponse**](TransactionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

