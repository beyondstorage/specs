---
author: xxchan <xxchan22f@gmail.com>
status: draft
updated_at: 2021-06-09
deprecates:
  - 13-remove-config-string.md
---

# GSP-90: Support Initialization Via Config String and Map

## Background

We had config strings like `s3://bucket_name?access_key=xxxx&secret_key=xxxx` before. This is introduced in [GSP-3] but deprecated by [GSP-13].

There are two reasons why config strings are removed, as mentioned in [GSP-13]:

1. Usability: If an application uses [go-storage]'s config string, then it either exposes the format of config string to end user directly (additional burden for end users) or it has to manipulate its config into our format (additional burden for application developers). And if the application wants to combine its other config into a config string, this is also inconvenient.
2. Implementation: We didn't find a good way to parse pairs from strings easily and safely.

So we let the users of [go-storage] construct pairs directly using `WithXxx()`. 

But passing string config is indeed more convenient, if not much. And actually parsing string into pairs is not that difficult. We can use a registry mechanism like [GSP-48]'s implementation.

## Proposal

So I propose to support service init from config string and map:

We add the following APIs:
```go
func NewServicerFromMap(ty string, m map[string]string) (types.Servicer, error) {}
func NewStoragerFromMap(ty string, m map[string]string) (types.Storager, error) {}

func NewServicerFromString(config string) (types.Servicer, error) {}
func NewStoragerFromString(config string) (types.Storager, error) {}
```

### Format

The format of the config string is:

TODO

### Implementation

`New*FromString` will first parse config string into `(ty string, m map[string]string)`, and then call `New*FromMap(ty, m)`.

`New*FromMap` will first parse `(ty string, m map[string]string)` into `ps []Pairs`, and then call `New*(ty, ps)`.
To support this, we have to know that a name is a pair (global or service), and its type, so we implement:

#### Pair Registry

We register the types of global pairs and service pairs.

From [go-storage] side:

```go
type PairMap map[string]string
func RegisterGlobalPairMap(m PairMap) {}
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

`PairMap`s are generated.

## Rationale

Possible problems:

1. In the pair registry `PairMap`, we may need more pair information beside types in the future, so we can also create a custom type `PairInfo` as the value type of the map. But to keep things simple, we'd better not overdesign now.

2. It is possible that a service pair have the same name as a global pair. Internally, the name of a service pair has prefix `<type>_` so it wasn't a problem.  
   
   But in `New*FromMap` and `New*FromString`, we won't want to let users add the prefix, so if two pairs have the same name, here comes ambiguity. So we should forbid a service pair from having the same name with a global pair. We can add check in service pair registry.

3. We can also export parsing APIs like 
   ```go 
   func ParseMap(ty string, m map[string]string) []Pair {} 
   func ParseString(config string) []Pair {}
   ```
   Pros:
   - Users can combine other pairs into a config string flexibly. (Solve the usability problem mentioned in the background section.)
   
   Cons:
   - We only want to make initialization easier, instead of providing a general pair parser, which may involve additional complexity. But if we provide an API like `ParseMap` then we can't stop users from using it in other places, and thus it's effectively a general pair parser.


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
