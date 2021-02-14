# API Service Template

## Endpoints

```shell
curl -XGET localhost:8080/v1/ping
curl -XGET localhost:8080 /v1/panic

curl -XPOST /v1/records -d '{"name": "one", "description":"The 1st one"}'
curl -XGET /v1/record/:id
```

## Test

```shell
make test
```

## Build

```shell
make build
```