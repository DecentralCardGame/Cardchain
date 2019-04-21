rm ~/.csd/ -r
csd init --chain-id testCardchain
printf 'asdfasdf\n' | cscli keys delete alice
printf 'asdfasdf\n' | cscli keys delete bob
printf 'asdfasdf\nasdfasdf\n' | cscli keys add alice
printf 'asdfasdf\nasdfasdf\n' | cscli keys add bob
csd add-genesis-account $(cscli keys show alice --address) 1000mycoin,1000credits
csd add-genesis-account $(cscli keys show bob --address) 1000mycoin,1000credits
