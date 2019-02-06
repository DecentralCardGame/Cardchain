printf 'asdfasdf\n' | nscli tx cardservice save-card-content 0 this_is_shitcard --from $(nscli keys show alice --address) --chain-id testCardchain
