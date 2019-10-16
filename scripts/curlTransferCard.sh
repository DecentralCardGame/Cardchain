#!/bin/bash


curl -XPOST \
    http://127.0.0.1:1317/cardservice/transfer_card \
    -d '{
    "base_req":{
        "from":"'$(cscli keys show alice --address)'",
        "chain_id":"testCardchain",
        "gas": "auto",
        "gas_adjustment": "1.5"
      },
      "sender":"'$(cscli keys show alice --address)'",
      "receiver":"'$(cscli keys show bob --address)'",
      "cardid":"1"
    }' > unsignedTx.json
