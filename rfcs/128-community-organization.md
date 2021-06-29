- Author: Xuanwo <github@xuanwo.io>
- Start Date: 2021-06-29
- RFC PR: [beyondstorage/specs#128](https://github.com/beyondstorage/specs/pull/128)
- Tracking Issue: [beyondstorage/go-storage#0](https://github.com/beyondstorage/go-storage/issues/0)

# GSP-128: Community Organization

## Background

BeyondStorage has been grown up, and we need to define our community organization to make it clear that everyone's rights and responsibilities.

## Proposal

I propose to add five positions for BeyondStorage, and they are separate in different projects.

For now, there are two main projects:

- `go-storage`: Focus on implement storage abstraction for golang, including `go-storage`, `go-service-*`, `go-endpoint`, `site`, `specs` and so on.
- `dm`: Focus on implement data migration services, including `dm`.

### Leader

- ID: `leader`
- Permissions: `Admin`

TBD

### Maintainer

- ID: `maintainer`
- Permissions: `Maintain`

TBD

### Committer

- ID: `committer`
- Permissions: `Write`

TBD

### Reviewer

- ID: `reviewer`
- Permissions: `Triage`

TBD

### Contributor

- ID: `contributor`
- Permissions: `Read`

TBD

## Rationale

N/A

## Compatibility

No code changes.

## Implementation

- Create a repo called `community` in which maintains all members and teams
- Implement team member permission automatically maintain in `go-community`.
