# \AdminAPI

All URIs are relative to *http://localhost:8443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminDeleteSetting**](AdminAPI.md#AdminDeleteSetting) | **Delete** /v1/admin/settings/{key} | Delete a platform setting
[**AdminDeleteUser**](AdminAPI.md#AdminDeleteUser) | **Delete** /v1/admin/users/{user_id} | Delete a user (cascade)
[**AdminGetOrgLimits**](AdminAPI.md#AdminGetOrgLimits) | **Get** /v1/admin/orgs/{org_id}/limits | Get org limits
[**AdminGetX402Config**](AdminAPI.md#AdminGetX402Config) | **Get** /v1/admin/x402 | Get x402 payment config
[**AdminListSettings**](AdminAPI.md#AdminListSettings) | **Get** /v1/admin/settings | List platform settings
[**AdminListUsers**](AdminAPI.md#AdminListUsers) | **Get** /v1/admin/users | List all platform users
[**AdminResetUsageEvents**](AdminAPI.md#AdminResetUsageEvents) | **Post** /v1/admin/usage/reset | Reset all API usage events (testing)
[**AdminResetUsageForUserByEmail**](AdminAPI.md#AdminResetUsageForUserByEmail) | **Post** /v1/admin/usage/reset-for-user | Reset API usage for a user&#39;s organization
[**AdminSetBillingTier**](AdminAPI.md#AdminSetBillingTier) | **Put** /v1/admin/orgs/{org_id}/billing-tier | Set org billing tier (without Stripe)
[**AdminUpdateOrgLimits**](AdminAPI.md#AdminUpdateOrgLimits) | **Put** /v1/admin/orgs/{org_id}/limits | Update org limits
[**AdminUpdateSetting**](AdminAPI.md#AdminUpdateSetting) | **Put** /v1/admin/settings/{key} | Update a platform setting
[**AdminUpdateX402Config**](AdminAPI.md#AdminUpdateX402Config) | **Put** /v1/admin/x402 | Update x402 payment config



## AdminDeleteSetting

> AdminDeleteSetting(ctx, key).Execute()

Delete a platform setting

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
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminAPI.AdminDeleteSetting(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminDeleteSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminDeleteSettingRequest struct via the builder pattern


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


## AdminDeleteUser

> AdminDeleteUser(ctx, userId).Execute()

Delete a user (cascade)

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminAPI.AdminDeleteUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminDeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminDeleteUserRequest struct via the builder pattern


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


## AdminGetOrgLimits

> OrgLimitsResponse AdminGetOrgLimits(ctx, orgId).Execute()

Get org limits

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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminGetOrgLimits(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminGetOrgLimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminGetOrgLimits`: OrgLimitsResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminGetOrgLimits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminGetOrgLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrgLimitsResponse**](OrgLimitsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminGetX402Config

> X402ConfigResponse AdminGetX402Config(ctx).Execute()

Get x402 payment config

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
	resp, r, err := apiClient.AdminAPI.AdminGetX402Config(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminGetX402Config``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminGetX402Config`: X402ConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminGetX402Config`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminGetX402ConfigRequest struct via the builder pattern


### Return type

[**X402ConfigResponse**](X402ConfigResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminListSettings

> SettingsListResponse AdminListSettings(ctx).Execute()

List platform settings

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
	resp, r, err := apiClient.AdminAPI.AdminListSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminListSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminListSettings`: SettingsListResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminListSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminListSettingsRequest struct via the builder pattern


### Return type

[**SettingsListResponse**](SettingsListResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminListUsers

> AdminUsersListResponse AdminListUsers(ctx).Execute()

List all platform users

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
	resp, r, err := apiClient.AdminAPI.AdminListUsers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminListUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminListUsers`: AdminUsersListResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminListUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminListUsersRequest struct via the builder pattern


### Return type

[**AdminUsersListResponse**](AdminUsersListResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminResetUsageEvents

> ResetUsageEventsResponse AdminResetUsageEvents(ctx).Execute()

Reset all API usage events (testing)



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
	resp, r, err := apiClient.AdminAPI.AdminResetUsageEvents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminResetUsageEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminResetUsageEvents`: ResetUsageEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminResetUsageEvents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminResetUsageEventsRequest struct via the builder pattern


### Return type

[**ResetUsageEventsResponse**](ResetUsageEventsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminResetUsageForUserByEmail

> ResetUsageForUserEmailResponse AdminResetUsageForUserByEmail(ctx).ResetUsageForUserEmailRequest(resetUsageForUserEmailRequest).Execute()

Reset API usage for a user's organization



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
	resetUsageForUserEmailRequest := *openapiclient.NewResetUsageForUserEmailRequest("Email_example") // ResetUsageForUserEmailRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminResetUsageForUserByEmail(context.Background()).ResetUsageForUserEmailRequest(resetUsageForUserEmailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminResetUsageForUserByEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminResetUsageForUserByEmail`: ResetUsageForUserEmailResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminResetUsageForUserByEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminResetUsageForUserByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetUsageForUserEmailRequest** | [**ResetUsageForUserEmailRequest**](ResetUsageForUserEmailRequest.md) |  | 

### Return type

[**ResetUsageForUserEmailResponse**](ResetUsageForUserEmailResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminSetBillingTier

> AdminSetBillingTier(ctx, orgId).SetBillingTierRequest(setBillingTierRequest).Execute()

Set org billing tier (without Stripe)



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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	setBillingTierRequest := *openapiclient.NewSetBillingTierRequest("Tier_example") // SetBillingTierRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminAPI.AdminSetBillingTier(context.Background(), orgId).SetBillingTierRequest(setBillingTierRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminSetBillingTier``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminSetBillingTierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setBillingTierRequest** | [**SetBillingTierRequest**](SetBillingTierRequest.md) |  | 

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


## AdminUpdateOrgLimits

> OrgLimitsResponse AdminUpdateOrgLimits(ctx, orgId).UpdateOrgLimitsRequest(updateOrgLimitsRequest).Execute()

Update org limits

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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateOrgLimitsRequest := *openapiclient.NewUpdateOrgLimitsRequest() // UpdateOrgLimitsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminUpdateOrgLimits(context.Background(), orgId).UpdateOrgLimitsRequest(updateOrgLimitsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminUpdateOrgLimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminUpdateOrgLimits`: OrgLimitsResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminUpdateOrgLimits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminUpdateOrgLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrgLimitsRequest** | [**UpdateOrgLimitsRequest**](UpdateOrgLimitsRequest.md) |  | 

### Return type

[**OrgLimitsResponse**](OrgLimitsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminUpdateSetting

> SettingResponse AdminUpdateSetting(ctx, key).UpdateSettingRequest(updateSettingRequest).Execute()

Update a platform setting

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
	key := "key_example" // string | 
	updateSettingRequest := *openapiclient.NewUpdateSettingRequest("Value_example") // UpdateSettingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminUpdateSetting(context.Background(), key).UpdateSettingRequest(updateSettingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminUpdateSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminUpdateSetting`: SettingResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminUpdateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminUpdateSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSettingRequest** | [**UpdateSettingRequest**](UpdateSettingRequest.md) |  | 

### Return type

[**SettingResponse**](SettingResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminUpdateX402Config

> X402ConfigResponse AdminUpdateX402Config(ctx).X402ConfigResponse(x402ConfigResponse).Execute()

Update x402 payment config

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
	x402ConfigResponse := *openapiclient.NewX402ConfigResponse() // X402ConfigResponse | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminUpdateX402Config(context.Background()).X402ConfigResponse(x402ConfigResponse).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminUpdateX402Config``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminUpdateX402Config`: X402ConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminUpdateX402Config`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminUpdateX402ConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **x402ConfigResponse** | [**X402ConfigResponse**](X402ConfigResponse.md) |  | 

### Return type

[**X402ConfigResponse**](X402ConfigResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

