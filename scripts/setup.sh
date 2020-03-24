#!/bin/bash

rm ~/.csd/ -r
csd init --chain-id testCardchain
printf 'asdfasdf\n' | cscli keys delete alice
printf 'asdfasdf\n' | cscli keys delete bob
printf 'asdfasdf\nasdfasdf\n' | cscli keys add alice &>> keysAliceBob.log
printf 'asdfasdf\nasdfasdf\n' | cscli keys add bob &>> keysAliceBob.log
csd add-genesis-account $(cscli keys show alice --address) 1000stake,1000credits
csd add-genesis-account $(cscli keys show bob --address) 1000stake,1000credits
csd add-genesis-account cosmos178x4cwg7zuppfgypdd7c0wy0kp304wad9v0awe 1stake,1credits

#sed -i 's/create_empty_blocks = true/create_empty_blocks = false/g' ~/.csd/config/config.toml
