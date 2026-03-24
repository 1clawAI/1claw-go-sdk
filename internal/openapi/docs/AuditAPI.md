# \AuditAPI

All URIs are relative to *http://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryAuditEvents**](AuditAPI.md#QueryAuditEvents) | **Get** /v1/audit/events | Query audit events



## QueryAuditEvents

> AuditEventsResponse QueryAuditEvents(ctx).ResourceId(resourceId).ActorId(actorId).Action(action).From(from).To(to).Limit(limit).Offset(offset).Execute()

Query audit events

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
	resourceId := "resourceId_example" // string |  (optional)
	actorId := "actorId_example" // string |  (optional)
	action := "action_example" // string |  (optional)
	from := time.Now() // time.Time |  (optional)
	to := time.Now() // time.Time |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditAPI.QueryAuditEvents(context.Background()).ResourceId(resourceId).ActorId(actorId).Action(action).From(from).To(to).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditAPI.QueryAuditEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryAuditEvents`: AuditEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `AuditAPI.QueryAuditEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryAuditEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceId** | **string** |  | 
 **actorId** | **string** |  | 
 **action** | **string** |  | 
 **from** | **time.Time** |  | 
 **to** | **time.Time** |  | 
 **limit** | **int32** |  | [default to 100]
 **offset** | **int32** |  | [default to 0]

### Return type

[**AuditEventsResponse**](AuditEventsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

