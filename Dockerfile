FROM golang:1.18 AS build
WORKDIR /go/src/github.com/VATUSA/api-v3/
COPY go.mod ./
COPY go.sum ./
COPY cmd ./cmd
COPY internal ./internal
COPY pkg ./pkg
RUN go build -o bin/conversion ./cmd/conversion/main.go
RUN go build -o bin/api ./cmd/api/main.go

FROM alpine:latest AS conversion
WORKDIR /app
COPY --from=build /go/src/github.com/VATUSA/api-v3/bin/conversion ./
CMD ["./conversion"]

FROM alpine:latest AS api
WORKDIR /app
COPY --from=build /go/src/github.com/VATUSA/api-v3/bin/api ./
CMD ["./api"]
