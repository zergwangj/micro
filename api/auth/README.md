# Auth Service

This is the Auth service

Generated with

```
micro new github.com/zergwangj/micro/api/auth --namespace=zergwangj --alias=auth --type=api
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: zergwangj.api.auth
- Type: api
- Alias: auth

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./auth-api
```

Build a docker image
```
make docker
```