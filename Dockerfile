FROM scratch

# ARG SERVICE_NAME=${SERVICE_NAME:go-service-template}

COPY bin/go-service-template /go-service-template

WORKDIR /

# Main service port
EXPOSE 80

# (optional) metrics exposing port
EXPOSE 81

ENTRYPOINT ["./go-service-template"]

# Maybe better approach: build inside the container

# FROM golang:1.8 as builder
# WORKDIR /go/src/<PATH-TO-PROJECT>
# RUN curl https://glide.sh/get | sh
# WORKDIR /go/src/<PATH-TO-PROJECT>
# RUN make build

# FROM golang:1.8
# COPY --from=builder /go/src/<PATH-TO-PROJECT>/bin/$SERVICE_NAME .
# ENTRYPOINT ["./$SERVICE_NAME"]
