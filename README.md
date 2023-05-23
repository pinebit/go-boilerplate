# go-boilerplate

[![Go](https://github.com/pinebit/go-boilerplate/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/pinebit/go-boilerplate/actions/workflows/go.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Golang server boilerplate.

## Motivation

To create a battle-tested server template that implements the most common best practices.

## Features

* Dockerfile
* TOML configuration with [go-toml](https://github.com/pelletier/go-toml)
* Structured logging with [zap](https://github.com/uber-go/zap)
* Graceful boot/shutdown with [go-boot](https://github.com/pinebit/go-boot)
* [Prometheus](https://github.com/prometheus/client_golang) metrics
* Simple HTTP server with [gin framework](https://github.com/gin-gonic/gin)
* Dependency injection with [dig](https://github.com/uber-go/dig)

## Usage

1. Find and replace all "boilerplate" words for your product.
2. Building: `docker build .` or `make build`
3. Testing: `make test` and `make lint`
4. Running: `make run` or simply `go run .`
