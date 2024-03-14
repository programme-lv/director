#! /usr/bin/env bash

set -ex # exit on first error, print commands

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    msg/director.proto