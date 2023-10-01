#!/usr/bin/env bash

Cardchaind tx cardchain createuser $(cat backup/faucetaddress.txt) faucet --from jannik --gas auto