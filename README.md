# echo-server

A simple echo server that spits out exactly what you want from it.

## Table of Contents

- [How to use](#how-to-use)
  - [Install with Go](#install-with-go)
  - [Run it with Docker](#run-with-docker)
  - [Clone the Repo](#clone-the-repo)
- [Prerequisites](#prerequisites)

## How to use

Wow such options. Ensure you have met all of the [prerequisites](#prerequisites) first.

### Install with Go

1. Install with Go

```shell
go install github.com/cybergronk/echo-server
```

1. Run server

```shell
# NOTE: You can change the port if you wish
echo-server 8080
```

1. Stop server with `Ctrl+C`

### Run with Docker

> TODO!

### Clone the Repo

1. Clone this repo

```shell
git clone https://github.com/cybergronk/echo-server
```

1. Change directories to the newly cloned repo

```shell
cd echo-server
```

1. Run server

```shell
# NOTE: You can change the port if you wish
go run cmd/main.go 8080
```

1. You can now reach the server with your favorite client

```shell
http://localhost:8080/
```

## Prerequisites

- [Go v1.22](https://go.dev/dl/)

## Roadmap

- [] Unit test routes
- [] CI/CD Pipeline
- [] Docker image
- [] Running with Docker steps
