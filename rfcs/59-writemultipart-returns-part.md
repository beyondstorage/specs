---
author: JinnyYi <github.com/JinnyYi>
status: draft
updated_at: 2021-05-11
---

# AOS-59: WriteMultipart Returns Part

## Background

`Multiparter` is designed for multipart upload. Multipart upload is a three-step process: use `CreateMultipart` to initiate the upload, use `WriteMultipart` to upload the object parts, and after you have uploaded all the parts, use `CompleteMultipart` to complete the multipart upload. You can use `ListMultipart` to list all of your in-progress multipart uploads or get a list of the parts that you have uploaded for a specific multipart upload.

`CompleteMultipart` request must include the upload ID and a list of both part numbers and corresponding ETag values returned after that part was uploaded in some services. The ETag uniquely identifies the combined object data, not necessarily an MD5 hash of the object data. We need return ETag that we got from services to make it possible. 

## Proposal

So I propose to change `WriteMultipart` to: 

```go
WriteMultipart(o *Object, r io.Reader, size int64, index int, pairs ...Pair) (n int 64, p Part, err error)
```

`Part` defines the upload part:

```go
type Part struct {
	Index int
	Size  int64
	ETag  string
}
```

For service provider:

- Services SHOULD return `Part` if there are limitations for multipart upload.

For user:

- Users SHOULD hold and pass in the complete parameter `Part` to `CompleteMultipart` returned by `WriteMultipart`.

## Rationale

This design is highly influenced by [Amazon Simple Storage Service API](https://docs.aws.amazon.com/AmazonS3/latest/userguide/mpuoverview.html).

### Alternative Way: Call `ListMultipart` before `CompleteMultipart`

The alternative way is call `ListMultipart` to obtain `Part` list before `CompleteMultipart`.

This will need extra requests.

## Compatibility

This is a breaking change, we should

- Bump go-storage to v4.
- Bump all service that already tagged to a higher major version.

## Implementation

- Update [go-storage](https://github.com/aos-dev/go-storage)
  - Update the return value of `WriteMultipart`.
- Update [go-integration-test](https://github.com/aos-dev/go-integration-test)
  - Add the argument ETag in `CompletePart` test case.
- Make sure all services implement `WriteMultipart` and `CompleteMultipart` correctly.