#!/usr/bin/bash

ARGS="--node tcp://localhost:26659"

Cardchaind keys add "Cooler Typ"

Cardchaind keys add "Cooler Artist"

Cardchaind tx cardchain createuser $(Cardchaind keys show "Cooler Typ" --address) "Cooler Typ" --from alice $ARGS

Cardchaind tx cardchain createuser $(Cardchaind keys show "Cooler Artist" --address) "Cooler Artist" --from alice $ARGS

Cardchaind tx cardchain buy-card-scheme 8000000000ucredits --from "Cooler Typ" $ARGS

Cardchaind tx cardchain save-card-content 0 '{"Entity":{"CardName":"Yees","CastingCost":13,"Class":{"Nature":false,"Mysticism":false,"Technology":true,"Culture":true},"Abilities":[{"Pay":{"ManaAmount":2,"Effects":[{"Arm":{"Target":"RANDOM","Amount":{"SimpleIntValue":3}}}]}}],"Attack":10,"Health":10,"FlavourText":"-.-","Tags":["SPIRIT"],"Keywords":["[\"Periodic\", \"Ravage\"]"],"RulesTexts":["This could be your RulesText",""]}}' "no notes" "$(Cardchaind keys show "Cooler Artist" --address)" --from "Cooler Typ" --gas 99000000 $ARGS

Cardchaind tx cardchain add-artwork 0 "data:image/jpeg;base64,$(base64 ./pic.jpg)" true --from "Cooler Artist" --gas 990000000 $ARGS
