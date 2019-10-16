#!/bin/bash

cscli query cardservice whois $(cscli keys show alice --address) --chain-id testCardchain

cscli query account $(cscli keys show alice --address) --chain-id testCardchain
