FROM scratch

# ARG SERVICE_NAME=${SERVICE_NAME:-my-cool-service}

COPY bin/my-cool-service /my-cool-service

WORKDIR /

ENTRYPOINT ["./my-cool-service"]

# Maybe better approach: build inside the container

# FROM golang:1.8 as builder
# WORKDIR /go/src/<PATH-TO-PROJECT>
# RUN curl https://glide.sh/get | sh
# WORKDIR /go/src/<PATH-TO-PROJECT>
# RUN make build

# FROM golang:1.8
# COPY --from=builder /go/src/<PATH-TO-PROJECT>/bin/$SERVICE_NAME .
# ENTRYPOINT ["./$SERVICE_NAME"]
