
curl -XPOST \
  http://127.0.0.1:1317/cardservice/buy_card_scheme \
  -d '{
  "base_req":{
      "from":"'$(cscli keys show alice --address)'",
      "password":"asdfasdf",
      "chain_id":"testCardchain",
      "sequence":"0",
      "account_number":"0",
      "gas": "auto",
      "gas_adjustment": "1.5"
    },
    "amount":"800credits",
    "buyer":"'$(cscli keys show alice --address)'"
  }' > unsignedTx.json
