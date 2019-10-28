#!/bin/bash

cscli query cardservice user $(cscli keys show alice --address) --chain-id testCardchain

cscli query account $(cscli keys show alice --address) --chain-id testCardchain
