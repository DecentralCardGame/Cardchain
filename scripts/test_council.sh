./test_setup.sh

ARGS="--node tcp://localhost:26657 --chain-id Testnet3 --gas 8000000"

for i in {0..5}
do
./createUser.sh $i
done

Cardchaind tx cardchain create-council 1 --from "Cooler Typ" $ARGS

exit

for i in {0..5}
do
Cardchaind tx cardchain register-for-council --from $i $ARGS
done

for i in {0..5}
do
Cardchaind tx cardchain commit-council-response 0 Yes $i "" --from $i $ARGS
done

Cardchaind q cardchain q-council 0 $ARGS

for i in {0..5}
do
Cardchaind tx cardchain reveal-council-response 0 Yes $i --from $i $ARGS
done

Cardchaind q cardchain q-council 0 $ARGS
