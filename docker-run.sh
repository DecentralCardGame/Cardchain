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

Cardchaind start --consensus.create_empty_blocks false

# the following line is evaluated if csd is terminated via pkill (docker-stop-and-export.sh)
Cardchaind export > ~/.ignite/local-chains/Cardchain/exported_genesis.json
