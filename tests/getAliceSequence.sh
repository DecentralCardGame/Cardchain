cscli query account $(cscli keys show alice --address) --chain-id testCardchain -o json | jq '.value.sequence | tonumber'
