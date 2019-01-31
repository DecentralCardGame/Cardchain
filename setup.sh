rm ~/.nsd/ -r
nsd init --chain-id testCardchain
printf 'asdfasdf\n' | nscli keys delete alice
printf 'asdfasdf\n' | nscli keys delete bob
printf 'asdfasdf\nasdfasdf\n' | nscli keys add alice
printf 'asdfasdf\nasdfasdf\n' | nscli keys add bob
nsd add-genesis-account $(nscli keys show alice --address) 1000mycoin,1000credits
nsd add-genesis-account $(nscli keys show bob --address) 1000mycoin,1000credits
