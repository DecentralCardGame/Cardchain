rm ~/.csd/ -r
csd init prawda --chain-id testCardchain
printf 'asdfasdf\n' | cscli keys delete alice
printf 'asdfasdf\n' | cscli keys delete bob
printf 'asdfasdf\nasdfasdf\n' | cscli keys add alice
printf 'asdfasdf\nasdfasdf\n' | cscli keys add bob
csd add-genesis-account $(cscli keys show alice --address) 1000credits,100000000stake
csd add-genesis-account $(cscli keys show bob --address) 1000credits,100000000stake

cscli config chain-id testCardchain
cscli config output json
cscli config indent true
cscli config trust-node true

printf 'asdfasdf\n' | csd gentx --name alice

csd collect-gentxs
csd validate-genesis
