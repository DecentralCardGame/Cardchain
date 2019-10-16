#!/bin/bash

curl -X GET http://127.0.0.1:1317/cardservice/votable_cards/$(cscli keys show alice -a) | jq .
