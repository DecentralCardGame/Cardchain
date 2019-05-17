printf 'Y\nasdfasdf\n' | cscli tx cardservice save-card-content 1 "this_is_shitcard" --from $(cscli keys show alice --address) --chain-id testCardchain
