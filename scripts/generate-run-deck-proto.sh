#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/.."
GO=${GO:-go}
tmp=$(mktemp -d)
"$GO" build -o "$tmp/protoc-gen-gocosmos" github.com/cosmos/gogoproto/protoc-gen-gocosmos
"$GO" build -o "$tmp/protoc-gen-grpc-gateway" github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
moddir() { "$GO" list -m -f '{{.Dir}}' "$1"; }
includes=(-I proto -I "$(moddir github.com/cosmos/cosmos-sdk)/proto" -I "$(moddir github.com/cosmos/gogoproto)" -I "$(moddir github.com/cosmos/cosmos-proto)/proto" -I "$(moddir github.com/grpc-ecosystem/grpc-gateway)/third_party/googleapis")
PATH="$tmp:$PATH" protoc "${includes[@]}" --gocosmos_out="plugins=grpc:$tmp" cardchain/cardchain/run_deck.proto cardchain/cardchain/genesis.proto
PATH="$tmp:$PATH" protoc "${includes[@]}" --grpc-gateway_out="logtostderr=true:$tmp" cardchain/cardchain/run_deck.proto
out="$tmp/github.com/DecentralCardGame/cardchain/x/cardchain/types"
cp "$out/run_deck.pb.go" "$out/genesis.pb.go" "$out/run_deck.pb.gw.go" x/cardchain/types/
printf 'Generated run deck types, services and gateway; staging: %s\n' "$tmp"
