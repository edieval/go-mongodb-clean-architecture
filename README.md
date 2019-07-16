# Go-fasthttp version of the API Challenge

Simple Go server using [fast-http](https://github.com/valyala/fasthttp/) whitch is quite faster than net/http.

For a beautiful code, I use a routing package [fasthttp-routing](https://github.com/qiangxue/fasthttp-routing)

Inspire by https://github.com/irahardianto/service-pattern-go

## Setup

Install the needed dependencies

```sh
go get -u
```

## Run
Run with go installed :

```sh
go run api.go
```

The API will be available here : http://localhost:8080/customers

## Docker

```sh
docker build -t go-fasthttp-api .
docker run -p 8080:8080 go-fasthttp-api
```
