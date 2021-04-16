#!/bin/bash

printf 'Y\nasdfasdf\n' | cscli tx cardservice save-card-content 1 '{
  "Entity":{
    "CardName":"Name",
    "Tags":["SPIRITUAL"],
    "FlavourText":"-.-",
    "Class":{
      "Technology":true,
      "Nature":false,
      "Mysticism":false,
      "Culture":false,
      "Mana":false
    },
    "CastingCost":13,
    "Attack":10,
    "Health":10,
    "Abilities":[
      {
        "Pay":{
          "ManaAmount":2,
          "Effects":[
            {
              "Arm": {
                "Amount":2
              }
            }
          ]
        }
      }
    ]
  }
}' 'no image' 'no notes' --from $(cscli keys show alice --address) --chain-id testCardchain
