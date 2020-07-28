# Golang es_ES service

## env vars

```
export PORT=8001 
```

## install requirements

```bash
$ go build ./...
```

## Docker

```bash
$ docker build -t say-what:v1.0 .
$ docker run --rm -d -p 8001:8001/tcp say-what:v1.0
```

## run test

```bash
$ go test -v
```

## run service

```
$ go run main.go
```

## Authors

* Maximo Miranda <maximo.miranda@cuemby.com>