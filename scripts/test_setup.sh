#!/usr/bin/bash

Cardchaind keys add "Cooler Typ"

Cardchaind keys add "Cooler Artist"

Cardchaind tx cardchain createuser $(Cardchaind keys show "Cooler Typ" --address) "Cooler Typ" --from alice

Cardchaind tx cardchain createuser $(Cardchaind keys show "Cooler Artist" --address) "Cooler Artist" --from alice

Cardchaind tx cardchain buy-card-scheme 8000000000 --from "Cooler Typ"

Cardchaind tx cardchain save-card-content 0 '{"Entity":{"CardName":"Name","CastingCost":13,"Class":{"Nature":false,"Mysticism":false,"Technology":true,"Culture":true},"Abilities":[{"Pay":{"ManaAmount":2,"Effects":[{"Arm":{"Target":"RANDOM","Amount":{"SimpleIntValue":3}}}]}}],"Attack":10,"Health":10,"FlavourText":"-.-","Tags":["SPIRIT"],"Keywords":["ARM","PAY"],"RulesTexts":["This could be your RulesText",""]}}' "no notes" "$(Cardchaind keys show "Cooler Artist" --address)" --from "Cooler Typ"
