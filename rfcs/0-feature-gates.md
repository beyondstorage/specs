---
author: Xuanwo <github@xuanwo.io>
status: draft
updated_at: 2021-05-27
---

# Proposal: Feature Gates

## Background

Behavior consistence is the most important thing for go-storage. However, we do have different capabilities and limitations in storage services. So the problem comes into how to handle them.

Our goals are:

- Behavior consistent by default (invalid operations should be return error)
- Give user abilities to loose the restriction
- Allow user to enable the features they want

In [GSP-16], we introduced loose mode which is a global flag that controls the behavior when services meet the unsupported pairs.

- If `loose` is on: This error will be ignored.
- If `loose` is off: Storager returns a compatibility related error.

However, we removed `loose` in [GSP-20] because we can't figure out [error could be returned too early while in loose mode](https://github.com/beyondstorage/go-storage/issues/233). And loose is so general that it affects nearly all behavior of Storager.

In [types: Implement pair policy](https://github.com/beyondstorage/go-storage/pull/453), we try to figure out this problem by introducing `PairPolicy`.

`PairPolicy` controls the behavior of pairs:

```go
type PairPolicy struct {
	All bool
	
	...

	// pairs for interface Storager
	Create           bool
	Delete           bool
	List             bool
	ListListMode     bool
	Metadata         bool
	Read             bool
	ReadSize         bool
	ReadOffset       bool
	ReadIoCallback   bool
	Stat             bool
	Write            bool
	WriteContentType bool
	WriteContentMd5  bool
	WriteIoCallback  bool
}
```

If `PairPolicy.Write` is on, Storager will return `services.PairUnsupportedError` while meeting not supported pairs.

But it doesn't solve the problem:

- `PairPolicy` is generated in go-storage and can't reflect the capabilities in service.
- `PairPolicy` only used for pair capabilities check, and can't fix the problem introduced in [GSP-86]

## Proposal

So I propose to treat loose behavior consistence as a feature, and introduce feature gates in [go-storage]. 

### New Feature: Loose Operation

By default, [go-storage] will return errors for not supported pairs. If loose operation feature has been enabled, [go-storage] will ignore those errors.

To enable loose operation, users need to add pairs like:

```go
s3.WithStorageFeatures(s3.StorageFeatures{
	LooseOpeartionAll: true,
	LooseOperationWrite: true,
})
```

> `New` function is special, and we will always enable `loose` for it.

### New Feature: Virtual Operation

```go
s3.WithStorageFeatures(s3.StorageFeatures{
	VirtualOperationAll: true,
})
```

### Feature Struct

[go-storage] will generate feature structs for all services.

```go
type StorageFeatures struct {
	LooseOpeartionAll bool
	LooseOperationWrite bool
	
	VirtualOperationAll bool
	VirtualOperationCreateDir bool
}
```

## Rationale

<proposal's rationale content, other implementations>

## Compatibility

This proposal will deprecate `PairPolicy`.

## Implementation

<proposal's implementation>

[GSP-16]: ./16-loose-mode.md
[GSP-20]: ./20-remove-loose-mode.md
[GSP-86]: https://github.com/beyondstorage/specs/pull/86
[go-storage]: https://github.com/beyondstorage/go-storage