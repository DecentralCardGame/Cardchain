#!/bin/bash

/go/bin/csd start --rpc.laddr tcp://0.0.0.0:26657
csd export --for-zero-height > /var/blockchain/genesis.json
