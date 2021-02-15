#!/usr/bin/env bash

set -eo pipefail

PROJECT_PROTO_DIR=proto/cns
COSMOS_SDK_DIR=${COSMOS_SDK_DIR:-$(go list -f "{{ .Dir }}" -m github.com/cosmos/cosmos-sdk)}

# Generate Go types from protobuf
protoc \
  -I=. \
  -I="$COSMOS_SDK_DIR/third_party/proto" \
  -I="$COSMOS_SDK_DIR/proto" \
  --gocosmos_out=plugins=interfacetype+grpc,\
Mgoogle/protobuf/any.proto=github.com/cosmos/cosmos-sdk/codec/types:. \
  --grpc-gateway_out=logtostderr=true:. \
  --doc_out=./docs \
  --doc_opt=markdown,proto.md \
  $(find "${PROJECT_PROTO_DIR}" -maxdepth 1 -name '*.proto')

# move proto files to the right places
cp -r github.com/tendermint/cns/* ./
rm -rf github.com

