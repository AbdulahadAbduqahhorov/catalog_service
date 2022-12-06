#!/bin/bash
CURRENT_DIR=$(pwd)
for x in $(find ${CURRENT_DIR}/ecommerce-protos/* -type d); do
  sudo protoc --plugin="protoc-gen-go=${GOPATH}/bin/protoc-gen-go" --plugin="protoc-gen-go-grpc=${GOPATH}/bin/protoc-gen-go-grpc" -I=${x} -I=${CURRENT_DIR}/ecommerce-protos -I /usr/local/include --go_out=${CURRENT_DIR}/genproto \
   --go-grpc_out=${CURRENT_DIR}/genproto ${x}/*.proto
done


