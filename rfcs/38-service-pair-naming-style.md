---
author: Xuanwo <github@xuanwo.io>
status: draft
updated_at: 2021-04-21
---

# AOS-38: Service Pair Naming Style

## Background

There are 10 services have been implemented and more services are on the way, and every service may have their only pairs. We need to design a naming style that meets following needs:

- No ambiguity: the pair should be confused with other pairs.
- Cross Language: the pair's name should not contain the language related details.
- Easy to follow and learn: the pair's naming style should be easy to follow.

## Proposal

So I propose following service pair naming style.

### Scope

the style only applies for service's pairs and infos, global pairs' name should be discussed in related RFCs.

### Rule

We should adopt a style called `API Native Style` which means our pairs name should be native in original API.

For example, AWS S3 support three kinds of [Server-side encryption](https://docs.aws.amazon.com/AmazonS3/latest/userguide/serv-side-encryption.html)

- Server-Side Encryption with Amazon S3-Managed Keys (SSE-S3)
- Server-Side Encryption with Customer Master Keys (CMKs) Stored in AWS Key Management Service (SSE-KMS)
- Server-Side Encryption with Customer-Provided Keys (SSE-C)


For `SSE-S3`, S3 supports following HTTP headers:

- `x-amz-server-side-encryption` (should be `AES256`)

For `SSE-KMS`, S3 supports following HTTP headers:

- `x-amz-server-side-encryption` (should be `aws:kms`)
- `x-amz-server-side-encryption-aws-kms-key-id`
- `x-amz-server-side-encryption-context`
- `x-amz-server-side-encryption-bucket-key-enabled`

For `SSE-C`, S3 supports following HTTP headers:

- `x-amz-server-side-encryption-customer-algorithm`
- `x-amz-server-side-encryption-customer-key`
- `x-amz-server-side-encryption-customer-key-MD5`

So we should add following pairs:

- `server-side-encryption`
- `server-side-encryption-aws-kms-key-id`
- `server-side-encryption-context`
- `server-side-encryption-bucket-key-enabled`
- `server-side-encryption-customer-algorithm`
- `server-side-encryption-customer-key`
- `server-side-encryption-customer-key-md5`

There is an exception for Servicer & Storager's `new` operation's pairs. Those pairs should follow the SDK's option name.

## Rationale

N/A

## Compatibility

This naming style will affect all service pairs that not released.

## Implementation

This proposal is a rule and will be followed by all contributors.