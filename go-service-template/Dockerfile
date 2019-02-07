FROM golang:1.8 as builder

WORKDIR /go/src/github.com/nchern/go-service-template
COPY . .

# Build the project inside an intermediate container
RUN make install-deps build

# Stat
FROM golang:1.8

WORKDIR /

COPY --from=builder /go/src/github.com/nchern/go-service-template/bin/go-service-template /go-service-template

# Main service port
EXPOSE 80

# (optional) metrics exposing port
EXPOSE 81

ENTRYPOINT /go-service-template
