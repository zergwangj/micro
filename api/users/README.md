# Users Service

This is the Users service

Generated with

```
micro new github.com/zergwangj/micro/api/users --namespace=zergwangj --alias=users --type=api
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: zergwangj.api.users
- Type: api
- Alias: users

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
./users-api
```

Build a docker image
```
make docker
```