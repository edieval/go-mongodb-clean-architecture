FROM golang:1.12.6 as builder

COPY  /app/api.go
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum

WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=readonly -ldflags="-w -s" -o /app/main

FROM scratch

COPY --from=builder /app/main /go/bin/main

EXPOSE 8080

CMD ["/go/bin/main"]
