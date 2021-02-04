---
author: Xuanwo <github@xuanwo.io>
status: draft
updated_at: 2021-02-03
---

# Proposal: Unify Object Metadata

## Background

Object may have different kind of metadata:

- last updated time
- content length
- content type
- storage class
- ...

Some of them are defined in standards, such as: `content-type`, `content-length` which defined in [RFC 2616](https://tools.ietf.org/html/rfc2616). Some of them are only used in special services, such as: `x-amz-storage-class` which only used in the AWS S3 service.

In order to unify object metadata behavior so that we can handle them in the same way, we need to clearly define and classify them.

[Proposal: Normalize Metadata](./6-normalize-metadata.md) is an attempt, it normalized metadata by [message-header](https://www.iana.org/assignments/message-headers/message-headers.xhtml), and handle all private metadata seperately.

- `Content-MD5` -> `content-md5`
- `x-aws-storage-class` / `x-qs-storage-class` -> `storage-class`

The problem is: How to implement cross-storage operations？

Giving a set object, how to migrate them to another storage services without changing it's behavior? For example, permission on local file system, storage class on object storage class.


## Proposal

Split all metadata into three groups:

- global metadata: defined via framework
    - id
    - link target
    - mode
    - multipart id
    - path
    - ...
- standard metadata: defined via existing RFCs
    - content md5
    - content type
    - etag
    - last-modified
    - ...
- service metadata: defined via service implemantations
    - x-amz-storage-class
    - x-qs-storage-class
    - ...
- user metadata: defined via user input.

`global` and `standard` metadata will be included in object type. `service` and `user-defined` metadata will be stored in a hash map inside object type.

Use `golang` as example:

- `id` will be `id string`
- `content md5` will be `contentMd5 string`
- service metadata will be `serviceMetadata map[string]interface{}`
- user metadata will be `userMetadata map[string]string`

## Rationale

TODO

## Compatibility

TODO

## Implementation

TODO