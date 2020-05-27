#!/bin/bash

if test -f "/var/blockchain/genesis.json"; then
  /go/bin/csd unsafe-reset-all
  mv /var/blockchain/genesis.json ~/.csd/config/
fi

/go/bin/csd start --rpc.laddr tcp://0.0.0.0:26657
csd export --for-zero-height > /var/blockchain/genesis.json
