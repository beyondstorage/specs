---
author: Xuanwo <github@xuanwo.io>
status: draft
updated_at: 2021-05-08
---

# Proposal: Add CreateDir Operation

## Background

Applications need the ability to create a directory. For now, our support is a bit wired.

In [fs](https://github.com/aos-dev/go-service-fs), we support `CreateDir` by dirty hack:

```go
if s.isDirPath(rp) {
    // FIXME: Do we need to check r == nil && size == 0 ?
    return 0, s.createDir(rp)
}
```

In other storage service, user needs to create by special `content-type`:

```go
store.Write("abc/", nil, 0, ps.WithContentType("application/x-directory"))
```

## Proposal

So I propose to add a new operation `CreateDir`:

```go
type Direr interface {
	CreateDir(path string, pairs ...Pair) (o *Object, err error)
}
```

## Rationale

<proposal's rationale content, other implementations>

## Compatibility

<proposal's compatibility statement>

## Implementation

<proposal's implementation>