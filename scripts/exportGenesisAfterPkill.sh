#!/bin/bash

pkill -f csd
csd export --for-zero-height > /var/blockchain/genesis.json
