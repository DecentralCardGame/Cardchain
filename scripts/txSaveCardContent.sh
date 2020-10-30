#!/bin/bash

printf 'Y\nasdfasdf\n' | cscli tx cardservice save-card-content 1 '{"Place":{"CardName":"Mühle","CastingCost":2,"CostType":{"Energy":false,"Food":false,"Lumber":true,"Iron":false,"Mana":false},"Abilities":[],"Health":3,"FlavourText":"lulul","Tags":["PRIMITIVE"]}}
' 'no image' 'no notes' --from $(cscli keys show alice --address) --chain-id testCardchain
