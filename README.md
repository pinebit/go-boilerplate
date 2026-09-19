# go-boilerplate

A small Go HTTP server template using the standard library for routing, structured
logging, dependency wiring, and graceful shutdown. TOML configuration and
Prometheus metrics are the only direct external dependencies.

## Quick start

With Go 1.27 or newer installed, run from the repository root:

```sh
go mod download
make run
```

In another terminal, check the greeting and metrics:

```sh
curl http://localhost:3333/
curl http://localhost:3333/metrics
```

Press Ctrl+C to shut down gracefully.

## Requirements and development

The module requires Go 1.27.0; the Docker builder uses Go 1.27.1. Docker is
optional for local development. Race-enabled tests require CGO and a C compiler.
Run commands from the repository root:

```sh
make build          # Build ./boilerplate
make run            # Run with config/config.toml
make test           # Test all packages with the race detector
make vet            # Run Go's static checks
make fmt            # Format Go files
make lint           # Run vet and golangci-lint v2
```

Install the linter version used by CI:

```sh
go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.13.2
```

`GET /` returns `Hello world!`. `GET /metrics` exposes the `hello_world`
counter plus Go runtime and process metrics. Unknown paths return 404;
unsupported methods return 405. Standard-library GET patterns also accept HEAD.

## Configuration

```sh
go run . -config config/config.toml
```

The included [sample configuration](config/config.toml) is:

```toml
DevMode = true
ShutdownTimeout = '5s'

[HttpServer]
Address = ''
Port = 3333
ReadHeaderTimeout = '5s'
ReadTimeout = '15s'
WriteTimeout = '30s'
IdleTimeout = '60s'
```

The configuration path defaults to `config/config.toml`. The sample listens on
port 3333; `NewDefaultConfig` uses 8080. Empty `HttpServer.Address` listens on all
interfaces; use `127.0.0.1` for local-only access. Partial TOML files inherit
defaults. Unknown keys, invalid durations, nonpositive timeouts, and port zero
are rejected before startup. Missing configuration files are errors.

Set `DevMode = false` for JSON logs; development uses readable text and enables
debug logging. Configure HTTP header, read, write, idle, and shutdown timeouts in
the sample TOML. SIGINT and SIGTERM stop new connections and allow active requests
to finish until `ShutdownTimeout`, after which connections are closed. Startup
and shutdown failures exit with a nonzero status.

## Containers

```sh
docker build -t boilerplate .
docker run --rm -p 3333:3333 boilerplate
# Mount a custom configuration:
docker run --rm -p 3333:3333 \
  -v "$PWD/config/config.toml:/app/config/config.toml:ro" boilerplate
```

The image includes the sample configuration and CA certificates and runs as a
non-root user. Ensure mounted configuration files are readable by UID 65532.
Set `DevMode = false` in production and restrict access to `/metrics` at your
network or reverse proxy boundary. CI builds the image and checks startup,
endpoints, and graceful shutdown.

## Project structure

- `main.go`: CLI, signals, and explicit dependency wiring.
- `config/`: defaults, validation, TOML, and duration helpers.
- `logger/`: `log/slog` setup.
- `services/http/`: `net/http` routes, metrics, and lifecycle.

Extend the `ServeMux` in `services/http/router.go` and wire new dependencies
explicitly in `main.go`. When using this template for another project, update the
module path in `go.mod`, internal imports, and the binary and image names.

## Migration from the original template

Gin, Zap, Dig, and go-boot have been removed. Constructors now accept
`*slog.Logger`, `NewRouter` returns `http.Handler`, and `Server.Run` owns the
blocking lifecycle. Tests use Go's `testing` package without Testify. Existing
partial TOML configurations inherit the new timeout defaults; unknown keys are
now rejected.

## License

Copyright (c) 2026 Andrei Smirnov. Licensed under the [MIT License](LICENSE).
