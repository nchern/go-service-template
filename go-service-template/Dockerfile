FROM golang:1.13 as builder

WORKDIR /go/src/github.com/nchern/go-service-template/go-service-template
COPY . .

# Build the project inside an intermediate container
RUN make install-deps install

FROM golang:1.13

WORKDIR /

COPY --from=builder /go/bin/go-service-template /go-service-template

# Main service port
EXPOSE 80

# (optional) metrics exposing port
EXPOSE 81

ENTRYPOINT /go-service-template
