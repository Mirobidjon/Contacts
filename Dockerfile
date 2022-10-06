# workspace (GOPATH) configured at /go
FROM golang:1.16 as builder

#
RUN mkdir -p $GOPATH/src/github.com/Mirobidjon/contact-list
WORKDIR $GOPATH/src/github.com/Mirobidjon/contact-list

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && \
    make build && \
    mv ./bin/contact /

FROM alpine
COPY --from=builder contact .
RUN mkdir config
ENTRYPOINT ["/contact"]
