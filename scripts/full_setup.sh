yes | ./test_council.sh
yes | Cardchaind tx bank send alice $(Cardchaind keys show "Cooler Typ" --address) 2000000000000000000000000000000000000000000000000000000ucredits
./import_old_genesis.py ./genesis9mar.json
