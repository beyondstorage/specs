---
author: JinnyYi <github.com/JinnyYi>
status: draft
updated_at: 2021-06-09
---

# GSP: Add Maximum File Size for Storage Write

## Background

Depending on the size of the data you are uploading, services offer different options, like upload a file in a single operation, in parts, or through append operation.

We use `Write` to handle single write operation. And the size is not unlimited.

- S3 alike storage services have 5GB limit for a single PUT operation.
- azblob limited to 5000 MiB <https://docs.microsoft.com/en-us/rest/api/storageservices/put-blob>

If we upload a file with size out of limit in a single operation, we will get an error like `Request Entity Too Large` in azblob. Same with `copy`,`fetch`, etc.

We should figure out the error before sending request to reduce the consumption of network resources.

## Proposal

So I propose to add maximum file size for storage write related operations.

A storager will be accompanied by the settings after initialization. We could add a variable with type `interface{}` into the current `Storage` structure in services to hold these settings and possible limitations in the future.

```go
type Storage struct {
	service *<service type>

	name    string
	workDir string

	m       interface{}

	defaultPairs DefaultStoragePairs
	features     StorageFeatures

	typ.UnimplementedStorager
	...
}
```

To assign the added variable, we could generate the `StorageMetadata` structure through the fields `[infos.storage.meta.<storage meta name>]` like the `ObjectMetadata`. Then assign the corresponding restriction constants to `StorageMetadata` and the variable when initiate `Storage`.

```go
[infos.storage.meta.<storage meta name>]
type = "<type>"
description = "<description>"
}
```

## Rationale

### Alternative Way: Ignore size limit

The alternative way is to send request directly without size check and get an error when the `size` larger then the limit.

This will need invalid request and get the ambiguous error.

## Compatibility

All changes are compatible.

## Implementation

- `go-storage` should support generate `StorageMetadata` distinguish with `ObjectMetadata` by `Infos.Scope`.
- `go-service-*`
   - Add the variable above and assign it on initialization.
   - Check `size` for write related operations and return `ErrRestrictionDissatisfied` when `size` out of limit.
