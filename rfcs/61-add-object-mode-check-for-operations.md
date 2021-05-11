---
author: Prnyself <lanceren@yunify.com>
status: draft
updated_at: 2021-05-11
---

# AOS-61: Add object mode check for operations

## Background

In [AOS-25], we added support for object modes by bit map. All available object modes are listed below:

```go
type ObjectMode uint32

// All available object mode
const (
	// ModeDir means this Object represents a dir which can be used to list with dir mode.
	ModeDir ObjectMode = 1 << iota
	// ModeRead means this Object can be used to read content.
	ModeRead
	// ModeLink means this Object is a link which targets to another Object.
	ModeLink
	// ModePart means this Object is a Multipart Object which can be used for multipart operations.
	ModePart
	// ModeBlock means this Object is a Block Object which can be used for block operations.
	ModeBlock
	// ModePage means this Object is a Page Object which can be used for random write with offset.
	ModePage
	// ModeAppend means this Object is a Append Object which can be used for append.
	ModeAppend
)
```

It is intended to check object mode at the start of specific operation. For instance, both `WritePart`
and `WriteAppend` got a pointer to `Object` as an input, we need to ensure this `Object` is available 
for certain operation, so we should add object mode check and return `ObjectModeInvalidError`(introduced in [AOS-47]) asap if `Object` not fit.

### Current Practice

The check is implemented by service in the actual method call. 
For example, in [go-service-qingstor](https://github.com/aos-dev/go-service-qingstor/blob/master/storage.go#L534):
```go
func (s *Storage) writeAppend(ctx context.Context, o *Object, r io.Reader, size int64, opt pairStorageWriteAppend) (n int64, err error) {
	if !o.Mode.IsAppend() {
		err = services.ObjectModeInvalidError{Expected: ModeAppend, Actual: o.Mode}
		return
	}
	...
}
```

We check `*Object` mode `IsAppend` at the start of `writeAppend` which is called by `WriteAppendWithContext` which is generated.

```go
func (s *Storage) writeMultipart(ctx context.Context, o *Object, r io.Reader, size int64, index int, opt pairStorageWriteMultipart) (n int64, err error) {
	if o.Mode&ModePart == 0 {
		return 0, services.ObjectModeInvalidError{Expected: ModePart, Actual: o.Mode}
	}
	...
}	
```

We check `*Object` mode `ModePart` at the start of `writeMultipart` which is called by `WriteMultipartWithContext` which is generated, too.

### Operation with Object Input

All function listed below contains `XXXWithContext` method 

- Appender
  - CommitAppend
  - WriteAppend
- Blocker
  - CombineBlock
  - ListBlock
  - WriteBlock
- Multiparter
  - CompleteMultipart
  - ListMultipart
  - WriteMultipart
- Pager
  - WritePage

## Proposal

So I propose that we should add mode check in specific operation, return `ObjectModeInvalidError` if mode not meet, and the check should be generated.

In this way, we can ensure the `*Object` input must meet the operation, and service do not need to care about the check in specific operation.

## Rationale

None.

## Compatibility

No break changes

## Implementation

Most of the work would be done by the author of this proposal, including:
- Implement mode check
- Upgrade go-service-*
- Remove all mode check in go-sevice-*

[AOS-25]: ./25-object-mode.md
[AOS-47]: ./47-additional-error-specification.md