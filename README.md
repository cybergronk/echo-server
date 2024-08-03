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
    echo-server
    ```

    > NOTE: You can change the server's listening port with a PORT environment variable

    ```shell
    # This tells the server to listen on port 8080
    PORT=8080 echo-server
    ```

### Run with Docker

1. Run server in a container

    ```shell
    docker run -p 127.0.0.1:8080:36110 rogvc/echo-server:0.0.1
    ```

    > TIP: You can bind the server's listening port with a port in your host with the `-p` flag

    ```shell
    # This binds the container to port 8080 on your host
    docker run -p 127.0.0.1:8080:36110 rogvc/echo-server:0.0.1
    ```

    > NOTE: You can change the server's listening port with a PORT environment variable

    ```shell
    # This tells the server to listen on port 8080
    docker run -e PORT=8080 rogvc/echo-server:0.0.1
    ```

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
    go run main.go
    ```

1. You can now reach the server with your favorite client

    ```shell
    http://localhost:36110/
    ```

> NOTE: You can change the server's listening port with a PORT environment variable

```shell
# This tells the server to listen on port 8080
PORT=8080 go run main.go
```

## Prerequisites

- [Go v1.22](https://go.dev/dl/)

## Roadmap

- [] Unit test routes
- [] CI/CD Pipeline
- [X] Docker image
- [X] Running with Docker steps
