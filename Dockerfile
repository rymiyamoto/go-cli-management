ARG GO_VERSION=1.21.4

FROM golang:${GO_VERSION} AS dev
RUN go install github.com/spf13/cobra-cli@v1.3.0
