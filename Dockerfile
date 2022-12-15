FROM golang:1.19-alpine as builder

#
RUN mkdir -p $GOPATH/src/github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/catalog_service 
WORKDIR $GOPATH/src/github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/catalog_service

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && \
    make build && \
    mv ./bin/catalog_service /

FROM alpine
COPY --from=builder catalog_service .
ENTRYPOINT ["/catalog_service"]