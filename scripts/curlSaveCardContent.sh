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
      "content": "{\r\n  \"Entity\":{\r\n    \"CardName\":\"Name\",\r\n    \"Tags\":[\"SPIRITUAL\"],\r\n    \"FlavourText\":\"-.-\",\r\n    \"CostType\":{\r\n      \"Lumber\":true,\r\n      \"Energy\":false,\r\n      \"Food\":false,\r\n      \"Iron\":false,\r\n      \"Mana\":false\r\n    },\r\n    \"CastingCost\":13,\r\n    \"Attack\":10,\r\n    \"Health\":10,\r\n    \"Abilities\":[\r\n      {\r\n        \"Pay\":{\r\n          \"RessourceAmount\":2,\r\n          \"Effects\":[\r\n            {\r\n              \"Arm\": {\r\n                \"Amount\":2\r\n              }\r\n            }\r\n          ]\r\n        }\r\n      }\r\n    ]\r\n  }\r\n}",
      "notes": "no notes",
      "image": "no image",
      "cardid":"1"
    }' > unsignedTx.json
