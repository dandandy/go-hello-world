# Simple Go Web API Application


This repo contains a simple Web API Application written in Go. Due to the simplicity of the challenge I have decided to not use any frameworks and instead only use the default library.

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
docker build .
```

***

## Environment Variables

The application requires the following environment variables to be set in order for it to start:

- APPLICATION_NAME
- VERSION
- LAST_COMMIT_SHA
- DESCRIPTION

Any valid string is acceptable for these variables.

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
    "name":"foo",
    "version":"1.0.0",
    "lastCommitSha":"abcd1234",
    "description":"Simple Go Web API Application"
}
```

***

## Risks

This application has the following risks:

- If any of the required environment variables are not set it will not start. It is intentionally designed to fail fast, however it instead could be designed to ignore any missing environment variables since usually metadata is not strictly required to run workloads or it could bake in these environment variables into the docker container at build time.

- The health check endpoint doesn't actually do anything. Ideally a health check endpoint would check downstream services are available. This provides an interface to do that, but it doesn't actually check anything.

- The current CI setup does not contain a versioning scheme. This makes it difficult to map code commits to artifact versions when diagnosing issues. Whilst the metadata endpoint somewhat helps in this matter, you still have to go to the trouble to start the application to find the commit sha.