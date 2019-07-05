printf 'Y\nasdfasdf\n' | cscli tx cardservice vote-card 3 underpowered --from $(cscli keys show alice --address) --chain-id testCardchain
