#!/bin/bash

curl -XPUT \
    http://127.0.0.1:1317/cardservice/save_card_content \
    -d '{
    "base_req":{
        "from":"'$(cscli keys show alice --address)'",
        "chain_id":"testCardchain",
        "gas": "auto",
        "gas_adjustment": "1.5"
      },
      "owner":"'$(cscli keys show alice --address)'",
      "content":"this_is_shitcard",
      "cardid":"1"
    }' > unsignedTx.json
