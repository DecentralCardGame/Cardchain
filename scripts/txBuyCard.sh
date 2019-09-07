printf 'y\nasdfasdf\n' | cscli tx cardservice buy-card-scheme 990credits --from $(cscli keys show alice --address) --chain-id testCardchain
