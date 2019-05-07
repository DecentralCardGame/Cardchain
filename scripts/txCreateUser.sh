printf 'Y\nasdfasdf\n' | cscli tx cardservice create-user $(cscli keys show alice --address) --from $(cscli keys show alice --address) --chain-id testCardchain
