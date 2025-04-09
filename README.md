# openpaths

A CLI tool to extract and display endpoints from an OpenAPI specification.

## Usage

```sh
openpaths [filepath]
```

## example

```sh
$ openpaths tests/01/01.yaml
GET    /pets
POST   /pets
GET    /users
POST   /users
```
