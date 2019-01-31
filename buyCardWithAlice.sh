printf 'asdfasdf\n' | nscli tx cardservice buy-card-scheme 5credits --from $(nscli keys show alice --address) --chain-id testCardchain
