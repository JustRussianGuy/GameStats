#!/bin/bash
set -e

PROTO_DIR="./api/gamestats_api"
OUT_DIR="./internal/pb/gamestats_api"

mkdir -p $OUT_DIR

protoc \
  -I $PROTO_DIR \
  --go_out=$OUT_DIR --go_opt=paths=source_relative \
  --go-grpc_out=$OUT_DIR --go-grpc_opt=paths=source_relative \
  $PROTO_DIR/gamestats.proto
