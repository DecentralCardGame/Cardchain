#!/bin/bash

printf 'asdfasdf\n' | cscli tx sign unsignedTx.json --from alice --offline --chain-id testCardchain --sequence $(cscli query account $(cscli keys show alice --address) --chain-id testCardchain -o json | jq '.value.sequence | tonumber') --account-number 0 > signedTx.json
