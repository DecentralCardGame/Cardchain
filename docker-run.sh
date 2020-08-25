#!/bin/bash

if test -f "/var/blockchain/genesis.json"; then
  echo -e "'\033[0;31m' genesis.json found - reloading from there'\033[0m'"
  /go/bin/csd unsafe-reset-all
  cp /var/blockchain/genesis.json /var/blockchain/genesis.json.backup
  mv /var/blockchain/genesis.json ~/.csd/config/
fi

/go/bin/csd start --rpc.laddr tcp://0.0.0.0:26657
# the following line is evaluated if csd is terminated via pkill (docker-stop-and-export.sh)
csd export --for-zero-height > /var/blockchain/genesis.json
