#!/bin/bash

# TODO
if false && test -f "/var/blockchain/exported_genesis.json"; then
  echo -e "'\033[0;31m' exported_genesis.json found - reloading from there'\033[0m'"
  mkdir ~/.ignite/local-chains/Cardchain
  cp /var/blockchain/exported_genesis.json ~/.ignite/local-chains/Cardchain/exported_genesis.json
  cp /var/blockchain/binary_checksum.txt ~/.ignite/local-chains/Cardchain/
  cp /var/blockchain/config_checksum.txt ~/.ignite/local-chains/Cardchain/
  cp /var/blockchain/source_checksum.txt ~/.ignite/local-chains/Cardchain/
fi

ignite chain serve

#/go/bin/csd start --rpc.laddr tcp://0.0.0.0:26657
# the following line is evaluated if csd is terminated via pkill (docker-stop-and-export.sh)
#csd export --for-zero-height > /var/blockchain/genesis.json
#.ignite/local-chains/Cardchain ?
