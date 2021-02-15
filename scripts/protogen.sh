#!/usr/bin/env bash

set -eo pipefail

COSMOS_SDK_DIR=${COSMOS_SDK_DIR:-$(go list -f "{{ .Dir }}" -m github.com/cosmos/cosmos-sdk)}

proto_dirs=$(find ./proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
   protoc \
  -I "proto" \
  -I="$COSMOS_SDK_DIR/third_party/proto" \
  -I="$COSMOS_SDK_DIR/proto" \
  --gocosmos_out=plugins=interfacetype+grpc,\
Mgoogle/protobuf/any.proto=github.com/cosmos/cosmos-sdk/codec/types:. \
  --doc_out=./docs \
  --doc_opt=markdown,proto.md \
  $(find "${dir}" -maxdepth 1 -name '*.proto')

  # command to generate gRPC gateway (*.pb.gw.go in respective modules) files
  protoc \
  -I "proto" \
  -I="$COSMOS_SDK_DIR/third_party/proto" \
  -I="$COSMOS_SDK_DIR/proto" \
  --grpc-gateway_out=logtostderr=true:. \
  $(find "${dir}" -maxdepth 1 -name '*.proto')
done

# move proto files to the right places
cp -r github.com/tendermint/cns/* ./
rm -rf github.com