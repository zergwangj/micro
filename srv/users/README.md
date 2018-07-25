# Users Service

This is the Users service

Generated with

```
micro new github.com/zergwangj/micro/users --namespace=zergwangj --alias=users --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: zergwangj.srv.users
- Type: srv
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
./users-srv
```

Build a docker image
```
make docker
```