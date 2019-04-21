printf 'Y\nasdfasdf\n' | cscli tx cardservice vote-card 1 fairenough --from $(cscli keys show alice --address) --chain-id testCardchain
