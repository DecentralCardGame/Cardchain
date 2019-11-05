#!/bin/bash

curl -XPUT \
  http://127.0.0.1:1317/cardservice/create_user \
  -d '{
  "base_req":{
      "from":"'$(cscli keys show alice --address)'",
      "chain_id":"testCardchain",
      "gas": "auto",
      "gas_adjustment": "1.5"
    },
    "new_user":"'$(cscli keys show alice --address)'",
    "creator":"'$(cscli keys show alice --address)'",
    "alias":"Alice"
  }' > unsignedTx.json
