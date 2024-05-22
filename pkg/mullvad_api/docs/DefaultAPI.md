# \DefaultAPI

All URIs are relative to */app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountsPost**](DefaultAPI.md#V1AccountsPost) | **Post** /v1/accounts | Create account
[**V1ApiAddrsGet**](DefaultAPI.md#V1ApiAddrsGet) | **Get** /v1/api-addrs | List IP addresses for reaching the API
[**V1CreateApplePaymentPost**](DefaultAPI.md#V1CreateApplePaymentPost) | **Post** /v1/create-apple-payment | Create an Apple In-App payment
[**V1MeGet**](DefaultAPI.md#V1MeGet) | **Get** /v1/me | Get info about account
[**V1ProblemReportPost**](DefaultAPI.md#V1ProblemReportPost) | **Post** /v1/problem-report | Submit a problem report
[**V1RelaysGet**](DefaultAPI.md#V1RelaysGet) | **Get** /v1/relays | Relay list
[**V1ReleasesPlatformVersionGet**](DefaultAPI.md#V1ReleasesPlatformVersionGet) | **Get** /v1/releases/{platform}/{version} | Information about app release
[**V1ReplaceWireguardKeyPost**](DefaultAPI.md#V1ReplaceWireguardKeyPost) | **Post** /v1/replace-wireguard-key | Replace a WireGuard pubkey
[**V1SubmitVoucherPost**](DefaultAPI.md#V1SubmitVoucherPost) | **Post** /v1/submit-voucher | Submit a voucher
[**V1WireguardKeysPost**](DefaultAPI.md#V1WireguardKeysPost) | **Post** /v1/wireguard-keys | Add WireGuard public key
[**V1WireguardKeysPubkeyDelete**](DefaultAPI.md#V1WireguardKeysPubkeyDelete) | **Delete** /v1/wireguard-keys/{pubkey} | Delete WireGuard public key
[**V1WireguardKeysPubkeyGet**](DefaultAPI.md#V1WireguardKeysPubkeyGet) | **Get** /v1/wireguard-keys/{pubkey} | Get WireGuard public key
[**V1WwwAuthTokenPost**](DefaultAPI.md#V1WwwAuthTokenPost) | **Post** /v1/www-auth-token | Request a website auth token (valid for 1 hour)



## V1AccountsPost

> V1MeGet200Response V1AccountsPost(ctx).Execute()

Create account

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/mullvad_api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V1AccountsPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V1AccountsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountsPost`: V1MeGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V1AccountsPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountsPostRequest struct via the builder pattern


### Return type

[**V1MeGet200Response**](V1MeGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ApiAddrsGet

> []string V1ApiAddrsGet(ctx).Execute()

List IP addresses for reaching the API

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/mullvad_api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V1ApiAddrsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V1ApiAddrsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ApiAddrsGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V1ApiAddrsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1ApiAddrsGetRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CreateApplePaymentPost

> V1SubmitVoucherPost200Response V1CreateApplePaymentPost(ctx).V1CreateApplePaymentPostRequest(v1CreateApplePaymentPostRequest).Execute()

Create an Apple In-App payment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/mullvad_api"
)

func main() {
	v1CreateApplePaymentPostRequest := *openapiclient.NewV1CreateApplePaymentPostRequest("ReceiptString_example") // V1CreateApplePaymentPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V1CreateApplePaymentPost(context.Background()).V1CreateApplePaymentPostRequest(v1CreateApplePaymentPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V1CreateApplePaymentPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CreateApplePaymentPost`: V1SubmitVoucherPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V1CreateApplePaymentPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CreateApplePaymentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1CreateApplePaymentPostRequest** | [**V1CreateApplePaymentPostRequest**](V1CreateApplePaymentPostRequest.md) |  | 

### Return type

[**V1SubmitVoucherPost200Response**](V1SubmitVoucherPost200Response.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MeGet

> V1MeGet200Response V1MeGet(ctx).Execute()

Get info about account

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/mullvad_api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V1MeGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V1MeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MeGet`: V1MeGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V1MeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1MeGetRequest struct via the builder pattern


### Return type

[**V1MeGet200Response**](V1MeGet200Response.md)

### Authorization

[AccessToken](../README.md#AccessToken), [TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProblemReportPost

> V1ProblemReportPost(ctx).V1ProblemReportPostRequest(v1ProblemReportPostRequest).Execute()

Submit a problem report

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/mullvad_api"
)

func main() {
	v1ProblemReportPostRequest := *openapiclient.NewV1ProblemReportPostRequest("Log_example", map[string]interface{}(123)) // V1ProblemReportPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V1ProblemReportPost(context.Background()).V1ProblemReportPostRequest(v1ProblemReportPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V1ProblemReportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ProblemReportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1ProblemReportPostRequest** | [**V1ProblemReportPostRequest**](V1ProblemReportPostRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1RelaysGet

> V1RelaysGet200Response V1RelaysGet(ctx).Execute()

Relay list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/mullvad_api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V1RelaysGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V1RelaysGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1RelaysGet`: V1RelaysGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V1RelaysGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1RelaysGetRequest struct via the builder pattern


### Return type

[**V1RelaysGet200Response**](V1RelaysGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ReleasesPlatformVersionGet

> V1ReleasesPlatformVersionGet200Response V1ReleasesPlatformVersionGet(ctx, platform, version).Execute()

Information about app release

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/mullvad_api"
)

func main() {
	platform := "platform_example" // string | 
	version := "version_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V1ReleasesPlatformVersionGet(context.Background(), platform, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V1ReleasesPlatformVersionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ReleasesPlatformVersionGet`: V1ReleasesPlatformVersionGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V1ReleasesPlatformVersionGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ReleasesPlatformVersionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V1ReleasesPlatformVersionGet200Response**](V1ReleasesPlatformVersionGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ReplaceWireguardKeyPost

> WireGuardPubkey V1ReplaceWireguardKeyPost(ctx).V1ReplaceWireguardKeyPostRequest(v1ReplaceWireguardKeyPostRequest).Execute()

Replace a WireGuard pubkey

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/mullvad_api"
)

func main() {
	v1ReplaceWireguardKeyPostRequest := *openapiclient.NewV1ReplaceWireguardKeyPostRequest("Old_example", "New_example") // V1ReplaceWireguardKeyPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V1ReplaceWireguardKeyPost(context.Background()).V1ReplaceWireguardKeyPostRequest(v1ReplaceWireguardKeyPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V1ReplaceWireguardKeyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ReplaceWireguardKeyPost`: WireGuardPubkey
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V1ReplaceWireguardKeyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ReplaceWireguardKeyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1ReplaceWireguardKeyPostRequest** | [**V1ReplaceWireguardKeyPostRequest**](V1ReplaceWireguardKeyPostRequest.md) |  | 

### Return type

[**WireGuardPubkey**](WireGuardPubkey.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SubmitVoucherPost

> V1SubmitVoucherPost200Response V1SubmitVoucherPost(ctx).V1SubmitVoucherPostRequest(v1SubmitVoucherPostRequest).Execute()

Submit a voucher

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/mullvad_api"
)

func main() {
	v1SubmitVoucherPostRequest := *openapiclient.NewV1SubmitVoucherPostRequest("VoucherCode_example") // V1SubmitVoucherPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V1SubmitVoucherPost(context.Background()).V1SubmitVoucherPostRequest(v1SubmitVoucherPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V1SubmitVoucherPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SubmitVoucherPost`: V1SubmitVoucherPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V1SubmitVoucherPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SubmitVoucherPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1SubmitVoucherPostRequest** | [**V1SubmitVoucherPostRequest**](V1SubmitVoucherPostRequest.md) |  | 

### Return type

[**V1SubmitVoucherPost200Response**](V1SubmitVoucherPost200Response.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1WireguardKeysPost

> WireGuardPubkey V1WireguardKeysPost(ctx).V1WireguardKeysPostRequest(v1WireguardKeysPostRequest).Execute()

Add WireGuard public key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/mullvad_api"
)

func main() {
	v1WireguardKeysPostRequest := *openapiclient.NewV1WireguardKeysPostRequest("Pubkey_example") // V1WireguardKeysPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V1WireguardKeysPost(context.Background()).V1WireguardKeysPostRequest(v1WireguardKeysPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V1WireguardKeysPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1WireguardKeysPost`: WireGuardPubkey
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V1WireguardKeysPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1WireguardKeysPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1WireguardKeysPostRequest** | [**V1WireguardKeysPostRequest**](V1WireguardKeysPostRequest.md) |  | 

### Return type

[**WireGuardPubkey**](WireGuardPubkey.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1WireguardKeysPubkeyDelete

> V1WireguardKeysPubkeyDelete(ctx, pubkey).Execute()

Delete WireGuard public key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/mullvad_api"
)

func main() {
	pubkey := "pubkey_example" // string | The WireGuard public key, urlencoded.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V1WireguardKeysPubkeyDelete(context.Background(), pubkey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V1WireguardKeysPubkeyDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pubkey** | **string** | The WireGuard public key, urlencoded. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1WireguardKeysPubkeyDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1WireguardKeysPubkeyGet

> WireGuardPubkey V1WireguardKeysPubkeyGet(ctx, pubkey).Execute()

Get WireGuard public key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/mullvad_api"
)

func main() {
	pubkey := "pubkey_example" // string | The WireGuard public key, urlencoded.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V1WireguardKeysPubkeyGet(context.Background(), pubkey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V1WireguardKeysPubkeyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1WireguardKeysPubkeyGet`: WireGuardPubkey
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V1WireguardKeysPubkeyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pubkey** | **string** | The WireGuard public key, urlencoded. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1WireguardKeysPubkeyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WireGuardPubkey**](WireGuardPubkey.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1WwwAuthTokenPost

> V1WwwAuthTokenPost200Response V1WwwAuthTokenPost(ctx).Execute()

Request a website auth token (valid for 1 hour)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/mullvad_api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V1WwwAuthTokenPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V1WwwAuthTokenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1WwwAuthTokenPost`: V1WwwAuthTokenPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V1WwwAuthTokenPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1WwwAuthTokenPostRequest struct via the builder pattern


### Return type

[**V1WwwAuthTokenPost200Response**](V1WwwAuthTokenPost200Response.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

