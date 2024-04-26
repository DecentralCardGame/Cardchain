yes | ./test_council.sh
yes | cardchaind tx bank send alice $(cardchaind keys show "Cooler Typ" --address) 2000000000000000000000000000000000000000000000000000000ucredits
./import_old_genesis.py ./genesis9mar.json
