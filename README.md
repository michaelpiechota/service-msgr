# service-msgr

## Introduction

A simple messenger API.

## Installation

Install Go: https://golang.org/doc/install

## Features

- HTTP Server
- Configuration file imports
- Request logging
- Response encoding
- Graceful shutdown

The following third-party packages are used:
- [github.com/spf13/viper](https://github.com/spf13/viper) (config)
- [github.com/go-chi/chi](https://github.com/go-chi/chi) (routing and middleware) (v5)
- [github.com/uber-go/zap](https://github.com/uber-go/zap) (structured logging)
- [github.com/InVisionApp/go-health](https://github.com/InVisionApp/go-health) (health checking)

## Run

```bash
make build
./bin/serve
```

## Clear Port (if needed)
```bash
kill -9 $(lsof -i:3000 -t) 2> /dev/null
```

## Environment Variables
#### Runtime

| name                     | description                                                     | type    | optional | default     |
|--------------------------|-----------------------------------------------------------------|---------|----------|-------------|
| PORT                     | The server port                                                 | string  | yes      | 3000        |
| LOGGER                   | The logger type (TEST, DEVELOPMENT)                             | string  | yes      | DEVELOPMENT |
| ENV_FILE                 | location of the .env file                                       | string  | yes      | .env        |
| HEALTH_CHECK_ENDPOINT    | Endpoint to send a GET requeset for the health of service       | string  | yes      | /healthz    |
| READY_CHECK_ENDPOINT     | Endpoint to send a GET requeset for the services ready check    | string  | yes      | /readyz     |


## Test

```bash
go test ./...
```