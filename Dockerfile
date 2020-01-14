FROM golang:1.13.6 AS builder
WORKDIR /go/src/platform-test/src
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X 'main.version=$(git describe --always)' -X 'main.commit=$(git rev-parse HEAD)'" -o app main.go app.go handler.go

# final stage
FROM scratch
COPY --from=builder /go/src/platform-test/src/app /go/bin/
EXPOSE 8000
ENTRYPOINT ["/go/bin/app"]