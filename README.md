# HugeINT

Arbitrary-precision unsigned integer arithmetic for Go, built from scratch on top of
primitive machine-word operations — a software full-adder over `uint32` words with
ripple carry — rather than wrapping the standard `math/big`.

> **Status: early work in progress.** The building blocks live under `internal/`;
> there is no stable public API yet and `main` does not currently build (see below), so
> it is not usable as a dependency at this stage.

## What's implemented

- **Word arithmetic** (`internal/instruction`) — addition of 32-bit words returning a
  `(high, low)` pair so the carry survives, in two forms (unrolled and looped) with
  benchmarks. Further operations — subtract, compare, multiply, divide, shifts,
  invert — are in progress on the `test` branch.
- **Word storage** (`internal/storage`) — a concurrency-safe singly-linked list of
  `uint32` nodes that holds a number's magnitude.

## Build & test

```bash
make test    # go test -race ./...
make bench   # benchmarks
make cover   # coverage report
```

`main` is mid-refactor: `internal/instruction` does not compile there yet (a constants
file is missing). The continuation, including a first public API, lives on the `test`
branch.

## License

[MIT](./LICENSE) © Seilbek
