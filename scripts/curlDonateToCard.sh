curl -XPOST \
    http://127.0.0.1:1317/cardservice/donate_to_card \
    -d '{
    "base_req":{
        "from":"cosmos189ujy6ep9pflja9j00tt4pk3q0kqenayf3l7hh",
        "chain_id":"testCardchain",
        "sequence":"18",
        "account_number":"0",
        "gas": "auto",
        "gas_adjustment": "1.5"
      },
      "cardid":"1",
      "amount":"1credits",
      "donator":"cosmos189ujy6ep9pflja9j00tt4pk3q0kqenayf3l7hh"
    }'
