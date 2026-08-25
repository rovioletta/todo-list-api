#!/bin/bash

# Stop on errors
set -e

# Paths
PROTO_DIR="api/todo-list"
OUT_DIR="pkg/pb"
VENDOR_DIR=".proto-deps"

# Create out folders if not exist
mkdir -p $OUT_DIR/user
mkdir -p $OUT_DIR/task

# Install buf/validate dependencies
if [ ! -d "$VENDOR_DIR/buf/validate" ]; then
    echo "Downloading buf/validate definitions..."
    mkdir -p $VENDOR_DIR/buf
    git clone --depth 1 --sparse https://github.com/bufbuild/protovalidate.git $VENDOR_DIR/protovalidate_tmp
    cd $VENDOR_DIR/protovalidate_tmp
    git sparse-checkout set proto/protovalidate/buf/validate
    cd ../..
    mv $VENDOR_DIR/protovalidate_tmp/proto/protovalidate/buf/validate $VENDOR_DIR/buf/
    rm -rf $VENDOR_DIR/protovalidate_tmp
fi

echo "Generating gRPC code from proto files..."

protoc --proto_path=$PROTO_DIR \
       --proto_path=$VENDOR_DIR \
       --go_out=$OUT_DIR/user --go_opt=paths=source_relative \
       --go-grpc_out=$OUT_DIR/user --go-grpc_opt=paths=source_relative \
       $PROTO_DIR/user.proto

protoc --proto_path=$PROTO_DIR \
       --proto_path=$VENDOR_DIR \
       --go_out=$OUT_DIR/task --go_opt=paths=source_relative \
       --go-grpc_out=$OUT_DIR/task --go-grpc_opt=paths=source_relative \
       $PROTO_DIR/task.proto

go mod tidy

echo "Done! Code successfully generated in $OUT_DIR"