./test_setup.sh

for i in {0..5}
do
./createUser.sh $i
done

Cardchaind tx cardchain create-council 0 --from "Cooler Typ"

for i in {0..5}
do
Cardchaind tx cardchain register-for-council --from $i
done

for i in {0..5}
do
Cardchaind tx cardchain commit-council-response 0 Yes $i --from $i
done

Cardchaind q cardchain q-council 0

for i in {0..5}
do
Cardchaind tx cardchain reveal-council-response 0 Yes $i --from $i
done

Cardchaind q cardchain q-council 0
