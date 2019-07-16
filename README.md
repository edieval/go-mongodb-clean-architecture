# Go-fasthttp version of the API Challenge

Go http server using [fast-http](https://github.com/valyala/fasthttp/) whitch is quite faster than net/http.

It uses a routing package [fasthttp-routing](https://github.com/qiangxue/fasthttp-routing).

Inspired from https://github.com/irahardianto/service-pattern-go

## Setup
Install the needed dependencies

```sh
go get -u
```

## Build
Run with go installed :

```sh
go build -o serve
```

## Run
Run:

```sh
./serve
```

The API will be available here : http://localhost:8080/category/:code

## Docker

```sh
docker build -t go-mongodb-clean-architecture .
docker run -p 8080:8080 go-mongodb-clean-architecture
```
