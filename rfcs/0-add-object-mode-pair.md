---
author: Xuanwo <github@xuanwo.io>
status: draft
updated_at: 2021-06-08
---

# Proposal: Add ObjectMode Pair

## Background

We have introduced `ObjectMode` in [GSP-25]:

```go
type ObjectMode uint32

const (
	ModeIrregular ObjectMode = 0

	ModeDir ObjectMode = 1 << iota
	ModeRead
	ModeLink
)
```

But it's not enough, we need `ObjectMode` pair too. For example, `Create` operation needs `ObjectMode` to create a new object. And `ObjectMode` COULD be used as input restriction for operations. Let me declare the later case more clearly.

[GSP-49] added create dir operation, and [GSP-87] introduced features gates. With those GSP, our user can use `CreateDir` to create a dir object. In services like s3 which doesn't have native `CreateDir` support, we usually simulated it via create an object ends with `/`. But the behavior doesn't work in `Stat` and `Delete`.

- If user call `Stat("test")` after `CreateDir("test")`, he will get `ObjectNotExist` error.
- If user call `Delete("test")` after `CreateDir("test")`, no object will be removed and `test/` will be kept in service.

We need to find a way to figure it out.

## Proposal

So I propose to add `ObjectMode` pair. This pair COULD be used in following operations:

- `Create`: set the output object's `ObjectMode`
- `Stat`: ObjectMode hint
- `Delete`: ObjectMode hint

For `Stat` and `Delete`

- Service SHOULD use `ObjectMode` pair as hint.
- Service could have different implementations for different `ObjectMode`.
    
Take `s3` as example, we simulate `CreateDir` via create object ends with `/`. `CreateDir("test")` will create an object `test/` in s3. And we can 

- stat this object via `Stat("test", pairs.WithObjectMode(types.ObjectModeDir))`
- delete this object via `Delete("test", pairs.WithObjectMode(types.ObjectModeDir))`

## Rationale

<proposal's rationale content, other implementations>

## Compatibility

<proposal's compatibility statement>

## Implementation

<proposal's implementation>


[GSP-25]: ./25-object-mode.md
[GSP-49]: ./49-add-create-dir-operation.md
[GSP-87]: ./87-feature-gates.md