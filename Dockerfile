FROM golang:1.16-alpine AS build

WORKDIR /build
COPY . .

# TEST
ENV CGO_ENABLED 0
RUN go test -cover ./...

# BUILD
RUN apk add git
RUN go build -ldflags="-X github.com/dandandy/go-hello-world/internal/configuration.lastCommitSha=$(git rev-list -1 HEAD)\
    -X github.com/dandandy/go-hello-world/internal/configuration.version=$(git describe --tags --abbrev=0)"

FROM scratch

COPY --from=build /build/go-hello-world /app/go-hello-world
ENTRYPOINT [ "/app/go-hello-world" ]