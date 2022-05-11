#!/usr/bin/env bash

ARGS="--node tcp://localhost:26659 --gas auto"

Cardchaind keys add $1

Cardchaind tx cardchain createuser $(Cardchaind keys show $1 --address) $1 --from alice $ARGS
