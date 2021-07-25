# Simple Go Web API Application


This repo contains a simple Web API Application written in Go. I decided to not use any frameworks and instead only use the default library since I felt that frameworks are overkill for these endpoints.

This project was built using Go 1.16 but should work with most versions.

***

## Requirements for build and test the application

The following tools should be installed on the system

- Go 1.16
- Docker

***

## How to build, test and run the application

The following shell command is used to build the project:

```bash
go build
```

The following shell command is used to run the unit test suite with unit test coverage:

```bash
go test -cover ./...
```

The following shell command is used to run the application binary:

```bash
./go-hello-world
```

Alternatively to run the application without building it first:

```bash
go run main.go
```

All of the above commands are run in multistage the Dockerfile. To build the Dockerfile:

```bash
docker build . -t go-hello-world
```

Then to run it:

```bash
docker run -p 8080:8080 go-hello-world
```

***

## Endpoints

This application supports the following endpoints.

### Hello World: GET /

```text
HTTP Status code: 200
Body: "hello world"
```

### Health Check: GET /health

```text
HTTP Status code: 200
```

Body:

```json
{
    "healthy":true,
    "dependencies":[]
}
```

### Metadata: GET /metadata

```text
HTTP Status code: 200
```

Body:

```json
{
    "name":"go-hello-world",
    "version":"v1.0.0",
    "lastCommitSha":"dfc20f365734a26dfa9afa7a90edd92a8ec88a57",
    "description":"A simple Go Web API application"
}
```

***

## Risks

This application has the following risks:

- The health check endpoint doesn't actually do anything. Ideally a health check endpoint would check downstream services are available. This provides an interface to do that, but it doesn't actually check anything.

- The current CI setup does not contain an automatic versioning scheme, instead it relies on developers to push a git tag. This means that developers may forget to update the versions.

- From a security standpoint, someone could tamper with the container by swapping out the binary for a mallicious one. To detect if this has occured could take a HASH of the final image at build time to give us the ability to detect if it has been tampered with.