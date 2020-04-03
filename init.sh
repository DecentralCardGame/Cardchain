#!/usr/bin/env bash

rm -rf ~/.csd
rm -rf ~/.cscli

csd init test --chain-id=testCardchain

cscli config output json
cscli config indent true
cscli config trust-node true
cscli config chain-id testCardchain
cscli config keyring-backend test

cscli keys add alice
cscli keys add bob

csd add-genesis-account $(cscli keys show alice -a) 10000000000000credits,100000000stake
csd add-genesis-account $(cscli keys show bob -a) 10000000000000credits,100000000stake
csd add-genesis-account cosmos178x4cwg7zuppfgypdd7c0wy0kp304wad9v0awe 1stake,1credits

csd gentx --name bob --keyring-backend test

echo "Collecting genesis txs..."
csd collect-gentxs

echo "Validating genesis file..."
csd validate-genesis
