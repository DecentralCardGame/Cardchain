printf 'Y\nasdfasdf\n' | cscli tx cardservice donate-to-card 1 10credits --from $(cscli keys show alice --address) --chain-id testCardchain
