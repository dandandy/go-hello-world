FROM golang:1.16-alpine AS build

WORKDIR /build
COPY . .

ENV CGO_ENABLED 0
RUN go test -cover ./...
RUN go build

FROM scratch

COPY --from=build /build/go-hello-world /app/go-hello-world
ENTRYPOINT [ "/app/go-hello-world" ]