curl -XPUT \
    http://127.0.0.1:1317/cardservice/vote_card \
    -d '{
    "base_req":{
        "from":"'$(cscli keys show alice --address)'",
        "chain_id":"testCardchain",
        "sequence":"18",
        "account_number":"0",
        "gas": "auto",
        "gas_adjustment": "1.5"
      },
      "voter":"'$(cscli keys show alice --address)'",
      "votetype":"fair_enough",
      "cardid":"1"
    }' > unsignedTx.json
