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
      "content":"{\"Entity\":{\"CardName\":\"A\u00FChle\",\"CastingCost\":1,\"CostType\":{\"Energy\":false,\"Food\":false,\"Lumber\":true,\"Iron\":false,\"Mana\":false},\"Abilities\":[],\"Health\":3,\"FlavourText\":\"lulul\",\"Tags\":[\"PRIMITIVE\"]}}",
      "notes": "no notes",
      "image": "no image",
      "cardid":"4"
    }' > unsignedTx.json
