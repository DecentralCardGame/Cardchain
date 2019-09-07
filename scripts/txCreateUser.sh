printf 'y\nasdfasdf\n' | cscli tx cardservice create-user $(cscli keys show alice --address) alice --from $(cscli keys show alice --address) --chain-id testCardchain
