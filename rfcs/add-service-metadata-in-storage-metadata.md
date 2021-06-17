---
author: JinnyYi <github.com/JinnyYi>
status: draft
updated_at: 2021-06-17
---

# GSP: Add Service Metadata In Storage Metadata

## Background

In [GSP-6], we split storage meta and object meta. In [GSP-40], we split all object metadata into four groups: `global` metadata, `standard` metadata, `service` metadata and `user` metadata. For storage, there's no `standard` metadata and user defined metadata is not allowed.

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

It only defines the `global` metadata for storage which could be used in all services, but no field to carry the metadata used in special services.

## Proposal

So I propose to add service metadata in storage metadata.

`service` metadata is provided by service, could only be used in current service.

The `global` metadata has been included in object type. `service` metadata will be stored in a struct that defined by service:

- service metadata will be `serviceMetadata interface{}` in `StorageMeta`.

For services side:

Example of adding service pair for storage metadata in `service.toml`:

```go
[infos.storage.meta.<service-meta>]
type = "<type>"
description = "<description>"
```

For service metadata, we will introduce `Strong Typed Service Metadata` like `ObjectMetadata`. We will generate following struct for specific service according to service pairs:

```go
type ServiceMetadata struct {
    <service meta>  <type>
    ...
}
```

And add following generated functions in service packages:

```go
// Only be used in service to set ServiceMetadata into StorageMeta.
func setServiceMetadata(s *StorageMeta, sm ServiceMetadata) {}
// GetServiceMetadata will get ServiceMetadata from StorageMeta.
func GetServiceMetadata(s *StorageMeta) ServiceMetadata {}
```

## Rationale

This design is highly influenced by [GSP-40].

## Compatibility

No break changes.

## Implementation

- `specs`
  - Add `service-metadata` with type `any` in [info_storage_meta.toml].
- `go-storage`
  - Generate `serviceMetadata` and corresponding `get/set` functions into `StorageMeta`.
  - Support generate `ServiceMetadata` distinguish with `ObjectMetadata` by `Infos.Scope`.


[GSP-6]: ./6-normalize-metadata.md
[GSP-40]: ./40-unify-object-metadata.md
[info_storage_meta.toml]: ../definitions/info_storage_meta.toml