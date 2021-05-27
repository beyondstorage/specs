---
author: Xuanwo <github@xuanwo.io>
status: draft
updated_at: 2021-05-27
---

# Proposal: Add Create Link Operation

## Background

We have `ModeLink` for Object which means this Object is a link which targets to another Object. A link object could be returned in `Stat` or `List`. But there is no way to create a link object.

As discussed in [Link / Symlink support](https://github.com/beyondstorage/specs/issues/85), link related support is very different in different services:

- fs: Native support for hardlink and symlink
- oss: Native [Symlink](https://help.aliyun.com/document_detail/45126.html) API
- s3: No native support, [x-amz-website-redirect-location](https://docs.aws.amazon.com/AmazonS3/latest/userguide/how-to-page-redirect.html) can redirect pages but only works for website, see <https://stackoverflow.com/questions/35042316/amazon-s3-multiple-keys-to-one-object#comment57863171_35043462>
- `more information here`

## Proposal

I propose to add a new `CreateLink` operation:

```go
type Linker interface {
	CreateLink(src, dst string, pairs ...Pair) (o *Object, err error)
}
```

To implement this feature for those services that don't have native support, we will introduce a new pair:

```go
func WithEnableVirtualLink(v bool) Pair {
	return Pair{
		Key:   "enable_virtual_link",
		Value: v,
	}
}
```

- For services that have native support, they will ignore this option.
- For services that don't implement `Linker`, they will ignore this option.
- For services that don't have native support but implement `Link`
  - They MUST support `EnableVirtualLink` pair
  - If `EnableVirtualLink` is `true`, service should simulate a Link.
  - If `EnableVirtualLink` is `false`, service should return `ErrCapabilityInsufficient`.

For the implementations for simulated Link, services provider SHOULD choose the lowest cost method. For example, to implement `CreateLink` for s3, service provider should use user metadata instead of create an object with JSON.

## Rationale

### Introduce `EnableVirtualDir`

For object storage services, the `Dir` is also simulated. We may need to introduce `EnableVirtualDir` to make the behavior consistent.

## Compatibility

N/A

## Implementation

N/A