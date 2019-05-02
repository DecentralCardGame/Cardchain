
curl -XPOST \
  http://127.0.0.1:1317/cardservice/buy_card_scheme \
  -d '{
  "base_req":{
      "from":"cosmos189ujy6ep9pflja9j00tt4pk3q0kqenayf3l7hh",
      "password":"asdfasdf",
      "chain_id":"testCardchain",
      "sequence":"0",
      "account_number":"0",
      "gas": "auto",
      "gas_adjustment": "1.5"
    },
    "amount":"801credits",
    "buyer":"cosmos189ujy6ep9pflja9j00tt4pk3q0kqenayf3l7hh"
  }'
