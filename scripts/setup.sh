#!/bin/bash

rm ~/.csd/ -r
csd init --chain-id testCardchain
printf 'asdfasdf\n' | cscli keys delete alice
printf 'asdfasdf\n' | cscli keys delete bob
printf 'asdfasdf\nasdfasdf\n' | cscli keys add alice &>> keysAliceBob.log
printf 'asdfasdf\nasdfasdf\n' | cscli keys add bob &>> keysAliceBob.log
csd add-genesis-account $(cscli keys show alice --address) 1000mycoin,1000credits
csd add-genesis-account $(cscli keys show bob --address) 1000mycoin,1000credits
