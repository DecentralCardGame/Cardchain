printf 'Y\nasdfasdf\n' | cscli tx cardservice transfer-card 1 $(cscli keys show bob --address) --from $(cscli keys show alice --address) --chain-id testCardchain
