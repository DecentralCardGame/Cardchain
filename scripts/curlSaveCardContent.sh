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
      "content":"{\"Entity\":{\"Name\":\"Archer\",\"Tags\":[\"RANGE\"],\"Text\":\"Ping\",\"CostType\":{\"Food\":true,\"Lumber\":true,\"Generic\":true},\"CastingCost\":2,\"Abilities\":[{\"ActivatedAbility\":{\"AbilityCost\":1,\"MultipleUse\":true,\"Effects\":[{\"TargetEffect\":{\"EntityTargetEffect\":{\"EntitySelector\":{\"PlayerMode\":\"TARGET\",\"PlayerCondition\":{\"IntCondition\":{\"IntProperty\":\"HANDSIZE\",\"IntComparator\":\"GREATER\",\"IntValue\":1}},\"Zone\":\"FIELD\",\"CardMode\":\"TARGET\",\"CardCondition\":{\"IntCondition\":{\"IntProperty\":\"HEALTH\",\"IntComparator\":\"GREATER\",\"IntValue\":1}}},\"EntityManipulations\":[{\"EntityIntManipulation\":{\"IntProperty\":\"HEALTH\",\"IntOperator\":\"SUBTRACT\",\"IntValue\":1}}]}}}]}}],\"Health\":2,\"Attack\":1}}",
      "notes": "no notes",
      "image": "no image",
      "cardid":"1"
    }' > unsignedTx.json
