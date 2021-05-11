---
author: xxchan <xxchan22f@gmail.com>
status: draft
updated_at: 2021-05-11
---

# AOS-60: Network Error and Retryable Server-side Errors

## Background

We can classify all errors by the root cause of error:

- User-side: incapability, wrong usage, pass an append object into multipart operations (like `failed to send a request to the storage system`)
- Network-side: Network related error (like `failed to receive a response from the storage system`)
  - explicit error: connection refuse, connection reset
  - implicit error: timeout
- Service-side: not exist, permission denied, already exists... (like `received an error response from the storage system`)

All the errors will go into `formatError`, and be converted into either a defined error or an unexpected error. Currently, we have only defined some user-side errors and server-side errors. 

For errors which are returned by SDK, but not belong to the SDK's error type, they are either belong to user-side or network-side. 

```go
func formatError(err error) error {
	if _, ok := err.(services.AosError); ok {
		// client-side
		return err 
	}

	var e *qserror.QingStorError
	if !errors.As(err, &e) {
		// client-side or network-side
		return fmt.Errorf("%w: %v", services.ErrUnexpected, err)
	}

	// server-side
	switch e.Code {
	case "permission_denied":
		return fmt.Errorf("%w: %v", services.ErrPermissionDenied, e)
	case "object_not_exists":
		return fmt.Errorf("%w: %v", services.ErrObjectNotExist, e)
	default:
		return fmt.Errorf("%w: %v", services.ErrUnexpected, err)
	}
}
```

Network errors are common and should be expected. Besides, there's a common need to retry a network error. So we should provide network error codes. 

It will be better if a good retry mechanism can be built on top of network errors.

## Proposal

So I propose to add the following global error codes:

Network-side: 

- `ErrNetworkTemporary`: context.canceled, connection timeout, no response for some time 

And use the following function to judge and convert.

```go 
func IsErrNetworkTemporary(err error) bool {
	if errors.Is(err, io.ErrUnexpectedEOF) {
		return true
	}

	if _, ok := err.(interface{ Temporary() bool }); ok && e.Temporary() {
		return true
	}

	var urlErr *url.Error
	if errors.As(err, &urlErr) {
		retriable := []string{"connection refused", "connection reset"}
		for _, s := range retriable {
			if strings.Contains(e.Error(), s) {
				return true
			}
		}
	}

	var netErr *net.OpError
	if errors.As(err, &netErr) && netErr.Op == "dial" {
		return true
	}
	
	return false
}
```

Retryable server-side errors:

- `ErrServerInternal`: HTTP 5xx
- `ErrRequestThrottled`: HTTP 429 Too Many Requests/Limit Exceeded, 503 SlowDown, ...


## Rationale

1. The errors should apply to different type of services:

   - local file system (but may have network operation. e.g., fetch)
   - RPC-based 
   - HTTP-based 
   - other network protocol

2. The possible error types are of many kinds: `*net.OpError`, `*url.Error`, `*os.SyscallError`, ... (non-exhaustive) And they are probably not orthogonal, so `IsErrNetworkTemporary` uses `errors.As` and `errors.Is` to check, instead of type assertion. And there's a single `return false` in the end. In this way, the check may be more complete, but brings more overheads.

3. The granularity of network errors is very coarse in this proposal. 

	- For non-temporary network, there's little need to handle them specially in code, but it may be good if they are labelled. However, it may be hard to distinguish between non-temporary network errors and user-side errors.
	- An internal classification in temporary network errors is acceptable if there's a good idea.

4. Usually there's only bipartite communication between the client and the storage service. However, there's a `fetch` operation. Usually the server fetches the resource, and if there's an error, it's a server-side error. But if the service doesn't support `fetch`, and we implement `fetch` in client code, e.g., `go-service-fs`, what's the error's type? In our abstraction, we treat it as server-side. 

### Examples of judging retryability

Google SDK 

```go 
func shouldRetry(err error) bool {
	if err == io.ErrUnexpectedEOF {
		return true
	}
	switch e := err.(type) {
	case *googleapi.Error:
		// Retry on 429 and 5xx, according to
		// https://cloud.google.com/storage/docs/exponential-backoff.
		return e.Code == 429 || (e.Code >= 500 && e.Code < 600)
	case *url.Error:
		// Retry socket-level errors ECONNREFUSED and ENETUNREACH (from syscall).
		// Unfortunately the error type is unexported, so we resort to string
		// matching.
		retriable := []string{"connection refused", "connection reset"}
		for _, s := range retriable {
			if strings.Contains(e.Error(), s) {
				return true
			}
		}
		return false
	case interface{ Temporary() bool }:
		return e.Temporary()
	default:
		return false
	}
}
```

AWS SDK 

```go 
func shouldRetryError(origErr error) bool {
	switch err := origErr.(type) {
	case awserr.Error:
		// ...

	case *url.Error:
		if strings.Contains(err.Error(), "connection refused") {
			// Refused connections should be retried as the service may not yet
			// be running on the port. Go TCP dial considers refused
			// connections as not temporary.
			return true
		}
		// *url.Error only implements Temporary after golang 1.6 but since
		// url.Error only wraps the error:
		return shouldRetryError(err.Err)

	case temporary:
		if netErr, ok := err.(*net.OpError); ok && netErr.Op == "dial" {
			return true
		}
		// If the error is temporary, we want to allow continuation of the
		// retry process
		return err.Temporary() || isErrConnectionReset(origErr)

	case nil:
		// `awserr.Error.OrigErr()` can be nil, meaning there was an error but
		// because we don't know the cause, it is marked as retryable. See
		// TestRequest4xxUnretryable for an example.
		return true

	default:
		switch err.Error() {
		case "net/http: request canceled",
			"net/http: request canceled while waiting for connection":
			// known 1.5 error case when an http request is cancelled
			return false
		}
		// here we don't know the error; so we allow a retry.
		return true
	}
}
```

## Compatibility

No breaking changes.

## Implementation

Most of the work would be done by the author of this proposal.

Add the listed errors to [go-storage], and use them in `go-service-*`.

[go-storage]: https://github.com/aos-dev/go-storage
