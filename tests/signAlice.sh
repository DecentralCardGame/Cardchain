printf 'asdfasdf\n' | cscli tx sign unsignedTx.json --from alice --offline --chain-id testCardchain --sequence $(cscli query account $(cscli keys show alice --address) --chain-id testCardchain -o json | jq '.value.sequence | tonumber') --account-number 0 > signedTx.json


# alternate version with sequence by argument (not by query)
#printf 'asdfasdf\n' | cscli tx sign unsignedTx.json --from alice --offline --chain-id testCardchain --sequence $0 --account-number 0 > signedTx.json
