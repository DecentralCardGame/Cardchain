#!/usr/bin/env bash
parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )

cd "$parent_path"

Cardchaind tx cardchain createuser $(cat ../backup/faucetaddress.txt) faucet --from jannik --gas auto --node tcp://$(cat ../syncnode.txt):26657