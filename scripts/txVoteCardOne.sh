#!/bin/bash

printf 'Y\nasdfasdf\n' | cscli tx cardservice vote-card 1 underpowered --from $(cscli keys show alice --address) --chain-id testCardchain
