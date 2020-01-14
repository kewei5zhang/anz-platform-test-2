# ANZ Technical Test v2 - Test 2

The following test will require you to do the following:

- Create a simple application which has a single "/version" endpoint.
- Containerise your application as a single deployable artefact, encapsulating all dependencies. 
- Create a CI pipeline for your application

## Prerequisites

- `make` - 4.1+
- [docker](https://docs.docker.com/install/) - 18.09.0+

## Building Image
```
make go-build
```
## Running tests
```
make go-test
```
## Deployment
```
make go-deploy
```
To verify the application version api service on local machine, hit the following GET requests in Postman.
```
http://localhost:8080/
```
```
http://localhost:8080/version
```
## Summary

-----
This simple Restful API will return application version metadata in the following JSON format.
```
[
    {
        "version":"0.1",
        "lastCommitSha":"d9012fe47f2e80d25a408a7baa6c28b390d72ee0",
        "description":"Release for ANZ Platform Test"
    }
]
```
- *version* returns the latest git tag
- *lastCommitSha* returns the commit id

*Note that version will return abbreviated commit id if there is no availabe git tags*

This app is developed using GoLang with gorilla/mux package for requrest router implementation. All unit testing are addressed using GoLang testing framework. All CI pipelines are developed using Makefile with each build, test and deployment step compiled witin docker containers.

### Application / Deployment Versioning
As shown in the *RUN* step in the *builder* stage in the docker file below, all application and deployment verioning are managed by git tagging. 
```
FROM golang:1.13.6 AS builder
WORKDIR /go/src/platform-test/src
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X 'main.version=$(git describe --always)' -X 'main.commit=$(git rev-parse HEAD)'" -o app main.go app.go handler.go
```
There are several advantages with this approach.
1. Centralized app versioning control associated with SCM tool
2. Applications do not need to maintain a seperate file of app version and other environment variables and pass it through docker images during testing and building stage.
3. Empowering version control and release management with additional capabilities provided by a wide range of products that can be intergrated with SCM tools, e.g. JIRA, Logging Services, Jenkins with Git-webhook

### Limitations and Risks
- The deployment step is currently binding the container port *8000* with the host server port *8080*. Open port *8080* on the host server is required by this web services.
- Although the final stage in the Dockerfile uses the *scratch* base image which size is fairly small, the size of *golang:1.13.6* image in builder stage is quite large. Should consider using *golang:alpine* image instead in future interations.
- A CICD pipeline should be developed and tested with github integration before production deployed for better change monitoring, version control and enforcing branching strategy.
