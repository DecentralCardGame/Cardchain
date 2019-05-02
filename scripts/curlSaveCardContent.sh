curl -XPUT \
    http://127.0.0.1:1317/cardservice/save_card_content \
    -d '{
    "base_req":{
        "from":"cosmos189ujy6ep9pflja9j00tt4pk3q0kqenayf3l7hh",
        "chain_id":"testCardchain",
        "sequence":"2",
        "account_number":"0",
        "gas": "auto",
        "gas_adjustment": "1.5"
      },
      "owner":"cosmos189ujy6ep9pflja9j00tt4pk3q0kqenayf3l7hh",
      "content":"this_is_shitcard",
      "cardid":"1"
    }'
