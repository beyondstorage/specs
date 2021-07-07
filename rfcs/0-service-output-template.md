- Author: Xuanwo <github@xuanwo.io>
- Start Date: 2021-07-07
- RFC PR: [beyondstorage/specs#0](https://github.com/beyondstorage/specs/issues/0)
- Tracking Issue: [beyondstorage/go-storage#0](https://github.com/beyondstorage/go-storage/issues/0)

# GSP-0: Service Output Template

## Background

[Increase code coverage](https://github.com/beyondstorage/go-storage/issues/620) plans to move `cmd/definitions/tests` into `tests` and support generating files into `generate_test.go` so we won't export test services.

So `definitions` should have the ability to generate all files into a specified output filename template.

## Proposal

I propose to add a new field `output_template` in `service.toml`. In `output_template`, we will use format string like `%s_test.go`.

If service provider specify `output_template` to `%s_test.go`, `definitions` will generate files into:

- `generated_test.go`
- `service_test.go`
- `storage_test.go`

## Rationale

N/A

## Compatibility

This GSP introduces a compatible change.

## Implementation

- Add `output_template` field in specs.
- Support `output_template` in go-storage.
