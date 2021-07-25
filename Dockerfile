FROM golang:1.16-alpine AS build

WORKDIR /build
COPY . .
# Required for golangci-lint and go test
ENV CGO_ENABLED 0

# LINT
COPY --from=golangci/golangci-lint /usr/bin/golangci-lint /usr/bin/golangci-lint
RUN /usr/bin/golangci-lint run

# TEST
RUN go test -cover ./...

# BUILD
RUN apk add git
RUN go build -ldflags="-X github.com/dandandy/go-hello-world/internal/configuration.lastCommitSha=$(git rev-list -1 HEAD)\
    -X github.com/dandandy/go-hello-world/internal/configuration.version=$(git describe --tags --abbrev=0)"

FROM scratch

COPY --from=build /build/go-hello-world /app/go-hello-world
ENTRYPOINT [ "/app/go-hello-world" ]