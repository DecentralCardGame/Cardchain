#!/bin/bash

printf 'Y\nasdfasdf\n' | cscli tx cardservice save-card-content 1 '{"Field":{"Name":"Mühle","Tag":"primitiv","Text":"Brot für alle","Cost":["LUMBER","LUMBER","GENERIC"],"CastSpeed":1,"Abilities":[{"TriggeredAbility":{"Effect":{"Production":["FOOD"],"Manipulation":[],"ZoneChange":[]},"Cause":{"TimeEventListener":{"TimeEvent":"TICKSTART"}}}}],"AbilitySpeed":1,"Health":2}}' --from $(cscli keys show alice --address) --chain-id testCardchain
