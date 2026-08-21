#!/bin/bash

# Stop on errors
set -e

# Paths
PROTO_DIR="api/todo-list"
OUT_DIR="pkg/pb"

# Create out folders if not exist
mkdir -p $OUT_DIR/user

echo "Generating gRPC code from proto files..."

protoc --proto_path=$PROTO_DIR \
       --go_out=$OUT_DIR/user --go_opt=paths=source_relative \
       --go-grpc_out=$OUT_DIR/user --go-grpc_opt=paths=source_relative \
       $PROTO_DIR/user.proto

go mod tidy

echo "Done! Code successfully generated in $OUT_DIR"