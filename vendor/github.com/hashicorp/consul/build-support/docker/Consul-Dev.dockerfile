ARG CONSUL_BUILD_IMAGE
FROM ${CONSUL_BUILD_IMAGE}:latest as builder
# FROM golang:latest as builder
ARG GIT_COMMIT
ARG GIT_DIRTY
ARG GIT_DESCRIBE
# WORKDIR /go/src/github.com/hashicorp/consul
ENV CONSUL_DEV=1
ENV COLORIZE=0
Add . /go/src/github.com/hashicorp/consul/
RUN make dev


FROM consul:latest

COPY --from=builder /go/src/github.com/hashicorp/consul/bin/consul /bin
