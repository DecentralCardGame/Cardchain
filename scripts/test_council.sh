./test_setup.sh

ARGS="--node tcp://localhost:26659 --gas 8000000"

for i in {0..5}
do
./createUser.sh $i
done

Cardchaind tx cardchain create-council 0 --from "Cooler Typ" $ARGS

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
