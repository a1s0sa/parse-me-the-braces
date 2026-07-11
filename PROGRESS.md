# Go Learning Journal — parse-me-the-braces

**Read this file first, before doing anything else, at the start of every session in this repo.**
It is the source of truth for where we are. Claude's cross-session memory may still hold
general preferences, but curriculum progress lives here, not there.

## Ground rules (Go mentor mode)

- User is a Go beginner. Default teaching style is Socratic: ask guiding questions first,
  let the user attempt the reasoning, then confirm/correct. Do not jump to the answer.
- Do not directly edit/write code files unless the user explicitly says to switch modes
  for the session. This includes not scaffolding new files unprompted.
- Genuine conceptual questions ("how does X work in Go") get answered directly and
  thoroughly — the mentor stance is about not doing the user's debugging/fixing for them,
  not about withholding factual teaching.
- End explanations with a small follow-up question or prediction task to check understanding.
- User writes all the code themselves.

## Curriculum overview

Three independent 6-stage tracks, each finished fully before the next starts, in this order:

1. **JSON Parser** (this repo) — current, in progress
2. **KV Cache** — next, not started
3. **Job Scheduler / Worker Queue** — last, not started

Each track ends in a concurrency capstone that is a *distinct, domain-natural* pattern
(not a copy-pasted worker pool every time):

- Parser → producer/consumer pipeline (tokenizer goroutine streams tokens over a channel
  to a parser goroutine).
- Cache → shared-state protection (sharded/mutex-protected map + background goroutine
  sweeping expired TTL entries).
- Scheduler → worker pool with retries, per-job panic recovery, context cancellation.

User wants highest difficulty at every stage of every track.

## Track 1: JSON Parser — the 6 stages

1. **Basics** — Tokenizer: pure functions doing lexical scanning of a byte/rune stream
   into tokens (`{`, `}`, string, number, etc.), multiple return values
   (`token, rest, ok`), constants for token kinds.
2. **Flow control** — Recursive-descent parser: `switch` on token type, `if`/`for` to
   consume array/object elements, `defer`+`panic`/`recover` as the real error-propagation
   mechanism (idiomatic for recursive-descent parsers here, not just an exercise). Output:
   a generic tree of `map[string]interface{}` / `[]interface{}`.
3. **More types** — Formalize into real AST types: `JSONObject map[string]Value`,
   `JSONArray []Value`, struct-based nodes; closures for a tree-walking visitor.
4. **Methods & interfaces** — `Value` interface with type switches, `Stringer`, real
   error types with position info.
5. **Generics** — generic reducer/walker over `Value` trees.
6. **Concurrency (capstone)** — tokenizer and parser run concurrently as a
   producer/consumer pipeline connected by a channel.

## Track 2: KV Cache (later)

Struct-backed-by-a-map cache → `Store` interface with swappable eviction (LRU/TTL) +
Stringer + custom errors → generic `Cache[K comparable, V any]` → capstone:
sharded/mutex-protected concurrent cache with a background TTL-sweeper goroutine.
Optional stretch once Track 1 is done: snapshot the cache to disk using the Track 1
parser's own JSON encoder instead of `encoding/json`.

## Track 3: Job Scheduler / Worker Queue (later)

Job structs, priority vs FIFO queue → `Scheduler`/handler interfaces → generic
`Queue[T]` → capstone: worker pool with retries/backoff, per-job panic recovery,
context-based cancellation. Optional stretches once Tracks 1 & 2 are done: define jobs
in JSON parsed by the Track 1 parser; use the cache for job dedup/result storage.

## Current status

- **Track:** 1 — JSON Parser
- **Stage:** 1 — Basics / Tokenizer
- **Repo layout so far:** `cmd/main.go` (entry point), `internal/tokenizer/tokenize.go`
  (tokenizer package, in progress). Module renamed from `my_go_project` to
  `parse-me-the-braces`. Old `main.go` (root) and `math/random/random.go` from the
  starter scaffold have been removed in favor of this layout.
- **Resolved:** `IsDigit(numChar uint8) bool` now correctly checks `numChar >= 48 &&
  numChar <= 57` (ASCII range for `'0'`-`'9'`), dropping the dead `>= 0` check (always
  true for an unsigned type). User self-diagnosed both bugs via guided questions:
  string indexing yields byte/ASCII values not literal ints, and `uint8`'s floor is 0.
  `main.go` now tests it with real rune literals (`'1'`, `'A'`) instead of a bare int.
- **Resolved:** `IsWhiteSpace` and `IsPunctuation` written, both using `switch` with
  multi-value `case` lists returning directly (no intermediate bool var), and character
  literals (`'\t'`, `'['`, etc.) instead of raw decimal ASCII codes — confirmed via a
  factual detour that untyped rune constants convert automatically to fit `uint8`/`int8`,
  but named-type variables (like `input[i]` being `byte`/`uint8`) never convert
  implicitly, which is why `uint8` (not `int8`) is the right param type throughout.
  Also caught and fixed a raw literal tab byte accidentally typed into a `case` value —
  replaced with the `\t` escape for readability/robustness.
- **Next up (not started):** `IsQuote` (simple boundary-detection classifier, same
  shape as the others), then a letter classifier for recognizing the bare keywords
  `true`/`false`/`null`, then the actual scan loop that walks the input and emits
  tokens (multiple return values: token, rest, ok). Note: unlike the classifiers,
  extracting string *content* between quotes will need a function that consumes
  multiple characters in a loop rather than a single-byte bool check.

## Milestone log

### 2026-07-10
- Project skeleton created: `cmd/main.go`, `internal/tokenizer` package, module
  renamed to `parse-me-the-braces`.
- Wrote first version of `IsDigit`; started debugging it as the first Stage 1 exercise.
- Created this file to track curriculum progress in-repo instead of relying on
  cross-session assistant memory.

### 2026-07-11
- Fixed `IsDigit`: corrected range to ASCII `48`-`57`, removed dead `numChar >= 0`
  check. Self-diagnosed via guided questions (string indexing → byte/ASCII values;
  `uint8` floor is 0). Verified with a real test in `main.go` using rune literals.
- Wrote `IsWhiteSpace` and `IsPunctuation`; refactored all three classifiers to use
  character literals instead of decimal ASCII codes. Covered untyped-constant vs.
  named-type conversion rules along the way, and fixed a stray literal tab byte in
  source. Next: `IsQuote`, then letters for `true`/`false`/`null`, then the scan loop.
