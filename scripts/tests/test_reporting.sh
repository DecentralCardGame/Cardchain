#./createUser.sh player1
#./createUser.sh player2
#./createUser.sh reporter1

NUM=43

Cardchaind tx cardchain open-match `Cardchaind keys show player1 --address` `Cardchaind keys show player2 --address` "[1,2,3,7,8,624]" "[4,5,6,10,11,12,672]" --from reporter1 --gas auto

Cardchaind tx cardchain confirm-match $NUM BWon "[{\"cardId\": 672, \"voteType\": 3}]" --from player1 --gas auto
Cardchaind tx cardchain confirm-match $NUM BWon "[{\"cardId\": 624, \"voteType\": 2}]" --from player2 --gas auto
Cardchaind tx cardchain report-match $NUM "[2,3,8,624]" "[4,5,6,12,672]" BWon --from reporter1 --gas auto