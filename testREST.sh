#!/bin/sh

ALICEKEY=$(nscli keys show alice --address)
URL="https://localhost:1317/auth/accounts/${ALICEKEY}"
echo "REQ URL: ${URL}"

nohup nscli rest-server --chain-id testCardchain --trust-node &

while [ -z "$RES" ]; do
  sleep 1
  printf 'querying rest-server'
  RES=$(curl -s --insecure ${URL});

  printf '.'
done

printf '\n'

SEQUENCE=$(echo $RES | jq -r  '.value.sequence')
ACCNUM=$(echo $RES | jq -r  '.value.account_number')

echo $SEQUENCE
echo $ACCNUM

REQ='{"base_req":{"name":"alice","from":"alice","password":"asdfasdf","chain_id":"testCardchain","sequence":"'
REQ+=$SEQUENCE
REQ+='","account_number":"'
REQ+=$ACCNUM
REQ+='"},"amount":"990credits","buyer":"cosmos1grwugureenk22gft5nmdtux2a3h5uf0xa64kdk"}'

RES=$(curl -XPOST -s --insecure https://localhost:1317/cardservice/buy_card_scheme --data-binary $REQ)

echo $RES

kill %1
