#!/bin/bash

set -ex

protoc \
  -I/usr/local/include \
  -I. \
  -I$(GOPATH)/src \
  -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --gofast_out=plugins=grpc:. \
  ./api/admission/v1alpha1/generated.proto
protoc \
  -I/usr/local/include \
  -I. \
  -I$(GOPATH)/src \
  -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --gofast_out=plugins=grpc:. \
  ./api/account/v1alpha1/generated.proto