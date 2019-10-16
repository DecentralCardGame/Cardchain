#!/bin/bash

cscli tx cardservice buy-card-scheme 990credits --from $(cscli keys show alice --address) --gas auto --gas-adjustment 1.5 --chain-id testCardchain --generate-only
