---
author: xxchan <xxchan22f@gmail.com>
status: draft
updated_at: 2021-05-08
---

# AOS-51: Identify go-storage Error

## Background

Currently we use a function named `formatError` in generated code to turn SDK errors into our errors, as defined in [AOS-47]. However, sometimes we have used our errors before `formatError`, so our errors will also be formatted. Then it's hard for us to distinguish between non-SDK errors and our errors.

Take `go-storage-s3` as an example:

```go
func formatError(err error) error {
	e, ok := err.(awserr.RequestFailure)
	if !ok {
		return fmt.Errorf("%w: %v", services.ErrUnexpected, err)
	}
	switch e.Code() {
	// AWS SDK will use status code to generate awserr.Error, so "NotFound" should also be supported.
	case "NoSuchKey", "NotFound":
		return fmt.Errorf("%w: %v", services.ErrObjectNotExist, err)
	case "AccessDenied":
		return fmt.Errorf("%w: %v", services.ErrPermissionDenied, err)
	default:
		return fmt.Errorf("%w: %v", services.ErrUnexpected, err)
	}
}
```

This is problematic since our `PairRequiredError` will also be turned into `ErrUnexpected`.

## Proposal

So I propose to add support for recognizing errors defined in `go-storage`, including:

- Add a new type `ErrorCode` to replace sentinel errors.
- Add a special `ErrorCode` `ErrGoStorage`
- `ErrorCode` and every error `struct` should implement `Is(target error) bool { return target == ErrGoStorage}`

## Rationale

We cannot add methods to sentinel errors, so we need to replace them.

### Alternative

Add an interface and use type cast to judge.

```go
type ErrGoStorage interface {
    IsErrGoStorage()
}
```

## Compatibility

This proposal will not break users, since `errors.New()` create a private error type which cannot be used by users.

## Implementation

- Add `ErrorCode` `ErrGoStorage`
- Replace `errors.New` with `NewErrorCode`
- Implement `Is` for error `struct`s

```go 
// Create a new error code.
//
// Users should return/compare created ErrorCode instances, instead of calling this function.
func NewErrorCode(text string) ErrorCode {
	return ErrorCode{text}
}

type ErrorCode struct {
	s string
}

func (e ErrorCode) Error() string {
	return e.s
}

var (
	ErrGoStorage = NewErrorCode("go-storage error")

	ErrObjectNotExist = NewErrorCode("object not exist")
	//...
)

func (e ErrorCode) Is(target error) bool {
	return target == ErrGoStorage
}

func (e PairRequiredError) Is(target error) bool {
	return target == ErrGoStorage
}
```

[AOS-47]: ./47-additional-error-specification.md
