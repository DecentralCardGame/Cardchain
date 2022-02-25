#!/usr/bin/env bash

Cardchaind keys add $1

Cardchaind tx cardchain createuser $(Cardchaind keys show $1 --address) $1 --from alice
