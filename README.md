# service-msgr

## Introduction

A simple messenger API.

## Installation

Install Go: https://golang.org/doc/install

Install golangci-lint: https://github.com/golangci/golangci-lint

## Features

This template includes:

- HTTP Server
- Configuration file imports
- Request logging
- Response encoding
- Graceful shutdown

The following third-party packages are used:
- [github.com/spf13/viper](https://github.com/spf13/viper) (Configuration)
- [github.com/go-chi/chi](https://github.com/go-chi/chi) (Routing and middleware)
- [github.com/uber-go/zap](https://github.com/uber-go/zap) (Logging)
- [github.com/InVisionApp/go-health](https://github.com/InVisionApp/go-health) (Health checking)

Refer to the documentation above for implementation details.

## Run

```bash
go build -o ./bin/serve ./cmd/serve
./bin/serve
```

## Environment Variables
#### Runtime

| name                     | description                                                     | type    | optional | default     |
|--------------------------|-----------------------------------------------------------------|---------|----------|-------------|
| PORT                     | The server port                                                 | string  | yes      | 3000        |
| LOGGER                   | The logger type (DEVELOPMENT, PRODUCTION)                       | string  | yes      | DEVELOPMENT |
| ENV_FILE                 | location of the .env file                                       | string  | yes      | .env        |
| HEALTH_CHECK_ENDPOINT    | Endpoint to send a GET requeset for the health of service       | string  | yes      | /healthz    |
| READY_CHECK_ENDPOINT     | Endpoint to send a GET requeset for the services ready check    | string  | yes      | /readyz     |


## Test

```bash
go test ./...
```

## Linter

The Product Service uses [golangci-lint](https://github.com/golangci/golangci-lint) to lint the project files. It runs any time a change is pushed to the remote repository.

To run:
```bash
make lint
```

The output will list any issues that must be addressed.