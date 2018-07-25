# Auth Service

This is the Auth service

Generated with

```
micro new github.com/zergwangj/micro/auth --namespace=zergwangj --alias=auth --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: zergwangj.srv.auth
- Type: srv
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
./auth-srv
```

Build a docker image
```
make docker
```