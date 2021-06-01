---
author: xxchan <xxchan22f@gmail.com>
status: draft
updated_at: 2021-06-01
deprecates:
  - 13-remove-config-string.md
---

# GSP-90: Parse Pairs from String

## Background

We had config strings like `s3://bucket_name?access_key=xxxx&secret_key=xxxx` before. This is introduced in [GSP-3] but deprecated by [GSP-13].

There are two reasons why config strings are removed, as mentioned in [GSP-13]:

1. Usability: If an application uses [go-storage]'s config string, then it either exposes the format of config string to end user directly (additional burden for end users) or it has to manipulate its config into our format (additional burden for application developers). And if the application wants to combine its other config into a config string, this is also inconvenient.
2. Implementation: We didn't find a good way to parse pairs from strings easily and safely.

So we let the users of [go-storage] construct pairs directly using `WithXxx()`. 

But passing string config is indeed more convenient, if not much. And actually parsing string into pairs is not that difficult. We can use a registry mechanism like [GSP-48]'s implementation.

## Proposal

So I propose the following mechanism for parsing pairs:

### Pair Registry

We first register the types of global pairs and service pairs.

From [go-storage] side:

```go
type PairMap map[string]reflect.Type

func RegisterServicePairMap(ty string, m PairMap) {}
```

From services side, we can generate following code:

```go
func servicePairs() PairMap {}

func init() {
	// ...
	pairs.RegisterServicePairMap("<type>", servicePairs())
}
```

The pair information maps are also generated.

### Parse According to Type

For each type, we can have a parse function.

```go 
type parseFunc func(string) (interface{}, error)

var parseMap map[reflect.Type]parseFunc
```

Then we can parse a pair of string into a pair. We check the pair's name in a `PairMap` and get its type information, then use corresponding parse function to parse the pair's value.

```go
func (m PairMap) parse(k string, v string) (pair Pair, err error) {}
```

Then we can easily provide APIs for users to parse pairs:

```go 
func Parse(ty string, k string, v string) (pair Pair, err error) {}
func ParseMap(ty string, m map[string]string) (ps []Pair, err error) {}
```

## Rationale

We split the problem of parsing pairs into two orthogonal sub-problems to make things elegant and easy to do.

This proposal can function as a basis for further utilities on top of it. After this work, we can bring back config string by first parsing a config string into a `(string, map[string]string)`, then parsing the string map into pairs. Since more pairs can be added into those got from the config string easily, this approach has better composability than the old one, where services are directly opened with the config string. Of course, we can still provide the direct way as an option.

Possible problems:

1. In the pair registry `PairMap`, we may need more pair information beside types in the future, so we can also create a custom type `PairInfo` as the value type of the map. But to keep things simple, we'd better not overdesign now.
2. On how to recognize a service pair, we have these alternatives:
   1. Always specify service type like `Parse("s3", "pair_name", v)`. It is possible that a service pair have the same name as a global pair. We should forbid this situation. We can add check in service pair registry.
   2. Service pairs should always have a service type prefix, like `Parse("s3.pair_name", v)`. This way has no ambiguity, but is a little bit inconvenient.
   3. Support ambiguous parsing. We can support both `Parse("s3.pair_name", v)`, `Parse("s3", "pair_name", v)` and `Parse("pair_name", v)`. In the last one, we will try to guess the service type.
   
   The first one seems to be both user-friendly and easy to implement so it is the choice in the proposal.

## Compatibility

New additional utility, no break change.

## Implementation

- Implement pair registry and parsing in [go-storage].
- Implement service code generate in [go-storage] definitions.
- Make sure all service has been updated.


[GSP-3]: ./3-support-service-init-via-config-string.md
[GSP-13]: ./13-remove-config-string.md
[GSP-48]: ./48-service-registry.md
[go-storage]: https://github.com/beyondstorage/go-storage
