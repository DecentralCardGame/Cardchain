#!/usr/bin/bash

Cardchaind keys add "Cooler Typ"

Cardchaind tx cardchain createuser $(Cardchaind keys show "Cooler Typ" --address) "Cooler Typ" --from alice

Cardchaind tx cardchain buy-card-scheme 8000credits $(Cardchaind keys show "Cooler Typ" --address) --from alice

Cardchaind tx cardchain save-card-content "1" '{"Entity":{"CardName":"Name","CastingCost":13,"Class":{"Nature":false,"Mysticism":false,"Technology":true,"Culture":true},"Abilities":[{"Pay":{"ManaAmount":2,"Effects":[{"Arm":{"Target":"RANDOM","Amount":{"SimpleIntValue":3}}}]}}],"Attack":10,"Health":10,"FlavourText":"-.-","Tags":["SPIRIT"],"Keywords":["ARM","PAY"],"RulesTexts":["This could be your RulesText",""]}}' "no image" "no art" "no notes" "$(Cardchaind keys show "Cooler Typ" --address)" --from alice

