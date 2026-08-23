# CLAUDE.md

Generic Go conventions (style, declaration order, test structure, godoc, error discipline,
build hygiene, organisation) come from the `stack-go` plugin skills and are not restated
anywhere in this repo.

## What this is

`hugeint` is a from-scratch arbitrary-precision unsigned integer library for Go. It
builds big-number arithmetic up from primitive machine-word operations — a bitwise
software full-adder over `uint32` words with ripple carry — rather than wrapping the
standard `math/big`. A number's magnitude is held as a sequence of 32-bit words in a
linked-list store.

## Current state — WIP, `main` does not build

Read this before touching anything:

- `internal/storage` builds and its tests pass.
- `internal/instruction` on `main` does NOT compile: `add.go` references `bit01..bit32`,
  but `constants.go` (which defines those masks) is missing from `main`.
- The missing file, plus a fuller instruction set (`cmp`, `sub`, `div`, `mul`, `sal`,
  `sar`, `invert`) and a first public API (`integer.go`, `pkg/`), live on the unmerged
  `origin/test` branch (commit `b3ab517`, "WIP"). That branch compiles but is itself
  red — `TestSubV1` fails (subtracting a larger operand from a smaller one is wrong).
- No decision has been made on reconciling `main` and `test`. Do not merge, restore
  the missing file, or "fix" arithmetic without an explicit plan and the owner's
  sign-off.

## Build & test

```bash
make test      # go test -v -race -timeout 30s ./...
make bench     # go test -benchmem -bench=. ./...
make cover     # go test -v -cover ./...
make install   # go install
make help      # list targets
```

## Architecture

| Package | Role |
|---------|------|
| `internal/instruction` | Primitive arithmetic on 32-bit words, returning a `(high, low)` word pair so carry/overflow is preserved. `Add` is a software full-adder — `addV1` (unrolled ×32) and `addV2` (looped), benchmarked against each other. |
| `internal/storage` | Concurrency-safe singly-linked list of `uint32` word-nodes (RWMutex-guarded `Push`/`Pop`/`Begin`/`End`/`ForEach`/`Free`) — the backing store for a number's words. |

Everything is under `internal/`; there is no importable public API on `main` yet (the
`integer.go` / `pkg/` surface exists only on `origin/test`).

## Working agreement

Plan-first pipeline; the canonical procedure is the `pipeline:working-agreement` skill — load it
before starting non-trivial work. Project delta:

- **Gate:** `make test` — **currently red**, because `main` does not compile. A red gate here is
  pre-existing, not something your change broke; read "Current state" above before acting on it.
- **Lenses:** standard set — see `pipeline:working-agreement`, which includes lens O.
- **Branching:** standard (`type/<issue>-<slug>`, PR into `main`).
