
curl -XPOST \
    http://127.0.0.1:1317/cardservice/transfer_card \
    -d '{
    "base_req":{
        "from":"cosmos189ujy6ep9pflja9j00tt4pk3q0kqenayf3l7hh",
        "chain_id":"testCardchain",
        "sequence":"2",
        "account_number":"0",
        "gas": "auto",
        "gas_adjustment": "1.5"
      },
      "sender":"cosmos189ujy6ep9pflja9j00tt4pk3q0kqenayf3l7hh",
      "receiver":"cosmos1a2e728jkct4h45hhcpgz2l8ak9jhw5fkvyyy6k",
      "cardid":"1"
    }'
