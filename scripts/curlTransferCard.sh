
curl -XPOST \
    http://127.0.0.1:1317/cardservice/transfer_card \
    -d '{
    "base_req":{
        "from":"'$(cscli keys show alice --address)'",
        "chain_id":"testCardchain",
        "sequence":"2",
        "account_number":"0",
        "gas": "auto",
        "gas_adjustment": "1.5"
      },
      "sender":"'$(cscli keys show alice --address)'",
      "receiver":"'$(cscli keys show bob --address)'",
      "cardid":"1"
    }'
