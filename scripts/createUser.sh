#!/usr/bin/env bash

ARGS="--node tcp://localhost:26657 --gas auto"

cardchaind keys add $1

cardchaind tx cardchain createuser $(cardchaind keys show $1 --address) $1 --from alice $ARGS
