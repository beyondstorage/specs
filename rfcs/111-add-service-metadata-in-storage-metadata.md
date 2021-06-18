---
author: JinnyYi <github.com/JinnyYi>
status: draft
updated_at: 2021-06-17
---

# GSP-111: Add Service Metadata In Storage Metadata

## Background

In [GSP-6], we split storage meta and object meta. In [GSP-40], we split all metadata into four groups: `global metadata`, `standard metadata`, `service metadata` and `user metadata`. For storage, there's no `standard metadata` and user defined metadata is not allowed.

For now, storage related information carried in `StorageMeta`:

```go
type StorageMeta struct {
	location string
	Name     string
	WorkDir  string

	// bit used as a bitmap for object value, 0 means not set, 1 means set
	bit uint64
	m   map[string]interface{}
}
```

It only defines the `global metadata` for storage which could be used in all services, but no field to carry the metadata used in special services.

## Proposal

So I propose to add `service metadata` in storage metadata.

### Add service metadata in storage metadata

The `global metadata` has been included in object type. `service metadata` for storage will be stored in a struct that defined by service:

- `service metadata` will be `serviceMetadata interface{}` in `StorageMeta`.

For services:

Example of adding service's pair for storage metadata in `service.toml`:

```go
[infos.storage.meta.<service-meta>]
type = "<type>"
description = "<description>"
```

For `service metadata`, we will introduce `Strong Typed Service Metadata`. We will generate following struct for specific service according to service's pairs:

```go
type StorageServiceMetadata struct {
    <service meta>  <type>
    ...
}
```

And add following generated functions in service packages:

```go
// Only be used in service to set ServiceMetadata into StorageMeta.
func setStorageServiceMetadata(s *StorageMeta, sm StorageServiceMetadata) {}
// GetStorageServiceMetadata will get ServiceMetadata from StorageMeta.
func GetStorageServiceMetadata(s *StorageMeta) StorageServiceMetadata {}
```

### Rename `ObjectMetadata` to `ObjectServiceMetadata`

For now, `ObjectMetadata` is used to set `service metadata` for `Object`. To avoid confusion, we will rename `ObjectMetadata` to `ObjectServiceMetadata`. Correspondingly, the getter and setter will be updated as follows:

```go
// Only be used in service to set object service metadata.
func setObjectServiceMetadata(o *Object, om *ObjectServiceMetadata) {}
// Only be used outside service package to get object service metadata.
func GetObjectServiceMetadata(o *Object) *ObjectServiceMetadata {}
```

## Rationale

This design is highly influenced by [GSP-40].

## Compatibility

All API call that used object service metadata could be affected.

## Implementation

- `specs`
  - Add `service-metadata` with type `any` in [info_storage_meta.toml].
- `go-storage`
  - Generate `serviceMetadata` and corresponding `get/set` functions into `StorageMeta`.
  - Support generate `StorageServiceMetadata` distinguish with `ObjectServiceMetadata` by `Infos.Scope`.


[GSP-6]: ./6-normalize-metadata.md
[GSP-40]: ./40-unify-object-metadata.md
[info_storage_meta.toml]: ../definitions/info_storage_meta.toml
