curl -XPOST \
  http://127.0.0.1:1317/cardservice/buy_card_scheme \
  -d '{
  "base_req":{
      "from":"cosmos1prl2q05w8ua6jxk7m6tzkyq2ljjl7y7adzhqrg",
      "chain_id":"testCardchain",
      "sequence":"0",
      "account_number":"0",
      "gas": "auto",
      "gas_adjustment": "1.5"
    },
    "amount":"991credits",
    "buyer":"cosmos1prl2q05w8ua6jxk7m6tzkyq2ljjl7y7adzhqrg"
  }'
