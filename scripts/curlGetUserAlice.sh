#!/bin/bash

curl -X GET http://localhost:1317/cardservice/user/$(cscli keys show alice -a) | jq '.'
