# kin-starter-go

## About

This project is a simple demo of how to use the [kin-go](https://github.com/kinecosystem/kin-go) SDK.

## Requirements

- Basic Go knowledge
- Go 1.14 or higher
- Docker (optional)

## Running this project

### 1. Clone the repo

```shell
git clone https://github.com/kintegrate/kin-starter-go.git
cd kin-starter-go
```

### 2. Install the dependencies

```shell
go install
```

### 3. Run the demo

```shell
go run main.go
```

## Docker

You can also run this project inside a Docker container:

```shell
make docker-build
make docker-run
```

## What's next?

You can read the [Getting Started - Go](https://kintegrate.dev/tutorials/getting-started-go-sdk) to read how you can integrate the `kin-go` SDK in your own apps.

If you have questions or want to talk about how to integrate Kin, please join our [discord channel](https://discord.gg/kdRyUNmHDn).
