# API Service Template

A template for an API service.

## Endpoints

```shell
curl -XGET localhost:8080/v1/ping
curl -XGET localhost:8080 /v1/panic

curl -XPOST /v1/records -d '{"name": "one", "description":"The 1st one"}'
curl -XGET /v1/records/:id
```

## Test

```shell
docker-compose up -d
cd api
make test
```

## Build

```shell
cd api
make build
```