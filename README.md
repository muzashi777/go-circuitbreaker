# Go Circuit Breaker & Retry Demo

A small demonstration of circuit breaker and retry patterns implemented in Go.

This repository contains a lightweight example showing how to use the `gobreaker` and `retry-go` libraries through a tiny project layout with a reusable `circuitbreaker` package and a `mainmodule` that runs a sample service.

**Contents**

- `circuitbreaker/` — Core implementations and examples for circuit breaker and retry logic.
- `circuitbreaker/_config/` — Example configuration helpers used in the demo (`go-breakger.go`, `go-retry.go`).
- `circuitbreaker/service1/` — A sample service that exercises the circuit breaker / retry code.
- `entity/` — Shared entities and error types used by the demo.
- `mainmodule/` — Small main program that wires the demo and runs `service1.Service1()`.

**Prerequisites**

- Go 1.25 (the project `go.mod` files target `go 1.25.4`).

**Build & Run**

From the repository root you can run the demo main program directly:

```bash
cd ./mainmodule
go run .
```

The `mainmodule` uses `replace` statements in its `go.mod` to point at the local `circuitbreaker` module so no additional `go get` is required for local development.

**What this demo shows**

- How to integrate `github.com/sony/gobreaker` for circuit breaking behavior.
- How to integrate `github.com/avast/retry-go` for retry policies.
- How to structure a small Go project with a reusable package and a simple runner.

**Key files to inspect**

- `circuitbreaker/gobreaker.go` — Circuit breaker wrapper and example usage.
- `circuitbreaker/retry.go` — Retry helper functions.
- `circuitbreaker/_config/go-breakger.go` and `_config/go-retry.go` — Example configurations used by the demo.
- `circuitbreaker/service1/service1.go` — Example service called by `mainmodule`.
- `mainmodule/main.go` — Entry point that runs the service demo.



**Contributing**

Feel free to open issues or pull requests. If you add features, please include tests and update this README with any new usage steps.

**License**

This demo does not include a license file. Add a `LICENSE` file if you want to specify terms for reuse.


