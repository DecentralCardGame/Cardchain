#!/bin/bash

cscli query cardservice votable-cards $(cscli keys show alice --address) --chain-id testCardchain
