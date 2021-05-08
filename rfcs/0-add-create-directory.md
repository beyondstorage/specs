---
author: Xuanwo <github@xuanwo.io>
status: draft
updated_at: 2021-05-08
---

# Proposal: Add CreateDir Operation

## Background

<proposal's background>

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